# V0040Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | Pointer to [**[]V0040AssocShort**](V0040AssocShort.md) |  | [optional] 
**Coordinators** | Pointer to [**[]V0040Coord**](V0040Coord.md) |  | [optional] 
**Description** | **string** |  | 
**Name** | **string** |  | 
**Organization** | **string** |  | 
**Flags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0040Account

`func NewV0040Account(description string, name string, organization string, ) *V0040Account`

NewV0040Account instantiates a new V0040Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040AccountWithDefaults

`func NewV0040AccountWithDefaults() *V0040Account`

NewV0040AccountWithDefaults instantiates a new V0040Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *V0040Account) GetAssociations() []V0040AssocShort`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0040Account) GetAssociationsOk() (*[]V0040AssocShort, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0040Account) SetAssociations(v []V0040AssocShort)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0040Account) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *V0040Account) GetCoordinators() []V0040Coord`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *V0040Account) GetCoordinatorsOk() (*[]V0040Coord, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *V0040Account) SetCoordinators(v []V0040Coord)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *V0040Account) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDescription

`func (o *V0040Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0040Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0040Account) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *V0040Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0040Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0040Account) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *V0040Account) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V0040Account) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V0040Account) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetFlags

`func (o *V0040Account) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0040Account) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0040Account) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0040Account) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


