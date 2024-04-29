# V0038JobProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Charge resources used by this job to specified account. | [optional] 
**AccountGatherFrequency** | Pointer to **string** | Define the job accounting and profiling sampling intervals. | [optional] 
**Argv** | Pointer to **[]string** | Arguments to the script. | [optional] 
**Array** | Pointer to **string** | Submit a job array, multiple jobs to be executed with identical parameters. The indexes specification identifies what array index values should be used. | [optional] 
**BatchFeatures** | Pointer to **string** | features required for batch script&#39;s node | [optional] 
**BeginTime** | Pointer to **int64** | Submit the batch script to the Slurm controller immediately, like normal, but tell the controller to defer the allocation of the job until the specified time. | [optional] 
**BurstBuffer** | Pointer to **string** | Burst buffer specification. | [optional] 
**ClusterConstraint** | Pointer to **string** | Specifies features that a federated cluster must have to have a sibling job submitted to it. | [optional] 
**Comment** | Pointer to **string** | An arbitrary comment. | [optional] 
**Constraints** | Pointer to **string** | node features required by job. | [optional] 
**Container** | Pointer to **string** | absolute path to OCI container bundle | [optional] 
**CoreSpecification** | Pointer to **int32** | Count of specialized threads per node reserved by the job for system operations and not used by the application. | [optional] 
**CoresPerSocket** | Pointer to **int32** | Restrict node selection to nodes with at least the specified number of cores per socket. | [optional] 
**CpuBinding** | Pointer to **string** | Cpu binding | [optional] 
**CpuBindingHint** | Pointer to **string** | Cpu binding hint | [optional] 
**CpuFrequency** | Pointer to **string** | Request that job steps initiated by srun commands inside this sbatch script be run at some requested frequency if possible, on the CPUs selected for the step on the compute node(s). | [optional] 
**CpusPerGpu** | Pointer to **string** | Number of CPUs requested per allocated GPU. | [optional] 
**CpusPerTask** | Pointer to **int32** | Advise the Slurm controller that ensuing job steps will require ncpus number of processors per task. | [optional] 
**CurrentWorkingDirectory** | Pointer to **string** | Instruct Slurm to connect the batch script&#39;s standard output directly to the file name. | [optional] 
**Deadline** | Pointer to **string** | Remove the job if no ending is possible before this deadline (start &gt; (deadline - time[-min])). | [optional] 
**DelayBoot** | Pointer to **int32** | Do not reboot nodes in order to satisfied this job&#39;s feature specification if the job has been eligible to run for less than this time period. | [optional] 
**Dependency** | Pointer to **string** | Defer the start of this job until the specified dependencies have been satisfied completed. | [optional] 
**Distribution** | Pointer to **string** | Specify alternate distribution methods for remote processes. | [optional] 
**Environment** | **map[string]interface{}** | Dictionary of environment entries. | 
**Exclusive** | Pointer to **string** | The job allocation can share nodes just other users with the \&quot;user\&quot; option or with the \&quot;mcs\&quot; option). | [optional] 
**GetUserEnvironment** | Pointer to **bool** | Load new login environment for user on job node. | [optional] 
**Gres** | Pointer to **string** | Specifies a comma delimited list of generic consumable resources. | [optional] 
**GresFlags** | Pointer to **string** | Specify generic resource task binding options. | [optional] 
**GpuBinding** | Pointer to **string** | Requested binding of tasks to GPU. | [optional] 
**GpuFrequency** | Pointer to **string** | Requested GPU frequency. | [optional] 
**Gpus** | Pointer to **string** | GPUs per job. | [optional] 
**GpusPerNode** | Pointer to **string** | GPUs per node. | [optional] 
**GpusPerSocket** | Pointer to **string** | GPUs per socket. | [optional] 
**GpusPerTask** | Pointer to **string** | GPUs per task. | [optional] 
**Hold** | Pointer to **bool** | Specify the job is to be submitted in a held state (priority of zero). | [optional] 
**KillOnInvalidDependency** | Pointer to **bool** | If a job has an invalid dependency, then Slurm is to terminate it. | [optional] 
**Licenses** | Pointer to **string** | Specification of licenses (or other resources available on all nodes of the cluster) which must be allocated to this job. | [optional] 
**MailType** | Pointer to **string** | Notify user by email when certain event types occur. | [optional] 
**MailUser** | Pointer to **string** | User to receive email notification of state changes as defined by mail_type. | [optional] 
**McsLabel** | Pointer to **string** | This parameter is a group among the groups of the user. | [optional] 
**MemoryBinding** | Pointer to **string** | Bind tasks to memory. | [optional] 
**MemoryPerCpu** | Pointer to **int32** | Minimum real memory per cpu (MB). | [optional] 
**MemoryPerGpu** | Pointer to **int32** | Minimum memory required per allocated GPU. | [optional] 
**MemoryPerNode** | Pointer to **int32** | Minimum real memory per node (MB). | [optional] 
**MinimumCpusPerNode** | Pointer to **int32** | Minimum number of CPUs per node. | [optional] 
**MinimumNodes** | Pointer to **bool** | If a range of node counts is given, prefer the smaller count. | [optional] 
**Name** | Pointer to **string** | Specify a name for the job allocation. | [optional] 
**Nice** | Pointer to **int32** | Run the job with an adjusted scheduling priority within Slurm. | [optional] 
**NoKill** | Pointer to **bool** | Do not automatically terminate a job if one of the nodes it has been allocated fails. | [optional] 
**Nodes** | Pointer to **[]int32** | Request that a minimum of minnodes nodes and a maximum node count. | [optional] 
**OpenMode** | Pointer to **string** | Open the output and error files using append or truncate mode as specified. | [optional] [default to "append"]
**Oversubscribe** | Pointer to **bool** | The job allocation can over-subscribe resources with other running jobs. | [optional] [default to false]
**Partition** | Pointer to **string** | Request a specific partition for the resource allocation. | [optional] 
**Prefer** | Pointer to **string** | Comma delimited list of features for scheduler to prefer but not a strict requirement like a constraint. Value can be used for job submission but is only displayed for PENDING jobs. | [optional] 
**Priority** | Pointer to **string** | Request a specific job priority. | [optional] 
**Qos** | Pointer to **string** | Request a quality of service for the job. | [optional] 
**Requeue** | Pointer to **bool** | Specifies that the batch job should eligible to being requeue. | [optional] 
**Reservation** | Pointer to **string** | Allocate resources for the job from the named reservation. | [optional] 
**Signal** | Pointer to **string** | When a job is within sig_time seconds of its end time, send it the signal sig_num. | [optional] 
**SocketsPerNode** | Pointer to **int32** | Restrict node selection to nodes with at least the specified number of sockets. | [optional] 
**SpreadJob** | Pointer to **bool** | Spread the job allocation over as many nodes as possible and attempt to evenly distribute tasks across the allocated nodes. | [optional] 
**StandardError** | Pointer to **string** | Instruct Slurm to connect the batch script&#39;s standard error directly to the file name. | [optional] 
**StandardInput** | Pointer to **string** | Instruct Slurm to connect the batch script&#39;s standard input directly to the file name specified. | [optional] 
**StandardOutput** | Pointer to **string** | Instruct Slurm to connect the batch script&#39;s standard output directly to the file name. | [optional] 
**Tasks** | Pointer to **int32** | Advises the Slurm controller that job steps run within the allocation will launch a maximum of number tasks and to provide for sufficient resources. | [optional] 
**TasksPerCore** | Pointer to **int32** | Request the maximum ntasks be invoked on each core. | [optional] 
**TasksPerNode** | Pointer to **int32** | Request the maximum ntasks be invoked on each node. | [optional] 
**TasksPerSocket** | Pointer to **int32** | Request the maximum ntasks be invoked on each socket. | [optional] 
**ThreadSpecification** | Pointer to **int32** | Count of specialized threads per node reserved by the job for system operations and not used by the application. | [optional] 
**ThreadsPerCore** | Pointer to **int32** | Restrict node selection to nodes with at least the specified number of threads per core. | [optional] 
**TimeLimit** | Pointer to **int32** | Step time limit in minutes. | [optional] 
**TimeMinimum** | Pointer to **int32** | Minimum run time in minutes. | [optional] 
**WaitAllNodes** | Pointer to **bool** | Do not begin execution until all nodes are ready for use. | [optional] 
**Wckey** | Pointer to **string** | Specify wckey to be used with job. | [optional] 

## Methods

### NewV0038JobProperties

`func NewV0038JobProperties(environment map[string]interface{}, ) *V0038JobProperties`

NewV0038JobProperties instantiates a new V0038JobProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038JobPropertiesWithDefaults

`func NewV0038JobPropertiesWithDefaults() *V0038JobProperties`

NewV0038JobPropertiesWithDefaults instantiates a new V0038JobProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0038JobProperties) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0038JobProperties) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0038JobProperties) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0038JobProperties) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountGatherFrequency

`func (o *V0038JobProperties) GetAccountGatherFrequency() string`

GetAccountGatherFrequency returns the AccountGatherFrequency field if non-nil, zero value otherwise.

### GetAccountGatherFrequencyOk

`func (o *V0038JobProperties) GetAccountGatherFrequencyOk() (*string, bool)`

GetAccountGatherFrequencyOk returns a tuple with the AccountGatherFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGatherFrequency

`func (o *V0038JobProperties) SetAccountGatherFrequency(v string)`

SetAccountGatherFrequency sets AccountGatherFrequency field to given value.

### HasAccountGatherFrequency

`func (o *V0038JobProperties) HasAccountGatherFrequency() bool`

HasAccountGatherFrequency returns a boolean if a field has been set.

### GetArgv

`func (o *V0038JobProperties) GetArgv() []string`

GetArgv returns the Argv field if non-nil, zero value otherwise.

### GetArgvOk

`func (o *V0038JobProperties) GetArgvOk() (*[]string, bool)`

GetArgvOk returns a tuple with the Argv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgv

`func (o *V0038JobProperties) SetArgv(v []string)`

SetArgv sets Argv field to given value.

### HasArgv

`func (o *V0038JobProperties) HasArgv() bool`

HasArgv returns a boolean if a field has been set.

### GetArray

`func (o *V0038JobProperties) GetArray() string`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *V0038JobProperties) GetArrayOk() (*string, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *V0038JobProperties) SetArray(v string)`

SetArray sets Array field to given value.

### HasArray

`func (o *V0038JobProperties) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *V0038JobProperties) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *V0038JobProperties) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *V0038JobProperties) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *V0038JobProperties) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBeginTime

`func (o *V0038JobProperties) GetBeginTime() int64`

GetBeginTime returns the BeginTime field if non-nil, zero value otherwise.

### GetBeginTimeOk

`func (o *V0038JobProperties) GetBeginTimeOk() (*int64, bool)`

GetBeginTimeOk returns a tuple with the BeginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTime

`func (o *V0038JobProperties) SetBeginTime(v int64)`

SetBeginTime sets BeginTime field to given value.

### HasBeginTime

`func (o *V0038JobProperties) HasBeginTime() bool`

HasBeginTime returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0038JobProperties) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0038JobProperties) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0038JobProperties) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0038JobProperties) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetClusterConstraint

`func (o *V0038JobProperties) GetClusterConstraint() string`

GetClusterConstraint returns the ClusterConstraint field if non-nil, zero value otherwise.

### GetClusterConstraintOk

`func (o *V0038JobProperties) GetClusterConstraintOk() (*string, bool)`

GetClusterConstraintOk returns a tuple with the ClusterConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConstraint

`func (o *V0038JobProperties) SetClusterConstraint(v string)`

SetClusterConstraint sets ClusterConstraint field to given value.

### HasClusterConstraint

`func (o *V0038JobProperties) HasClusterConstraint() bool`

HasClusterConstraint returns a boolean if a field has been set.

### GetComment

`func (o *V0038JobProperties) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0038JobProperties) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0038JobProperties) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0038JobProperties) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConstraints

`func (o *V0038JobProperties) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *V0038JobProperties) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *V0038JobProperties) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *V0038JobProperties) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetContainer

`func (o *V0038JobProperties) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0038JobProperties) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0038JobProperties) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0038JobProperties) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCoreSpecification

`func (o *V0038JobProperties) GetCoreSpecification() int32`

GetCoreSpecification returns the CoreSpecification field if non-nil, zero value otherwise.

### GetCoreSpecificationOk

`func (o *V0038JobProperties) GetCoreSpecificationOk() (*int32, bool)`

GetCoreSpecificationOk returns a tuple with the CoreSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpecification

`func (o *V0038JobProperties) SetCoreSpecification(v int32)`

SetCoreSpecification sets CoreSpecification field to given value.

### HasCoreSpecification

`func (o *V0038JobProperties) HasCoreSpecification() bool`

HasCoreSpecification returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *V0038JobProperties) GetCoresPerSocket() int32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *V0038JobProperties) GetCoresPerSocketOk() (*int32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *V0038JobProperties) SetCoresPerSocket(v int32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *V0038JobProperties) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCpuBinding

`func (o *V0038JobProperties) GetCpuBinding() string`

GetCpuBinding returns the CpuBinding field if non-nil, zero value otherwise.

### GetCpuBindingOk

`func (o *V0038JobProperties) GetCpuBindingOk() (*string, bool)`

GetCpuBindingOk returns a tuple with the CpuBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBinding

`func (o *V0038JobProperties) SetCpuBinding(v string)`

SetCpuBinding sets CpuBinding field to given value.

### HasCpuBinding

`func (o *V0038JobProperties) HasCpuBinding() bool`

HasCpuBinding returns a boolean if a field has been set.

### GetCpuBindingHint

`func (o *V0038JobProperties) GetCpuBindingHint() string`

GetCpuBindingHint returns the CpuBindingHint field if non-nil, zero value otherwise.

### GetCpuBindingHintOk

`func (o *V0038JobProperties) GetCpuBindingHintOk() (*string, bool)`

GetCpuBindingHintOk returns a tuple with the CpuBindingHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBindingHint

`func (o *V0038JobProperties) SetCpuBindingHint(v string)`

SetCpuBindingHint sets CpuBindingHint field to given value.

### HasCpuBindingHint

`func (o *V0038JobProperties) HasCpuBindingHint() bool`

HasCpuBindingHint returns a boolean if a field has been set.

### GetCpuFrequency

`func (o *V0038JobProperties) GetCpuFrequency() string`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *V0038JobProperties) GetCpuFrequencyOk() (*string, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *V0038JobProperties) SetCpuFrequency(v string)`

SetCpuFrequency sets CpuFrequency field to given value.

### HasCpuFrequency

`func (o *V0038JobProperties) HasCpuFrequency() bool`

HasCpuFrequency returns a boolean if a field has been set.

### GetCpusPerGpu

`func (o *V0038JobProperties) GetCpusPerGpu() string`

GetCpusPerGpu returns the CpusPerGpu field if non-nil, zero value otherwise.

### GetCpusPerGpuOk

`func (o *V0038JobProperties) GetCpusPerGpuOk() (*string, bool)`

GetCpusPerGpuOk returns a tuple with the CpusPerGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerGpu

`func (o *V0038JobProperties) SetCpusPerGpu(v string)`

SetCpusPerGpu sets CpusPerGpu field to given value.

### HasCpusPerGpu

`func (o *V0038JobProperties) HasCpusPerGpu() bool`

HasCpusPerGpu returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *V0038JobProperties) GetCpusPerTask() int32`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *V0038JobProperties) GetCpusPerTaskOk() (*int32, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *V0038JobProperties) SetCpusPerTask(v int32)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *V0038JobProperties) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *V0038JobProperties) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *V0038JobProperties) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *V0038JobProperties) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *V0038JobProperties) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.

### GetDeadline

`func (o *V0038JobProperties) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *V0038JobProperties) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *V0038JobProperties) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *V0038JobProperties) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *V0038JobProperties) GetDelayBoot() int32`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *V0038JobProperties) GetDelayBootOk() (*int32, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *V0038JobProperties) SetDelayBoot(v int32)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *V0038JobProperties) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *V0038JobProperties) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *V0038JobProperties) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *V0038JobProperties) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *V0038JobProperties) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetDistribution

`func (o *V0038JobProperties) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *V0038JobProperties) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *V0038JobProperties) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *V0038JobProperties) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetEnvironment

`func (o *V0038JobProperties) GetEnvironment() map[string]interface{}`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *V0038JobProperties) GetEnvironmentOk() (*map[string]interface{}, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *V0038JobProperties) SetEnvironment(v map[string]interface{})`

SetEnvironment sets Environment field to given value.


### GetExclusive

`func (o *V0038JobProperties) GetExclusive() string`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *V0038JobProperties) GetExclusiveOk() (*string, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *V0038JobProperties) SetExclusive(v string)`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *V0038JobProperties) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### GetGetUserEnvironment

`func (o *V0038JobProperties) GetGetUserEnvironment() bool`

GetGetUserEnvironment returns the GetUserEnvironment field if non-nil, zero value otherwise.

### GetGetUserEnvironmentOk

`func (o *V0038JobProperties) GetGetUserEnvironmentOk() (*bool, bool)`

GetGetUserEnvironmentOk returns a tuple with the GetUserEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetUserEnvironment

`func (o *V0038JobProperties) SetGetUserEnvironment(v bool)`

SetGetUserEnvironment sets GetUserEnvironment field to given value.

### HasGetUserEnvironment

`func (o *V0038JobProperties) HasGetUserEnvironment() bool`

HasGetUserEnvironment returns a boolean if a field has been set.

### GetGres

`func (o *V0038JobProperties) GetGres() string`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *V0038JobProperties) GetGresOk() (*string, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *V0038JobProperties) SetGres(v string)`

SetGres sets Gres field to given value.

### HasGres

`func (o *V0038JobProperties) HasGres() bool`

HasGres returns a boolean if a field has been set.

### GetGresFlags

`func (o *V0038JobProperties) GetGresFlags() string`

GetGresFlags returns the GresFlags field if non-nil, zero value otherwise.

### GetGresFlagsOk

`func (o *V0038JobProperties) GetGresFlagsOk() (*string, bool)`

GetGresFlagsOk returns a tuple with the GresFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresFlags

`func (o *V0038JobProperties) SetGresFlags(v string)`

SetGresFlags sets GresFlags field to given value.

### HasGresFlags

`func (o *V0038JobProperties) HasGresFlags() bool`

HasGresFlags returns a boolean if a field has been set.

### GetGpuBinding

`func (o *V0038JobProperties) GetGpuBinding() string`

GetGpuBinding returns the GpuBinding field if non-nil, zero value otherwise.

### GetGpuBindingOk

`func (o *V0038JobProperties) GetGpuBindingOk() (*string, bool)`

GetGpuBindingOk returns a tuple with the GpuBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuBinding

`func (o *V0038JobProperties) SetGpuBinding(v string)`

SetGpuBinding sets GpuBinding field to given value.

### HasGpuBinding

`func (o *V0038JobProperties) HasGpuBinding() bool`

HasGpuBinding returns a boolean if a field has been set.

### GetGpuFrequency

`func (o *V0038JobProperties) GetGpuFrequency() string`

GetGpuFrequency returns the GpuFrequency field if non-nil, zero value otherwise.

### GetGpuFrequencyOk

`func (o *V0038JobProperties) GetGpuFrequencyOk() (*string, bool)`

GetGpuFrequencyOk returns a tuple with the GpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuFrequency

`func (o *V0038JobProperties) SetGpuFrequency(v string)`

SetGpuFrequency sets GpuFrequency field to given value.

### HasGpuFrequency

`func (o *V0038JobProperties) HasGpuFrequency() bool`

HasGpuFrequency returns a boolean if a field has been set.

### GetGpus

`func (o *V0038JobProperties) GetGpus() string`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *V0038JobProperties) GetGpusOk() (*string, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *V0038JobProperties) SetGpus(v string)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *V0038JobProperties) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetGpusPerNode

`func (o *V0038JobProperties) GetGpusPerNode() string`

GetGpusPerNode returns the GpusPerNode field if non-nil, zero value otherwise.

### GetGpusPerNodeOk

`func (o *V0038JobProperties) GetGpusPerNodeOk() (*string, bool)`

GetGpusPerNodeOk returns a tuple with the GpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpusPerNode

`func (o *V0038JobProperties) SetGpusPerNode(v string)`

SetGpusPerNode sets GpusPerNode field to given value.

### HasGpusPerNode

`func (o *V0038JobProperties) HasGpusPerNode() bool`

HasGpusPerNode returns a boolean if a field has been set.

### GetGpusPerSocket

`func (o *V0038JobProperties) GetGpusPerSocket() string`

GetGpusPerSocket returns the GpusPerSocket field if non-nil, zero value otherwise.

### GetGpusPerSocketOk

`func (o *V0038JobProperties) GetGpusPerSocketOk() (*string, bool)`

GetGpusPerSocketOk returns a tuple with the GpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpusPerSocket

`func (o *V0038JobProperties) SetGpusPerSocket(v string)`

SetGpusPerSocket sets GpusPerSocket field to given value.

### HasGpusPerSocket

`func (o *V0038JobProperties) HasGpusPerSocket() bool`

HasGpusPerSocket returns a boolean if a field has been set.

### GetGpusPerTask

`func (o *V0038JobProperties) GetGpusPerTask() string`

GetGpusPerTask returns the GpusPerTask field if non-nil, zero value otherwise.

### GetGpusPerTaskOk

`func (o *V0038JobProperties) GetGpusPerTaskOk() (*string, bool)`

GetGpusPerTaskOk returns a tuple with the GpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpusPerTask

`func (o *V0038JobProperties) SetGpusPerTask(v string)`

SetGpusPerTask sets GpusPerTask field to given value.

### HasGpusPerTask

`func (o *V0038JobProperties) HasGpusPerTask() bool`

HasGpusPerTask returns a boolean if a field has been set.

### GetHold

`func (o *V0038JobProperties) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *V0038JobProperties) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *V0038JobProperties) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *V0038JobProperties) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetKillOnInvalidDependency

`func (o *V0038JobProperties) GetKillOnInvalidDependency() bool`

GetKillOnInvalidDependency returns the KillOnInvalidDependency field if non-nil, zero value otherwise.

### GetKillOnInvalidDependencyOk

`func (o *V0038JobProperties) GetKillOnInvalidDependencyOk() (*bool, bool)`

GetKillOnInvalidDependencyOk returns a tuple with the KillOnInvalidDependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillOnInvalidDependency

`func (o *V0038JobProperties) SetKillOnInvalidDependency(v bool)`

SetKillOnInvalidDependency sets KillOnInvalidDependency field to given value.

### HasKillOnInvalidDependency

`func (o *V0038JobProperties) HasKillOnInvalidDependency() bool`

HasKillOnInvalidDependency returns a boolean if a field has been set.

### GetLicenses

`func (o *V0038JobProperties) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0038JobProperties) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0038JobProperties) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0038JobProperties) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMailType

`func (o *V0038JobProperties) GetMailType() string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *V0038JobProperties) GetMailTypeOk() (*string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *V0038JobProperties) SetMailType(v string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *V0038JobProperties) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMailUser

`func (o *V0038JobProperties) GetMailUser() string`

GetMailUser returns the MailUser field if non-nil, zero value otherwise.

### GetMailUserOk

`func (o *V0038JobProperties) GetMailUserOk() (*string, bool)`

GetMailUserOk returns a tuple with the MailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUser

`func (o *V0038JobProperties) SetMailUser(v string)`

SetMailUser sets MailUser field to given value.

### HasMailUser

`func (o *V0038JobProperties) HasMailUser() bool`

HasMailUser returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0038JobProperties) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0038JobProperties) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0038JobProperties) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0038JobProperties) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryBinding

`func (o *V0038JobProperties) GetMemoryBinding() string`

GetMemoryBinding returns the MemoryBinding field if non-nil, zero value otherwise.

### GetMemoryBindingOk

`func (o *V0038JobProperties) GetMemoryBindingOk() (*string, bool)`

GetMemoryBindingOk returns a tuple with the MemoryBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBinding

`func (o *V0038JobProperties) SetMemoryBinding(v string)`

SetMemoryBinding sets MemoryBinding field to given value.

### HasMemoryBinding

`func (o *V0038JobProperties) HasMemoryBinding() bool`

HasMemoryBinding returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0038JobProperties) GetMemoryPerCpu() int32`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0038JobProperties) GetMemoryPerCpuOk() (*int32, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0038JobProperties) SetMemoryPerCpu(v int32)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0038JobProperties) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerGpu

`func (o *V0038JobProperties) GetMemoryPerGpu() int32`

GetMemoryPerGpu returns the MemoryPerGpu field if non-nil, zero value otherwise.

### GetMemoryPerGpuOk

`func (o *V0038JobProperties) GetMemoryPerGpuOk() (*int32, bool)`

GetMemoryPerGpuOk returns a tuple with the MemoryPerGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerGpu

`func (o *V0038JobProperties) SetMemoryPerGpu(v int32)`

SetMemoryPerGpu sets MemoryPerGpu field to given value.

### HasMemoryPerGpu

`func (o *V0038JobProperties) HasMemoryPerGpu() bool`

HasMemoryPerGpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0038JobProperties) GetMemoryPerNode() int32`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0038JobProperties) GetMemoryPerNodeOk() (*int32, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0038JobProperties) SetMemoryPerNode(v int32)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0038JobProperties) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *V0038JobProperties) GetMinimumCpusPerNode() int32`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *V0038JobProperties) GetMinimumCpusPerNodeOk() (*int32, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *V0038JobProperties) SetMinimumCpusPerNode(v int32)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *V0038JobProperties) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMinimumNodes

`func (o *V0038JobProperties) GetMinimumNodes() bool`

GetMinimumNodes returns the MinimumNodes field if non-nil, zero value otherwise.

### GetMinimumNodesOk

`func (o *V0038JobProperties) GetMinimumNodesOk() (*bool, bool)`

GetMinimumNodesOk returns a tuple with the MinimumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNodes

`func (o *V0038JobProperties) SetMinimumNodes(v bool)`

SetMinimumNodes sets MinimumNodes field to given value.

### HasMinimumNodes

`func (o *V0038JobProperties) HasMinimumNodes() bool`

HasMinimumNodes returns a boolean if a field has been set.

### GetName

`func (o *V0038JobProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0038JobProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0038JobProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0038JobProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNice

`func (o *V0038JobProperties) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *V0038JobProperties) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *V0038JobProperties) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *V0038JobProperties) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetNoKill

`func (o *V0038JobProperties) GetNoKill() bool`

GetNoKill returns the NoKill field if non-nil, zero value otherwise.

### GetNoKillOk

`func (o *V0038JobProperties) GetNoKillOk() (*bool, bool)`

GetNoKillOk returns a tuple with the NoKill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoKill

`func (o *V0038JobProperties) SetNoKill(v bool)`

SetNoKill sets NoKill field to given value.

### HasNoKill

`func (o *V0038JobProperties) HasNoKill() bool`

HasNoKill returns a boolean if a field has been set.

### GetNodes

`func (o *V0038JobProperties) GetNodes() []int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0038JobProperties) GetNodesOk() (*[]int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0038JobProperties) SetNodes(v []int32)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0038JobProperties) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOpenMode

`func (o *V0038JobProperties) GetOpenMode() string`

GetOpenMode returns the OpenMode field if non-nil, zero value otherwise.

### GetOpenModeOk

`func (o *V0038JobProperties) GetOpenModeOk() (*string, bool)`

GetOpenModeOk returns a tuple with the OpenMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMode

`func (o *V0038JobProperties) SetOpenMode(v string)`

SetOpenMode sets OpenMode field to given value.

### HasOpenMode

`func (o *V0038JobProperties) HasOpenMode() bool`

HasOpenMode returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0038JobProperties) GetOversubscribe() bool`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0038JobProperties) GetOversubscribeOk() (*bool, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0038JobProperties) SetOversubscribe(v bool)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0038JobProperties) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetPartition

`func (o *V0038JobProperties) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0038JobProperties) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0038JobProperties) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0038JobProperties) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPrefer

`func (o *V0038JobProperties) GetPrefer() string`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *V0038JobProperties) GetPreferOk() (*string, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *V0038JobProperties) SetPrefer(v string)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *V0038JobProperties) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetPriority

`func (o *V0038JobProperties) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0038JobProperties) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0038JobProperties) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0038JobProperties) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *V0038JobProperties) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0038JobProperties) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0038JobProperties) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0038JobProperties) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetRequeue

`func (o *V0038JobProperties) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *V0038JobProperties) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *V0038JobProperties) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *V0038JobProperties) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetReservation

`func (o *V0038JobProperties) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *V0038JobProperties) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *V0038JobProperties) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *V0038JobProperties) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetSignal

`func (o *V0038JobProperties) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *V0038JobProperties) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *V0038JobProperties) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *V0038JobProperties) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *V0038JobProperties) GetSocketsPerNode() int32`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0038JobProperties) GetSocketsPerNodeOk() (*int32, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0038JobProperties) SetSocketsPerNode(v int32)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0038JobProperties) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetSpreadJob

`func (o *V0038JobProperties) GetSpreadJob() bool`

GetSpreadJob returns the SpreadJob field if non-nil, zero value otherwise.

### GetSpreadJobOk

`func (o *V0038JobProperties) GetSpreadJobOk() (*bool, bool)`

GetSpreadJobOk returns a tuple with the SpreadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadJob

`func (o *V0038JobProperties) SetSpreadJob(v bool)`

SetSpreadJob sets SpreadJob field to given value.

### HasSpreadJob

`func (o *V0038JobProperties) HasSpreadJob() bool`

HasSpreadJob returns a boolean if a field has been set.

### GetStandardError

`func (o *V0038JobProperties) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *V0038JobProperties) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *V0038JobProperties) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *V0038JobProperties) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStandardInput

`func (o *V0038JobProperties) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *V0038JobProperties) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *V0038JobProperties) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *V0038JobProperties) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *V0038JobProperties) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *V0038JobProperties) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *V0038JobProperties) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *V0038JobProperties) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetTasks

`func (o *V0038JobProperties) GetTasks() int32`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0038JobProperties) GetTasksOk() (*int32, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0038JobProperties) SetTasks(v int32)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0038JobProperties) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *V0038JobProperties) GetTasksPerCore() int32`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *V0038JobProperties) GetTasksPerCoreOk() (*int32, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *V0038JobProperties) SetTasksPerCore(v int32)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *V0038JobProperties) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerNode

`func (o *V0038JobProperties) GetTasksPerNode() int32`

GetTasksPerNode returns the TasksPerNode field if non-nil, zero value otherwise.

### GetTasksPerNodeOk

`func (o *V0038JobProperties) GetTasksPerNodeOk() (*int32, bool)`

GetTasksPerNodeOk returns a tuple with the TasksPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerNode

`func (o *V0038JobProperties) SetTasksPerNode(v int32)`

SetTasksPerNode sets TasksPerNode field to given value.

### HasTasksPerNode

`func (o *V0038JobProperties) HasTasksPerNode() bool`

HasTasksPerNode returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *V0038JobProperties) GetTasksPerSocket() int32`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *V0038JobProperties) GetTasksPerSocketOk() (*int32, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *V0038JobProperties) SetTasksPerSocket(v int32)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *V0038JobProperties) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetThreadSpecification

`func (o *V0038JobProperties) GetThreadSpecification() int32`

GetThreadSpecification returns the ThreadSpecification field if non-nil, zero value otherwise.

### GetThreadSpecificationOk

`func (o *V0038JobProperties) GetThreadSpecificationOk() (*int32, bool)`

GetThreadSpecificationOk returns a tuple with the ThreadSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpecification

`func (o *V0038JobProperties) SetThreadSpecification(v int32)`

SetThreadSpecification sets ThreadSpecification field to given value.

### HasThreadSpecification

`func (o *V0038JobProperties) HasThreadSpecification() bool`

HasThreadSpecification returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *V0038JobProperties) GetThreadsPerCore() int32`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0038JobProperties) GetThreadsPerCoreOk() (*int32, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0038JobProperties) SetThreadsPerCore(v int32)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *V0038JobProperties) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTimeLimit

`func (o *V0038JobProperties) GetTimeLimit() int32`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *V0038JobProperties) GetTimeLimitOk() (*int32, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *V0038JobProperties) SetTimeLimit(v int32)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *V0038JobProperties) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *V0038JobProperties) GetTimeMinimum() int32`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *V0038JobProperties) GetTimeMinimumOk() (*int32, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *V0038JobProperties) SetTimeMinimum(v int32)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *V0038JobProperties) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetWaitAllNodes

`func (o *V0038JobProperties) GetWaitAllNodes() bool`

GetWaitAllNodes returns the WaitAllNodes field if non-nil, zero value otherwise.

### GetWaitAllNodesOk

`func (o *V0038JobProperties) GetWaitAllNodesOk() (*bool, bool)`

GetWaitAllNodesOk returns a tuple with the WaitAllNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitAllNodes

`func (o *V0038JobProperties) SetWaitAllNodes(v bool)`

SetWaitAllNodes sets WaitAllNodes field to given value.

### HasWaitAllNodes

`func (o *V0038JobProperties) HasWaitAllNodes() bool`

HasWaitAllNodes returns a boolean if a field has been set.

### GetWckey

`func (o *V0038JobProperties) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0038JobProperties) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0038JobProperties) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0038JobProperties) HasWckey() bool`

HasWckey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


