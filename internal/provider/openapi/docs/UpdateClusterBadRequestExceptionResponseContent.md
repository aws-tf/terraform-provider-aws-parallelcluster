# UpdateClusterBadRequestExceptionResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**ConfigurationValidationErrors** | Pointer to [**[]ConfigValidationMessage**](ConfigValidationMessage.md) |  | [optional] 
**UpdateValidationErrors** | Pointer to [**[]UpdateError**](UpdateError.md) |  | [optional] 
**ChangeSet** | Pointer to [**[]Change**](Change.md) |  | [optional] 

## Methods

### NewUpdateClusterBadRequestExceptionResponseContent

`func NewUpdateClusterBadRequestExceptionResponseContent() *UpdateClusterBadRequestExceptionResponseContent`

NewUpdateClusterBadRequestExceptionResponseContent instantiates a new UpdateClusterBadRequestExceptionResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterBadRequestExceptionResponseContentWithDefaults

`func NewUpdateClusterBadRequestExceptionResponseContentWithDefaults() *UpdateClusterBadRequestExceptionResponseContent`

NewUpdateClusterBadRequestExceptionResponseContentWithDefaults instantiates a new UpdateClusterBadRequestExceptionResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateClusterBadRequestExceptionResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateClusterBadRequestExceptionResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateClusterBadRequestExceptionResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateClusterBadRequestExceptionResponseContent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetConfigurationValidationErrors

`func (o *UpdateClusterBadRequestExceptionResponseContent) GetConfigurationValidationErrors() []ConfigValidationMessage`

GetConfigurationValidationErrors returns the ConfigurationValidationErrors field if non-nil, zero value otherwise.

### GetConfigurationValidationErrorsOk

`func (o *UpdateClusterBadRequestExceptionResponseContent) GetConfigurationValidationErrorsOk() (*[]ConfigValidationMessage, bool)`

GetConfigurationValidationErrorsOk returns a tuple with the ConfigurationValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationValidationErrors

`func (o *UpdateClusterBadRequestExceptionResponseContent) SetConfigurationValidationErrors(v []ConfigValidationMessage)`

SetConfigurationValidationErrors sets ConfigurationValidationErrors field to given value.

### HasConfigurationValidationErrors

`func (o *UpdateClusterBadRequestExceptionResponseContent) HasConfigurationValidationErrors() bool`

HasConfigurationValidationErrors returns a boolean if a field has been set.

### GetUpdateValidationErrors

`func (o *UpdateClusterBadRequestExceptionResponseContent) GetUpdateValidationErrors() []UpdateError`

GetUpdateValidationErrors returns the UpdateValidationErrors field if non-nil, zero value otherwise.

### GetUpdateValidationErrorsOk

`func (o *UpdateClusterBadRequestExceptionResponseContent) GetUpdateValidationErrorsOk() (*[]UpdateError, bool)`

GetUpdateValidationErrorsOk returns a tuple with the UpdateValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateValidationErrors

`func (o *UpdateClusterBadRequestExceptionResponseContent) SetUpdateValidationErrors(v []UpdateError)`

SetUpdateValidationErrors sets UpdateValidationErrors field to given value.

### HasUpdateValidationErrors

`func (o *UpdateClusterBadRequestExceptionResponseContent) HasUpdateValidationErrors() bool`

HasUpdateValidationErrors returns a boolean if a field has been set.

### GetChangeSet

`func (o *UpdateClusterBadRequestExceptionResponseContent) GetChangeSet() []Change`

GetChangeSet returns the ChangeSet field if non-nil, zero value otherwise.

### GetChangeSetOk

`func (o *UpdateClusterBadRequestExceptionResponseContent) GetChangeSetOk() (*[]Change, bool)`

GetChangeSetOk returns a tuple with the ChangeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSet

`func (o *UpdateClusterBadRequestExceptionResponseContent) SetChangeSet(v []Change)`

SetChangeSet sets ChangeSet field to given value.

### HasChangeSet

`func (o *UpdateClusterBadRequestExceptionResponseContent) HasChangeSet() bool`

HasChangeSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


