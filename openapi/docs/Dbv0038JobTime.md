# Dbv0038JobTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | Total time elapsed | [optional] 
**Eligible** | Pointer to **int32** | Total time eligible to run | [optional] 
**End** | Pointer to **int32** | Timestamp of when job ended | [optional] 
**Start** | Pointer to **int32** | Timestamp of when job started | [optional] 
**Submission** | Pointer to **int32** | Timestamp of when job submitted | [optional] 
**Suspended** | Pointer to **int32** | Timestamp of when job last suspended | [optional] 
**System** | Pointer to [**Dbv0038JobTimeSystem**](Dbv0038JobTimeSystem.md) |  | [optional] 
**Total** | Pointer to [**Dbv0038JobTimeTotal**](Dbv0038JobTimeTotal.md) |  | [optional] 
**User** | Pointer to [**Dbv0038JobTimeUser**](Dbv0038JobTimeUser.md) |  | [optional] 
**Limit** | Pointer to **int32** | Job wall clock time limit | [optional] 

## Methods

### NewDbv0038JobTime

`func NewDbv0038JobTime() *Dbv0038JobTime`

NewDbv0038JobTime instantiates a new Dbv0038JobTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobTimeWithDefaults

`func NewDbv0038JobTimeWithDefaults() *Dbv0038JobTime`

NewDbv0038JobTimeWithDefaults instantiates a new Dbv0038JobTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *Dbv0038JobTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *Dbv0038JobTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *Dbv0038JobTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *Dbv0038JobTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEligible

`func (o *Dbv0038JobTime) GetEligible() int32`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *Dbv0038JobTime) GetEligibleOk() (*int32, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *Dbv0038JobTime) SetEligible(v int32)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *Dbv0038JobTime) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetEnd

`func (o *Dbv0038JobTime) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Dbv0038JobTime) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Dbv0038JobTime) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Dbv0038JobTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *Dbv0038JobTime) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Dbv0038JobTime) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Dbv0038JobTime) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *Dbv0038JobTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSubmission

`func (o *Dbv0038JobTime) GetSubmission() int32`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *Dbv0038JobTime) GetSubmissionOk() (*int32, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *Dbv0038JobTime) SetSubmission(v int32)`

SetSubmission sets Submission field to given value.

### HasSubmission

`func (o *Dbv0038JobTime) HasSubmission() bool`

HasSubmission returns a boolean if a field has been set.

### GetSuspended

`func (o *Dbv0038JobTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Dbv0038JobTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Dbv0038JobTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Dbv0038JobTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *Dbv0038JobTime) GetSystem() Dbv0038JobTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Dbv0038JobTime) GetSystemOk() (*Dbv0038JobTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Dbv0038JobTime) SetSystem(v Dbv0038JobTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Dbv0038JobTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTotal

`func (o *Dbv0038JobTime) GetTotal() Dbv0038JobTimeTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Dbv0038JobTime) GetTotalOk() (*Dbv0038JobTimeTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Dbv0038JobTime) SetTotal(v Dbv0038JobTimeTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Dbv0038JobTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *Dbv0038JobTime) GetUser() Dbv0038JobTimeUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Dbv0038JobTime) GetUserOk() (*Dbv0038JobTimeUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Dbv0038JobTime) SetUser(v Dbv0038JobTimeUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Dbv0038JobTime) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLimit

`func (o *Dbv0038JobTime) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Dbv0038JobTime) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Dbv0038JobTime) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Dbv0038JobTime) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


