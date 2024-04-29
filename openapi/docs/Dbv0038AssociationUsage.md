# Dbv0038AssociationUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrueJobCount** | Pointer to **int32** | Jobs accuring priority | [optional] 
**GroupUsedWallclock** | Pointer to **float32** | Group used wallclock time (s) | [optional] 
**FairshareFactor** | Pointer to **float32** | Fairshare factor | [optional] 
**FairshareShares** | Pointer to **int32** | Fairshare shares | [optional] 
**NormalizedPriority** | Pointer to **int32** | Currently active jobs | [optional] 
**NormalizedShares** | Pointer to **float32** | Normalized shares | [optional] 
**EffectiveNormalizedUsage** | Pointer to **float32** | Effective normalized usage | [optional] 
**RawUsage** | Pointer to **int32** | Raw usage | [optional] 
**JobCount** | Pointer to **int32** | Total jobs submitted | [optional] 
**FairshareLevel** | Pointer to **float32** | Fairshare level | [optional] 

## Methods

### NewDbv0038AssociationUsage

`func NewDbv0038AssociationUsage() *Dbv0038AssociationUsage`

NewDbv0038AssociationUsage instantiates a new Dbv0038AssociationUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038AssociationUsageWithDefaults

`func NewDbv0038AssociationUsageWithDefaults() *Dbv0038AssociationUsage`

NewDbv0038AssociationUsageWithDefaults instantiates a new Dbv0038AssociationUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrueJobCount

`func (o *Dbv0038AssociationUsage) GetAccrueJobCount() int32`

GetAccrueJobCount returns the AccrueJobCount field if non-nil, zero value otherwise.

### GetAccrueJobCountOk

`func (o *Dbv0038AssociationUsage) GetAccrueJobCountOk() (*int32, bool)`

GetAccrueJobCountOk returns a tuple with the AccrueJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrueJobCount

`func (o *Dbv0038AssociationUsage) SetAccrueJobCount(v int32)`

SetAccrueJobCount sets AccrueJobCount field to given value.

### HasAccrueJobCount

`func (o *Dbv0038AssociationUsage) HasAccrueJobCount() bool`

HasAccrueJobCount returns a boolean if a field has been set.

### GetGroupUsedWallclock

`func (o *Dbv0038AssociationUsage) GetGroupUsedWallclock() float32`

GetGroupUsedWallclock returns the GroupUsedWallclock field if non-nil, zero value otherwise.

### GetGroupUsedWallclockOk

`func (o *Dbv0038AssociationUsage) GetGroupUsedWallclockOk() (*float32, bool)`

GetGroupUsedWallclockOk returns a tuple with the GroupUsedWallclock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUsedWallclock

`func (o *Dbv0038AssociationUsage) SetGroupUsedWallclock(v float32)`

SetGroupUsedWallclock sets GroupUsedWallclock field to given value.

### HasGroupUsedWallclock

`func (o *Dbv0038AssociationUsage) HasGroupUsedWallclock() bool`

HasGroupUsedWallclock returns a boolean if a field has been set.

### GetFairshareFactor

`func (o *Dbv0038AssociationUsage) GetFairshareFactor() float32`

GetFairshareFactor returns the FairshareFactor field if non-nil, zero value otherwise.

### GetFairshareFactorOk

`func (o *Dbv0038AssociationUsage) GetFairshareFactorOk() (*float32, bool)`

GetFairshareFactorOk returns a tuple with the FairshareFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshareFactor

`func (o *Dbv0038AssociationUsage) SetFairshareFactor(v float32)`

SetFairshareFactor sets FairshareFactor field to given value.

### HasFairshareFactor

`func (o *Dbv0038AssociationUsage) HasFairshareFactor() bool`

HasFairshareFactor returns a boolean if a field has been set.

### GetFairshareShares

`func (o *Dbv0038AssociationUsage) GetFairshareShares() int32`

GetFairshareShares returns the FairshareShares field if non-nil, zero value otherwise.

### GetFairshareSharesOk

`func (o *Dbv0038AssociationUsage) GetFairshareSharesOk() (*int32, bool)`

GetFairshareSharesOk returns a tuple with the FairshareShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshareShares

`func (o *Dbv0038AssociationUsage) SetFairshareShares(v int32)`

SetFairshareShares sets FairshareShares field to given value.

### HasFairshareShares

`func (o *Dbv0038AssociationUsage) HasFairshareShares() bool`

HasFairshareShares returns a boolean if a field has been set.

### GetNormalizedPriority

`func (o *Dbv0038AssociationUsage) GetNormalizedPriority() int32`

GetNormalizedPriority returns the NormalizedPriority field if non-nil, zero value otherwise.

### GetNormalizedPriorityOk

`func (o *Dbv0038AssociationUsage) GetNormalizedPriorityOk() (*int32, bool)`

GetNormalizedPriorityOk returns a tuple with the NormalizedPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedPriority

`func (o *Dbv0038AssociationUsage) SetNormalizedPriority(v int32)`

SetNormalizedPriority sets NormalizedPriority field to given value.

### HasNormalizedPriority

`func (o *Dbv0038AssociationUsage) HasNormalizedPriority() bool`

HasNormalizedPriority returns a boolean if a field has been set.

### GetNormalizedShares

`func (o *Dbv0038AssociationUsage) GetNormalizedShares() float32`

GetNormalizedShares returns the NormalizedShares field if non-nil, zero value otherwise.

### GetNormalizedSharesOk

`func (o *Dbv0038AssociationUsage) GetNormalizedSharesOk() (*float32, bool)`

GetNormalizedSharesOk returns a tuple with the NormalizedShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedShares

`func (o *Dbv0038AssociationUsage) SetNormalizedShares(v float32)`

SetNormalizedShares sets NormalizedShares field to given value.

### HasNormalizedShares

`func (o *Dbv0038AssociationUsage) HasNormalizedShares() bool`

HasNormalizedShares returns a boolean if a field has been set.

### GetEffectiveNormalizedUsage

`func (o *Dbv0038AssociationUsage) GetEffectiveNormalizedUsage() float32`

GetEffectiveNormalizedUsage returns the EffectiveNormalizedUsage field if non-nil, zero value otherwise.

### GetEffectiveNormalizedUsageOk

`func (o *Dbv0038AssociationUsage) GetEffectiveNormalizedUsageOk() (*float32, bool)`

GetEffectiveNormalizedUsageOk returns a tuple with the EffectiveNormalizedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveNormalizedUsage

`func (o *Dbv0038AssociationUsage) SetEffectiveNormalizedUsage(v float32)`

SetEffectiveNormalizedUsage sets EffectiveNormalizedUsage field to given value.

### HasEffectiveNormalizedUsage

`func (o *Dbv0038AssociationUsage) HasEffectiveNormalizedUsage() bool`

HasEffectiveNormalizedUsage returns a boolean if a field has been set.

### GetRawUsage

`func (o *Dbv0038AssociationUsage) GetRawUsage() int32`

GetRawUsage returns the RawUsage field if non-nil, zero value otherwise.

### GetRawUsageOk

`func (o *Dbv0038AssociationUsage) GetRawUsageOk() (*int32, bool)`

GetRawUsageOk returns a tuple with the RawUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawUsage

`func (o *Dbv0038AssociationUsage) SetRawUsage(v int32)`

SetRawUsage sets RawUsage field to given value.

### HasRawUsage

`func (o *Dbv0038AssociationUsage) HasRawUsage() bool`

HasRawUsage returns a boolean if a field has been set.

### GetJobCount

`func (o *Dbv0038AssociationUsage) GetJobCount() int32`

GetJobCount returns the JobCount field if non-nil, zero value otherwise.

### GetJobCountOk

`func (o *Dbv0038AssociationUsage) GetJobCountOk() (*int32, bool)`

GetJobCountOk returns a tuple with the JobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCount

`func (o *Dbv0038AssociationUsage) SetJobCount(v int32)`

SetJobCount sets JobCount field to given value.

### HasJobCount

`func (o *Dbv0038AssociationUsage) HasJobCount() bool`

HasJobCount returns a boolean if a field has been set.

### GetFairshareLevel

`func (o *Dbv0038AssociationUsage) GetFairshareLevel() float32`

GetFairshareLevel returns the FairshareLevel field if non-nil, zero value otherwise.

### GetFairshareLevelOk

`func (o *Dbv0038AssociationUsage) GetFairshareLevelOk() (*float32, bool)`

GetFairshareLevelOk returns a tuple with the FairshareLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshareLevel

`func (o *Dbv0038AssociationUsage) SetFairshareLevel(v float32)`

SetFairshareLevel sets FairshareLevel field to given value.

### HasFairshareLevel

`func (o *Dbv0038AssociationUsage) HasFairshareLevel() bool`

HasFairshareLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


