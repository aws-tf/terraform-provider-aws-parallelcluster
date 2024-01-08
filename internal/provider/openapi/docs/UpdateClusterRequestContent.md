# UpdateClusterRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterConfiguration** | **string** | Cluster configuration as a YAML document. | 

## Methods

### NewUpdateClusterRequestContent

`func NewUpdateClusterRequestContent(clusterConfiguration string, ) *UpdateClusterRequestContent`

NewUpdateClusterRequestContent instantiates a new UpdateClusterRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterRequestContentWithDefaults

`func NewUpdateClusterRequestContentWithDefaults() *UpdateClusterRequestContent`

NewUpdateClusterRequestContentWithDefaults instantiates a new UpdateClusterRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterConfiguration

`func (o *UpdateClusterRequestContent) GetClusterConfiguration() string`

GetClusterConfiguration returns the ClusterConfiguration field if non-nil, zero value otherwise.

### GetClusterConfigurationOk

`func (o *UpdateClusterRequestContent) GetClusterConfigurationOk() (*string, bool)`

GetClusterConfigurationOk returns a tuple with the ClusterConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConfiguration

`func (o *UpdateClusterRequestContent) SetClusterConfiguration(v string)`

SetClusterConfiguration sets ClusterConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


