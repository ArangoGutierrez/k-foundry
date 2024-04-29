# Dbv0038UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Dbv0038Meta**](Dbv0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]Dbv0038Error**](Dbv0038Error.md) | Slurm errors | [optional] 
**Users** | Pointer to [**[]Dbv0038User**](Dbv0038User.md) | Array of users | [optional] 

## Methods

### NewDbv0038UserInfo

`func NewDbv0038UserInfo() *Dbv0038UserInfo`

NewDbv0038UserInfo instantiates a new Dbv0038UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038UserInfoWithDefaults

`func NewDbv0038UserInfoWithDefaults() *Dbv0038UserInfo`

NewDbv0038UserInfoWithDefaults instantiates a new Dbv0038UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Dbv0038UserInfo) GetMeta() Dbv0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Dbv0038UserInfo) GetMetaOk() (*Dbv0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Dbv0038UserInfo) SetMeta(v Dbv0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Dbv0038UserInfo) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *Dbv0038UserInfo) GetErrors() []Dbv0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0038UserInfo) GetErrorsOk() (*[]Dbv0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0038UserInfo) SetErrors(v []Dbv0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0038UserInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetUsers

`func (o *Dbv0038UserInfo) GetUsers() []Dbv0038User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Dbv0038UserInfo) GetUsersOk() (*[]Dbv0038User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Dbv0038UserInfo) SetUsers(v []Dbv0038User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Dbv0038UserInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


