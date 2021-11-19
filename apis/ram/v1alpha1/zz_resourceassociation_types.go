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

type ResourceAssociationObservation struct {
}

type ResourceAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Required
	ResourceShareArn *string `json:"resourceShareArn" tf:"resource_share_arn,omitempty"`
}

// ResourceAssociationSpec defines the desired state of ResourceAssociation
type ResourceAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceAssociationParameters `json:"forProvider"`
}

// ResourceAssociationStatus defines the observed state of ResourceAssociation.
type ResourceAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceAssociation is the Schema for the ResourceAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceAssociationSpec   `json:"spec"`
	Status            ResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceAssociationList contains a list of ResourceAssociations
type ResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceAssociation `json:"items"`
}

// Repository type metadata.
var (
	ResourceAssociation_Kind             = "ResourceAssociation"
	ResourceAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceAssociation_Kind}.String()
	ResourceAssociation_KindAPIVersion   = ResourceAssociation_Kind + "." + CRDGroupVersion.String()
	ResourceAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ResourceAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceAssociation{}, &ResourceAssociationList{})
}
