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

type APIStagesObservation struct {
}

type APIStagesParameters struct {

	// +kubebuilder:validation:Required
	APIID *string `json:"apiId" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Required
	Stage *string `json:"stage" tf:"stage,omitempty"`
}

type GatewayUsagePlanObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GatewayUsagePlanParameters struct {

	// +kubebuilder:validation:Optional
	APIStages []APIStagesParameters `json:"apiStages,omitempty" tf:"api_stages,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ProductCode *string `json:"productCode,omitempty" tf:"product_code,omitempty"`

	// +kubebuilder:validation:Optional
	QuotaSettings []QuotaSettingsParameters `json:"quotaSettings,omitempty" tf:"quota_settings,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ThrottleSettings []GatewayUsagePlanThrottleSettingsParameters `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

type GatewayUsagePlanThrottleSettingsObservation struct {
}

type GatewayUsagePlanThrottleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	BurstLimit *int64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	// +kubebuilder:validation:Optional
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type QuotaSettingsObservation struct {
}

type QuotaSettingsParameters struct {

	// +kubebuilder:validation:Required
	Limit *int64 `json:"limit" tf:"limit,omitempty"`

	// +kubebuilder:validation:Optional
	Offset *int64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// +kubebuilder:validation:Required
	Period *string `json:"period" tf:"period,omitempty"`
}

// GatewayUsagePlanSpec defines the desired state of GatewayUsagePlan
type GatewayUsagePlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayUsagePlanParameters `json:"forProvider"`
}

// GatewayUsagePlanStatus defines the observed state of GatewayUsagePlan.
type GatewayUsagePlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayUsagePlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayUsagePlan is the Schema for the GatewayUsagePlans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GatewayUsagePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayUsagePlanSpec   `json:"spec"`
	Status            GatewayUsagePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayUsagePlanList contains a list of GatewayUsagePlans
type GatewayUsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayUsagePlan `json:"items"`
}

// Repository type metadata.
var (
	GatewayUsagePlan_Kind             = "GatewayUsagePlan"
	GatewayUsagePlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayUsagePlan_Kind}.String()
	GatewayUsagePlan_KindAPIVersion   = GatewayUsagePlan_Kind + "." + CRDGroupVersion.String()
	GatewayUsagePlan_GroupVersionKind = CRDGroupVersion.WithKind(GatewayUsagePlan_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayUsagePlan{}, &GatewayUsagePlanList{})
}
