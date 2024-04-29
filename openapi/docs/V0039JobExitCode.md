# V0039JobExitCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | exit status | [optional] 
**ReturnCode** | Pointer to **int32** | return code (numeric) | [optional] 
**Signal** | Pointer to [**V0039JobExitCodeSignal**](V0039JobExitCodeSignal.md) |  | [optional] 

## Methods

### NewV0039JobExitCode

`func NewV0039JobExitCode() *V0039JobExitCode`

NewV0039JobExitCode instantiates a new V0039JobExitCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039JobExitCodeWithDefaults

`func NewV0039JobExitCodeWithDefaults() *V0039JobExitCode`

NewV0039JobExitCodeWithDefaults instantiates a new V0039JobExitCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V0039JobExitCode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0039JobExitCode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0039JobExitCode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0039JobExitCode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReturnCode

`func (o *V0039JobExitCode) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *V0039JobExitCode) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *V0039JobExitCode) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *V0039JobExitCode) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetSignal

`func (o *V0039JobExitCode) GetSignal() V0039JobExitCodeSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *V0039JobExitCode) GetSignalOk() (*V0039JobExitCodeSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *V0039JobExitCode) SetSignal(v V0039JobExitCodeSignal)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *V0039JobExitCode) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


