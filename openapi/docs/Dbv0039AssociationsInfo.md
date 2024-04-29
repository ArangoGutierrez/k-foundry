# Dbv0039AssociationsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0039Meta**](Dbv0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0039Error**](Dbv0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]Dbv0039Warning**](Dbv0039Warning.md) | Slurm warnings | [optional] 
**Associations** | Pointer to [**[]V0039Assoc**](V0039Assoc.md) |  | [optional] 

## Methods

### NewDbv0039AssociationsInfo

`func NewDbv0039AssociationsInfo() *Dbv0039AssociationsInfo`

NewDbv0039AssociationsInfo instantiates a new Dbv0039AssociationsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0039AssociationsInfoWithDefaults

`func NewDbv0039AssociationsInfoWithDefaults() *Dbv0039AssociationsInfo`

NewDbv0039AssociationsInfoWithDefaults instantiates a new Dbv0039AssociationsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0039AssociationsInfo) GetMeta() Dbv0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0039AssociationsInfo) GetMetaOk() (*Dbv0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0039AssociationsInfo) SetMeta(v Dbv0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0039AssociationsInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0039AssociationsInfo) GetErrors() []Dbv0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0039AssociationsInfo) GetErrorsOk() (*[]Dbv0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0039AssociationsInfo) SetErrors(v []Dbv0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0039AssociationsInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *Dbv0039AssociationsInfo) GetWarnings() []Dbv0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Dbv0039AssociationsInfo) GetWarningsOk() (*[]Dbv0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Dbv0039AssociationsInfo) SetWarnings(v []Dbv0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Dbv0039AssociationsInfo) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0039AssociationsInfo) GetAssociations() []V0039Assoc`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0039AssociationsInfo) GetAssociationsOk() (*[]V0039Assoc, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0039AssociationsInfo) SetAssociations(v []V0039Assoc)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0039AssociationsInfo) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


