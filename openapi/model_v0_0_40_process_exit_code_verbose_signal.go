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

// checks if the V0040ProcessExitCodeVerboseSignal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040ProcessExitCodeVerboseSignal{}

// V0040ProcessExitCodeVerboseSignal struct for V0040ProcessExitCodeVerboseSignal
type V0040ProcessExitCodeVerboseSignal struct {
	Id *V0040Uint16NoVal `json:"id,omitempty"`
	// Signal sent to process
	Name *string `json:"name,omitempty"`
}

// NewV0040ProcessExitCodeVerboseSignal instantiates a new V0040ProcessExitCodeVerboseSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040ProcessExitCodeVerboseSignal() *V0040ProcessExitCodeVerboseSignal {
	this := V0040ProcessExitCodeVerboseSignal{}
	return &this
}

// NewV0040ProcessExitCodeVerboseSignalWithDefaults instantiates a new V0040ProcessExitCodeVerboseSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040ProcessExitCodeVerboseSignalWithDefaults() *V0040ProcessExitCodeVerboseSignal {
	this := V0040ProcessExitCodeVerboseSignal{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0040ProcessExitCodeVerboseSignal) GetId() V0040Uint16NoVal {
	if o == nil || IsNil(o.Id) {
		var ret V0040Uint16NoVal
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ProcessExitCodeVerboseSignal) GetIdOk() (*V0040Uint16NoVal, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0040ProcessExitCodeVerboseSignal) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given V0040Uint16NoVal and assigns it to the Id field.
func (o *V0040ProcessExitCodeVerboseSignal) SetId(v V0040Uint16NoVal) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0040ProcessExitCodeVerboseSignal) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040ProcessExitCodeVerboseSignal) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0040ProcessExitCodeVerboseSignal) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0040ProcessExitCodeVerboseSignal) SetName(v string) {
	o.Name = &v
}

func (o V0040ProcessExitCodeVerboseSignal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040ProcessExitCodeVerboseSignal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV0040ProcessExitCodeVerboseSignal struct {
	value *V0040ProcessExitCodeVerboseSignal
	isSet bool
}

func (v NullableV0040ProcessExitCodeVerboseSignal) Get() *V0040ProcessExitCodeVerboseSignal {
	return v.value
}

func (v *NullableV0040ProcessExitCodeVerboseSignal) Set(val *V0040ProcessExitCodeVerboseSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040ProcessExitCodeVerboseSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040ProcessExitCodeVerboseSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040ProcessExitCodeVerboseSignal(val *V0040ProcessExitCodeVerboseSignal) *NullableV0040ProcessExitCodeVerboseSignal {
	return &NullableV0040ProcessExitCodeVerboseSignal{value: val, isSet: true}
}

func (v NullableV0040ProcessExitCodeVerboseSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040ProcessExitCodeVerboseSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


