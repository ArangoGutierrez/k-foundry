# V0038Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | computer architecture | [optional] 
**BurstbufferNetworkAddress** | Pointer to **string** | BcastAddr | [optional] 
**Boards** | Pointer to **int32** | total number of boards per node | [optional] 
**BootTime** | Pointer to **int64** | timestamp of node boot | [optional] 
**Cores** | Pointer to **int32** | number of cores per socket | [optional] 
**CpuBinding** | Pointer to **int32** | Default task binding | [optional] 
**CpuLoad** | Pointer to **int64** | CPU load * 100 | [optional] 
**FreeMemory** | Pointer to **int32** | free memory in MiB | [optional] 
**Cpus** | Pointer to **int32** | configured count of cpus running on the node | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**ActiveFeatures** | Pointer to **string** | list of a node&#39;s available features | [optional] 
**Gres** | Pointer to **string** | list of a node&#39;s generic resources | [optional] 
**GresDrained** | Pointer to **string** | list of drained GRES | [optional] 
**GresUsed** | Pointer to **string** | list of GRES in current use | [optional] 
**McsLabel** | Pointer to **string** | mcs label if mcs plugin in use | [optional] 
**Name** | Pointer to **string** | node name to slurm | [optional] 
**NextStateAfterReboot** | Pointer to **string** | state after reboot | [optional] 
**NextStateAfterRebootFlags** | Pointer to **[]string** | node state flags | [optional] 
**Address** | Pointer to **string** | state after reboot | [optional] 
**Hostname** | Pointer to **string** | node&#39;s hostname | [optional] 
**State** | Pointer to **string** | current node state | [optional] 
**StateFlags** | Pointer to **[]string** | node state flags | [optional] 
**OperatingSystem** | Pointer to **string** | operating system | [optional] 
**Owner** | Pointer to **string** | User allowed to use this node | [optional] 
**Partitions** | Pointer to **[]string** | assigned partitions | [optional] 
**Port** | Pointer to **int32** | TCP port number of the slurmd | [optional] 
**RealMemory** | Pointer to **int32** | configured MB of real memory on the node | [optional] 
**Reason** | Pointer to **string** | reason for node being DOWN or DRAINING | [optional] 
**ReasonChangedAt** | Pointer to **int32** | Time stamp when reason was set | [optional] 
**ReasonSetByUser** | Pointer to **string** | User that set the reason | [optional] 
**SlurmdStartTime** | Pointer to **int64** | timestamp of slurmd startup | [optional] 
**Sockets** | Pointer to **int32** | total number of sockets per node | [optional] 
**Threads** | Pointer to **int32** | number of threads per core | [optional] 
**TemporaryDisk** | Pointer to **int32** | configured MB of total disk in TMP_FS | [optional] 
**Weight** | Pointer to **int32** | arbitrary priority of node for scheduling | [optional] 
**Tres** | Pointer to **string** | TRES on node | [optional] 
**TresUsed** | Pointer to **string** | TRES used on node | [optional] 
**TresWeighted** | Pointer to **float64** | TRES weight used on node | [optional] 
**SlurmdVersion** | Pointer to **string** | Slurmd version | [optional] 
**AllocCpus** | Pointer to **int64** | Allocated CPUs | [optional] 
**IdleCpus** | Pointer to **int64** | Idle CPUs | [optional] 
**AllocMemory** | Pointer to **int64** | Allocated memory (MB) | [optional] 

## Methods

### NewV0038Node

`func NewV0038Node() *V0038Node`

NewV0038Node instantiates a new V0038Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038NodeWithDefaults

`func NewV0038NodeWithDefaults() *V0038Node`

NewV0038NodeWithDefaults instantiates a new V0038Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *V0038Node) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *V0038Node) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *V0038Node) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *V0038Node) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBurstbufferNetworkAddress

`func (o *V0038Node) GetBurstbufferNetworkAddress() string`

GetBurstbufferNetworkAddress returns the BurstbufferNetworkAddress field if non-nil, zero value otherwise.

### GetBurstbufferNetworkAddressOk

`func (o *V0038Node) GetBurstbufferNetworkAddressOk() (*string, bool)`

GetBurstbufferNetworkAddressOk returns a tuple with the BurstbufferNetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstbufferNetworkAddress

`func (o *V0038Node) SetBurstbufferNetworkAddress(v string)`

SetBurstbufferNetworkAddress sets BurstbufferNetworkAddress field to given value.

### HasBurstbufferNetworkAddress

`func (o *V0038Node) HasBurstbufferNetworkAddress() bool`

HasBurstbufferNetworkAddress returns a boolean if a field has been set.

### GetBoards

`func (o *V0038Node) GetBoards() int32`

GetBoards returns the Boards field if non-nil, zero value otherwise.

### GetBoardsOk

`func (o *V0038Node) GetBoardsOk() (*int32, bool)`

GetBoardsOk returns a tuple with the Boards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoards

`func (o *V0038Node) SetBoards(v int32)`

SetBoards sets Boards field to given value.

### HasBoards

`func (o *V0038Node) HasBoards() bool`

HasBoards returns a boolean if a field has been set.

### GetBootTime

`func (o *V0038Node) GetBootTime() int64`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *V0038Node) GetBootTimeOk() (*int64, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *V0038Node) SetBootTime(v int64)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *V0038Node) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetCores

`func (o *V0038Node) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *V0038Node) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *V0038Node) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *V0038Node) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpuBinding

`func (o *V0038Node) GetCpuBinding() int32`

GetCpuBinding returns the CpuBinding field if non-nil, zero value otherwise.

### GetCpuBindingOk

`func (o *V0038Node) GetCpuBindingOk() (*int32, bool)`

GetCpuBindingOk returns a tuple with the CpuBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBinding

`func (o *V0038Node) SetCpuBinding(v int32)`

SetCpuBinding sets CpuBinding field to given value.

### HasCpuBinding

`func (o *V0038Node) HasCpuBinding() bool`

HasCpuBinding returns a boolean if a field has been set.

### GetCpuLoad

`func (o *V0038Node) GetCpuLoad() int64`

GetCpuLoad returns the CpuLoad field if non-nil, zero value otherwise.

### GetCpuLoadOk

`func (o *V0038Node) GetCpuLoadOk() (*int64, bool)`

GetCpuLoadOk returns a tuple with the CpuLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLoad

`func (o *V0038Node) SetCpuLoad(v int64)`

SetCpuLoad sets CpuLoad field to given value.

### HasCpuLoad

`func (o *V0038Node) HasCpuLoad() bool`

HasCpuLoad returns a boolean if a field has been set.

### GetFreeMemory

`func (o *V0038Node) GetFreeMemory() int32`

GetFreeMemory returns the FreeMemory field if non-nil, zero value otherwise.

### GetFreeMemoryOk

`func (o *V0038Node) GetFreeMemoryOk() (*int32, bool)`

GetFreeMemoryOk returns a tuple with the FreeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemory

`func (o *V0038Node) SetFreeMemory(v int32)`

SetFreeMemory sets FreeMemory field to given value.

### HasFreeMemory

`func (o *V0038Node) HasFreeMemory() bool`

HasFreeMemory returns a boolean if a field has been set.

### GetCpus

`func (o *V0038Node) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0038Node) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0038Node) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0038Node) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetFeatures

`func (o *V0038Node) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0038Node) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0038Node) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0038Node) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetActiveFeatures

`func (o *V0038Node) GetActiveFeatures() string`

GetActiveFeatures returns the ActiveFeatures field if non-nil, zero value otherwise.

### GetActiveFeaturesOk

`func (o *V0038Node) GetActiveFeaturesOk() (*string, bool)`

GetActiveFeaturesOk returns a tuple with the ActiveFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFeatures

`func (o *V0038Node) SetActiveFeatures(v string)`

SetActiveFeatures sets ActiveFeatures field to given value.

### HasActiveFeatures

`func (o *V0038Node) HasActiveFeatures() bool`

HasActiveFeatures returns a boolean if a field has been set.

### GetGres

`func (o *V0038Node) GetGres() string`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *V0038Node) GetGresOk() (*string, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *V0038Node) SetGres(v string)`

SetGres sets Gres field to given value.

### HasGres

`func (o *V0038Node) HasGres() bool`

HasGres returns a boolean if a field has been set.

### GetGresDrained

`func (o *V0038Node) GetGresDrained() string`

GetGresDrained returns the GresDrained field if non-nil, zero value otherwise.

### GetGresDrainedOk

`func (o *V0038Node) GetGresDrainedOk() (*string, bool)`

GetGresDrainedOk returns a tuple with the GresDrained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresDrained

`func (o *V0038Node) SetGresDrained(v string)`

SetGresDrained sets GresDrained field to given value.

### HasGresDrained

`func (o *V0038Node) HasGresDrained() bool`

HasGresDrained returns a boolean if a field has been set.

### GetGresUsed

`func (o *V0038Node) GetGresUsed() string`

GetGresUsed returns the GresUsed field if non-nil, zero value otherwise.

### GetGresUsedOk

`func (o *V0038Node) GetGresUsedOk() (*string, bool)`

GetGresUsedOk returns a tuple with the GresUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresUsed

`func (o *V0038Node) SetGresUsed(v string)`

SetGresUsed sets GresUsed field to given value.

### HasGresUsed

`func (o *V0038Node) HasGresUsed() bool`

HasGresUsed returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0038Node) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0038Node) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0038Node) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0038Node) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetName

`func (o *V0038Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0038Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0038Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0038Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextStateAfterReboot

`func (o *V0038Node) GetNextStateAfterReboot() string`

GetNextStateAfterReboot returns the NextStateAfterReboot field if non-nil, zero value otherwise.

### GetNextStateAfterRebootOk

`func (o *V0038Node) GetNextStateAfterRebootOk() (*string, bool)`

GetNextStateAfterRebootOk returns a tuple with the NextStateAfterReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStateAfterReboot

`func (o *V0038Node) SetNextStateAfterReboot(v string)`

SetNextStateAfterReboot sets NextStateAfterReboot field to given value.

### HasNextStateAfterReboot

`func (o *V0038Node) HasNextStateAfterReboot() bool`

HasNextStateAfterReboot returns a boolean if a field has been set.

### GetNextStateAfterRebootFlags

`func (o *V0038Node) GetNextStateAfterRebootFlags() []string`

GetNextStateAfterRebootFlags returns the NextStateAfterRebootFlags field if non-nil, zero value otherwise.

### GetNextStateAfterRebootFlagsOk

`func (o *V0038Node) GetNextStateAfterRebootFlagsOk() (*[]string, bool)`

GetNextStateAfterRebootFlagsOk returns a tuple with the NextStateAfterRebootFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStateAfterRebootFlags

`func (o *V0038Node) SetNextStateAfterRebootFlags(v []string)`

SetNextStateAfterRebootFlags sets NextStateAfterRebootFlags field to given value.

### HasNextStateAfterRebootFlags

`func (o *V0038Node) HasNextStateAfterRebootFlags() bool`

HasNextStateAfterRebootFlags returns a boolean if a field has been set.

### GetAddress

`func (o *V0038Node) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V0038Node) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V0038Node) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V0038Node) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *V0038Node) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0038Node) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0038Node) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0038Node) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetState

`func (o *V0038Node) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0038Node) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0038Node) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V0038Node) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateFlags

`func (o *V0038Node) GetStateFlags() []string`

GetStateFlags returns the StateFlags field if non-nil, zero value otherwise.

### GetStateFlagsOk

`func (o *V0038Node) GetStateFlagsOk() (*[]string, bool)`

GetStateFlagsOk returns a tuple with the StateFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFlags

`func (o *V0038Node) SetStateFlags(v []string)`

SetStateFlags sets StateFlags field to given value.

### HasStateFlags

`func (o *V0038Node) HasStateFlags() bool`

HasStateFlags returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *V0038Node) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *V0038Node) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *V0038Node) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *V0038Node) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOwner

`func (o *V0038Node) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V0038Node) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V0038Node) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V0038Node) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPartitions

`func (o *V0038Node) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0038Node) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0038Node) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0038Node) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetPort

`func (o *V0038Node) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V0038Node) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V0038Node) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V0038Node) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRealMemory

`func (o *V0038Node) GetRealMemory() int32`

GetRealMemory returns the RealMemory field if non-nil, zero value otherwise.

### GetRealMemoryOk

`func (o *V0038Node) GetRealMemoryOk() (*int32, bool)`

GetRealMemoryOk returns a tuple with the RealMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealMemory

`func (o *V0038Node) SetRealMemory(v int32)`

SetRealMemory sets RealMemory field to given value.

### HasRealMemory

`func (o *V0038Node) HasRealMemory() bool`

HasRealMemory returns a boolean if a field has been set.

### GetReason

`func (o *V0038Node) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V0038Node) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V0038Node) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V0038Node) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonChangedAt

`func (o *V0038Node) GetReasonChangedAt() int32`

GetReasonChangedAt returns the ReasonChangedAt field if non-nil, zero value otherwise.

### GetReasonChangedAtOk

`func (o *V0038Node) GetReasonChangedAtOk() (*int32, bool)`

GetReasonChangedAtOk returns a tuple with the ReasonChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonChangedAt

`func (o *V0038Node) SetReasonChangedAt(v int32)`

SetReasonChangedAt sets ReasonChangedAt field to given value.

### HasReasonChangedAt

`func (o *V0038Node) HasReasonChangedAt() bool`

HasReasonChangedAt returns a boolean if a field has been set.

### GetReasonSetByUser

`func (o *V0038Node) GetReasonSetByUser() string`

GetReasonSetByUser returns the ReasonSetByUser field if non-nil, zero value otherwise.

### GetReasonSetByUserOk

`func (o *V0038Node) GetReasonSetByUserOk() (*string, bool)`

GetReasonSetByUserOk returns a tuple with the ReasonSetByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSetByUser

`func (o *V0038Node) SetReasonSetByUser(v string)`

SetReasonSetByUser sets ReasonSetByUser field to given value.

### HasReasonSetByUser

`func (o *V0038Node) HasReasonSetByUser() bool`

HasReasonSetByUser returns a boolean if a field has been set.

### GetSlurmdStartTime

`func (o *V0038Node) GetSlurmdStartTime() int64`

GetSlurmdStartTime returns the SlurmdStartTime field if non-nil, zero value otherwise.

### GetSlurmdStartTimeOk

`func (o *V0038Node) GetSlurmdStartTimeOk() (*int64, bool)`

GetSlurmdStartTimeOk returns a tuple with the SlurmdStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurmdStartTime

`func (o *V0038Node) SetSlurmdStartTime(v int64)`

SetSlurmdStartTime sets SlurmdStartTime field to given value.

### HasSlurmdStartTime

`func (o *V0038Node) HasSlurmdStartTime() bool`

HasSlurmdStartTime returns a boolean if a field has been set.

### GetSockets

`func (o *V0038Node) GetSockets() int32`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0038Node) GetSocketsOk() (*int32, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0038Node) SetSockets(v int32)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *V0038Node) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetThreads

`func (o *V0038Node) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *V0038Node) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *V0038Node) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *V0038Node) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetTemporaryDisk

`func (o *V0038Node) GetTemporaryDisk() int32`

GetTemporaryDisk returns the TemporaryDisk field if non-nil, zero value otherwise.

### GetTemporaryDiskOk

`func (o *V0038Node) GetTemporaryDiskOk() (*int32, bool)`

GetTemporaryDiskOk returns a tuple with the TemporaryDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDisk

`func (o *V0038Node) SetTemporaryDisk(v int32)`

SetTemporaryDisk sets TemporaryDisk field to given value.

### HasTemporaryDisk

`func (o *V0038Node) HasTemporaryDisk() bool`

HasTemporaryDisk returns a boolean if a field has been set.

### GetWeight

`func (o *V0038Node) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V0038Node) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V0038Node) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V0038Node) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetTres

`func (o *V0038Node) GetTres() string`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0038Node) GetTresOk() (*string, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0038Node) SetTres(v string)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0038Node) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetTresUsed

`func (o *V0038Node) GetTresUsed() string`

GetTresUsed returns the TresUsed field if non-nil, zero value otherwise.

### GetTresUsedOk

`func (o *V0038Node) GetTresUsedOk() (*string, bool)`

GetTresUsedOk returns a tuple with the TresUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresUsed

`func (o *V0038Node) SetTresUsed(v string)`

SetTresUsed sets TresUsed field to given value.

### HasTresUsed

`func (o *V0038Node) HasTresUsed() bool`

HasTresUsed returns a boolean if a field has been set.

### GetTresWeighted

`func (o *V0038Node) GetTresWeighted() float64`

GetTresWeighted returns the TresWeighted field if non-nil, zero value otherwise.

### GetTresWeightedOk

`func (o *V0038Node) GetTresWeightedOk() (*float64, bool)`

GetTresWeightedOk returns a tuple with the TresWeighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresWeighted

`func (o *V0038Node) SetTresWeighted(v float64)`

SetTresWeighted sets TresWeighted field to given value.

### HasTresWeighted

`func (o *V0038Node) HasTresWeighted() bool`

HasTresWeighted returns a boolean if a field has been set.

### GetSlurmdVersion

`func (o *V0038Node) GetSlurmdVersion() string`

GetSlurmdVersion returns the SlurmdVersion field if non-nil, zero value otherwise.

### GetSlurmdVersionOk

`func (o *V0038Node) GetSlurmdVersionOk() (*string, bool)`

GetSlurmdVersionOk returns a tuple with the SlurmdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurmdVersion

`func (o *V0038Node) SetSlurmdVersion(v string)`

SetSlurmdVersion sets SlurmdVersion field to given value.

### HasSlurmdVersion

`func (o *V0038Node) HasSlurmdVersion() bool`

HasSlurmdVersion returns a boolean if a field has been set.

### GetAllocCpus

`func (o *V0038Node) GetAllocCpus() int64`

GetAllocCpus returns the AllocCpus field if non-nil, zero value otherwise.

### GetAllocCpusOk

`func (o *V0038Node) GetAllocCpusOk() (*int64, bool)`

GetAllocCpusOk returns a tuple with the AllocCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocCpus

`func (o *V0038Node) SetAllocCpus(v int64)`

SetAllocCpus sets AllocCpus field to given value.

### HasAllocCpus

`func (o *V0038Node) HasAllocCpus() bool`

HasAllocCpus returns a boolean if a field has been set.

### GetIdleCpus

`func (o *V0038Node) GetIdleCpus() int64`

GetIdleCpus returns the IdleCpus field if non-nil, zero value otherwise.

### GetIdleCpusOk

`func (o *V0038Node) GetIdleCpusOk() (*int64, bool)`

GetIdleCpusOk returns a tuple with the IdleCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleCpus

`func (o *V0038Node) SetIdleCpus(v int64)`

SetIdleCpus sets IdleCpus field to given value.

### HasIdleCpus

`func (o *V0038Node) HasIdleCpus() bool`

HasIdleCpus returns a boolean if a field has been set.

### GetAllocMemory

`func (o *V0038Node) GetAllocMemory() int64`

GetAllocMemory returns the AllocMemory field if non-nil, zero value otherwise.

### GetAllocMemoryOk

`func (o *V0038Node) GetAllocMemoryOk() (*int64, bool)`

GetAllocMemoryOk returns a tuple with the AllocMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocMemory

`func (o *V0038Node) SetAllocMemory(v int64)`

SetAllocMemory sets AllocMemory field to given value.

### HasAllocMemory

`func (o *V0038Node) HasAllocMemory() bool`

HasAllocMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


