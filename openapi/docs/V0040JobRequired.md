# V0040JobRequired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CPUs** | Pointer to **int32** |  | [optional] 
**MemoryPerCpu** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0040Uint64NoVal**](V0040Uint64NoVal.md) |  | [optional] 

## Methods

### NewV0040JobRequired

`func NewV0040JobRequired() *V0040JobRequired`

NewV0040JobRequired instantiates a new V0040JobRequired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040JobRequiredWithDefaults

`func NewV0040JobRequiredWithDefaults() *V0040JobRequired`

NewV0040JobRequiredWithDefaults instantiates a new V0040JobRequired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCPUs

`func (o *V0040JobRequired) GetCPUs() int32`

GetCPUs returns the CPUs field if non-nil, zero value otherwise.

### GetCPUsOk

`func (o *V0040JobRequired) GetCPUsOk() (*int32, bool)`

GetCPUsOk returns a tuple with the CPUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPUs

`func (o *V0040JobRequired) SetCPUs(v int32)`

SetCPUs sets CPUs field to given value.

### HasCPUs

`func (o *V0040JobRequired) HasCPUs() bool`

HasCPUs returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0040JobRequired) GetMemoryPerCpu() V0040Uint64NoVal`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0040JobRequired) GetMemoryPerCpuOk() (*V0040Uint64NoVal, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0040JobRequired) SetMemoryPerCpu(v V0040Uint64NoVal)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0040JobRequired) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0040JobRequired) GetMemoryPerNode() V0040Uint64NoVal`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0040JobRequired) GetMemoryPerNodeOk() (*V0040Uint64NoVal, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0040JobRequired) SetMemoryPerNode(v V0040Uint64NoVal)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0040JobRequired) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


