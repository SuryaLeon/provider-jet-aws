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

type TransitGatewayPeeringAttachmentAccepterObservation struct {
	PeerAccountID *string `json:"peerAccountId,omitempty" tf:"peer_account_id,omitempty"`

	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	PeerTransitGatewayID *string `json:"peerTransitGatewayId,omitempty" tf:"peer_transit_gateway_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`
}

type TransitGatewayPeeringAttachmentAccepterParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId" tf:"transit_gateway_attachment_id,omitempty"`
}

// TransitGatewayPeeringAttachmentAccepterSpec defines the desired state of TransitGatewayPeeringAttachmentAccepter
type TransitGatewayPeeringAttachmentAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayPeeringAttachmentAccepterParameters `json:"forProvider"`
}

// TransitGatewayPeeringAttachmentAccepterStatus defines the observed state of TransitGatewayPeeringAttachmentAccepter.
type TransitGatewayPeeringAttachmentAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayPeeringAttachmentAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPeeringAttachmentAccepter is the Schema for the TransitGatewayPeeringAttachmentAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TransitGatewayPeeringAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayPeeringAttachmentAccepterSpec   `json:"spec"`
	Status            TransitGatewayPeeringAttachmentAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPeeringAttachmentAccepterList contains a list of TransitGatewayPeeringAttachmentAccepters
type TransitGatewayPeeringAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayPeeringAttachmentAccepter `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayPeeringAttachmentAccepter_Kind             = "TransitGatewayPeeringAttachmentAccepter"
	TransitGatewayPeeringAttachmentAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayPeeringAttachmentAccepter_Kind}.String()
	TransitGatewayPeeringAttachmentAccepter_KindAPIVersion   = TransitGatewayPeeringAttachmentAccepter_Kind + "." + CRDGroupVersion.String()
	TransitGatewayPeeringAttachmentAccepter_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayPeeringAttachmentAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayPeeringAttachmentAccepter{}, &TransitGatewayPeeringAttachmentAccepterList{})
}
