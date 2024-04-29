# V0040OpenapiAssocsRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovedAssociations** | **[]string** |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiAssocsRemovedResp

`func NewV0040OpenapiAssocsRemovedResp(removedAssociations []string, ) *V0040OpenapiAssocsRemovedResp`

NewV0040OpenapiAssocsRemovedResp instantiates a new V0040OpenapiAssocsRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiAssocsRemovedRespWithDefaults

`func NewV0040OpenapiAssocsRemovedRespWithDefaults() *V0040OpenapiAssocsRemovedResp`

NewV0040OpenapiAssocsRemovedRespWithDefaults instantiates a new V0040OpenapiAssocsRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovedAssociations

`func (o *V0040OpenapiAssocsRemovedResp) GetRemovedAssociations() []string`

GetRemovedAssociations returns the RemovedAssociations field if non-nil, zero value otherwise.

### GetRemovedAssociationsOk

`func (o *V0040OpenapiAssocsRemovedResp) GetRemovedAssociationsOk() (*[]string, bool)`

GetRemovedAssociationsOk returns a tuple with the RemovedAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAssociations

`func (o *V0040OpenapiAssocsRemovedResp) SetRemovedAssociations(v []string)`

SetRemovedAssociations sets RemovedAssociations field to given value.


### GetMeta

`func (o *V0040OpenapiAssocsRemovedResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiAssocsRemovedResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiAssocsRemovedResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiAssocsRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiAssocsRemovedResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiAssocsRemovedResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiAssocsRemovedResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiAssocsRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiAssocsRemovedResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiAssocsRemovedResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiAssocsRemovedResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiAssocsRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


