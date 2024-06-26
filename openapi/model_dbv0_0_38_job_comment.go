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

// checks if the Dbv0038JobComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobComment{}

// Dbv0038JobComment Job comments by type
type Dbv0038JobComment struct {
	// Administrator set comment
	Administrator *string `json:"administrator,omitempty"`
	// Job comment
	Job *string `json:"job,omitempty"`
	// System set comment
	System *string `json:"system,omitempty"`
}

// NewDbv0038JobComment instantiates a new Dbv0038JobComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobComment() *Dbv0038JobComment {
	this := Dbv0038JobComment{}
	return &this
}

// NewDbv0038JobCommentWithDefaults instantiates a new Dbv0038JobComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobCommentWithDefaults() *Dbv0038JobComment {
	this := Dbv0038JobComment{}
	return &this
}

// GetAdministrator returns the Administrator field value if set, zero value otherwise.
func (o *Dbv0038JobComment) GetAdministrator() string {
	if o == nil || IsNil(o.Administrator) {
		var ret string
		return ret
	}
	return *o.Administrator
}

// GetAdministratorOk returns a tuple with the Administrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobComment) GetAdministratorOk() (*string, bool) {
	if o == nil || IsNil(o.Administrator) {
		return nil, false
	}
	return o.Administrator, true
}

// HasAdministrator returns a boolean if a field has been set.
func (o *Dbv0038JobComment) HasAdministrator() bool {
	if o != nil && !IsNil(o.Administrator) {
		return true
	}

	return false
}

// SetAdministrator gets a reference to the given string and assigns it to the Administrator field.
func (o *Dbv0038JobComment) SetAdministrator(v string) {
	o.Administrator = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *Dbv0038JobComment) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobComment) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *Dbv0038JobComment) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *Dbv0038JobComment) SetJob(v string) {
	o.Job = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *Dbv0038JobComment) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobComment) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *Dbv0038JobComment) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *Dbv0038JobComment) SetSystem(v string) {
	o.System = &v
}

func (o Dbv0038JobComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Administrator) {
		toSerialize["administrator"] = o.Administrator
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	return toSerialize, nil
}

type NullableDbv0038JobComment struct {
	value *Dbv0038JobComment
	isSet bool
}

func (v NullableDbv0038JobComment) Get() *Dbv0038JobComment {
	return v.value
}

func (v *NullableDbv0038JobComment) Set(val *Dbv0038JobComment) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobComment) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobComment(val *Dbv0038JobComment) *NullableDbv0038JobComment {
	return &NullableDbv0038JobComment{value: val, isSet: true}
}

func (v NullableDbv0038JobComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


