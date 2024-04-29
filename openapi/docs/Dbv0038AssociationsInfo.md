# Dbv0038AssociationsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 
**Associations** | Pointer to [**[]Dbv0038Association**](Dbv0038Association.md) | Array of associations | [optional] 

## Methods

### NewDbv0038AssociationsInfo

`func NewDbv0038AssociationsInfo() *Dbv0038AssociationsInfo`

NewDbv0038AssociationsInfo instantiates a new Dbv0038AssociationsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038AssociationsInfoWithDefaults

`func NewDbv0038AssociationsInfoWithDefaults() *Dbv0038AssociationsInfo`

NewDbv0038AssociationsInfoWithDefaults instantiates a new Dbv0038AssociationsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038AssociationsInfo) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038AssociationsInfo) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038AssociationsInfo) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038AssociationsInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038AssociationsInfo) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038AssociationsInfo) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038AssociationsInfo) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038AssociationsInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0038AssociationsInfo) GetAssociations() []Dbv0038Association`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0038AssociationsInfo) GetAssociationsOk() (*[]Dbv0038Association, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0038AssociationsInfo) SetAssociations(v []Dbv0038Association)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0038AssociationsInfo) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


