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

// checks if the ConfigValidationMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigValidationMessage{}

// ConfigValidationMessage struct for ConfigValidationMessage
type ConfigValidationMessage struct {
	// Id of the validator.
	Id *string `json:"id,omitempty"`
	// Type of the validator.
	Type *string `json:"type,omitempty"`
	Level *ValidationLevel `json:"level,omitempty"`
	// Validation message
	Message *string `json:"message,omitempty"`
}

// NewConfigValidationMessage instantiates a new ConfigValidationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigValidationMessage() *ConfigValidationMessage {
	this := ConfigValidationMessage{}
	return &this
}

// NewConfigValidationMessageWithDefaults instantiates a new ConfigValidationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigValidationMessageWithDefaults() *ConfigValidationMessage {
	this := ConfigValidationMessage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigValidationMessage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigValidationMessage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigValidationMessage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigValidationMessage) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConfigValidationMessage) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigValidationMessage) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConfigValidationMessage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConfigValidationMessage) SetType(v string) {
	o.Type = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ConfigValidationMessage) GetLevel() ValidationLevel {
	if o == nil || IsNil(o.Level) {
		var ret ValidationLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigValidationMessage) GetLevelOk() (*ValidationLevel, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ConfigValidationMessage) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given ValidationLevel and assigns it to the Level field.
func (o *ConfigValidationMessage) SetLevel(v ValidationLevel) {
	o.Level = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConfigValidationMessage) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigValidationMessage) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConfigValidationMessage) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConfigValidationMessage) SetMessage(v string) {
	o.Message = &v
}

func (o ConfigValidationMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigValidationMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableConfigValidationMessage struct {
	value *ConfigValidationMessage
	isSet bool
}

func (v NullableConfigValidationMessage) Get() *ConfigValidationMessage {
	return v.value
}

func (v *NullableConfigValidationMessage) Set(val *ConfigValidationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigValidationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigValidationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigValidationMessage(val *ConfigValidationMessage) *NullableConfigValidationMessage {
	return &NullableConfigValidationMessage{value: val, isSet: true}
}

func (v NullableConfigValidationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigValidationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


