// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuota) DeepCopyInto(out *ServiceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuota.
func (in *ServiceQuota) DeepCopy() *ServiceQuota {
	if in == nil {
		return nil
	}
	out := new(ServiceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuotaList) DeepCopyInto(out *ServiceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuotaList.
func (in *ServiceQuotaList) DeepCopy() *ServiceQuotaList {
	if in == nil {
		return nil
	}
	out := new(ServiceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuotaObservation) DeepCopyInto(out *ServiceQuotaObservation) {
	*out = *in
	if in.Adjustable != nil {
		in, out := &in.Adjustable, &out.Adjustable
		*out = new(bool)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(float64)
		**out = **in
	}
	if in.QuotaName != nil {
		in, out := &in.QuotaName, &out.QuotaName
		*out = new(string)
		**out = **in
	}
	if in.RequestID != nil {
		in, out := &in.RequestID, &out.RequestID
		*out = new(string)
		**out = **in
	}
	if in.RequestStatus != nil {
		in, out := &in.RequestStatus, &out.RequestStatus
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuotaObservation.
func (in *ServiceQuotaObservation) DeepCopy() *ServiceQuotaObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceQuotaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuotaParameters) DeepCopyInto(out *ServiceQuotaParameters) {
	*out = *in
	if in.QuotaCode != nil {
		in, out := &in.QuotaCode, &out.QuotaCode
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServiceCode != nil {
		in, out := &in.ServiceCode, &out.ServiceCode
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuotaParameters.
func (in *ServiceQuotaParameters) DeepCopy() *ServiceQuotaParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceQuotaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuotaSpec) DeepCopyInto(out *ServiceQuotaSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuotaSpec.
func (in *ServiceQuotaSpec) DeepCopy() *ServiceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuotaStatus) DeepCopyInto(out *ServiceQuotaStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuotaStatus.
func (in *ServiceQuotaStatus) DeepCopy() *ServiceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}
