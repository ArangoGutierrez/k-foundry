# V0038JobResponseProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Charge resources used by this job to specified account | [optional] 
**AccrueTime** | Pointer to **int64** | time job is eligible for running | [optional] 
**AdminComment** | Pointer to **string** | administrator&#39;s arbitrary comment | [optional] 
**ArrayJobId** | Pointer to **int32** | job_id of a job array or 0 if N/A | [optional] 
**ArrayTaskId** | Pointer to **int32** | task_id of a job array | [optional] 
**ArrayMaxTasks** | Pointer to **int32** | Maximum number of running array tasks | [optional] 
**ArrayTaskString** | Pointer to **string** | string expression of task IDs in this record | [optional] 
**AssociationId** | Pointer to **int32** | association id for job | [optional] 
**BatchFeatures** | Pointer to **string** | features required for batch script&#39;s node | [optional] 
**BatchFlag** | Pointer to **bool** | if batch: queued job with script | [optional] 
**BatchHost** | Pointer to **string** | name of host running batch script | [optional] 
**Flags** | Pointer to **[]string** | Job flags | [optional] 
**BurstBuffer** | Pointer to **string** | burst buffer specifications | [optional] 
**BurstBufferState** | Pointer to **string** | burst buffer state info | [optional] 
**Cluster** | Pointer to **string** | name of cluster that the job is on | [optional] 
**ClusterFeatures** | Pointer to **string** | comma separated list of required cluster features | [optional] 
**Command** | Pointer to **string** | command to be executed | [optional] 
**Comment** | Pointer to **string** | arbitrary comment | [optional] 
**Container** | Pointer to **string** | absolute path to OCI container bundle | [optional] 
**Contiguous** | Pointer to **bool** | job requires contiguous nodes | [optional] 
**CoreSpec** | Pointer to **string** | specialized core count | [optional] 
**ThreadSpec** | Pointer to **string** | specialized thread count | [optional] 
**CoresPerSocket** | Pointer to **string** | cores per socket required by job | [optional] 
**BillableTres** | Pointer to **float32** | billable TRES | [optional] 
**CpusPerTask** | Pointer to **string** | number of processors required for each task | [optional] 
**CpuFrequencyMinimum** | Pointer to **string** | Minimum cpu frequency | [optional] 
**CpuFrequencyMaximum** | Pointer to **string** | Maximum cpu frequency | [optional] 
**CpuFrequencyGovernor** | Pointer to **string** | cpu frequency governor | [optional] 
**CpusPerTres** | Pointer to **string** | semicolon delimited list of TRES&#x3D;# values | [optional] 
**Deadline** | Pointer to **int32** | job start deadline  | [optional] 
**DelayBoot** | Pointer to **int32** | command to be executed | [optional] 
**Dependency** | Pointer to **string** | synchronize job execution with other jobs | [optional] 
**DerivedExitCode** | Pointer to **int32** | highest exit code of all job steps | [optional] 
**EligibleTime** | Pointer to **int64** | time job is eligible for running | [optional] 
**EndTime** | Pointer to **int64** | time of termination, actual or expected | [optional] 
**ExcludedNodes** | Pointer to **string** | comma separated list of excluded nodes | [optional] 
**ExitCode** | Pointer to **int32** | exit code for job | [optional] 
**Features** | Pointer to **string** | comma separated list of required features | [optional] 
**FederationOrigin** | Pointer to **string** | Origin cluster&#39;s name | [optional] 
**FederationSiblingsActive** | Pointer to **string** | string of active sibling names | [optional] 
**FederationSiblingsViable** | Pointer to **string** | string of viable sibling names | [optional] 
**GresDetail** | Pointer to **[]string** | Job flags | [optional] 
**GroupId** | Pointer to **int32** | group job submitted as | [optional] 
**JobId** | Pointer to **int32** | job ID | [optional] 
**JobResources** | Pointer to [**V0038JobResources**](V0038JobResources.md) |  | [optional] 
**JobState** | Pointer to **string** | state of the job | [optional] 
**LastSchedEvaluation** | Pointer to **int32** | last time job was evaluated for scheduling | [optional] 
**Licenses** | Pointer to **string** | licenses required by the job | [optional] 
**MaxCpus** | Pointer to **int32** | maximum number of cpus usable by job | [optional] 
**MaxNodes** | Pointer to **int32** | maximum number of nodes usable by job | [optional] 
**McsLabel** | Pointer to **string** | mcs_label if mcs plugin in use | [optional] 
**MemoryPerTres** | Pointer to **string** | semicolon delimited list of TRES&#x3D;# values | [optional] 
**Name** | Pointer to **string** | name of the job | [optional] 
**Nodes** | Pointer to **string** | list of nodes allocated to job | [optional] 
**Nice** | Pointer to **int32** | requested priority change | [optional] 
**TasksPerCore** | Pointer to **int32** | number of tasks to invoke on each core | [optional] 
**TasksPerSocket** | Pointer to **int32** | number of tasks to invoke on each socket | [optional] 
**TasksPerBoard** | Pointer to **int32** | number of tasks to invoke on each board | [optional] 
**Cpus** | Pointer to **int32** | minimum number of cpus required by job | [optional] 
**NodeCount** | Pointer to **int32** | minimum number of nodes required by job | [optional] 
**Tasks** | Pointer to **int32** | requested task count | [optional] 
**HetJobId** | Pointer to **int32** | job ID of hetjob leader | [optional] 
**HetJobIdSet** | Pointer to **string** | job IDs for all components | [optional] 
**HetJobOffset** | Pointer to **int32** | HetJob component offset from leader | [optional] 
**Partition** | Pointer to **string** | name of assigned partition | [optional] 
**MemoryPerNode** | Pointer to **int32** | minimum real memory per node | [optional] 
**MemoryPerCpu** | Pointer to **int32** | minimum real memory per cpu | [optional] 
**MinimumCpusPerNode** | Pointer to **int32** | minimum # CPUs per node | [optional] 
**MinimumTmpDiskPerNode** | Pointer to **int32** | minimum tmp disk per node | [optional] 
**PreemptTime** | Pointer to **int64** | preemption signal time | [optional] 
**PreSusTime** | Pointer to **int64** | time job ran prior to last suspend | [optional] 
**Priority** | Pointer to **int32** | relative priority of the job | [optional] 
**Profile** | Pointer to **[]string** | Job profiling requested | [optional] 
**Qos** | Pointer to **string** | Quality of Service | [optional] 
**Reboot** | Pointer to **bool** | node reboot requested before start | [optional] 
**RequiredNodes** | Pointer to **string** | comma separated list of required nodes | [optional] 
**Requeue** | Pointer to **bool** | enable or disable job requeue option | [optional] 
**ResizeTime** | Pointer to **int64** | time of latest size change | [optional] 
**RestartCnt** | Pointer to **int32** | count of job restarts | [optional] 
**ResvName** | Pointer to **string** | reservation name | [optional] 
**Shared** | Pointer to **string** | type and if job can share nodes with other jobs | [optional] 
**ShowFlags** | Pointer to **[]string** | details requested | [optional] 
**SocketsPerBoard** | Pointer to **int32** | sockets per board required by job | [optional] 
**SocketsPerNode** | Pointer to **int32** | sockets per node required by job | [optional] 
**StartTime** | Pointer to **int64** | time execution begins, actual or expected | [optional] 
**StateDescription** | Pointer to **string** | optional details for state_reason | [optional] 
**StateReason** | Pointer to **string** | reason job still pending or failed | [optional] 
**StandardError** | Pointer to **string** | pathname of job&#39;s stderr file | [optional] 
**StandardInput** | Pointer to **string** | pathname of job&#39;s stdin file | [optional] 
**StandardOutput** | Pointer to **string** | pathname of job&#39;s stdout file | [optional] 
**SubmitTime** | Pointer to **int64** | time of job submission | [optional] 
**SuspendTime** | Pointer to **int64** | time job last suspended or resumed | [optional] 
**SystemComment** | Pointer to **string** | slurmctld&#39;s arbitrary comment | [optional] 
**TimeLimit** | Pointer to **int64** | maximum run time in minutes | [optional] 
**TimeMinimum** | Pointer to **int64** | minimum run time in minutes | [optional] 
**ThreadsPerCore** | Pointer to **int32** | threads per core required by job | [optional] 
**TresBind** | Pointer to **string** | Task to TRES binding directives | [optional] 
**TresFreq** | Pointer to **string** | TRES frequency directives | [optional] 
**TresPerJob** | Pointer to **string** | semicolon delimited list of TRES&#x3D;# values | [optional] 
**TresPerNode** | Pointer to **string** | semicolon delimited list of TRES&#x3D;# values | [optional] 
**TresPerSocket** | Pointer to **string** | semicolon delimited list of TRES&#x3D;# values | [optional] 
**TresPerTask** | Pointer to **string** | semicolon delimited list of TRES&#x3D;# values | [optional] 
**TresReqStr** | Pointer to **string** | tres requested in the job | [optional] 
**TresAllocStr** | Pointer to **string** | tres used in the job | [optional] 
**UserId** | Pointer to **int64** | user id the job runs as | [optional] 
**UserName** | Pointer to **string** | user the job runs as | [optional] 
**Wckey** | Pointer to **string** | wckey for job | [optional] 
**CurrentWorkingDirectory** | Pointer to **string** | pathname of working directory | [optional] 

## Methods

### NewV0038JobResponseProperties

`func NewV0038JobResponseProperties() *V0038JobResponseProperties`

NewV0038JobResponseProperties instantiates a new V0038JobResponseProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038JobResponsePropertiesWithDefaults

`func NewV0038JobResponsePropertiesWithDefaults() *V0038JobResponseProperties`

NewV0038JobResponsePropertiesWithDefaults instantiates a new V0038JobResponseProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0038JobResponseProperties) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0038JobResponseProperties) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0038JobResponseProperties) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0038JobResponseProperties) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccrueTime

`func (o *V0038JobResponseProperties) GetAccrueTime() int64`

GetAccrueTime returns the AccrueTime field if non-nil, zero value otherwise.

### GetAccrueTimeOk

`func (o *V0038JobResponseProperties) GetAccrueTimeOk() (*int64, bool)`

GetAccrueTimeOk returns a tuple with the AccrueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrueTime

`func (o *V0038JobResponseProperties) SetAccrueTime(v int64)`

SetAccrueTime sets AccrueTime field to given value.

### HasAccrueTime

`func (o *V0038JobResponseProperties) HasAccrueTime() bool`

HasAccrueTime returns a boolean if a field has been set.

### GetAdminComment

`func (o *V0038JobResponseProperties) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *V0038JobResponseProperties) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *V0038JobResponseProperties) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *V0038JobResponseProperties) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetArrayJobId

`func (o *V0038JobResponseProperties) GetArrayJobId() int32`

GetArrayJobId returns the ArrayJobId field if non-nil, zero value otherwise.

### GetArrayJobIdOk

`func (o *V0038JobResponseProperties) GetArrayJobIdOk() (*int32, bool)`

GetArrayJobIdOk returns a tuple with the ArrayJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayJobId

`func (o *V0038JobResponseProperties) SetArrayJobId(v int32)`

SetArrayJobId sets ArrayJobId field to given value.

### HasArrayJobId

`func (o *V0038JobResponseProperties) HasArrayJobId() bool`

HasArrayJobId returns a boolean if a field has been set.

### GetArrayTaskId

`func (o *V0038JobResponseProperties) GetArrayTaskId() int32`

GetArrayTaskId returns the ArrayTaskId field if non-nil, zero value otherwise.

### GetArrayTaskIdOk

`func (o *V0038JobResponseProperties) GetArrayTaskIdOk() (*int32, bool)`

GetArrayTaskIdOk returns a tuple with the ArrayTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskId

`func (o *V0038JobResponseProperties) SetArrayTaskId(v int32)`

SetArrayTaskId sets ArrayTaskId field to given value.

### HasArrayTaskId

`func (o *V0038JobResponseProperties) HasArrayTaskId() bool`

HasArrayTaskId returns a boolean if a field has been set.

### GetArrayMaxTasks

`func (o *V0038JobResponseProperties) GetArrayMaxTasks() int32`

GetArrayMaxTasks returns the ArrayMaxTasks field if non-nil, zero value otherwise.

### GetArrayMaxTasksOk

`func (o *V0038JobResponseProperties) GetArrayMaxTasksOk() (*int32, bool)`

GetArrayMaxTasksOk returns a tuple with the ArrayMaxTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayMaxTasks

`func (o *V0038JobResponseProperties) SetArrayMaxTasks(v int32)`

SetArrayMaxTasks sets ArrayMaxTasks field to given value.

### HasArrayMaxTasks

`func (o *V0038JobResponseProperties) HasArrayMaxTasks() bool`

HasArrayMaxTasks returns a boolean if a field has been set.

### GetArrayTaskString

`func (o *V0038JobResponseProperties) GetArrayTaskString() string`

GetArrayTaskString returns the ArrayTaskString field if non-nil, zero value otherwise.

### GetArrayTaskStringOk

`func (o *V0038JobResponseProperties) GetArrayTaskStringOk() (*string, bool)`

GetArrayTaskStringOk returns a tuple with the ArrayTaskString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskString

`func (o *V0038JobResponseProperties) SetArrayTaskString(v string)`

SetArrayTaskString sets ArrayTaskString field to given value.

### HasArrayTaskString

`func (o *V0038JobResponseProperties) HasArrayTaskString() bool`

HasArrayTaskString returns a boolean if a field has been set.

### GetAssociationId

`func (o *V0038JobResponseProperties) GetAssociationId() int32`

GetAssociationId returns the AssociationId field if non-nil, zero value otherwise.

### GetAssociationIdOk

`func (o *V0038JobResponseProperties) GetAssociationIdOk() (*int32, bool)`

GetAssociationIdOk returns a tuple with the AssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationId

`func (o *V0038JobResponseProperties) SetAssociationId(v int32)`

SetAssociationId sets AssociationId field to given value.

### HasAssociationId

`func (o *V0038JobResponseProperties) HasAssociationId() bool`

HasAssociationId returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *V0038JobResponseProperties) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *V0038JobResponseProperties) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *V0038JobResponseProperties) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *V0038JobResponseProperties) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBatchFlag

`func (o *V0038JobResponseProperties) GetBatchFlag() bool`

GetBatchFlag returns the BatchFlag field if non-nil, zero value otherwise.

### GetBatchFlagOk

`func (o *V0038JobResponseProperties) GetBatchFlagOk() (*bool, bool)`

GetBatchFlagOk returns a tuple with the BatchFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFlag

`func (o *V0038JobResponseProperties) SetBatchFlag(v bool)`

SetBatchFlag sets BatchFlag field to given value.

### HasBatchFlag

`func (o *V0038JobResponseProperties) HasBatchFlag() bool`

HasBatchFlag returns a boolean if a field has been set.

### GetBatchHost

`func (o *V0038JobResponseProperties) GetBatchHost() string`

GetBatchHost returns the BatchHost field if non-nil, zero value otherwise.

### GetBatchHostOk

`func (o *V0038JobResponseProperties) GetBatchHostOk() (*string, bool)`

GetBatchHostOk returns a tuple with the BatchHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchHost

`func (o *V0038JobResponseProperties) SetBatchHost(v string)`

SetBatchHost sets BatchHost field to given value.

### HasBatchHost

`func (o *V0038JobResponseProperties) HasBatchHost() bool`

HasBatchHost returns a boolean if a field has been set.

### GetFlags

`func (o *V0038JobResponseProperties) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0038JobResponseProperties) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0038JobResponseProperties) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0038JobResponseProperties) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0038JobResponseProperties) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0038JobResponseProperties) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0038JobResponseProperties) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0038JobResponseProperties) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetBurstBufferState

`func (o *V0038JobResponseProperties) GetBurstBufferState() string`

GetBurstBufferState returns the BurstBufferState field if non-nil, zero value otherwise.

### GetBurstBufferStateOk

`func (o *V0038JobResponseProperties) GetBurstBufferStateOk() (*string, bool)`

GetBurstBufferStateOk returns a tuple with the BurstBufferState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBufferState

`func (o *V0038JobResponseProperties) SetBurstBufferState(v string)`

SetBurstBufferState sets BurstBufferState field to given value.

### HasBurstBufferState

`func (o *V0038JobResponseProperties) HasBurstBufferState() bool`

HasBurstBufferState returns a boolean if a field has been set.

### GetCluster

`func (o *V0038JobResponseProperties) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0038JobResponseProperties) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0038JobResponseProperties) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0038JobResponseProperties) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterFeatures

`func (o *V0038JobResponseProperties) GetClusterFeatures() string`

GetClusterFeatures returns the ClusterFeatures field if non-nil, zero value otherwise.

### GetClusterFeaturesOk

`func (o *V0038JobResponseProperties) GetClusterFeaturesOk() (*string, bool)`

GetClusterFeaturesOk returns a tuple with the ClusterFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterFeatures

`func (o *V0038JobResponseProperties) SetClusterFeatures(v string)`

SetClusterFeatures sets ClusterFeatures field to given value.

### HasClusterFeatures

`func (o *V0038JobResponseProperties) HasClusterFeatures() bool`

HasClusterFeatures returns a boolean if a field has been set.

### GetCommand

`func (o *V0038JobResponseProperties) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0038JobResponseProperties) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0038JobResponseProperties) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0038JobResponseProperties) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetComment

`func (o *V0038JobResponseProperties) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0038JobResponseProperties) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0038JobResponseProperties) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0038JobResponseProperties) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContainer

`func (o *V0038JobResponseProperties) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0038JobResponseProperties) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0038JobResponseProperties) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0038JobResponseProperties) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContiguous

`func (o *V0038JobResponseProperties) GetContiguous() bool`

GetContiguous returns the Contiguous field if non-nil, zero value otherwise.

### GetContiguousOk

`func (o *V0038JobResponseProperties) GetContiguousOk() (*bool, bool)`

GetContiguousOk returns a tuple with the Contiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContiguous

`func (o *V0038JobResponseProperties) SetContiguous(v bool)`

SetContiguous sets Contiguous field to given value.

### HasContiguous

`func (o *V0038JobResponseProperties) HasContiguous() bool`

HasContiguous returns a boolean if a field has been set.

### GetCoreSpec

`func (o *V0038JobResponseProperties) GetCoreSpec() string`

GetCoreSpec returns the CoreSpec field if non-nil, zero value otherwise.

### GetCoreSpecOk

`func (o *V0038JobResponseProperties) GetCoreSpecOk() (*string, bool)`

GetCoreSpecOk returns a tuple with the CoreSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpec

`func (o *V0038JobResponseProperties) SetCoreSpec(v string)`

SetCoreSpec sets CoreSpec field to given value.

### HasCoreSpec

`func (o *V0038JobResponseProperties) HasCoreSpec() bool`

HasCoreSpec returns a boolean if a field has been set.

### GetThreadSpec

`func (o *V0038JobResponseProperties) GetThreadSpec() string`

GetThreadSpec returns the ThreadSpec field if non-nil, zero value otherwise.

### GetThreadSpecOk

`func (o *V0038JobResponseProperties) GetThreadSpecOk() (*string, bool)`

GetThreadSpecOk returns a tuple with the ThreadSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpec

`func (o *V0038JobResponseProperties) SetThreadSpec(v string)`

SetThreadSpec sets ThreadSpec field to given value.

### HasThreadSpec

`func (o *V0038JobResponseProperties) HasThreadSpec() bool`

HasThreadSpec returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *V0038JobResponseProperties) GetCoresPerSocket() string`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *V0038JobResponseProperties) GetCoresPerSocketOk() (*string, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *V0038JobResponseProperties) SetCoresPerSocket(v string)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *V0038JobResponseProperties) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetBillableTres

`func (o *V0038JobResponseProperties) GetBillableTres() float32`

GetBillableTres returns the BillableTres field if non-nil, zero value otherwise.

### GetBillableTresOk

`func (o *V0038JobResponseProperties) GetBillableTresOk() (*float32, bool)`

GetBillableTresOk returns a tuple with the BillableTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableTres

`func (o *V0038JobResponseProperties) SetBillableTres(v float32)`

SetBillableTres sets BillableTres field to given value.

### HasBillableTres

`func (o *V0038JobResponseProperties) HasBillableTres() bool`

HasBillableTres returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *V0038JobResponseProperties) GetCpusPerTask() string`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *V0038JobResponseProperties) GetCpusPerTaskOk() (*string, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *V0038JobResponseProperties) SetCpusPerTask(v string)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *V0038JobResponseProperties) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetCpuFrequencyMinimum

`func (o *V0038JobResponseProperties) GetCpuFrequencyMinimum() string`

GetCpuFrequencyMinimum returns the CpuFrequencyMinimum field if non-nil, zero value otherwise.

### GetCpuFrequencyMinimumOk

`func (o *V0038JobResponseProperties) GetCpuFrequencyMinimumOk() (*string, bool)`

GetCpuFrequencyMinimumOk returns a tuple with the CpuFrequencyMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMinimum

`func (o *V0038JobResponseProperties) SetCpuFrequencyMinimum(v string)`

SetCpuFrequencyMinimum sets CpuFrequencyMinimum field to given value.

### HasCpuFrequencyMinimum

`func (o *V0038JobResponseProperties) HasCpuFrequencyMinimum() bool`

HasCpuFrequencyMinimum returns a boolean if a field has been set.

### GetCpuFrequencyMaximum

`func (o *V0038JobResponseProperties) GetCpuFrequencyMaximum() string`

GetCpuFrequencyMaximum returns the CpuFrequencyMaximum field if non-nil, zero value otherwise.

### GetCpuFrequencyMaximumOk

`func (o *V0038JobResponseProperties) GetCpuFrequencyMaximumOk() (*string, bool)`

GetCpuFrequencyMaximumOk returns a tuple with the CpuFrequencyMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMaximum

`func (o *V0038JobResponseProperties) SetCpuFrequencyMaximum(v string)`

SetCpuFrequencyMaximum sets CpuFrequencyMaximum field to given value.

### HasCpuFrequencyMaximum

`func (o *V0038JobResponseProperties) HasCpuFrequencyMaximum() bool`

HasCpuFrequencyMaximum returns a boolean if a field has been set.

### GetCpuFrequencyGovernor

`func (o *V0038JobResponseProperties) GetCpuFrequencyGovernor() string`

GetCpuFrequencyGovernor returns the CpuFrequencyGovernor field if non-nil, zero value otherwise.

### GetCpuFrequencyGovernorOk

`func (o *V0038JobResponseProperties) GetCpuFrequencyGovernorOk() (*string, bool)`

GetCpuFrequencyGovernorOk returns a tuple with the CpuFrequencyGovernor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyGovernor

`func (o *V0038JobResponseProperties) SetCpuFrequencyGovernor(v string)`

SetCpuFrequencyGovernor sets CpuFrequencyGovernor field to given value.

### HasCpuFrequencyGovernor

`func (o *V0038JobResponseProperties) HasCpuFrequencyGovernor() bool`

HasCpuFrequencyGovernor returns a boolean if a field has been set.

### GetCpusPerTres

`func (o *V0038JobResponseProperties) GetCpusPerTres() string`

GetCpusPerTres returns the CpusPerTres field if non-nil, zero value otherwise.

### GetCpusPerTresOk

`func (o *V0038JobResponseProperties) GetCpusPerTresOk() (*string, bool)`

GetCpusPerTresOk returns a tuple with the CpusPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTres

`func (o *V0038JobResponseProperties) SetCpusPerTres(v string)`

SetCpusPerTres sets CpusPerTres field to given value.

### HasCpusPerTres

`func (o *V0038JobResponseProperties) HasCpusPerTres() bool`

HasCpusPerTres returns a boolean if a field has been set.

### GetDeadline

`func (o *V0038JobResponseProperties) GetDeadline() int32`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *V0038JobResponseProperties) GetDeadlineOk() (*int32, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *V0038JobResponseProperties) SetDeadline(v int32)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *V0038JobResponseProperties) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *V0038JobResponseProperties) GetDelayBoot() int32`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *V0038JobResponseProperties) GetDelayBootOk() (*int32, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *V0038JobResponseProperties) SetDelayBoot(v int32)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *V0038JobResponseProperties) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *V0038JobResponseProperties) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *V0038JobResponseProperties) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *V0038JobResponseProperties) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *V0038JobResponseProperties) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *V0038JobResponseProperties) GetDerivedExitCode() int32`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *V0038JobResponseProperties) GetDerivedExitCodeOk() (*int32, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *V0038JobResponseProperties) SetDerivedExitCode(v int32)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *V0038JobResponseProperties) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetEligibleTime

`func (o *V0038JobResponseProperties) GetEligibleTime() int64`

GetEligibleTime returns the EligibleTime field if non-nil, zero value otherwise.

### GetEligibleTimeOk

`func (o *V0038JobResponseProperties) GetEligibleTimeOk() (*int64, bool)`

GetEligibleTimeOk returns a tuple with the EligibleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleTime

`func (o *V0038JobResponseProperties) SetEligibleTime(v int64)`

SetEligibleTime sets EligibleTime field to given value.

### HasEligibleTime

`func (o *V0038JobResponseProperties) HasEligibleTime() bool`

HasEligibleTime returns a boolean if a field has been set.

### GetEndTime

`func (o *V0038JobResponseProperties) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0038JobResponseProperties) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0038JobResponseProperties) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0038JobResponseProperties) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExcludedNodes

`func (o *V0038JobResponseProperties) GetExcludedNodes() string`

GetExcludedNodes returns the ExcludedNodes field if non-nil, zero value otherwise.

### GetExcludedNodesOk

`func (o *V0038JobResponseProperties) GetExcludedNodesOk() (*string, bool)`

GetExcludedNodesOk returns a tuple with the ExcludedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNodes

`func (o *V0038JobResponseProperties) SetExcludedNodes(v string)`

SetExcludedNodes sets ExcludedNodes field to given value.

### HasExcludedNodes

`func (o *V0038JobResponseProperties) HasExcludedNodes() bool`

HasExcludedNodes returns a boolean if a field has been set.

### GetExitCode

`func (o *V0038JobResponseProperties) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0038JobResponseProperties) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0038JobResponseProperties) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0038JobResponseProperties) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetFeatures

`func (o *V0038JobResponseProperties) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0038JobResponseProperties) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0038JobResponseProperties) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0038JobResponseProperties) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFederationOrigin

`func (o *V0038JobResponseProperties) GetFederationOrigin() string`

GetFederationOrigin returns the FederationOrigin field if non-nil, zero value otherwise.

### GetFederationOriginOk

`func (o *V0038JobResponseProperties) GetFederationOriginOk() (*string, bool)`

GetFederationOriginOk returns a tuple with the FederationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationOrigin

`func (o *V0038JobResponseProperties) SetFederationOrigin(v string)`

SetFederationOrigin sets FederationOrigin field to given value.

### HasFederationOrigin

`func (o *V0038JobResponseProperties) HasFederationOrigin() bool`

HasFederationOrigin returns a boolean if a field has been set.

### GetFederationSiblingsActive

`func (o *V0038JobResponseProperties) GetFederationSiblingsActive() string`

GetFederationSiblingsActive returns the FederationSiblingsActive field if non-nil, zero value otherwise.

### GetFederationSiblingsActiveOk

`func (o *V0038JobResponseProperties) GetFederationSiblingsActiveOk() (*string, bool)`

GetFederationSiblingsActiveOk returns a tuple with the FederationSiblingsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsActive

`func (o *V0038JobResponseProperties) SetFederationSiblingsActive(v string)`

SetFederationSiblingsActive sets FederationSiblingsActive field to given value.

### HasFederationSiblingsActive

`func (o *V0038JobResponseProperties) HasFederationSiblingsActive() bool`

HasFederationSiblingsActive returns a boolean if a field has been set.

### GetFederationSiblingsViable

`func (o *V0038JobResponseProperties) GetFederationSiblingsViable() string`

GetFederationSiblingsViable returns the FederationSiblingsViable field if non-nil, zero value otherwise.

### GetFederationSiblingsViableOk

`func (o *V0038JobResponseProperties) GetFederationSiblingsViableOk() (*string, bool)`

GetFederationSiblingsViableOk returns a tuple with the FederationSiblingsViable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsViable

`func (o *V0038JobResponseProperties) SetFederationSiblingsViable(v string)`

SetFederationSiblingsViable sets FederationSiblingsViable field to given value.

### HasFederationSiblingsViable

`func (o *V0038JobResponseProperties) HasFederationSiblingsViable() bool`

HasFederationSiblingsViable returns a boolean if a field has been set.

### GetGresDetail

`func (o *V0038JobResponseProperties) GetGresDetail() []string`

GetGresDetail returns the GresDetail field if non-nil, zero value otherwise.

### GetGresDetailOk

`func (o *V0038JobResponseProperties) GetGresDetailOk() (*[]string, bool)`

GetGresDetailOk returns a tuple with the GresDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresDetail

`func (o *V0038JobResponseProperties) SetGresDetail(v []string)`

SetGresDetail sets GresDetail field to given value.

### HasGresDetail

`func (o *V0038JobResponseProperties) HasGresDetail() bool`

HasGresDetail returns a boolean if a field has been set.

### GetGroupId

`func (o *V0038JobResponseProperties) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V0038JobResponseProperties) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V0038JobResponseProperties) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V0038JobResponseProperties) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetJobId

`func (o *V0038JobResponseProperties) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0038JobResponseProperties) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0038JobResponseProperties) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0038JobResponseProperties) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobResources

`func (o *V0038JobResponseProperties) GetJobResources() V0038JobResources`

GetJobResources returns the JobResources field if non-nil, zero value otherwise.

### GetJobResourcesOk

`func (o *V0038JobResponseProperties) GetJobResourcesOk() (*V0038JobResources, bool)`

GetJobResourcesOk returns a tuple with the JobResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResources

`func (o *V0038JobResponseProperties) SetJobResources(v V0038JobResources)`

SetJobResources sets JobResources field to given value.

### HasJobResources

`func (o *V0038JobResponseProperties) HasJobResources() bool`

HasJobResources returns a boolean if a field has been set.

### GetJobState

`func (o *V0038JobResponseProperties) GetJobState() string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *V0038JobResponseProperties) GetJobStateOk() (*string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *V0038JobResponseProperties) SetJobState(v string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *V0038JobResponseProperties) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetLastSchedEvaluation

`func (o *V0038JobResponseProperties) GetLastSchedEvaluation() int32`

GetLastSchedEvaluation returns the LastSchedEvaluation field if non-nil, zero value otherwise.

### GetLastSchedEvaluationOk

`func (o *V0038JobResponseProperties) GetLastSchedEvaluationOk() (*int32, bool)`

GetLastSchedEvaluationOk returns a tuple with the LastSchedEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchedEvaluation

`func (o *V0038JobResponseProperties) SetLastSchedEvaluation(v int32)`

SetLastSchedEvaluation sets LastSchedEvaluation field to given value.

### HasLastSchedEvaluation

`func (o *V0038JobResponseProperties) HasLastSchedEvaluation() bool`

HasLastSchedEvaluation returns a boolean if a field has been set.

### GetLicenses

`func (o *V0038JobResponseProperties) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0038JobResponseProperties) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0038JobResponseProperties) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0038JobResponseProperties) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMaxCpus

`func (o *V0038JobResponseProperties) GetMaxCpus() int32`

GetMaxCpus returns the MaxCpus field if non-nil, zero value otherwise.

### GetMaxCpusOk

`func (o *V0038JobResponseProperties) GetMaxCpusOk() (*int32, bool)`

GetMaxCpusOk returns a tuple with the MaxCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpus

`func (o *V0038JobResponseProperties) SetMaxCpus(v int32)`

SetMaxCpus sets MaxCpus field to given value.

### HasMaxCpus

`func (o *V0038JobResponseProperties) HasMaxCpus() bool`

HasMaxCpus returns a boolean if a field has been set.

### GetMaxNodes

`func (o *V0038JobResponseProperties) GetMaxNodes() int32`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *V0038JobResponseProperties) GetMaxNodesOk() (*int32, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *V0038JobResponseProperties) SetMaxNodes(v int32)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *V0038JobResponseProperties) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0038JobResponseProperties) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0038JobResponseProperties) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0038JobResponseProperties) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0038JobResponseProperties) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryPerTres

`func (o *V0038JobResponseProperties) GetMemoryPerTres() string`

GetMemoryPerTres returns the MemoryPerTres field if non-nil, zero value otherwise.

### GetMemoryPerTresOk

`func (o *V0038JobResponseProperties) GetMemoryPerTresOk() (*string, bool)`

GetMemoryPerTresOk returns a tuple with the MemoryPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerTres

`func (o *V0038JobResponseProperties) SetMemoryPerTres(v string)`

SetMemoryPerTres sets MemoryPerTres field to given value.

### HasMemoryPerTres

`func (o *V0038JobResponseProperties) HasMemoryPerTres() bool`

HasMemoryPerTres returns a boolean if a field has been set.

### GetName

`func (o *V0038JobResponseProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0038JobResponseProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0038JobResponseProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0038JobResponseProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *V0038JobResponseProperties) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0038JobResponseProperties) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0038JobResponseProperties) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0038JobResponseProperties) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNice

`func (o *V0038JobResponseProperties) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *V0038JobResponseProperties) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *V0038JobResponseProperties) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *V0038JobResponseProperties) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *V0038JobResponseProperties) GetTasksPerCore() int32`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *V0038JobResponseProperties) GetTasksPerCoreOk() (*int32, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *V0038JobResponseProperties) SetTasksPerCore(v int32)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *V0038JobResponseProperties) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *V0038JobResponseProperties) GetTasksPerSocket() int32`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *V0038JobResponseProperties) GetTasksPerSocketOk() (*int32, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *V0038JobResponseProperties) SetTasksPerSocket(v int32)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *V0038JobResponseProperties) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetTasksPerBoard

`func (o *V0038JobResponseProperties) GetTasksPerBoard() int32`

GetTasksPerBoard returns the TasksPerBoard field if non-nil, zero value otherwise.

### GetTasksPerBoardOk

`func (o *V0038JobResponseProperties) GetTasksPerBoardOk() (*int32, bool)`

GetTasksPerBoardOk returns a tuple with the TasksPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerBoard

`func (o *V0038JobResponseProperties) SetTasksPerBoard(v int32)`

SetTasksPerBoard sets TasksPerBoard field to given value.

### HasTasksPerBoard

`func (o *V0038JobResponseProperties) HasTasksPerBoard() bool`

HasTasksPerBoard returns a boolean if a field has been set.

### GetCpus

`func (o *V0038JobResponseProperties) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0038JobResponseProperties) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0038JobResponseProperties) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0038JobResponseProperties) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0038JobResponseProperties) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0038JobResponseProperties) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0038JobResponseProperties) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0038JobResponseProperties) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTasks

`func (o *V0038JobResponseProperties) GetTasks() int32`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0038JobResponseProperties) GetTasksOk() (*int32, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0038JobResponseProperties) SetTasks(v int32)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0038JobResponseProperties) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetHetJobId

`func (o *V0038JobResponseProperties) GetHetJobId() int32`

GetHetJobId returns the HetJobId field if non-nil, zero value otherwise.

### GetHetJobIdOk

`func (o *V0038JobResponseProperties) GetHetJobIdOk() (*int32, bool)`

GetHetJobIdOk returns a tuple with the HetJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobId

`func (o *V0038JobResponseProperties) SetHetJobId(v int32)`

SetHetJobId sets HetJobId field to given value.

### HasHetJobId

`func (o *V0038JobResponseProperties) HasHetJobId() bool`

HasHetJobId returns a boolean if a field has been set.

### GetHetJobIdSet

`func (o *V0038JobResponseProperties) GetHetJobIdSet() string`

GetHetJobIdSet returns the HetJobIdSet field if non-nil, zero value otherwise.

### GetHetJobIdSetOk

`func (o *V0038JobResponseProperties) GetHetJobIdSetOk() (*string, bool)`

GetHetJobIdSetOk returns a tuple with the HetJobIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobIdSet

`func (o *V0038JobResponseProperties) SetHetJobIdSet(v string)`

SetHetJobIdSet sets HetJobIdSet field to given value.

### HasHetJobIdSet

`func (o *V0038JobResponseProperties) HasHetJobIdSet() bool`

HasHetJobIdSet returns a boolean if a field has been set.

### GetHetJobOffset

`func (o *V0038JobResponseProperties) GetHetJobOffset() int32`

GetHetJobOffset returns the HetJobOffset field if non-nil, zero value otherwise.

### GetHetJobOffsetOk

`func (o *V0038JobResponseProperties) GetHetJobOffsetOk() (*int32, bool)`

GetHetJobOffsetOk returns a tuple with the HetJobOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobOffset

`func (o *V0038JobResponseProperties) SetHetJobOffset(v int32)`

SetHetJobOffset sets HetJobOffset field to given value.

### HasHetJobOffset

`func (o *V0038JobResponseProperties) HasHetJobOffset() bool`

HasHetJobOffset returns a boolean if a field has been set.

### GetPartition

`func (o *V0038JobResponseProperties) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0038JobResponseProperties) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0038JobResponseProperties) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0038JobResponseProperties) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0038JobResponseProperties) GetMemoryPerNode() int32`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0038JobResponseProperties) GetMemoryPerNodeOk() (*int32, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0038JobResponseProperties) SetMemoryPerNode(v int32)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0038JobResponseProperties) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0038JobResponseProperties) GetMemoryPerCpu() int32`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0038JobResponseProperties) GetMemoryPerCpuOk() (*int32, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0038JobResponseProperties) SetMemoryPerCpu(v int32)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0038JobResponseProperties) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *V0038JobResponseProperties) GetMinimumCpusPerNode() int32`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *V0038JobResponseProperties) GetMinimumCpusPerNodeOk() (*int32, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *V0038JobResponseProperties) SetMinimumCpusPerNode(v int32)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *V0038JobResponseProperties) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMinimumTmpDiskPerNode

`func (o *V0038JobResponseProperties) GetMinimumTmpDiskPerNode() int32`

GetMinimumTmpDiskPerNode returns the MinimumTmpDiskPerNode field if non-nil, zero value otherwise.

### GetMinimumTmpDiskPerNodeOk

`func (o *V0038JobResponseProperties) GetMinimumTmpDiskPerNodeOk() (*int32, bool)`

GetMinimumTmpDiskPerNodeOk returns a tuple with the MinimumTmpDiskPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTmpDiskPerNode

`func (o *V0038JobResponseProperties) SetMinimumTmpDiskPerNode(v int32)`

SetMinimumTmpDiskPerNode sets MinimumTmpDiskPerNode field to given value.

### HasMinimumTmpDiskPerNode

`func (o *V0038JobResponseProperties) HasMinimumTmpDiskPerNode() bool`

HasMinimumTmpDiskPerNode returns a boolean if a field has been set.

### GetPreemptTime

`func (o *V0038JobResponseProperties) GetPreemptTime() int64`

GetPreemptTime returns the PreemptTime field if non-nil, zero value otherwise.

### GetPreemptTimeOk

`func (o *V0038JobResponseProperties) GetPreemptTimeOk() (*int64, bool)`

GetPreemptTimeOk returns a tuple with the PreemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptTime

`func (o *V0038JobResponseProperties) SetPreemptTime(v int64)`

SetPreemptTime sets PreemptTime field to given value.

### HasPreemptTime

`func (o *V0038JobResponseProperties) HasPreemptTime() bool`

HasPreemptTime returns a boolean if a field has been set.

### GetPreSusTime

`func (o *V0038JobResponseProperties) GetPreSusTime() int64`

GetPreSusTime returns the PreSusTime field if non-nil, zero value otherwise.

### GetPreSusTimeOk

`func (o *V0038JobResponseProperties) GetPreSusTimeOk() (*int64, bool)`

GetPreSusTimeOk returns a tuple with the PreSusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSusTime

`func (o *V0038JobResponseProperties) SetPreSusTime(v int64)`

SetPreSusTime sets PreSusTime field to given value.

### HasPreSusTime

`func (o *V0038JobResponseProperties) HasPreSusTime() bool`

HasPreSusTime returns a boolean if a field has been set.

### GetPriority

`func (o *V0038JobResponseProperties) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0038JobResponseProperties) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0038JobResponseProperties) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0038JobResponseProperties) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfile

`func (o *V0038JobResponseProperties) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *V0038JobResponseProperties) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *V0038JobResponseProperties) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *V0038JobResponseProperties) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQos

`func (o *V0038JobResponseProperties) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0038JobResponseProperties) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0038JobResponseProperties) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0038JobResponseProperties) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReboot

`func (o *V0038JobResponseProperties) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *V0038JobResponseProperties) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *V0038JobResponseProperties) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *V0038JobResponseProperties) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRequiredNodes

`func (o *V0038JobResponseProperties) GetRequiredNodes() string`

GetRequiredNodes returns the RequiredNodes field if non-nil, zero value otherwise.

### GetRequiredNodesOk

`func (o *V0038JobResponseProperties) GetRequiredNodesOk() (*string, bool)`

GetRequiredNodesOk returns a tuple with the RequiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNodes

`func (o *V0038JobResponseProperties) SetRequiredNodes(v string)`

SetRequiredNodes sets RequiredNodes field to given value.

### HasRequiredNodes

`func (o *V0038JobResponseProperties) HasRequiredNodes() bool`

HasRequiredNodes returns a boolean if a field has been set.

### GetRequeue

`func (o *V0038JobResponseProperties) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *V0038JobResponseProperties) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *V0038JobResponseProperties) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *V0038JobResponseProperties) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetResizeTime

`func (o *V0038JobResponseProperties) GetResizeTime() int64`

GetResizeTime returns the ResizeTime field if non-nil, zero value otherwise.

### GetResizeTimeOk

`func (o *V0038JobResponseProperties) GetResizeTimeOk() (*int64, bool)`

GetResizeTimeOk returns a tuple with the ResizeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeTime

`func (o *V0038JobResponseProperties) SetResizeTime(v int64)`

SetResizeTime sets ResizeTime field to given value.

### HasResizeTime

`func (o *V0038JobResponseProperties) HasResizeTime() bool`

HasResizeTime returns a boolean if a field has been set.

### GetRestartCnt

`func (o *V0038JobResponseProperties) GetRestartCnt() int32`

GetRestartCnt returns the RestartCnt field if non-nil, zero value otherwise.

### GetRestartCntOk

`func (o *V0038JobResponseProperties) GetRestartCntOk() (*int32, bool)`

GetRestartCntOk returns a tuple with the RestartCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCnt

`func (o *V0038JobResponseProperties) SetRestartCnt(v int32)`

SetRestartCnt sets RestartCnt field to given value.

### HasRestartCnt

`func (o *V0038JobResponseProperties) HasRestartCnt() bool`

HasRestartCnt returns a boolean if a field has been set.

### GetResvName

`func (o *V0038JobResponseProperties) GetResvName() string`

GetResvName returns the ResvName field if non-nil, zero value otherwise.

### GetResvNameOk

`func (o *V0038JobResponseProperties) GetResvNameOk() (*string, bool)`

GetResvNameOk returns a tuple with the ResvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResvName

`func (o *V0038JobResponseProperties) SetResvName(v string)`

SetResvName sets ResvName field to given value.

### HasResvName

`func (o *V0038JobResponseProperties) HasResvName() bool`

HasResvName returns a boolean if a field has been set.

### GetShared

`func (o *V0038JobResponseProperties) GetShared() string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *V0038JobResponseProperties) GetSharedOk() (*string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *V0038JobResponseProperties) SetShared(v string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *V0038JobResponseProperties) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetShowFlags

`func (o *V0038JobResponseProperties) GetShowFlags() []string`

GetShowFlags returns the ShowFlags field if non-nil, zero value otherwise.

### GetShowFlagsOk

`func (o *V0038JobResponseProperties) GetShowFlagsOk() (*[]string, bool)`

GetShowFlagsOk returns a tuple with the ShowFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFlags

`func (o *V0038JobResponseProperties) SetShowFlags(v []string)`

SetShowFlags sets ShowFlags field to given value.

### HasShowFlags

`func (o *V0038JobResponseProperties) HasShowFlags() bool`

HasShowFlags returns a boolean if a field has been set.

### GetSocketsPerBoard

`func (o *V0038JobResponseProperties) GetSocketsPerBoard() int32`

GetSocketsPerBoard returns the SocketsPerBoard field if non-nil, zero value otherwise.

### GetSocketsPerBoardOk

`func (o *V0038JobResponseProperties) GetSocketsPerBoardOk() (*int32, bool)`

GetSocketsPerBoardOk returns a tuple with the SocketsPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerBoard

`func (o *V0038JobResponseProperties) SetSocketsPerBoard(v int32)`

SetSocketsPerBoard sets SocketsPerBoard field to given value.

### HasSocketsPerBoard

`func (o *V0038JobResponseProperties) HasSocketsPerBoard() bool`

HasSocketsPerBoard returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *V0038JobResponseProperties) GetSocketsPerNode() int32`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0038JobResponseProperties) GetSocketsPerNodeOk() (*int32, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0038JobResponseProperties) SetSocketsPerNode(v int32)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0038JobResponseProperties) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetStartTime

`func (o *V0038JobResponseProperties) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0038JobResponseProperties) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0038JobResponseProperties) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0038JobResponseProperties) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStateDescription

`func (o *V0038JobResponseProperties) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *V0038JobResponseProperties) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *V0038JobResponseProperties) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *V0038JobResponseProperties) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetStateReason

`func (o *V0038JobResponseProperties) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *V0038JobResponseProperties) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *V0038JobResponseProperties) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *V0038JobResponseProperties) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStandardError

`func (o *V0038JobResponseProperties) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *V0038JobResponseProperties) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *V0038JobResponseProperties) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *V0038JobResponseProperties) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStandardInput

`func (o *V0038JobResponseProperties) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *V0038JobResponseProperties) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *V0038JobResponseProperties) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *V0038JobResponseProperties) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *V0038JobResponseProperties) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *V0038JobResponseProperties) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *V0038JobResponseProperties) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *V0038JobResponseProperties) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetSubmitTime

`func (o *V0038JobResponseProperties) GetSubmitTime() int64`

GetSubmitTime returns the SubmitTime field if non-nil, zero value otherwise.

### GetSubmitTimeOk

`func (o *V0038JobResponseProperties) GetSubmitTimeOk() (*int64, bool)`

GetSubmitTimeOk returns a tuple with the SubmitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitTime

`func (o *V0038JobResponseProperties) SetSubmitTime(v int64)`

SetSubmitTime sets SubmitTime field to given value.

### HasSubmitTime

`func (o *V0038JobResponseProperties) HasSubmitTime() bool`

HasSubmitTime returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0038JobResponseProperties) GetSuspendTime() int64`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0038JobResponseProperties) GetSuspendTimeOk() (*int64, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0038JobResponseProperties) SetSuspendTime(v int64)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0038JobResponseProperties) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.

### GetSystemComment

`func (o *V0038JobResponseProperties) GetSystemComment() string`

GetSystemComment returns the SystemComment field if non-nil, zero value otherwise.

### GetSystemCommentOk

`func (o *V0038JobResponseProperties) GetSystemCommentOk() (*string, bool)`

GetSystemCommentOk returns a tuple with the SystemComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemComment

`func (o *V0038JobResponseProperties) SetSystemComment(v string)`

SetSystemComment sets SystemComment field to given value.

### HasSystemComment

`func (o *V0038JobResponseProperties) HasSystemComment() bool`

HasSystemComment returns a boolean if a field has been set.

### GetTimeLimit

`func (o *V0038JobResponseProperties) GetTimeLimit() int64`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *V0038JobResponseProperties) GetTimeLimitOk() (*int64, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *V0038JobResponseProperties) SetTimeLimit(v int64)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *V0038JobResponseProperties) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *V0038JobResponseProperties) GetTimeMinimum() int64`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *V0038JobResponseProperties) GetTimeMinimumOk() (*int64, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *V0038JobResponseProperties) SetTimeMinimum(v int64)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *V0038JobResponseProperties) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *V0038JobResponseProperties) GetThreadsPerCore() int32`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0038JobResponseProperties) GetThreadsPerCoreOk() (*int32, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0038JobResponseProperties) SetThreadsPerCore(v int32)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *V0038JobResponseProperties) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTresBind

`func (o *V0038JobResponseProperties) GetTresBind() string`

GetTresBind returns the TresBind field if non-nil, zero value otherwise.

### GetTresBindOk

`func (o *V0038JobResponseProperties) GetTresBindOk() (*string, bool)`

GetTresBindOk returns a tuple with the TresBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresBind

`func (o *V0038JobResponseProperties) SetTresBind(v string)`

SetTresBind sets TresBind field to given value.

### HasTresBind

`func (o *V0038JobResponseProperties) HasTresBind() bool`

HasTresBind returns a boolean if a field has been set.

### GetTresFreq

`func (o *V0038JobResponseProperties) GetTresFreq() string`

GetTresFreq returns the TresFreq field if non-nil, zero value otherwise.

### GetTresFreqOk

`func (o *V0038JobResponseProperties) GetTresFreqOk() (*string, bool)`

GetTresFreqOk returns a tuple with the TresFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresFreq

`func (o *V0038JobResponseProperties) SetTresFreq(v string)`

SetTresFreq sets TresFreq field to given value.

### HasTresFreq

`func (o *V0038JobResponseProperties) HasTresFreq() bool`

HasTresFreq returns a boolean if a field has been set.

### GetTresPerJob

`func (o *V0038JobResponseProperties) GetTresPerJob() string`

GetTresPerJob returns the TresPerJob field if non-nil, zero value otherwise.

### GetTresPerJobOk

`func (o *V0038JobResponseProperties) GetTresPerJobOk() (*string, bool)`

GetTresPerJobOk returns a tuple with the TresPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerJob

`func (o *V0038JobResponseProperties) SetTresPerJob(v string)`

SetTresPerJob sets TresPerJob field to given value.

### HasTresPerJob

`func (o *V0038JobResponseProperties) HasTresPerJob() bool`

HasTresPerJob returns a boolean if a field has been set.

### GetTresPerNode

`func (o *V0038JobResponseProperties) GetTresPerNode() string`

GetTresPerNode returns the TresPerNode field if non-nil, zero value otherwise.

### GetTresPerNodeOk

`func (o *V0038JobResponseProperties) GetTresPerNodeOk() (*string, bool)`

GetTresPerNodeOk returns a tuple with the TresPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerNode

`func (o *V0038JobResponseProperties) SetTresPerNode(v string)`

SetTresPerNode sets TresPerNode field to given value.

### HasTresPerNode

`func (o *V0038JobResponseProperties) HasTresPerNode() bool`

HasTresPerNode returns a boolean if a field has been set.

### GetTresPerSocket

`func (o *V0038JobResponseProperties) GetTresPerSocket() string`

GetTresPerSocket returns the TresPerSocket field if non-nil, zero value otherwise.

### GetTresPerSocketOk

`func (o *V0038JobResponseProperties) GetTresPerSocketOk() (*string, bool)`

GetTresPerSocketOk returns a tuple with the TresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerSocket

`func (o *V0038JobResponseProperties) SetTresPerSocket(v string)`

SetTresPerSocket sets TresPerSocket field to given value.

### HasTresPerSocket

`func (o *V0038JobResponseProperties) HasTresPerSocket() bool`

HasTresPerSocket returns a boolean if a field has been set.

### GetTresPerTask

`func (o *V0038JobResponseProperties) GetTresPerTask() string`

GetTresPerTask returns the TresPerTask field if non-nil, zero value otherwise.

### GetTresPerTaskOk

`func (o *V0038JobResponseProperties) GetTresPerTaskOk() (*string, bool)`

GetTresPerTaskOk returns a tuple with the TresPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerTask

`func (o *V0038JobResponseProperties) SetTresPerTask(v string)`

SetTresPerTask sets TresPerTask field to given value.

### HasTresPerTask

`func (o *V0038JobResponseProperties) HasTresPerTask() bool`

HasTresPerTask returns a boolean if a field has been set.

### GetTresReqStr

`func (o *V0038JobResponseProperties) GetTresReqStr() string`

GetTresReqStr returns the TresReqStr field if non-nil, zero value otherwise.

### GetTresReqStrOk

`func (o *V0038JobResponseProperties) GetTresReqStrOk() (*string, bool)`

GetTresReqStrOk returns a tuple with the TresReqStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresReqStr

`func (o *V0038JobResponseProperties) SetTresReqStr(v string)`

SetTresReqStr sets TresReqStr field to given value.

### HasTresReqStr

`func (o *V0038JobResponseProperties) HasTresReqStr() bool`

HasTresReqStr returns a boolean if a field has been set.

### GetTresAllocStr

`func (o *V0038JobResponseProperties) GetTresAllocStr() string`

GetTresAllocStr returns the TresAllocStr field if non-nil, zero value otherwise.

### GetTresAllocStrOk

`func (o *V0038JobResponseProperties) GetTresAllocStrOk() (*string, bool)`

GetTresAllocStrOk returns a tuple with the TresAllocStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresAllocStr

`func (o *V0038JobResponseProperties) SetTresAllocStr(v string)`

SetTresAllocStr sets TresAllocStr field to given value.

### HasTresAllocStr

`func (o *V0038JobResponseProperties) HasTresAllocStr() bool`

HasTresAllocStr returns a boolean if a field has been set.

### GetUserId

`func (o *V0038JobResponseProperties) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0038JobResponseProperties) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0038JobResponseProperties) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0038JobResponseProperties) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *V0038JobResponseProperties) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *V0038JobResponseProperties) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *V0038JobResponseProperties) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *V0038JobResponseProperties) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWckey

`func (o *V0038JobResponseProperties) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0038JobResponseProperties) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0038JobResponseProperties) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0038JobResponseProperties) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *V0038JobResponseProperties) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *V0038JobResponseProperties) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *V0038JobResponseProperties) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *V0038JobResponseProperties) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


