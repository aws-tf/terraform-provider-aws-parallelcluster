# EC2Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** |  | 
**InstanceType** | **string** |  | 
**LaunchTime** | **time.Time** |  | 
**PrivateIpAddress** | **string** |  | 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**State** | [**InstanceState**](InstanceState.md) |  | 

## Methods

### NewEC2Instance

`func NewEC2Instance(instanceId string, instanceType string, launchTime time.Time, privateIpAddress string, state InstanceState, ) *EC2Instance`

NewEC2Instance instantiates a new EC2Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEC2InstanceWithDefaults

`func NewEC2InstanceWithDefaults() *EC2Instance`

NewEC2InstanceWithDefaults instantiates a new EC2Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *EC2Instance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *EC2Instance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *EC2Instance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceType

`func (o *EC2Instance) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *EC2Instance) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *EC2Instance) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetLaunchTime

`func (o *EC2Instance) GetLaunchTime() time.Time`

GetLaunchTime returns the LaunchTime field if non-nil, zero value otherwise.

### GetLaunchTimeOk

`func (o *EC2Instance) GetLaunchTimeOk() (*time.Time, bool)`

GetLaunchTimeOk returns a tuple with the LaunchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchTime

`func (o *EC2Instance) SetLaunchTime(v time.Time)`

SetLaunchTime sets LaunchTime field to given value.


### GetPrivateIpAddress

`func (o *EC2Instance) GetPrivateIpAddress() string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *EC2Instance) GetPrivateIpAddressOk() (*string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *EC2Instance) SetPrivateIpAddress(v string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.


### GetPublicIpAddress

`func (o *EC2Instance) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *EC2Instance) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *EC2Instance) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *EC2Instance) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetState

`func (o *EC2Instance) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EC2Instance) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EC2Instance) SetState(v InstanceState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


