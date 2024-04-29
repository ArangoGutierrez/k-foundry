# Dbv0038JobStepTres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | Pointer to [**Dbv0038JobStepTresRequested**](Dbv0038JobStepTresRequested.md) |  | [optional] 
**Consumed** | Pointer to [**Dbv0038JobStepTresRequested**](Dbv0038JobStepTresRequested.md) |  | [optional] 
**Allocated** | Pointer to [**[]Dbv0038TresListInner**](Dbv0038TresListInner.md) | TRES list of attributes | [optional] 

## Methods

### NewDbv0038JobStepTres

`func NewDbv0038JobStepTres() *Dbv0038JobStepTres`

NewDbv0038JobStepTres instantiates a new Dbv0038JobStepTres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobStepTresWithDefaults

`func NewDbv0038JobStepTresWithDefaults() *Dbv0038JobStepTres`

NewDbv0038JobStepTresWithDefaults instantiates a new Dbv0038JobStepTres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequested

`func (o *Dbv0038JobStepTres) GetRequested() Dbv0038JobStepTresRequested`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *Dbv0038JobStepTres) GetRequestedOk() (*Dbv0038JobStepTresRequested, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *Dbv0038JobStepTres) SetRequested(v Dbv0038JobStepTresRequested)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *Dbv0038JobStepTres) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetConsumed

`func (o *Dbv0038JobStepTres) GetConsumed() Dbv0038JobStepTresRequested`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *Dbv0038JobStepTres) GetConsumedOk() (*Dbv0038JobStepTresRequested, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *Dbv0038JobStepTres) SetConsumed(v Dbv0038JobStepTresRequested)`

SetConsumed sets Consumed field to given value.

### HasConsumed

`func (o *Dbv0038JobStepTres) HasConsumed() bool`

HasConsumed returns a boolean if a field has been set.

### GetAllocated

`func (o *Dbv0038JobStepTres) GetAllocated() []Dbv0038TresListInner`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *Dbv0038JobStepTres) GetAllocatedOk() (*[]Dbv0038TresListInner, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *Dbv0038JobStepTres) SetAllocated(v []Dbv0038TresListInner)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *Dbv0038JobStepTres) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


