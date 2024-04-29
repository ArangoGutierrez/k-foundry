# Dbv0038QosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 
**Qos** | Pointer to [**[]Dbv0038Qos**](Dbv0038Qos.md) | Array of QOS | [optional] 

## Methods

### NewDbv0038QosInfo

`func NewDbv0038QosInfo() *Dbv0038QosInfo`

NewDbv0038QosInfo instantiates a new Dbv0038QosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038QosInfoWithDefaults

`func NewDbv0038QosInfoWithDefaults() *Dbv0038QosInfo`

NewDbv0038QosInfoWithDefaults instantiates a new Dbv0038QosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038QosInfo) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038QosInfo) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038QosInfo) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038QosInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038QosInfo) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038QosInfo) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038QosInfo) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038QosInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0038QosInfo) GetQos() []Dbv0038Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0038QosInfo) GetQosOk() (*[]Dbv0038Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0038QosInfo) SetQos(v []Dbv0038Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0038QosInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


