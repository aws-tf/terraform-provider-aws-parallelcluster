# UpdateClusterResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**ClusterInfoSummary**](ClusterInfoSummary.md) |  | 
**ValidationMessages** | Pointer to [**[]ConfigValidationMessage**](ConfigValidationMessage.md) | List of messages collected during cluster config validation whose level is lower than the &#39;validationFailureLevel&#39; set by the user. | [optional] 
**ChangeSet** | [**[]Change**](Change.md) | List of configuration changes requested by the update operation. | 

## Methods

### NewUpdateClusterResponseContent

`func NewUpdateClusterResponseContent(cluster ClusterInfoSummary, changeSet []Change, ) *UpdateClusterResponseContent`

NewUpdateClusterResponseContent instantiates a new UpdateClusterResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterResponseContentWithDefaults

`func NewUpdateClusterResponseContentWithDefaults() *UpdateClusterResponseContent`

NewUpdateClusterResponseContentWithDefaults instantiates a new UpdateClusterResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *UpdateClusterResponseContent) GetCluster() ClusterInfoSummary`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpdateClusterResponseContent) GetClusterOk() (*ClusterInfoSummary, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpdateClusterResponseContent) SetCluster(v ClusterInfoSummary)`

SetCluster sets Cluster field to given value.


### GetValidationMessages

`func (o *UpdateClusterResponseContent) GetValidationMessages() []ConfigValidationMessage`

GetValidationMessages returns the ValidationMessages field if non-nil, zero value otherwise.

### GetValidationMessagesOk

`func (o *UpdateClusterResponseContent) GetValidationMessagesOk() (*[]ConfigValidationMessage, bool)`

GetValidationMessagesOk returns a tuple with the ValidationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessages

`func (o *UpdateClusterResponseContent) SetValidationMessages(v []ConfigValidationMessage)`

SetValidationMessages sets ValidationMessages field to given value.

### HasValidationMessages

`func (o *UpdateClusterResponseContent) HasValidationMessages() bool`

HasValidationMessages returns a boolean if a field has been set.

### GetChangeSet

`func (o *UpdateClusterResponseContent) GetChangeSet() []Change`

GetChangeSet returns the ChangeSet field if non-nil, zero value otherwise.

### GetChangeSetOk

`func (o *UpdateClusterResponseContent) GetChangeSetOk() (*[]Change, bool)`

GetChangeSetOk returns a tuple with the ChangeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSet

`func (o *UpdateClusterResponseContent) SetChangeSet(v []Change)`

SetChangeSet sets ChangeSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


