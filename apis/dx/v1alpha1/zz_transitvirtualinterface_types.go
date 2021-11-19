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

type TransitVirtualInterfaceObservation struct {
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TransitVirtualInterfaceParameters struct {

	// +kubebuilder:validation:Required
	AddressFamily *string `json:"addressFamily" tf:"address_family,omitempty"`

	// +kubebuilder:validation:Optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// +kubebuilder:validation:Required
	BgpAsn *int64 `json:"bgpAsn" tf:"bgp_asn,omitempty"`

	// +kubebuilder:validation:Optional
	BgpAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionID *string `json:"connectionId" tf:"connection_id,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// +kubebuilder:validation:Required
	DxGatewayID *string `json:"dxGatewayId" tf:"dx_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Vlan *int64 `json:"vlan" tf:"vlan,omitempty"`
}

// TransitVirtualInterfaceSpec defines the desired state of TransitVirtualInterface
type TransitVirtualInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitVirtualInterfaceParameters `json:"forProvider"`
}

// TransitVirtualInterfaceStatus defines the observed state of TransitVirtualInterface.
type TransitVirtualInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitVirtualInterface is the Schema for the TransitVirtualInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TransitVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitVirtualInterfaceSpec   `json:"spec"`
	Status            TransitVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitVirtualInterfaceList contains a list of TransitVirtualInterfaces
type TransitVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	TransitVirtualInterface_Kind             = "TransitVirtualInterface"
	TransitVirtualInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitVirtualInterface_Kind}.String()
	TransitVirtualInterface_KindAPIVersion   = TransitVirtualInterface_Kind + "." + CRDGroupVersion.String()
	TransitVirtualInterface_GroupVersionKind = CRDGroupVersion.WithKind(TransitVirtualInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitVirtualInterface{}, &TransitVirtualInterfaceList{})
}
