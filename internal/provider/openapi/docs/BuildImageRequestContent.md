# BuildImageRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageConfiguration** | **string** | Image configuration as a YAML document. | 
**ImageId** | **string** | Id of the Image that will be built. | 

## Methods

### NewBuildImageRequestContent

`func NewBuildImageRequestContent(imageConfiguration string, imageId string, ) *BuildImageRequestContent`

NewBuildImageRequestContent instantiates a new BuildImageRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildImageRequestContentWithDefaults

`func NewBuildImageRequestContentWithDefaults() *BuildImageRequestContent`

NewBuildImageRequestContentWithDefaults instantiates a new BuildImageRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageConfiguration

`func (o *BuildImageRequestContent) GetImageConfiguration() string`

GetImageConfiguration returns the ImageConfiguration field if non-nil, zero value otherwise.

### GetImageConfigurationOk

`func (o *BuildImageRequestContent) GetImageConfigurationOk() (*string, bool)`

GetImageConfigurationOk returns a tuple with the ImageConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfiguration

`func (o *BuildImageRequestContent) SetImageConfiguration(v string)`

SetImageConfiguration sets ImageConfiguration field to given value.


### GetImageId

`func (o *BuildImageRequestContent) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *BuildImageRequestContent) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *BuildImageRequestContent) SetImageId(v string)`

SetImageId sets ImageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


