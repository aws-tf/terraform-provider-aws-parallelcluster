# GetClusterLogEventsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | Token to use for paginated requests. | [optional] 
**PrevToken** | Pointer to **string** | Token to use for paginated requests. | [optional] 
**Events** | Pointer to [**[]LogEvent**](LogEvent.md) |  | [optional] 

## Methods

### NewGetClusterLogEventsResponseContent

`func NewGetClusterLogEventsResponseContent() *GetClusterLogEventsResponseContent`

NewGetClusterLogEventsResponseContent instantiates a new GetClusterLogEventsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterLogEventsResponseContentWithDefaults

`func NewGetClusterLogEventsResponseContentWithDefaults() *GetClusterLogEventsResponseContent`

NewGetClusterLogEventsResponseContentWithDefaults instantiates a new GetClusterLogEventsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *GetClusterLogEventsResponseContent) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetClusterLogEventsResponseContent) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetClusterLogEventsResponseContent) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetClusterLogEventsResponseContent) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetPrevToken

`func (o *GetClusterLogEventsResponseContent) GetPrevToken() string`

GetPrevToken returns the PrevToken field if non-nil, zero value otherwise.

### GetPrevTokenOk

`func (o *GetClusterLogEventsResponseContent) GetPrevTokenOk() (*string, bool)`

GetPrevTokenOk returns a tuple with the PrevToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevToken

`func (o *GetClusterLogEventsResponseContent) SetPrevToken(v string)`

SetPrevToken sets PrevToken field to given value.

### HasPrevToken

`func (o *GetClusterLogEventsResponseContent) HasPrevToken() bool`

HasPrevToken returns a boolean if a field has been set.

### GetEvents

`func (o *GetClusterLogEventsResponseContent) GetEvents() []LogEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetClusterLogEventsResponseContent) GetEventsOk() (*[]LogEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetClusterLogEventsResponseContent) SetEvents(v []LogEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetClusterLogEventsResponseContent) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


