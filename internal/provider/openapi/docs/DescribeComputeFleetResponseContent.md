# DescribeComputeFleetResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ComputeFleetStatus**](ComputeFleetStatus.md) |  | 
**LastStatusUpdatedTime** | Pointer to **time.Time** | Timestamp representing the last status update time. | [optional] 

## Methods

### NewDescribeComputeFleetResponseContent

`func NewDescribeComputeFleetResponseContent(status ComputeFleetStatus, ) *DescribeComputeFleetResponseContent`

NewDescribeComputeFleetResponseContent instantiates a new DescribeComputeFleetResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeComputeFleetResponseContentWithDefaults

`func NewDescribeComputeFleetResponseContentWithDefaults() *DescribeComputeFleetResponseContent`

NewDescribeComputeFleetResponseContentWithDefaults instantiates a new DescribeComputeFleetResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DescribeComputeFleetResponseContent) GetStatus() ComputeFleetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeComputeFleetResponseContent) GetStatusOk() (*ComputeFleetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeComputeFleetResponseContent) SetStatus(v ComputeFleetStatus)`

SetStatus sets Status field to given value.


### GetLastStatusUpdatedTime

`func (o *DescribeComputeFleetResponseContent) GetLastStatusUpdatedTime() time.Time`

GetLastStatusUpdatedTime returns the LastStatusUpdatedTime field if non-nil, zero value otherwise.

### GetLastStatusUpdatedTimeOk

`func (o *DescribeComputeFleetResponseContent) GetLastStatusUpdatedTimeOk() (*time.Time, bool)`

GetLastStatusUpdatedTimeOk returns a tuple with the LastStatusUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdatedTime

`func (o *DescribeComputeFleetResponseContent) SetLastStatusUpdatedTime(v time.Time)`

SetLastStatusUpdatedTime sets LastStatusUpdatedTime field to given value.

### HasLastStatusUpdatedTime

`func (o *DescribeComputeFleetResponseContent) HasLastStatusUpdatedTime() bool`

HasLastStatusUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


