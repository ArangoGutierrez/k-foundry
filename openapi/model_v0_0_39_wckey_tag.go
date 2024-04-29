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

// checks if the V0039WckeyTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039WckeyTag{}

// V0039WckeyTag wckey details
type V0039WckeyTag struct {
	// wckey
	Wckey *string `json:"wckey,omitempty"`
	// active flags
	Flags []string `json:"flags,omitempty"`
}

// NewV0039WckeyTag instantiates a new V0039WckeyTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039WckeyTag() *V0039WckeyTag {
	this := V0039WckeyTag{}
	return &this
}

// NewV0039WckeyTagWithDefaults instantiates a new V0039WckeyTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039WckeyTagWithDefaults() *V0039WckeyTag {
	this := V0039WckeyTag{}
	return &this
}

// GetWckey returns the Wckey field value if set, zero value otherwise.
func (o *V0039WckeyTag) GetWckey() string {
	if o == nil || IsNil(o.Wckey) {
		var ret string
		return ret
	}
	return *o.Wckey
}

// GetWckeyOk returns a tuple with the Wckey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039WckeyTag) GetWckeyOk() (*string, bool) {
	if o == nil || IsNil(o.Wckey) {
		return nil, false
	}
	return o.Wckey, true
}

// HasWckey returns a boolean if a field has been set.
func (o *V0039WckeyTag) HasWckey() bool {
	if o != nil && !IsNil(o.Wckey) {
		return true
	}

	return false
}

// SetWckey gets a reference to the given string and assigns it to the Wckey field.
func (o *V0039WckeyTag) SetWckey(v string) {
	o.Wckey = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0039WckeyTag) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039WckeyTag) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0039WckeyTag) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0039WckeyTag) SetFlags(v []string) {
	o.Flags = v
}

func (o V0039WckeyTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039WckeyTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Wckey) {
		toSerialize["wckey"] = o.Wckey
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableV0039WckeyTag struct {
	value *V0039WckeyTag
	isSet bool
}

func (v NullableV0039WckeyTag) Get() *V0039WckeyTag {
	return v.value
}

func (v *NullableV0039WckeyTag) Set(val *V0039WckeyTag) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039WckeyTag) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039WckeyTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039WckeyTag(val *V0039WckeyTag) *NullableV0039WckeyTag {
	return &NullableV0039WckeyTag{value: val, isSet: true}
}

func (v NullableV0039WckeyTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039WckeyTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


