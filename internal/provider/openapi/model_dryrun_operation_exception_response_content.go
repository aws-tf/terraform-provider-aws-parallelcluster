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

// checks if the DryrunOperationExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DryrunOperationExceptionResponseContent{}

// DryrunOperationExceptionResponseContent Communicates that the operation would have succeeded without the dryrun flag.
type DryrunOperationExceptionResponseContent struct {
	Message *string `json:"message,omitempty"`
	// List of configuration changes requested by the operation.
	ChangeSet []Change `json:"changeSet,omitempty"`
	// List of messages collected during cluster config validation whose level is lower than the 'validationFailureLevel' set by the user.
	ValidationMessages []ConfigValidationMessage `json:"validationMessages,omitempty"`
}

// NewDryrunOperationExceptionResponseContent instantiates a new DryrunOperationExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDryrunOperationExceptionResponseContent() *DryrunOperationExceptionResponseContent {
	this := DryrunOperationExceptionResponseContent{}
	return &this
}

// NewDryrunOperationExceptionResponseContentWithDefaults instantiates a new DryrunOperationExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDryrunOperationExceptionResponseContentWithDefaults() *DryrunOperationExceptionResponseContent {
	this := DryrunOperationExceptionResponseContent{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DryrunOperationExceptionResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryrunOperationExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DryrunOperationExceptionResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DryrunOperationExceptionResponseContent) SetMessage(v string) {
	o.Message = &v
}

// GetChangeSet returns the ChangeSet field value if set, zero value otherwise.
func (o *DryrunOperationExceptionResponseContent) GetChangeSet() []Change {
	if o == nil || IsNil(o.ChangeSet) {
		var ret []Change
		return ret
	}
	return o.ChangeSet
}

// GetChangeSetOk returns a tuple with the ChangeSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryrunOperationExceptionResponseContent) GetChangeSetOk() ([]Change, bool) {
	if o == nil || IsNil(o.ChangeSet) {
		return nil, false
	}
	return o.ChangeSet, true
}

// HasChangeSet returns a boolean if a field has been set.
func (o *DryrunOperationExceptionResponseContent) HasChangeSet() bool {
	if o != nil && !IsNil(o.ChangeSet) {
		return true
	}

	return false
}

// SetChangeSet gets a reference to the given []Change and assigns it to the ChangeSet field.
func (o *DryrunOperationExceptionResponseContent) SetChangeSet(v []Change) {
	o.ChangeSet = v
}

// GetValidationMessages returns the ValidationMessages field value if set, zero value otherwise.
func (o *DryrunOperationExceptionResponseContent) GetValidationMessages() []ConfigValidationMessage {
	if o == nil || IsNil(o.ValidationMessages) {
		var ret []ConfigValidationMessage
		return ret
	}
	return o.ValidationMessages
}

// GetValidationMessagesOk returns a tuple with the ValidationMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryrunOperationExceptionResponseContent) GetValidationMessagesOk() ([]ConfigValidationMessage, bool) {
	if o == nil || IsNil(o.ValidationMessages) {
		return nil, false
	}
	return o.ValidationMessages, true
}

// HasValidationMessages returns a boolean if a field has been set.
func (o *DryrunOperationExceptionResponseContent) HasValidationMessages() bool {
	if o != nil && !IsNil(o.ValidationMessages) {
		return true
	}

	return false
}

// SetValidationMessages gets a reference to the given []ConfigValidationMessage and assigns it to the ValidationMessages field.
func (o *DryrunOperationExceptionResponseContent) SetValidationMessages(v []ConfigValidationMessage) {
	o.ValidationMessages = v
}

func (o DryrunOperationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DryrunOperationExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.ChangeSet) {
		toSerialize["changeSet"] = o.ChangeSet
	}
	if !IsNil(o.ValidationMessages) {
		toSerialize["validationMessages"] = o.ValidationMessages
	}
	return toSerialize, nil
}

type NullableDryrunOperationExceptionResponseContent struct {
	value *DryrunOperationExceptionResponseContent
	isSet bool
}

func (v NullableDryrunOperationExceptionResponseContent) Get() *DryrunOperationExceptionResponseContent {
	return v.value
}

func (v *NullableDryrunOperationExceptionResponseContent) Set(val *DryrunOperationExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDryrunOperationExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDryrunOperationExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDryrunOperationExceptionResponseContent(val *DryrunOperationExceptionResponseContent) *NullableDryrunOperationExceptionResponseContent {
	return &NullableDryrunOperationExceptionResponseContent{value: val, isSet: true}
}

func (v NullableDryrunOperationExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDryrunOperationExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

