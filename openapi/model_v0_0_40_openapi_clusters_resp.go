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

// checks if the V0040OpenapiClustersResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiClustersResp{}

// V0040OpenapiClustersResp struct for V0040OpenapiClustersResp
type V0040OpenapiClustersResp struct {
	Clusters []V0040ClusterRec `json:"clusters"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiClustersResp V0040OpenapiClustersResp

// NewV0040OpenapiClustersResp instantiates a new V0040OpenapiClustersResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiClustersResp(clusters []V0040ClusterRec) *V0040OpenapiClustersResp {
	this := V0040OpenapiClustersResp{}
	this.Clusters = clusters
	return &this
}

// NewV0040OpenapiClustersRespWithDefaults instantiates a new V0040OpenapiClustersResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiClustersRespWithDefaults() *V0040OpenapiClustersResp {
	this := V0040OpenapiClustersResp{}
	return &this
}

// GetClusters returns the Clusters field value
func (o *V0040OpenapiClustersResp) GetClusters() []V0040ClusterRec {
	if o == nil {
		var ret []V0040ClusterRec
		return ret
	}

	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiClustersResp) GetClustersOk() ([]V0040ClusterRec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clusters, true
}

// SetClusters sets field value
func (o *V0040OpenapiClustersResp) SetClusters(v []V0040ClusterRec) {
	o.Clusters = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiClustersResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiClustersResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiClustersResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiClustersResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiClustersResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiClustersResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiClustersResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiClustersResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiClustersResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiClustersResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiClustersResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiClustersResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiClustersResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiClustersResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusters"] = o.Clusters
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

func (o *V0040OpenapiClustersResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusters",
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

	varV0040OpenapiClustersResp := _V0040OpenapiClustersResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiClustersResp)

	if err != nil {
		return err
	}

	*o = V0040OpenapiClustersResp(varV0040OpenapiClustersResp)

	return err
}

type NullableV0040OpenapiClustersResp struct {
	value *V0040OpenapiClustersResp
	isSet bool
}

func (v NullableV0040OpenapiClustersResp) Get() *V0040OpenapiClustersResp {
	return v.value
}

func (v *NullableV0040OpenapiClustersResp) Set(val *V0040OpenapiClustersResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiClustersResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiClustersResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiClustersResp(val *V0040OpenapiClustersResp) *NullableV0040OpenapiClustersResp {
	return &NullableV0040OpenapiClustersResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiClustersResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiClustersResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


