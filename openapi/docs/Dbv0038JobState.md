# Dbv0038JobState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **string** | Current state of job | [optional] 
**Reason** | Pointer to **string** | Last reason job didn&#39;t run | [optional] 

## Methods

### NewDbv0038JobState

`func NewDbv0038JobState() *Dbv0038JobState`

NewDbv0038JobState instantiates a new Dbv0038JobState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobStateWithDefaults

`func NewDbv0038JobStateWithDefaults() *Dbv0038JobState`

NewDbv0038JobStateWithDefaults instantiates a new Dbv0038JobState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *Dbv0038JobState) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Dbv0038JobState) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Dbv0038JobState) SetCurrent(v string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *Dbv0038JobState) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetReason

`func (o *Dbv0038JobState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Dbv0038JobState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Dbv0038JobState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Dbv0038JobState) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


