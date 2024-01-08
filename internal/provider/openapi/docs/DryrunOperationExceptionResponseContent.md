# DryrunOperationExceptionResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**ChangeSet** | Pointer to [**[]Change**](Change.md) | List of configuration changes requested by the operation. | [optional] 
**ValidationMessages** | Pointer to [**[]ConfigValidationMessage**](ConfigValidationMessage.md) | List of messages collected during cluster config validation whose level is lower than the &#39;validationFailureLevel&#39; set by the user. | [optional] 

## Methods

### NewDryrunOperationExceptionResponseContent

`func NewDryrunOperationExceptionResponseContent() *DryrunOperationExceptionResponseContent`

NewDryrunOperationExceptionResponseContent instantiates a new DryrunOperationExceptionResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDryrunOperationExceptionResponseContentWithDefaults

`func NewDryrunOperationExceptionResponseContentWithDefaults() *DryrunOperationExceptionResponseContent`

NewDryrunOperationExceptionResponseContentWithDefaults instantiates a new DryrunOperationExceptionResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DryrunOperationExceptionResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DryrunOperationExceptionResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DryrunOperationExceptionResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DryrunOperationExceptionResponseContent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetChangeSet

`func (o *DryrunOperationExceptionResponseContent) GetChangeSet() []Change`

GetChangeSet returns the ChangeSet field if non-nil, zero value otherwise.

### GetChangeSetOk

`func (o *DryrunOperationExceptionResponseContent) GetChangeSetOk() (*[]Change, bool)`

GetChangeSetOk returns a tuple with the ChangeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSet

`func (o *DryrunOperationExceptionResponseContent) SetChangeSet(v []Change)`

SetChangeSet sets ChangeSet field to given value.

### HasChangeSet

`func (o *DryrunOperationExceptionResponseContent) HasChangeSet() bool`

HasChangeSet returns a boolean if a field has been set.

### GetValidationMessages

`func (o *DryrunOperationExceptionResponseContent) GetValidationMessages() []ConfigValidationMessage`

GetValidationMessages returns the ValidationMessages field if non-nil, zero value otherwise.

### GetValidationMessagesOk

`func (o *DryrunOperationExceptionResponseContent) GetValidationMessagesOk() (*[]ConfigValidationMessage, bool)`

GetValidationMessagesOk returns a tuple with the ValidationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessages

`func (o *DryrunOperationExceptionResponseContent) SetValidationMessages(v []ConfigValidationMessage)`

SetValidationMessages sets ValidationMessages field to given value.

### HasValidationMessages

`func (o *DryrunOperationExceptionResponseContent) HasValidationMessages() bool`

HasValidationMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


