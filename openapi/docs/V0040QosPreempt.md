# V0040QosPreempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to **[]string** |  | [optional] 
**Mode** | Pointer to **[]string** |  | [optional] 
**ExemptTime** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 

## Methods

### NewV0040QosPreempt

`func NewV0040QosPreempt() *V0040QosPreempt`

NewV0040QosPreempt instantiates a new V0040QosPreempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040QosPreemptWithDefaults

`func NewV0040QosPreemptWithDefaults() *V0040QosPreempt`

NewV0040QosPreemptWithDefaults instantiates a new V0040QosPreempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *V0040QosPreempt) GetList() []string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *V0040QosPreempt) GetListOk() (*[]string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *V0040QosPreempt) SetList(v []string)`

SetList sets List field to given value.

### HasList

`func (o *V0040QosPreempt) HasList() bool`

HasList returns a boolean if a field has been set.

### GetMode

`func (o *V0040QosPreempt) GetMode() []string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V0040QosPreempt) GetModeOk() (*[]string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V0040QosPreempt) SetMode(v []string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *V0040QosPreempt) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetExemptTime

`func (o *V0040QosPreempt) GetExemptTime() V0040Uint32NoVal`

GetExemptTime returns the ExemptTime field if non-nil, zero value otherwise.

### GetExemptTimeOk

`func (o *V0040QosPreempt) GetExemptTimeOk() (*V0040Uint32NoVal, bool)`

GetExemptTimeOk returns a tuple with the ExemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptTime

`func (o *V0040QosPreempt) SetExemptTime(v V0040Uint32NoVal)`

SetExemptTime sets ExemptTime field to given value.

### HasExemptTime

`func (o *V0040QosPreempt) HasExemptTime() bool`

HasExemptTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


