# ListImageLogStreamsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | Token to use for paginated requests. | [optional] 
**LogStreams** | [**[]LogStream**](LogStream.md) |  | 

## Methods

### NewListImageLogStreamsResponseContent

`func NewListImageLogStreamsResponseContent(logStreams []LogStream, ) *ListImageLogStreamsResponseContent`

NewListImageLogStreamsResponseContent instantiates a new ListImageLogStreamsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageLogStreamsResponseContentWithDefaults

`func NewListImageLogStreamsResponseContentWithDefaults() *ListImageLogStreamsResponseContent`

NewListImageLogStreamsResponseContentWithDefaults instantiates a new ListImageLogStreamsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *ListImageLogStreamsResponseContent) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListImageLogStreamsResponseContent) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListImageLogStreamsResponseContent) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListImageLogStreamsResponseContent) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetLogStreams

`func (o *ListImageLogStreamsResponseContent) GetLogStreams() []LogStream`

GetLogStreams returns the LogStreams field if non-nil, zero value otherwise.

### GetLogStreamsOk

`func (o *ListImageLogStreamsResponseContent) GetLogStreamsOk() (*[]LogStream, bool)`

GetLogStreamsOk returns a tuple with the LogStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStreams

`func (o *ListImageLogStreamsResponseContent) SetLogStreams(v []LogStream)`

SetLogStreams sets LogStreams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


