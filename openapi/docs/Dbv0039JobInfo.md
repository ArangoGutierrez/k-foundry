# Dbv0039JobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0039Meta**](Dbv0039Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0039Error**](Dbv0039Error.md) | Slurm errors | [optional] 
**Warnings** | Pointer to [**[]Dbv0039Warning**](Dbv0039Warning.md) | Slurm warnings | [optional] 
**Jobs** | Pointer to [**[]V0039Job**](V0039Job.md) |  | [optional] 

## Methods

### NewDbv0039JobInfo

`func NewDbv0039JobInfo() *Dbv0039JobInfo`

NewDbv0039JobInfo instantiates a new Dbv0039JobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0039JobInfoWithDefaults

`func NewDbv0039JobInfoWithDefaults() *Dbv0039JobInfo`

NewDbv0039JobInfoWithDefaults instantiates a new Dbv0039JobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0039JobInfo) GetMeta() Dbv0039Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0039JobInfo) GetMetaOk() (*Dbv0039Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0039JobInfo) SetMeta(v Dbv0039Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0039JobInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0039JobInfo) GetErrors() []Dbv0039Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0039JobInfo) GetErrorsOk() (*[]Dbv0039Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0039JobInfo) SetErrors(v []Dbv0039Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0039JobInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *Dbv0039JobInfo) GetWarnings() []Dbv0039Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Dbv0039JobInfo) GetWarningsOk() (*[]Dbv0039Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Dbv0039JobInfo) SetWarnings(v []Dbv0039Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Dbv0039JobInfo) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetJobs

`func (o *Dbv0039JobInfo) GetJobs() []V0039Job`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *Dbv0039JobInfo) GetJobsOk() (*[]V0039Job, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *Dbv0039JobInfo) SetJobs(v []V0039Job)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *Dbv0039JobInfo) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


