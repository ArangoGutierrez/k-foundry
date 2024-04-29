# V0039QosLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraceTime** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 
**Max** | Pointer to [**V0039QosLimitsMax**](V0039QosLimitsMax.md) |  | [optional] 
**Factor** | Pointer to **float64** |  | [optional] 
**Min** | Pointer to [**V0039QosLimitsMin**](V0039QosLimitsMin.md) |  | [optional] 

## Methods

### NewV0039QosLimits

`func NewV0039QosLimits() *V0039QosLimits`

NewV0039QosLimits instantiates a new V0039QosLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039QosLimitsWithDefaults

`func NewV0039QosLimitsWithDefaults() *V0039QosLimits`

NewV0039QosLimitsWithDefaults instantiates a new V0039QosLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraceTime

`func (o *V0039QosLimits) GetGraceTime() V0039Uint32NoVal`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0039QosLimits) GetGraceTimeOk() (*V0039Uint32NoVal, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0039QosLimits) SetGraceTime(v V0039Uint32NoVal)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0039QosLimits) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMax

`func (o *V0039QosLimits) GetMax() V0039QosLimitsMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0039QosLimits) GetMaxOk() (*V0039QosLimitsMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0039QosLimits) SetMax(v V0039QosLimitsMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0039QosLimits) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetFactor

`func (o *V0039QosLimits) GetFactor() float64`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *V0039QosLimits) GetFactorOk() (*float64, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *V0039QosLimits) SetFactor(v float64)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *V0039QosLimits) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMin

`func (o *V0039QosLimits) GetMin() V0039QosLimitsMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0039QosLimits) GetMinOk() (*V0039QosLimitsMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0039QosLimits) SetMin(v V0039QosLimitsMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0039QosLimits) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


