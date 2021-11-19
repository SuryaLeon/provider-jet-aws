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

type LocalGatewayRouteTableVpcAssociationObservation struct {
	LocalGatewayID *string `json:"localGatewayId,omitempty" tf:"local_gateway_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LocalGatewayRouteTableVpcAssociationParameters struct {

	// +kubebuilder:validation:Required
	LocalGatewayRouteTableID *string `json:"localGatewayRouteTableId" tf:"local_gateway_route_table_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VpcID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

// LocalGatewayRouteTableVpcAssociationSpec defines the desired state of LocalGatewayRouteTableVpcAssociation
type LocalGatewayRouteTableVpcAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocalGatewayRouteTableVpcAssociationParameters `json:"forProvider"`
}

// LocalGatewayRouteTableVpcAssociationStatus defines the observed state of LocalGatewayRouteTableVpcAssociation.
type LocalGatewayRouteTableVpcAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocalGatewayRouteTableVpcAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocalGatewayRouteTableVpcAssociation is the Schema for the LocalGatewayRouteTableVpcAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LocalGatewayRouteTableVpcAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocalGatewayRouteTableVpcAssociationSpec   `json:"spec"`
	Status            LocalGatewayRouteTableVpcAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocalGatewayRouteTableVpcAssociationList contains a list of LocalGatewayRouteTableVpcAssociations
type LocalGatewayRouteTableVpcAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocalGatewayRouteTableVpcAssociation `json:"items"`
}

// Repository type metadata.
var (
	LocalGatewayRouteTableVpcAssociation_Kind             = "LocalGatewayRouteTableVpcAssociation"
	LocalGatewayRouteTableVpcAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocalGatewayRouteTableVpcAssociation_Kind}.String()
	LocalGatewayRouteTableVpcAssociation_KindAPIVersion   = LocalGatewayRouteTableVpcAssociation_Kind + "." + CRDGroupVersion.String()
	LocalGatewayRouteTableVpcAssociation_GroupVersionKind = CRDGroupVersion.WithKind(LocalGatewayRouteTableVpcAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&LocalGatewayRouteTableVpcAssociation{}, &LocalGatewayRouteTableVpcAssociationList{})
}
