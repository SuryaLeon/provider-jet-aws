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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ManagedPolicyAttachmentObservation struct {
	ManagedPolicyName *string `json:"managedPolicyName,omitempty" tf:"managed_policy_name,omitempty"`
}

type ManagedPolicyAttachmentParameters struct {

	// +kubebuilder:validation:Required
	InstanceArn *string `json:"instanceArn" tf:"instance_arn,omitempty"`

	// +kubebuilder:validation:Required
	ManagedPolicyArn *string `json:"managedPolicyArn" tf:"managed_policy_arn,omitempty"`

	// +kubebuilder:validation:Required
	PermissionSetArn *string `json:"permissionSetArn" tf:"permission_set_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ManagedPolicyAttachmentSpec defines the desired state of ManagedPolicyAttachment
type ManagedPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedPolicyAttachmentParameters `json:"forProvider"`
}

// ManagedPolicyAttachmentStatus defines the observed state of ManagedPolicyAttachment.
type ManagedPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPolicyAttachment is the Schema for the ManagedPolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ManagedPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedPolicyAttachmentSpec   `json:"spec"`
	Status            ManagedPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPolicyAttachmentList contains a list of ManagedPolicyAttachments
type ManagedPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	ManagedPolicyAttachment_Kind             = "ManagedPolicyAttachment"
	ManagedPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedPolicyAttachment_Kind}.String()
	ManagedPolicyAttachment_KindAPIVersion   = ManagedPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	ManagedPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(ManagedPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedPolicyAttachment{}, &ManagedPolicyAttachmentList{})
}
