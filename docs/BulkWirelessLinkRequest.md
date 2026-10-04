# BulkWirelessLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**InterfaceA** | [**BulkVirtualCircuitTerminationRequestInterface**](BulkVirtualCircuitTerminationRequestInterface.md) |  | 
**InterfaceB** | [**BulkVirtualCircuitTerminationRequestInterface**](BulkVirtualCircuitTerminationRequestInterface.md) |  | 
**Ssid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkCableRequestStatus**](BulkCableRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**AuthType** | Pointer to [**BulkWirelessLANRequestAuthType**](BulkWirelessLANRequestAuthType.md) |  | [optional] 
**AuthCipher** | Pointer to [**BulkWirelessLANRequestAuthCipher**](BulkWirelessLANRequestAuthCipher.md) |  | [optional] 
**AuthPsk** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **NullableFloat64** |  | [optional] 
**DistanceUnit** | Pointer to [**NullableBulkCircuitRequestDistanceUnit**](BulkCircuitRequestDistanceUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkWirelessLinkRequest

`func NewBulkWirelessLinkRequest(id int32, interfaceA BulkVirtualCircuitTerminationRequestInterface, interfaceB BulkVirtualCircuitTerminationRequestInterface, ) *BulkWirelessLinkRequest`

NewBulkWirelessLinkRequest instantiates a new BulkWirelessLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWirelessLinkRequestWithDefaults

`func NewBulkWirelessLinkRequestWithDefaults() *BulkWirelessLinkRequest`

NewBulkWirelessLinkRequestWithDefaults instantiates a new BulkWirelessLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWirelessLinkRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWirelessLinkRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWirelessLinkRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetInterfaceA

`func (o *BulkWirelessLinkRequest) GetInterfaceA() BulkVirtualCircuitTerminationRequestInterface`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *BulkWirelessLinkRequest) GetInterfaceAOk() (*BulkVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *BulkWirelessLinkRequest) SetInterfaceA(v BulkVirtualCircuitTerminationRequestInterface)`

SetInterfaceA sets InterfaceA field to given value.


### GetInterfaceB

`func (o *BulkWirelessLinkRequest) GetInterfaceB() BulkVirtualCircuitTerminationRequestInterface`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *BulkWirelessLinkRequest) GetInterfaceBOk() (*BulkVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *BulkWirelessLinkRequest) SetInterfaceB(v BulkVirtualCircuitTerminationRequestInterface)`

SetInterfaceB sets InterfaceB field to given value.


### GetSsid

`func (o *BulkWirelessLinkRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *BulkWirelessLinkRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *BulkWirelessLinkRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *BulkWirelessLinkRequest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStatus

`func (o *BulkWirelessLinkRequest) GetStatus() BulkCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWirelessLinkRequest) GetStatusOk() (*BulkCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWirelessLinkRequest) SetStatus(v BulkCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkWirelessLinkRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *BulkWirelessLinkRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWirelessLinkRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWirelessLinkRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWirelessLinkRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWirelessLinkRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWirelessLinkRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *BulkWirelessLinkRequest) GetAuthType() BulkWirelessLANRequestAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *BulkWirelessLinkRequest) GetAuthTypeOk() (*BulkWirelessLANRequestAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *BulkWirelessLinkRequest) SetAuthType(v BulkWirelessLANRequestAuthType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *BulkWirelessLinkRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthCipher

`func (o *BulkWirelessLinkRequest) GetAuthCipher() BulkWirelessLANRequestAuthCipher`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *BulkWirelessLinkRequest) GetAuthCipherOk() (*BulkWirelessLANRequestAuthCipher, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *BulkWirelessLinkRequest) SetAuthCipher(v BulkWirelessLANRequestAuthCipher)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *BulkWirelessLinkRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### GetAuthPsk

`func (o *BulkWirelessLinkRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *BulkWirelessLinkRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *BulkWirelessLinkRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *BulkWirelessLinkRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetDistance

`func (o *BulkWirelessLinkRequest) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *BulkWirelessLinkRequest) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *BulkWirelessLinkRequest) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *BulkWirelessLinkRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *BulkWirelessLinkRequest) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *BulkWirelessLinkRequest) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetDistanceUnit

`func (o *BulkWirelessLinkRequest) GetDistanceUnit() BulkCircuitRequestDistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *BulkWirelessLinkRequest) GetDistanceUnitOk() (*BulkCircuitRequestDistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *BulkWirelessLinkRequest) SetDistanceUnit(v BulkCircuitRequestDistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *BulkWirelessLinkRequest) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.

### SetDistanceUnitNil

`func (o *BulkWirelessLinkRequest) SetDistanceUnitNil(b bool)`

 SetDistanceUnitNil sets the value for DistanceUnit to be an explicit nil

### UnsetDistanceUnit
`func (o *BulkWirelessLinkRequest) UnsetDistanceUnit()`

UnsetDistanceUnit ensures that no value is present for DistanceUnit, not even an explicit nil
### GetDescription

`func (o *BulkWirelessLinkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWirelessLinkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWirelessLinkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWirelessLinkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkWirelessLinkRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkWirelessLinkRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkWirelessLinkRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkWirelessLinkRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkWirelessLinkRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkWirelessLinkRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkWirelessLinkRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWirelessLinkRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWirelessLinkRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWirelessLinkRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkWirelessLinkRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWirelessLinkRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWirelessLinkRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWirelessLinkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWirelessLinkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWirelessLinkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWirelessLinkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWirelessLinkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


