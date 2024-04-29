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

// checks if the V0040OpenapiJobInfoResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiJobInfoResp{}

// V0040OpenapiJobInfoResp struct for V0040OpenapiJobInfoResp
type V0040OpenapiJobInfoResp struct {
	Jobs []V0040JobInfo `json:"jobs"`
	LastBackfill V0040Uint64NoVal `json:"last_backfill"`
	LastUpdate V0040Uint64NoVal `json:"last_update"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiJobInfoResp V0040OpenapiJobInfoResp

// NewV0040OpenapiJobInfoResp instantiates a new V0040OpenapiJobInfoResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiJobInfoResp(jobs []V0040JobInfo, lastBackfill V0040Uint64NoVal, lastUpdate V0040Uint64NoVal) *V0040OpenapiJobInfoResp {
	this := V0040OpenapiJobInfoResp{}
	this.Jobs = jobs
	this.LastBackfill = lastBackfill
	this.LastUpdate = lastUpdate
	return &this
}

// NewV0040OpenapiJobInfoRespWithDefaults instantiates a new V0040OpenapiJobInfoResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiJobInfoRespWithDefaults() *V0040OpenapiJobInfoResp {
	this := V0040OpenapiJobInfoResp{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *V0040OpenapiJobInfoResp) GetJobs() []V0040JobInfo {
	if o == nil {
		var ret []V0040JobInfo
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobInfoResp) GetJobsOk() ([]V0040JobInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *V0040OpenapiJobInfoResp) SetJobs(v []V0040JobInfo) {
	o.Jobs = v
}

// GetLastBackfill returns the LastBackfill field value
func (o *V0040OpenapiJobInfoResp) GetLastBackfill() V0040Uint64NoVal {
	if o == nil {
		var ret V0040Uint64NoVal
		return ret
	}

	return o.LastBackfill
}

// GetLastBackfillOk returns a tuple with the LastBackfill field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobInfoResp) GetLastBackfillOk() (*V0040Uint64NoVal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastBackfill, true
}

// SetLastBackfill sets field value
func (o *V0040OpenapiJobInfoResp) SetLastBackfill(v V0040Uint64NoVal) {
	o.LastBackfill = v
}

// GetLastUpdate returns the LastUpdate field value
func (o *V0040OpenapiJobInfoResp) GetLastUpdate() V0040Uint64NoVal {
	if o == nil {
		var ret V0040Uint64NoVal
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobInfoResp) GetLastUpdateOk() (*V0040Uint64NoVal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *V0040OpenapiJobInfoResp) SetLastUpdate(v V0040Uint64NoVal) {
	o.LastUpdate = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiJobInfoResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobInfoResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiJobInfoResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiJobInfoResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiJobInfoResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobInfoResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiJobInfoResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiJobInfoResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiJobInfoResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiJobInfoResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiJobInfoResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiJobInfoResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiJobInfoResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiJobInfoResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobs"] = o.Jobs
	toSerialize["last_backfill"] = o.LastBackfill
	toSerialize["last_update"] = o.LastUpdate
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

func (o *V0040OpenapiJobInfoResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobs",
		"last_backfill",
		"last_update",
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

	varV0040OpenapiJobInfoResp := _V0040OpenapiJobInfoResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiJobInfoResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiJobInfoResp(varV0040OpenapiJobInfoResp)

	return err
}

type NullableV0040OpenapiJobInfoResp struct {
	value *V0040OpenapiJobInfoResp
	isSet bool
}

func (v NullableV0040OpenapiJobInfoResp) Get() *V0040OpenapiJobInfoResp {
	return v.value
}

func (v *NullableV0040OpenapiJobInfoResp) Set(val *V0040OpenapiJobInfoResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiJobInfoResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiJobInfoResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiJobInfoResp(val *V0040OpenapiJobInfoResp) *NullableV0040OpenapiJobInfoResp {
	return &NullableV0040OpenapiJobInfoResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiJobInfoResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiJobInfoResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


