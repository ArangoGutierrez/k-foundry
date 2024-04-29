# Dbv0039WckeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0039Meta**](Dbv0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0039Error**](Dbv0039Error.md) | Slurm errors | [optional] 
**Wckeys** | Pointer to [**[]V0039Wckey**](V0039Wckey.md) |  | [optional] 

## Methods

### NewDbv0039WckeyInfo

`func NewDbv0039WckeyInfo() *Dbv0039WckeyInfo`

NewDbv0039WckeyInfo instantiates a new Dbv0039WckeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0039WckeyInfoWithDefaults

`func NewDbv0039WckeyInfoWithDefaults() *Dbv0039WckeyInfo`

NewDbv0039WckeyInfoWithDefaults instantiates a new Dbv0039WckeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0039WckeyInfo) GetMeta() Dbv0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0039WckeyInfo) GetMetaOk() (*Dbv0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0039WckeyInfo) SetMeta(v Dbv0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0039WckeyInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0039WckeyInfo) GetErrors() []Dbv0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0039WckeyInfo) GetErrorsOk() (*[]Dbv0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0039WckeyInfo) SetErrors(v []Dbv0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0039WckeyInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWckeys

`func (o *Dbv0039WckeyInfo) GetWckeys() []V0039Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *Dbv0039WckeyInfo) GetWckeysOk() (*[]V0039Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *Dbv0039WckeyInfo) SetWckeys(v []V0039Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *Dbv0039WckeyInfo) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


