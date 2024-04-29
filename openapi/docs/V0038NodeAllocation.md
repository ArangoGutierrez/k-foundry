# V0038NodeAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to **int32** | amount of assigned job memory | [optional] 
**Cpus** | Pointer to **int32** | number of assigned job CPUs | [optional] 
**Sockets** | Pointer to [**V0038NodeAllocationSockets**](V0038NodeAllocationSockets.md) |  | [optional] 
**Nodename** | Pointer to **string** | node name | [optional] 

## Methods

### NewV0038NodeAllocation

`func NewV0038NodeAllocation() *V0038NodeAllocation`

NewV0038NodeAllocation instantiates a new V0038NodeAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038NodeAllocationWithDefaults

`func NewV0038NodeAllocationWithDefaults() *V0038NodeAllocation`

NewV0038NodeAllocationWithDefaults instantiates a new V0038NodeAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *V0038NodeAllocation) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *V0038NodeAllocation) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *V0038NodeAllocation) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *V0038NodeAllocation) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetCpus

`func (o *V0038NodeAllocation) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0038NodeAllocation) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0038NodeAllocation) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0038NodeAllocation) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetSockets

`func (o *V0038NodeAllocation) GetSockets() V0038NodeAllocationSockets`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0038NodeAllocation) GetSocketsOk() (*V0038NodeAllocationSockets, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0038NodeAllocation) SetSockets(v V0038NodeAllocationSockets)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *V0038NodeAllocation) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetNodename

`func (o *V0038NodeAllocation) GetNodename() string`

GetNodename returns the Nodename field if non-nil, zero value otherwise.

### GetNodenameOk

`func (o *V0038NodeAllocation) GetNodenameOk() (*string, bool)`

GetNodenameOk returns a tuple with the Nodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodename

`func (o *V0038NodeAllocation) SetNodename(v string)`

SetNodename sets Nodename field to given value.

### HasNodename

`func (o *V0038NodeAllocation) HasNodename() bool`

HasNodename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


