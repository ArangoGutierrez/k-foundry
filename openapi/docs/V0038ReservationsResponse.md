# V0038ReservationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**V0038Meta**](V0038Meta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0038Error**](V0038Error.md) | slurm errors | [optional] 
**Reservations** | Pointer to [**[]V0038Reservation**](V0038Reservation.md) | reservation info | [optional] 

## Methods

### NewV0038ReservationsResponse

`func NewV0038ReservationsResponse() *V0038ReservationsResponse`

NewV0038ReservationsResponse instantiates a new V0038ReservationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0038ReservationsResponseWithDefaults

`func NewV0038ReservationsResponseWithDefaults() *V0038ReservationsResponse`

NewV0038ReservationsResponseWithDefaults instantiates a new V0038ReservationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V0038ReservationsResponse) GetMeta() V0038Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0038ReservationsResponse) GetMetaOk() (*V0038Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0038ReservationsResponse) SetMeta(v V0038Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0038ReservationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0038ReservationsResponse) GetErrors() []V0038Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0038ReservationsResponse) GetErrorsOk() (*[]V0038Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0038ReservationsResponse) SetErrors(v []V0038Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0038ReservationsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetReservations

`func (o *V0038ReservationsResponse) GetReservations() []V0038Reservation`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *V0038ReservationsResponse) GetReservationsOk() (*[]V0038Reservation, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *V0038ReservationsResponse) SetReservations(v []V0038Reservation)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *V0038ReservationsResponse) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


