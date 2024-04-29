# Dbv0038ConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 
**Tres** | Pointer to [**[][]Dbv0038TresListInner**]([]Dbv0038TresListInner.md) | Array of TRES | [optional] 
**Accounts** | Pointer to [**[]Dbv0038Account**](Dbv0038Account.md) | Array of accounts | [optional] 
**Associations** | Pointer to [**[]Dbv0038Association**](Dbv0038Association.md) | Array of associations | [optional] 
**Users** | Pointer to [**[]Dbv0038User**](Dbv0038User.md) | Array of users | [optional] 
**Qos** | Pointer to [**[]Dbv0038Qos**](Dbv0038Qos.md) | Array of qos | [optional] 
**Wckeys** | Pointer to [**[]Dbv0038Wckey**](Dbv0038Wckey.md) | Array of wckeys | [optional] 

## Methods

### NewDbv0038ConfigInfo

`func NewDbv0038ConfigInfo() *Dbv0038ConfigInfo`

NewDbv0038ConfigInfo instantiates a new Dbv0038ConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038ConfigInfoWithDefaults

`func NewDbv0038ConfigInfoWithDefaults() *Dbv0038ConfigInfo`

NewDbv0038ConfigInfoWithDefaults instantiates a new Dbv0038ConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038ConfigInfo) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038ConfigInfo) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038ConfigInfo) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038ConfigInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038ConfigInfo) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038ConfigInfo) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038ConfigInfo) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038ConfigInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0038ConfigInfo) GetTres() [][]Dbv0038TresListInner`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0038ConfigInfo) GetTresOk() (*[][]Dbv0038TresListInner, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0038ConfigInfo) SetTres(v [][]Dbv0038TresListInner)`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0038ConfigInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetAccounts

`func (o *Dbv0038ConfigInfo) GetAccounts() []Dbv0038Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Dbv0038ConfigInfo) GetAccountsOk() (*[]Dbv0038Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Dbv0038ConfigInfo) SetAccounts(v []Dbv0038Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Dbv0038ConfigInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0038ConfigInfo) GetAssociations() []Dbv0038Association`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0038ConfigInfo) GetAssociationsOk() (*[]Dbv0038Association, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0038ConfigInfo) SetAssociations(v []Dbv0038Association)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0038ConfigInfo) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetUsers

`func (o *Dbv0038ConfigInfo) GetUsers() []Dbv0038User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Dbv0038ConfigInfo) GetUsersOk() (*[]Dbv0038User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Dbv0038ConfigInfo) SetUsers(v []Dbv0038User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Dbv0038ConfigInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0038ConfigInfo) GetQos() []Dbv0038Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0038ConfigInfo) GetQosOk() (*[]Dbv0038Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0038ConfigInfo) SetQos(v []Dbv0038Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0038ConfigInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *Dbv0038ConfigInfo) GetWckeys() []Dbv0038Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *Dbv0038ConfigInfo) GetWckeysOk() (*[]Dbv0038Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *Dbv0038ConfigInfo) SetWckeys(v []Dbv0038Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *Dbv0038ConfigInfo) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


