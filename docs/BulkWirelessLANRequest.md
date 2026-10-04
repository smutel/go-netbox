# BulkWirelessLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Ssid** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**NullableBulkWirelessLANRequestGroup**](BulkWirelessLANRequestGroup.md) |  | [optional] 
**Status** | Pointer to [**BulkWirelessLANRequestStatus**](BulkWirelessLANRequestStatus.md) |  | [optional] 
**Vlan** | Pointer to [**NullableBulkInterfaceRequestUntaggedVlan**](BulkInterfaceRequestUntaggedVlan.md) |  | [optional] 
**ScopeType** | Pointer to **NullableString** |  | [optional] 
**ScopeId** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**AuthType** | Pointer to [**BulkWirelessLANRequestAuthType**](BulkWirelessLANRequestAuthType.md) |  | [optional] 
**AuthCipher** | Pointer to [**BulkWirelessLANRequestAuthCipher**](BulkWirelessLANRequestAuthCipher.md) |  | [optional] 
**AuthPsk** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkWirelessLANRequest

`func NewBulkWirelessLANRequest(id int32, ssid string, ) *BulkWirelessLANRequest`

NewBulkWirelessLANRequest instantiates a new BulkWirelessLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWirelessLANRequestWithDefaults

`func NewBulkWirelessLANRequestWithDefaults() *BulkWirelessLANRequest`

NewBulkWirelessLANRequestWithDefaults instantiates a new BulkWirelessLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWirelessLANRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWirelessLANRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWirelessLANRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSsid

`func (o *BulkWirelessLANRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *BulkWirelessLANRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *BulkWirelessLANRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetDescription

`func (o *BulkWirelessLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWirelessLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWirelessLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWirelessLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *BulkWirelessLANRequest) GetGroup() BulkWirelessLANRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkWirelessLANRequest) GetGroupOk() (*BulkWirelessLANRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkWirelessLANRequest) SetGroup(v BulkWirelessLANRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BulkWirelessLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BulkWirelessLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BulkWirelessLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetStatus

`func (o *BulkWirelessLANRequest) GetStatus() BulkWirelessLANRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWirelessLANRequest) GetStatusOk() (*BulkWirelessLANRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWirelessLANRequest) SetStatus(v BulkWirelessLANRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkWirelessLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlan

`func (o *BulkWirelessLANRequest) GetVlan() BulkInterfaceRequestUntaggedVlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *BulkWirelessLANRequest) GetVlanOk() (*BulkInterfaceRequestUntaggedVlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *BulkWirelessLANRequest) SetVlan(v BulkInterfaceRequestUntaggedVlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *BulkWirelessLANRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *BulkWirelessLANRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *BulkWirelessLANRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetScopeType

`func (o *BulkWirelessLANRequest) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *BulkWirelessLANRequest) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *BulkWirelessLANRequest) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *BulkWirelessLANRequest) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### SetScopeTypeNil

`func (o *BulkWirelessLANRequest) SetScopeTypeNil(b bool)`

 SetScopeTypeNil sets the value for ScopeType to be an explicit nil

### UnsetScopeType
`func (o *BulkWirelessLANRequest) UnsetScopeType()`

UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
### GetScopeId

`func (o *BulkWirelessLANRequest) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *BulkWirelessLANRequest) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *BulkWirelessLANRequest) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *BulkWirelessLANRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### SetScopeIdNil

`func (o *BulkWirelessLANRequest) SetScopeIdNil(b bool)`

 SetScopeIdNil sets the value for ScopeId to be an explicit nil

### UnsetScopeId
`func (o *BulkWirelessLANRequest) UnsetScopeId()`

UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
### GetTenant

`func (o *BulkWirelessLANRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWirelessLANRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWirelessLANRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWirelessLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWirelessLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWirelessLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *BulkWirelessLANRequest) GetAuthType() BulkWirelessLANRequestAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *BulkWirelessLANRequest) GetAuthTypeOk() (*BulkWirelessLANRequestAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *BulkWirelessLANRequest) SetAuthType(v BulkWirelessLANRequestAuthType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *BulkWirelessLANRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthCipher

`func (o *BulkWirelessLANRequest) GetAuthCipher() BulkWirelessLANRequestAuthCipher`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *BulkWirelessLANRequest) GetAuthCipherOk() (*BulkWirelessLANRequestAuthCipher, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *BulkWirelessLANRequest) SetAuthCipher(v BulkWirelessLANRequestAuthCipher)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *BulkWirelessLANRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### GetAuthPsk

`func (o *BulkWirelessLANRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *BulkWirelessLANRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *BulkWirelessLANRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *BulkWirelessLANRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetOwner

`func (o *BulkWirelessLANRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkWirelessLANRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkWirelessLANRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkWirelessLANRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkWirelessLANRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkWirelessLANRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkWirelessLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWirelessLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWirelessLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWirelessLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkWirelessLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWirelessLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWirelessLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWirelessLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWirelessLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWirelessLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWirelessLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWirelessLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


