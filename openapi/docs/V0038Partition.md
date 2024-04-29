# V0038Partition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **[]string** | partition options | [optional] 
**PreemptionMode** | Pointer to **[]string** | preemption type | [optional] 
**AllowedAllocationNodes** | Pointer to **string** | list names of allowed allocating nodes | [optional] 
**AllowedAccounts** | Pointer to **string** | comma delimited list of accounts | [optional] 
**AllowedGroups** | Pointer to **string** | comma delimited list of groups | [optional] 
**AllowedQos** | Pointer to **string** | comma delimited list of qos | [optional] 
**Alternative** | Pointer to **string** | name of alternate partition | [optional] 
**BillingWeights** | Pointer to **string** | TRES billing weights | [optional] 
**DefaultMemoryPerCpu** | Pointer to **int64** | default MB memory per allocated CPU | [optional] 
**DefaultTimeLimit** | Pointer to **int64** | default time limit (minutes) | [optional] 
**DeniedAccounts** | Pointer to **string** | comma delimited list of denied accounts | [optional] 
**DeniedQos** | Pointer to **string** | comma delimited list of denied qos | [optional] 
**PreemptionGraceTime** | Pointer to **int64** | preemption grace time (seconds) | [optional] 
**MaximumCpusPerNode** | Pointer to **int32** | maximum allocated CPUs per node | [optional] 
**MaximumMemoryPerNode** | Pointer to **int64** | maximum memory per allocated node (MiB) | [optional] 
**MaximumNodesPerJob** | Pointer to **int32** | Max nodes per job | [optional] 
**MaxTimeLimit** | Pointer to **int64** | Max time limit per job | [optional] 
**MinNodesPerJob** | Pointer to **int32** | Min number of nodes per job | [optional] 
**Name** | Pointer to **string** | Partition name | [optional] 
**Nodes** | Pointer to **string** | list names of nodes in partition | [optional] 
**OverTimeLimit** | Pointer to **int32** | job&#39;s time limit can be exceeded by this number of minutes before cancellation | [optional] 
**PriorityJobFactor** | Pointer to **int32** | job priority weight factor | [optional] 
**PriorityTier** | Pointer to **int32** | tier for scheduling and preemption | [optional] 
**Qos** | Pointer to **string** | partition QOS name | [optional] 
**State** | Pointer to **string** | Partition state | [optional] 
**TotalCpus** | Pointer to **int32** | Total cpus in partition | [optional] 
**TotalNodes** | Pointer to **int32** | Total number of nodes in partition | [optional] 
**Tres** | Pointer to **string** | configured TRES in partition | [optional] 
**MaximumMemoryPerCpu** | Pointer to **int64** | maximum memory per allocated CPU (MiB) | [optional] 
**DefaultMemoryPerNode** | Pointer to **int64** | default MB memory per allocated node | [optional] 

## Methods

### NewV0038Partition

`func NewV0038Partition() *V0038Partition`

NewV0038Partition instantiates a new V0038Partition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038PartitionWithDefaults

`func NewV0038PartitionWithDefaults() *V0038Partition`

NewV0038PartitionWithDefaults instantiates a new V0038Partition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *V0038Partition) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0038Partition) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0038Partition) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0038Partition) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetPreemptionMode

`func (o *V0038Partition) GetPreemptionMode() []string`

GetPreemptionMode returns the PreemptionMode field if non-nil, zero value otherwise.

### GetPreemptionModeOk

`func (o *V0038Partition) GetPreemptionModeOk() (*[]string, bool)`

GetPreemptionModeOk returns a tuple with the PreemptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptionMode

`func (o *V0038Partition) SetPreemptionMode(v []string)`

SetPreemptionMode sets PreemptionMode field to given value.

### HasPreemptionMode

`func (o *V0038Partition) HasPreemptionMode() bool`

HasPreemptionMode returns a boolean if a field has been set.

### GetAllowedAllocationNodes

`func (o *V0038Partition) GetAllowedAllocationNodes() string`

GetAllowedAllocationNodes returns the AllowedAllocationNodes field if non-nil, zero value otherwise.

### GetAllowedAllocationNodesOk

`func (o *V0038Partition) GetAllowedAllocationNodesOk() (*string, bool)`

GetAllowedAllocationNodesOk returns a tuple with the AllowedAllocationNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAllocationNodes

`func (o *V0038Partition) SetAllowedAllocationNodes(v string)`

SetAllowedAllocationNodes sets AllowedAllocationNodes field to given value.

### HasAllowedAllocationNodes

`func (o *V0038Partition) HasAllowedAllocationNodes() bool`

HasAllowedAllocationNodes returns a boolean if a field has been set.

### GetAllowedAccounts

`func (o *V0038Partition) GetAllowedAccounts() string`

GetAllowedAccounts returns the AllowedAccounts field if non-nil, zero value otherwise.

### GetAllowedAccountsOk

`func (o *V0038Partition) GetAllowedAccountsOk() (*string, bool)`

GetAllowedAccountsOk returns a tuple with the AllowedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAccounts

`func (o *V0038Partition) SetAllowedAccounts(v string)`

SetAllowedAccounts sets AllowedAccounts field to given value.

### HasAllowedAccounts

`func (o *V0038Partition) HasAllowedAccounts() bool`

HasAllowedAccounts returns a boolean if a field has been set.

### GetAllowedGroups

`func (o *V0038Partition) GetAllowedGroups() string`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *V0038Partition) GetAllowedGroupsOk() (*string, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *V0038Partition) SetAllowedGroups(v string)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *V0038Partition) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetAllowedQos

`func (o *V0038Partition) GetAllowedQos() string`

GetAllowedQos returns the AllowedQos field if non-nil, zero value otherwise.

### GetAllowedQosOk

`func (o *V0038Partition) GetAllowedQosOk() (*string, bool)`

GetAllowedQosOk returns a tuple with the AllowedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedQos

`func (o *V0038Partition) SetAllowedQos(v string)`

SetAllowedQos sets AllowedQos field to given value.

### HasAllowedQos

`func (o *V0038Partition) HasAllowedQos() bool`

HasAllowedQos returns a boolean if a field has been set.

### GetAlternative

`func (o *V0038Partition) GetAlternative() string`

GetAlternative returns the Alternative field if non-nil, zero value otherwise.

### GetAlternativeOk

`func (o *V0038Partition) GetAlternativeOk() (*string, bool)`

GetAlternativeOk returns a tuple with the Alternative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternative

`func (o *V0038Partition) SetAlternative(v string)`

SetAlternative sets Alternative field to given value.

### HasAlternative

`func (o *V0038Partition) HasAlternative() bool`

HasAlternative returns a boolean if a field has been set.

### GetBillingWeights

`func (o *V0038Partition) GetBillingWeights() string`

GetBillingWeights returns the BillingWeights field if non-nil, zero value otherwise.

### GetBillingWeightsOk

`func (o *V0038Partition) GetBillingWeightsOk() (*string, bool)`

GetBillingWeightsOk returns a tuple with the BillingWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingWeights

`func (o *V0038Partition) SetBillingWeights(v string)`

SetBillingWeights sets BillingWeights field to given value.

### HasBillingWeights

`func (o *V0038Partition) HasBillingWeights() bool`

HasBillingWeights returns a boolean if a field has been set.

### GetDefaultMemoryPerCpu

`func (o *V0038Partition) GetDefaultMemoryPerCpu() int64`

GetDefaultMemoryPerCpu returns the DefaultMemoryPerCpu field if non-nil, zero value otherwise.

### GetDefaultMemoryPerCpuOk

`func (o *V0038Partition) GetDefaultMemoryPerCpuOk() (*int64, bool)`

GetDefaultMemoryPerCpuOk returns a tuple with the DefaultMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemoryPerCpu

`func (o *V0038Partition) SetDefaultMemoryPerCpu(v int64)`

SetDefaultMemoryPerCpu sets DefaultMemoryPerCpu field to given value.

### HasDefaultMemoryPerCpu

`func (o *V0038Partition) HasDefaultMemoryPerCpu() bool`

HasDefaultMemoryPerCpu returns a boolean if a field has been set.

### GetDefaultTimeLimit

`func (o *V0038Partition) GetDefaultTimeLimit() int64`

GetDefaultTimeLimit returns the DefaultTimeLimit field if non-nil, zero value otherwise.

### GetDefaultTimeLimitOk

`func (o *V0038Partition) GetDefaultTimeLimitOk() (*int64, bool)`

GetDefaultTimeLimitOk returns a tuple with the DefaultTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeLimit

`func (o *V0038Partition) SetDefaultTimeLimit(v int64)`

SetDefaultTimeLimit sets DefaultTimeLimit field to given value.

### HasDefaultTimeLimit

`func (o *V0038Partition) HasDefaultTimeLimit() bool`

HasDefaultTimeLimit returns a boolean if a field has been set.

### GetDeniedAccounts

`func (o *V0038Partition) GetDeniedAccounts() string`

GetDeniedAccounts returns the DeniedAccounts field if non-nil, zero value otherwise.

### GetDeniedAccountsOk

`func (o *V0038Partition) GetDeniedAccountsOk() (*string, bool)`

GetDeniedAccountsOk returns a tuple with the DeniedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedAccounts

`func (o *V0038Partition) SetDeniedAccounts(v string)`

SetDeniedAccounts sets DeniedAccounts field to given value.

### HasDeniedAccounts

`func (o *V0038Partition) HasDeniedAccounts() bool`

HasDeniedAccounts returns a boolean if a field has been set.

### GetDeniedQos

`func (o *V0038Partition) GetDeniedQos() string`

GetDeniedQos returns the DeniedQos field if non-nil, zero value otherwise.

### GetDeniedQosOk

`func (o *V0038Partition) GetDeniedQosOk() (*string, bool)`

GetDeniedQosOk returns a tuple with the DeniedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedQos

`func (o *V0038Partition) SetDeniedQos(v string)`

SetDeniedQos sets DeniedQos field to given value.

### HasDeniedQos

`func (o *V0038Partition) HasDeniedQos() bool`

HasDeniedQos returns a boolean if a field has been set.

### GetPreemptionGraceTime

`func (o *V0038Partition) GetPreemptionGraceTime() int64`

GetPreemptionGraceTime returns the PreemptionGraceTime field if non-nil, zero value otherwise.

### GetPreemptionGraceTimeOk

`func (o *V0038Partition) GetPreemptionGraceTimeOk() (*int64, bool)`

GetPreemptionGraceTimeOk returns a tuple with the PreemptionGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptionGraceTime

`func (o *V0038Partition) SetPreemptionGraceTime(v int64)`

SetPreemptionGraceTime sets PreemptionGraceTime field to given value.

### HasPreemptionGraceTime

`func (o *V0038Partition) HasPreemptionGraceTime() bool`

HasPreemptionGraceTime returns a boolean if a field has been set.

### GetMaximumCpusPerNode

`func (o *V0038Partition) GetMaximumCpusPerNode() int32`

GetMaximumCpusPerNode returns the MaximumCpusPerNode field if non-nil, zero value otherwise.

### GetMaximumCpusPerNodeOk

`func (o *V0038Partition) GetMaximumCpusPerNodeOk() (*int32, bool)`

GetMaximumCpusPerNodeOk returns a tuple with the MaximumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpusPerNode

`func (o *V0038Partition) SetMaximumCpusPerNode(v int32)`

SetMaximumCpusPerNode sets MaximumCpusPerNode field to given value.

### HasMaximumCpusPerNode

`func (o *V0038Partition) HasMaximumCpusPerNode() bool`

HasMaximumCpusPerNode returns a boolean if a field has been set.

### GetMaximumMemoryPerNode

`func (o *V0038Partition) GetMaximumMemoryPerNode() int64`

GetMaximumMemoryPerNode returns the MaximumMemoryPerNode field if non-nil, zero value otherwise.

### GetMaximumMemoryPerNodeOk

`func (o *V0038Partition) GetMaximumMemoryPerNodeOk() (*int64, bool)`

GetMaximumMemoryPerNodeOk returns a tuple with the MaximumMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemoryPerNode

`func (o *V0038Partition) SetMaximumMemoryPerNode(v int64)`

SetMaximumMemoryPerNode sets MaximumMemoryPerNode field to given value.

### HasMaximumMemoryPerNode

`func (o *V0038Partition) HasMaximumMemoryPerNode() bool`

HasMaximumMemoryPerNode returns a boolean if a field has been set.

### GetMaximumNodesPerJob

`func (o *V0038Partition) GetMaximumNodesPerJob() int32`

GetMaximumNodesPerJob returns the MaximumNodesPerJob field if non-nil, zero value otherwise.

### GetMaximumNodesPerJobOk

`func (o *V0038Partition) GetMaximumNodesPerJobOk() (*int32, bool)`

GetMaximumNodesPerJobOk returns a tuple with the MaximumNodesPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNodesPerJob

`func (o *V0038Partition) SetMaximumNodesPerJob(v int32)`

SetMaximumNodesPerJob sets MaximumNodesPerJob field to given value.

### HasMaximumNodesPerJob

`func (o *V0038Partition) HasMaximumNodesPerJob() bool`

HasMaximumNodesPerJob returns a boolean if a field has been set.

### GetMaxTimeLimit

`func (o *V0038Partition) GetMaxTimeLimit() int64`

GetMaxTimeLimit returns the MaxTimeLimit field if non-nil, zero value otherwise.

### GetMaxTimeLimitOk

`func (o *V0038Partition) GetMaxTimeLimitOk() (*int64, bool)`

GetMaxTimeLimitOk returns a tuple with the MaxTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeLimit

`func (o *V0038Partition) SetMaxTimeLimit(v int64)`

SetMaxTimeLimit sets MaxTimeLimit field to given value.

### HasMaxTimeLimit

`func (o *V0038Partition) HasMaxTimeLimit() bool`

HasMaxTimeLimit returns a boolean if a field has been set.

### GetMinNodesPerJob

`func (o *V0038Partition) GetMinNodesPerJob() int32`

GetMinNodesPerJob returns the MinNodesPerJob field if non-nil, zero value otherwise.

### GetMinNodesPerJobOk

`func (o *V0038Partition) GetMinNodesPerJobOk() (*int32, bool)`

GetMinNodesPerJobOk returns a tuple with the MinNodesPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodesPerJob

`func (o *V0038Partition) SetMinNodesPerJob(v int32)`

SetMinNodesPerJob sets MinNodesPerJob field to given value.

### HasMinNodesPerJob

`func (o *V0038Partition) HasMinNodesPerJob() bool`

HasMinNodesPerJob returns a boolean if a field has been set.

### GetName

`func (o *V0038Partition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0038Partition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0038Partition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0038Partition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *V0038Partition) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0038Partition) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0038Partition) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0038Partition) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOverTimeLimit

`func (o *V0038Partition) GetOverTimeLimit() int32`

GetOverTimeLimit returns the OverTimeLimit field if non-nil, zero value otherwise.

### GetOverTimeLimitOk

`func (o *V0038Partition) GetOverTimeLimitOk() (*int32, bool)`

GetOverTimeLimitOk returns a tuple with the OverTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTimeLimit

`func (o *V0038Partition) SetOverTimeLimit(v int32)`

SetOverTimeLimit sets OverTimeLimit field to given value.

### HasOverTimeLimit

`func (o *V0038Partition) HasOverTimeLimit() bool`

HasOverTimeLimit returns a boolean if a field has been set.

### GetPriorityJobFactor

`func (o *V0038Partition) GetPriorityJobFactor() int32`

GetPriorityJobFactor returns the PriorityJobFactor field if non-nil, zero value otherwise.

### GetPriorityJobFactorOk

`func (o *V0038Partition) GetPriorityJobFactorOk() (*int32, bool)`

GetPriorityJobFactorOk returns a tuple with the PriorityJobFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityJobFactor

`func (o *V0038Partition) SetPriorityJobFactor(v int32)`

SetPriorityJobFactor sets PriorityJobFactor field to given value.

### HasPriorityJobFactor

`func (o *V0038Partition) HasPriorityJobFactor() bool`

HasPriorityJobFactor returns a boolean if a field has been set.

### GetPriorityTier

`func (o *V0038Partition) GetPriorityTier() int32`

GetPriorityTier returns the PriorityTier field if non-nil, zero value otherwise.

### GetPriorityTierOk

`func (o *V0038Partition) GetPriorityTierOk() (*int32, bool)`

GetPriorityTierOk returns a tuple with the PriorityTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityTier

`func (o *V0038Partition) SetPriorityTier(v int32)`

SetPriorityTier sets PriorityTier field to given value.

### HasPriorityTier

`func (o *V0038Partition) HasPriorityTier() bool`

HasPriorityTier returns a boolean if a field has been set.

### GetQos

`func (o *V0038Partition) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0038Partition) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0038Partition) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0038Partition) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetState

`func (o *V0038Partition) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0038Partition) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0038Partition) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V0038Partition) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTotalCpus

`func (o *V0038Partition) GetTotalCpus() int32`

GetTotalCpus returns the TotalCpus field if non-nil, zero value otherwise.

### GetTotalCpusOk

`func (o *V0038Partition) GetTotalCpusOk() (*int32, bool)`

GetTotalCpusOk returns a tuple with the TotalCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpus

`func (o *V0038Partition) SetTotalCpus(v int32)`

SetTotalCpus sets TotalCpus field to given value.

### HasTotalCpus

`func (o *V0038Partition) HasTotalCpus() bool`

HasTotalCpus returns a boolean if a field has been set.

### GetTotalNodes

`func (o *V0038Partition) GetTotalNodes() int32`

GetTotalNodes returns the TotalNodes field if non-nil, zero value otherwise.

### GetTotalNodesOk

`func (o *V0038Partition) GetTotalNodesOk() (*int32, bool)`

GetTotalNodesOk returns a tuple with the TotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNodes

`func (o *V0038Partition) SetTotalNodes(v int32)`

SetTotalNodes sets TotalNodes field to given value.

### HasTotalNodes

`func (o *V0038Partition) HasTotalNodes() bool`

HasTotalNodes returns a boolean if a field has been set.

### GetTres

`func (o *V0038Partition) GetTres() string`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0038Partition) GetTresOk() (*string, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0038Partition) SetTres(v string)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0038Partition) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetMaximumMemoryPerCpu

`func (o *V0038Partition) GetMaximumMemoryPerCpu() int64`

GetMaximumMemoryPerCpu returns the MaximumMemoryPerCpu field if non-nil, zero value otherwise.

### GetMaximumMemoryPerCpuOk

`func (o *V0038Partition) GetMaximumMemoryPerCpuOk() (*int64, bool)`

GetMaximumMemoryPerCpuOk returns a tuple with the MaximumMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemoryPerCpu

`func (o *V0038Partition) SetMaximumMemoryPerCpu(v int64)`

SetMaximumMemoryPerCpu sets MaximumMemoryPerCpu field to given value.

### HasMaximumMemoryPerCpu

`func (o *V0038Partition) HasMaximumMemoryPerCpu() bool`

HasMaximumMemoryPerCpu returns a boolean if a field has been set.

### GetDefaultMemoryPerNode

`func (o *V0038Partition) GetDefaultMemoryPerNode() int64`

GetDefaultMemoryPerNode returns the DefaultMemoryPerNode field if non-nil, zero value otherwise.

### GetDefaultMemoryPerNodeOk

`func (o *V0038Partition) GetDefaultMemoryPerNodeOk() (*int64, bool)`

GetDefaultMemoryPerNodeOk returns a tuple with the DefaultMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemoryPerNode

`func (o *V0038Partition) SetDefaultMemoryPerNode(v int64)`

SetDefaultMemoryPerNode sets DefaultMemoryPerNode field to given value.

### HasDefaultMemoryPerNode

`func (o *V0038Partition) HasDefaultMemoryPerNode() bool`

HasDefaultMemoryPerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


