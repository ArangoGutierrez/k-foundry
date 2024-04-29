# V0038Pings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0038Meta**](V0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0038Error**](V0038Error.md) | slurm errors | [optional] 
**Pings** | Pointer to [**[]V0038Ping**](V0038Ping.md) | slurm controller pings | [optional] 

## Methods

### NewV0038Pings

`func NewV0038Pings() *V0038Pings`

NewV0038Pings instantiates a new V0038Pings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038PingsWithDefaults

`func NewV0038PingsWithDefaults() *V0038Pings`

NewV0038PingsWithDefaults instantiates a new V0038Pings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0038Pings) GetMeta() V0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0038Pings) GetMetaOk() (*V0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0038Pings) SetMeta(v V0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0038Pings) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0038Pings) GetErrors() []V0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0038Pings) GetErrorsOk() (*[]V0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0038Pings) SetErrors(v []V0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0038Pings) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetPings

`func (o *V0038Pings) GetPings() []V0038Ping`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0038Pings) GetPingsOk() (*[]V0038Ping, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0038Pings) SetPings(v []V0038Ping)`

SetPings sets Pings field to given value.

### HasPings

`func (o *V0038Pings) HasPings() bool`

HasPings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


