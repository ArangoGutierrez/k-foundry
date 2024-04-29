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
	"bytes"
	"fmt"
)

// checks if the V0040OpenapiSlurmdbdJobsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiSlurmdbdJobsResp{}

// V0040OpenapiSlurmdbdJobsResp struct for V0040OpenapiSlurmdbdJobsResp
type V0040OpenapiSlurmdbdJobsResp struct {
	Jobs []V0040Job `json:"jobs"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiSlurmdbdJobsResp V0040OpenapiSlurmdbdJobsResp

// NewV0040OpenapiSlurmdbdJobsResp instantiates a new V0040OpenapiSlurmdbdJobsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiSlurmdbdJobsResp(jobs []V0040Job) *V0040OpenapiSlurmdbdJobsResp {
	this := V0040OpenapiSlurmdbdJobsResp{}
	this.Jobs = jobs
	return &this
}

// NewV0040OpenapiSlurmdbdJobsRespWithDefaults instantiates a new V0040OpenapiSlurmdbdJobsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiSlurmdbdJobsRespWithDefaults() *V0040OpenapiSlurmdbdJobsResp {
	this := V0040OpenapiSlurmdbdJobsResp{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *V0040OpenapiSlurmdbdJobsResp) GetJobs() []V0040Job {
	if o == nil {
		var ret []V0040Job
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdJobsResp) GetJobsOk() ([]V0040Job, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *V0040OpenapiSlurmdbdJobsResp) SetJobs(v []V0040Job) {
	o.Jobs = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdJobsResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdJobsResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdJobsResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiSlurmdbdJobsResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdJobsResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdJobsResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdJobsResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiSlurmdbdJobsResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdJobsResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdJobsResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdJobsResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiSlurmdbdJobsResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiSlurmdbdJobsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiSlurmdbdJobsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobs"] = o.Jobs
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *V0040OpenapiSlurmdbdJobsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0040OpenapiSlurmdbdJobsResp := _V0040OpenapiSlurmdbdJobsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiSlurmdbdJobsResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiSlurmdbdJobsResp(varV0040OpenapiSlurmdbdJobsResp)

	return err
}

type NullableV0040OpenapiSlurmdbdJobsResp struct {
	value *V0040OpenapiSlurmdbdJobsResp
	isSet bool
}

func (v NullableV0040OpenapiSlurmdbdJobsResp) Get() *V0040OpenapiSlurmdbdJobsResp {
	return v.value
}

func (v *NullableV0040OpenapiSlurmdbdJobsResp) Set(val *V0040OpenapiSlurmdbdJobsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiSlurmdbdJobsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiSlurmdbdJobsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiSlurmdbdJobsResp(val *V0040OpenapiSlurmdbdJobsResp) *NullableV0040OpenapiSlurmdbdJobsResp {
	return &NullableV0040OpenapiSlurmdbdJobsResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiSlurmdbdJobsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiSlurmdbdJobsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

