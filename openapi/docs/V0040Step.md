# V0040Step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**V0040StepTime**](V0040StepTime.md) |  | [optional] 
**ExitCode** | Pointer to [**V0040ProcessExitCodeVerbose**](V0040ProcessExitCodeVerbose.md) |  | [optional] 
**Nodes** | Pointer to [**V0040StepNodes**](V0040StepNodes.md) |  | [optional] 
**Tasks** | Pointer to [**V0039StepTasks**](V0039StepTasks.md) |  | [optional] 
**Pid** | Pointer to **string** |  | [optional] 
**CPU** | Pointer to [**V0040StepCPU**](V0040StepCPU.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** |  | [optional] 
**State** | Pointer to **[]string** |  | [optional] 
**Statistics** | Pointer to [**V0040StepStatistics**](V0040StepStatistics.md) |  | [optional] 
**Step** | Pointer to [**V0040StepStep**](V0040StepStep.md) |  | [optional] 
**Task** | Pointer to [**V0039StepTask**](V0039StepTask.md) |  | [optional] 
**Tres** | Pointer to [**V0040StepTres**](V0040StepTres.md) |  | [optional] 

## Methods

### NewV0040Step

`func NewV0040Step() *V0040Step`

NewV0040Step instantiates a new V0040Step object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040StepWithDefaults

`func NewV0040StepWithDefaults() *V0040Step`

NewV0040StepWithDefaults instantiates a new V0040Step object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *V0040Step) GetTime() V0040StepTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0040Step) GetTimeOk() (*V0040StepTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0040Step) SetTime(v V0040StepTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0040Step) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *V0040Step) GetExitCode() V0040ProcessExitCodeVerbose`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0040Step) GetExitCodeOk() (*V0040ProcessExitCodeVerbose, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0040Step) SetExitCode(v V0040ProcessExitCodeVerbose)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0040Step) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetNodes

`func (o *V0040Step) GetNodes() V0040StepNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0040Step) GetNodesOk() (*V0040StepNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0040Step) SetNodes(v V0040StepNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0040Step) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTasks

`func (o *V0040Step) GetTasks() V0039StepTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0040Step) GetTasksOk() (*V0039StepTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0040Step) SetTasks(v V0039StepTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0040Step) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPid

`func (o *V0040Step) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *V0040Step) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *V0040Step) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *V0040Step) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetCPU

`func (o *V0040Step) GetCPU() V0040StepCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *V0040Step) GetCPUOk() (*V0040StepCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *V0040Step) SetCPU(v V0040StepCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *V0040Step) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *V0040Step) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *V0040Step) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *V0040Step) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *V0040Step) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetState

`func (o *V0040Step) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0040Step) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0040Step) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0040Step) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatistics

`func (o *V0040Step) GetStatistics() V0040StepStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0040Step) GetStatisticsOk() (*V0040StepStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0040Step) SetStatistics(v V0040StepStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V0040Step) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetStep

`func (o *V0040Step) GetStep() V0040StepStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *V0040Step) GetStepOk() (*V0040StepStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *V0040Step) SetStep(v V0040StepStep)`

SetStep sets Step field to given value.

### HasStep

`func (o *V0040Step) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTask

`func (o *V0040Step) GetTask() V0039StepTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V0040Step) GetTaskOk() (*V0039StepTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V0040Step) SetTask(v V0039StepTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *V0040Step) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTres

`func (o *V0040Step) GetTres() V0040StepTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0040Step) GetTresOk() (*V0040StepTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0040Step) SetTres(v V0040StepTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0040Step) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


