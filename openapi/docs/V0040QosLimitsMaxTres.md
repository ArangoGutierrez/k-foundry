# V0040QosLimitsMaxTres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Minutes** | Pointer to [**V0040QosLimitsMaxTresMinutes**](V0040QosLimitsMaxTresMinutes.md) |  | [optional] 
**Per** | Pointer to [**V0040QosLimitsMaxTresPer**](V0040QosLimitsMaxTresPer.md) |  | [optional] 

## Methods

### NewV0040QosLimitsMaxTres

`func NewV0040QosLimitsMaxTres() *V0040QosLimitsMaxTres`

NewV0040QosLimitsMaxTres instantiates a new V0040QosLimitsMaxTres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040QosLimitsMaxTresWithDefaults

`func NewV0040QosLimitsMaxTresWithDefaults() *V0040QosLimitsMaxTres`

NewV0040QosLimitsMaxTresWithDefaults instantiates a new V0040QosLimitsMaxTres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0040QosLimitsMaxTres) GetTotal() []V0040Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0040QosLimitsMaxTres) GetTotalOk() (*[]V0040Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0040QosLimitsMaxTres) SetTotal(v []V0040Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0040QosLimitsMaxTres) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMinutes

`func (o *V0040QosLimitsMaxTres) GetMinutes() V0040QosLimitsMaxTresMinutes`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *V0040QosLimitsMaxTres) GetMinutesOk() (*V0040QosLimitsMaxTresMinutes, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *V0040QosLimitsMaxTres) SetMinutes(v V0040QosLimitsMaxTresMinutes)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *V0040QosLimitsMaxTres) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetPer

`func (o *V0040QosLimitsMaxTres) GetPer() V0040QosLimitsMaxTresPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0040QosLimitsMaxTres) GetPerOk() (*V0040QosLimitsMaxTresPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0040QosLimitsMaxTres) SetPer(v V0040QosLimitsMaxTresPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0040QosLimitsMaxTres) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


