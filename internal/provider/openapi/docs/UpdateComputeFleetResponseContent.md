# UpdateComputeFleetResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ComputeFleetStatus**](ComputeFleetStatus.md) |  | 
**LastStatusUpdatedTime** | Pointer to **time.Time** | Timestamp representing the last status update time. | [optional] 

## Methods

### NewUpdateComputeFleetResponseContent

`func NewUpdateComputeFleetResponseContent(status ComputeFleetStatus, ) *UpdateComputeFleetResponseContent`

NewUpdateComputeFleetResponseContent instantiates a new UpdateComputeFleetResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateComputeFleetResponseContentWithDefaults

`func NewUpdateComputeFleetResponseContentWithDefaults() *UpdateComputeFleetResponseContent`

NewUpdateComputeFleetResponseContentWithDefaults instantiates a new UpdateComputeFleetResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateComputeFleetResponseContent) GetStatus() ComputeFleetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateComputeFleetResponseContent) GetStatusOk() (*ComputeFleetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateComputeFleetResponseContent) SetStatus(v ComputeFleetStatus)`

SetStatus sets Status field to given value.


### GetLastStatusUpdatedTime

`func (o *UpdateComputeFleetResponseContent) GetLastStatusUpdatedTime() time.Time`

GetLastStatusUpdatedTime returns the LastStatusUpdatedTime field if non-nil, zero value otherwise.

### GetLastStatusUpdatedTimeOk

`func (o *UpdateComputeFleetResponseContent) GetLastStatusUpdatedTimeOk() (*time.Time, bool)`

GetLastStatusUpdatedTimeOk returns a tuple with the LastStatusUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdatedTime

`func (o *UpdateComputeFleetResponseContent) SetLastStatusUpdatedTime(v time.Time)`

SetLastStatusUpdatedTime sets LastStatusUpdatedTime field to given value.

### HasLastStatusUpdatedTime

`func (o *UpdateComputeFleetResponseContent) HasLastStatusUpdatedTime() bool`

HasLastStatusUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


