# V0040UserShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adminlevel** | Pointer to **[]string** | Admin level of user.  Valid levels are None, Operator, and Admin. | [optional] 
**Defaultaccount** | Pointer to **string** | Identify the default bank account name to be used for a job if none is specified at submission time. | [optional] 
**Defaultwckey** | Pointer to **string** | Identify the default Workload Characterization Key. | [optional] 

## Methods

### NewV0040UserShort

`func NewV0040UserShort() *V0040UserShort`

NewV0040UserShort instantiates a new V0040UserShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040UserShortWithDefaults

`func NewV0040UserShortWithDefaults() *V0040UserShort`

NewV0040UserShortWithDefaults instantiates a new V0040UserShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminlevel

`func (o *V0040UserShort) GetAdminlevel() []string`

GetAdminlevel returns the Adminlevel field if non-nil, zero value otherwise.

### GetAdminlevelOk

`func (o *V0040UserShort) GetAdminlevelOk() (*[]string, bool)`

GetAdminlevelOk returns a tuple with the Adminlevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminlevel

`func (o *V0040UserShort) SetAdminlevel(v []string)`

SetAdminlevel sets Adminlevel field to given value.

### HasAdminlevel

`func (o *V0040UserShort) HasAdminlevel() bool`

HasAdminlevel returns a boolean if a field has been set.

### GetDefaultaccount

`func (o *V0040UserShort) GetDefaultaccount() string`

GetDefaultaccount returns the Defaultaccount field if non-nil, zero value otherwise.

### GetDefaultaccountOk

`func (o *V0040UserShort) GetDefaultaccountOk() (*string, bool)`

GetDefaultaccountOk returns a tuple with the Defaultaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultaccount

`func (o *V0040UserShort) SetDefaultaccount(v string)`

SetDefaultaccount sets Defaultaccount field to given value.

### HasDefaultaccount

`func (o *V0040UserShort) HasDefaultaccount() bool`

HasDefaultaccount returns a boolean if a field has been set.

### GetDefaultwckey

`func (o *V0040UserShort) GetDefaultwckey() string`

GetDefaultwckey returns the Defaultwckey field if non-nil, zero value otherwise.

### GetDefaultwckeyOk

`func (o *V0040UserShort) GetDefaultwckeyOk() (*string, bool)`

GetDefaultwckeyOk returns a tuple with the Defaultwckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultwckey

`func (o *V0040UserShort) SetDefaultwckey(v string)`

SetDefaultwckey sets Defaultwckey field to given value.

### HasDefaultwckey

`func (o *V0040UserShort) HasDefaultwckey() bool`

HasDefaultwckey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


