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

// checks if the Dbv0038QosInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosInfo{}

// Dbv0038QosInfo struct for Dbv0038QosInfo
type Dbv0038QosInfo struct {
	Meta *Dbv0038Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0038Error `json:"errors,omitempty"`
	// Array of QOS
	Qos []Dbv0038Qos `json:"qos,omitempty"`
}

// NewDbv0038QosInfo instantiates a new Dbv0038QosInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosInfo() *Dbv0038QosInfo {
	this := Dbv0038QosInfo{}
	return &this
}

// NewDbv0038QosInfoWithDefaults instantiates a new Dbv0038QosInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosInfoWithDefaults() *Dbv0038QosInfo {
	this := Dbv0038QosInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0038QosInfo) GetMeta() Dbv0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosInfo) GetMetaOk() (*Dbv0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0038QosInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0038Meta and assigns it to the Meta field.
func (o *Dbv0038QosInfo) SetMeta(v Dbv0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0038QosInfo) GetErrors() []Dbv0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosInfo) GetErrorsOk() ([]Dbv0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0038QosInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0038Error and assigns it to the Errors field.
func (o *Dbv0038QosInfo) SetErrors(v []Dbv0038Error) {
	o.Errors = v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0038QosInfo) GetQos() []Dbv0038Qos {
	if o == nil || IsNil(o.Qos) {
		var ret []Dbv0038Qos
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosInfo) GetQosOk() ([]Dbv0038Qos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0038QosInfo) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []Dbv0038Qos and assigns it to the Qos field.
func (o *Dbv0038QosInfo) SetQos(v []Dbv0038Qos) {
	o.Qos = v
}

func (o Dbv0038QosInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	return toSerialize, nil
}

type NullableDbv0038QosInfo struct {
	value *Dbv0038QosInfo
	isSet bool
}

func (v NullableDbv0038QosInfo) Get() *Dbv0038QosInfo {
	return v.value
}

func (v *NullableDbv0038QosInfo) Set(val *Dbv0038QosInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosInfo(val *Dbv0038QosInfo) *NullableDbv0038QosInfo {
	return &NullableDbv0038QosInfo{value: val, isSet: true}
}

func (v NullableDbv0038QosInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


