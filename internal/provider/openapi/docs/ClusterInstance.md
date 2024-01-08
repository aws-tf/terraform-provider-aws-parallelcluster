# ClusterInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** |  | 
**InstanceType** | **string** |  | 
**LaunchTime** | **time.Time** |  | 
**PrivateIpAddress** | **string** |  | 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**State** | [**InstanceState**](InstanceState.md) |  | 
**NodeType** | [**NodeType**](NodeType.md) |  | 
**QueueName** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterInstance

`func NewClusterInstance(instanceId string, instanceType string, launchTime time.Time, privateIpAddress string, state InstanceState, nodeType NodeType, ) *ClusterInstance`

NewClusterInstance instantiates a new ClusterInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInstanceWithDefaults

`func NewClusterInstanceWithDefaults() *ClusterInstance`

NewClusterInstanceWithDefaults instantiates a new ClusterInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *ClusterInstance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ClusterInstance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ClusterInstance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceType

`func (o *ClusterInstance) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterInstance) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterInstance) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetLaunchTime

`func (o *ClusterInstance) GetLaunchTime() time.Time`

GetLaunchTime returns the LaunchTime field if non-nil, zero value otherwise.

### GetLaunchTimeOk

`func (o *ClusterInstance) GetLaunchTimeOk() (*time.Time, bool)`

GetLaunchTimeOk returns a tuple with the LaunchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchTime

`func (o *ClusterInstance) SetLaunchTime(v time.Time)`

SetLaunchTime sets LaunchTime field to given value.


### GetPrivateIpAddress

`func (o *ClusterInstance) GetPrivateIpAddress() string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *ClusterInstance) GetPrivateIpAddressOk() (*string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *ClusterInstance) SetPrivateIpAddress(v string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.


### GetPublicIpAddress

`func (o *ClusterInstance) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *ClusterInstance) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *ClusterInstance) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *ClusterInstance) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetState

`func (o *ClusterInstance) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterInstance) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterInstance) SetState(v InstanceState)`

SetState sets State field to given value.


### GetNodeType

`func (o *ClusterInstance) GetNodeType() NodeType`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ClusterInstance) GetNodeTypeOk() (*NodeType, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ClusterInstance) SetNodeType(v NodeType)`

SetNodeType sets NodeType field to given value.


### GetQueueName

`func (o *ClusterInstance) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *ClusterInstance) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *ClusterInstance) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *ClusterInstance) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


