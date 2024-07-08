/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// KubernetesNodeProperties struct for KubernetesNodeProperties
type KubernetesNodeProperties struct {
	// The Kubernetes version running in the node pool. Note that this imposes restrictions on which Kubernetes versions can run in the node pools of a cluster. Also, not all Kubernetes versions are suitable upgrade targets for all earlier versions.
	K8sVersion *string `json:"k8sVersion"`
	// The Kubernetes node name.
	Name *string `json:"name"`
	// The private IP associated with the node.
	PrivateIP *string `json:"privateIP,omitempty"`
	// The public IP associated with the node.
	PublicIP *string `json:"publicIP,omitempty"`
}

// NewKubernetesNodeProperties instantiates a new KubernetesNodeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeProperties(k8sVersion string, name string) *KubernetesNodeProperties {
	this := KubernetesNodeProperties{}

	this.K8sVersion = &k8sVersion
	this.Name = &name

	return &this
}

// NewKubernetesNodePropertiesWithDefaults instantiates a new KubernetesNodeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodePropertiesWithDefaults() *KubernetesNodeProperties {
	this := KubernetesNodeProperties{}
	return &this
}

// GetK8sVersion returns the K8sVersion field value
// If the value is explicit nil, nil is returned
func (o *KubernetesNodeProperties) GetK8sVersion() *string {
	if o == nil {
		return nil
	}

	return o.K8sVersion

}

// GetK8sVersionOk returns a tuple with the K8sVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProperties) GetK8sVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.K8sVersion, true
}

// SetK8sVersion sets field value
func (o *KubernetesNodeProperties) SetK8sVersion(v string) {

	o.K8sVersion = &v

}

// HasK8sVersion returns a boolean if a field has been set.
func (o *KubernetesNodeProperties) HasK8sVersion() bool {
	if o != nil && o.K8sVersion != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, nil is returned
func (o *KubernetesNodeProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *KubernetesNodeProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNodeProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetPrivateIP returns the PrivateIP field value
// If the value is explicit nil, nil is returned
func (o *KubernetesNodeProperties) GetPrivateIP() *string {
	if o == nil {
		return nil
	}

	return o.PrivateIP

}

// GetPrivateIPOk returns a tuple with the PrivateIP field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProperties) GetPrivateIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PrivateIP, true
}

// SetPrivateIP sets field value
func (o *KubernetesNodeProperties) SetPrivateIP(v string) {

	o.PrivateIP = &v

}

// HasPrivateIP returns a boolean if a field has been set.
func (o *KubernetesNodeProperties) HasPrivateIP() bool {
	if o != nil && o.PrivateIP != nil {
		return true
	}

	return false
}

// GetPublicIP returns the PublicIP field value
// If the value is explicit nil, nil is returned
func (o *KubernetesNodeProperties) GetPublicIP() *string {
	if o == nil {
		return nil
	}

	return o.PublicIP

}

// GetPublicIPOk returns a tuple with the PublicIP field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProperties) GetPublicIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PublicIP, true
}

// SetPublicIP sets field value
func (o *KubernetesNodeProperties) SetPublicIP(v string) {

	o.PublicIP = &v

}

// HasPublicIP returns a boolean if a field has been set.
func (o *KubernetesNodeProperties) HasPublicIP() bool {
	if o != nil && o.PublicIP != nil {
		return true
	}

	return false
}

func (o KubernetesNodeProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.K8sVersion != nil {
		toSerialize["k8sVersion"] = o.K8sVersion
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.PrivateIP != nil {
		toSerialize["privateIP"] = o.PrivateIP
	}

	if o.PublicIP != nil {
		toSerialize["publicIP"] = o.PublicIP
	}

	return json.Marshal(toSerialize)
}

type NullableKubernetesNodeProperties struct {
	value *KubernetesNodeProperties
	isSet bool
}

func (v NullableKubernetesNodeProperties) Get() *KubernetesNodeProperties {
	return v.value
}

func (v *NullableKubernetesNodeProperties) Set(val *KubernetesNodeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeProperties(val *KubernetesNodeProperties) *NullableKubernetesNodeProperties {
	return &NullableKubernetesNodeProperties{value: val, isSet: true}
}

func (v NullableKubernetesNodeProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}