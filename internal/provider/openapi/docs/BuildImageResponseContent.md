# BuildImageResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | [**ImageInfoSummary**](ImageInfoSummary.md) |  | 
**ValidationMessages** | Pointer to [**[]ConfigValidationMessage**](ConfigValidationMessage.md) | List of messages collected during image config validation whose level is lower than the &#39;validationFailureLevel&#39; set by the user. | [optional] 

## Methods

### NewBuildImageResponseContent

`func NewBuildImageResponseContent(image ImageInfoSummary, ) *BuildImageResponseContent`

NewBuildImageResponseContent instantiates a new BuildImageResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildImageResponseContentWithDefaults

`func NewBuildImageResponseContentWithDefaults() *BuildImageResponseContent`

NewBuildImageResponseContentWithDefaults instantiates a new BuildImageResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *BuildImageResponseContent) GetImage() ImageInfoSummary`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BuildImageResponseContent) GetImageOk() (*ImageInfoSummary, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BuildImageResponseContent) SetImage(v ImageInfoSummary)`

SetImage sets Image field to given value.


### GetValidationMessages

`func (o *BuildImageResponseContent) GetValidationMessages() []ConfigValidationMessage`

GetValidationMessages returns the ValidationMessages field if non-nil, zero value otherwise.

### GetValidationMessagesOk

`func (o *BuildImageResponseContent) GetValidationMessagesOk() (*[]ConfigValidationMessage, bool)`

GetValidationMessagesOk returns a tuple with the ValidationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessages

`func (o *BuildImageResponseContent) SetValidationMessages(v []ConfigValidationMessage)`

SetValidationMessages sets ValidationMessages field to given value.

### HasValidationMessages

`func (o *BuildImageResponseContent) HasValidationMessages() bool`

HasValidationMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


