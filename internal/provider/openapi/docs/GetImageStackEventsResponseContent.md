# GetImageStackEventsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | Token to use for paginated requests. | [optional] 
**Events** | Pointer to [**[]StackEvent**](StackEvent.md) |  | [optional] 

## Methods

### NewGetImageStackEventsResponseContent

`func NewGetImageStackEventsResponseContent() *GetImageStackEventsResponseContent`

NewGetImageStackEventsResponseContent instantiates a new GetImageStackEventsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageStackEventsResponseContentWithDefaults

`func NewGetImageStackEventsResponseContentWithDefaults() *GetImageStackEventsResponseContent`

NewGetImageStackEventsResponseContentWithDefaults instantiates a new GetImageStackEventsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *GetImageStackEventsResponseContent) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetImageStackEventsResponseContent) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetImageStackEventsResponseContent) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetImageStackEventsResponseContent) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetEvents

`func (o *GetImageStackEventsResponseContent) GetEvents() []StackEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetImageStackEventsResponseContent) GetEventsOk() (*[]StackEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetImageStackEventsResponseContent) SetEvents(v []StackEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetImageStackEventsResponseContent) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


