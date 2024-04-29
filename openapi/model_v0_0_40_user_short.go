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

// checks if the V0040UserShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040UserShort{}

// V0040UserShort struct for V0040UserShort
type V0040UserShort struct {
	// Admin level of user.  Valid levels are None, Operator, and Admin.
	Adminlevel []string `json:"adminlevel,omitempty"`
	// Identify the default bank account name to be used for a job if none is specified at submission time.
	Defaultaccount *string `json:"defaultaccount,omitempty"`
	// Identify the default Workload Characterization Key.
	Defaultwckey *string `json:"defaultwckey,omitempty"`
}

// NewV0040UserShort instantiates a new V0040UserShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040UserShort() *V0040UserShort {
	this := V0040UserShort{}
	return &this
}

// NewV0040UserShortWithDefaults instantiates a new V0040UserShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040UserShortWithDefaults() *V0040UserShort {
	this := V0040UserShort{}
	return &this
}

// GetAdminlevel returns the Adminlevel field value if set, zero value otherwise.
func (o *V0040UserShort) GetAdminlevel() []string {
	if o == nil || IsNil(o.Adminlevel) {
		var ret []string
		return ret
	}
	return o.Adminlevel
}

// GetAdminlevelOk returns a tuple with the Adminlevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040UserShort) GetAdminlevelOk() ([]string, bool) {
	if o == nil || IsNil(o.Adminlevel) {
		return nil, false
	}
	return o.Adminlevel, true
}

// HasAdminlevel returns a boolean if a field has been set.
func (o *V0040UserShort) HasAdminlevel() bool {
	if o != nil && !IsNil(o.Adminlevel) {
		return true
	}

	return false
}

// SetAdminlevel gets a reference to the given []string and assigns it to the Adminlevel field.
func (o *V0040UserShort) SetAdminlevel(v []string) {
	o.Adminlevel = v
}

// GetDefaultaccount returns the Defaultaccount field value if set, zero value otherwise.
func (o *V0040UserShort) GetDefaultaccount() string {
	if o == nil || IsNil(o.Defaultaccount) {
		var ret string
		return ret
	}
	return *o.Defaultaccount
}

// GetDefaultaccountOk returns a tuple with the Defaultaccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040UserShort) GetDefaultaccountOk() (*string, bool) {
	if o == nil || IsNil(o.Defaultaccount) {
		return nil, false
	}
	return o.Defaultaccount, true
}

// HasDefaultaccount returns a boolean if a field has been set.
func (o *V0040UserShort) HasDefaultaccount() bool {
	if o != nil && !IsNil(o.Defaultaccount) {
		return true
	}

	return false
}

// SetDefaultaccount gets a reference to the given string and assigns it to the Defaultaccount field.
func (o *V0040UserShort) SetDefaultaccount(v string) {
	o.Defaultaccount = &v
}

// GetDefaultwckey returns the Defaultwckey field value if set, zero value otherwise.
func (o *V0040UserShort) GetDefaultwckey() string {
	if o == nil || IsNil(o.Defaultwckey) {
		var ret string
		return ret
	}
	return *o.Defaultwckey
}

// GetDefaultwckeyOk returns a tuple with the Defaultwckey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040UserShort) GetDefaultwckeyOk() (*string, bool) {
	if o == nil || IsNil(o.Defaultwckey) {
		return nil, false
	}
	return o.Defaultwckey, true
}

// HasDefaultwckey returns a boolean if a field has been set.
func (o *V0040UserShort) HasDefaultwckey() bool {
	if o != nil && !IsNil(o.Defaultwckey) {
		return true
	}

	return false
}

// SetDefaultwckey gets a reference to the given string and assigns it to the Defaultwckey field.
func (o *V0040UserShort) SetDefaultwckey(v string) {
	o.Defaultwckey = &v
}

func (o V0040UserShort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040UserShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adminlevel) {
		toSerialize["adminlevel"] = o.Adminlevel
	}
	if !IsNil(o.Defaultaccount) {
		toSerialize["defaultaccount"] = o.Defaultaccount
	}
	if !IsNil(o.Defaultwckey) {
		toSerialize["defaultwckey"] = o.Defaultwckey
	}
	return toSerialize, nil
}

type NullableV0040UserShort struct {
	value *V0040UserShort
	isSet bool
}

func (v NullableV0040UserShort) Get() *V0040UserShort {
	return v.value
}

func (v *NullableV0040UserShort) Set(val *V0040UserShort) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040UserShort) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040UserShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040UserShort(val *V0040UserShort) *NullableV0040UserShort {
	return &NullableV0040UserShort{value: val, isSet: true}
}

func (v NullableV0040UserShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040UserShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


