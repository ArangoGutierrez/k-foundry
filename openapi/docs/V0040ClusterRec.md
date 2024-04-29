# V0040ClusterRec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**V0039ClusterRecController**](V0039ClusterRecController.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to **string** |  | [optional] 
**SelectPlugin** | Pointer to **string** |  | [optional] 
**Associations** | Pointer to [**V0040ClusterRecAssociations**](V0040ClusterRecAssociations.md) |  | [optional] 
**RpcVersion** | Pointer to **int32** |  | [optional] 
**Tres** | Pointer to [**[]V0040Tres**](V0040Tres.md) |  | [optional] 

## Methods

### NewV0040ClusterRec

`func NewV0040ClusterRec() *V0040ClusterRec`

NewV0040ClusterRec instantiates a new V0040ClusterRec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040ClusterRecWithDefaults

`func NewV0040ClusterRecWithDefaults() *V0040ClusterRec`

NewV0040ClusterRecWithDefaults instantiates a new V0040ClusterRec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *V0040ClusterRec) GetController() V0039ClusterRecController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *V0040ClusterRec) GetControllerOk() (*V0039ClusterRecController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *V0040ClusterRec) SetController(v V0039ClusterRecController)`

SetController sets Controller field to given value.

### HasController

`func (o *V0040ClusterRec) HasController() bool`

HasController returns a boolean if a field has been set.

### GetFlags

`func (o *V0040ClusterRec) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0040ClusterRec) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0040ClusterRec) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0040ClusterRec) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *V0040ClusterRec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040ClusterRec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040ClusterRec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0040ClusterRec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *V0040ClusterRec) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0040ClusterRec) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0040ClusterRec) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0040ClusterRec) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSelectPlugin

`func (o *V0040ClusterRec) GetSelectPlugin() string`

GetSelectPlugin returns the SelectPlugin field if non-nil, zero value otherwise.

### GetSelectPluginOk

`func (o *V0040ClusterRec) GetSelectPluginOk() (*string, bool)`

GetSelectPluginOk returns a tuple with the SelectPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectPlugin

`func (o *V0040ClusterRec) SetSelectPlugin(v string)`

SetSelectPlugin sets SelectPlugin field to given value.

### HasSelectPlugin

`func (o *V0040ClusterRec) HasSelectPlugin() bool`

HasSelectPlugin returns a boolean if a field has been set.

### GetAssociations

`func (o *V0040ClusterRec) GetAssociations() V0040ClusterRecAssociations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0040ClusterRec) GetAssociationsOk() (*V0040ClusterRecAssociations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0040ClusterRec) SetAssociations(v V0040ClusterRecAssociations)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0040ClusterRec) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetRpcVersion

`func (o *V0040ClusterRec) GetRpcVersion() int32`

GetRpcVersion returns the RpcVersion field if non-nil, zero value otherwise.

### GetRpcVersionOk

`func (o *V0040ClusterRec) GetRpcVersionOk() (*int32, bool)`

GetRpcVersionOk returns a tuple with the RpcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcVersion

`func (o *V0040ClusterRec) SetRpcVersion(v int32)`

SetRpcVersion sets RpcVersion field to given value.

### HasRpcVersion

`func (o *V0040ClusterRec) HasRpcVersion() bool`

HasRpcVersion returns a boolean if a field has been set.

### GetTres

`func (o *V0040ClusterRec) GetTres() []V0040Tres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0040ClusterRec) GetTresOk() (*[]V0040Tres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0040ClusterRec) SetTres(v []V0040Tres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0040ClusterRec) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


