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

type MemberAccountAssociationObservation struct {
}

type MemberAccountAssociationParameters struct {

	// +kubebuilder:validation:Required
	MemberAccountID *string `json:"memberAccountId" tf:"member_account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// MemberAccountAssociationSpec defines the desired state of MemberAccountAssociation
type MemberAccountAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberAccountAssociationParameters `json:"forProvider"`
}

// MemberAccountAssociationStatus defines the observed state of MemberAccountAssociation.
type MemberAccountAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberAccountAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MemberAccountAssociation is the Schema for the MemberAccountAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type MemberAccountAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MemberAccountAssociationSpec   `json:"spec"`
	Status            MemberAccountAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberAccountAssociationList contains a list of MemberAccountAssociations
type MemberAccountAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemberAccountAssociation `json:"items"`
}

// Repository type metadata.
var (
	MemberAccountAssociation_Kind             = "MemberAccountAssociation"
	MemberAccountAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MemberAccountAssociation_Kind}.String()
	MemberAccountAssociation_KindAPIVersion   = MemberAccountAssociation_Kind + "." + CRDGroupVersion.String()
	MemberAccountAssociation_GroupVersionKind = CRDGroupVersion.WithKind(MemberAccountAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&MemberAccountAssociation{}, &MemberAccountAssociationList{})
}
