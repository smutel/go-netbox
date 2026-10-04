# PatchedBulkWirelessLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Ssid** | Pointer to **string** |  | [optional] 
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

### NewPatchedBulkWirelessLANRequest

`func NewPatchedBulkWirelessLANRequest(id int32, ) *PatchedBulkWirelessLANRequest`

NewPatchedBulkWirelessLANRequest instantiates a new PatchedBulkWirelessLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWirelessLANRequestWithDefaults

`func NewPatchedBulkWirelessLANRequestWithDefaults() *PatchedBulkWirelessLANRequest`

NewPatchedBulkWirelessLANRequestWithDefaults instantiates a new PatchedBulkWirelessLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWirelessLANRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWirelessLANRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWirelessLANRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSsid

`func (o *PatchedBulkWirelessLANRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *PatchedBulkWirelessLANRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *PatchedBulkWirelessLANRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *PatchedBulkWirelessLANRequest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWirelessLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWirelessLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWirelessLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWirelessLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedBulkWirelessLANRequest) GetGroup() BulkWirelessLANRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkWirelessLANRequest) GetGroupOk() (*BulkWirelessLANRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkWirelessLANRequest) SetGroup(v BulkWirelessLANRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkWirelessLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedBulkWirelessLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedBulkWirelessLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetStatus

`func (o *PatchedBulkWirelessLANRequest) GetStatus() BulkWirelessLANRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWirelessLANRequest) GetStatusOk() (*BulkWirelessLANRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWirelessLANRequest) SetStatus(v BulkWirelessLANRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWirelessLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlan

`func (o *PatchedBulkWirelessLANRequest) GetVlan() BulkInterfaceRequestUntaggedVlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PatchedBulkWirelessLANRequest) GetVlanOk() (*BulkInterfaceRequestUntaggedVlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PatchedBulkWirelessLANRequest) SetVlan(v BulkInterfaceRequestUntaggedVlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PatchedBulkWirelessLANRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *PatchedBulkWirelessLANRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *PatchedBulkWirelessLANRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetScopeType

`func (o *PatchedBulkWirelessLANRequest) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *PatchedBulkWirelessLANRequest) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *PatchedBulkWirelessLANRequest) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *PatchedBulkWirelessLANRequest) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### SetScopeTypeNil

`func (o *PatchedBulkWirelessLANRequest) SetScopeTypeNil(b bool)`

 SetScopeTypeNil sets the value for ScopeType to be an explicit nil

### UnsetScopeType
`func (o *PatchedBulkWirelessLANRequest) UnsetScopeType()`

UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
### GetScopeId

`func (o *PatchedBulkWirelessLANRequest) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *PatchedBulkWirelessLANRequest) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *PatchedBulkWirelessLANRequest) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *PatchedBulkWirelessLANRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### SetScopeIdNil

`func (o *PatchedBulkWirelessLANRequest) SetScopeIdNil(b bool)`

 SetScopeIdNil sets the value for ScopeId to be an explicit nil

### UnsetScopeId
`func (o *PatchedBulkWirelessLANRequest) UnsetScopeId()`

UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWirelessLANRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWirelessLANRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWirelessLANRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWirelessLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWirelessLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWirelessLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *PatchedBulkWirelessLANRequest) GetAuthType() BulkWirelessLANRequestAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *PatchedBulkWirelessLANRequest) GetAuthTypeOk() (*BulkWirelessLANRequestAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *PatchedBulkWirelessLANRequest) SetAuthType(v BulkWirelessLANRequestAuthType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *PatchedBulkWirelessLANRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthCipher

`func (o *PatchedBulkWirelessLANRequest) GetAuthCipher() BulkWirelessLANRequestAuthCipher`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *PatchedBulkWirelessLANRequest) GetAuthCipherOk() (*BulkWirelessLANRequestAuthCipher, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *PatchedBulkWirelessLANRequest) SetAuthCipher(v BulkWirelessLANRequestAuthCipher)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *PatchedBulkWirelessLANRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### GetAuthPsk

`func (o *PatchedBulkWirelessLANRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *PatchedBulkWirelessLANRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *PatchedBulkWirelessLANRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *PatchedBulkWirelessLANRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkWirelessLANRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkWirelessLANRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkWirelessLANRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkWirelessLANRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkWirelessLANRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkWirelessLANRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkWirelessLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkWirelessLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkWirelessLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkWirelessLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWirelessLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWirelessLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWirelessLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWirelessLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWirelessLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWirelessLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWirelessLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWirelessLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


