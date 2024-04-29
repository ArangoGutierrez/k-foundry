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

// checks if the V0039StepTasks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepTasks{}

// V0039StepTasks struct for V0039StepTasks
type V0039StepTasks struct {
	Count *int32 `json:"count,omitempty"`
}

// NewV0039StepTasks instantiates a new V0039StepTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepTasks() *V0039StepTasks {
	this := V0039StepTasks{}
	return &this
}

// NewV0039StepTasksWithDefaults instantiates a new V0039StepTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepTasksWithDefaults() *V0039StepTasks {
	this := V0039StepTasks{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0039StepTasks) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTasks) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0039StepTasks) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0039StepTasks) SetCount(v int32) {
	o.Count = &v
}

func (o V0039StepTasks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepTasks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableV0039StepTasks struct {
	value *V0039StepTasks
	isSet bool
}

func (v NullableV0039StepTasks) Get() *V0039StepTasks {
	return v.value
}

func (v *NullableV0039StepTasks) Set(val *V0039StepTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepTasks(val *V0039StepTasks) *NullableV0039StepTasks {
	return &NullableV0039StepTasks{value: val, isSet: true}
}

func (v NullableV0039StepTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

