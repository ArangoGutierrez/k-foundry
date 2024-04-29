# V0039Assoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**Default** | Pointer to [**V0039AssocDefault**](V0039AssocDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Max** | Pointer to [**V0039AssocMax**](V0039AssocMax.md) |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Min** | Pointer to [**V0039AssocMin**](V0039AssocMin.md) |  | [optional] 
**ParentAccount** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 
**Qos** | Pointer to **[]string** | List of QOS names | [optional] 
**SharesRaw** | Pointer to **int32** |  | [optional] 
**Usage** | Pointer to [**V0039AssocUsage**](V0039AssocUsage.md) |  | [optional] 
**User** | **string** |  | 

## Methods

### NewV0039Assoc

`func NewV0039Assoc(user string, ) *V0039Assoc`

NewV0039Assoc instantiates a new V0039Assoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039AssocWithDefaults

`func NewV0039AssocWithDefaults() *V0039Assoc`

NewV0039AssocWithDefaults instantiates a new V0039Assoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0039Assoc) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0039Assoc) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0039Assoc) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0039Assoc) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCluster

`func (o *V0039Assoc) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0039Assoc) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0039Assoc) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0039Assoc) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDefault

`func (o *V0039Assoc) GetDefault() V0039AssocDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0039Assoc) GetDefaultOk() (*V0039AssocDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0039Assoc) SetDefault(v V0039AssocDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0039Assoc) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0039Assoc) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0039Assoc) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0039Assoc) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0039Assoc) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMax

`func (o *V0039Assoc) GetMax() V0039AssocMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0039Assoc) GetMaxOk() (*V0039AssocMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0039Assoc) SetMax(v V0039AssocMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0039Assoc) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetIsDefault

`func (o *V0039Assoc) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *V0039Assoc) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *V0039Assoc) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *V0039Assoc) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMin

`func (o *V0039Assoc) GetMin() V0039AssocMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0039Assoc) GetMinOk() (*V0039AssocMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0039Assoc) SetMin(v V0039AssocMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0039Assoc) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetParentAccount

`func (o *V0039Assoc) GetParentAccount() string`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *V0039Assoc) GetParentAccountOk() (*string, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *V0039Assoc) SetParentAccount(v string)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *V0039Assoc) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### GetPartition

`func (o *V0039Assoc) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0039Assoc) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0039Assoc) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0039Assoc) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *V0039Assoc) GetPriority() V0039Uint32NoVal`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0039Assoc) GetPriorityOk() (*V0039Uint32NoVal, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0039Assoc) SetPriority(v V0039Uint32NoVal)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0039Assoc) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *V0039Assoc) GetQos() []string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0039Assoc) GetQosOk() (*[]string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0039Assoc) SetQos(v []string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0039Assoc) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetSharesRaw

`func (o *V0039Assoc) GetSharesRaw() int32`

GetSharesRaw returns the SharesRaw field if non-nil, zero value otherwise.

### GetSharesRawOk

`func (o *V0039Assoc) GetSharesRawOk() (*int32, bool)`

GetSharesRawOk returns a tuple with the SharesRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesRaw

`func (o *V0039Assoc) SetSharesRaw(v int32)`

SetSharesRaw sets SharesRaw field to given value.

### HasSharesRaw

`func (o *V0039Assoc) HasSharesRaw() bool`

HasSharesRaw returns a boolean if a field has been set.

### GetUsage

`func (o *V0039Assoc) GetUsage() V0039AssocUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *V0039Assoc) GetUsageOk() (*V0039AssocUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *V0039Assoc) SetUsage(v V0039AssocUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *V0039Assoc) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *V0039Assoc) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0039Assoc) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0039Assoc) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


