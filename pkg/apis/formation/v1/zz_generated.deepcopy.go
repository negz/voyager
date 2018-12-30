// +build !ignore_autogenerated

// Generated code
// run `make generate` to update

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	conditionv1 "github.com/atlassian/ctrl/apis/condition/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationDescriptor) DeepCopyInto(out *LocationDescriptor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationDescriptor.
func (in *LocationDescriptor) DeepCopy() *LocationDescriptor {
	if in == nil {
		return nil
	}
	out := new(LocationDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocationDescriptor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationDescriptorConfigMapNames) DeepCopyInto(out *LocationDescriptorConfigMapNames) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationDescriptorConfigMapNames.
func (in *LocationDescriptorConfigMapNames) DeepCopy() *LocationDescriptorConfigMapNames {
	if in == nil {
		return nil
	}
	out := new(LocationDescriptorConfigMapNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationDescriptorDependency.
func (in *LocationDescriptorDependency) DeepCopy() *LocationDescriptorDependency {
	if in == nil {
		return nil
	}
	out := new(LocationDescriptorDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationDescriptorList) DeepCopyInto(out *LocationDescriptorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocationDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationDescriptorList.
func (in *LocationDescriptorList) DeepCopy() *LocationDescriptorList {
	if in == nil {
		return nil
	}
	out := new(LocationDescriptorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocationDescriptorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationDescriptorResource) DeepCopyInto(out *LocationDescriptorResource) {
	*out = *in
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = make([]LocationDescriptorDependency, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationDescriptorResource.
func (in *LocationDescriptorResource) DeepCopy() *LocationDescriptorResource {
	if in == nil {
		return nil
	}
	out := new(LocationDescriptorResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationDescriptorSpec) DeepCopyInto(out *LocationDescriptorSpec) {
	*out = *in
	out.ConfigMapNames = in.ConfigMapNames
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]LocationDescriptorResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationDescriptorSpec.
func (in *LocationDescriptorSpec) DeepCopy() *LocationDescriptorSpec {
	if in == nil {
		return nil
	}
	out := new(LocationDescriptorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationDescriptorStatus) DeepCopyInto(out *LocationDescriptorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditionv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceStatuses != nil {
		in, out := &in.ResourceStatuses, &out.ResourceStatuses
		*out = make([]ResourceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationDescriptorStatus.
func (in *LocationDescriptorStatus) DeepCopy() *LocationDescriptorStatus {
	if in == nil {
		return nil
	}
	out := new(LocationDescriptorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatus) DeepCopyInto(out *ResourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditionv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatus.
func (in *ResourceStatus) DeepCopy() *ResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceStatus)
	in.DeepCopyInto(out)
	return out
}
