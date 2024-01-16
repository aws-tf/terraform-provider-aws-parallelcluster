/*
ParallelCluster

ParallelCluster API

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the EC2Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EC2Instance{}

// EC2Instance struct for EC2Instance
type EC2Instance struct {
	InstanceId string `json:"instanceId"`
	InstanceType string `json:"instanceType"`
	LaunchTime time.Time `json:"launchTime"`
	PrivateIpAddress string `json:"privateIpAddress"`
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
	State InstanceState `json:"state"`
}

type _EC2Instance EC2Instance

// NewEC2Instance instantiates a new EC2Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEC2Instance(instanceId string, instanceType string, launchTime time.Time, privateIpAddress string, state InstanceState) *EC2Instance {
	this := EC2Instance{}
	this.InstanceId = instanceId
	this.InstanceType = instanceType
	this.LaunchTime = launchTime
	this.PrivateIpAddress = privateIpAddress
	this.State = state
	return &this
}

// NewEC2InstanceWithDefaults instantiates a new EC2Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEC2InstanceWithDefaults() *EC2Instance {
	this := EC2Instance{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *EC2Instance) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *EC2Instance) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *EC2Instance) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetInstanceType returns the InstanceType field value
func (o *EC2Instance) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *EC2Instance) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *EC2Instance) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetLaunchTime returns the LaunchTime field value
func (o *EC2Instance) GetLaunchTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LaunchTime
}

// GetLaunchTimeOk returns a tuple with the LaunchTime field value
// and a boolean to check if the value has been set.
func (o *EC2Instance) GetLaunchTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LaunchTime, true
}

// SetLaunchTime sets field value
func (o *EC2Instance) SetLaunchTime(v time.Time) {
	o.LaunchTime = v
}

// GetPrivateIpAddress returns the PrivateIpAddress field value
func (o *EC2Instance) GetPrivateIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateIpAddress
}

// GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field value
// and a boolean to check if the value has been set.
func (o *EC2Instance) GetPrivateIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateIpAddress, true
}

// SetPrivateIpAddress sets field value
func (o *EC2Instance) SetPrivateIpAddress(v string) {
	o.PrivateIpAddress = v
}

// GetPublicIpAddress returns the PublicIpAddress field value if set, zero value otherwise.
func (o *EC2Instance) GetPublicIpAddress() string {
	if o == nil || IsNil(o.PublicIpAddress) {
		var ret string
		return ret
	}
	return *o.PublicIpAddress
}

// GetPublicIpAddressOk returns a tuple with the PublicIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EC2Instance) GetPublicIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIpAddress) {
		return nil, false
	}
	return o.PublicIpAddress, true
}

// HasPublicIpAddress returns a boolean if a field has been set.
func (o *EC2Instance) HasPublicIpAddress() bool {
	if o != nil && !IsNil(o.PublicIpAddress) {
		return true
	}

	return false
}

// SetPublicIpAddress gets a reference to the given string and assigns it to the PublicIpAddress field.
func (o *EC2Instance) SetPublicIpAddress(v string) {
	o.PublicIpAddress = &v
}

// GetState returns the State field value
func (o *EC2Instance) GetState() InstanceState {
	if o == nil {
		var ret InstanceState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *EC2Instance) GetStateOk() (*InstanceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *EC2Instance) SetState(v InstanceState) {
	o.State = v
}

func (o EC2Instance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EC2Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["instanceType"] = o.InstanceType
	toSerialize["launchTime"] = o.LaunchTime
	toSerialize["privateIpAddress"] = o.PrivateIpAddress
	if !IsNil(o.PublicIpAddress) {
		toSerialize["publicIpAddress"] = o.PublicIpAddress
	}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *EC2Instance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instanceId",
		"instanceType",
		"launchTime",
		"privateIpAddress",
		"state",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEC2Instance := _EC2Instance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEC2Instance)

	if err != nil {
		return err
	}

	*o = EC2Instance(varEC2Instance)

	return err
}

type NullableEC2Instance struct {
	value *EC2Instance
	isSet bool
}

func (v NullableEC2Instance) Get() *EC2Instance {
	return v.value
}

func (v *NullableEC2Instance) Set(val *EC2Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableEC2Instance) IsSet() bool {
	return v.isSet
}

func (v *NullableEC2Instance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEC2Instance(val *EC2Instance) *NullableEC2Instance {
	return &NullableEC2Instance{value: val, isSet: true}
}

func (v NullableEC2Instance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEC2Instance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

