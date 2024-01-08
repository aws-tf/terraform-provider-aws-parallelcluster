/*
ParallelCluster

ParallelCluster API

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnauthorizedClientErrorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnauthorizedClientErrorResponseContent{}

// UnauthorizedClientErrorResponseContent This exception is thrown when the client is not authorized to perform an action.
type UnauthorizedClientErrorResponseContent struct {
	Message *string `json:"message,omitempty"`
}

// NewUnauthorizedClientErrorResponseContent instantiates a new UnauthorizedClientErrorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthorizedClientErrorResponseContent() *UnauthorizedClientErrorResponseContent {
	this := UnauthorizedClientErrorResponseContent{}
	return &this
}

// NewUnauthorizedClientErrorResponseContentWithDefaults instantiates a new UnauthorizedClientErrorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthorizedClientErrorResponseContentWithDefaults() *UnauthorizedClientErrorResponseContent {
	this := UnauthorizedClientErrorResponseContent{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UnauthorizedClientErrorResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnauthorizedClientErrorResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UnauthorizedClientErrorResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UnauthorizedClientErrorResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o UnauthorizedClientErrorResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnauthorizedClientErrorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableUnauthorizedClientErrorResponseContent struct {
	value *UnauthorizedClientErrorResponseContent
	isSet bool
}

func (v NullableUnauthorizedClientErrorResponseContent) Get() *UnauthorizedClientErrorResponseContent {
	return v.value
}

func (v *NullableUnauthorizedClientErrorResponseContent) Set(val *UnauthorizedClientErrorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorizedClientErrorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorizedClientErrorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorizedClientErrorResponseContent(val *UnauthorizedClientErrorResponseContent) *NullableUnauthorizedClientErrorResponseContent {
	return &NullableUnauthorizedClientErrorResponseContent{value: val, isSet: true}
}

func (v NullableUnauthorizedClientErrorResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnauthorizedClientErrorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


