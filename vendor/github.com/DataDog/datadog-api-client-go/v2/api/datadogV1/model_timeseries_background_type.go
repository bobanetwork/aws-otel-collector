// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesBackgroundType Timeseries is made using an area or bars.
type TimeseriesBackgroundType string

// List of TimeseriesBackgroundType.
const (
	TIMESERIESBACKGROUNDTYPE_BARS TimeseriesBackgroundType = "bars"
	TIMESERIESBACKGROUNDTYPE_AREA TimeseriesBackgroundType = "area"
)

var allowedTimeseriesBackgroundTypeEnumValues = []TimeseriesBackgroundType{
	TIMESERIESBACKGROUNDTYPE_BARS,
	TIMESERIESBACKGROUNDTYPE_AREA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesBackgroundType) GetAllowedValues() []TimeseriesBackgroundType {
	return allowedTimeseriesBackgroundTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesBackgroundType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesBackgroundType(value)
	return nil
}

// NewTimeseriesBackgroundTypeFromValue returns a pointer to a valid TimeseriesBackgroundType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesBackgroundTypeFromValue(v string) (*TimeseriesBackgroundType, error) {
	ev := TimeseriesBackgroundType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesBackgroundType: valid values are %v", v, allowedTimeseriesBackgroundTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesBackgroundType) IsValid() bool {
	for _, existing := range allowedTimeseriesBackgroundTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesBackgroundType value.
func (v TimeseriesBackgroundType) Ptr() *TimeseriesBackgroundType {
	return &v
}
