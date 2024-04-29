# V0040OpenapiMetaSlurm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**V0040OpenapiMetaSlurmVersion**](V0040OpenapiMetaSlurmVersion.md) |  | [optional] 
**Release** | Pointer to **string** | Slurm release string | [optional] 
**Cluster** | Pointer to **string** | Slurm cluster name | [optional] 

## Methods

### NewV0040OpenapiMetaSlurm

`func NewV0040OpenapiMetaSlurm() *V0040OpenapiMetaSlurm`

NewV0040OpenapiMetaSlurm instantiates a new V0040OpenapiMetaSlurm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiMetaSlurmWithDefaults

`func NewV0040OpenapiMetaSlurmWithDefaults() *V0040OpenapiMetaSlurm`

NewV0040OpenapiMetaSlurmWithDefaults instantiates a new V0040OpenapiMetaSlurm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *V0040OpenapiMetaSlurm) GetVersion() V0040OpenapiMetaSlurmVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V0040OpenapiMetaSlurm) GetVersionOk() (*V0040OpenapiMetaSlurmVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V0040OpenapiMetaSlurm) SetVersion(v V0040OpenapiMetaSlurmVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V0040OpenapiMetaSlurm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *V0040OpenapiMetaSlurm) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *V0040OpenapiMetaSlurm) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *V0040OpenapiMetaSlurm) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *V0040OpenapiMetaSlurm) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetCluster

`func (o *V0040OpenapiMetaSlurm) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0040OpenapiMetaSlurm) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0040OpenapiMetaSlurm) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0040OpenapiMetaSlurm) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


