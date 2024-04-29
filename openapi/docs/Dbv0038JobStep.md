# Dbv0038JobStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**Dbv0038JobStepTime**](Dbv0038JobStepTime.md) |  | [optional] 
**ExitCode** | Pointer to [**Dbv0038JobExitCode**](Dbv0038JobExitCode.md) |  | [optional] 
**Nodes** | Pointer to [**Dbv0038JobStepNodes**](Dbv0038JobStepNodes.md) |  | [optional] 
**Tasks** | Pointer to [**Dbv0038JobStepTasks**](Dbv0038JobStepTasks.md) |  | [optional] 
**Pid** | Pointer to **string** | First process PID | [optional] 
**CPU** | Pointer to [**Dbv0038JobStepCPU**](Dbv0038JobStepCPU.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** | User who requested job killed | [optional] 
**State** | Pointer to **string** | State of job step | [optional] 
**Statistics** | Pointer to [**Dbv0038JobStepStatistics**](Dbv0038JobStepStatistics.md) |  | [optional] 
**Step** | Pointer to [**Dbv0038JobStepStep**](Dbv0038JobStepStep.md) |  | [optional] 
**Task** | Pointer to **string** | Task distribution properties | [optional] 
**Tres** | Pointer to [**Dbv0038JobStepTres**](Dbv0038JobStepTres.md) |  | [optional] 

## Methods

### NewDbv0038JobStep

`func NewDbv0038JobStep() *Dbv0038JobStep`

NewDbv0038JobStep instantiates a new Dbv0038JobStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobStepWithDefaults

`func NewDbv0038JobStepWithDefaults() *Dbv0038JobStep`

NewDbv0038JobStepWithDefaults instantiates a new Dbv0038JobStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *Dbv0038JobStep) GetTime() Dbv0038JobStepTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Dbv0038JobStep) GetTimeOk() (*Dbv0038JobStepTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Dbv0038JobStep) SetTime(v Dbv0038JobStepTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *Dbv0038JobStep) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *Dbv0038JobStep) GetExitCode() Dbv0038JobExitCode`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Dbv0038JobStep) GetExitCodeOk() (*Dbv0038JobExitCode, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Dbv0038JobStep) SetExitCode(v Dbv0038JobExitCode)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Dbv0038JobStep) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetNodes

`func (o *Dbv0038JobStep) GetNodes() Dbv0038JobStepNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Dbv0038JobStep) GetNodesOk() (*Dbv0038JobStepNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Dbv0038JobStep) SetNodes(v Dbv0038JobStepNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Dbv0038JobStep) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTasks

`func (o *Dbv0038JobStep) GetTasks() Dbv0038JobStepTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Dbv0038JobStep) GetTasksOk() (*Dbv0038JobStepTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Dbv0038JobStep) SetTasks(v Dbv0038JobStepTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Dbv0038JobStep) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPid

`func (o *Dbv0038JobStep) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *Dbv0038JobStep) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *Dbv0038JobStep) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *Dbv0038JobStep) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetCPU

`func (o *Dbv0038JobStep) GetCPU() Dbv0038JobStepCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *Dbv0038JobStep) GetCPUOk() (*Dbv0038JobStepCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *Dbv0038JobStep) SetCPU(v Dbv0038JobStepCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *Dbv0038JobStep) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *Dbv0038JobStep) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *Dbv0038JobStep) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *Dbv0038JobStep) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *Dbv0038JobStep) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetState

`func (o *Dbv0038JobStep) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Dbv0038JobStep) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Dbv0038JobStep) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Dbv0038JobStep) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatistics

`func (o *Dbv0038JobStep) GetStatistics() Dbv0038JobStepStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Dbv0038JobStep) GetStatisticsOk() (*Dbv0038JobStepStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Dbv0038JobStep) SetStatistics(v Dbv0038JobStepStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Dbv0038JobStep) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetStep

`func (o *Dbv0038JobStep) GetStep() Dbv0038JobStepStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *Dbv0038JobStep) GetStepOk() (*Dbv0038JobStepStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *Dbv0038JobStep) SetStep(v Dbv0038JobStepStep)`

SetStep sets Step field to given value.

### HasStep

`func (o *Dbv0038JobStep) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTask

`func (o *Dbv0038JobStep) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Dbv0038JobStep) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Dbv0038JobStep) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *Dbv0038JobStep) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0038JobStep) GetTres() Dbv0038JobStepTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0038JobStep) GetTresOk() (*Dbv0038JobStepTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0038JobStep) SetTres(v Dbv0038JobStepTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0038JobStep) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


