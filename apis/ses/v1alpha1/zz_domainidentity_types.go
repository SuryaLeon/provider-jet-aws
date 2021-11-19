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

type DomainIdentityObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	VerificationToken *string `json:"verificationToken,omitempty" tf:"verification_token,omitempty"`
}

type DomainIdentityParameters struct {

	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainIdentitySpec defines the desired state of DomainIdentity
type DomainIdentitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainIdentityParameters `json:"forProvider"`
}

// DomainIdentityStatus defines the observed state of DomainIdentity.
type DomainIdentityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainIdentity is the Schema for the DomainIdentitys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DomainIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainIdentitySpec   `json:"spec"`
	Status            DomainIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainIdentityList contains a list of DomainIdentitys
type DomainIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainIdentity `json:"items"`
}

// Repository type metadata.
var (
	DomainIdentity_Kind             = "DomainIdentity"
	DomainIdentity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainIdentity_Kind}.String()
	DomainIdentity_KindAPIVersion   = DomainIdentity_Kind + "." + CRDGroupVersion.String()
	DomainIdentity_GroupVersionKind = CRDGroupVersion.WithKind(DomainIdentity_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainIdentity{}, &DomainIdentityList{})
}
