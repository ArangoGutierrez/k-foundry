# Dbv0038Diag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 
**Statistics** | Pointer to [**Dbv0038DiagStatistics**](Dbv0038DiagStatistics.md) |  | [optional] 

## Methods

### NewDbv0038Diag

`func NewDbv0038Diag() *Dbv0038Diag`

NewDbv0038Diag instantiates a new Dbv0038Diag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038DiagWithDefaults

`func NewDbv0038DiagWithDefaults() *Dbv0038Diag`

NewDbv0038DiagWithDefaults instantiates a new Dbv0038Diag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038Diag) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038Diag) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038Diag) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038Diag) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038Diag) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038Diag) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038Diag) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038Diag) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatistics

`func (o *Dbv0038Diag) GetStatistics() Dbv0038DiagStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Dbv0038Diag) GetStatisticsOk() (*Dbv0038DiagStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Dbv0038Diag) SetStatistics(v Dbv0038DiagStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Dbv0038Diag) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


