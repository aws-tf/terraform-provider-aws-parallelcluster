# StackEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **string** | The unique ID name of the instance of the stack. | 
**EventId** | **string** | The unique ID of this event. | 
**StackName** | **string** | The name associated with a stack. | 
**LogicalResourceId** | **string** | The logical name of the resource specified in the template. | 
**PhysicalResourceId** | **string** | The name or unique identifier associated with the physical instance of the resource. | 
**ResourceType** | **string** | Type of resource. | 
**Timestamp** | **time.Time** | Time the status was updated. | 
**ResourceStatus** | [**CloudFormationResourceStatus**](CloudFormationResourceStatus.md) |  | 
**ResourceStatusReason** | Pointer to **string** | Success/failure message associated with the resource. | [optional] 
**ResourceProperties** | Pointer to **string** | BLOB of the properties used to create the resource. | [optional] 
**ClientRequestToken** | Pointer to **string** | The token passed to the operation that generated this event. | [optional] 

## Methods

### NewStackEvent

`func NewStackEvent(stackId string, eventId string, stackName string, logicalResourceId string, physicalResourceId string, resourceType string, timestamp time.Time, resourceStatus CloudFormationResourceStatus, ) *StackEvent`

NewStackEvent instantiates a new StackEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackEventWithDefaults

`func NewStackEventWithDefaults() *StackEvent`

NewStackEventWithDefaults instantiates a new StackEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *StackEvent) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *StackEvent) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *StackEvent) SetStackId(v string)`

SetStackId sets StackId field to given value.


### GetEventId

`func (o *StackEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *StackEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *StackEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetStackName

`func (o *StackEvent) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *StackEvent) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *StackEvent) SetStackName(v string)`

SetStackName sets StackName field to given value.


### GetLogicalResourceId

`func (o *StackEvent) GetLogicalResourceId() string`

GetLogicalResourceId returns the LogicalResourceId field if non-nil, zero value otherwise.

### GetLogicalResourceIdOk

`func (o *StackEvent) GetLogicalResourceIdOk() (*string, bool)`

GetLogicalResourceIdOk returns a tuple with the LogicalResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalResourceId

`func (o *StackEvent) SetLogicalResourceId(v string)`

SetLogicalResourceId sets LogicalResourceId field to given value.


### GetPhysicalResourceId

`func (o *StackEvent) GetPhysicalResourceId() string`

GetPhysicalResourceId returns the PhysicalResourceId field if non-nil, zero value otherwise.

### GetPhysicalResourceIdOk

`func (o *StackEvent) GetPhysicalResourceIdOk() (*string, bool)`

GetPhysicalResourceIdOk returns a tuple with the PhysicalResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalResourceId

`func (o *StackEvent) SetPhysicalResourceId(v string)`

SetPhysicalResourceId sets PhysicalResourceId field to given value.


### GetResourceType

`func (o *StackEvent) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *StackEvent) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *StackEvent) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetTimestamp

`func (o *StackEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StackEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StackEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetResourceStatus

`func (o *StackEvent) GetResourceStatus() CloudFormationResourceStatus`

GetResourceStatus returns the ResourceStatus field if non-nil, zero value otherwise.

### GetResourceStatusOk

`func (o *StackEvent) GetResourceStatusOk() (*CloudFormationResourceStatus, bool)`

GetResourceStatusOk returns a tuple with the ResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceStatus

`func (o *StackEvent) SetResourceStatus(v CloudFormationResourceStatus)`

SetResourceStatus sets ResourceStatus field to given value.


### GetResourceStatusReason

`func (o *StackEvent) GetResourceStatusReason() string`

GetResourceStatusReason returns the ResourceStatusReason field if non-nil, zero value otherwise.

### GetResourceStatusReasonOk

`func (o *StackEvent) GetResourceStatusReasonOk() (*string, bool)`

GetResourceStatusReasonOk returns a tuple with the ResourceStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceStatusReason

`func (o *StackEvent) SetResourceStatusReason(v string)`

SetResourceStatusReason sets ResourceStatusReason field to given value.

### HasResourceStatusReason

`func (o *StackEvent) HasResourceStatusReason() bool`

HasResourceStatusReason returns a boolean if a field has been set.

### GetResourceProperties

`func (o *StackEvent) GetResourceProperties() string`

GetResourceProperties returns the ResourceProperties field if non-nil, zero value otherwise.

### GetResourcePropertiesOk

`func (o *StackEvent) GetResourcePropertiesOk() (*string, bool)`

GetResourcePropertiesOk returns a tuple with the ResourceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProperties

`func (o *StackEvent) SetResourceProperties(v string)`

SetResourceProperties sets ResourceProperties field to given value.

### HasResourceProperties

`func (o *StackEvent) HasResourceProperties() bool`

HasResourceProperties returns a boolean if a field has been set.

### GetClientRequestToken

`func (o *StackEvent) GetClientRequestToken() string`

GetClientRequestToken returns the ClientRequestToken field if non-nil, zero value otherwise.

### GetClientRequestTokenOk

`func (o *StackEvent) GetClientRequestTokenOk() (*string, bool)`

GetClientRequestTokenOk returns a tuple with the ClientRequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRequestToken

`func (o *StackEvent) SetClientRequestToken(v string)`

SetClientRequestToken sets ClientRequestToken field to given value.

### HasClientRequestToken

`func (o *StackEvent) HasClientRequestToken() bool`

HasClientRequestToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


