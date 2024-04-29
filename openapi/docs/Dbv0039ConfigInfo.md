# Dbv0039ConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0039Meta**](Dbv0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0039Error**](Dbv0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]Dbv0039Warning**](Dbv0039Warning.md) | Slurm warnings | [optional] 
**Tres** | Pointer to [**[]V0039Tres**](V0039Tres.md) |  | [optional] 
**Accounts** | Pointer to [**[]V0039Account**](V0039Account.md) |  | [optional] 
**Associations** | Pointer to [**[]V0039Assoc**](V0039Assoc.md) |  | [optional] 
**Users** | Pointer to [**[]V0039User**](V0039User.md) |  | [optional] 
**Qos** | Pointer to [**[]V0039Qos**](V0039Qos.md) |  | [optional] 
**Wckeys** | Pointer to [**[]V0039Wckey**](V0039Wckey.md) |  | [optional] 
**Clusters** | Pointer to [**[]V0039ClusterRec**](V0039ClusterRec.md) |  | [optional] 

## Methods

### NewDbv0039ConfigInfo

`func NewDbv0039ConfigInfo() *Dbv0039ConfigInfo`

NewDbv0039ConfigInfo instantiates a new Dbv0039ConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0039ConfigInfoWithDefaults

`func NewDbv0039ConfigInfoWithDefaults() *Dbv0039ConfigInfo`

NewDbv0039ConfigInfoWithDefaults instantiates a new Dbv0039ConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0039ConfigInfo) GetMeta() Dbv0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0039ConfigInfo) GetMetaOk() (*Dbv0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0039ConfigInfo) SetMeta(v Dbv0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0039ConfigInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0039ConfigInfo) GetErrors() []Dbv0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0039ConfigInfo) GetErrorsOk() (*[]Dbv0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0039ConfigInfo) SetErrors(v []Dbv0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0039ConfigInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *Dbv0039ConfigInfo) GetWarnings() []Dbv0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Dbv0039ConfigInfo) GetWarningsOk() (*[]Dbv0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Dbv0039ConfigInfo) SetWarnings(v []Dbv0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Dbv0039ConfigInfo) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0039ConfigInfo) GetTres() []V0039Tres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0039ConfigInfo) GetTresOk() (*[]V0039Tres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0039ConfigInfo) SetTres(v []V0039Tres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0039ConfigInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetAccounts

`func (o *Dbv0039ConfigInfo) GetAccounts() []V0039Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Dbv0039ConfigInfo) GetAccountsOk() (*[]V0039Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Dbv0039ConfigInfo) SetAccounts(v []V0039Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Dbv0039ConfigInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0039ConfigInfo) GetAssociations() []V0039Assoc`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0039ConfigInfo) GetAssociationsOk() (*[]V0039Assoc, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0039ConfigInfo) SetAssociations(v []V0039Assoc)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0039ConfigInfo) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetUsers

`func (o *Dbv0039ConfigInfo) GetUsers() []V0039User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Dbv0039ConfigInfo) GetUsersOk() (*[]V0039User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Dbv0039ConfigInfo) SetUsers(v []V0039User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Dbv0039ConfigInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0039ConfigInfo) GetQos() []V0039Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0039ConfigInfo) GetQosOk() (*[]V0039Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0039ConfigInfo) SetQos(v []V0039Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0039ConfigInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *Dbv0039ConfigInfo) GetWckeys() []V0039Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *Dbv0039ConfigInfo) GetWckeysOk() (*[]V0039Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *Dbv0039ConfigInfo) SetWckeys(v []V0039Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *Dbv0039ConfigInfo) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.

### GetClusters

`func (o *Dbv0039ConfigInfo) GetClusters() []V0039ClusterRec`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *Dbv0039ConfigInfo) GetClustersOk() (*[]V0039ClusterRec, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *Dbv0039ConfigInfo) SetClusters(v []V0039ClusterRec)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *Dbv0039ConfigInfo) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


