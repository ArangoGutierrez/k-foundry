# V0040OpenapiAccountsAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | Pointer to [**V0040AccountsAddCond**](V0040AccountsAddCond.md) |  | [optional] 
**Account** | Pointer to [**V0040AccountShort**](V0040AccountShort.md) |  | [optional] 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiAccountsAddCondResp

`func NewV0040OpenapiAccountsAddCondResp() *V0040OpenapiAccountsAddCondResp`

NewV0040OpenapiAccountsAddCondResp instantiates a new V0040OpenapiAccountsAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiAccountsAddCondRespWithDefaults

`func NewV0040OpenapiAccountsAddCondRespWithDefaults() *V0040OpenapiAccountsAddCondResp`

NewV0040OpenapiAccountsAddCondRespWithDefaults instantiates a new V0040OpenapiAccountsAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0040OpenapiAccountsAddCondResp) GetAssociationCondition() V0040AccountsAddCond`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0040OpenapiAccountsAddCondResp) GetAssociationConditionOk() (*V0040AccountsAddCond, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0040OpenapiAccountsAddCondResp) SetAssociationCondition(v V0040AccountsAddCond)`

SetAssociationCondition sets AssociationCondition field to given value.

### HasAssociationCondition

`func (o *V0040OpenapiAccountsAddCondResp) HasAssociationCondition() bool`

HasAssociationCondition returns a boolean if a field has been set.

### GetAccount

`func (o *V0040OpenapiAccountsAddCondResp) GetAccount() V0040AccountShort`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0040OpenapiAccountsAddCondResp) GetAccountOk() (*V0040AccountShort, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0040OpenapiAccountsAddCondResp) SetAccount(v V0040AccountShort)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0040OpenapiAccountsAddCondResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetMeta

`func (o *V0040OpenapiAccountsAddCondResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiAccountsAddCondResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiAccountsAddCondResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiAccountsAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiAccountsAddCondResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiAccountsAddCondResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiAccountsAddCondResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiAccountsAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiAccountsAddCondResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiAccountsAddCondResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiAccountsAddCondResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiAccountsAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


