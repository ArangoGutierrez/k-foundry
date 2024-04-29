# V0040StatsRec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | Pointer to **int64** |  | [optional] 
**Rollups** | Pointer to [**[]V0039RollupStatsInner**](V0039RollupStatsInner.md) | list of recorded rollup statistics | [optional] 
**RPCs** | Pointer to [**[]V0040StatsRpc**](V0040StatsRpc.md) |  | [optional] 
**Users** | Pointer to [**[]V0040StatsUser**](V0040StatsUser.md) |  | [optional] 

## Methods

### NewV0040StatsRec

`func NewV0040StatsRec() *V0040StatsRec`

NewV0040StatsRec instantiates a new V0040StatsRec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040StatsRecWithDefaults

`func NewV0040StatsRecWithDefaults() *V0040StatsRec`

NewV0040StatsRecWithDefaults instantiates a new V0040StatsRec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *V0040StatsRec) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *V0040StatsRec) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *V0040StatsRec) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *V0040StatsRec) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetRollups

`func (o *V0040StatsRec) GetRollups() []V0039RollupStatsInner`

GetRollups returns the Rollups field if non-nil, zero value otherwise.

### GetRollupsOk

`func (o *V0040StatsRec) GetRollupsOk() (*[]V0039RollupStatsInner, bool)`

GetRollupsOk returns a tuple with the Rollups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollups

`func (o *V0040StatsRec) SetRollups(v []V0039RollupStatsInner)`

SetRollups sets Rollups field to given value.

### HasRollups

`func (o *V0040StatsRec) HasRollups() bool`

HasRollups returns a boolean if a field has been set.

### GetRPCs

`func (o *V0040StatsRec) GetRPCs() []V0040StatsRpc`

GetRPCs returns the RPCs field if non-nil, zero value otherwise.

### GetRPCsOk

`func (o *V0040StatsRec) GetRPCsOk() (*[]V0040StatsRpc, bool)`

GetRPCsOk returns a tuple with the RPCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRPCs

`func (o *V0040StatsRec) SetRPCs(v []V0040StatsRpc)`

SetRPCs sets RPCs field to given value.

### HasRPCs

`func (o *V0040StatsRec) HasRPCs() bool`

HasRPCs returns a boolean if a field has been set.

### GetUsers

`func (o *V0040StatsRec) GetUsers() []V0040StatsUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0040StatsRec) GetUsersOk() (*[]V0040StatsUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0040StatsRec) SetUsers(v []V0040StatsUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0040StatsRec) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


