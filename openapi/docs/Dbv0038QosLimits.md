# Dbv0038QosLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Factor** | Pointer to **float32** | factor to apply to TRES count for associations using this QOS | [optional] 
**Max** | Pointer to [**Dbv0038QosLimitsMax**](Dbv0038QosLimitsMax.md) |  | [optional] 
**Min** | Pointer to [**Dbv0038QosLimitsMin**](Dbv0038QosLimitsMin.md) |  | [optional] 

## Methods

### NewDbv0038QosLimits

`func NewDbv0038QosLimits() *Dbv0038QosLimits`

NewDbv0038QosLimits instantiates a new Dbv0038QosLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038QosLimitsWithDefaults

`func NewDbv0038QosLimitsWithDefaults() *Dbv0038QosLimits`

NewDbv0038QosLimitsWithDefaults instantiates a new Dbv0038QosLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactor

`func (o *Dbv0038QosLimits) GetFactor() float32`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *Dbv0038QosLimits) GetFactorOk() (*float32, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *Dbv0038QosLimits) SetFactor(v float32)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *Dbv0038QosLimits) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMax

`func (o *Dbv0038QosLimits) GetMax() Dbv0038QosLimitsMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Dbv0038QosLimits) GetMaxOk() (*Dbv0038QosLimitsMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Dbv0038QosLimits) SetMax(v Dbv0038QosLimitsMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *Dbv0038QosLimits) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *Dbv0038QosLimits) GetMin() Dbv0038QosLimitsMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *Dbv0038QosLimits) GetMinOk() (*Dbv0038QosLimitsMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *Dbv0038QosLimits) SetMin(v Dbv0038QosLimitsMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *Dbv0038QosLimits) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


