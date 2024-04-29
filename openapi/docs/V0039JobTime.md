# V0039JobTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** |  | [optional] 
**Eligible** | Pointer to **int64** |  | [optional] 
**End** | Pointer to **int64** |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**Submission** | Pointer to **int64** |  | [optional] 
**Suspended** | Pointer to **int32** |  | [optional] 
**System** | Pointer to [**V0039JobTimeSystem**](V0039JobTimeSystem.md) |  | [optional] 
**Limit** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 
**Total** | Pointer to [**V0039JobTimeSystem**](V0039JobTimeSystem.md) |  | [optional] 
**User** | Pointer to [**V0039JobTimeSystem**](V0039JobTimeSystem.md) |  | [optional] 

## Methods

### NewV0039JobTime

`func NewV0039JobTime() *V0039JobTime`

NewV0039JobTime instantiates a new V0039JobTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039JobTimeWithDefaults

`func NewV0039JobTimeWithDefaults() *V0039JobTime`

NewV0039JobTimeWithDefaults instantiates a new V0039JobTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *V0039JobTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *V0039JobTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *V0039JobTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *V0039JobTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEligible

`func (o *V0039JobTime) GetEligible() int64`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *V0039JobTime) GetEligibleOk() (*int64, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *V0039JobTime) SetEligible(v int64)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *V0039JobTime) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetEnd

`func (o *V0039JobTime) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V0039JobTime) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V0039JobTime) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V0039JobTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *V0039JobTime) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0039JobTime) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0039JobTime) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0039JobTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSubmission

`func (o *V0039JobTime) GetSubmission() int64`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *V0039JobTime) GetSubmissionOk() (*int64, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *V0039JobTime) SetSubmission(v int64)`

SetSubmission sets Submission field to given value.

### HasSubmission

`func (o *V0039JobTime) HasSubmission() bool`

HasSubmission returns a boolean if a field has been set.

### GetSuspended

`func (o *V0039JobTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *V0039JobTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *V0039JobTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *V0039JobTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *V0039JobTime) GetSystem() V0039JobTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0039JobTime) GetSystemOk() (*V0039JobTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0039JobTime) SetSystem(v V0039JobTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0039JobTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetLimit

`func (o *V0039JobTime) GetLimit() V0039Uint32NoVal`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V0039JobTime) GetLimitOk() (*V0039Uint32NoVal, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V0039JobTime) SetLimit(v V0039Uint32NoVal)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V0039JobTime) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *V0039JobTime) GetTotal() V0039JobTimeSystem`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0039JobTime) GetTotalOk() (*V0039JobTimeSystem, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0039JobTime) SetTotal(v V0039JobTimeSystem)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0039JobTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *V0039JobTime) GetUser() V0039JobTimeSystem`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0039JobTime) GetUserOk() (*V0039JobTimeSystem, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0039JobTime) SetUser(v V0039JobTimeSystem)`

SetUser sets User field to given value.

### HasUser

`func (o *V0039JobTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


