# V0040StepStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CPU** | Pointer to [**V0039StepStatisticsCPU**](V0039StepStatisticsCPU.md) |  | [optional] 
**Energy** | Pointer to [**V0040StepStatisticsEnergy**](V0040StepStatisticsEnergy.md) |  | [optional] 

## Methods

### NewV0040StepStatistics

`func NewV0040StepStatistics() *V0040StepStatistics`

NewV0040StepStatistics instantiates a new V0040StepStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040StepStatisticsWithDefaults

`func NewV0040StepStatisticsWithDefaults() *V0040StepStatistics`

NewV0040StepStatisticsWithDefaults instantiates a new V0040StepStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCPU

`func (o *V0040StepStatistics) GetCPU() V0039StepStatisticsCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *V0040StepStatistics) GetCPUOk() (*V0039StepStatisticsCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *V0040StepStatistics) SetCPU(v V0039StepStatisticsCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *V0040StepStatistics) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetEnergy

`func (o *V0040StepStatistics) GetEnergy() V0040StepStatisticsEnergy`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *V0040StepStatistics) GetEnergyOk() (*V0040StepStatisticsEnergy, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *V0040StepStatistics) SetEnergy(v V0040StepStatisticsEnergy)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *V0040StepStatistics) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


