# ListImagesResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | Token to use for paginated requests. | [optional] 
**Images** | [**[]ImageInfoSummary**](ImageInfoSummary.md) |  | 

## Methods

### NewListImagesResponseContent

`func NewListImagesResponseContent(images []ImageInfoSummary, ) *ListImagesResponseContent`

NewListImagesResponseContent instantiates a new ListImagesResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImagesResponseContentWithDefaults

`func NewListImagesResponseContentWithDefaults() *ListImagesResponseContent`

NewListImagesResponseContentWithDefaults instantiates a new ListImagesResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *ListImagesResponseContent) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListImagesResponseContent) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListImagesResponseContent) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListImagesResponseContent) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetImages

`func (o *ListImagesResponseContent) GetImages() []ImageInfoSummary`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ListImagesResponseContent) GetImagesOk() (*[]ImageInfoSummary, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ListImagesResponseContent) SetImages(v []ImageInfoSummary)`

SetImages sets Images field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


