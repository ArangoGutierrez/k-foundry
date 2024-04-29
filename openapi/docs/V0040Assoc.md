# V0040Assoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounting** | Pointer to [**[]V0040Accounting**](V0040Accounting.md) |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** | comment for the association | [optional] 
**Default** | Pointer to [**V0039AssocDefault**](V0039AssocDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Max** | Pointer to [**V0040AssocMax**](V0040AssocMax.md) |  | [optional] 
**Id** | Pointer to [**V0040AssocShort**](V0040AssocShort.md) |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Lineage** | Pointer to **string** | Complete path up the hierarchy to the root association | [optional] 
**Min** | Pointer to [**V0040AssocMin**](V0040AssocMin.md) |  | [optional] 
**ParentAccount** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Qos** | Pointer to **[]string** | List of QOS names | [optional] 
**SharesRaw** | Pointer to **int32** |  | [optional] 
**User** | **string** |  | 

## Methods

### NewV0040Assoc

`func NewV0040Assoc(user string, ) *V0040Assoc`

NewV0040Assoc instantiates a new V0040Assoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040AssocWithDefaults

`func NewV0040AssocWithDefaults() *V0040Assoc`

NewV0040AssocWithDefaults instantiates a new V0040Assoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounting

`func (o *V0040Assoc) GetAccounting() []V0040Accounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *V0040Assoc) GetAccountingOk() (*[]V0040Accounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *V0040Assoc) SetAccounting(v []V0040Accounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *V0040Assoc) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetAccount

`func (o *V0040Assoc) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0040Assoc) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0040Assoc) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0040Assoc) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCluster

`func (o *V0040Assoc) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0040Assoc) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0040Assoc) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0040Assoc) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetComment

`func (o *V0040Assoc) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0040Assoc) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0040Assoc) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0040Assoc) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefault

`func (o *V0040Assoc) GetDefault() V0039AssocDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0040Assoc) GetDefaultOk() (*V0039AssocDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0040Assoc) SetDefault(v V0039AssocDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0040Assoc) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0040Assoc) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0040Assoc) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0040Assoc) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0040Assoc) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMax

`func (o *V0040Assoc) GetMax() V0040AssocMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0040Assoc) GetMaxOk() (*V0040AssocMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0040Assoc) SetMax(v V0040AssocMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0040Assoc) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetId

`func (o *V0040Assoc) GetId() V0040AssocShort`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0040Assoc) GetIdOk() (*V0040AssocShort, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0040Assoc) SetId(v V0040AssocShort)`

SetId sets Id field to given value.

### HasId

`func (o *V0040Assoc) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *V0040Assoc) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *V0040Assoc) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *V0040Assoc) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *V0040Assoc) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLineage

`func (o *V0040Assoc) GetLineage() string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V0040Assoc) GetLineageOk() (*string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V0040Assoc) SetLineage(v string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V0040Assoc) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetMin

`func (o *V0040Assoc) GetMin() V0040AssocMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0040Assoc) GetMinOk() (*V0040AssocMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0040Assoc) SetMin(v V0040AssocMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0040Assoc) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetParentAccount

`func (o *V0040Assoc) GetParentAccount() string`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *V0040Assoc) GetParentAccountOk() (*string, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *V0040Assoc) SetParentAccount(v string)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *V0040Assoc) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### GetPartition

`func (o *V0040Assoc) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0040Assoc) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0040Assoc) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0040Assoc) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *V0040Assoc) GetPriority() V0040Uint32NoVal`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0040Assoc) GetPriorityOk() (*V0040Uint32NoVal, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0040Assoc) SetPriority(v V0040Uint32NoVal)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0040Assoc) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *V0040Assoc) GetQos() []string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0040Assoc) GetQosOk() (*[]string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0040Assoc) SetQos(v []string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0040Assoc) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetSharesRaw

`func (o *V0040Assoc) GetSharesRaw() int32`

GetSharesRaw returns the SharesRaw field if non-nil, zero value otherwise.

### GetSharesRawOk

`func (o *V0040Assoc) GetSharesRawOk() (*int32, bool)`

GetSharesRawOk returns a tuple with the SharesRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesRaw

`func (o *V0040Assoc) SetSharesRaw(v int32)`

SetSharesRaw sets SharesRaw field to given value.

### HasSharesRaw

`func (o *V0040Assoc) HasSharesRaw() bool`

HasSharesRaw returns a boolean if a field has been set.

### GetUser

`func (o *V0040Assoc) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0040Assoc) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0040Assoc) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


