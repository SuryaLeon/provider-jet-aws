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

type EndpointConfigurationObservation struct {
}

type EndpointConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Types []*string `json:"types" tf:"types,omitempty"`
}

type GatewayDomainNameObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CertificateUploadDate *string `json:"certificateUploadDate,omitempty" tf:"certificate_upload_date,omitempty"`

	CloudfrontDomainName *string `json:"cloudfrontDomainName,omitempty" tf:"cloudfront_domain_name,omitempty"`

	CloudfrontZoneID *string `json:"cloudfrontZoneId,omitempty" tf:"cloudfront_zone_id,omitempty"`

	RegionalDomainName *string `json:"regionalDomainName,omitempty" tf:"regional_domain_name,omitempty"`

	RegionalZoneID *string `json:"regionalZoneId,omitempty" tf:"regional_zone_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GatewayDomainNameParameters struct {

	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// +kubebuilder:validation:Optional
	CertificatePrivateKeySecretRef *v1.SecretKeySelector `json:"certificatePrivateKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointConfiguration []EndpointConfigurationParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	MutualTLSAuthentication []MutualTLSAuthenticationParameters `json:"mutualTlsAuthentication,omitempty" tf:"mutual_tls_authentication,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RegionalCertificateArn *string `json:"regionalCertificateArn,omitempty" tf:"regional_certificate_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RegionalCertificateName *string `json:"regionalCertificateName,omitempty" tf:"regional_certificate_name,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityPolicy *string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MutualTLSAuthenticationObservation struct {
}

type MutualTLSAuthenticationParameters struct {

	// +kubebuilder:validation:Required
	TruststoreURI *string `json:"truststoreUri" tf:"truststore_uri,omitempty"`

	// +kubebuilder:validation:Optional
	TruststoreVersion *string `json:"truststoreVersion,omitempty" tf:"truststore_version,omitempty"`
}

// GatewayDomainNameSpec defines the desired state of GatewayDomainName
type GatewayDomainNameSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayDomainNameParameters `json:"forProvider"`
}

// GatewayDomainNameStatus defines the observed state of GatewayDomainName.
type GatewayDomainNameStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayDomainNameObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayDomainName is the Schema for the GatewayDomainNames API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GatewayDomainName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayDomainNameSpec   `json:"spec"`
	Status            GatewayDomainNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayDomainNameList contains a list of GatewayDomainNames
type GatewayDomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayDomainName `json:"items"`
}

// Repository type metadata.
var (
	GatewayDomainName_Kind             = "GatewayDomainName"
	GatewayDomainName_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayDomainName_Kind}.String()
	GatewayDomainName_KindAPIVersion   = GatewayDomainName_Kind + "." + CRDGroupVersion.String()
	GatewayDomainName_GroupVersionKind = CRDGroupVersion.WithKind(GatewayDomainName_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayDomainName{}, &GatewayDomainNameList{})
}
