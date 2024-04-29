# V0040OpenapiWckeyRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedWckeys** | **[]string** |  | 
**Meta** | Pointer to [**V0040OpenapiMeta**](V0040OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0040OpenapiError**](V0040OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0040OpenapiWarning**](V0040OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0040OpenapiWckeyRemovedResp

`func NewV0040OpenapiWckeyRemovedResp(deletedWckeys []string, ) *V0040OpenapiWckeyRemovedResp`

NewV0040OpenapiWckeyRemovedResp instantiates a new V0040OpenapiWckeyRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiWckeyRemovedRespWithDefaults

`func NewV0040OpenapiWckeyRemovedRespWithDefaults() *V0040OpenapiWckeyRemovedResp`

NewV0040OpenapiWckeyRemovedRespWithDefaults instantiates a new V0040OpenapiWckeyRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedWckeys

`func (o *V0040OpenapiWckeyRemovedResp) GetDeletedWckeys() []string`

GetDeletedWckeys returns the DeletedWckeys field if non-nil, zero value otherwise.

### GetDeletedWckeysOk

`func (o *V0040OpenapiWckeyRemovedResp) GetDeletedWckeysOk() (*[]string, bool)`

GetDeletedWckeysOk returns a tuple with the DeletedWckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedWckeys

`func (o *V0040OpenapiWckeyRemovedResp) SetDeletedWckeys(v []string)`

SetDeletedWckeys sets DeletedWckeys field to given value.


### GetMeta

`func (o *V0040OpenapiWckeyRemovedResp) GetMeta() V0040OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0040OpenapiWckeyRemovedResp) GetMetaOk() (*V0040OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0040OpenapiWckeyRemovedResp) SetMeta(v V0040OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0040OpenapiWckeyRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0040OpenapiWckeyRemovedResp) GetErrors() []V0040OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0040OpenapiWckeyRemovedResp) GetErrorsOk() (*[]V0040OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0040OpenapiWckeyRemovedResp) SetErrors(v []V0040OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0040OpenapiWckeyRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0040OpenapiWckeyRemovedResp) GetWarnings() []V0040OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0040OpenapiWckeyRemovedResp) GetWarningsOk() (*[]V0040OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0040OpenapiWckeyRemovedResp) SetWarnings(v []V0040OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0040OpenapiWckeyRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


