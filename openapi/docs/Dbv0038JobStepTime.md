# Dbv0038JobStepTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | Total time elapsed | [optional] 
**End** | Pointer to **int32** | Timestamp of when job ended | [optional] 
**Start** | Pointer to **int32** | Timestamp of when job started | [optional] 
**Suspended** | Pointer to **int32** | Timestamp of when job last suspended | [optional] 
**System** | Pointer to [**Dbv0038JobTimeSystem**](Dbv0038JobTimeSystem.md) |  | [optional] 
**Total** | Pointer to [**Dbv0038JobTimeTotal**](Dbv0038JobTimeTotal.md) |  | [optional] 
**User** | Pointer to [**Dbv0038JobTimeUser**](Dbv0038JobTimeUser.md) |  | [optional] 

## Methods

### NewDbv0038JobStepTime

`func NewDbv0038JobStepTime() *Dbv0038JobStepTime`

NewDbv0038JobStepTime instantiates a new Dbv0038JobStepTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobStepTimeWithDefaults

`func NewDbv0038JobStepTimeWithDefaults() *Dbv0038JobStepTime`

NewDbv0038JobStepTimeWithDefaults instantiates a new Dbv0038JobStepTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *Dbv0038JobStepTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *Dbv0038JobStepTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *Dbv0038JobStepTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *Dbv0038JobStepTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEnd

`func (o *Dbv0038JobStepTime) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Dbv0038JobStepTime) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Dbv0038JobStepTime) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Dbv0038JobStepTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *Dbv0038JobStepTime) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Dbv0038JobStepTime) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Dbv0038JobStepTime) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *Dbv0038JobStepTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSuspended

`func (o *Dbv0038JobStepTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Dbv0038JobStepTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Dbv0038JobStepTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Dbv0038JobStepTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *Dbv0038JobStepTime) GetSystem() Dbv0038JobTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Dbv0038JobStepTime) GetSystemOk() (*Dbv0038JobTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Dbv0038JobStepTime) SetSystem(v Dbv0038JobTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Dbv0038JobStepTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTotal

`func (o *Dbv0038JobStepTime) GetTotal() Dbv0038JobTimeTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Dbv0038JobStepTime) GetTotalOk() (*Dbv0038JobTimeTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Dbv0038JobStepTime) SetTotal(v Dbv0038JobTimeTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Dbv0038JobStepTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *Dbv0038JobStepTime) GetUser() Dbv0038JobTimeUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Dbv0038JobStepTime) GetUserOk() (*Dbv0038JobTimeUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Dbv0038JobStepTime) SetUser(v Dbv0038JobTimeUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Dbv0038JobStepTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


