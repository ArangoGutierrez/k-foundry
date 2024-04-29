# V0038Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | error message | [optional] 
**ErrorNumber** | Pointer to **int32** | Slurm internal error number | [optional] 

## Methods

### NewV0038Error

`func NewV0038Error() *V0038Error`

NewV0038Error instantiates a new V0038Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038ErrorWithDefaults

`func NewV0038ErrorWithDefaults() *V0038Error`

NewV0038ErrorWithDefaults instantiates a new V0038Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V0038Error) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0038Error) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0038Error) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V0038Error) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorNumber

`func (o *V0038Error) GetErrorNumber() int32`

GetErrorNumber returns the ErrorNumber field if non-nil, zero value otherwise.

### GetErrorNumberOk

`func (o *V0038Error) GetErrorNumberOk() (*int32, bool)`

GetErrorNumberOk returns a tuple with the ErrorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNumber

`func (o *V0038Error) SetErrorNumber(v int32)`

SetErrorNumber sets ErrorNumber field to given value.

### HasErrorNumber

`func (o *V0038Error) HasErrorNumber() bool`

HasErrorNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


