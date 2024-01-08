# ListClusterLogStreamsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | Token to use for paginated requests. | [optional] 
**LogStreams** | [**[]LogStream**](LogStream.md) |  | 

## Methods

### NewListClusterLogStreamsResponseContent

`func NewListClusterLogStreamsResponseContent(logStreams []LogStream, ) *ListClusterLogStreamsResponseContent`

NewListClusterLogStreamsResponseContent instantiates a new ListClusterLogStreamsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterLogStreamsResponseContentWithDefaults

`func NewListClusterLogStreamsResponseContentWithDefaults() *ListClusterLogStreamsResponseContent`

NewListClusterLogStreamsResponseContentWithDefaults instantiates a new ListClusterLogStreamsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *ListClusterLogStreamsResponseContent) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListClusterLogStreamsResponseContent) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListClusterLogStreamsResponseContent) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListClusterLogStreamsResponseContent) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetLogStreams

`func (o *ListClusterLogStreamsResponseContent) GetLogStreams() []LogStream`

GetLogStreams returns the LogStreams field if non-nil, zero value otherwise.

### GetLogStreamsOk

`func (o *ListClusterLogStreamsResponseContent) GetLogStreamsOk() (*[]LogStream, bool)`

GetLogStreamsOk returns a tuple with the LogStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStreams

`func (o *ListClusterLogStreamsResponseContent) SetLogStreams(v []LogStream)`

SetLogStreams sets LogStreams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


