# Dbv0038Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account charged by job | [optional] 
**Comment** | Pointer to [**Dbv0038JobComment**](Dbv0038JobComment.md) |  | [optional] 
**AllocationNodes** | Pointer to **string** | Nodes allocated to job | [optional] 
**Array** | Pointer to [**Dbv0038JobArray**](Dbv0038JobArray.md) |  | [optional] 
**Time** | Pointer to [**Dbv0038JobTime**](Dbv0038JobTime.md) |  | [optional] 
**Association** | Pointer to [**Dbv0038AssociationShortInfo**](Dbv0038AssociationShortInfo.md) |  | [optional] 
**Cluster** | Pointer to **string** | Assigned cluster | [optional] 
**Constraints** | Pointer to **string** | Constraints on job | [optional] 
**DerivedExitCode** | Pointer to [**Dbv0038JobExitCode**](Dbv0038JobExitCode.md) |  | [optional] 
**ExitCode** | Pointer to [**Dbv0038JobExitCode**](Dbv0038JobExitCode.md) |  | [optional] 
**Flags** | Pointer to **[]string** | List of properties of job | [optional] 
**Group** | Pointer to **string** | User&#39;s group to run job | [optional] 
**Het** | Pointer to [**Dbv0038JobHet**](Dbv0038JobHet.md) |  | [optional] 
**JobId** | Pointer to **int32** | Job id | [optional] 
**Name** | Pointer to **string** | Assigned job name | [optional] 
**Mcs** | Pointer to [**Dbv0038JobMcs**](Dbv0038JobMcs.md) |  | [optional] 
**Nodes** | Pointer to **string** | List of nodes allocated for job | [optional] 
**Partition** | Pointer to **string** | Assigned job&#39;s partition | [optional] 
**Priority** | Pointer to **int32** | Priority | [optional] 
**Qos** | Pointer to **string** | Assigned qos name | [optional] 
**Required** | Pointer to [**Dbv0038JobRequired**](Dbv0038JobRequired.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** | User who requested job killed | [optional] 
**Reservation** | Pointer to [**Dbv0038JobReservation**](Dbv0038JobReservation.md) |  | [optional] 
**State** | Pointer to [**Dbv0038JobState**](Dbv0038JobState.md) |  | [optional] 
**Steps** | Pointer to [**[]Dbv0038JobStep**](Dbv0038JobStep.md) | Job step description | [optional] 
**Tres** | Pointer to [**Dbv0038JobTres**](Dbv0038JobTres.md) |  | [optional] 
**User** | Pointer to **string** | Job user | [optional] 
**Wckey** | Pointer to [**Dbv0038JobWckey**](Dbv0038JobWckey.md) |  | [optional] 
**WorkingDirectory** | Pointer to **string** | Directory where job was initially started | [optional] 
**Container** | Pointer to **string** | absolute path to OCI container bundle | [optional] 

## Methods

### NewDbv0038Job

`func NewDbv0038Job() *Dbv0038Job`

NewDbv0038Job instantiates a new Dbv0038Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobWithDefaults

`func NewDbv0038JobWithDefaults() *Dbv0038Job`

NewDbv0038JobWithDefaults instantiates a new Dbv0038Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Dbv0038Job) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Dbv0038Job) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Dbv0038Job) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Dbv0038Job) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetComment

`func (o *Dbv0038Job) GetComment() Dbv0038JobComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Dbv0038Job) GetCommentOk() (*Dbv0038JobComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Dbv0038Job) SetComment(v Dbv0038JobComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Dbv0038Job) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAllocationNodes

`func (o *Dbv0038Job) GetAllocationNodes() string`

GetAllocationNodes returns the AllocationNodes field if non-nil, zero value otherwise.

### GetAllocationNodesOk

`func (o *Dbv0038Job) GetAllocationNodesOk() (*string, bool)`

GetAllocationNodesOk returns a tuple with the AllocationNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodes

`func (o *Dbv0038Job) SetAllocationNodes(v string)`

SetAllocationNodes sets AllocationNodes field to given value.

### HasAllocationNodes

`func (o *Dbv0038Job) HasAllocationNodes() bool`

HasAllocationNodes returns a boolean if a field has been set.

### GetArray

`func (o *Dbv0038Job) GetArray() Dbv0038JobArray`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *Dbv0038Job) GetArrayOk() (*Dbv0038JobArray, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *Dbv0038Job) SetArray(v Dbv0038JobArray)`

SetArray sets Array field to given value.

### HasArray

`func (o *Dbv0038Job) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetTime

`func (o *Dbv0038Job) GetTime() Dbv0038JobTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Dbv0038Job) GetTimeOk() (*Dbv0038JobTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Dbv0038Job) SetTime(v Dbv0038JobTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *Dbv0038Job) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAssociation

`func (o *Dbv0038Job) GetAssociation() Dbv0038AssociationShortInfo`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *Dbv0038Job) GetAssociationOk() (*Dbv0038AssociationShortInfo, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *Dbv0038Job) SetAssociation(v Dbv0038AssociationShortInfo)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *Dbv0038Job) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCluster

`func (o *Dbv0038Job) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Dbv0038Job) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Dbv0038Job) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Dbv0038Job) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConstraints

`func (o *Dbv0038Job) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Dbv0038Job) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Dbv0038Job) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *Dbv0038Job) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *Dbv0038Job) GetDerivedExitCode() Dbv0038JobExitCode`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *Dbv0038Job) GetDerivedExitCodeOk() (*Dbv0038JobExitCode, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *Dbv0038Job) SetDerivedExitCode(v Dbv0038JobExitCode)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *Dbv0038Job) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetExitCode

`func (o *Dbv0038Job) GetExitCode() Dbv0038JobExitCode`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Dbv0038Job) GetExitCodeOk() (*Dbv0038JobExitCode, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Dbv0038Job) SetExitCode(v Dbv0038JobExitCode)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Dbv0038Job) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetFlags

`func (o *Dbv0038Job) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbv0038Job) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Dbv0038Job) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbv0038Job) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroup

`func (o *Dbv0038Job) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Dbv0038Job) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Dbv0038Job) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Dbv0038Job) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHet

`func (o *Dbv0038Job) GetHet() Dbv0038JobHet`

GetHet returns the Het field if non-nil, zero value otherwise.

### GetHetOk

`func (o *Dbv0038Job) GetHetOk() (*Dbv0038JobHet, bool)`

GetHetOk returns a tuple with the Het field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHet

`func (o *Dbv0038Job) SetHet(v Dbv0038JobHet)`

SetHet sets Het field to given value.

### HasHet

`func (o *Dbv0038Job) HasHet() bool`

HasHet returns a boolean if a field has been set.

### GetJobId

`func (o *Dbv0038Job) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Dbv0038Job) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Dbv0038Job) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *Dbv0038Job) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetName

`func (o *Dbv0038Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0038Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dbv0038Job) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0038Job) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMcs

`func (o *Dbv0038Job) GetMcs() Dbv0038JobMcs`

GetMcs returns the Mcs field if non-nil, zero value otherwise.

### GetMcsOk

`func (o *Dbv0038Job) GetMcsOk() (*Dbv0038JobMcs, bool)`

GetMcsOk returns a tuple with the Mcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcs

`func (o *Dbv0038Job) SetMcs(v Dbv0038JobMcs)`

SetMcs sets Mcs field to given value.

### HasMcs

`func (o *Dbv0038Job) HasMcs() bool`

HasMcs returns a boolean if a field has been set.

### GetNodes

`func (o *Dbv0038Job) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Dbv0038Job) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Dbv0038Job) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Dbv0038Job) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPartition

`func (o *Dbv0038Job) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Dbv0038Job) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Dbv0038Job) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *Dbv0038Job) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *Dbv0038Job) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Dbv0038Job) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Dbv0038Job) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Dbv0038Job) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0038Job) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0038Job) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0038Job) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0038Job) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetRequired

`func (o *Dbv0038Job) GetRequired() Dbv0038JobRequired`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Dbv0038Job) GetRequiredOk() (*Dbv0038JobRequired, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Dbv0038Job) SetRequired(v Dbv0038JobRequired)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Dbv0038Job) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *Dbv0038Job) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *Dbv0038Job) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *Dbv0038Job) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *Dbv0038Job) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetReservation

`func (o *Dbv0038Job) GetReservation() Dbv0038JobReservation`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Dbv0038Job) GetReservationOk() (*Dbv0038JobReservation, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *Dbv0038Job) SetReservation(v Dbv0038JobReservation)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *Dbv0038Job) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetState

`func (o *Dbv0038Job) GetState() Dbv0038JobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Dbv0038Job) GetStateOk() (*Dbv0038JobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Dbv0038Job) SetState(v Dbv0038JobState)`

SetState sets State field to given value.

### HasState

`func (o *Dbv0038Job) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSteps

`func (o *Dbv0038Job) GetSteps() []Dbv0038JobStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Dbv0038Job) GetStepsOk() (*[]Dbv0038JobStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Dbv0038Job) SetSteps(v []Dbv0038JobStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Dbv0038Job) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0038Job) GetTres() Dbv0038JobTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0038Job) GetTresOk() (*Dbv0038JobTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0038Job) SetTres(v Dbv0038JobTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0038Job) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetUser

`func (o *Dbv0038Job) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Dbv0038Job) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Dbv0038Job) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Dbv0038Job) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWckey

`func (o *Dbv0038Job) GetWckey() Dbv0038JobWckey`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *Dbv0038Job) GetWckeyOk() (*Dbv0038JobWckey, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *Dbv0038Job) SetWckey(v Dbv0038JobWckey)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *Dbv0038Job) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *Dbv0038Job) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *Dbv0038Job) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *Dbv0038Job) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *Dbv0038Job) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetContainer

`func (o *Dbv0038Job) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *Dbv0038Job) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *Dbv0038Job) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *Dbv0038Job) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


