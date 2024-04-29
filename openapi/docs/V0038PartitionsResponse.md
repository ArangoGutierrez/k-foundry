# V0038PartitionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0038Meta**](V0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0038Error**](V0038Error.md) | slurm errors | [optional] 
**Partitions** | Pointer to [**[]V0038Partition**](V0038Partition.md) | partition info | [optional] 

## Methods

### NewV0038PartitionsResponse

`func NewV0038PartitionsResponse() *V0038PartitionsResponse`

NewV0038PartitionsResponse instantiates a new V0038PartitionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038PartitionsResponseWithDefaults

`func NewV0038PartitionsResponseWithDefaults() *V0038PartitionsResponse`

NewV0038PartitionsResponseWithDefaults instantiates a new V0038PartitionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0038PartitionsResponse) GetMeta() V0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0038PartitionsResponse) GetMetaOk() (*V0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0038PartitionsResponse) SetMeta(v V0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0038PartitionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0038PartitionsResponse) GetErrors() []V0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0038PartitionsResponse) GetErrorsOk() (*[]V0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0038PartitionsResponse) SetErrors(v []V0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0038PartitionsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetPartitions

`func (o *V0038PartitionsResponse) GetPartitions() []V0038Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0038PartitionsResponse) GetPartitionsOk() (*[]V0038Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0038PartitionsResponse) SetPartitions(v []V0038Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0038PartitionsResponse) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


