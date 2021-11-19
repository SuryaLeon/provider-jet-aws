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

type GatewayAssociationProposalObservation struct {
	AssociatedGatewayOwnerAccountID *string `json:"associatedGatewayOwnerAccountId,omitempty" tf:"associated_gateway_owner_account_id,omitempty"`

	AssociatedGatewayType *string `json:"associatedGatewayType,omitempty" tf:"associated_gateway_type,omitempty"`
}

type GatewayAssociationProposalParameters struct {

	// +kubebuilder:validation:Optional
	AllowedPrefixes []*string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes,omitempty"`

	// +kubebuilder:validation:Required
	AssociatedGatewayID *string `json:"associatedGatewayId" tf:"associated_gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	DxGatewayID *string `json:"dxGatewayId" tf:"dx_gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	DxGatewayOwnerAccountID *string `json:"dxGatewayOwnerAccountId" tf:"dx_gateway_owner_account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// GatewayAssociationProposalSpec defines the desired state of GatewayAssociationProposal
type GatewayAssociationProposalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayAssociationProposalParameters `json:"forProvider"`
}

// GatewayAssociationProposalStatus defines the observed state of GatewayAssociationProposal.
type GatewayAssociationProposalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayAssociationProposalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAssociationProposal is the Schema for the GatewayAssociationProposals API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayAssociationProposalSpec   `json:"spec"`
	Status            GatewayAssociationProposalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAssociationProposalList contains a list of GatewayAssociationProposals
type GatewayAssociationProposalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayAssociationProposal `json:"items"`
}

// Repository type metadata.
var (
	GatewayAssociationProposal_Kind             = "GatewayAssociationProposal"
	GatewayAssociationProposal_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayAssociationProposal_Kind}.String()
	GatewayAssociationProposal_KindAPIVersion   = GatewayAssociationProposal_Kind + "." + CRDGroupVersion.String()
	GatewayAssociationProposal_GroupVersionKind = CRDGroupVersion.WithKind(GatewayAssociationProposal_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayAssociationProposal{}, &GatewayAssociationProposalList{})
}
