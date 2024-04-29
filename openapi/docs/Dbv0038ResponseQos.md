# Dbv0038ResponseQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 

## Methods

### NewDbv0038ResponseQos

`func NewDbv0038ResponseQos() *Dbv0038ResponseQos`

NewDbv0038ResponseQos instantiates a new Dbv0038ResponseQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038ResponseQosWithDefaults

`func NewDbv0038ResponseQosWithDefaults() *Dbv0038ResponseQos`

NewDbv0038ResponseQosWithDefaults instantiates a new Dbv0038ResponseQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038ResponseQos) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038ResponseQos) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038ResponseQos) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038ResponseQos) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038ResponseQos) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038ResponseQos) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038ResponseQos) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038ResponseQos) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


