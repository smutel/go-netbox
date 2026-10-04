# PatchedBulkWirelessLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**InterfaceA** | Pointer to [**BulkVirtualCircuitTerminationRequestInterface**](BulkVirtualCircuitTerminationRequestInterface.md) |  | [optional] 
**InterfaceB** | Pointer to [**BulkVirtualCircuitTerminationRequestInterface**](BulkVirtualCircuitTerminationRequestInterface.md) |  | [optional] 
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

### NewPatchedBulkWirelessLinkRequest

`func NewPatchedBulkWirelessLinkRequest(id int32, ) *PatchedBulkWirelessLinkRequest`

NewPatchedBulkWirelessLinkRequest instantiates a new PatchedBulkWirelessLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWirelessLinkRequestWithDefaults

`func NewPatchedBulkWirelessLinkRequestWithDefaults() *PatchedBulkWirelessLinkRequest`

NewPatchedBulkWirelessLinkRequestWithDefaults instantiates a new PatchedBulkWirelessLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWirelessLinkRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWirelessLinkRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWirelessLinkRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetInterfaceA

`func (o *PatchedBulkWirelessLinkRequest) GetInterfaceA() BulkVirtualCircuitTerminationRequestInterface`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *PatchedBulkWirelessLinkRequest) GetInterfaceAOk() (*BulkVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *PatchedBulkWirelessLinkRequest) SetInterfaceA(v BulkVirtualCircuitTerminationRequestInterface)`

SetInterfaceA sets InterfaceA field to given value.

### HasInterfaceA

`func (o *PatchedBulkWirelessLinkRequest) HasInterfaceA() bool`

HasInterfaceA returns a boolean if a field has been set.

### GetInterfaceB

`func (o *PatchedBulkWirelessLinkRequest) GetInterfaceB() BulkVirtualCircuitTerminationRequestInterface`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *PatchedBulkWirelessLinkRequest) GetInterfaceBOk() (*BulkVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *PatchedBulkWirelessLinkRequest) SetInterfaceB(v BulkVirtualCircuitTerminationRequestInterface)`

SetInterfaceB sets InterfaceB field to given value.

### HasInterfaceB

`func (o *PatchedBulkWirelessLinkRequest) HasInterfaceB() bool`

HasInterfaceB returns a boolean if a field has been set.

### GetSsid

`func (o *PatchedBulkWirelessLinkRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *PatchedBulkWirelessLinkRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *PatchedBulkWirelessLinkRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *PatchedBulkWirelessLinkRequest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWirelessLinkRequest) GetStatus() BulkCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWirelessLinkRequest) GetStatusOk() (*BulkCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWirelessLinkRequest) SetStatus(v BulkCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWirelessLinkRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkWirelessLinkRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWirelessLinkRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWirelessLinkRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWirelessLinkRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWirelessLinkRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWirelessLinkRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *PatchedBulkWirelessLinkRequest) GetAuthType() BulkWirelessLANRequestAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *PatchedBulkWirelessLinkRequest) GetAuthTypeOk() (*BulkWirelessLANRequestAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *PatchedBulkWirelessLinkRequest) SetAuthType(v BulkWirelessLANRequestAuthType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *PatchedBulkWirelessLinkRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthCipher

`func (o *PatchedBulkWirelessLinkRequest) GetAuthCipher() BulkWirelessLANRequestAuthCipher`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *PatchedBulkWirelessLinkRequest) GetAuthCipherOk() (*BulkWirelessLANRequestAuthCipher, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *PatchedBulkWirelessLinkRequest) SetAuthCipher(v BulkWirelessLANRequestAuthCipher)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *PatchedBulkWirelessLinkRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### GetAuthPsk

`func (o *PatchedBulkWirelessLinkRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *PatchedBulkWirelessLinkRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *PatchedBulkWirelessLinkRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *PatchedBulkWirelessLinkRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetDistance

`func (o *PatchedBulkWirelessLinkRequest) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *PatchedBulkWirelessLinkRequest) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *PatchedBulkWirelessLinkRequest) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *PatchedBulkWirelessLinkRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *PatchedBulkWirelessLinkRequest) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *PatchedBulkWirelessLinkRequest) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetDistanceUnit

`func (o *PatchedBulkWirelessLinkRequest) GetDistanceUnit() BulkCircuitRequestDistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *PatchedBulkWirelessLinkRequest) GetDistanceUnitOk() (*BulkCircuitRequestDistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *PatchedBulkWirelessLinkRequest) SetDistanceUnit(v BulkCircuitRequestDistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *PatchedBulkWirelessLinkRequest) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.

### SetDistanceUnitNil

`func (o *PatchedBulkWirelessLinkRequest) SetDistanceUnitNil(b bool)`

 SetDistanceUnitNil sets the value for DistanceUnit to be an explicit nil

### UnsetDistanceUnit
`func (o *PatchedBulkWirelessLinkRequest) UnsetDistanceUnit()`

UnsetDistanceUnit ensures that no value is present for DistanceUnit, not even an explicit nil
### GetDescription

`func (o *PatchedBulkWirelessLinkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWirelessLinkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWirelessLinkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWirelessLinkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkWirelessLinkRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkWirelessLinkRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkWirelessLinkRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkWirelessLinkRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkWirelessLinkRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkWirelessLinkRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkWirelessLinkRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkWirelessLinkRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkWirelessLinkRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkWirelessLinkRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWirelessLinkRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWirelessLinkRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWirelessLinkRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWirelessLinkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWirelessLinkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWirelessLinkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWirelessLinkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWirelessLinkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


