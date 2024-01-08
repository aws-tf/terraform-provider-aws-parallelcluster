# BuildImageBadRequestExceptionResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**ConfigurationValidationErrors** | Pointer to [**[]ConfigValidationMessage**](ConfigValidationMessage.md) |  | [optional] 

## Methods

### NewBuildImageBadRequestExceptionResponseContent

`func NewBuildImageBadRequestExceptionResponseContent() *BuildImageBadRequestExceptionResponseContent`

NewBuildImageBadRequestExceptionResponseContent instantiates a new BuildImageBadRequestExceptionResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildImageBadRequestExceptionResponseContentWithDefaults

`func NewBuildImageBadRequestExceptionResponseContentWithDefaults() *BuildImageBadRequestExceptionResponseContent`

NewBuildImageBadRequestExceptionResponseContentWithDefaults instantiates a new BuildImageBadRequestExceptionResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BuildImageBadRequestExceptionResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BuildImageBadRequestExceptionResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BuildImageBadRequestExceptionResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BuildImageBadRequestExceptionResponseContent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetConfigurationValidationErrors

`func (o *BuildImageBadRequestExceptionResponseContent) GetConfigurationValidationErrors() []ConfigValidationMessage`

GetConfigurationValidationErrors returns the ConfigurationValidationErrors field if non-nil, zero value otherwise.

### GetConfigurationValidationErrorsOk

`func (o *BuildImageBadRequestExceptionResponseContent) GetConfigurationValidationErrorsOk() (*[]ConfigValidationMessage, bool)`

GetConfigurationValidationErrorsOk returns a tuple with the ConfigurationValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationValidationErrors

`func (o *BuildImageBadRequestExceptionResponseContent) SetConfigurationValidationErrors(v []ConfigValidationMessage)`

SetConfigurationValidationErrors sets ConfigurationValidationErrors field to given value.

### HasConfigurationValidationErrors

`func (o *BuildImageBadRequestExceptionResponseContent) HasConfigurationValidationErrors() bool`

HasConfigurationValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


