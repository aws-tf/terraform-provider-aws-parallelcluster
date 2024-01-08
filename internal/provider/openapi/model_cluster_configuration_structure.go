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

// checks if the ClusterConfigurationStructure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterConfigurationStructure{}

// ClusterConfigurationStructure struct for ClusterConfigurationStructure
type ClusterConfigurationStructure struct {
	// URL of the cluster configuration file.
	Url *string `json:"url,omitempty"`
}

// NewClusterConfigurationStructure instantiates a new ClusterConfigurationStructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfigurationStructure() *ClusterConfigurationStructure {
	this := ClusterConfigurationStructure{}
	return &this
}

// NewClusterConfigurationStructureWithDefaults instantiates a new ClusterConfigurationStructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigurationStructureWithDefaults() *ClusterConfigurationStructure {
	this := ClusterConfigurationStructure{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ClusterConfigurationStructure) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationStructure) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ClusterConfigurationStructure) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ClusterConfigurationStructure) SetUrl(v string) {
	o.Url = &v
}

func (o ClusterConfigurationStructure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterConfigurationStructure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableClusterConfigurationStructure struct {
	value *ClusterConfigurationStructure
	isSet bool
}

func (v NullableClusterConfigurationStructure) Get() *ClusterConfigurationStructure {
	return v.value
}

func (v *NullableClusterConfigurationStructure) Set(val *ClusterConfigurationStructure) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterConfigurationStructure) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterConfigurationStructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterConfigurationStructure(val *ClusterConfigurationStructure) *NullableClusterConfigurationStructure {
	return &NullableClusterConfigurationStructure{value: val, isSet: true}
}

func (v NullableClusterConfigurationStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterConfigurationStructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


