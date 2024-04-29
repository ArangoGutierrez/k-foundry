# V0040AssocRecSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment for the association | [optional] 
**Defaultqos** | Pointer to **string** | Which QOS id is this association default | [optional] 
**Grpjobs** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Grpjobsaccrue** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Grpsubmitjobs** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Grptres** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Grptresmins** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Grptresrunmins** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Grpwall** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Maxjobs** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Maxjobsaccrue** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Maxsubmitjobs** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Maxtresminsperjob** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Maxtresrunmins** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Maxtresperjob** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Maxtrespernode** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 
**Maxwalldurationperjob** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Minpriothresh** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Parent** | Pointer to **string** | Name of parent account | [optional] 
**Priority** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**Qoslevel** | Pointer to **[]string** | List of QOS names | [optional] 
**Fairshare** | Pointer to **int32** | Number of shares allocated to this association | [optional] 

## Methods

### NewV0040AssocRecSet

`func NewV0040AssocRecSet() *V0040AssocRecSet`

NewV0040AssocRecSet instantiates a new V0040AssocRecSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040AssocRecSetWithDefaults

`func NewV0040AssocRecSetWithDefaults() *V0040AssocRecSet`

NewV0040AssocRecSetWithDefaults instantiates a new V0040AssocRecSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V0040AssocRecSet) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0040AssocRecSet) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0040AssocRecSet) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0040AssocRecSet) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultqos

`func (o *V0040AssocRecSet) GetDefaultqos() string`

GetDefaultqos returns the Defaultqos field if non-nil, zero value otherwise.

### GetDefaultqosOk

`func (o *V0040AssocRecSet) GetDefaultqosOk() (*string, bool)`

GetDefaultqosOk returns a tuple with the Defaultqos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultqos

`func (o *V0040AssocRecSet) SetDefaultqos(v string)`

SetDefaultqos sets Defaultqos field to given value.

### HasDefaultqos

`func (o *V0040AssocRecSet) HasDefaultqos() bool`

HasDefaultqos returns a boolean if a field has been set.

### GetGrpjobs

`func (o *V0040AssocRecSet) GetGrpjobs() V0040Uint32NoVal`

GetGrpjobs returns the Grpjobs field if non-nil, zero value otherwise.

### GetGrpjobsOk

`func (o *V0040AssocRecSet) GetGrpjobsOk() (*V0040Uint32NoVal, bool)`

GetGrpjobsOk returns a tuple with the Grpjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobs

`func (o *V0040AssocRecSet) SetGrpjobs(v V0040Uint32NoVal)`

SetGrpjobs sets Grpjobs field to given value.

### HasGrpjobs

`func (o *V0040AssocRecSet) HasGrpjobs() bool`

HasGrpjobs returns a boolean if a field has been set.

### GetGrpjobsaccrue

`func (o *V0040AssocRecSet) GetGrpjobsaccrue() V0040Uint32NoVal`

GetGrpjobsaccrue returns the Grpjobsaccrue field if non-nil, zero value otherwise.

### GetGrpjobsaccrueOk

`func (o *V0040AssocRecSet) GetGrpjobsaccrueOk() (*V0040Uint32NoVal, bool)`

GetGrpjobsaccrueOk returns a tuple with the Grpjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobsaccrue

`func (o *V0040AssocRecSet) SetGrpjobsaccrue(v V0040Uint32NoVal)`

SetGrpjobsaccrue sets Grpjobsaccrue field to given value.

### HasGrpjobsaccrue

`func (o *V0040AssocRecSet) HasGrpjobsaccrue() bool`

HasGrpjobsaccrue returns a boolean if a field has been set.

### GetGrpsubmitjobs

`func (o *V0040AssocRecSet) GetGrpsubmitjobs() V0040Uint32NoVal`

GetGrpsubmitjobs returns the Grpsubmitjobs field if non-nil, zero value otherwise.

### GetGrpsubmitjobsOk

`func (o *V0040AssocRecSet) GetGrpsubmitjobsOk() (*V0040Uint32NoVal, bool)`

GetGrpsubmitjobsOk returns a tuple with the Grpsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpsubmitjobs

`func (o *V0040AssocRecSet) SetGrpsubmitjobs(v V0040Uint32NoVal)`

SetGrpsubmitjobs sets Grpsubmitjobs field to given value.

### HasGrpsubmitjobs

`func (o *V0040AssocRecSet) HasGrpsubmitjobs() bool`

HasGrpsubmitjobs returns a boolean if a field has been set.

### GetGrptres

`func (o *V0040AssocRecSet) GetGrptres() []V0040Tres`

GetGrptres returns the Grptres field if non-nil, zero value otherwise.

### GetGrptresOk

`func (o *V0040AssocRecSet) GetGrptresOk() (*[]V0040Tres, bool)`

GetGrptresOk returns a tuple with the Grptres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptres

`func (o *V0040AssocRecSet) SetGrptres(v []V0040Tres)`

SetGrptres sets Grptres field to given value.

### HasGrptres

`func (o *V0040AssocRecSet) HasGrptres() bool`

HasGrptres returns a boolean if a field has been set.

### GetGrptresmins

`func (o *V0040AssocRecSet) GetGrptresmins() []V0040Tres`

GetGrptresmins returns the Grptresmins field if non-nil, zero value otherwise.

### GetGrptresminsOk

`func (o *V0040AssocRecSet) GetGrptresminsOk() (*[]V0040Tres, bool)`

GetGrptresminsOk returns a tuple with the Grptresmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresmins

`func (o *V0040AssocRecSet) SetGrptresmins(v []V0040Tres)`

SetGrptresmins sets Grptresmins field to given value.

### HasGrptresmins

`func (o *V0040AssocRecSet) HasGrptresmins() bool`

HasGrptresmins returns a boolean if a field has been set.

### GetGrptresrunmins

`func (o *V0040AssocRecSet) GetGrptresrunmins() []V0040Tres`

GetGrptresrunmins returns the Grptresrunmins field if non-nil, zero value otherwise.

### GetGrptresrunminsOk

`func (o *V0040AssocRecSet) GetGrptresrunminsOk() (*[]V0040Tres, bool)`

GetGrptresrunminsOk returns a tuple with the Grptresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresrunmins

`func (o *V0040AssocRecSet) SetGrptresrunmins(v []V0040Tres)`

SetGrptresrunmins sets Grptresrunmins field to given value.

### HasGrptresrunmins

`func (o *V0040AssocRecSet) HasGrptresrunmins() bool`

HasGrptresrunmins returns a boolean if a field has been set.

### GetGrpwall

`func (o *V0040AssocRecSet) GetGrpwall() V0040Uint32NoVal`

GetGrpwall returns the Grpwall field if non-nil, zero value otherwise.

### GetGrpwallOk

`func (o *V0040AssocRecSet) GetGrpwallOk() (*V0040Uint32NoVal, bool)`

GetGrpwallOk returns a tuple with the Grpwall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpwall

`func (o *V0040AssocRecSet) SetGrpwall(v V0040Uint32NoVal)`

SetGrpwall sets Grpwall field to given value.

### HasGrpwall

`func (o *V0040AssocRecSet) HasGrpwall() bool`

HasGrpwall returns a boolean if a field has been set.

### GetMaxjobs

`func (o *V0040AssocRecSet) GetMaxjobs() V0040Uint32NoVal`

GetMaxjobs returns the Maxjobs field if non-nil, zero value otherwise.

### GetMaxjobsOk

`func (o *V0040AssocRecSet) GetMaxjobsOk() (*V0040Uint32NoVal, bool)`

GetMaxjobsOk returns a tuple with the Maxjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobs

`func (o *V0040AssocRecSet) SetMaxjobs(v V0040Uint32NoVal)`

SetMaxjobs sets Maxjobs field to given value.

### HasMaxjobs

`func (o *V0040AssocRecSet) HasMaxjobs() bool`

HasMaxjobs returns a boolean if a field has been set.

### GetMaxjobsaccrue

`func (o *V0040AssocRecSet) GetMaxjobsaccrue() V0040Uint32NoVal`

GetMaxjobsaccrue returns the Maxjobsaccrue field if non-nil, zero value otherwise.

### GetMaxjobsaccrueOk

`func (o *V0040AssocRecSet) GetMaxjobsaccrueOk() (*V0040Uint32NoVal, bool)`

GetMaxjobsaccrueOk returns a tuple with the Maxjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobsaccrue

`func (o *V0040AssocRecSet) SetMaxjobsaccrue(v V0040Uint32NoVal)`

SetMaxjobsaccrue sets Maxjobsaccrue field to given value.

### HasMaxjobsaccrue

`func (o *V0040AssocRecSet) HasMaxjobsaccrue() bool`

HasMaxjobsaccrue returns a boolean if a field has been set.

### GetMaxsubmitjobs

`func (o *V0040AssocRecSet) GetMaxsubmitjobs() V0040Uint32NoVal`

GetMaxsubmitjobs returns the Maxsubmitjobs field if non-nil, zero value otherwise.

### GetMaxsubmitjobsOk

`func (o *V0040AssocRecSet) GetMaxsubmitjobsOk() (*V0040Uint32NoVal, bool)`

GetMaxsubmitjobsOk returns a tuple with the Maxsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxsubmitjobs

`func (o *V0040AssocRecSet) SetMaxsubmitjobs(v V0040Uint32NoVal)`

SetMaxsubmitjobs sets Maxsubmitjobs field to given value.

### HasMaxsubmitjobs

`func (o *V0040AssocRecSet) HasMaxsubmitjobs() bool`

HasMaxsubmitjobs returns a boolean if a field has been set.

### GetMaxtresminsperjob

`func (o *V0040AssocRecSet) GetMaxtresminsperjob() []V0040Tres`

GetMaxtresminsperjob returns the Maxtresminsperjob field if non-nil, zero value otherwise.

### GetMaxtresminsperjobOk

`func (o *V0040AssocRecSet) GetMaxtresminsperjobOk() (*[]V0040Tres, bool)`

GetMaxtresminsperjobOk returns a tuple with the Maxtresminsperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresminsperjob

`func (o *V0040AssocRecSet) SetMaxtresminsperjob(v []V0040Tres)`

SetMaxtresminsperjob sets Maxtresminsperjob field to given value.

### HasMaxtresminsperjob

`func (o *V0040AssocRecSet) HasMaxtresminsperjob() bool`

HasMaxtresminsperjob returns a boolean if a field has been set.

### GetMaxtresrunmins

`func (o *V0040AssocRecSet) GetMaxtresrunmins() []V0040Tres`

GetMaxtresrunmins returns the Maxtresrunmins field if non-nil, zero value otherwise.

### GetMaxtresrunminsOk

`func (o *V0040AssocRecSet) GetMaxtresrunminsOk() (*[]V0040Tres, bool)`

GetMaxtresrunminsOk returns a tuple with the Maxtresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresrunmins

`func (o *V0040AssocRecSet) SetMaxtresrunmins(v []V0040Tres)`

SetMaxtresrunmins sets Maxtresrunmins field to given value.

### HasMaxtresrunmins

`func (o *V0040AssocRecSet) HasMaxtresrunmins() bool`

HasMaxtresrunmins returns a boolean if a field has been set.

### GetMaxtresperjob

`func (o *V0040AssocRecSet) GetMaxtresperjob() []V0040Tres`

GetMaxtresperjob returns the Maxtresperjob field if non-nil, zero value otherwise.

### GetMaxtresperjobOk

`func (o *V0040AssocRecSet) GetMaxtresperjobOk() (*[]V0040Tres, bool)`

GetMaxtresperjobOk returns a tuple with the Maxtresperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresperjob

`func (o *V0040AssocRecSet) SetMaxtresperjob(v []V0040Tres)`

SetMaxtresperjob sets Maxtresperjob field to given value.

### HasMaxtresperjob

`func (o *V0040AssocRecSet) HasMaxtresperjob() bool`

HasMaxtresperjob returns a boolean if a field has been set.

### GetMaxtrespernode

`func (o *V0040AssocRecSet) GetMaxtrespernode() []V0040Tres`

GetMaxtrespernode returns the Maxtrespernode field if non-nil, zero value otherwise.

### GetMaxtrespernodeOk

`func (o *V0040AssocRecSet) GetMaxtrespernodeOk() (*[]V0040Tres, bool)`

GetMaxtrespernodeOk returns a tuple with the Maxtrespernode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtrespernode

`func (o *V0040AssocRecSet) SetMaxtrespernode(v []V0040Tres)`

SetMaxtrespernode sets Maxtrespernode field to given value.

### HasMaxtrespernode

`func (o *V0040AssocRecSet) HasMaxtrespernode() bool`

HasMaxtrespernode returns a boolean if a field has been set.

### GetMaxwalldurationperjob

`func (o *V0040AssocRecSet) GetMaxwalldurationperjob() V0040Uint32NoVal`

GetMaxwalldurationperjob returns the Maxwalldurationperjob field if non-nil, zero value otherwise.

### GetMaxwalldurationperjobOk

`func (o *V0040AssocRecSet) GetMaxwalldurationperjobOk() (*V0040Uint32NoVal, bool)`

GetMaxwalldurationperjobOk returns a tuple with the Maxwalldurationperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxwalldurationperjob

`func (o *V0040AssocRecSet) SetMaxwalldurationperjob(v V0040Uint32NoVal)`

SetMaxwalldurationperjob sets Maxwalldurationperjob field to given value.

### HasMaxwalldurationperjob

`func (o *V0040AssocRecSet) HasMaxwalldurationperjob() bool`

HasMaxwalldurationperjob returns a boolean if a field has been set.

### GetMinpriothresh

`func (o *V0040AssocRecSet) GetMinpriothresh() V0040Uint32NoVal`

GetMinpriothresh returns the Minpriothresh field if non-nil, zero value otherwise.

### GetMinpriothreshOk

`func (o *V0040AssocRecSet) GetMinpriothreshOk() (*V0040Uint32NoVal, bool)`

GetMinpriothreshOk returns a tuple with the Minpriothresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinpriothresh

`func (o *V0040AssocRecSet) SetMinpriothresh(v V0040Uint32NoVal)`

SetMinpriothresh sets Minpriothresh field to given value.

### HasMinpriothresh

`func (o *V0040AssocRecSet) HasMinpriothresh() bool`

HasMinpriothresh returns a boolean if a field has been set.

### GetParent

`func (o *V0040AssocRecSet) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *V0040AssocRecSet) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *V0040AssocRecSet) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *V0040AssocRecSet) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPriority

`func (o *V0040AssocRecSet) GetPriority() V0040Uint32NoVal`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0040AssocRecSet) GetPriorityOk() (*V0040Uint32NoVal, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0040AssocRecSet) SetPriority(v V0040Uint32NoVal)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0040AssocRecSet) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQoslevel

`func (o *V0040AssocRecSet) GetQoslevel() []string`

GetQoslevel returns the Qoslevel field if non-nil, zero value otherwise.

### GetQoslevelOk

`func (o *V0040AssocRecSet) GetQoslevelOk() (*[]string, bool)`

GetQoslevelOk returns a tuple with the Qoslevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoslevel

`func (o *V0040AssocRecSet) SetQoslevel(v []string)`

SetQoslevel sets Qoslevel field to given value.

### HasQoslevel

`func (o *V0040AssocRecSet) HasQoslevel() bool`

HasQoslevel returns a boolean if a field has been set.

### GetFairshare

`func (o *V0040AssocRecSet) GetFairshare() int32`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *V0040AssocRecSet) GetFairshareOk() (*int32, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *V0040AssocRecSet) SetFairshare(v int32)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *V0040AssocRecSet) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


