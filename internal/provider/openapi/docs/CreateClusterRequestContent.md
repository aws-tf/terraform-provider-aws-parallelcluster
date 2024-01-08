# CreateClusterRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | Name of the cluster that will be created. | 
**ClusterConfiguration** | **string** | Cluster configuration as a YAML document. | 

## Methods

### NewCreateClusterRequestContent

`func NewCreateClusterRequestContent(clusterName string, clusterConfiguration string, ) *CreateClusterRequestContent`

NewCreateClusterRequestContent instantiates a new CreateClusterRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterRequestContentWithDefaults

`func NewCreateClusterRequestContentWithDefaults() *CreateClusterRequestContent`

NewCreateClusterRequestContentWithDefaults instantiates a new CreateClusterRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *CreateClusterRequestContent) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreateClusterRequestContent) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreateClusterRequestContent) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetClusterConfiguration

`func (o *CreateClusterRequestContent) GetClusterConfiguration() string`

GetClusterConfiguration returns the ClusterConfiguration field if non-nil, zero value otherwise.

### GetClusterConfigurationOk

`func (o *CreateClusterRequestContent) GetClusterConfigurationOk() (*string, bool)`

GetClusterConfigurationOk returns a tuple with the ClusterConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConfiguration

`func (o *CreateClusterRequestContent) SetClusterConfiguration(v string)`

SetClusterConfiguration sets ClusterConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


