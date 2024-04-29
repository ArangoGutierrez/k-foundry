# V0040OpenapiSlurmdbdQosRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovedQos** | **[]string** |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiSlurmdbdQosRemovedResp

`func NewV0040OpenapiSlurmdbdQosRemovedResp(removedQos []string, ) *V0040OpenapiSlurmdbdQosRemovedResp`

NewV0040OpenapiSlurmdbdQosRemovedResp instantiates a new V0040OpenapiSlurmdbdQosRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiSlurmdbdQosRemovedRespWithDefaults

`func NewV0040OpenapiSlurmdbdQosRemovedRespWithDefaults() *V0040OpenapiSlurmdbdQosRemovedResp`

NewV0040OpenapiSlurmdbdQosRemovedRespWithDefaults instantiates a new V0040OpenapiSlurmdbdQosRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovedQos

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetRemovedQos() []string`

GetRemovedQos returns the RemovedQos field if non-nil, zero value otherwise.

### GetRemovedQosOk

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetRemovedQosOk() (*[]string, bool)`

GetRemovedQosOk returns a tuple with the RemovedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedQos

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) SetRemovedQos(v []string)`

SetRemovedQos sets RemovedQos field to given value.


### GetMeta

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiSlurmdbdQosRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


