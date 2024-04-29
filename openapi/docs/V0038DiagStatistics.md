# V0038DiagStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartsPacked** | Pointer to **int32** | partition records packed | [optional] 
**ReqTime** | Pointer to **int32** | generation time | [optional] 
**ReqTimeStart** | Pointer to **int32** | data since | [optional] 
**ServerThreadCount** | Pointer to **int32** | Server thread count | [optional] 
**AgentQueueSize** | Pointer to **int32** | Agent queue size | [optional] 
**AgentCount** | Pointer to **int32** | Agent count | [optional] 
**AgentThreadCount** | Pointer to **int32** | Agent thread count | [optional] 
**DbdAgentQueueSize** | Pointer to **int32** | DBD Agent queue size | [optional] 
**GettimeofdayLatency** | Pointer to **int32** | Latency for 1000 calls to gettimeofday() | [optional] 
**ScheduleCycleMax** | Pointer to **int32** | Main Schedule max cycle | [optional] 
**ScheduleCycleLast** | Pointer to **int32** | Main Schedule last cycle | [optional] 
**ScheduleCycleTotal** | Pointer to **int32** | Main Schedule cycle iterations | [optional] 
**ScheduleCycleMean** | Pointer to **int32** | Average time for Schedule Max cycle | [optional] 
**ScheduleCycleMeanDepth** | Pointer to **int32** | Average depth for Schedule Max cycle | [optional] 
**ScheduleCyclePerMinute** | Pointer to **int32** | Main Schedule Cycles per minute | [optional] 
**ScheduleQueueLength** | Pointer to **int32** | Main Schedule Last queue length | [optional] 
**JobsSubmitted** | Pointer to **int32** | Job submitted | [optional] 
**JobsStarted** | Pointer to **int32** | Job started | [optional] 
**JobsCompleted** | Pointer to **int32** | Job completed | [optional] 
**JobsCanceled** | Pointer to **int32** | Job cancelled | [optional] 
**JobsFailed** | Pointer to **int32** | Job failed | [optional] 
**JobsPending** | Pointer to **int32** | Job pending | [optional] 
**JobsRunning** | Pointer to **int32** | Job running | [optional] 
**JobStatesTs** | Pointer to **int32** | Job states timestamp | [optional] 
**BfBackfilledJobs** | Pointer to **int32** | Total backfilled jobs (since last slurm start) | [optional] 
**BfLastBackfilledJobs** | Pointer to **int32** | Total backfilled jobs (since last stats cycle start) | [optional] 
**BfBackfilledHetJobs** | Pointer to **int32** | Total backfilled heterogeneous job components | [optional] 
**BfCycleCounter** | Pointer to **int32** | Backfill Schedule Total cycles | [optional] 
**BfCycleMean** | Pointer to **int32** | Backfill Schedule Mean cycle | [optional] 
**BfCycleMax** | Pointer to **int32** | Backfill Schedule Max cycle time | [optional] 
**BfLastDepth** | Pointer to **int32** | Backfill Schedule Last depth cycle | [optional] 
**BfLastDepthTry** | Pointer to **int32** | Backfill Schedule Mean cycle (try sched) | [optional] 
**BfDepthMean** | Pointer to **int32** | Backfill Schedule Depth Mean | [optional] 
**BfDepthMeanTry** | Pointer to **int32** | Backfill Schedule Depth Mean (try sched) | [optional] 
**BfCycleLast** | Pointer to **int32** | Backfill Schedule Last cycle time | [optional] 
**BfQueueLen** | Pointer to **int32** | Backfill Schedule Last queue length | [optional] 
**BfQueueLenMean** | Pointer to **int32** | Backfill Schedule Mean queue length | [optional] 
**BfTableSize** | Pointer to **int32** | Backfill Schedule Last table size | [optional] 
**BfTableSizeMean** | Pointer to **int32** | Backfill Schedule Mean table size | [optional] 
**BfWhenLastCycle** | Pointer to **int32** | Last cycle timestamp | [optional] 
**BfActive** | Pointer to **bool** | Backfill Schedule currently active | [optional] 
**RpcsByMessageType** | Pointer to [**[]V0038DiagRpcm**](V0038DiagRpcm.md) | Remote Procedure Call statistics by message type | [optional] 
**RpcsByUser** | Pointer to [**[]V0038DiagRpcu**](V0038DiagRpcu.md) | Remote Procedure Call statistics by user | [optional] 

## Methods

### NewV0038DiagStatistics

`func NewV0038DiagStatistics() *V0038DiagStatistics`

NewV0038DiagStatistics instantiates a new V0038DiagStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038DiagStatisticsWithDefaults

`func NewV0038DiagStatisticsWithDefaults() *V0038DiagStatistics`

NewV0038DiagStatisticsWithDefaults instantiates a new V0038DiagStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartsPacked

`func (o *V0038DiagStatistics) GetPartsPacked() int32`

GetPartsPacked returns the PartsPacked field if non-nil, zero value otherwise.

### GetPartsPackedOk

`func (o *V0038DiagStatistics) GetPartsPackedOk() (*int32, bool)`

GetPartsPackedOk returns a tuple with the PartsPacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsPacked

`func (o *V0038DiagStatistics) SetPartsPacked(v int32)`

SetPartsPacked sets PartsPacked field to given value.

### HasPartsPacked

`func (o *V0038DiagStatistics) HasPartsPacked() bool`

HasPartsPacked returns a boolean if a field has been set.

### GetReqTime

`func (o *V0038DiagStatistics) GetReqTime() int32`

GetReqTime returns the ReqTime field if non-nil, zero value otherwise.

### GetReqTimeOk

`func (o *V0038DiagStatistics) GetReqTimeOk() (*int32, bool)`

GetReqTimeOk returns a tuple with the ReqTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTime

`func (o *V0038DiagStatistics) SetReqTime(v int32)`

SetReqTime sets ReqTime field to given value.

### HasReqTime

`func (o *V0038DiagStatistics) HasReqTime() bool`

HasReqTime returns a boolean if a field has been set.

### GetReqTimeStart

`func (o *V0038DiagStatistics) GetReqTimeStart() int32`

GetReqTimeStart returns the ReqTimeStart field if non-nil, zero value otherwise.

### GetReqTimeStartOk

`func (o *V0038DiagStatistics) GetReqTimeStartOk() (*int32, bool)`

GetReqTimeStartOk returns a tuple with the ReqTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTimeStart

`func (o *V0038DiagStatistics) SetReqTimeStart(v int32)`

SetReqTimeStart sets ReqTimeStart field to given value.

### HasReqTimeStart

`func (o *V0038DiagStatistics) HasReqTimeStart() bool`

HasReqTimeStart returns a boolean if a field has been set.

### GetServerThreadCount

`func (o *V0038DiagStatistics) GetServerThreadCount() int32`

GetServerThreadCount returns the ServerThreadCount field if non-nil, zero value otherwise.

### GetServerThreadCountOk

`func (o *V0038DiagStatistics) GetServerThreadCountOk() (*int32, bool)`

GetServerThreadCountOk returns a tuple with the ServerThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerThreadCount

`func (o *V0038DiagStatistics) SetServerThreadCount(v int32)`

SetServerThreadCount sets ServerThreadCount field to given value.

### HasServerThreadCount

`func (o *V0038DiagStatistics) HasServerThreadCount() bool`

HasServerThreadCount returns a boolean if a field has been set.

### GetAgentQueueSize

`func (o *V0038DiagStatistics) GetAgentQueueSize() int32`

GetAgentQueueSize returns the AgentQueueSize field if non-nil, zero value otherwise.

### GetAgentQueueSizeOk

`func (o *V0038DiagStatistics) GetAgentQueueSizeOk() (*int32, bool)`

GetAgentQueueSizeOk returns a tuple with the AgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentQueueSize

`func (o *V0038DiagStatistics) SetAgentQueueSize(v int32)`

SetAgentQueueSize sets AgentQueueSize field to given value.

### HasAgentQueueSize

`func (o *V0038DiagStatistics) HasAgentQueueSize() bool`

HasAgentQueueSize returns a boolean if a field has been set.

### GetAgentCount

`func (o *V0038DiagStatistics) GetAgentCount() int32`

GetAgentCount returns the AgentCount field if non-nil, zero value otherwise.

### GetAgentCountOk

`func (o *V0038DiagStatistics) GetAgentCountOk() (*int32, bool)`

GetAgentCountOk returns a tuple with the AgentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCount

`func (o *V0038DiagStatistics) SetAgentCount(v int32)`

SetAgentCount sets AgentCount field to given value.

### HasAgentCount

`func (o *V0038DiagStatistics) HasAgentCount() bool`

HasAgentCount returns a boolean if a field has been set.

### GetAgentThreadCount

`func (o *V0038DiagStatistics) GetAgentThreadCount() int32`

GetAgentThreadCount returns the AgentThreadCount field if non-nil, zero value otherwise.

### GetAgentThreadCountOk

`func (o *V0038DiagStatistics) GetAgentThreadCountOk() (*int32, bool)`

GetAgentThreadCountOk returns a tuple with the AgentThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentThreadCount

`func (o *V0038DiagStatistics) SetAgentThreadCount(v int32)`

SetAgentThreadCount sets AgentThreadCount field to given value.

### HasAgentThreadCount

`func (o *V0038DiagStatistics) HasAgentThreadCount() bool`

HasAgentThreadCount returns a boolean if a field has been set.

### GetDbdAgentQueueSize

`func (o *V0038DiagStatistics) GetDbdAgentQueueSize() int32`

GetDbdAgentQueueSize returns the DbdAgentQueueSize field if non-nil, zero value otherwise.

### GetDbdAgentQueueSizeOk

`func (o *V0038DiagStatistics) GetDbdAgentQueueSizeOk() (*int32, bool)`

GetDbdAgentQueueSizeOk returns a tuple with the DbdAgentQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbdAgentQueueSize

`func (o *V0038DiagStatistics) SetDbdAgentQueueSize(v int32)`

SetDbdAgentQueueSize sets DbdAgentQueueSize field to given value.

### HasDbdAgentQueueSize

`func (o *V0038DiagStatistics) HasDbdAgentQueueSize() bool`

HasDbdAgentQueueSize returns a boolean if a field has been set.

### GetGettimeofdayLatency

`func (o *V0038DiagStatistics) GetGettimeofdayLatency() int32`

GetGettimeofdayLatency returns the GettimeofdayLatency field if non-nil, zero value otherwise.

### GetGettimeofdayLatencyOk

`func (o *V0038DiagStatistics) GetGettimeofdayLatencyOk() (*int32, bool)`

GetGettimeofdayLatencyOk returns a tuple with the GettimeofdayLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGettimeofdayLatency

`func (o *V0038DiagStatistics) SetGettimeofdayLatency(v int32)`

SetGettimeofdayLatency sets GettimeofdayLatency field to given value.

### HasGettimeofdayLatency

`func (o *V0038DiagStatistics) HasGettimeofdayLatency() bool`

HasGettimeofdayLatency returns a boolean if a field has been set.

### GetScheduleCycleMax

`func (o *V0038DiagStatistics) GetScheduleCycleMax() int32`

GetScheduleCycleMax returns the ScheduleCycleMax field if non-nil, zero value otherwise.

### GetScheduleCycleMaxOk

`func (o *V0038DiagStatistics) GetScheduleCycleMaxOk() (*int32, bool)`

GetScheduleCycleMaxOk returns a tuple with the ScheduleCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMax

`func (o *V0038DiagStatistics) SetScheduleCycleMax(v int32)`

SetScheduleCycleMax sets ScheduleCycleMax field to given value.

### HasScheduleCycleMax

`func (o *V0038DiagStatistics) HasScheduleCycleMax() bool`

HasScheduleCycleMax returns a boolean if a field has been set.

### GetScheduleCycleLast

`func (o *V0038DiagStatistics) GetScheduleCycleLast() int32`

GetScheduleCycleLast returns the ScheduleCycleLast field if non-nil, zero value otherwise.

### GetScheduleCycleLastOk

`func (o *V0038DiagStatistics) GetScheduleCycleLastOk() (*int32, bool)`

GetScheduleCycleLastOk returns a tuple with the ScheduleCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleLast

`func (o *V0038DiagStatistics) SetScheduleCycleLast(v int32)`

SetScheduleCycleLast sets ScheduleCycleLast field to given value.

### HasScheduleCycleLast

`func (o *V0038DiagStatistics) HasScheduleCycleLast() bool`

HasScheduleCycleLast returns a boolean if a field has been set.

### GetScheduleCycleTotal

`func (o *V0038DiagStatistics) GetScheduleCycleTotal() int32`

GetScheduleCycleTotal returns the ScheduleCycleTotal field if non-nil, zero value otherwise.

### GetScheduleCycleTotalOk

`func (o *V0038DiagStatistics) GetScheduleCycleTotalOk() (*int32, bool)`

GetScheduleCycleTotalOk returns a tuple with the ScheduleCycleTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleTotal

`func (o *V0038DiagStatistics) SetScheduleCycleTotal(v int32)`

SetScheduleCycleTotal sets ScheduleCycleTotal field to given value.

### HasScheduleCycleTotal

`func (o *V0038DiagStatistics) HasScheduleCycleTotal() bool`

HasScheduleCycleTotal returns a boolean if a field has been set.

### GetScheduleCycleMean

`func (o *V0038DiagStatistics) GetScheduleCycleMean() int32`

GetScheduleCycleMean returns the ScheduleCycleMean field if non-nil, zero value otherwise.

### GetScheduleCycleMeanOk

`func (o *V0038DiagStatistics) GetScheduleCycleMeanOk() (*int32, bool)`

GetScheduleCycleMeanOk returns a tuple with the ScheduleCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMean

`func (o *V0038DiagStatistics) SetScheduleCycleMean(v int32)`

SetScheduleCycleMean sets ScheduleCycleMean field to given value.

### HasScheduleCycleMean

`func (o *V0038DiagStatistics) HasScheduleCycleMean() bool`

HasScheduleCycleMean returns a boolean if a field has been set.

### GetScheduleCycleMeanDepth

`func (o *V0038DiagStatistics) GetScheduleCycleMeanDepth() int32`

GetScheduleCycleMeanDepth returns the ScheduleCycleMeanDepth field if non-nil, zero value otherwise.

### GetScheduleCycleMeanDepthOk

`func (o *V0038DiagStatistics) GetScheduleCycleMeanDepthOk() (*int32, bool)`

GetScheduleCycleMeanDepthOk returns a tuple with the ScheduleCycleMeanDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCycleMeanDepth

`func (o *V0038DiagStatistics) SetScheduleCycleMeanDepth(v int32)`

SetScheduleCycleMeanDepth sets ScheduleCycleMeanDepth field to given value.

### HasScheduleCycleMeanDepth

`func (o *V0038DiagStatistics) HasScheduleCycleMeanDepth() bool`

HasScheduleCycleMeanDepth returns a boolean if a field has been set.

### GetScheduleCyclePerMinute

`func (o *V0038DiagStatistics) GetScheduleCyclePerMinute() int32`

GetScheduleCyclePerMinute returns the ScheduleCyclePerMinute field if non-nil, zero value otherwise.

### GetScheduleCyclePerMinuteOk

`func (o *V0038DiagStatistics) GetScheduleCyclePerMinuteOk() (*int32, bool)`

GetScheduleCyclePerMinuteOk returns a tuple with the ScheduleCyclePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCyclePerMinute

`func (o *V0038DiagStatistics) SetScheduleCyclePerMinute(v int32)`

SetScheduleCyclePerMinute sets ScheduleCyclePerMinute field to given value.

### HasScheduleCyclePerMinute

`func (o *V0038DiagStatistics) HasScheduleCyclePerMinute() bool`

HasScheduleCyclePerMinute returns a boolean if a field has been set.

### GetScheduleQueueLength

`func (o *V0038DiagStatistics) GetScheduleQueueLength() int32`

GetScheduleQueueLength returns the ScheduleQueueLength field if non-nil, zero value otherwise.

### GetScheduleQueueLengthOk

`func (o *V0038DiagStatistics) GetScheduleQueueLengthOk() (*int32, bool)`

GetScheduleQueueLengthOk returns a tuple with the ScheduleQueueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleQueueLength

`func (o *V0038DiagStatistics) SetScheduleQueueLength(v int32)`

SetScheduleQueueLength sets ScheduleQueueLength field to given value.

### HasScheduleQueueLength

`func (o *V0038DiagStatistics) HasScheduleQueueLength() bool`

HasScheduleQueueLength returns a boolean if a field has been set.

### GetJobsSubmitted

`func (o *V0038DiagStatistics) GetJobsSubmitted() int32`

GetJobsSubmitted returns the JobsSubmitted field if non-nil, zero value otherwise.

### GetJobsSubmittedOk

`func (o *V0038DiagStatistics) GetJobsSubmittedOk() (*int32, bool)`

GetJobsSubmittedOk returns a tuple with the JobsSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsSubmitted

`func (o *V0038DiagStatistics) SetJobsSubmitted(v int32)`

SetJobsSubmitted sets JobsSubmitted field to given value.

### HasJobsSubmitted

`func (o *V0038DiagStatistics) HasJobsSubmitted() bool`

HasJobsSubmitted returns a boolean if a field has been set.

### GetJobsStarted

`func (o *V0038DiagStatistics) GetJobsStarted() int32`

GetJobsStarted returns the JobsStarted field if non-nil, zero value otherwise.

### GetJobsStartedOk

`func (o *V0038DiagStatistics) GetJobsStartedOk() (*int32, bool)`

GetJobsStartedOk returns a tuple with the JobsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsStarted

`func (o *V0038DiagStatistics) SetJobsStarted(v int32)`

SetJobsStarted sets JobsStarted field to given value.

### HasJobsStarted

`func (o *V0038DiagStatistics) HasJobsStarted() bool`

HasJobsStarted returns a boolean if a field has been set.

### GetJobsCompleted

`func (o *V0038DiagStatistics) GetJobsCompleted() int32`

GetJobsCompleted returns the JobsCompleted field if non-nil, zero value otherwise.

### GetJobsCompletedOk

`func (o *V0038DiagStatistics) GetJobsCompletedOk() (*int32, bool)`

GetJobsCompletedOk returns a tuple with the JobsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCompleted

`func (o *V0038DiagStatistics) SetJobsCompleted(v int32)`

SetJobsCompleted sets JobsCompleted field to given value.

### HasJobsCompleted

`func (o *V0038DiagStatistics) HasJobsCompleted() bool`

HasJobsCompleted returns a boolean if a field has been set.

### GetJobsCanceled

`func (o *V0038DiagStatistics) GetJobsCanceled() int32`

GetJobsCanceled returns the JobsCanceled field if non-nil, zero value otherwise.

### GetJobsCanceledOk

`func (o *V0038DiagStatistics) GetJobsCanceledOk() (*int32, bool)`

GetJobsCanceledOk returns a tuple with the JobsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCanceled

`func (o *V0038DiagStatistics) SetJobsCanceled(v int32)`

SetJobsCanceled sets JobsCanceled field to given value.

### HasJobsCanceled

`func (o *V0038DiagStatistics) HasJobsCanceled() bool`

HasJobsCanceled returns a boolean if a field has been set.

### GetJobsFailed

`func (o *V0038DiagStatistics) GetJobsFailed() int32`

GetJobsFailed returns the JobsFailed field if non-nil, zero value otherwise.

### GetJobsFailedOk

`func (o *V0038DiagStatistics) GetJobsFailedOk() (*int32, bool)`

GetJobsFailedOk returns a tuple with the JobsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsFailed

`func (o *V0038DiagStatistics) SetJobsFailed(v int32)`

SetJobsFailed sets JobsFailed field to given value.

### HasJobsFailed

`func (o *V0038DiagStatistics) HasJobsFailed() bool`

HasJobsFailed returns a boolean if a field has been set.

### GetJobsPending

`func (o *V0038DiagStatistics) GetJobsPending() int32`

GetJobsPending returns the JobsPending field if non-nil, zero value otherwise.

### GetJobsPendingOk

`func (o *V0038DiagStatistics) GetJobsPendingOk() (*int32, bool)`

GetJobsPendingOk returns a tuple with the JobsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsPending

`func (o *V0038DiagStatistics) SetJobsPending(v int32)`

SetJobsPending sets JobsPending field to given value.

### HasJobsPending

`func (o *V0038DiagStatistics) HasJobsPending() bool`

HasJobsPending returns a boolean if a field has been set.

### GetJobsRunning

`func (o *V0038DiagStatistics) GetJobsRunning() int32`

GetJobsRunning returns the JobsRunning field if non-nil, zero value otherwise.

### GetJobsRunningOk

`func (o *V0038DiagStatistics) GetJobsRunningOk() (*int32, bool)`

GetJobsRunningOk returns a tuple with the JobsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsRunning

`func (o *V0038DiagStatistics) SetJobsRunning(v int32)`

SetJobsRunning sets JobsRunning field to given value.

### HasJobsRunning

`func (o *V0038DiagStatistics) HasJobsRunning() bool`

HasJobsRunning returns a boolean if a field has been set.

### GetJobStatesTs

`func (o *V0038DiagStatistics) GetJobStatesTs() int32`

GetJobStatesTs returns the JobStatesTs field if non-nil, zero value otherwise.

### GetJobStatesTsOk

`func (o *V0038DiagStatistics) GetJobStatesTsOk() (*int32, bool)`

GetJobStatesTsOk returns a tuple with the JobStatesTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatesTs

`func (o *V0038DiagStatistics) SetJobStatesTs(v int32)`

SetJobStatesTs sets JobStatesTs field to given value.

### HasJobStatesTs

`func (o *V0038DiagStatistics) HasJobStatesTs() bool`

HasJobStatesTs returns a boolean if a field has been set.

### GetBfBackfilledJobs

`func (o *V0038DiagStatistics) GetBfBackfilledJobs() int32`

GetBfBackfilledJobs returns the BfBackfilledJobs field if non-nil, zero value otherwise.

### GetBfBackfilledJobsOk

`func (o *V0038DiagStatistics) GetBfBackfilledJobsOk() (*int32, bool)`

GetBfBackfilledJobsOk returns a tuple with the BfBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledJobs

`func (o *V0038DiagStatistics) SetBfBackfilledJobs(v int32)`

SetBfBackfilledJobs sets BfBackfilledJobs field to given value.

### HasBfBackfilledJobs

`func (o *V0038DiagStatistics) HasBfBackfilledJobs() bool`

HasBfBackfilledJobs returns a boolean if a field has been set.

### GetBfLastBackfilledJobs

`func (o *V0038DiagStatistics) GetBfLastBackfilledJobs() int32`

GetBfLastBackfilledJobs returns the BfLastBackfilledJobs field if non-nil, zero value otherwise.

### GetBfLastBackfilledJobsOk

`func (o *V0038DiagStatistics) GetBfLastBackfilledJobsOk() (*int32, bool)`

GetBfLastBackfilledJobsOk returns a tuple with the BfLastBackfilledJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastBackfilledJobs

`func (o *V0038DiagStatistics) SetBfLastBackfilledJobs(v int32)`

SetBfLastBackfilledJobs sets BfLastBackfilledJobs field to given value.

### HasBfLastBackfilledJobs

`func (o *V0038DiagStatistics) HasBfLastBackfilledJobs() bool`

HasBfLastBackfilledJobs returns a boolean if a field has been set.

### GetBfBackfilledHetJobs

`func (o *V0038DiagStatistics) GetBfBackfilledHetJobs() int32`

GetBfBackfilledHetJobs returns the BfBackfilledHetJobs field if non-nil, zero value otherwise.

### GetBfBackfilledHetJobsOk

`func (o *V0038DiagStatistics) GetBfBackfilledHetJobsOk() (*int32, bool)`

GetBfBackfilledHetJobsOk returns a tuple with the BfBackfilledHetJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfBackfilledHetJobs

`func (o *V0038DiagStatistics) SetBfBackfilledHetJobs(v int32)`

SetBfBackfilledHetJobs sets BfBackfilledHetJobs field to given value.

### HasBfBackfilledHetJobs

`func (o *V0038DiagStatistics) HasBfBackfilledHetJobs() bool`

HasBfBackfilledHetJobs returns a boolean if a field has been set.

### GetBfCycleCounter

`func (o *V0038DiagStatistics) GetBfCycleCounter() int32`

GetBfCycleCounter returns the BfCycleCounter field if non-nil, zero value otherwise.

### GetBfCycleCounterOk

`func (o *V0038DiagStatistics) GetBfCycleCounterOk() (*int32, bool)`

GetBfCycleCounterOk returns a tuple with the BfCycleCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleCounter

`func (o *V0038DiagStatistics) SetBfCycleCounter(v int32)`

SetBfCycleCounter sets BfCycleCounter field to given value.

### HasBfCycleCounter

`func (o *V0038DiagStatistics) HasBfCycleCounter() bool`

HasBfCycleCounter returns a boolean if a field has been set.

### GetBfCycleMean

`func (o *V0038DiagStatistics) GetBfCycleMean() int32`

GetBfCycleMean returns the BfCycleMean field if non-nil, zero value otherwise.

### GetBfCycleMeanOk

`func (o *V0038DiagStatistics) GetBfCycleMeanOk() (*int32, bool)`

GetBfCycleMeanOk returns a tuple with the BfCycleMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMean

`func (o *V0038DiagStatistics) SetBfCycleMean(v int32)`

SetBfCycleMean sets BfCycleMean field to given value.

### HasBfCycleMean

`func (o *V0038DiagStatistics) HasBfCycleMean() bool`

HasBfCycleMean returns a boolean if a field has been set.

### GetBfCycleMax

`func (o *V0038DiagStatistics) GetBfCycleMax() int32`

GetBfCycleMax returns the BfCycleMax field if non-nil, zero value otherwise.

### GetBfCycleMaxOk

`func (o *V0038DiagStatistics) GetBfCycleMaxOk() (*int32, bool)`

GetBfCycleMaxOk returns a tuple with the BfCycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleMax

`func (o *V0038DiagStatistics) SetBfCycleMax(v int32)`

SetBfCycleMax sets BfCycleMax field to given value.

### HasBfCycleMax

`func (o *V0038DiagStatistics) HasBfCycleMax() bool`

HasBfCycleMax returns a boolean if a field has been set.

### GetBfLastDepth

`func (o *V0038DiagStatistics) GetBfLastDepth() int32`

GetBfLastDepth returns the BfLastDepth field if non-nil, zero value otherwise.

### GetBfLastDepthOk

`func (o *V0038DiagStatistics) GetBfLastDepthOk() (*int32, bool)`

GetBfLastDepthOk returns a tuple with the BfLastDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepth

`func (o *V0038DiagStatistics) SetBfLastDepth(v int32)`

SetBfLastDepth sets BfLastDepth field to given value.

### HasBfLastDepth

`func (o *V0038DiagStatistics) HasBfLastDepth() bool`

HasBfLastDepth returns a boolean if a field has been set.

### GetBfLastDepthTry

`func (o *V0038DiagStatistics) GetBfLastDepthTry() int32`

GetBfLastDepthTry returns the BfLastDepthTry field if non-nil, zero value otherwise.

### GetBfLastDepthTryOk

`func (o *V0038DiagStatistics) GetBfLastDepthTryOk() (*int32, bool)`

GetBfLastDepthTryOk returns a tuple with the BfLastDepthTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfLastDepthTry

`func (o *V0038DiagStatistics) SetBfLastDepthTry(v int32)`

SetBfLastDepthTry sets BfLastDepthTry field to given value.

### HasBfLastDepthTry

`func (o *V0038DiagStatistics) HasBfLastDepthTry() bool`

HasBfLastDepthTry returns a boolean if a field has been set.

### GetBfDepthMean

`func (o *V0038DiagStatistics) GetBfDepthMean() int32`

GetBfDepthMean returns the BfDepthMean field if non-nil, zero value otherwise.

### GetBfDepthMeanOk

`func (o *V0038DiagStatistics) GetBfDepthMeanOk() (*int32, bool)`

GetBfDepthMeanOk returns a tuple with the BfDepthMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMean

`func (o *V0038DiagStatistics) SetBfDepthMean(v int32)`

SetBfDepthMean sets BfDepthMean field to given value.

### HasBfDepthMean

`func (o *V0038DiagStatistics) HasBfDepthMean() bool`

HasBfDepthMean returns a boolean if a field has been set.

### GetBfDepthMeanTry

`func (o *V0038DiagStatistics) GetBfDepthMeanTry() int32`

GetBfDepthMeanTry returns the BfDepthMeanTry field if non-nil, zero value otherwise.

### GetBfDepthMeanTryOk

`func (o *V0038DiagStatistics) GetBfDepthMeanTryOk() (*int32, bool)`

GetBfDepthMeanTryOk returns a tuple with the BfDepthMeanTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfDepthMeanTry

`func (o *V0038DiagStatistics) SetBfDepthMeanTry(v int32)`

SetBfDepthMeanTry sets BfDepthMeanTry field to given value.

### HasBfDepthMeanTry

`func (o *V0038DiagStatistics) HasBfDepthMeanTry() bool`

HasBfDepthMeanTry returns a boolean if a field has been set.

### GetBfCycleLast

`func (o *V0038DiagStatistics) GetBfCycleLast() int32`

GetBfCycleLast returns the BfCycleLast field if non-nil, zero value otherwise.

### GetBfCycleLastOk

`func (o *V0038DiagStatistics) GetBfCycleLastOk() (*int32, bool)`

GetBfCycleLastOk returns a tuple with the BfCycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfCycleLast

`func (o *V0038DiagStatistics) SetBfCycleLast(v int32)`

SetBfCycleLast sets BfCycleLast field to given value.

### HasBfCycleLast

`func (o *V0038DiagStatistics) HasBfCycleLast() bool`

HasBfCycleLast returns a boolean if a field has been set.

### GetBfQueueLen

`func (o *V0038DiagStatistics) GetBfQueueLen() int32`

GetBfQueueLen returns the BfQueueLen field if non-nil, zero value otherwise.

### GetBfQueueLenOk

`func (o *V0038DiagStatistics) GetBfQueueLenOk() (*int32, bool)`

GetBfQueueLenOk returns a tuple with the BfQueueLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLen

`func (o *V0038DiagStatistics) SetBfQueueLen(v int32)`

SetBfQueueLen sets BfQueueLen field to given value.

### HasBfQueueLen

`func (o *V0038DiagStatistics) HasBfQueueLen() bool`

HasBfQueueLen returns a boolean if a field has been set.

### GetBfQueueLenMean

`func (o *V0038DiagStatistics) GetBfQueueLenMean() int32`

GetBfQueueLenMean returns the BfQueueLenMean field if non-nil, zero value otherwise.

### GetBfQueueLenMeanOk

`func (o *V0038DiagStatistics) GetBfQueueLenMeanOk() (*int32, bool)`

GetBfQueueLenMeanOk returns a tuple with the BfQueueLenMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfQueueLenMean

`func (o *V0038DiagStatistics) SetBfQueueLenMean(v int32)`

SetBfQueueLenMean sets BfQueueLenMean field to given value.

### HasBfQueueLenMean

`func (o *V0038DiagStatistics) HasBfQueueLenMean() bool`

HasBfQueueLenMean returns a boolean if a field has been set.

### GetBfTableSize

`func (o *V0038DiagStatistics) GetBfTableSize() int32`

GetBfTableSize returns the BfTableSize field if non-nil, zero value otherwise.

### GetBfTableSizeOk

`func (o *V0038DiagStatistics) GetBfTableSizeOk() (*int32, bool)`

GetBfTableSizeOk returns a tuple with the BfTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSize

`func (o *V0038DiagStatistics) SetBfTableSize(v int32)`

SetBfTableSize sets BfTableSize field to given value.

### HasBfTableSize

`func (o *V0038DiagStatistics) HasBfTableSize() bool`

HasBfTableSize returns a boolean if a field has been set.

### GetBfTableSizeMean

`func (o *V0038DiagStatistics) GetBfTableSizeMean() int32`

GetBfTableSizeMean returns the BfTableSizeMean field if non-nil, zero value otherwise.

### GetBfTableSizeMeanOk

`func (o *V0038DiagStatistics) GetBfTableSizeMeanOk() (*int32, bool)`

GetBfTableSizeMeanOk returns a tuple with the BfTableSizeMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfTableSizeMean

`func (o *V0038DiagStatistics) SetBfTableSizeMean(v int32)`

SetBfTableSizeMean sets BfTableSizeMean field to given value.

### HasBfTableSizeMean

`func (o *V0038DiagStatistics) HasBfTableSizeMean() bool`

HasBfTableSizeMean returns a boolean if a field has been set.

### GetBfWhenLastCycle

`func (o *V0038DiagStatistics) GetBfWhenLastCycle() int32`

GetBfWhenLastCycle returns the BfWhenLastCycle field if non-nil, zero value otherwise.

### GetBfWhenLastCycleOk

`func (o *V0038DiagStatistics) GetBfWhenLastCycleOk() (*int32, bool)`

GetBfWhenLastCycleOk returns a tuple with the BfWhenLastCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfWhenLastCycle

`func (o *V0038DiagStatistics) SetBfWhenLastCycle(v int32)`

SetBfWhenLastCycle sets BfWhenLastCycle field to given value.

### HasBfWhenLastCycle

`func (o *V0038DiagStatistics) HasBfWhenLastCycle() bool`

HasBfWhenLastCycle returns a boolean if a field has been set.

### GetBfActive

`func (o *V0038DiagStatistics) GetBfActive() bool`

GetBfActive returns the BfActive field if non-nil, zero value otherwise.

### GetBfActiveOk

`func (o *V0038DiagStatistics) GetBfActiveOk() (*bool, bool)`

GetBfActiveOk returns a tuple with the BfActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfActive

`func (o *V0038DiagStatistics) SetBfActive(v bool)`

SetBfActive sets BfActive field to given value.

### HasBfActive

`func (o *V0038DiagStatistics) HasBfActive() bool`

HasBfActive returns a boolean if a field has been set.

### GetRpcsByMessageType

`func (o *V0038DiagStatistics) GetRpcsByMessageType() []V0038DiagRpcm`

GetRpcsByMessageType returns the RpcsByMessageType field if non-nil, zero value otherwise.

### GetRpcsByMessageTypeOk

`func (o *V0038DiagStatistics) GetRpcsByMessageTypeOk() (*[]V0038DiagRpcm, bool)`

GetRpcsByMessageTypeOk returns a tuple with the RpcsByMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByMessageType

`func (o *V0038DiagStatistics) SetRpcsByMessageType(v []V0038DiagRpcm)`

SetRpcsByMessageType sets RpcsByMessageType field to given value.

### HasRpcsByMessageType

`func (o *V0038DiagStatistics) HasRpcsByMessageType() bool`

HasRpcsByMessageType returns a boolean if a field has been set.

### GetRpcsByUser

`func (o *V0038DiagStatistics) GetRpcsByUser() []V0038DiagRpcu`

GetRpcsByUser returns the RpcsByUser field if non-nil, zero value otherwise.

### GetRpcsByUserOk

`func (o *V0038DiagStatistics) GetRpcsByUserOk() (*[]V0038DiagRpcu, bool)`

GetRpcsByUserOk returns a tuple with the RpcsByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcsByUser

`func (o *V0038DiagStatistics) SetRpcsByUser(v []V0038DiagRpcu)`

SetRpcsByUser sets RpcsByUser field to given value.

### HasRpcsByUser

`func (o *V0038DiagStatistics) HasRpcsByUser() bool`

HasRpcsByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


