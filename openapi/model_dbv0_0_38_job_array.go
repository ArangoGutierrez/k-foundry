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

// checks if the Dbv0038JobArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobArray{}

// Dbv0038JobArray Array properties (optional)
type Dbv0038JobArray struct {
	// Job id of array
	JobId *int32 `json:"job_id,omitempty"`
	Limits *Dbv0038JobArrayLimits `json:"limits,omitempty"`
	// Array task
	Task *string `json:"task,omitempty"`
	// Array task id
	TaskId *int32 `json:"task_id,omitempty"`
}

// NewDbv0038JobArray instantiates a new Dbv0038JobArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobArray() *Dbv0038JobArray {
	this := Dbv0038JobArray{}
	return &this
}

// NewDbv0038JobArrayWithDefaults instantiates a new Dbv0038JobArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobArrayWithDefaults() *Dbv0038JobArray {
	this := Dbv0038JobArray{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *Dbv0038JobArray) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobArray) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *Dbv0038JobArray) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *Dbv0038JobArray) SetJobId(v int32) {
	o.JobId = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *Dbv0038JobArray) GetLimits() Dbv0038JobArrayLimits {
	if o == nil || IsNil(o.Limits) {
		var ret Dbv0038JobArrayLimits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobArray) GetLimitsOk() (*Dbv0038JobArrayLimits, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *Dbv0038JobArray) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given Dbv0038JobArrayLimits and assigns it to the Limits field.
func (o *Dbv0038JobArray) SetLimits(v Dbv0038JobArrayLimits) {
	o.Limits = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *Dbv0038JobArray) GetTask() string {
	if o == nil || IsNil(o.Task) {
		var ret string
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobArray) GetTaskOk() (*string, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *Dbv0038JobArray) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given string and assigns it to the Task field.
func (o *Dbv0038JobArray) SetTask(v string) {
	o.Task = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *Dbv0038JobArray) GetTaskId() int32 {
	if o == nil || IsNil(o.TaskId) {
		var ret int32
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobArray) GetTaskIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *Dbv0038JobArray) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.
func (o *Dbv0038JobArray) SetTaskId(v int32) {
	o.TaskId = &v
}

func (o Dbv0038JobArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.TaskId) {
		toSerialize["task_id"] = o.TaskId
	}
	return toSerialize, nil
}

type NullableDbv0038JobArray struct {
	value *Dbv0038JobArray
	isSet bool
}

func (v NullableDbv0038JobArray) Get() *Dbv0038JobArray {
	return v.value
}

func (v *NullableDbv0038JobArray) Set(val *Dbv0038JobArray) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobArray) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobArray(val *Dbv0038JobArray) *NullableDbv0038JobArray {
	return &NullableDbv0038JobArray{value: val, isSet: true}
}

func (v NullableDbv0038JobArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


