# V0038NodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0038Meta**](V0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0038Error**](V0038Error.md) | slurm errors | [optional] 
**Nodes** | Pointer to [**[]V0038Node**](V0038Node.md) | nodes info | [optional] 

## Methods

### NewV0038NodesResponse

`func NewV0038NodesResponse() *V0038NodesResponse`

NewV0038NodesResponse instantiates a new V0038NodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038NodesResponseWithDefaults

`func NewV0038NodesResponseWithDefaults() *V0038NodesResponse`

NewV0038NodesResponseWithDefaults instantiates a new V0038NodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0038NodesResponse) GetMeta() V0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0038NodesResponse) GetMetaOk() (*V0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0038NodesResponse) SetMeta(v V0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0038NodesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0038NodesResponse) GetErrors() []V0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0038NodesResponse) GetErrorsOk() (*[]V0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0038NodesResponse) SetErrors(v []V0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0038NodesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetNodes

`func (o *V0038NodesResponse) GetNodes() []V0038Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0038NodesResponse) GetNodesOk() (*[]V0038Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0038NodesResponse) SetNodes(v []V0038Node)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0038NodesResponse) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


