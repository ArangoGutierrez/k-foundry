# Dbv0038Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugin** | Pointer to [**Dbv0039MetaPlugin**](Dbv0039MetaPlugin.md) |  | [optional] 
**Slurm** | Pointer to [**Dbv0038MetaSlurm**](Dbv0038MetaSlurm.md) |  | [optional] 

## Methods

### NewDbv0038Meta

`func NewDbv0038Meta() *Dbv0038Meta`

NewDbv0038Meta instantiates a new Dbv0038Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038MetaWithDefaults

`func NewDbv0038MetaWithDefaults() *Dbv0038Meta`

NewDbv0038MetaWithDefaults instantiates a new Dbv0038Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugin

`func (o *Dbv0038Meta) GetPlugin() Dbv0039MetaPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *Dbv0038Meta) GetPluginOk() (*Dbv0039MetaPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *Dbv0038Meta) SetPlugin(v Dbv0039MetaPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *Dbv0038Meta) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetSlurm

`func (o *Dbv0038Meta) GetSlurm() Dbv0038MetaSlurm`

GetSlurm returns the Slurm field if non-nil, zero value otherwise.

### GetSlurmOk

`func (o *Dbv0038Meta) GetSlurmOk() (*Dbv0038MetaSlurm, bool)`

GetSlurmOk returns a tuple with the Slurm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurm

`func (o *Dbv0038Meta) SetSlurm(v Dbv0038MetaSlurm)`

SetSlurm sets Slurm field to given value.

### HasSlurm

`func (o *Dbv0038Meta) HasSlurm() bool`

HasSlurm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


