# Dbv0038JobStepCPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedFrequency** | Pointer to [**Dbv0038JobStepCPURequestedFrequency**](Dbv0038JobStepCPURequestedFrequency.md) |  | [optional] 
**Governor** | Pointer to **[]string** | CPU governor | [optional] 

## Methods

### NewDbv0038JobStepCPU

`func NewDbv0038JobStepCPU() *Dbv0038JobStepCPU`

NewDbv0038JobStepCPU instantiates a new Dbv0038JobStepCPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobStepCPUWithDefaults

`func NewDbv0038JobStepCPUWithDefaults() *Dbv0038JobStepCPU`

NewDbv0038JobStepCPUWithDefaults instantiates a new Dbv0038JobStepCPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedFrequency

`func (o *Dbv0038JobStepCPU) GetRequestedFrequency() Dbv0038JobStepCPURequestedFrequency`

GetRequestedFrequency returns the RequestedFrequency field if non-nil, zero value otherwise.

### GetRequestedFrequencyOk

`func (o *Dbv0038JobStepCPU) GetRequestedFrequencyOk() (*Dbv0038JobStepCPURequestedFrequency, bool)`

GetRequestedFrequencyOk returns a tuple with the RequestedFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFrequency

`func (o *Dbv0038JobStepCPU) SetRequestedFrequency(v Dbv0038JobStepCPURequestedFrequency)`

SetRequestedFrequency sets RequestedFrequency field to given value.

### HasRequestedFrequency

`func (o *Dbv0038JobStepCPU) HasRequestedFrequency() bool`

HasRequestedFrequency returns a boolean if a field has been set.

### GetGovernor

`func (o *Dbv0038JobStepCPU) GetGovernor() []string`

GetGovernor returns the Governor field if non-nil, zero value otherwise.

### GetGovernorOk

`func (o *Dbv0038JobStepCPU) GetGovernorOk() (*[]string, bool)`

GetGovernorOk returns a tuple with the Governor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernor

`func (o *Dbv0038JobStepCPU) SetGovernor(v []string)`

SetGovernor sets Governor field to given value.

### HasGovernor

`func (o *Dbv0038JobStepCPU) HasGovernor() bool`

HasGovernor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


