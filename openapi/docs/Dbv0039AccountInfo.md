# Dbv0039AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0039Meta**](Dbv0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0039Error**](Dbv0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]Dbv0039Warning**](Dbv0039Warning.md) | Slurm warnings | [optional] 
**Accounts** | Pointer to [**[]V0039Account**](V0039Account.md) |  | [optional] 

## Methods

### NewDbv0039AccountInfo

`func NewDbv0039AccountInfo() *Dbv0039AccountInfo`

NewDbv0039AccountInfo instantiates a new Dbv0039AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0039AccountInfoWithDefaults

`func NewDbv0039AccountInfoWithDefaults() *Dbv0039AccountInfo`

NewDbv0039AccountInfoWithDefaults instantiates a new Dbv0039AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0039AccountInfo) GetMeta() Dbv0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0039AccountInfo) GetMetaOk() (*Dbv0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0039AccountInfo) SetMeta(v Dbv0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0039AccountInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0039AccountInfo) GetErrors() []Dbv0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0039AccountInfo) GetErrorsOk() (*[]Dbv0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0039AccountInfo) SetErrors(v []Dbv0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0039AccountInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *Dbv0039AccountInfo) GetWarnings() []Dbv0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Dbv0039AccountInfo) GetWarningsOk() (*[]Dbv0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Dbv0039AccountInfo) SetWarnings(v []Dbv0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Dbv0039AccountInfo) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetAccounts

`func (o *Dbv0039AccountInfo) GetAccounts() []V0039Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Dbv0039AccountInfo) GetAccountsOk() (*[]V0039Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Dbv0039AccountInfo) SetAccounts(v []V0039Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Dbv0039AccountInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


