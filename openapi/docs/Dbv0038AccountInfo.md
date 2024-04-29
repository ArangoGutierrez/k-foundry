# Dbv0038AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 
**Accounts** | Pointer to [**[]Dbv0038Account**](Dbv0038Account.md) | List of accounts | [optional] 

## Methods

### NewDbv0038AccountInfo

`func NewDbv0038AccountInfo() *Dbv0038AccountInfo`

NewDbv0038AccountInfo instantiates a new Dbv0038AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038AccountInfoWithDefaults

`func NewDbv0038AccountInfoWithDefaults() *Dbv0038AccountInfo`

NewDbv0038AccountInfoWithDefaults instantiates a new Dbv0038AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038AccountInfo) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038AccountInfo) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038AccountInfo) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038AccountInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038AccountInfo) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038AccountInfo) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038AccountInfo) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038AccountInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetAccounts

`func (o *Dbv0038AccountInfo) GetAccounts() []Dbv0038Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Dbv0038AccountInfo) GetAccountsOk() (*[]Dbv0038Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Dbv0038AccountInfo) SetAccounts(v []Dbv0038Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Dbv0038AccountInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


