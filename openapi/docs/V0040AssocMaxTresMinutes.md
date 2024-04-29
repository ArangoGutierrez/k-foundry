# V0040AssocMaxTresMinutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Per** | Pointer to [**V0040QosLimitsMinTresPer**](V0040QosLimitsMinTresPer.md) |  | [optional] 

## Methods

### NewV0040AssocMaxTresMinutes

`func NewV0040AssocMaxTresMinutes() *V0040AssocMaxTresMinutes`

NewV0040AssocMaxTresMinutes instantiates a new V0040AssocMaxTresMinutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040AssocMaxTresMinutesWithDefaults

`func NewV0040AssocMaxTresMinutesWithDefaults() *V0040AssocMaxTresMinutes`

NewV0040AssocMaxTresMinutesWithDefaults instantiates a new V0040AssocMaxTresMinutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0040AssocMaxTresMinutes) GetTotal() []V0040Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0040AssocMaxTresMinutes) GetTotalOk() (*[]V0040Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0040AssocMaxTresMinutes) SetTotal(v []V0040Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0040AssocMaxTresMinutes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPer

`func (o *V0040AssocMaxTresMinutes) GetPer() V0040QosLimitsMinTresPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0040AssocMaxTresMinutes) GetPerOk() (*V0040QosLimitsMinTresPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0040AssocMaxTresMinutes) SetPer(v V0040QosLimitsMinTresPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0040AssocMaxTresMinutes) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


