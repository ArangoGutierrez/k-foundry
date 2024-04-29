# V0038JobSubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0038Meta**](V0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0038Error**](V0038Error.md) | slurm errors | [optional] 
**JobId** | Pointer to **int32** | new job ID | [optional] 
**StepId** | Pointer to **string** | new job step ID | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | Message to user from job_submit plugin | [optional] 

## Methods

### NewV0038JobSubmissionResponse

`func NewV0038JobSubmissionResponse() *V0038JobSubmissionResponse`

NewV0038JobSubmissionResponse instantiates a new V0038JobSubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038JobSubmissionResponseWithDefaults

`func NewV0038JobSubmissionResponseWithDefaults() *V0038JobSubmissionResponse`

NewV0038JobSubmissionResponseWithDefaults instantiates a new V0038JobSubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0038JobSubmissionResponse) GetMeta() V0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0038JobSubmissionResponse) GetMetaOk() (*V0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0038JobSubmissionResponse) SetMeta(v V0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0038JobSubmissionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0038JobSubmissionResponse) GetErrors() []V0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0038JobSubmissionResponse) GetErrorsOk() (*[]V0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0038JobSubmissionResponse) SetErrors(v []V0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0038JobSubmissionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetJobId

`func (o *V0038JobSubmissionResponse) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0038JobSubmissionResponse) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0038JobSubmissionResponse) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0038JobSubmissionResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *V0038JobSubmissionResponse) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0038JobSubmissionResponse) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0038JobSubmissionResponse) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0038JobSubmissionResponse) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *V0038JobSubmissionResponse) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *V0038JobSubmissionResponse) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *V0038JobSubmissionResponse) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *V0038JobSubmissionResponse) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


