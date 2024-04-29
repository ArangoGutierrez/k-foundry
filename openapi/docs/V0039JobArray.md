# V0039JobArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** |  | [optional] 
**Limits** | Pointer to [**V0039JobArrayLimits**](V0039JobArrayLimits.md) |  | [optional] 
**TaskId** | Pointer to [**V0039Uint32NoVal**](V0039Uint32NoVal.md) |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 

## Methods

### NewV0039JobArray

`func NewV0039JobArray() *V0039JobArray`

NewV0039JobArray instantiates a new V0039JobArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039JobArrayWithDefaults

`func NewV0039JobArrayWithDefaults() *V0039JobArray`

NewV0039JobArrayWithDefaults instantiates a new V0039JobArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *V0039JobArray) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0039JobArray) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0039JobArray) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0039JobArray) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetLimits

`func (o *V0039JobArray) GetLimits() V0039JobArrayLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *V0039JobArray) GetLimitsOk() (*V0039JobArrayLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *V0039JobArray) SetLimits(v V0039JobArrayLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *V0039JobArray) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetTaskId

`func (o *V0039JobArray) GetTaskId() V0039Uint32NoVal`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *V0039JobArray) GetTaskIdOk() (*V0039Uint32NoVal, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *V0039JobArray) SetTaskId(v V0039Uint32NoVal)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *V0039JobArray) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTask

`func (o *V0039JobArray) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V0039JobArray) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V0039JobArray) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *V0039JobArray) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


