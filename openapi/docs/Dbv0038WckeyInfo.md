# Dbv0038WckeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 
**Wckeys** | Pointer to [**[]Dbv0038Wckey**](Dbv0038Wckey.md) | List of wckeys | [optional] 

## Methods

### NewDbv0038WckeyInfo

`func NewDbv0038WckeyInfo() *Dbv0038WckeyInfo`

NewDbv0038WckeyInfo instantiates a new Dbv0038WckeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038WckeyInfoWithDefaults

`func NewDbv0038WckeyInfoWithDefaults() *Dbv0038WckeyInfo`

NewDbv0038WckeyInfoWithDefaults instantiates a new Dbv0038WckeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038WckeyInfo) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038WckeyInfo) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038WckeyInfo) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038WckeyInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038WckeyInfo) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038WckeyInfo) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038WckeyInfo) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038WckeyInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWckeys

`func (o *Dbv0038WckeyInfo) GetWckeys() []Dbv0038Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *Dbv0038WckeyInfo) GetWckeysOk() (*[]Dbv0038Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *Dbv0038WckeyInfo) SetWckeys(v []Dbv0038Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *Dbv0038WckeyInfo) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


