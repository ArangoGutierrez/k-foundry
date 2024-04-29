# V0040AccountsAddCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **[]string** |  | 
**Association** | Pointer to [**V0040AssocRecSet**](V0040AssocRecSet.md) |  | [optional] 
**Clusters** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0040AccountsAddCond

`func NewV0040AccountsAddCond(accounts []string, ) *V0040AccountsAddCond`

NewV0040AccountsAddCond instantiates a new V0040AccountsAddCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040AccountsAddCondWithDefaults

`func NewV0040AccountsAddCondWithDefaults() *V0040AccountsAddCond`

NewV0040AccountsAddCondWithDefaults instantiates a new V0040AccountsAddCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0040AccountsAddCond) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0040AccountsAddCond) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0040AccountsAddCond) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.


### GetAssociation

`func (o *V0040AccountsAddCond) GetAssociation() V0040AssocRecSet`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *V0040AccountsAddCond) GetAssociationOk() (*V0040AssocRecSet, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *V0040AccountsAddCond) SetAssociation(v V0040AssocRecSet)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *V0040AccountsAddCond) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetClusters

`func (o *V0040AccountsAddCond) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0040AccountsAddCond) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0040AccountsAddCond) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0040AccountsAddCond) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


