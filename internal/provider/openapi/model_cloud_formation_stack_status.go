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

// CloudFormationStackStatus the model 'CloudFormationStackStatus'
type CloudFormationStackStatus string

// List of CloudFormationStackStatus
const (
	CLOUDFORMATIONSTACKSTATUS_CREATE_IN_PROGRESS CloudFormationStackStatus = "CREATE_IN_PROGRESS"
	CLOUDFORMATIONSTACKSTATUS_CREATE_FAILED CloudFormationStackStatus = "CREATE_FAILED"
	CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE CloudFormationStackStatus = "CREATE_COMPLETE"
	CLOUDFORMATIONSTACKSTATUS_ROLLBACK_IN_PROGRESS CloudFormationStackStatus = "ROLLBACK_IN_PROGRESS"
	CLOUDFORMATIONSTACKSTATUS_ROLLBACK_FAILED CloudFormationStackStatus = "ROLLBACK_FAILED"
	CLOUDFORMATIONSTACKSTATUS_ROLLBACK_COMPLETE CloudFormationStackStatus = "ROLLBACK_COMPLETE"
	CLOUDFORMATIONSTACKSTATUS_DELETE_IN_PROGRESS CloudFormationStackStatus = "DELETE_IN_PROGRESS"
	CLOUDFORMATIONSTACKSTATUS_DELETE_FAILED CloudFormationStackStatus = "DELETE_FAILED"
	CLOUDFORMATIONSTACKSTATUS_DELETE_COMPLETE CloudFormationStackStatus = "DELETE_COMPLETE"
	CLOUDFORMATIONSTACKSTATUS_UPDATE_IN_PROGRESS CloudFormationStackStatus = "UPDATE_IN_PROGRESS"
	CLOUDFORMATIONSTACKSTATUS_UPDATE_COMPLETE_CLEANUP_IN_PROGRESS CloudFormationStackStatus = "UPDATE_COMPLETE_CLEANUP_IN_PROGRESS"
	CLOUDFORMATIONSTACKSTATUS_UPDATE_COMPLETE CloudFormationStackStatus = "UPDATE_COMPLETE"
	CLOUDFORMATIONSTACKSTATUS_UPDATE_ROLLBACK_IN_PROGRESS CloudFormationStackStatus = "UPDATE_ROLLBACK_IN_PROGRESS"
	CLOUDFORMATIONSTACKSTATUS_UPDATE_ROLLBACK_FAILED CloudFormationStackStatus = "UPDATE_ROLLBACK_FAILED"
	CLOUDFORMATIONSTACKSTATUS_UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS CloudFormationStackStatus = "UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS"
	CLOUDFORMATIONSTACKSTATUS_UPDATE_ROLLBACK_COMPLETE CloudFormationStackStatus = "UPDATE_ROLLBACK_COMPLETE"
)

// All allowed values of CloudFormationStackStatus enum
var AllowedCloudFormationStackStatusEnumValues = []CloudFormationStackStatus{
	"CREATE_IN_PROGRESS",
	"CREATE_FAILED",
	"CREATE_COMPLETE",
	"ROLLBACK_IN_PROGRESS",
	"ROLLBACK_FAILED",
	"ROLLBACK_COMPLETE",
	"DELETE_IN_PROGRESS",
	"DELETE_FAILED",
	"DELETE_COMPLETE",
	"UPDATE_IN_PROGRESS",
	"UPDATE_COMPLETE_CLEANUP_IN_PROGRESS",
	"UPDATE_COMPLETE",
	"UPDATE_ROLLBACK_IN_PROGRESS",
	"UPDATE_ROLLBACK_FAILED",
	"UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS",
	"UPDATE_ROLLBACK_COMPLETE",
}

func (v *CloudFormationStackStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudFormationStackStatus(value)
	for _, existing := range AllowedCloudFormationStackStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudFormationStackStatus", value)
}

// NewCloudFormationStackStatusFromValue returns a pointer to a valid CloudFormationStackStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudFormationStackStatusFromValue(v string) (*CloudFormationStackStatus, error) {
	ev := CloudFormationStackStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudFormationStackStatus: valid values are %v", v, AllowedCloudFormationStackStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudFormationStackStatus) IsValid() bool {
	for _, existing := range AllowedCloudFormationStackStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudFormationStackStatus value
func (v CloudFormationStackStatus) Ptr() *CloudFormationStackStatus {
	return &v
}

type NullableCloudFormationStackStatus struct {
	value *CloudFormationStackStatus
	isSet bool
}

func (v NullableCloudFormationStackStatus) Get() *CloudFormationStackStatus {
	return v.value
}

func (v *NullableCloudFormationStackStatus) Set(val *CloudFormationStackStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFormationStackStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFormationStackStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFormationStackStatus(val *CloudFormationStackStatus) *NullableCloudFormationStackStatus {
	return &NullableCloudFormationStackStatus{value: val, isSet: true}
}

func (v NullableCloudFormationStackStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFormationStackStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
