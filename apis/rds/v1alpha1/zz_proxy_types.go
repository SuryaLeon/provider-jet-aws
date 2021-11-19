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

type AuthObservation struct {
}

type AuthParameters struct {

	// +kubebuilder:validation:Optional
	AuthScheme *string `json:"authScheme,omitempty" tf:"auth_scheme,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IamAuth *string `json:"iamAuth,omitempty" tf:"iam_auth,omitempty"`

	// +kubebuilder:validation:Optional
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`
}

type ProxyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ProxyParameters struct {

	// +kubebuilder:validation:Required
	Auth []AuthParameters `json:"auth" tf:"auth,omitempty"`

	// +kubebuilder:validation:Optional
	DebugLogging *bool `json:"debugLogging,omitempty" tf:"debug_logging,omitempty"`

	// +kubebuilder:validation:Required
	EngineFamily *string `json:"engineFamily" tf:"engine_family,omitempty"`

	// +kubebuilder:validation:Optional
	IdleClientTimeout *int64 `json:"idleClientTimeout,omitempty" tf:"idle_client_timeout,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequireTLS *bool `json:"requireTls,omitempty" tf:"require_tls,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VpcSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	VpcSubnetIds []*string `json:"vpcSubnetIds" tf:"vpc_subnet_ids,omitempty"`
}

// ProxySpec defines the desired state of Proxy
type ProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProxyParameters `json:"forProvider"`
}

// ProxyStatus defines the observed state of Proxy.
type ProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Proxy is the Schema for the Proxys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Proxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxySpec   `json:"spec"`
	Status            ProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyList contains a list of Proxys
type ProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Proxy `json:"items"`
}

// Repository type metadata.
var (
	Proxy_Kind             = "Proxy"
	Proxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Proxy_Kind}.String()
	Proxy_KindAPIVersion   = Proxy_Kind + "." + CRDGroupVersion.String()
	Proxy_GroupVersionKind = CRDGroupVersion.WithKind(Proxy_Kind)
)

func init() {
	SchemeBuilder.Register(&Proxy{}, &ProxyList{})
}
