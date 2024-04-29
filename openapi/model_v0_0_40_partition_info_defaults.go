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

// checks if the V0040PartitionInfoDefaults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoDefaults{}

// V0040PartitionInfoDefaults struct for V0040PartitionInfoDefaults
type V0040PartitionInfoDefaults struct {
	MemoryPerCpu *int64 `json:"memory_per_cpu,omitempty"`
	PartitionMemoryPerCpu *V0040Uint64NoVal `json:"partition_memory_per_cpu,omitempty"`
	PartitionMemoryPerNode *V0040Uint64NoVal `json:"partition_memory_per_node,omitempty"`
	Time *V0040Uint32NoVal `json:"time,omitempty"`
	Job *string `json:"job,omitempty"`
}

// NewV0040PartitionInfoDefaults instantiates a new V0040PartitionInfoDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoDefaults() *V0040PartitionInfoDefaults {
	this := V0040PartitionInfoDefaults{}
	return &this
}

// NewV0040PartitionInfoDefaultsWithDefaults instantiates a new V0040PartitionInfoDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoDefaultsWithDefaults() *V0040PartitionInfoDefaults {
	this := V0040PartitionInfoDefaults{}
	return &this
}

// GetMemoryPerCpu returns the MemoryPerCpu field value if set, zero value otherwise.
func (o *V0040PartitionInfoDefaults) GetMemoryPerCpu() int64 {
	if o == nil || IsNil(o.MemoryPerCpu) {
		var ret int64
		return ret
	}
	return *o.MemoryPerCpu
}

// GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoDefaults) GetMemoryPerCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryPerCpu) {
		return nil, false
	}
	return o.MemoryPerCpu, true
}

// HasMemoryPerCpu returns a boolean if a field has been set.
func (o *V0040PartitionInfoDefaults) HasMemoryPerCpu() bool {
	if o != nil && !IsNil(o.MemoryPerCpu) {
		return true
	}

	return false
}

// SetMemoryPerCpu gets a reference to the given int64 and assigns it to the MemoryPerCpu field.
func (o *V0040PartitionInfoDefaults) SetMemoryPerCpu(v int64) {
	o.MemoryPerCpu = &v
}

// GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field value if set, zero value otherwise.
func (o *V0040PartitionInfoDefaults) GetPartitionMemoryPerCpu() V0040Uint64NoVal {
	if o == nil || IsNil(o.PartitionMemoryPerCpu) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.PartitionMemoryPerCpu
}

// GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoDefaults) GetPartitionMemoryPerCpuOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.PartitionMemoryPerCpu) {
		return nil, false
	}
	return o.PartitionMemoryPerCpu, true
}

// HasPartitionMemoryPerCpu returns a boolean if a field has been set.
func (o *V0040PartitionInfoDefaults) HasPartitionMemoryPerCpu() bool {
	if o != nil && !IsNil(o.PartitionMemoryPerCpu) {
		return true
	}

	return false
}

// SetPartitionMemoryPerCpu gets a reference to the given V0040Uint64NoVal and assigns it to the PartitionMemoryPerCpu field.
func (o *V0040PartitionInfoDefaults) SetPartitionMemoryPerCpu(v V0040Uint64NoVal) {
	o.PartitionMemoryPerCpu = &v
}

// GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field value if set, zero value otherwise.
func (o *V0040PartitionInfoDefaults) GetPartitionMemoryPerNode() V0040Uint64NoVal {
	if o == nil || IsNil(o.PartitionMemoryPerNode) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.PartitionMemoryPerNode
}

// GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoDefaults) GetPartitionMemoryPerNodeOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.PartitionMemoryPerNode) {
		return nil, false
	}
	return o.PartitionMemoryPerNode, true
}

// HasPartitionMemoryPerNode returns a boolean if a field has been set.
func (o *V0040PartitionInfoDefaults) HasPartitionMemoryPerNode() bool {
	if o != nil && !IsNil(o.PartitionMemoryPerNode) {
		return true
	}

	return false
}

// SetPartitionMemoryPerNode gets a reference to the given V0040Uint64NoVal and assigns it to the PartitionMemoryPerNode field.
func (o *V0040PartitionInfoDefaults) SetPartitionMemoryPerNode(v V0040Uint64NoVal) {
	o.PartitionMemoryPerNode = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0040PartitionInfoDefaults) GetTime() V0040Uint32NoVal {
	if o == nil || IsNil(o.Time) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoDefaults) GetTimeOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0040PartitionInfoDefaults) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0040Uint32NoVal and assigns it to the Time field.
func (o *V0040PartitionInfoDefaults) SetTime(v V0040Uint32NoVal) {
	o.Time = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0040PartitionInfoDefaults) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoDefaults) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0040PartitionInfoDefaults) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *V0040PartitionInfoDefaults) SetJob(v string) {
	o.Job = &v
}

func (o V0040PartitionInfoDefaults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoDefaults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemoryPerCpu) {
		toSerialize["memory_per_cpu"] = o.MemoryPerCpu
	}
	if !IsNil(o.PartitionMemoryPerCpu) {
		toSerialize["partition_memory_per_cpu"] = o.PartitionMemoryPerCpu
	}
	if !IsNil(o.PartitionMemoryPerNode) {
		toSerialize["partition_memory_per_node"] = o.PartitionMemoryPerNode
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoDefaults struct {
	value *V0040PartitionInfoDefaults
	isSet bool
}

func (v NullableV0040PartitionInfoDefaults) Get() *V0040PartitionInfoDefaults {
	return v.value
}

func (v *NullableV0040PartitionInfoDefaults) Set(val *V0040PartitionInfoDefaults) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoDefaults) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoDefaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoDefaults(val *V0040PartitionInfoDefaults) *NullableV0040PartitionInfoDefaults {
	return &NullableV0040PartitionInfoDefaults{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoDefaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


