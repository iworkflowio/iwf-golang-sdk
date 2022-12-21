/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
	"fmt"
)

// DeciderTriggerType the model 'DeciderTriggerType'
type DeciderTriggerType string

// List of DeciderTriggerType
const (
	ALL_COMMAND_COMPLETED DeciderTriggerType = "ALL_COMMAND_COMPLETED"
	ANY_COMMAND_COMPLETED DeciderTriggerType = "ANY_COMMAND_COMPLETED"
)

// All allowed values of DeciderTriggerType enum
var AllowedDeciderTriggerTypeEnumValues = []DeciderTriggerType{
	"ALL_COMMAND_COMPLETED",
	"ANY_COMMAND_COMPLETED",
}

func (v *DeciderTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeciderTriggerType(value)
	for _, existing := range AllowedDeciderTriggerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeciderTriggerType", value)
}

// NewDeciderTriggerTypeFromValue returns a pointer to a valid DeciderTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeciderTriggerTypeFromValue(v string) (*DeciderTriggerType, error) {
	ev := DeciderTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeciderTriggerType: valid values are %v", v, AllowedDeciderTriggerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeciderTriggerType) IsValid() bool {
	for _, existing := range AllowedDeciderTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeciderTriggerType value
func (v DeciderTriggerType) Ptr() *DeciderTriggerType {
	return &v
}

type NullableDeciderTriggerType struct {
	value *DeciderTriggerType
	isSet bool
}

func (v NullableDeciderTriggerType) Get() *DeciderTriggerType {
	return v.value
}

func (v *NullableDeciderTriggerType) Set(val *DeciderTriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeciderTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeciderTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeciderTriggerType(val *DeciderTriggerType) *NullableDeciderTriggerType {
	return &NullableDeciderTriggerType{value: val, isSet: true}
}

func (v NullableDeciderTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeciderTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

