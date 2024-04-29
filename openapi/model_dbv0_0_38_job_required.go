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

// checks if the Dbv0038JobRequired type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobRequired{}

// Dbv0038JobRequired Job run requirements
type Dbv0038JobRequired struct {
	// Required number of CPUs
	CPUs *int32 `json:"CPUs,omitempty"`
	// Required amount of memory (MiB)
	Memory *int32 `json:"memory,omitempty"`
}

// NewDbv0038JobRequired instantiates a new Dbv0038JobRequired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobRequired() *Dbv0038JobRequired {
	this := Dbv0038JobRequired{}
	return &this
}

// NewDbv0038JobRequiredWithDefaults instantiates a new Dbv0038JobRequired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobRequiredWithDefaults() *Dbv0038JobRequired {
	this := Dbv0038JobRequired{}
	return &this
}

// GetCPUs returns the CPUs field value if set, zero value otherwise.
func (o *Dbv0038JobRequired) GetCPUs() int32 {
	if o == nil || IsNil(o.CPUs) {
		var ret int32
		return ret
	}
	return *o.CPUs
}

// GetCPUsOk returns a tuple with the CPUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobRequired) GetCPUsOk() (*int32, bool) {
	if o == nil || IsNil(o.CPUs) {
		return nil, false
	}
	return o.CPUs, true
}

// HasCPUs returns a boolean if a field has been set.
func (o *Dbv0038JobRequired) HasCPUs() bool {
	if o != nil && !IsNil(o.CPUs) {
		return true
	}

	return false
}

// SetCPUs gets a reference to the given int32 and assigns it to the CPUs field.
func (o *Dbv0038JobRequired) SetCPUs(v int32) {
	o.CPUs = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Dbv0038JobRequired) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobRequired) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Dbv0038JobRequired) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *Dbv0038JobRequired) SetMemory(v int32) {
	o.Memory = &v
}

func (o Dbv0038JobRequired) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobRequired) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CPUs) {
		toSerialize["CPUs"] = o.CPUs
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableDbv0038JobRequired struct {
	value *Dbv0038JobRequired
	isSet bool
}

func (v NullableDbv0038JobRequired) Get() *Dbv0038JobRequired {
	return v.value
}

func (v *NullableDbv0038JobRequired) Set(val *Dbv0038JobRequired) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobRequired(val *Dbv0038JobRequired) *NullableDbv0038JobRequired {
	return &NullableDbv0038JobRequired{value: val, isSet: true}
}

func (v NullableDbv0038JobRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

