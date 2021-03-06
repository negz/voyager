package svccatadmission

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/atlassian/voyager/pkg/k8s"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	serviceInstanceName = "foo"
	defaultNamespace    = "somenamespace"
)

func TestMicrosAdmitFunc(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	scStore := setupSCMock()

	tests := []struct {
		name            string
		admissionReview admissionv1beta1.AdmissionReview
		want            *admissionv1beta1.AdmissionResponse
		wantErr         bool
	}{
		{
			"NotServiceInstance",
			buildAdmissionReview(defaultNamespace, k8s.ServiceBindingGVR, admissionv1beta1.Create, []byte(`{}`)),
			nil,
			true,
		},
		{
			"ServiceInstanceNotMicros",
			buildAdmissionReview(dougMicros2Service, k8s.ServiceInstanceGVR, admissionv1beta1.Create, buildServiceInstance(t, "serviceid", "planid", nil)),
			buildAdmissionResponse(true, 0, metav1.StatusReasonUnknown, nil, `ServiceInstance "foo" is not a micros compute instance`),
			false,
		},
		{
			"ErrorIfNoNamespace",
			buildAdmissionReview("", k8s.ServiceInstanceGVR, admissionv1beta1.Create, buildV1ServiceInstance(t, missingService)),
			nil,
			true,
		},
		{
			"ServiceMissingIsOk",
			buildAdmissionReview(dougMicros2Service, k8s.ServiceInstanceGVR, admissionv1beta1.Create, buildV1ServiceInstance(t, missingService)),
			buildAdmissionResponse(true, 0, metav1.StatusReasonUnknown, nil, "compute service \"missing-service\" doesn't exist in Service Central"),
			false,
		},
		{
			"ServiceHasSameOwner",
			buildAdmissionReview(elsieMicros2Service, k8s.ServiceInstanceGVR, admissionv1beta1.Create, buildV1ServiceInstance(t, elsieComputeService)),
			buildAdmissionResponse(true, 0, metav1.StatusReasonUnknown, nil, `service central owner of service "elsie-compute-service" (elsie) is same as micros2 service "elsie-micros2-service" (elsie)`),
			false,
		},
		{
			"ServiceHasDifferentOwnerForbidden",
			buildAdmissionReview(elsieMicros2Service, k8s.ServiceInstanceGVR, admissionv1beta1.Create, buildV1ServiceInstance(t, dougComputeService)),
			buildAdmissionResponse(false, http.StatusForbidden, metav1.StatusReasonForbidden, nil, `service central owner of service "doug-compute-service" (doug) is different to micros2 service "elsie-micros2-service" (elsie)`),
			false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := MicrosAdmitFunc(ctx, scStore, tc.admissionReview)
			if (err != nil) != tc.wantErr {
				t.Errorf("MicrosAdmitFunc() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MicrosAdmitFunc() = %v, want %v", got, tc.want)
			}
		})
	}
}

func buildV1ServiceInstance(t *testing.T, serviceName string) []byte {
	type Parameters struct {
		Name string `json:"name"`
	}
	return buildServiceInstance(
		t, microsClusterServiceClassName, microsV1ClusterServicePlanName,
		Parameters{serviceName})
}

func buildOtherServiceInstance(t *testing.T, serviceName string) []byte {
	type Service struct {
		ID string `json:"id"`
	}

	type Parameters struct {
		Service Service `json:"service"`
	}

	return buildServiceInstance(
		t, microsClusterServiceClassName, "foo",
		Parameters{Service{serviceName}})
}
