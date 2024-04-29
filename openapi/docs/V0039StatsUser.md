# V0039StatsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to [**V0039StatsRpcTime**](V0039StatsRpcTime.md) |  | [optional] 

## Methods

### NewV0039StatsUser

`func NewV0039StatsUser() *V0039StatsUser`

NewV0039StatsUser instantiates a new V0039StatsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039StatsUserWithDefaults

`func NewV0039StatsUserWithDefaults() *V0039StatsUser`

NewV0039StatsUserWithDefaults instantiates a new V0039StatsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *V0039StatsUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0039StatsUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0039StatsUser) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V0039StatsUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCount

`func (o *V0039StatsUser) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0039StatsUser) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0039StatsUser) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0039StatsUser) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTime

`func (o *V0039StatsUser) GetTime() V0039StatsRpcTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0039StatsUser) GetTimeOk() (*V0039StatsRpcTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0039StatsUser) SetTime(v V0039StatsRpcTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0039StatsUser) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


