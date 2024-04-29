/*
Slurm Rest API

API to access and control Slurm DB.

API version: Slurm-23.11.5&openapi/dbv0.0.39&openapi/slurmctld&openapi/v0.0.39&openapi/dbv0.0.38&openapi/slurmdbd&openapi/v0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V0040AssocMaxTresMinutes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocMaxTresMinutes{}

// V0040AssocMaxTresMinutes struct for V0040AssocMaxTresMinutes
type V0040AssocMaxTresMinutes struct {
	Total []V0040Tres `json:"total,omitempty"`
	Per *V0040QosLimitsMinTresPer `json:"per,omitempty"`
}

// NewV0040AssocMaxTresMinutes instantiates a new V0040AssocMaxTresMinutes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocMaxTresMinutes() *V0040AssocMaxTresMinutes {
	this := V0040AssocMaxTresMinutes{}
	return &this
}

// NewV0040AssocMaxTresMinutesWithDefaults instantiates a new V0040AssocMaxTresMinutes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocMaxTresMinutesWithDefaults() *V0040AssocMaxTresMinutes {
	this := V0040AssocMaxTresMinutes{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0040AssocMaxTresMinutes) GetTotal() []V0040Tres {
	if o == nil || IsNil(o.Total) {
		var ret []V0040Tres
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxTresMinutes) GetTotalOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0040AssocMaxTresMinutes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given []V0040Tres and assigns it to the Total field.
func (o *V0040AssocMaxTresMinutes) SetTotal(v []V0040Tres) {
	o.Total = v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0040AssocMaxTresMinutes) GetPer() V0040QosLimitsMinTresPer {
	if o == nil || IsNil(o.Per) {
		var ret V0040QosLimitsMinTresPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMaxTresMinutes) GetPerOk() (*V0040QosLimitsMinTresPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0040AssocMaxTresMinutes) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0040QosLimitsMinTresPer and assigns it to the Per field.
func (o *V0040AssocMaxTresMinutes) SetPer(v V0040QosLimitsMinTresPer) {
	o.Per = &v
}

func (o V0040AssocMaxTresMinutes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocMaxTresMinutes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0040AssocMaxTresMinutes struct {
	value *V0040AssocMaxTresMinutes
	isSet bool
}

func (v NullableV0040AssocMaxTresMinutes) Get() *V0040AssocMaxTresMinutes {
	return v.value
}

func (v *NullableV0040AssocMaxTresMinutes) Set(val *V0040AssocMaxTresMinutes) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocMaxTresMinutes) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocMaxTresMinutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocMaxTresMinutes(val *V0040AssocMaxTresMinutes) *NullableV0040AssocMaxTresMinutes {
	return &NullableV0040AssocMaxTresMinutes{value: val, isSet: true}
}

func (v NullableV0040AssocMaxTresMinutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocMaxTresMinutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


