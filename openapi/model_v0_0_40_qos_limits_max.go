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

// checks if the V0040QosLimitsMax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040QosLimitsMax{}

// V0040QosLimitsMax struct for V0040QosLimitsMax
type V0040QosLimitsMax struct {
	ActiveJobs *V0040QosLimitsMaxActiveJobs `json:"active_jobs,omitempty"`
	Tres *V0040QosLimitsMaxTres `json:"tres,omitempty"`
	WallClock *V0040QosLimitsMaxWallClock `json:"wall_clock,omitempty"`
	Jobs *V0040QosLimitsMaxJobs `json:"jobs,omitempty"`
	Accruing *V0040QosLimitsMaxJobsActiveJobs `json:"accruing,omitempty"`
}

// NewV0040QosLimitsMax instantiates a new V0040QosLimitsMax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040QosLimitsMax() *V0040QosLimitsMax {
	this := V0040QosLimitsMax{}
	return &this
}

// NewV0040QosLimitsMaxWithDefaults instantiates a new V0040QosLimitsMax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040QosLimitsMaxWithDefaults() *V0040QosLimitsMax {
	this := V0040QosLimitsMax{}
	return &this
}

// GetActiveJobs returns the ActiveJobs field value if set, zero value otherwise.
func (o *V0040QosLimitsMax) GetActiveJobs() V0040QosLimitsMaxActiveJobs {
	if o == nil || IsNil(o.ActiveJobs) {
		var ret V0040QosLimitsMaxActiveJobs
		return ret
	}
	return *o.ActiveJobs
}

// GetActiveJobsOk returns a tuple with the ActiveJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMax) GetActiveJobsOk() (*V0040QosLimitsMaxActiveJobs, bool) {
	if o == nil || IsNil(o.ActiveJobs) {
		return nil, false
	}
	return o.ActiveJobs, true
}

// HasActiveJobs returns a boolean if a field has been set.
func (o *V0040QosLimitsMax) HasActiveJobs() bool {
	if o != nil && !IsNil(o.ActiveJobs) {
		return true
	}

	return false
}

// SetActiveJobs gets a reference to the given V0040QosLimitsMaxActiveJobs and assigns it to the ActiveJobs field.
func (o *V0040QosLimitsMax) SetActiveJobs(v V0040QosLimitsMaxActiveJobs) {
	o.ActiveJobs = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0040QosLimitsMax) GetTres() V0040QosLimitsMaxTres {
	if o == nil || IsNil(o.Tres) {
		var ret V0040QosLimitsMaxTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMax) GetTresOk() (*V0040QosLimitsMaxTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0040QosLimitsMax) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given V0040QosLimitsMaxTres and assigns it to the Tres field.
func (o *V0040QosLimitsMax) SetTres(v V0040QosLimitsMaxTres) {
	o.Tres = &v
}

// GetWallClock returns the WallClock field value if set, zero value otherwise.
func (o *V0040QosLimitsMax) GetWallClock() V0040QosLimitsMaxWallClock {
	if o == nil || IsNil(o.WallClock) {
		var ret V0040QosLimitsMaxWallClock
		return ret
	}
	return *o.WallClock
}

// GetWallClockOk returns a tuple with the WallClock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMax) GetWallClockOk() (*V0040QosLimitsMaxWallClock, bool) {
	if o == nil || IsNil(o.WallClock) {
		return nil, false
	}
	return o.WallClock, true
}

// HasWallClock returns a boolean if a field has been set.
func (o *V0040QosLimitsMax) HasWallClock() bool {
	if o != nil && !IsNil(o.WallClock) {
		return true
	}

	return false
}

// SetWallClock gets a reference to the given V0040QosLimitsMaxWallClock and assigns it to the WallClock field.
func (o *V0040QosLimitsMax) SetWallClock(v V0040QosLimitsMaxWallClock) {
	o.WallClock = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0040QosLimitsMax) GetJobs() V0040QosLimitsMaxJobs {
	if o == nil || IsNil(o.Jobs) {
		var ret V0040QosLimitsMaxJobs
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMax) GetJobsOk() (*V0040QosLimitsMaxJobs, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0040QosLimitsMax) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given V0040QosLimitsMaxJobs and assigns it to the Jobs field.
func (o *V0040QosLimitsMax) SetJobs(v V0040QosLimitsMaxJobs) {
	o.Jobs = &v
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *V0040QosLimitsMax) GetAccruing() V0040QosLimitsMaxJobsActiveJobs {
	if o == nil || IsNil(o.Accruing) {
		var ret V0040QosLimitsMaxJobsActiveJobs
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMax) GetAccruingOk() (*V0040QosLimitsMaxJobsActiveJobs, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *V0040QosLimitsMax) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given V0040QosLimitsMaxJobsActiveJobs and assigns it to the Accruing field.
func (o *V0040QosLimitsMax) SetAccruing(v V0040QosLimitsMaxJobsActiveJobs) {
	o.Accruing = &v
}

func (o V0040QosLimitsMax) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040QosLimitsMax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveJobs) {
		toSerialize["active_jobs"] = o.ActiveJobs
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.WallClock) {
		toSerialize["wall_clock"] = o.WallClock
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	return toSerialize, nil
}

type NullableV0040QosLimitsMax struct {
	value *V0040QosLimitsMax
	isSet bool
}

func (v NullableV0040QosLimitsMax) Get() *V0040QosLimitsMax {
	return v.value
}

func (v *NullableV0040QosLimitsMax) Set(val *V0040QosLimitsMax) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040QosLimitsMax) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040QosLimitsMax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040QosLimitsMax(val *V0040QosLimitsMax) *NullableV0040QosLimitsMax {
	return &NullableV0040QosLimitsMax{value: val, isSet: true}
}

func (v NullableV0040QosLimitsMax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040QosLimitsMax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


