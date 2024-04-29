# Dbv0038JobTimeSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seconds** | Pointer to **int32** | Total number of CPU-seconds used by the system on behalf of the process (in kernel mode), in seconds | [optional] 
**Microseconds** | Pointer to **int32** | Total number of CPU-seconds used by the system on behalf of the process (in kernel mode), in microseconds | [optional] 

## Methods

### NewDbv0038JobTimeSystem

`func NewDbv0038JobTimeSystem() *Dbv0038JobTimeSystem`

NewDbv0038JobTimeSystem instantiates a new Dbv0038JobTimeSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobTimeSystemWithDefaults

`func NewDbv0038JobTimeSystemWithDefaults() *Dbv0038JobTimeSystem`

NewDbv0038JobTimeSystemWithDefaults instantiates a new Dbv0038JobTimeSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeconds

`func (o *Dbv0038JobTimeSystem) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *Dbv0038JobTimeSystem) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *Dbv0038JobTimeSystem) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *Dbv0038JobTimeSystem) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetMicroseconds

`func (o *Dbv0038JobTimeSystem) GetMicroseconds() int32`

GetMicroseconds returns the Microseconds field if non-nil, zero value otherwise.

### GetMicrosecondsOk

`func (o *Dbv0038JobTimeSystem) GetMicrosecondsOk() (*int32, bool)`

GetMicrosecondsOk returns a tuple with the Microseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroseconds

`func (o *Dbv0038JobTimeSystem) SetMicroseconds(v int32)`

SetMicroseconds sets Microseconds field to given value.

### HasMicroseconds

`func (o *Dbv0038JobTimeSystem) HasMicroseconds() bool`

HasMicroseconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


