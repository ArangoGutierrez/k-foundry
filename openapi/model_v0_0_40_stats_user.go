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

// checks if the V0040StatsUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StatsUser{}

// V0040StatsUser struct for V0040StatsUser
type V0040StatsUser struct {
	User *string `json:"user,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Time *V0039StatsRpcTime `json:"time,omitempty"`
}

// NewV0040StatsUser instantiates a new V0040StatsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StatsUser() *V0040StatsUser {
	this := V0040StatsUser{}
	return &this
}

// NewV0040StatsUserWithDefaults instantiates a new V0040StatsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StatsUserWithDefaults() *V0040StatsUser {
	this := V0040StatsUser{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0040StatsUser) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsUser) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0040StatsUser) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *V0040StatsUser) SetUser(v string) {
	o.User = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0040StatsUser) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsUser) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0040StatsUser) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0040StatsUser) SetCount(v int32) {
	o.Count = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0040StatsUser) GetTime() V0039StatsRpcTime {
	if o == nil || IsNil(o.Time) {
		var ret V0039StatsRpcTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsUser) GetTimeOk() (*V0039StatsRpcTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0040StatsUser) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0039StatsRpcTime and assigns it to the Time field.
func (o *V0040StatsUser) SetTime(v V0039StatsRpcTime) {
	o.Time = &v
}

func (o V0040StatsUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StatsUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableV0040StatsUser struct {
	value *V0040StatsUser
	isSet bool
}

func (v NullableV0040StatsUser) Get() *V0040StatsUser {
	return v.value
}

func (v *NullableV0040StatsUser) Set(val *V0040StatsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StatsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StatsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StatsUser(val *V0040StatsUser) *NullableV0040StatsUser {
	return &NullableV0040StatsUser{value: val, isSet: true}
}

func (v NullableV0040StatsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StatsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


