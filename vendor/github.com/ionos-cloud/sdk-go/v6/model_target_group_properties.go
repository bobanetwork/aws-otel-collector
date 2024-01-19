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

// TargetGroupProperties struct for TargetGroupProperties
type TargetGroupProperties struct {
	// The balancing algorithm. A balancing algorithm consists of predefined rules with the logic that a load balancer uses to distribute network traffic between servers.  - **Round Robin**: Targets are served alternately according to their weighting.  - **Least Connection**: The target with the least active connection is served.  - **Random**: The targets are served based on a consistent pseudorandom algorithm.  - **Source IP**: It is ensured that the same client IP address reaches the same target.
	Algorithm       *string                     `json:"algorithm"`
	HealthCheck     *TargetGroupHealthCheck     `json:"healthCheck,omitempty"`
	HttpHealthCheck *TargetGroupHttpHealthCheck `json:"httpHealthCheck,omitempty"`
	// The target group name.
	Name *string `json:"name"`
	// The forwarding protocol. Only the value 'HTTP' is allowed.
	Protocol *string `json:"protocol"`
	// Array of items in the collection.
	Targets *[]TargetGroupTarget `json:"targets,omitempty"`
}

// NewTargetGroupProperties instantiates a new TargetGroupProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetGroupProperties(algorithm string, name string, protocol string) *TargetGroupProperties {
	this := TargetGroupProperties{}

	this.Algorithm = &algorithm
	this.Name = &name
	this.Protocol = &protocol

	return &this
}

// NewTargetGroupPropertiesWithDefaults instantiates a new TargetGroupProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetGroupPropertiesWithDefaults() *TargetGroupProperties {
	this := TargetGroupProperties{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupProperties) GetAlgorithm() *string {
	if o == nil {
		return nil
	}

	return o.Algorithm

}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupProperties) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *TargetGroupProperties) SetAlgorithm(v string) {

	o.Algorithm = &v

}

// HasAlgorithm returns a boolean if a field has been set.
func (o *TargetGroupProperties) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// GetHealthCheck returns the HealthCheck field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupProperties) GetHealthCheck() *TargetGroupHealthCheck {
	if o == nil {
		return nil
	}

	return o.HealthCheck

}

// GetHealthCheckOk returns a tuple with the HealthCheck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupProperties) GetHealthCheckOk() (*TargetGroupHealthCheck, bool) {
	if o == nil {
		return nil, false
	}

	return o.HealthCheck, true
}

// SetHealthCheck sets field value
func (o *TargetGroupProperties) SetHealthCheck(v TargetGroupHealthCheck) {

	o.HealthCheck = &v

}

// HasHealthCheck returns a boolean if a field has been set.
func (o *TargetGroupProperties) HasHealthCheck() bool {
	if o != nil && o.HealthCheck != nil {
		return true
	}

	return false
}

// GetHttpHealthCheck returns the HttpHealthCheck field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupProperties) GetHttpHealthCheck() *TargetGroupHttpHealthCheck {
	if o == nil {
		return nil
	}

	return o.HttpHealthCheck

}

// GetHttpHealthCheckOk returns a tuple with the HttpHealthCheck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupProperties) GetHttpHealthCheckOk() (*TargetGroupHttpHealthCheck, bool) {
	if o == nil {
		return nil, false
	}

	return o.HttpHealthCheck, true
}

// SetHttpHealthCheck sets field value
func (o *TargetGroupProperties) SetHttpHealthCheck(v TargetGroupHttpHealthCheck) {

	o.HttpHealthCheck = &v

}

// HasHttpHealthCheck returns a boolean if a field has been set.
func (o *TargetGroupProperties) HasHttpHealthCheck() bool {
	if o != nil && o.HttpHealthCheck != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *TargetGroupProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *TargetGroupProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetProtocol returns the Protocol field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupProperties) GetProtocol() *string {
	if o == nil {
		return nil
	}

	return o.Protocol

}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupProperties) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Protocol, true
}

// SetProtocol sets field value
func (o *TargetGroupProperties) SetProtocol(v string) {

	o.Protocol = &v

}

// HasProtocol returns a boolean if a field has been set.
func (o *TargetGroupProperties) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// GetTargets returns the Targets field value
// If the value is explicit nil, nil is returned
func (o *TargetGroupProperties) GetTargets() *[]TargetGroupTarget {
	if o == nil {
		return nil
	}

	return o.Targets

}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetGroupProperties) GetTargetsOk() (*[]TargetGroupTarget, bool) {
	if o == nil {
		return nil, false
	}

	return o.Targets, true
}

// SetTargets sets field value
func (o *TargetGroupProperties) SetTargets(v []TargetGroupTarget) {

	o.Targets = &v

}

// HasTargets returns a boolean if a field has been set.
func (o *TargetGroupProperties) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

func (o TargetGroupProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}

	if o.HealthCheck != nil {
		toSerialize["healthCheck"] = o.HealthCheck
	}

	if o.HttpHealthCheck != nil {
		toSerialize["httpHealthCheck"] = o.HttpHealthCheck
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}

	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}

	return json.Marshal(toSerialize)
}

type NullableTargetGroupProperties struct {
	value *TargetGroupProperties
	isSet bool
}

func (v NullableTargetGroupProperties) Get() *TargetGroupProperties {
	return v.value
}

func (v *NullableTargetGroupProperties) Set(val *TargetGroupProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetGroupProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetGroupProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetGroupProperties(val *TargetGroupProperties) *NullableTargetGroupProperties {
	return &NullableTargetGroupProperties{value: val, isSet: true}
}

func (v NullableTargetGroupProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetGroupProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
