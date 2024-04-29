# V0038JobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0038Meta**](V0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0038Error**](V0038Error.md) | slurm errors | [optional] 
**Jobs** | Pointer to [**[]V0038JobResponseProperties**](V0038JobResponseProperties.md) | job descriptions | [optional] 

## Methods

### NewV0038JobsResponse

`func NewV0038JobsResponse() *V0038JobsResponse`

NewV0038JobsResponse instantiates a new V0038JobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038JobsResponseWithDefaults

`func NewV0038JobsResponseWithDefaults() *V0038JobsResponse`

NewV0038JobsResponseWithDefaults instantiates a new V0038JobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0038JobsResponse) GetMeta() V0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0038JobsResponse) GetMetaOk() (*V0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0038JobsResponse) SetMeta(v V0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0038JobsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0038JobsResponse) GetErrors() []V0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0038JobsResponse) GetErrorsOk() (*[]V0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0038JobsResponse) SetErrors(v []V0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0038JobsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetJobs

`func (o *V0038JobsResponse) GetJobs() []V0038JobResponseProperties`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0038JobsResponse) GetJobsOk() (*[]V0038JobResponseProperties, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0038JobsResponse) SetJobs(v []V0038JobResponseProperties)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0038JobsResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


