# DescribeClusterInstancesResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | Token to use for paginated requests. | [optional] 
**Instances** | [**[]ClusterInstance**](ClusterInstance.md) |  | 

## Methods

### NewDescribeClusterInstancesResponseContent

`func NewDescribeClusterInstancesResponseContent(instances []ClusterInstance, ) *DescribeClusterInstancesResponseContent`

NewDescribeClusterInstancesResponseContent instantiates a new DescribeClusterInstancesResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeClusterInstancesResponseContentWithDefaults

`func NewDescribeClusterInstancesResponseContentWithDefaults() *DescribeClusterInstancesResponseContent`

NewDescribeClusterInstancesResponseContentWithDefaults instantiates a new DescribeClusterInstancesResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *DescribeClusterInstancesResponseContent) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *DescribeClusterInstancesResponseContent) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *DescribeClusterInstancesResponseContent) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *DescribeClusterInstancesResponseContent) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetInstances

`func (o *DescribeClusterInstancesResponseContent) GetInstances() []ClusterInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *DescribeClusterInstancesResponseContent) GetInstancesOk() (*[]ClusterInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *DescribeClusterInstancesResponseContent) SetInstances(v []ClusterInstance)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


