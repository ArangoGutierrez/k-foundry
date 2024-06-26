# Dbv0038Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorNumber** | Pointer to **int32** | Slurm internal error number | [optional] 
**Error** | Pointer to **string** | Error message | [optional] 
**Source** | Pointer to **string** | Where error occurred in the source | [optional] 
**Description** | Pointer to **string** | Explanation of cause of error | [optional] 

## Methods

### NewDbv0038Error

`func NewDbv0038Error() *Dbv0038Error`

NewDbv0038Error instantiates a new Dbv0038Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038ErrorWithDefaults

`func NewDbv0038ErrorWithDefaults() *Dbv0038Error`

NewDbv0038ErrorWithDefaults instantiates a new Dbv0038Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorNumber

`func (o *Dbv0038Error) GetErrorNumber() int32`

GetErrorNumber returns the ErrorNumber field if non-nil, zero value otherwise.

### GetErrorNumberOk

`func (o *Dbv0038Error) GetErrorNumberOk() (*int32, bool)`

GetErrorNumberOk returns a tuple with the ErrorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNumber

`func (o *Dbv0038Error) SetErrorNumber(v int32)`

SetErrorNumber sets ErrorNumber field to given value.

### HasErrorNumber

`func (o *Dbv0038Error) HasErrorNumber() bool`

HasErrorNumber returns a boolean if a field has been set.

### GetError

`func (o *Dbv0038Error) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Dbv0038Error) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Dbv0038Error) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Dbv0038Error) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSource

`func (o *Dbv0038Error) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Dbv0038Error) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Dbv0038Error) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Dbv0038Error) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDescription

`func (o *Dbv0038Error) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dbv0038Error) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dbv0038Error) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dbv0038Error) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


