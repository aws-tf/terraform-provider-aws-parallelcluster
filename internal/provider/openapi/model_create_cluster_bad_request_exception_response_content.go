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

// checks if the CreateClusterBadRequestExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClusterBadRequestExceptionResponseContent{}

// CreateClusterBadRequestExceptionResponseContent This exception is thrown when a client calls the CreateCluster API with an invalid request. This includes an error due to invalid cluster configuration.
type CreateClusterBadRequestExceptionResponseContent struct {
	Message *string `json:"message,omitempty"`
	ConfigurationValidationErrors []ConfigValidationMessage `json:"configurationValidationErrors,omitempty"`
}

// NewCreateClusterBadRequestExceptionResponseContent instantiates a new CreateClusterBadRequestExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterBadRequestExceptionResponseContent() *CreateClusterBadRequestExceptionResponseContent {
	this := CreateClusterBadRequestExceptionResponseContent{}
	return &this
}

// NewCreateClusterBadRequestExceptionResponseContentWithDefaults instantiates a new CreateClusterBadRequestExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterBadRequestExceptionResponseContentWithDefaults() *CreateClusterBadRequestExceptionResponseContent {
	this := CreateClusterBadRequestExceptionResponseContent{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateClusterBadRequestExceptionResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterBadRequestExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateClusterBadRequestExceptionResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateClusterBadRequestExceptionResponseContent) SetMessage(v string) {
	o.Message = &v
}

// GetConfigurationValidationErrors returns the ConfigurationValidationErrors field value if set, zero value otherwise.
func (o *CreateClusterBadRequestExceptionResponseContent) GetConfigurationValidationErrors() []ConfigValidationMessage {
	if o == nil || IsNil(o.ConfigurationValidationErrors) {
		var ret []ConfigValidationMessage
		return ret
	}
	return o.ConfigurationValidationErrors
}

// GetConfigurationValidationErrorsOk returns a tuple with the ConfigurationValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterBadRequestExceptionResponseContent) GetConfigurationValidationErrorsOk() ([]ConfigValidationMessage, bool) {
	if o == nil || IsNil(o.ConfigurationValidationErrors) {
		return nil, false
	}
	return o.ConfigurationValidationErrors, true
}

// HasConfigurationValidationErrors returns a boolean if a field has been set.
func (o *CreateClusterBadRequestExceptionResponseContent) HasConfigurationValidationErrors() bool {
	if o != nil && !IsNil(o.ConfigurationValidationErrors) {
		return true
	}

	return false
}

// SetConfigurationValidationErrors gets a reference to the given []ConfigValidationMessage and assigns it to the ConfigurationValidationErrors field.
func (o *CreateClusterBadRequestExceptionResponseContent) SetConfigurationValidationErrors(v []ConfigValidationMessage) {
	o.ConfigurationValidationErrors = v
}

func (o CreateClusterBadRequestExceptionResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClusterBadRequestExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.ConfigurationValidationErrors) {
		toSerialize["configurationValidationErrors"] = o.ConfigurationValidationErrors
	}
	return toSerialize, nil
}

type NullableCreateClusterBadRequestExceptionResponseContent struct {
	value *CreateClusterBadRequestExceptionResponseContent
	isSet bool
}

func (v NullableCreateClusterBadRequestExceptionResponseContent) Get() *CreateClusterBadRequestExceptionResponseContent {
	return v.value
}

func (v *NullableCreateClusterBadRequestExceptionResponseContent) Set(val *CreateClusterBadRequestExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterBadRequestExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterBadRequestExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterBadRequestExceptionResponseContent(val *CreateClusterBadRequestExceptionResponseContent) *NullableCreateClusterBadRequestExceptionResponseContent {
	return &NullableCreateClusterBadRequestExceptionResponseContent{value: val, isSet: true}
}

func (v NullableCreateClusterBadRequestExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterBadRequestExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

