# Dbv0038Accounting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocated** | Pointer to **int32** | total seconds allocated | [optional] 
**Id** | Pointer to **int32** | association/wckey ID | [optional] 
**Start** | Pointer to **int32** | UNIX timestamp when accounting period started | [optional] 
**TRES** | Pointer to [**[]Dbv0038TresListInner**](Dbv0038TresListInner.md) | TRES list of attributes | [optional] 

## Methods

### NewDbv0038Accounting

`func NewDbv0038Accounting() *Dbv0038Accounting`

NewDbv0038Accounting instantiates a new Dbv0038Accounting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038AccountingWithDefaults

`func NewDbv0038AccountingWithDefaults() *Dbv0038Accounting`

NewDbv0038AccountingWithDefaults instantiates a new Dbv0038Accounting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocated

`func (o *Dbv0038Accounting) GetAllocated() int32`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *Dbv0038Accounting) GetAllocatedOk() (*int32, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *Dbv0038Accounting) SetAllocated(v int32)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *Dbv0038Accounting) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.

### GetId

`func (o *Dbv0038Accounting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dbv0038Accounting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dbv0038Accounting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Dbv0038Accounting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStart

`func (o *Dbv0038Accounting) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Dbv0038Accounting) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Dbv0038Accounting) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *Dbv0038Accounting) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTRES

`func (o *Dbv0038Accounting) GetTRES() []Dbv0038TresListInner`

GetTRES returns the TRES field if non-nil, zero value otherwise.

### GetTRESOk

`func (o *Dbv0038Accounting) GetTRESOk() (*[]Dbv0038TresListInner, bool)`

GetTRESOk returns a tuple with the TRES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTRES

`func (o *Dbv0038Accounting) SetTRES(v []Dbv0038TresListInner)`

SetTRES sets TRES field to given value.

### HasTRES

`func (o *Dbv0038Accounting) HasTRES() bool`

HasTRES returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


