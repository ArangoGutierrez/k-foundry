# V0040Wckey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounting** | Pointer to [**[]V0040Accounting**](V0040Accounting.md) |  | [optional] 
**Cluster** | **string** |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**User** | **string** |  | 
**Flags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0040Wckey

`func NewV0040Wckey(cluster string, name string, user string, ) *V0040Wckey`

NewV0040Wckey instantiates a new V0040Wckey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040WckeyWithDefaults

`func NewV0040WckeyWithDefaults() *V0040Wckey`

NewV0040WckeyWithDefaults instantiates a new V0040Wckey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounting

`func (o *V0040Wckey) GetAccounting() []V0040Accounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *V0040Wckey) GetAccountingOk() (*[]V0040Accounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *V0040Wckey) SetAccounting(v []V0040Accounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *V0040Wckey) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetCluster

`func (o *V0040Wckey) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0040Wckey) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0040Wckey) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetId

`func (o *V0040Wckey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0040Wckey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0040Wckey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0040Wckey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0040Wckey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040Wckey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040Wckey) SetName(v string)`

SetName sets Name field to given value.


### GetUser

`func (o *V0040Wckey) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0040Wckey) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0040Wckey) SetUser(v string)`

SetUser sets User field to given value.


### GetFlags

`func (o *V0040Wckey) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0040Wckey) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0040Wckey) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0040Wckey) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


