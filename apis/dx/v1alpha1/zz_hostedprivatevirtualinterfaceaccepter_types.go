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

type HostedPrivateVirtualInterfaceAccepterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HostedPrivateVirtualInterfaceAccepterParameters struct {

	// +kubebuilder:validation:Optional
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualInterfaceID *string `json:"virtualInterfaceId" tf:"virtual_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpnGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`
}

// HostedPrivateVirtualInterfaceAccepterSpec defines the desired state of HostedPrivateVirtualInterfaceAccepter
type HostedPrivateVirtualInterfaceAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedPrivateVirtualInterfaceAccepterParameters `json:"forProvider"`
}

// HostedPrivateVirtualInterfaceAccepterStatus defines the observed state of HostedPrivateVirtualInterfaceAccepter.
type HostedPrivateVirtualInterfaceAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedPrivateVirtualInterfaceAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPrivateVirtualInterfaceAccepter is the Schema for the HostedPrivateVirtualInterfaceAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type HostedPrivateVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedPrivateVirtualInterfaceAccepterSpec   `json:"spec"`
	Status            HostedPrivateVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPrivateVirtualInterfaceAccepterList contains a list of HostedPrivateVirtualInterfaceAccepters
type HostedPrivateVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedPrivateVirtualInterfaceAccepter `json:"items"`
}

// Repository type metadata.
var (
	HostedPrivateVirtualInterfaceAccepter_Kind             = "HostedPrivateVirtualInterfaceAccepter"
	HostedPrivateVirtualInterfaceAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedPrivateVirtualInterfaceAccepter_Kind}.String()
	HostedPrivateVirtualInterfaceAccepter_KindAPIVersion   = HostedPrivateVirtualInterfaceAccepter_Kind + "." + CRDGroupVersion.String()
	HostedPrivateVirtualInterfaceAccepter_GroupVersionKind = CRDGroupVersion.WithKind(HostedPrivateVirtualInterfaceAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedPrivateVirtualInterfaceAccepter{}, &HostedPrivateVirtualInterfaceAccepterList{})
}
