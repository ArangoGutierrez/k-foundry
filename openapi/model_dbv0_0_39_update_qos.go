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

// checks if the Dbv0039UpdateQos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039UpdateQos{}

// Dbv0039UpdateQos struct for Dbv0039UpdateQos
type Dbv0039UpdateQos struct {
	Qos []V0039Qos `json:"qos,omitempty"`
}

// NewDbv0039UpdateQos instantiates a new Dbv0039UpdateQos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039UpdateQos() *Dbv0039UpdateQos {
	this := Dbv0039UpdateQos{}
	return &this
}

// NewDbv0039UpdateQosWithDefaults instantiates a new Dbv0039UpdateQos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039UpdateQosWithDefaults() *Dbv0039UpdateQos {
	this := Dbv0039UpdateQos{}
	return &this
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0039UpdateQos) GetQos() []V0039Qos {
	if o == nil || IsNil(o.Qos) {
		var ret []V0039Qos
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039UpdateQos) GetQosOk() ([]V0039Qos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0039UpdateQos) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []V0039Qos and assigns it to the Qos field.
func (o *Dbv0039UpdateQos) SetQos(v []V0039Qos) {
	o.Qos = v
}

func (o Dbv0039UpdateQos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039UpdateQos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	return toSerialize, nil
}

type NullableDbv0039UpdateQos struct {
	value *Dbv0039UpdateQos
	isSet bool
}

func (v NullableDbv0039UpdateQos) Get() *Dbv0039UpdateQos {
	return v.value
}

func (v *NullableDbv0039UpdateQos) Set(val *Dbv0039UpdateQos) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039UpdateQos) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039UpdateQos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039UpdateQos(val *Dbv0039UpdateQos) *NullableDbv0039UpdateQos {
	return &NullableDbv0039UpdateQos{value: val, isSet: true}
}

func (v NullableDbv0039UpdateQos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039UpdateQos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


