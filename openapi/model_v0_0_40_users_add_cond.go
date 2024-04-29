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

// checks if the V0040UsersAddCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040UsersAddCond{}

// V0040UsersAddCond struct for V0040UsersAddCond
type V0040UsersAddCond struct {
	Accounts []string `json:"accounts,omitempty"`
	Association *V0040AssocRecSet `json:"association,omitempty"`
	Clusters []string `json:"clusters,omitempty"`
	Partitions []string `json:"partitions,omitempty"`
	Users []string `json:"users"`
	Wckeys []string `json:"wckeys,omitempty"`
}

type _V0040UsersAddCond V0040UsersAddCond

// NewV0040UsersAddCond instantiates a new V0040UsersAddCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040UsersAddCond(users []string) *V0040UsersAddCond {
	this := V0040UsersAddCond{}
	this.Users = users
	return &this
}

// NewV0040UsersAddCondWithDefaults instantiates a new V0040UsersAddCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040UsersAddCondWithDefaults() *V0040UsersAddCond {
	this := V0040UsersAddCond{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *V0040UsersAddCond) GetAccounts() []string {
	if o == nil || IsNil(o.Accounts) {
		var ret []string
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040UsersAddCond) GetAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *V0040UsersAddCond) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []string and assigns it to the Accounts field.
func (o *V0040UsersAddCond) SetAccounts(v []string) {
	o.Accounts = v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *V0040UsersAddCond) GetAssociation() V0040AssocRecSet {
	if o == nil || IsNil(o.Association) {
		var ret V0040AssocRecSet
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040UsersAddCond) GetAssociationOk() (*V0040AssocRecSet, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *V0040UsersAddCond) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given V0040AssocRecSet and assigns it to the Association field.
func (o *V0040UsersAddCond) SetAssociation(v V0040AssocRecSet) {
	o.Association = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V0040UsersAddCond) GetClusters() []string {
	if o == nil || IsNil(o.Clusters) {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040UsersAddCond) GetClustersOk() ([]string, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V0040UsersAddCond) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *V0040UsersAddCond) SetClusters(v []string) {
	o.Clusters = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *V0040UsersAddCond) GetPartitions() []string {
	if o == nil || IsNil(o.Partitions) {
		var ret []string
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040UsersAddCond) GetPartitionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *V0040UsersAddCond) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []string and assigns it to the Partitions field.
func (o *V0040UsersAddCond) SetPartitions(v []string) {
	o.Partitions = v
}

// GetUsers returns the Users field value
func (o *V0040UsersAddCond) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *V0040UsersAddCond) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *V0040UsersAddCond) SetUsers(v []string) {
	o.Users = v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *V0040UsersAddCond) GetWckeys() []string {
	if o == nil || IsNil(o.Wckeys) {
		var ret []string
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040UsersAddCond) GetWckeysOk() ([]string, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *V0040UsersAddCond) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []string and assigns it to the Wckeys field.
func (o *V0040UsersAddCond) SetWckeys(v []string) {
	o.Wckeys = v
}

func (o V0040UsersAddCond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040UsersAddCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	toSerialize["users"] = o.Users
	if !IsNil(o.Wckeys) {
		toSerialize["wckeys"] = o.Wckeys
	}
	return toSerialize, nil
}

func (o *V0040UsersAddCond) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
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

	varV0040UsersAddCond := _V0040UsersAddCond{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040UsersAddCond)

	if err != nil {
		return err
	}

	*o = V0040UsersAddCond(varV0040UsersAddCond)

	return err
}

type NullableV0040UsersAddCond struct {
	value *V0040UsersAddCond
	isSet bool
}

func (v NullableV0040UsersAddCond) Get() *V0040UsersAddCond {
	return v.value
}

func (v *NullableV0040UsersAddCond) Set(val *V0040UsersAddCond) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040UsersAddCond) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040UsersAddCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040UsersAddCond(val *V0040UsersAddCond) *NullableV0040UsersAddCond {
	return &NullableV0040UsersAddCond{value: val, isSet: true}
}

func (v NullableV0040UsersAddCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040UsersAddCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


