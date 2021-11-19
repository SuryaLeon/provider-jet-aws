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

type GatewayAuthorizerObservation struct {
}

type GatewayAuthorizerParameters struct {

	// +kubebuilder:validation:Optional
	AuthorizerCredentials *string `json:"authorizerCredentials,omitempty" tf:"authorizer_credentials,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizerResultTTLInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// +kubebuilder:validation:Optional
	IdentitySource *string `json:"identitySource,omitempty" tf:"identity_source,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ProviderArns []*string `json:"providerArns,omitempty" tf:"provider_arns,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restApiId" tf:"rest_api_id,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// GatewayAuthorizerSpec defines the desired state of GatewayAuthorizer
type GatewayAuthorizerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayAuthorizerParameters `json:"forProvider"`
}

// GatewayAuthorizerStatus defines the observed state of GatewayAuthorizer.
type GatewayAuthorizerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayAuthorizerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAuthorizer is the Schema for the GatewayAuthorizers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GatewayAuthorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayAuthorizerSpec   `json:"spec"`
	Status            GatewayAuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAuthorizerList contains a list of GatewayAuthorizers
type GatewayAuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayAuthorizer `json:"items"`
}

// Repository type metadata.
var (
	GatewayAuthorizer_Kind             = "GatewayAuthorizer"
	GatewayAuthorizer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayAuthorizer_Kind}.String()
	GatewayAuthorizer_KindAPIVersion   = GatewayAuthorizer_Kind + "." + CRDGroupVersion.String()
	GatewayAuthorizer_GroupVersionKind = CRDGroupVersion.WithKind(GatewayAuthorizer_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayAuthorizer{}, &GatewayAuthorizerList{})
}
