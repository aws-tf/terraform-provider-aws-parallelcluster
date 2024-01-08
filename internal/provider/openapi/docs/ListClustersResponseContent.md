# ListClustersResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | Token to use for paginated requests. | [optional] 
**Clusters** | [**[]ClusterInfoSummary**](ClusterInfoSummary.md) |  | 

## Methods

### NewListClustersResponseContent

`func NewListClustersResponseContent(clusters []ClusterInfoSummary, ) *ListClustersResponseContent`

NewListClustersResponseContent instantiates a new ListClustersResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClustersResponseContentWithDefaults

`func NewListClustersResponseContentWithDefaults() *ListClustersResponseContent`

NewListClustersResponseContentWithDefaults instantiates a new ListClustersResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *ListClustersResponseContent) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListClustersResponseContent) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListClustersResponseContent) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListClustersResponseContent) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetClusters

`func (o *ListClustersResponseContent) GetClusters() []ClusterInfoSummary`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ListClustersResponseContent) GetClustersOk() (*[]ClusterInfoSummary, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ListClustersResponseContent) SetClusters(v []ClusterInfoSummary)`

SetClusters sets Clusters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


