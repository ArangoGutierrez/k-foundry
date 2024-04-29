# V0040StepTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** |  | [optional] 
**End** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Start** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**Suspended** | Pointer to **int32** |  | [optional] 
**System** | Pointer to [**V0039StepTimeSystem**](V0039StepTimeSystem.md) |  | [optional] 
**Total** | Pointer to [**V0039StepTimeSystem**](V0039StepTimeSystem.md) |  | [optional] 
**User** | Pointer to [**V0039StepTimeSystem**](V0039StepTimeSystem.md) |  | [optional] 

## Methods

### NewV0040StepTime

`func NewV0040StepTime() *V0040StepTime`

NewV0040StepTime instantiates a new V0040StepTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040StepTimeWithDefaults

`func NewV0040StepTimeWithDefaults() *V0040StepTime`

NewV0040StepTimeWithDefaults instantiates a new V0040StepTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *V0040StepTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *V0040StepTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *V0040StepTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *V0040StepTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEnd

`func (o *V0040StepTime) GetEnd() V0040Uint64NoVal`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V0040StepTime) GetEndOk() (*V0040Uint64NoVal, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V0040StepTime) SetEnd(v V0040Uint64NoVal)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V0040StepTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *V0040StepTime) GetStart() V0040Uint64NoVal`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0040StepTime) GetStartOk() (*V0040Uint64NoVal, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0040StepTime) SetStart(v V0040Uint64NoVal)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0040StepTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSuspended

`func (o *V0040StepTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *V0040StepTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *V0040StepTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *V0040StepTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *V0040StepTime) GetSystem() V0039StepTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0040StepTime) GetSystemOk() (*V0039StepTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0040StepTime) SetSystem(v V0039StepTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0040StepTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTotal

`func (o *V0040StepTime) GetTotal() V0039StepTimeSystem`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0040StepTime) GetTotalOk() (*V0039StepTimeSystem, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0040StepTime) SetTotal(v V0039StepTimeSystem)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0040StepTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *V0040StepTime) GetUser() V0039StepTimeSystem`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0040StepTime) GetUserOk() (*V0039StepTimeSystem, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0040StepTime) SetUser(v V0039StepTimeSystem)`

SetUser sets User field to given value.

### HasUser

`func (o *V0040StepTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


