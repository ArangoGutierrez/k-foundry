# Dbv0038Qos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | QOS description | [optional] 
**Flags** | Pointer to **[]string** | List of properties of QOS | [optional] 
**Id** | Pointer to **string** | Database id | [optional] 
**Limits** | Pointer to [**Dbv0038QosLimits**](Dbv0038QosLimits.md) |  | [optional] 
**Preempt** | Pointer to [**Dbv0038QosPreempt**](Dbv0038QosPreempt.md) |  | [optional] 
**Priority** | Pointer to **int32** | QOS priority | [optional] 
**UsageFactor** | Pointer to **float32** | Usage factor | [optional] 
**UsageThreshold** | Pointer to **float32** | Usage threshold | [optional] 
**Name** | Pointer to **string** | Assigned name of QOS | [optional] 

## Methods

### NewDbv0038Qos

`func NewDbv0038Qos() *Dbv0038Qos`

NewDbv0038Qos instantiates a new Dbv0038Qos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038QosWithDefaults

`func NewDbv0038QosWithDefaults() *Dbv0038Qos`

NewDbv0038QosWithDefaults instantiates a new Dbv0038Qos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Dbv0038Qos) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dbv0038Qos) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dbv0038Qos) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dbv0038Qos) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlags

`func (o *Dbv0038Qos) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbv0038Qos) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Dbv0038Qos) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbv0038Qos) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetId

`func (o *Dbv0038Qos) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dbv0038Qos) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dbv0038Qos) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dbv0038Qos) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimits

`func (o *Dbv0038Qos) GetLimits() Dbv0038QosLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Dbv0038Qos) GetLimitsOk() (*Dbv0038QosLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Dbv0038Qos) SetLimits(v Dbv0038QosLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *Dbv0038Qos) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetPreempt

`func (o *Dbv0038Qos) GetPreempt() Dbv0038QosPreempt`

GetPreempt returns the Preempt field if non-nil, zero value otherwise.

### GetPreemptOk

`func (o *Dbv0038Qos) GetPreemptOk() (*Dbv0038QosPreempt, bool)`

GetPreemptOk returns a tuple with the Preempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreempt

`func (o *Dbv0038Qos) SetPreempt(v Dbv0038QosPreempt)`

SetPreempt sets Preempt field to given value.

### HasPreempt

`func (o *Dbv0038Qos) HasPreempt() bool`

HasPreempt returns a boolean if a field has been set.

### GetPriority

`func (o *Dbv0038Qos) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Dbv0038Qos) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Dbv0038Qos) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Dbv0038Qos) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUsageFactor

`func (o *Dbv0038Qos) GetUsageFactor() float32`

GetUsageFactor returns the UsageFactor field if non-nil, zero value otherwise.

### GetUsageFactorOk

`func (o *Dbv0038Qos) GetUsageFactorOk() (*float32, bool)`

GetUsageFactorOk returns a tuple with the UsageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageFactor

`func (o *Dbv0038Qos) SetUsageFactor(v float32)`

SetUsageFactor sets UsageFactor field to given value.

### HasUsageFactor

`func (o *Dbv0038Qos) HasUsageFactor() bool`

HasUsageFactor returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *Dbv0038Qos) GetUsageThreshold() float32`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *Dbv0038Qos) GetUsageThresholdOk() (*float32, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *Dbv0038Qos) SetUsageThreshold(v float32)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *Dbv0038Qos) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### GetName

`func (o *Dbv0038Qos) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0038Qos) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dbv0038Qos) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0038Qos) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


