# Dbv0038DiagStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | Pointer to **int32** | Unix timestamp of start time | [optional] 
**Rollups** | Pointer to [**[]Dbv0038DiagStatisticsRollupsInner**](Dbv0038DiagStatisticsRollupsInner.md) |  | [optional] 
**RPCs** | Pointer to [**[]Dbv0038DiagStatisticsRPCsInner**](Dbv0038DiagStatisticsRPCsInner.md) |  | [optional] 
**Users** | Pointer to [**[]Dbv0038DiagStatisticsUsersInner**](Dbv0038DiagStatisticsUsersInner.md) |  | [optional] 

## Methods

### NewDbv0038DiagStatistics

`func NewDbv0038DiagStatistics() *Dbv0038DiagStatistics`

NewDbv0038DiagStatistics instantiates a new Dbv0038DiagStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038DiagStatisticsWithDefaults

`func NewDbv0038DiagStatisticsWithDefaults() *Dbv0038DiagStatistics`

NewDbv0038DiagStatisticsWithDefaults instantiates a new Dbv0038DiagStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *Dbv0038DiagStatistics) GetTimeStart() int32`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *Dbv0038DiagStatistics) GetTimeStartOk() (*int32, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *Dbv0038DiagStatistics) SetTimeStart(v int32)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *Dbv0038DiagStatistics) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetRollups

`func (o *Dbv0038DiagStatistics) GetRollups() []Dbv0038DiagStatisticsRollupsInner`

GetRollups returns the Rollups field if non-nil, zero value otherwise.

### GetRollupsOk

`func (o *Dbv0038DiagStatistics) GetRollupsOk() (*[]Dbv0038DiagStatisticsRollupsInner, bool)`

GetRollupsOk returns a tuple with the Rollups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollups

`func (o *Dbv0038DiagStatistics) SetRollups(v []Dbv0038DiagStatisticsRollupsInner)`

SetRollups sets Rollups field to given value.

### HasRollups

`func (o *Dbv0038DiagStatistics) HasRollups() bool`

HasRollups returns a boolean if a field has been set.

### GetRPCs

`func (o *Dbv0038DiagStatistics) GetRPCs() []Dbv0038DiagStatisticsRPCsInner`

GetRPCs returns the RPCs field if non-nil, zero value otherwise.

### GetRPCsOk

`func (o *Dbv0038DiagStatistics) GetRPCsOk() (*[]Dbv0038DiagStatisticsRPCsInner, bool)`

GetRPCsOk returns a tuple with the RPCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRPCs

`func (o *Dbv0038DiagStatistics) SetRPCs(v []Dbv0038DiagStatisticsRPCsInner)`

SetRPCs sets RPCs field to given value.

### HasRPCs

`func (o *Dbv0038DiagStatistics) HasRPCs() bool`

HasRPCs returns a boolean if a field has been set.

### GetUsers

`func (o *Dbv0038DiagStatistics) GetUsers() []Dbv0038DiagStatisticsUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Dbv0038DiagStatistics) GetUsersOk() (*[]Dbv0038DiagStatisticsUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Dbv0038DiagStatistics) SetUsers(v []Dbv0038DiagStatisticsUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Dbv0038DiagStatistics) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


