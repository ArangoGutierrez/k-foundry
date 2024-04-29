# Dbv0038SetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]Dbv0038ClustersProperties**](Dbv0038ClustersProperties.md) |  | [optional] 
**TRES** | Pointer to [**[][]Dbv0038TresListInner**]([]Dbv0038TresListInner.md) |  | [optional] 
**Accounts** | Pointer to [**[]Dbv0038UpdateAccount**](Dbv0038UpdateAccount.md) |  | [optional] 
**Users** | Pointer to [**[]Dbv0038User**](Dbv0038User.md) |  | [optional] 
**Qos** | Pointer to [**[]Dbv0038Qos**](Dbv0038Qos.md) |  | [optional] 
**Wckeys** | Pointer to [**[]Dbv0038Wckey**](Dbv0038Wckey.md) |  | [optional] 
**Associations** | Pointer to [**[]Dbv0038Association**](Dbv0038Association.md) |  | [optional] 

## Methods

### NewDbv0038SetConfig

`func NewDbv0038SetConfig() *Dbv0038SetConfig`

NewDbv0038SetConfig instantiates a new Dbv0038SetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038SetConfigWithDefaults

`func NewDbv0038SetConfigWithDefaults() *Dbv0038SetConfig`

NewDbv0038SetConfigWithDefaults instantiates a new Dbv0038SetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *Dbv0038SetConfig) GetClusters() []Dbv0038ClustersProperties`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *Dbv0038SetConfig) GetClustersOk() (*[]Dbv0038ClustersProperties, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *Dbv0038SetConfig) SetClusters(v []Dbv0038ClustersProperties)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *Dbv0038SetConfig) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTRES

`func (o *Dbv0038SetConfig) GetTRES() [][]Dbv0038TresListInner`

GetTRES returns the TRES field if non-nil, zero value otherwise.

### GetTRESOk

`func (o *Dbv0038SetConfig) GetTRESOk() (*[][]Dbv0038TresListInner, bool)`

GetTRESOk returns a tuple with the TRES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTRES

`func (o *Dbv0038SetConfig) SetTRES(v [][]Dbv0038TresListInner)`

SetTRES sets TRES field to given value.

### HasTRES

`func (o *Dbv0038SetConfig) HasTRES() bool`

HasTRES returns a boolean if a field has been set.

### GetAccounts

`func (o *Dbv0038SetConfig) GetAccounts() []Dbv0038UpdateAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Dbv0038SetConfig) GetAccountsOk() (*[]Dbv0038UpdateAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Dbv0038SetConfig) SetAccounts(v []Dbv0038UpdateAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Dbv0038SetConfig) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetUsers

`func (o *Dbv0038SetConfig) GetUsers() []Dbv0038User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Dbv0038SetConfig) GetUsersOk() (*[]Dbv0038User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Dbv0038SetConfig) SetUsers(v []Dbv0038User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Dbv0038SetConfig) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0038SetConfig) GetQos() []Dbv0038Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0038SetConfig) GetQosOk() (*[]Dbv0038Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0038SetConfig) SetQos(v []Dbv0038Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0038SetConfig) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *Dbv0038SetConfig) GetWckeys() []Dbv0038Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *Dbv0038SetConfig) GetWckeysOk() (*[]Dbv0038Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *Dbv0038SetConfig) SetWckeys(v []Dbv0038Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *Dbv0038SetConfig) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0038SetConfig) GetAssociations() []Dbv0038Association`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0038SetConfig) GetAssociationsOk() (*[]Dbv0038Association, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0038SetConfig) SetAssociations(v []Dbv0038Association)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0038SetConfig) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


