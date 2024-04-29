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

// checks if the V0039JobArrayLimitsMaxRunning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobArrayLimitsMaxRunning{}

// V0039JobArrayLimitsMaxRunning struct for V0039JobArrayLimitsMaxRunning
type V0039JobArrayLimitsMaxRunning struct {
	Tasks *int32 `json:"tasks,omitempty"`
}

// NewV0039JobArrayLimitsMaxRunning instantiates a new V0039JobArrayLimitsMaxRunning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobArrayLimitsMaxRunning() *V0039JobArrayLimitsMaxRunning {
	this := V0039JobArrayLimitsMaxRunning{}
	return &this
}

// NewV0039JobArrayLimitsMaxRunningWithDefaults instantiates a new V0039JobArrayLimitsMaxRunning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobArrayLimitsMaxRunningWithDefaults() *V0039JobArrayLimitsMaxRunning {
	this := V0039JobArrayLimitsMaxRunning{}
	return &this
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *V0039JobArrayLimitsMaxRunning) GetTasks() int32 {
	if o == nil || IsNil(o.Tasks) {
		var ret int32
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobArrayLimitsMaxRunning) GetTasksOk() (*int32, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *V0039JobArrayLimitsMaxRunning) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given int32 and assigns it to the Tasks field.
func (o *V0039JobArrayLimitsMaxRunning) SetTasks(v int32) {
	o.Tasks = &v
}

func (o V0039JobArrayLimitsMaxRunning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobArrayLimitsMaxRunning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableV0039JobArrayLimitsMaxRunning struct {
	value *V0039JobArrayLimitsMaxRunning
	isSet bool
}

func (v NullableV0039JobArrayLimitsMaxRunning) Get() *V0039JobArrayLimitsMaxRunning {
	return v.value
}

func (v *NullableV0039JobArrayLimitsMaxRunning) Set(val *V0039JobArrayLimitsMaxRunning) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobArrayLimitsMaxRunning) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobArrayLimitsMaxRunning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobArrayLimitsMaxRunning(val *V0039JobArrayLimitsMaxRunning) *NullableV0039JobArrayLimitsMaxRunning {
	return &NullableV0039JobArrayLimitsMaxRunning{value: val, isSet: true}
}

func (v NullableV0039JobArrayLimitsMaxRunning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobArrayLimitsMaxRunning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

