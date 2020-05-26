/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SLOTimeframe The SLO time window options.
type SLOTimeframe string

// List of SLOTimeframe
const (
	SLOTIMEFRAME_SEVEN_DAYS  SLOTimeframe = "7d"
	SLOTIMEFRAME_THIRTY_DAYS SLOTimeframe = "30d"
	SLOTIMEFRAME_NINETY_DAYS SLOTimeframe = "90d"
)

func (v *SLOTimeframe) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SLOTimeframe(value)
	for _, existing := range []SLOTimeframe{"7d", "30d", "90d"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SLOTimeframe", *v)
}

// Ptr returns reference to SLOTimeframe value
func (v SLOTimeframe) Ptr() *SLOTimeframe {
	return &v
}

type NullableSLOTimeframe struct {
	value *SLOTimeframe
	isSet bool
}

func (v NullableSLOTimeframe) Get() *SLOTimeframe {
	return v.value
}

func (v *NullableSLOTimeframe) Set(val *SLOTimeframe) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOTimeframe) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOTimeframe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOTimeframe(val *SLOTimeframe) *NullableSLOTimeframe {
	return &NullableSLOTimeframe{value: val, isSet: true}
}

func (v NullableSLOTimeframe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOTimeframe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
