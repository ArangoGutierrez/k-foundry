# V0039JobRequired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CPUs** | Pointer to **int32** |  | [optional] 
**MemoryPerCpu** | Pointer to [**V0039Uint64NoVal**](V0039Uint64NoVal.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0039Uint64NoVal**](V0039Uint64NoVal.md) |  | [optional] 
**Memory** | Pointer to **int64** |  | [optional] 

## Methods

### NewV0039JobRequired

`func NewV0039JobRequired() *V0039JobRequired`

NewV0039JobRequired instantiates a new V0039JobRequired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039JobRequiredWithDefaults

`func NewV0039JobRequiredWithDefaults() *V0039JobRequired`

NewV0039JobRequiredWithDefaults instantiates a new V0039JobRequired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCPUs

`func (o *V0039JobRequired) GetCPUs() int32`

GetCPUs returns the CPUs field if non-nil, zero value otherwise.

### GetCPUsOk

`func (o *V0039JobRequired) GetCPUsOk() (*int32, bool)`

GetCPUsOk returns a tuple with the CPUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPUs

`func (o *V0039JobRequired) SetCPUs(v int32)`

SetCPUs sets CPUs field to given value.

### HasCPUs

`func (o *V0039JobRequired) HasCPUs() bool`

HasCPUs returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0039JobRequired) GetMemoryPerCpu() V0039Uint64NoVal`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0039JobRequired) GetMemoryPerCpuOk() (*V0039Uint64NoVal, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0039JobRequired) SetMemoryPerCpu(v V0039Uint64NoVal)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0039JobRequired) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0039JobRequired) GetMemoryPerNode() V0039Uint64NoVal`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0039JobRequired) GetMemoryPerNodeOk() (*V0039Uint64NoVal, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0039JobRequired) SetMemoryPerNode(v V0039Uint64NoVal)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0039JobRequired) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetMemory

`func (o *V0039JobRequired) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *V0039JobRequired) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *V0039JobRequired) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *V0039JobRequired) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


