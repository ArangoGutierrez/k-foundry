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

// checks if the Dbv0038ClustersProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038ClustersProperties{}

// Dbv0038ClustersProperties struct for Dbv0038ClustersProperties
type Dbv0038ClustersProperties struct {
	Clusters *Dbv0038ClusterInfo `json:"clusters,omitempty"`
}

// NewDbv0038ClustersProperties instantiates a new Dbv0038ClustersProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038ClustersProperties() *Dbv0038ClustersProperties {
	this := Dbv0038ClustersProperties{}
	return &this
}

// NewDbv0038ClustersPropertiesWithDefaults instantiates a new Dbv0038ClustersProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038ClustersPropertiesWithDefaults() *Dbv0038ClustersProperties {
	this := Dbv0038ClustersProperties{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *Dbv0038ClustersProperties) GetClusters() Dbv0038ClusterInfo {
	if o == nil || IsNil(o.Clusters) {
		var ret Dbv0038ClusterInfo
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ClustersProperties) GetClustersOk() (*Dbv0038ClusterInfo, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *Dbv0038ClustersProperties) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given Dbv0038ClusterInfo and assigns it to the Clusters field.
func (o *Dbv0038ClustersProperties) SetClusters(v Dbv0038ClusterInfo) {
	o.Clusters = &v
}

func (o Dbv0038ClustersProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038ClustersProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	return toSerialize, nil
}

type NullableDbv0038ClustersProperties struct {
	value *Dbv0038ClustersProperties
	isSet bool
}

func (v NullableDbv0038ClustersProperties) Get() *Dbv0038ClustersProperties {
	return v.value
}

func (v *NullableDbv0038ClustersProperties) Set(val *Dbv0038ClustersProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038ClustersProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038ClustersProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038ClustersProperties(val *Dbv0038ClustersProperties) *NullableDbv0038ClustersProperties {
	return &NullableDbv0038ClustersProperties{value: val, isSet: true}
}

func (v NullableDbv0038ClustersProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038ClustersProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


