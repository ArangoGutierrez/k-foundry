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

// checks if the Dbv0038JobStepTasks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobStepTasks{}

// Dbv0038JobStepTasks Task properties
type Dbv0038JobStepTasks struct {
	// Number of tasks in step
	Count *int32 `json:"count,omitempty"`
}

// NewDbv0038JobStepTasks instantiates a new Dbv0038JobStepTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobStepTasks() *Dbv0038JobStepTasks {
	this := Dbv0038JobStepTasks{}
	return &this
}

// NewDbv0038JobStepTasksWithDefaults instantiates a new Dbv0038JobStepTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobStepTasksWithDefaults() *Dbv0038JobStepTasks {
	this := Dbv0038JobStepTasks{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Dbv0038JobStepTasks) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStepTasks) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Dbv0038JobStepTasks) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Dbv0038JobStepTasks) SetCount(v int32) {
	o.Count = &v
}

func (o Dbv0038JobStepTasks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobStepTasks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableDbv0038JobStepTasks struct {
	value *Dbv0038JobStepTasks
	isSet bool
}

func (v NullableDbv0038JobStepTasks) Get() *Dbv0038JobStepTasks {
	return v.value
}

func (v *NullableDbv0038JobStepTasks) Set(val *Dbv0038JobStepTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobStepTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobStepTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobStepTasks(val *Dbv0038JobStepTasks) *NullableDbv0038JobStepTasks {
	return &NullableDbv0038JobStepTasks{value: val, isSet: true}
}

func (v NullableDbv0038JobStepTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobStepTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


