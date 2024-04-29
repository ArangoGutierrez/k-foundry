# Dbv0039QosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0039Meta**](Dbv0039Meta.md) |  | [optional] 
**Warnings** | Pointer to [**[]Dbv0039Warning**](Dbv0039Warning.md) | Slurm warnings | [optional] 
**Errors** | Pointer to [**[]Dbv0039Error**](Dbv0039Error.md) | Slurm errors | [optional] 
**Qos** | Pointer to [**[]V0039Qos**](V0039Qos.md) |  | [optional] 

## Methods

### NewDbv0039QosInfo

`func NewDbv0039QosInfo() *Dbv0039QosInfo`

NewDbv0039QosInfo instantiates a new Dbv0039QosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0039QosInfoWithDefaults

`func NewDbv0039QosInfoWithDefaults() *Dbv0039QosInfo`

NewDbv0039QosInfoWithDefaults instantiates a new Dbv0039QosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0039QosInfo) GetMeta() Dbv0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0039QosInfo) GetMetaOk() (*Dbv0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0039QosInfo) SetMeta(v Dbv0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0039QosInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWarnings

`func (o *Dbv0039QosInfo) GetWarnings() []Dbv0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Dbv0039QosInfo) GetWarningsOk() (*[]Dbv0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Dbv0039QosInfo) SetWarnings(v []Dbv0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Dbv0039QosInfo) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0039QosInfo) GetErrors() []Dbv0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0039QosInfo) GetErrorsOk() (*[]Dbv0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0039QosInfo) SetErrors(v []Dbv0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0039QosInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0039QosInfo) GetQos() []V0039Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0039QosInfo) GetQosOk() (*[]V0039Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0039QosInfo) SetQos(v []V0039Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0039QosInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


