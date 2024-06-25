/*
ParallelCluster

ParallelCluster API

API version: 3.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetImageStackEventsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetImageStackEventsResponseContent{}

// GetImageStackEventsResponseContent struct for GetImageStackEventsResponseContent
type GetImageStackEventsResponseContent struct {
	// Token to use for paginated requests.
	NextToken *string `json:"nextToken,omitempty"`
	Events []StackEvent `json:"events,omitempty"`
}

// NewGetImageStackEventsResponseContent instantiates a new GetImageStackEventsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetImageStackEventsResponseContent() *GetImageStackEventsResponseContent {
	this := GetImageStackEventsResponseContent{}
	return &this
}

// NewGetImageStackEventsResponseContentWithDefaults instantiates a new GetImageStackEventsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetImageStackEventsResponseContentWithDefaults() *GetImageStackEventsResponseContent {
	this := GetImageStackEventsResponseContent{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetImageStackEventsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetImageStackEventsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetImageStackEventsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetImageStackEventsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GetImageStackEventsResponseContent) GetEvents() []StackEvent {
	if o == nil || IsNil(o.Events) {
		var ret []StackEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetImageStackEventsResponseContent) GetEventsOk() ([]StackEvent, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GetImageStackEventsResponseContent) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StackEvent and assigns it to the Events field.
func (o *GetImageStackEventsResponseContent) SetEvents(v []StackEvent) {
	o.Events = v
}

func (o GetImageStackEventsResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetImageStackEventsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableGetImageStackEventsResponseContent struct {
	value *GetImageStackEventsResponseContent
	isSet bool
}

func (v NullableGetImageStackEventsResponseContent) Get() *GetImageStackEventsResponseContent {
	return v.value
}

func (v *NullableGetImageStackEventsResponseContent) Set(val *GetImageStackEventsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetImageStackEventsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetImageStackEventsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetImageStackEventsResponseContent(val *GetImageStackEventsResponseContent) *NullableGetImageStackEventsResponseContent {
	return &NullableGetImageStackEventsResponseContent{value: val, isSet: true}
}

func (v NullableGetImageStackEventsResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetImageStackEventsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
