# V0040OpenapiAccountsAddCondRespStr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAccounts** | **string** | added_accounts | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiAccountsAddCondRespStr

`func NewV0040OpenapiAccountsAddCondRespStr(addedAccounts string, ) *V0040OpenapiAccountsAddCondRespStr`

NewV0040OpenapiAccountsAddCondRespStr instantiates a new V0040OpenapiAccountsAddCondRespStr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiAccountsAddCondRespStrWithDefaults

`func NewV0040OpenapiAccountsAddCondRespStrWithDefaults() *V0040OpenapiAccountsAddCondRespStr`

NewV0040OpenapiAccountsAddCondRespStrWithDefaults instantiates a new V0040OpenapiAccountsAddCondRespStr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAccounts

`func (o *V0040OpenapiAccountsAddCondRespStr) GetAddedAccounts() string`

GetAddedAccounts returns the AddedAccounts field if non-nil, zero value otherwise.

### GetAddedAccountsOk

`func (o *V0040OpenapiAccountsAddCondRespStr) GetAddedAccountsOk() (*string, bool)`

GetAddedAccountsOk returns a tuple with the AddedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAccounts

`func (o *V0040OpenapiAccountsAddCondRespStr) SetAddedAccounts(v string)`

SetAddedAccounts sets AddedAccounts field to given value.


### GetMeta

`func (o *V0040OpenapiAccountsAddCondRespStr) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiAccountsAddCondRespStr) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiAccountsAddCondRespStr) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiAccountsAddCondRespStr) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiAccountsAddCondRespStr) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiAccountsAddCondRespStr) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiAccountsAddCondRespStr) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiAccountsAddCondRespStr) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiAccountsAddCondRespStr) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiAccountsAddCondRespStr) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiAccountsAddCondRespStr) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiAccountsAddCondRespStr) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


