# Dbv0038User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministratorLevel** | Pointer to **string** | Description of administrator level | [optional] 
**Associations** | Pointer to [**[]Dbv0038AssociationShortInfo**](Dbv0038AssociationShortInfo.md) | Assigned associations | [optional] 
**Coordinators** | Pointer to [**[]Dbv0038CoordinatorInfo**](Dbv0038CoordinatorInfo.md) | List of assigned coordinators | [optional] 
**Default** | Pointer to [**Dbv0038UserDefault**](Dbv0038UserDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** | List of properties of user | [optional] 
**Name** | Pointer to **string** | User name | [optional] 

## Methods

### NewDbv0038User

`func NewDbv0038User() *Dbv0038User`

NewDbv0038User instantiates a new Dbv0038User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0038UserWithDefaults

`func NewDbv0038UserWithDefaults() *Dbv0038User`

NewDbv0038UserWithDefaults instantiates a new Dbv0038User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministratorLevel

`func (o *Dbv0038User) GetAdministratorLevel() string`

GetAdministratorLevel returns the AdministratorLevel field if non-nil, zero value otherwise.

### GetAdministratorLevelOk

`func (o *Dbv0038User) GetAdministratorLevelOk() (*string, bool)`

GetAdministratorLevelOk returns a tuple with the AdministratorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorLevel

`func (o *Dbv0038User) SetAdministratorLevel(v string)`

SetAdministratorLevel sets AdministratorLevel field to given value.

### HasAdministratorLevel

`func (o *Dbv0038User) HasAdministratorLevel() bool`

HasAdministratorLevel returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0038User) GetAssociations() []Dbv0038AssociationShortInfo`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0038User) GetAssociationsOk() (*[]Dbv0038AssociationShortInfo, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0038User) SetAssociations(v []Dbv0038AssociationShortInfo)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0038User) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *Dbv0038User) GetCoordinators() []Dbv0038CoordinatorInfo`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *Dbv0038User) GetCoordinatorsOk() (*[]Dbv0038CoordinatorInfo, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *Dbv0038User) SetCoordinators(v []Dbv0038CoordinatorInfo)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *Dbv0038User) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDefault

`func (o *Dbv0038User) GetDefault() Dbv0038UserDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Dbv0038User) GetDefaultOk() (*Dbv0038UserDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Dbv0038User) SetDefault(v Dbv0038UserDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Dbv0038User) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *Dbv0038User) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbv0038User) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Dbv0038User) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbv0038User) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *Dbv0038User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0038User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dbv0038User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0038User) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


