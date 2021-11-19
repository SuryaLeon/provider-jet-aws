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

type SpecListenerPortMappingObservation struct {
}

type SpecListenerPortMappingParameters struct {

	// +kubebuilder:validation:Required
	Port *int64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type VirtualRouterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VirtualRouterParameters struct {

	// +kubebuilder:validation:Required
	MeshName *string `json:"meshName" tf:"mesh_name,omitempty"`

	// +kubebuilder:validation:Optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Spec []VirtualRouterSpecParameters `json:"spec" tf:"spec,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualRouterSpecListenerObservation struct {
}

type VirtualRouterSpecListenerParameters struct {

	// +kubebuilder:validation:Required
	PortMapping []SpecListenerPortMappingParameters `json:"portMapping" tf:"port_mapping,omitempty"`
}

type VirtualRouterSpecObservation struct {
}

type VirtualRouterSpecParameters struct {

	// +kubebuilder:validation:Required
	Listener []VirtualRouterSpecListenerParameters `json:"listener" tf:"listener,omitempty"`
}

// VirtualRouterSpec defines the desired state of VirtualRouter
type VirtualRouterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualRouterParameters `json:"forProvider"`
}

// VirtualRouterStatus defines the observed state of VirtualRouter.
type VirtualRouterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualRouterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualRouter is the Schema for the VirtualRouters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type VirtualRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualRouterSpec   `json:"spec"`
	Status            VirtualRouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualRouterList contains a list of VirtualRouters
type VirtualRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualRouter `json:"items"`
}

// Repository type metadata.
var (
	VirtualRouter_Kind             = "VirtualRouter"
	VirtualRouter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualRouter_Kind}.String()
	VirtualRouter_KindAPIVersion   = VirtualRouter_Kind + "." + CRDGroupVersion.String()
	VirtualRouter_GroupVersionKind = CRDGroupVersion.WithKind(VirtualRouter_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualRouter{}, &VirtualRouterList{})
}
