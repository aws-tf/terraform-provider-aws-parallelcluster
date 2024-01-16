/*
ParallelCluster

ParallelCluster API

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RequestedComputeFleetStatus the model 'RequestedComputeFleetStatus'
type RequestedComputeFleetStatus string

// List of RequestedComputeFleetStatus
const (
	REQUESTEDCOMPUTEFLEETSTATUS_START_REQUESTED RequestedComputeFleetStatus = "START_REQUESTED"
	REQUESTEDCOMPUTEFLEETSTATUS_STOP_REQUESTED RequestedComputeFleetStatus = "STOP_REQUESTED"
	REQUESTEDCOMPUTEFLEETSTATUS_ENABLED RequestedComputeFleetStatus = "ENABLED"
	REQUESTEDCOMPUTEFLEETSTATUS_DISABLED RequestedComputeFleetStatus = "DISABLED"
)

// All allowed values of RequestedComputeFleetStatus enum
var AllowedRequestedComputeFleetStatusEnumValues = []RequestedComputeFleetStatus{
	"START_REQUESTED",
	"STOP_REQUESTED",
	"ENABLED",
	"DISABLED",
}

func (v *RequestedComputeFleetStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestedComputeFleetStatus(value)
	for _, existing := range AllowedRequestedComputeFleetStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestedComputeFleetStatus", value)
}

// NewRequestedComputeFleetStatusFromValue returns a pointer to a valid RequestedComputeFleetStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestedComputeFleetStatusFromValue(v string) (*RequestedComputeFleetStatus, error) {
	ev := RequestedComputeFleetStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestedComputeFleetStatus: valid values are %v", v, AllowedRequestedComputeFleetStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestedComputeFleetStatus) IsValid() bool {
	for _, existing := range AllowedRequestedComputeFleetStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestedComputeFleetStatus value
func (v RequestedComputeFleetStatus) Ptr() *RequestedComputeFleetStatus {
	return &v
}

type NullableRequestedComputeFleetStatus struct {
	value *RequestedComputeFleetStatus
	isSet bool
}

func (v NullableRequestedComputeFleetStatus) Get() *RequestedComputeFleetStatus {
	return v.value
}

func (v *NullableRequestedComputeFleetStatus) Set(val *RequestedComputeFleetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedComputeFleetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedComputeFleetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedComputeFleetStatus(val *RequestedComputeFleetStatus) *NullableRequestedComputeFleetStatus {
	return &NullableRequestedComputeFleetStatus{value: val, isSet: true}
}

func (v NullableRequestedComputeFleetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedComputeFleetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
