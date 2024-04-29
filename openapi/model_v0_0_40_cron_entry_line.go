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

// checks if the V0040CronEntryLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040CronEntryLine{}

// V0040CronEntryLine struct for V0040CronEntryLine
type V0040CronEntryLine struct {
	Start *int32 `json:"start,omitempty"`
	End *int32 `json:"end,omitempty"`
}

// NewV0040CronEntryLine instantiates a new V0040CronEntryLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040CronEntryLine() *V0040CronEntryLine {
	this := V0040CronEntryLine{}
	return &this
}

// NewV0040CronEntryLineWithDefaults instantiates a new V0040CronEntryLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040CronEntryLineWithDefaults() *V0040CronEntryLine {
	this := V0040CronEntryLine{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *V0040CronEntryLine) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040CronEntryLine) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *V0040CronEntryLine) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *V0040CronEntryLine) SetStart(v int32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *V0040CronEntryLine) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040CronEntryLine) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *V0040CronEntryLine) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *V0040CronEntryLine) SetEnd(v int32) {
	o.End = &v
}

func (o V0040CronEntryLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040CronEntryLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableV0040CronEntryLine struct {
	value *V0040CronEntryLine
	isSet bool
}

func (v NullableV0040CronEntryLine) Get() *V0040CronEntryLine {
	return v.value
}

func (v *NullableV0040CronEntryLine) Set(val *V0040CronEntryLine) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040CronEntryLine) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040CronEntryLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040CronEntryLine(val *V0040CronEntryLine) *NullableV0040CronEntryLine {
	return &NullableV0040CronEntryLine{value: val, isSet: true}
}

func (v NullableV0040CronEntryLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040CronEntryLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

