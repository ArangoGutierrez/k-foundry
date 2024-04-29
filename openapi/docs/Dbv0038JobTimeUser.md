# Dbv0038JobTimeUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seconds** | Pointer to **int32** | Total number of CPU-seconds used by the job in user land, in seconds | [optional] 
**Microseconds** | Pointer to **int32** | Total number of CPU-seconds used by the job in user land, in microseconds | [optional] 

## Methods

### NewDbv0038JobTimeUser

`func NewDbv0038JobTimeUser() *Dbv0038JobTimeUser`

NewDbv0038JobTimeUser instantiates a new Dbv0038JobTimeUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobTimeUserWithDefaults

`func NewDbv0038JobTimeUserWithDefaults() *Dbv0038JobTimeUser`

NewDbv0038JobTimeUserWithDefaults instantiates a new Dbv0038JobTimeUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeconds

`func (o *Dbv0038JobTimeUser) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *Dbv0038JobTimeUser) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *Dbv0038JobTimeUser) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *Dbv0038JobTimeUser) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetMicroseconds

`func (o *Dbv0038JobTimeUser) GetMicroseconds() int32`

GetMicroseconds returns the Microseconds field if non-nil, zero value otherwise.

### GetMicrosecondsOk

`func (o *Dbv0038JobTimeUser) GetMicrosecondsOk() (*int32, bool)`

GetMicrosecondsOk returns a tuple with the Microseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroseconds

`func (o *Dbv0038JobTimeUser) SetMicroseconds(v int32)`

SetMicroseconds sets Microseconds field to given value.

### HasMicroseconds

`func (o *Dbv0038JobTimeUser) HasMicroseconds() bool`

HasMicroseconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


