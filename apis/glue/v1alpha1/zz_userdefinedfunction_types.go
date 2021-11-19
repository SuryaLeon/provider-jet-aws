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

type ResourceUrisObservation struct {
}

type ResourceUrisParameters struct {

	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`

	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type UserDefinedFunctionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`
}

type UserDefinedFunctionParameters struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// +kubebuilder:validation:Required
	ClassName *string `json:"className" tf:"class_name,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	OwnerName *string `json:"ownerName" tf:"owner_name,omitempty"`

	// +kubebuilder:validation:Required
	OwnerType *string `json:"ownerType" tf:"owner_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceUris []ResourceUrisParameters `json:"resourceUris,omitempty" tf:"resource_uris,omitempty"`
}

// UserDefinedFunctionSpec defines the desired state of UserDefinedFunction
type UserDefinedFunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserDefinedFunctionParameters `json:"forProvider"`
}

// UserDefinedFunctionStatus defines the observed state of UserDefinedFunction.
type UserDefinedFunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserDefinedFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserDefinedFunction is the Schema for the UserDefinedFunctions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type UserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserDefinedFunctionSpec   `json:"spec"`
	Status            UserDefinedFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserDefinedFunctionList contains a list of UserDefinedFunctions
type UserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserDefinedFunction `json:"items"`
}

// Repository type metadata.
var (
	UserDefinedFunction_Kind             = "UserDefinedFunction"
	UserDefinedFunction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserDefinedFunction_Kind}.String()
	UserDefinedFunction_KindAPIVersion   = UserDefinedFunction_Kind + "." + CRDGroupVersion.String()
	UserDefinedFunction_GroupVersionKind = CRDGroupVersion.WithKind(UserDefinedFunction_Kind)
)

func init() {
	SchemeBuilder.Register(&UserDefinedFunction{}, &UserDefinedFunctionList{})
}
