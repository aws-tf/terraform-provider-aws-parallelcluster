/*
ParallelCluster

ParallelCluster API

API version: 3.11.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ComputeFleetStatus the model 'ComputeFleetStatus'
type ComputeFleetStatus string

// List of ComputeFleetStatus
const (
	COMPUTEFLEETSTATUS_START_REQUESTED ComputeFleetStatus = "START_REQUESTED"
	COMPUTEFLEETSTATUS_STARTING ComputeFleetStatus = "STARTING"
	COMPUTEFLEETSTATUS_RUNNING ComputeFleetStatus = "RUNNING"
	COMPUTEFLEETSTATUS_PROTECTED ComputeFleetStatus = "PROTECTED"
	COMPUTEFLEETSTATUS_STOP_REQUESTED ComputeFleetStatus = "STOP_REQUESTED"
	COMPUTEFLEETSTATUS_STOPPING ComputeFleetStatus = "STOPPING"
	COMPUTEFLEETSTATUS_STOPPED ComputeFleetStatus = "STOPPED"
	COMPUTEFLEETSTATUS_UNKNOWN ComputeFleetStatus = "UNKNOWN"
	COMPUTEFLEETSTATUS_ENABLED ComputeFleetStatus = "ENABLED"
	COMPUTEFLEETSTATUS_DISABLED ComputeFleetStatus = "DISABLED"
)

// All allowed values of ComputeFleetStatus enum
var AllowedComputeFleetStatusEnumValues = []ComputeFleetStatus{
	"START_REQUESTED",
	"STARTING",
	"RUNNING",
	"PROTECTED",
	"STOP_REQUESTED",
	"STOPPING",
	"STOPPED",
	"UNKNOWN",
	"ENABLED",
	"DISABLED",
}

func (v *ComputeFleetStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComputeFleetStatus(value)
	for _, existing := range AllowedComputeFleetStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComputeFleetStatus", value)
}

// NewComputeFleetStatusFromValue returns a pointer to a valid ComputeFleetStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputeFleetStatusFromValue(v string) (*ComputeFleetStatus, error) {
	ev := ComputeFleetStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ComputeFleetStatus: valid values are %v", v, AllowedComputeFleetStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputeFleetStatus) IsValid() bool {
	for _, existing := range AllowedComputeFleetStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComputeFleetStatus value
func (v ComputeFleetStatus) Ptr() *ComputeFleetStatus {
	return &v
}

type NullableComputeFleetStatus struct {
	value *ComputeFleetStatus
	isSet bool
}

func (v NullableComputeFleetStatus) Get() *ComputeFleetStatus {
	return v.value
}

func (v *NullableComputeFleetStatus) Set(val *ComputeFleetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeFleetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeFleetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeFleetStatus(val *ComputeFleetStatus) *NullableComputeFleetStatus {
	return &NullableComputeFleetStatus{value: val, isSet: true}
}

func (v NullableComputeFleetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeFleetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

