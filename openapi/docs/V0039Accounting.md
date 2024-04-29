# V0039Accounting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocated** | Pointer to [**V0039AccountingAllocated**](V0039AccountingAllocated.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**TRES** | Pointer to [**V0039Tres**](V0039Tres.md) |  | [optional] 

## Methods

### NewV0039Accounting

`func NewV0039Accounting() *V0039Accounting`

NewV0039Accounting instantiates a new V0039Accounting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0039AccountingWithDefaults

`func NewV0039AccountingWithDefaults() *V0039Accounting`

NewV0039AccountingWithDefaults instantiates a new V0039Accounting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocated

`func (o *V0039Accounting) GetAllocated() V0039AccountingAllocated`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *V0039Accounting) GetAllocatedOk() (*V0039AccountingAllocated, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *V0039Accounting) SetAllocated(v V0039AccountingAllocated)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *V0039Accounting) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.

### GetId

`func (o *V0039Accounting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0039Accounting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0039Accounting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0039Accounting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStart

`func (o *V0039Accounting) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0039Accounting) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0039Accounting) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0039Accounting) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTRES

`func (o *V0039Accounting) GetTRES() V0039Tres`

GetTRES returns the TRES field if non-nil, zero value otherwise.

### GetTRESOk

`func (o *V0039Accounting) GetTRESOk() (*V0039Tres, bool)`

GetTRESOk returns a tuple with the TRES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTRES

`func (o *V0039Accounting) SetTRES(v V0039Tres)`

SetTRES sets TRES field to given value.

### HasTRES

`func (o *V0039Accounting) HasTRES() bool`

HasTRES returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


