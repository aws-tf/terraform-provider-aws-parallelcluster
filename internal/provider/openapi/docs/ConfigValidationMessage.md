# ConfigValidationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the validator. | [optional] 
**Type** | Pointer to **string** | Type of the validator. | [optional] 
**Level** | Pointer to [**ValidationLevel**](ValidationLevel.md) |  | [optional] 
**Message** | Pointer to **string** | Validation message | [optional] 

## Methods

### NewConfigValidationMessage

`func NewConfigValidationMessage() *ConfigValidationMessage`

NewConfigValidationMessage instantiates a new ConfigValidationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigValidationMessageWithDefaults

`func NewConfigValidationMessageWithDefaults() *ConfigValidationMessage`

NewConfigValidationMessageWithDefaults instantiates a new ConfigValidationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigValidationMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigValidationMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigValidationMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigValidationMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ConfigValidationMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigValidationMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigValidationMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigValidationMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *ConfigValidationMessage) GetLevel() ValidationLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ConfigValidationMessage) GetLevelOk() (*ValidationLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ConfigValidationMessage) SetLevel(v ValidationLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ConfigValidationMessage) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigValidationMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigValidationMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigValidationMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigValidationMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


