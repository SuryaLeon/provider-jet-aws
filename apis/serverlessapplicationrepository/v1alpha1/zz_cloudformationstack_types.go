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

type CloudformationStackObservation struct {
	Outputs map[string]*string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CloudformationStackParameters struct {

	// +kubebuilder:validation:Required
	ApplicationID *string `json:"applicationId" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Required
	Capabilities []*string `json:"capabilities" tf:"capabilities,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SemanticVersion *string `json:"semanticVersion,omitempty" tf:"semantic_version,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// CloudformationStackSpec defines the desired state of CloudformationStack
type CloudformationStackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudformationStackParameters `json:"forProvider"`
}

// CloudformationStackStatus defines the observed state of CloudformationStack.
type CloudformationStackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudformationStackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudformationStack is the Schema for the CloudformationStacks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type CloudformationStack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudformationStackSpec   `json:"spec"`
	Status            CloudformationStackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudformationStackList contains a list of CloudformationStacks
type CloudformationStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudformationStack `json:"items"`
}

// Repository type metadata.
var (
	CloudformationStack_Kind             = "CloudformationStack"
	CloudformationStack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudformationStack_Kind}.String()
	CloudformationStack_KindAPIVersion   = CloudformationStack_Kind + "." + CRDGroupVersion.String()
	CloudformationStack_GroupVersionKind = CRDGroupVersion.WithKind(CloudformationStack_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudformationStack{}, &CloudformationStackList{})
}
