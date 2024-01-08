# LogEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**Message** | **string** |  | 

## Methods

### NewLogEvent

`func NewLogEvent(timestamp time.Time, message string, ) *LogEvent`

NewLogEvent instantiates a new LogEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEventWithDefaults

`func NewLogEventWithDefaults() *LogEvent`

NewLogEventWithDefaults instantiates a new LogEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *LogEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMessage

`func (o *LogEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogEvent) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


