# Dbv0038ResponseAssociationsDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 
**RemovedAssociations** | Pointer to **[]string** | the associations | [optional] 

## Methods

### NewDbv0038ResponseAssociationsDelete

`func NewDbv0038ResponseAssociationsDelete() *Dbv0038ResponseAssociationsDelete`

NewDbv0038ResponseAssociationsDelete instantiates a new Dbv0038ResponseAssociationsDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038ResponseAssociationsDeleteWithDefaults

`func NewDbv0038ResponseAssociationsDeleteWithDefaults() *Dbv0038ResponseAssociationsDelete`

NewDbv0038ResponseAssociationsDeleteWithDefaults instantiates a new Dbv0038ResponseAssociationsDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038ResponseAssociationsDelete) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038ResponseAssociationsDelete) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038ResponseAssociationsDelete) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038ResponseAssociationsDelete) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038ResponseAssociationsDelete) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038ResponseAssociationsDelete) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038ResponseAssociationsDelete) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038ResponseAssociationsDelete) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRemovedAssociations

`func (o *Dbv0038ResponseAssociationsDelete) GetRemovedAssociations() []string`

GetRemovedAssociations returns the RemovedAssociations field if non-nil, zero value otherwise.

### GetRemovedAssociationsOk

`func (o *Dbv0038ResponseAssociationsDelete) GetRemovedAssociationsOk() (*[]string, bool)`

GetRemovedAssociationsOk returns a tuple with the RemovedAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAssociations

`func (o *Dbv0038ResponseAssociationsDelete) SetRemovedAssociations(v []string)`

SetRemovedAssociations sets RemovedAssociations field to given value.

### HasRemovedAssociations

`func (o *Dbv0038ResponseAssociationsDelete) HasRemovedAssociations() bool`

HasRemovedAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


