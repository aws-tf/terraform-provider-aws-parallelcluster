# LoginNodesPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**LoginNodesState**](LoginNodesState.md) |  | 
**PoolName** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Scheme** | Pointer to **string** |  | [optional] 
**HealthyNodes** | Pointer to **int32** |  | [optional] 
**UnhealthyNodes** | Pointer to **int32** |  | [optional] 

## Methods

### NewLoginNodesPool

`func NewLoginNodesPool(status LoginNodesState, ) *LoginNodesPool`

NewLoginNodesPool instantiates a new LoginNodesPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginNodesPoolWithDefaults

`func NewLoginNodesPoolWithDefaults() *LoginNodesPool`

NewLoginNodesPoolWithDefaults instantiates a new LoginNodesPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LoginNodesPool) GetStatus() LoginNodesState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoginNodesPool) GetStatusOk() (*LoginNodesState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoginNodesPool) SetStatus(v LoginNodesState)`

SetStatus sets Status field to given value.


### GetPoolName

`func (o *LoginNodesPool) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *LoginNodesPool) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *LoginNodesPool) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *LoginNodesPool) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetAddress

`func (o *LoginNodesPool) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoginNodesPool) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoginNodesPool) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LoginNodesPool) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetScheme

`func (o *LoginNodesPool) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *LoginNodesPool) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *LoginNodesPool) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *LoginNodesPool) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetHealthyNodes

`func (o *LoginNodesPool) GetHealthyNodes() int32`

GetHealthyNodes returns the HealthyNodes field if non-nil, zero value otherwise.

### GetHealthyNodesOk

`func (o *LoginNodesPool) GetHealthyNodesOk() (*int32, bool)`

GetHealthyNodesOk returns a tuple with the HealthyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyNodes

`func (o *LoginNodesPool) SetHealthyNodes(v int32)`

SetHealthyNodes sets HealthyNodes field to given value.

### HasHealthyNodes

`func (o *LoginNodesPool) HasHealthyNodes() bool`

HasHealthyNodes returns a boolean if a field has been set.

### GetUnhealthyNodes

`func (o *LoginNodesPool) GetUnhealthyNodes() int32`

GetUnhealthyNodes returns the UnhealthyNodes field if non-nil, zero value otherwise.

### GetUnhealthyNodesOk

`func (o *LoginNodesPool) GetUnhealthyNodesOk() (*int32, bool)`

GetUnhealthyNodesOk returns a tuple with the UnhealthyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyNodes

`func (o *LoginNodesPool) SetUnhealthyNodes(v int32)`

SetUnhealthyNodes sets UnhealthyNodes field to given value.

### HasUnhealthyNodes

`func (o *LoginNodesPool) HasUnhealthyNodes() bool`

HasUnhealthyNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


