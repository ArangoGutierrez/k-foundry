# V0039Wckey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounting** | Pointer to [**[]V0039Accounting**](V0039Accounting.md) |  | [optional] 
**Cluster** | **string** |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**User** | **string** |  | 
**Flags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0039Wckey

`func NewV0039Wckey(cluster string, name string, user string, ) *V0039Wckey`

NewV0039Wckey instantiates a new V0039Wckey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039WckeyWithDefaults

`func NewV0039WckeyWithDefaults() *V0039Wckey`

NewV0039WckeyWithDefaults instantiates a new V0039Wckey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounting

`func (o *V0039Wckey) GetAccounting() []V0039Accounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *V0039Wckey) GetAccountingOk() (*[]V0039Accounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *V0039Wckey) SetAccounting(v []V0039Accounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *V0039Wckey) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetCluster

`func (o *V0039Wckey) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0039Wckey) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0039Wckey) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetId

`func (o *V0039Wckey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0039Wckey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0039Wckey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0039Wckey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0039Wckey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0039Wckey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0039Wckey) SetName(v string)`

SetName sets Name field to given value.


### GetUser

`func (o *V0039Wckey) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0039Wckey) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0039Wckey) SetUser(v string)`

SetUser sets User field to given value.


### GetFlags

`func (o *V0039Wckey) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0039Wckey) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0039Wckey) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0039Wckey) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


