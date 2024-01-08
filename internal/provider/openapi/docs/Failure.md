# Failure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCode** | Pointer to **string** | Failure code when the cluster stack is in CREATE_FAILED status. | [optional] 
**FailureReason** | Pointer to **string** | Failure reason when the cluster stack is in CREATE_FAILED status. | [optional] 

## Methods

### NewFailure

`func NewFailure() *Failure`

NewFailure instantiates a new Failure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureWithDefaults

`func NewFailureWithDefaults() *Failure`

NewFailureWithDefaults instantiates a new Failure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCode

`func (o *Failure) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *Failure) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *Failure) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *Failure) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetFailureReason

`func (o *Failure) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *Failure) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *Failure) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *Failure) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


