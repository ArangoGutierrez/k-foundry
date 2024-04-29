# V0040OpenapiMetaPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Slurm plugin type (if applicable) | [optional] 
**Name** | Pointer to **string** | Slurm plugin name (if applicable) | [optional] 
**DataParser** | Pointer to **string** | Slurm data_parser plugin | [optional] 
**AccountingStorage** | Pointer to **string** | Slurm accounting plugin | [optional] 

## Methods

### NewV0040OpenapiMetaPlugin

`func NewV0040OpenapiMetaPlugin() *V0040OpenapiMetaPlugin`

NewV0040OpenapiMetaPlugin instantiates a new V0040OpenapiMetaPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040OpenapiMetaPluginWithDefaults

`func NewV0040OpenapiMetaPluginWithDefaults() *V0040OpenapiMetaPlugin`

NewV0040OpenapiMetaPluginWithDefaults instantiates a new V0040OpenapiMetaPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V0040OpenapiMetaPlugin) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0040OpenapiMetaPlugin) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0040OpenapiMetaPlugin) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V0040OpenapiMetaPlugin) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *V0040OpenapiMetaPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040OpenapiMetaPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040OpenapiMetaPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0040OpenapiMetaPlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataParser

`func (o *V0040OpenapiMetaPlugin) GetDataParser() string`

GetDataParser returns the DataParser field if non-nil, zero value otherwise.

### GetDataParserOk

`func (o *V0040OpenapiMetaPlugin) GetDataParserOk() (*string, bool)`

GetDataParserOk returns a tuple with the DataParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataParser

`func (o *V0040OpenapiMetaPlugin) SetDataParser(v string)`

SetDataParser sets DataParser field to given value.

### HasDataParser

`func (o *V0040OpenapiMetaPlugin) HasDataParser() bool`

HasDataParser returns a boolean if a field has been set.

### GetAccountingStorage

`func (o *V0040OpenapiMetaPlugin) GetAccountingStorage() string`

GetAccountingStorage returns the AccountingStorage field if non-nil, zero value otherwise.

### GetAccountingStorageOk

`func (o *V0040OpenapiMetaPlugin) GetAccountingStorageOk() (*string, bool)`

GetAccountingStorageOk returns a tuple with the AccountingStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingStorage

`func (o *V0040OpenapiMetaPlugin) SetAccountingStorage(v string)`

SetAccountingStorage sets AccountingStorage field to given value.

### HasAccountingStorage

`func (o *V0040OpenapiMetaPlugin) HasAccountingStorage() bool`

HasAccountingStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


