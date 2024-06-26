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

// checks if the Dbv0038ResponseAssociationsDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038ResponseAssociationsDelete{}

// Dbv0038ResponseAssociationsDelete struct for Dbv0038ResponseAssociationsDelete
type Dbv0038ResponseAssociationsDelete struct {
	Meta *Dbv0038Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0038Error `json:"errors,omitempty"`
	// the associations
	RemovedAssociations []string `json:"removed_associations,omitempty"`
}

// NewDbv0038ResponseAssociationsDelete instantiates a new Dbv0038ResponseAssociationsDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038ResponseAssociationsDelete() *Dbv0038ResponseAssociationsDelete {
	this := Dbv0038ResponseAssociationsDelete{}
	return &this
}

// NewDbv0038ResponseAssociationsDeleteWithDefaults instantiates a new Dbv0038ResponseAssociationsDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038ResponseAssociationsDeleteWithDefaults() *Dbv0038ResponseAssociationsDelete {
	this := Dbv0038ResponseAssociationsDelete{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0038ResponseAssociationsDelete) GetMeta() Dbv0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseAssociationsDelete) GetMetaOk() (*Dbv0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0038ResponseAssociationsDelete) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0038Meta and assigns it to the Meta field.
func (o *Dbv0038ResponseAssociationsDelete) SetMeta(v Dbv0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0038ResponseAssociationsDelete) GetErrors() []Dbv0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseAssociationsDelete) GetErrorsOk() ([]Dbv0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0038ResponseAssociationsDelete) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0038Error and assigns it to the Errors field.
func (o *Dbv0038ResponseAssociationsDelete) SetErrors(v []Dbv0038Error) {
	o.Errors = v
}

// GetRemovedAssociations returns the RemovedAssociations field value if set, zero value otherwise.
func (o *Dbv0038ResponseAssociationsDelete) GetRemovedAssociations() []string {
	if o == nil || IsNil(o.RemovedAssociations) {
		var ret []string
		return ret
	}
	return o.RemovedAssociations
}

// GetRemovedAssociationsOk returns a tuple with the RemovedAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ResponseAssociationsDelete) GetRemovedAssociationsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemovedAssociations) {
		return nil, false
	}
	return o.RemovedAssociations, true
}

// HasRemovedAssociations returns a boolean if a field has been set.
func (o *Dbv0038ResponseAssociationsDelete) HasRemovedAssociations() bool {
	if o != nil && !IsNil(o.RemovedAssociations) {
		return true
	}

	return false
}

// SetRemovedAssociations gets a reference to the given []string and assigns it to the RemovedAssociations field.
func (o *Dbv0038ResponseAssociationsDelete) SetRemovedAssociations(v []string) {
	o.RemovedAssociations = v
}

func (o Dbv0038ResponseAssociationsDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038ResponseAssociationsDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.RemovedAssociations) {
		toSerialize["removed_associations"] = o.RemovedAssociations
	}
	return toSerialize, nil
}

type NullableDbv0038ResponseAssociationsDelete struct {
	value *Dbv0038ResponseAssociationsDelete
	isSet bool
}

func (v NullableDbv0038ResponseAssociationsDelete) Get() *Dbv0038ResponseAssociationsDelete {
	return v.value
}

func (v *NullableDbv0038ResponseAssociationsDelete) Set(val *Dbv0038ResponseAssociationsDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038ResponseAssociationsDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038ResponseAssociationsDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038ResponseAssociationsDelete(val *Dbv0038ResponseAssociationsDelete) *NullableDbv0038ResponseAssociationsDelete {
	return &NullableDbv0038ResponseAssociationsDelete{value: val, isSet: true}
}

func (v NullableDbv0038ResponseAssociationsDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038ResponseAssociationsDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


