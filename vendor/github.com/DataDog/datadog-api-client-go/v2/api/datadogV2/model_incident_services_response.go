// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentServicesResponse Response with a list of incident service payloads.
type IncidentServicesResponse struct {
	// An array of incident services.
	Data []IncidentServiceResponseData `json:"data"`
	// Included related resources which the user requested.
	Included []IncidentServiceIncludedItems `json:"included,omitempty"`
	// The metadata object containing pagination metadata.
	Meta *IncidentResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentServicesResponse instantiates a new IncidentServicesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentServicesResponse(data []IncidentServiceResponseData) *IncidentServicesResponse {
	this := IncidentServicesResponse{}
	this.Data = data
	return &this
}

// NewIncidentServicesResponseWithDefaults instantiates a new IncidentServicesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentServicesResponseWithDefaults() *IncidentServicesResponse {
	this := IncidentServicesResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentServicesResponse) GetData() []IncidentServiceResponseData {
	if o == nil {
		var ret []IncidentServiceResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponse) GetDataOk() (*[]IncidentServiceResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentServicesResponse) SetData(v []IncidentServiceResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentServicesResponse) GetIncluded() []IncidentServiceIncludedItems {
	if o == nil || o.Included == nil {
		var ret []IncidentServiceIncludedItems
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponse) GetIncludedOk() (*[]IncidentServiceIncludedItems, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentServicesResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentServiceIncludedItems and assigns it to the Included field.
func (o *IncidentServicesResponse) SetIncluded(v []IncidentServiceIncludedItems) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentServicesResponse) GetMeta() IncidentResponseMeta {
	if o == nil || o.Meta == nil {
		var ret IncidentResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponse) GetMetaOk() (*IncidentResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentServicesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given IncidentResponseMeta and assigns it to the Meta field.
func (o *IncidentServicesResponse) SetMeta(v IncidentResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentServicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentServicesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]IncidentServiceResponseData `json:"data"`
		Included []IncidentServiceIncludedItems `json:"included,omitempty"`
		Meta     *IncidentResponseMeta          `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	o.Included = all.Included
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
