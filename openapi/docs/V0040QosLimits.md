# V0040QosLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraceTime** | Pointer to **int32** |  | [optional] 
**Max** | Pointer to [**V0040QosLimitsMax**](V0040QosLimitsMax.md) |  | [optional] 
**Factor** | Pointer to [**V0040Float64NoVal**](V0040Float64NoVal.md) |  | [optional] 
**Min** | Pointer to [**V0040QosLimitsMin**](V0040QosLimitsMin.md) |  | [optional] 

## Methods

### NewV0040QosLimits

`func NewV0040QosLimits() *V0040QosLimits`

NewV0040QosLimits instantiates a new V0040QosLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040QosLimitsWithDefaults

`func NewV0040QosLimitsWithDefaults() *V0040QosLimits`

NewV0040QosLimitsWithDefaults instantiates a new V0040QosLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraceTime

`func (o *V0040QosLimits) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0040QosLimits) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0040QosLimits) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0040QosLimits) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMax

`func (o *V0040QosLimits) GetMax() V0040QosLimitsMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0040QosLimits) GetMaxOk() (*V0040QosLimitsMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0040QosLimits) SetMax(v V0040QosLimitsMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0040QosLimits) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetFactor

`func (o *V0040QosLimits) GetFactor() V0040Float64NoVal`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *V0040QosLimits) GetFactorOk() (*V0040Float64NoVal, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *V0040QosLimits) SetFactor(v V0040Float64NoVal)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *V0040QosLimits) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMin

`func (o *V0040QosLimits) GetMin() V0040QosLimitsMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0040QosLimits) GetMinOk() (*V0040QosLimitsMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0040QosLimits) SetMin(v V0040QosLimitsMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0040QosLimits) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


