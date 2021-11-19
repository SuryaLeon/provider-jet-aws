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

type HlsIngestObservation struct {
	IngestEndpoints []IngestEndpointsObservation `json:"ingestEndpoints,omitempty" tf:"ingest_endpoints,omitempty"`
}

type HlsIngestParameters struct {
}

type IngestEndpointsObservation struct {
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type IngestEndpointsParameters struct {
}

type PackageChannelObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	HlsIngest []HlsIngestObservation `json:"hlsIngest,omitempty" tf:"hls_ingest,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PackageChannelParameters struct {

	// +kubebuilder:validation:Required
	ChannelID *string `json:"channelId" tf:"channel_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PackageChannelSpec defines the desired state of PackageChannel
type PackageChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PackageChannelParameters `json:"forProvider"`
}

// PackageChannelStatus defines the observed state of PackageChannel.
type PackageChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PackageChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PackageChannel is the Schema for the PackageChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type PackageChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PackageChannelSpec   `json:"spec"`
	Status            PackageChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PackageChannelList contains a list of PackageChannels
type PackageChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PackageChannel `json:"items"`
}

// Repository type metadata.
var (
	PackageChannel_Kind             = "PackageChannel"
	PackageChannel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PackageChannel_Kind}.String()
	PackageChannel_KindAPIVersion   = PackageChannel_Kind + "." + CRDGroupVersion.String()
	PackageChannel_GroupVersionKind = CRDGroupVersion.WithKind(PackageChannel_Kind)
)

func init() {
	SchemeBuilder.Register(&PackageChannel{}, &PackageChannelList{})
}
