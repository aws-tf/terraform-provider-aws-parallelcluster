# CreateClusterResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**ClusterInfoSummary**](ClusterInfoSummary.md) |  | 
**ValidationMessages** | Pointer to [**[]ConfigValidationMessage**](ConfigValidationMessage.md) | List of messages collected during cluster config validation whose level is lower than the &#39;validationFailureLevel&#39; set by the user. | [optional] 

## Methods

### NewCreateClusterResponseContent

`func NewCreateClusterResponseContent(cluster ClusterInfoSummary, ) *CreateClusterResponseContent`

NewCreateClusterResponseContent instantiates a new CreateClusterResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterResponseContentWithDefaults

`func NewCreateClusterResponseContentWithDefaults() *CreateClusterResponseContent`

NewCreateClusterResponseContentWithDefaults instantiates a new CreateClusterResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *CreateClusterResponseContent) GetCluster() ClusterInfoSummary`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CreateClusterResponseContent) GetClusterOk() (*ClusterInfoSummary, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CreateClusterResponseContent) SetCluster(v ClusterInfoSummary)`

SetCluster sets Cluster field to given value.


### GetValidationMessages

`func (o *CreateClusterResponseContent) GetValidationMessages() []ConfigValidationMessage`

GetValidationMessages returns the ValidationMessages field if non-nil, zero value otherwise.

### GetValidationMessagesOk

`func (o *CreateClusterResponseContent) GetValidationMessagesOk() (*[]ConfigValidationMessage, bool)`

GetValidationMessagesOk returns a tuple with the ValidationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessages

`func (o *CreateClusterResponseContent) SetValidationMessages(v []ConfigValidationMessage)`

SetValidationMessages sets ValidationMessages field to given value.

### HasValidationMessages

`func (o *CreateClusterResponseContent) HasValidationMessages() bool`

HasValidationMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


