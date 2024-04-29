# V0039StatsRec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | Pointer to **int64** |  | [optional] 
**Rollups** | Pointer to [**[]V0039RollupStatsInner**](V0039RollupStatsInner.md) | list of recorded rollup statistics | [optional] 
**RPCs** | Pointer to [**[]V0039StatsRpc**](V0039StatsRpc.md) |  | [optional] 
**Users** | Pointer to [**[]V0039StatsUser**](V0039StatsUser.md) |  | [optional] 

## Methods

### NewV0039StatsRec

`func NewV0039StatsRec() *V0039StatsRec`

NewV0039StatsRec instantiates a new V0039StatsRec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039StatsRecWithDefaults

`func NewV0039StatsRecWithDefaults() *V0039StatsRec`

NewV0039StatsRecWithDefaults instantiates a new V0039StatsRec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *V0039StatsRec) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *V0039StatsRec) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *V0039StatsRec) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *V0039StatsRec) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetRollups

`func (o *V0039StatsRec) GetRollups() []V0039RollupStatsInner`

GetRollups returns the Rollups field if non-nil, zero value otherwise.

### GetRollupsOk

`func (o *V0039StatsRec) GetRollupsOk() (*[]V0039RollupStatsInner, bool)`

GetRollupsOk returns a tuple with the Rollups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollups

`func (o *V0039StatsRec) SetRollups(v []V0039RollupStatsInner)`

SetRollups sets Rollups field to given value.

### HasRollups

`func (o *V0039StatsRec) HasRollups() bool`

HasRollups returns a boolean if a field has been set.

### GetRPCs

`func (o *V0039StatsRec) GetRPCs() []V0039StatsRpc`

GetRPCs returns the RPCs field if non-nil, zero value otherwise.

### GetRPCsOk

`func (o *V0039StatsRec) GetRPCsOk() (*[]V0039StatsRpc, bool)`

GetRPCsOk returns a tuple with the RPCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRPCs

`func (o *V0039StatsRec) SetRPCs(v []V0039StatsRpc)`

SetRPCs sets RPCs field to given value.

### HasRPCs

`func (o *V0039StatsRec) HasRPCs() bool`

HasRPCs returns a boolean if a field has been set.

### GetUsers

`func (o *V0039StatsRec) GetUsers() []V0039StatsUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0039StatsRec) GetUsersOk() (*[]V0039StatsUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0039StatsRec) SetUsers(v []V0039StatsUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0039StatsRec) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


