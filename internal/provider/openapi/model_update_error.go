/*
ParallelCluster

ParallelCluster API

API version: 3.11.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateError{}

// UpdateError struct for UpdateError
type UpdateError struct {
	Parameter *string `json:"parameter,omitempty"`
	CurrentValue *string `json:"currentValue,omitempty"`
	RequestedValue *string `json:"requestedValue,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewUpdateError instantiates a new UpdateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateError() *UpdateError {
	this := UpdateError{}
	return &this
}

// NewUpdateErrorWithDefaults instantiates a new UpdateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateErrorWithDefaults() *UpdateError {
	this := UpdateError{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *UpdateError) GetParameter() string {
	if o == nil || IsNil(o.Parameter) {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateError) GetParameterOk() (*string, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *UpdateError) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *UpdateError) SetParameter(v string) {
	o.Parameter = &v
}

// GetCurrentValue returns the CurrentValue field value if set, zero value otherwise.
func (o *UpdateError) GetCurrentValue() string {
	if o == nil || IsNil(o.CurrentValue) {
		var ret string
		return ret
	}
	return *o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateError) GetCurrentValueOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentValue) {
		return nil, false
	}
	return o.CurrentValue, true
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *UpdateError) HasCurrentValue() bool {
	if o != nil && !IsNil(o.CurrentValue) {
		return true
	}

	return false
}

// SetCurrentValue gets a reference to the given string and assigns it to the CurrentValue field.
func (o *UpdateError) SetCurrentValue(v string) {
	o.CurrentValue = &v
}

// GetRequestedValue returns the RequestedValue field value if set, zero value otherwise.
func (o *UpdateError) GetRequestedValue() string {
	if o == nil || IsNil(o.RequestedValue) {
		var ret string
		return ret
	}
	return *o.RequestedValue
}

// GetRequestedValueOk returns a tuple with the RequestedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateError) GetRequestedValueOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedValue) {
		return nil, false
	}
	return o.RequestedValue, true
}

// HasRequestedValue returns a boolean if a field has been set.
func (o *UpdateError) HasRequestedValue() bool {
	if o != nil && !IsNil(o.RequestedValue) {
		return true
	}

	return false
}

// SetRequestedValue gets a reference to the given string and assigns it to the RequestedValue field.
func (o *UpdateError) SetRequestedValue(v string) {
	o.RequestedValue = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateError) SetMessage(v string) {
	o.Message = &v
}

func (o UpdateError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.CurrentValue) {
		toSerialize["currentValue"] = o.CurrentValue
	}
	if !IsNil(o.RequestedValue) {
		toSerialize["requestedValue"] = o.RequestedValue
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableUpdateError struct {
	value *UpdateError
	isSet bool
}

func (v NullableUpdateError) Get() *UpdateError {
	return v.value
}

func (v *NullableUpdateError) Set(val *UpdateError) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateError) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateError(val *UpdateError) *NullableUpdateError {
	return &NullableUpdateError{value: val, isSet: true}
}

func (v NullableUpdateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


