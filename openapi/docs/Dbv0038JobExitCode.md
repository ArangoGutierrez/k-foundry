# Dbv0038JobExitCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Job exit status | [optional] 
**ReturnCode** | Pointer to **int32** | Return code from parent process | [optional] 
**Signal** | Pointer to [**Dbv0038JobExitCodeSignal**](Dbv0038JobExitCodeSignal.md) |  | [optional] 

## Methods

### NewDbv0038JobExitCode

`func NewDbv0038JobExitCode() *Dbv0038JobExitCode`

NewDbv0038JobExitCode instantiates a new Dbv0038JobExitCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038JobExitCodeWithDefaults

`func NewDbv0038JobExitCodeWithDefaults() *Dbv0038JobExitCode`

NewDbv0038JobExitCodeWithDefaults instantiates a new Dbv0038JobExitCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Dbv0038JobExitCode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Dbv0038JobExitCode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Dbv0038JobExitCode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Dbv0038JobExitCode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReturnCode

`func (o *Dbv0038JobExitCode) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *Dbv0038JobExitCode) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *Dbv0038JobExitCode) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *Dbv0038JobExitCode) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetSignal

`func (o *Dbv0038JobExitCode) GetSignal() Dbv0038JobExitCodeSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *Dbv0038JobExitCode) GetSignalOk() (*Dbv0038JobExitCodeSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *Dbv0038JobExitCode) SetSignal(v Dbv0038JobExitCodeSignal)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *Dbv0038JobExitCode) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


