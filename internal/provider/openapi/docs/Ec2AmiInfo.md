# Ec2AmiInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiId** | **string** | EC2 AMI id | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | EC2 AMI Tags | [optional] 
**AmiName** | Pointer to **string** | EC2 AMI name | [optional] 
**Architecture** | Pointer to **string** | EC2 AMI architecture | [optional] 
**State** | Pointer to [**Ec2AmiState**](Ec2AmiState.md) |  | [optional] 
**Description** | Pointer to **string** | EC2 AMI description | [optional] 

## Methods

### NewEc2AmiInfo

`func NewEc2AmiInfo(amiId string, ) *Ec2AmiInfo`

NewEc2AmiInfo instantiates a new Ec2AmiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEc2AmiInfoWithDefaults

`func NewEc2AmiInfoWithDefaults() *Ec2AmiInfo`

NewEc2AmiInfoWithDefaults instantiates a new Ec2AmiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmiId

`func (o *Ec2AmiInfo) GetAmiId() string`

GetAmiId returns the AmiId field if non-nil, zero value otherwise.

### GetAmiIdOk

`func (o *Ec2AmiInfo) GetAmiIdOk() (*string, bool)`

GetAmiIdOk returns a tuple with the AmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiId

`func (o *Ec2AmiInfo) SetAmiId(v string)`

SetAmiId sets AmiId field to given value.


### GetTags

`func (o *Ec2AmiInfo) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Ec2AmiInfo) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Ec2AmiInfo) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Ec2AmiInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAmiName

`func (o *Ec2AmiInfo) GetAmiName() string`

GetAmiName returns the AmiName field if non-nil, zero value otherwise.

### GetAmiNameOk

`func (o *Ec2AmiInfo) GetAmiNameOk() (*string, bool)`

GetAmiNameOk returns a tuple with the AmiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiName

`func (o *Ec2AmiInfo) SetAmiName(v string)`

SetAmiName sets AmiName field to given value.

### HasAmiName

`func (o *Ec2AmiInfo) HasAmiName() bool`

HasAmiName returns a boolean if a field has been set.

### GetArchitecture

`func (o *Ec2AmiInfo) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Ec2AmiInfo) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Ec2AmiInfo) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *Ec2AmiInfo) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetState

`func (o *Ec2AmiInfo) GetState() Ec2AmiState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Ec2AmiInfo) GetStateOk() (*Ec2AmiState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Ec2AmiInfo) SetState(v Ec2AmiState)`

SetState sets State field to given value.

### HasState

`func (o *Ec2AmiInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDescription

`func (o *Ec2AmiInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ec2AmiInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Ec2AmiInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Ec2AmiInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


