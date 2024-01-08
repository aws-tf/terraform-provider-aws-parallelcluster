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

// ImageBuilderImageStatus the model 'ImageBuilderImageStatus'
type ImageBuilderImageStatus string

// List of ImageBuilderImageStatus
const (
	IMAGEBUILDERIMAGESTATUS_PENDING ImageBuilderImageStatus = "PENDING"
	IMAGEBUILDERIMAGESTATUS_CREATING ImageBuilderImageStatus = "CREATING"
	IMAGEBUILDERIMAGESTATUS_BUILDING ImageBuilderImageStatus = "BUILDING"
	IMAGEBUILDERIMAGESTATUS_TESTING ImageBuilderImageStatus = "TESTING"
	IMAGEBUILDERIMAGESTATUS_DISTRIBUTING ImageBuilderImageStatus = "DISTRIBUTING"
	IMAGEBUILDERIMAGESTATUS_INTEGRATING ImageBuilderImageStatus = "INTEGRATING"
	IMAGEBUILDERIMAGESTATUS_AVAILABLE ImageBuilderImageStatus = "AVAILABLE"
	IMAGEBUILDERIMAGESTATUS_CANCELLED ImageBuilderImageStatus = "CANCELLED"
	IMAGEBUILDERIMAGESTATUS_FAILED ImageBuilderImageStatus = "FAILED"
	IMAGEBUILDERIMAGESTATUS_DEPRECATED ImageBuilderImageStatus = "DEPRECATED"
	IMAGEBUILDERIMAGESTATUS_DELETED ImageBuilderImageStatus = "DELETED"
)

// All allowed values of ImageBuilderImageStatus enum
var AllowedImageBuilderImageStatusEnumValues = []ImageBuilderImageStatus{
	"PENDING",
	"CREATING",
	"BUILDING",
	"TESTING",
	"DISTRIBUTING",
	"INTEGRATING",
	"AVAILABLE",
	"CANCELLED",
	"FAILED",
	"DEPRECATED",
	"DELETED",
}

func (v *ImageBuilderImageStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageBuilderImageStatus(value)
	for _, existing := range AllowedImageBuilderImageStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageBuilderImageStatus", value)
}

// NewImageBuilderImageStatusFromValue returns a pointer to a valid ImageBuilderImageStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageBuilderImageStatusFromValue(v string) (*ImageBuilderImageStatus, error) {
	ev := ImageBuilderImageStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageBuilderImageStatus: valid values are %v", v, AllowedImageBuilderImageStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageBuilderImageStatus) IsValid() bool {
	for _, existing := range AllowedImageBuilderImageStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageBuilderImageStatus value
func (v ImageBuilderImageStatus) Ptr() *ImageBuilderImageStatus {
	return &v
}

type NullableImageBuilderImageStatus struct {
	value *ImageBuilderImageStatus
	isSet bool
}

func (v NullableImageBuilderImageStatus) Get() *ImageBuilderImageStatus {
	return v.value
}

func (v *NullableImageBuilderImageStatus) Set(val *ImageBuilderImageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBuilderImageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBuilderImageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBuilderImageStatus(val *ImageBuilderImageStatus) *NullableImageBuilderImageStatus {
	return &NullableImageBuilderImageStatus{value: val, isSet: true}
}

func (v NullableImageBuilderImageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBuilderImageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

