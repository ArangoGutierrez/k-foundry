# V0040OpenapiSlurmdbdConfigResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]V0040ClusterRec**](V0040ClusterRec.md) |  | [optional] 
**Tres** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Accounts** | Pointer to [**[]V0040Account**](V0040Account.md) |  | [optional] 
**Users** | Pointer to [**[]V0040User**](V0040User.md) |  | [optional] 
**Qos** | Pointer to [**[]V0040Qos**](V0040Qos.md) |  | [optional] 
**Wckeys** | Pointer to [**[]V0040Wckey**](V0040Wckey.md) |  | [optional] 
**Associations** | Pointer to [**[]V0040Assoc**](V0040Assoc.md) |  | [optional] 
**Instances** | Pointer to [**[]V0040Instance**](V0040Instance.md) |  | [optional] 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiSlurmdbdConfigResp

`func NewV0040OpenapiSlurmdbdConfigResp() *V0040OpenapiSlurmdbdConfigResp`

NewV0040OpenapiSlurmdbdConfigResp instantiates a new V0040OpenapiSlurmdbdConfigResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiSlurmdbdConfigRespWithDefaults

`func NewV0040OpenapiSlurmdbdConfigRespWithDefaults() *V0040OpenapiSlurmdbdConfigResp`

NewV0040OpenapiSlurmdbdConfigRespWithDefaults instantiates a new V0040OpenapiSlurmdbdConfigResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *V0040OpenapiSlurmdbdConfigResp) GetClusters() []V0040ClusterRec`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetClustersOk() (*[]V0040ClusterRec, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0040OpenapiSlurmdbdConfigResp) SetClusters(v []V0040ClusterRec)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0040OpenapiSlurmdbdConfigResp) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTres

`func (o *V0040OpenapiSlurmdbdConfigResp) GetTres() []V0040Tres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetTresOk() (*[]V0040Tres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0040OpenapiSlurmdbdConfigResp) SetTres(v []V0040Tres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0040OpenapiSlurmdbdConfigResp) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetAccounts

`func (o *V0040OpenapiSlurmdbdConfigResp) GetAccounts() []V0040Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetAccountsOk() (*[]V0040Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0040OpenapiSlurmdbdConfigResp) SetAccounts(v []V0040Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0040OpenapiSlurmdbdConfigResp) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetUsers

`func (o *V0040OpenapiSlurmdbdConfigResp) GetUsers() []V0040User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetUsersOk() (*[]V0040User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0040OpenapiSlurmdbdConfigResp) SetUsers(v []V0040User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0040OpenapiSlurmdbdConfigResp) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *V0040OpenapiSlurmdbdConfigResp) GetQos() []V0040Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetQosOk() (*[]V0040Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0040OpenapiSlurmdbdConfigResp) SetQos(v []V0040Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0040OpenapiSlurmdbdConfigResp) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *V0040OpenapiSlurmdbdConfigResp) GetWckeys() []V0040Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetWckeysOk() (*[]V0040Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0040OpenapiSlurmdbdConfigResp) SetWckeys(v []V0040Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0040OpenapiSlurmdbdConfigResp) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.

### GetAssociations

`func (o *V0040OpenapiSlurmdbdConfigResp) GetAssociations() []V0040Assoc`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetAssociationsOk() (*[]V0040Assoc, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0040OpenapiSlurmdbdConfigResp) SetAssociations(v []V0040Assoc)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0040OpenapiSlurmdbdConfigResp) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetInstances

`func (o *V0040OpenapiSlurmdbdConfigResp) GetInstances() []V0040Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetInstancesOk() (*[]V0040Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *V0040OpenapiSlurmdbdConfigResp) SetInstances(v []V0040Instance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *V0040OpenapiSlurmdbdConfigResp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMeta

`func (o *V0040OpenapiSlurmdbdConfigResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiSlurmdbdConfigResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiSlurmdbdConfigResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiSlurmdbdConfigResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiSlurmdbdConfigResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiSlurmdbdConfigResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiSlurmdbdConfigResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiSlurmdbdConfigResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiSlurmdbdConfigResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiSlurmdbdConfigResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


