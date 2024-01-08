# LogStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogStreamName** | **string** | Name of the log stream. | 
**CreationTime** | **time.Time** | The creation time of the stream. | 
**FirstEventTimestamp** | **time.Time** | The time of the first event of the stream. | 
**LastEventTimestamp** | **time.Time** | The time of the last event of the stream. The lastEventTime value updates on an eventual consistency basis. It typically updates in less than an hour from ingestion, but in rare situations might take longer. | 
**LastIngestionTime** | **time.Time** | The last ingestion time. | 
**UploadSequenceToken** | **string** | The sequence token. | 
**LogStreamArn** | **string** | The Amazon Resource Name (ARN) of the log stream. | 

## Methods

### NewLogStream

`func NewLogStream(logStreamName string, creationTime time.Time, firstEventTimestamp time.Time, lastEventTimestamp time.Time, lastIngestionTime time.Time, uploadSequenceToken string, logStreamArn string, ) *LogStream`

NewLogStream instantiates a new LogStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamWithDefaults

`func NewLogStreamWithDefaults() *LogStream`

NewLogStreamWithDefaults instantiates a new LogStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogStreamName

`func (o *LogStream) GetLogStreamName() string`

GetLogStreamName returns the LogStreamName field if non-nil, zero value otherwise.

### GetLogStreamNameOk

`func (o *LogStream) GetLogStreamNameOk() (*string, bool)`

GetLogStreamNameOk returns a tuple with the LogStreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStreamName

`func (o *LogStream) SetLogStreamName(v string)`

SetLogStreamName sets LogStreamName field to given value.


### GetCreationTime

`func (o *LogStream) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *LogStream) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *LogStream) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetFirstEventTimestamp

`func (o *LogStream) GetFirstEventTimestamp() time.Time`

GetFirstEventTimestamp returns the FirstEventTimestamp field if non-nil, zero value otherwise.

### GetFirstEventTimestampOk

`func (o *LogStream) GetFirstEventTimestampOk() (*time.Time, bool)`

GetFirstEventTimestampOk returns a tuple with the FirstEventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEventTimestamp

`func (o *LogStream) SetFirstEventTimestamp(v time.Time)`

SetFirstEventTimestamp sets FirstEventTimestamp field to given value.


### GetLastEventTimestamp

`func (o *LogStream) GetLastEventTimestamp() time.Time`

GetLastEventTimestamp returns the LastEventTimestamp field if non-nil, zero value otherwise.

### GetLastEventTimestampOk

`func (o *LogStream) GetLastEventTimestampOk() (*time.Time, bool)`

GetLastEventTimestampOk returns a tuple with the LastEventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventTimestamp

`func (o *LogStream) SetLastEventTimestamp(v time.Time)`

SetLastEventTimestamp sets LastEventTimestamp field to given value.


### GetLastIngestionTime

`func (o *LogStream) GetLastIngestionTime() time.Time`

GetLastIngestionTime returns the LastIngestionTime field if non-nil, zero value otherwise.

### GetLastIngestionTimeOk

`func (o *LogStream) GetLastIngestionTimeOk() (*time.Time, bool)`

GetLastIngestionTimeOk returns a tuple with the LastIngestionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIngestionTime

`func (o *LogStream) SetLastIngestionTime(v time.Time)`

SetLastIngestionTime sets LastIngestionTime field to given value.


### GetUploadSequenceToken

`func (o *LogStream) GetUploadSequenceToken() string`

GetUploadSequenceToken returns the UploadSequenceToken field if non-nil, zero value otherwise.

### GetUploadSequenceTokenOk

`func (o *LogStream) GetUploadSequenceTokenOk() (*string, bool)`

GetUploadSequenceTokenOk returns a tuple with the UploadSequenceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSequenceToken

`func (o *LogStream) SetUploadSequenceToken(v string)`

SetUploadSequenceToken sets UploadSequenceToken field to given value.


### GetLogStreamArn

`func (o *LogStream) GetLogStreamArn() string`

GetLogStreamArn returns the LogStreamArn field if non-nil, zero value otherwise.

### GetLogStreamArnOk

`func (o *LogStream) GetLogStreamArnOk() (*string, bool)`

GetLogStreamArnOk returns a tuple with the LogStreamArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStreamArn

`func (o *LogStream) SetLogStreamArn(v string)`

SetLogStreamArn sets LogStreamArn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


