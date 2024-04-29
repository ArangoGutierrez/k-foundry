# V0038License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseName** | Pointer to **string** | name of license | [optional] 
**Total** | Pointer to **int32** | total number of licenses | [optional] 
**Used** | Pointer to **int32** | number of licenses in use | [optional] 
**Free** | Pointer to **int32** | number of licenses available | [optional] 
**Reserved** | Pointer to **int32** | number of licenses reserved | [optional] 
**Remote** | Pointer to **bool** | license is remote | [optional] 

## Methods

### NewV0038License

`func NewV0038License() *V0038License`

NewV0038License instantiates a new V0038License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038LicenseWithDefaults

`func NewV0038LicenseWithDefaults() *V0038License`

NewV0038LicenseWithDefaults instantiates a new V0038License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseName

`func (o *V0038License) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *V0038License) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *V0038License) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *V0038License) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### GetTotal

`func (o *V0038License) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0038License) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0038License) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0038License) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *V0038License) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *V0038License) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *V0038License) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *V0038License) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetFree

`func (o *V0038License) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *V0038License) GetFreeOk() (*int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *V0038License) SetFree(v int32)`

SetFree sets Free field to given value.

### HasFree

`func (o *V0038License) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetReserved

`func (o *V0038License) GetReserved() int32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *V0038License) GetReservedOk() (*int32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *V0038License) SetReserved(v int32)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *V0038License) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetRemote

`func (o *V0038License) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *V0038License) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *V0038License) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *V0038License) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


