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

// checks if the Dbv0038QosLimitsMaxJobsActiveJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosLimitsMaxJobsActiveJobs{}

// Dbv0038QosLimitsMaxJobsActiveJobs Limits on active jobs settings
type Dbv0038QosLimitsMaxJobsActiveJobs struct {
	Per *Dbv0038QosLimitsMaxJobsActiveJobsPer `json:"per,omitempty"`
}

// NewDbv0038QosLimitsMaxJobsActiveJobs instantiates a new Dbv0038QosLimitsMaxJobsActiveJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosLimitsMaxJobsActiveJobs() *Dbv0038QosLimitsMaxJobsActiveJobs {
	this := Dbv0038QosLimitsMaxJobsActiveJobs{}
	return &this
}

// NewDbv0038QosLimitsMaxJobsActiveJobsWithDefaults instantiates a new Dbv0038QosLimitsMaxJobsActiveJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosLimitsMaxJobsActiveJobsWithDefaults() *Dbv0038QosLimitsMaxJobsActiveJobs {
	this := Dbv0038QosLimitsMaxJobsActiveJobs{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxJobsActiveJobs) GetPer() Dbv0038QosLimitsMaxJobsActiveJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret Dbv0038QosLimitsMaxJobsActiveJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxJobsActiveJobs) GetPerOk() (*Dbv0038QosLimitsMaxJobsActiveJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxJobsActiveJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given Dbv0038QosLimitsMaxJobsActiveJobsPer and assigns it to the Per field.
func (o *Dbv0038QosLimitsMaxJobsActiveJobs) SetPer(v Dbv0038QosLimitsMaxJobsActiveJobsPer) {
	o.Per = &v
}

func (o Dbv0038QosLimitsMaxJobsActiveJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosLimitsMaxJobsActiveJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableDbv0038QosLimitsMaxJobsActiveJobs struct {
	value *Dbv0038QosLimitsMaxJobsActiveJobs
	isSet bool
}

func (v NullableDbv0038QosLimitsMaxJobsActiveJobs) Get() *Dbv0038QosLimitsMaxJobsActiveJobs {
	return v.value
}

func (v *NullableDbv0038QosLimitsMaxJobsActiveJobs) Set(val *Dbv0038QosLimitsMaxJobsActiveJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosLimitsMaxJobsActiveJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosLimitsMaxJobsActiveJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosLimitsMaxJobsActiveJobs(val *Dbv0038QosLimitsMaxJobsActiveJobs) *NullableDbv0038QosLimitsMaxJobsActiveJobs {
	return &NullableDbv0038QosLimitsMaxJobsActiveJobs{value: val, isSet: true}
}

func (v NullableDbv0038QosLimitsMaxJobsActiveJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosLimitsMaxJobsActiveJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


