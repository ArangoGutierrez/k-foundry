# V0040OpenapiUsersAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | Pointer to [**V0040UsersAddCond**](V0040UsersAddCond.md) |  | [optional] 
**User** | Pointer to [**V0040UserShort**](V0040UserShort.md) |  | [optional] 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiUsersAddCondResp

`func NewV0040OpenapiUsersAddCondResp() *V0040OpenapiUsersAddCondResp`

NewV0040OpenapiUsersAddCondResp instantiates a new V0040OpenapiUsersAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiUsersAddCondRespWithDefaults

`func NewV0040OpenapiUsersAddCondRespWithDefaults() *V0040OpenapiUsersAddCondResp`

NewV0040OpenapiUsersAddCondRespWithDefaults instantiates a new V0040OpenapiUsersAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0040OpenapiUsersAddCondResp) GetAssociationCondition() V0040UsersAddCond`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0040OpenapiUsersAddCondResp) GetAssociationConditionOk() (*V0040UsersAddCond, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0040OpenapiUsersAddCondResp) SetAssociationCondition(v V0040UsersAddCond)`

SetAssociationCondition sets AssociationCondition field to given value.

### HasAssociationCondition

`func (o *V0040OpenapiUsersAddCondResp) HasAssociationCondition() bool`

HasAssociationCondition returns a boolean if a field has been set.

### GetUser

`func (o *V0040OpenapiUsersAddCondResp) GetUser() V0040UserShort`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0040OpenapiUsersAddCondResp) GetUserOk() (*V0040UserShort, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0040OpenapiUsersAddCondResp) SetUser(v V0040UserShort)`

SetUser sets User field to given value.

### HasUser

`func (o *V0040OpenapiUsersAddCondResp) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMeta

`func (o *V0040OpenapiUsersAddCondResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiUsersAddCondResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiUsersAddCondResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiUsersAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiUsersAddCondResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiUsersAddCondResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiUsersAddCondResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiUsersAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiUsersAddCondResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiUsersAddCondResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiUsersAddCondResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiUsersAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


