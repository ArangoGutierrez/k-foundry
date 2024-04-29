# V0038Licenses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0038Error**](V0038Error.md) | slurm errors | [optional] 
**Licenses** | Pointer to [**[]V0038License**](V0038License.md) | licenses info | [optional] 

## Methods

### NewV0038Licenses

`func NewV0038Licenses() *V0038Licenses`

NewV0038Licenses instantiates a new V0038Licenses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038LicensesWithDefaults

`func NewV0038LicensesWithDefaults() *V0038Licenses`

NewV0038LicensesWithDefaults instantiates a new V0038Licenses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0038Licenses) GetErrors() []V0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0038Licenses) GetErrorsOk() (*[]V0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0038Licenses) SetErrors(v []V0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0038Licenses) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLicenses

`func (o *V0038Licenses) GetLicenses() []V0038License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0038Licenses) GetLicensesOk() (*[]V0038License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0038Licenses) SetLicenses(v []V0038License)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0038Licenses) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


