# V0040AssocShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Association account (if assigned) | [optional] 
**Cluster** | Pointer to **string** | Association cluster (if assigned) | [optional] 
**Partition** | Pointer to **string** | Association partition (if assigned) | [optional] 
**User** | **string** | Assocation user (if assigned) | 
**Id** | Pointer to **int32** | Numeric Association ID (if known) | [optional] 

## Methods

### NewV0040AssocShort

`func NewV0040AssocShort(user string, ) *V0040AssocShort`

NewV0040AssocShort instantiates a new V0040AssocShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040AssocShortWithDefaults

`func NewV0040AssocShortWithDefaults() *V0040AssocShort`

NewV0040AssocShortWithDefaults instantiates a new V0040AssocShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0040AssocShort) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0040AssocShort) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0040AssocShort) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0040AssocShort) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCluster

`func (o *V0040AssocShort) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0040AssocShort) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0040AssocShort) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0040AssocShort) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetPartition

`func (o *V0040AssocShort) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0040AssocShort) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0040AssocShort) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0040AssocShort) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetUser

`func (o *V0040AssocShort) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0040AssocShort) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0040AssocShort) SetUser(v string)`

SetUser sets User field to given value.


### GetId

`func (o *V0040AssocShort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0040AssocShort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0040AssocShort) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0040AssocShort) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


