# PatchedBulkTunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkTunnelRequestStatus**](BulkTunnelRequestStatus.md) |  | [optional] 
**Group** | Pointer to [**NullableBulkTunnelRequestGroup**](BulkTunnelRequestGroup.md) |  | [optional] 
**Encapsulation** | Pointer to [**BulkTunnelRequestEncapsulation**](BulkTunnelRequestEncapsulation.md) |  | [optional] 
**IpsecProfile** | Pointer to [**NullableBulkTunnelRequestIpsecProfile**](BulkTunnelRequestIpsecProfile.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**TunnelId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkTunnelRequest

`func NewPatchedBulkTunnelRequest(id int32, ) *PatchedBulkTunnelRequest`

NewPatchedBulkTunnelRequest instantiates a new PatchedBulkTunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkTunnelRequestWithDefaults

`func NewPatchedBulkTunnelRequestWithDefaults() *PatchedBulkTunnelRequest`

NewPatchedBulkTunnelRequestWithDefaults instantiates a new PatchedBulkTunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkTunnelRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkTunnelRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkTunnelRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkTunnelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkTunnelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkTunnelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkTunnelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkTunnelRequest) GetStatus() BulkTunnelRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkTunnelRequest) GetStatusOk() (*BulkTunnelRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkTunnelRequest) SetStatus(v BulkTunnelRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkTunnelRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedBulkTunnelRequest) GetGroup() BulkTunnelRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkTunnelRequest) GetGroupOk() (*BulkTunnelRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkTunnelRequest) SetGroup(v BulkTunnelRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkTunnelRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedBulkTunnelRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedBulkTunnelRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetEncapsulation

`func (o *PatchedBulkTunnelRequest) GetEncapsulation() BulkTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *PatchedBulkTunnelRequest) GetEncapsulationOk() (*BulkTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *PatchedBulkTunnelRequest) SetEncapsulation(v BulkTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *PatchedBulkTunnelRequest) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetIpsecProfile

`func (o *PatchedBulkTunnelRequest) GetIpsecProfile() BulkTunnelRequestIpsecProfile`

GetIpsecProfile returns the IpsecProfile field if non-nil, zero value otherwise.

### GetIpsecProfileOk

`func (o *PatchedBulkTunnelRequest) GetIpsecProfileOk() (*BulkTunnelRequestIpsecProfile, bool)`

GetIpsecProfileOk returns a tuple with the IpsecProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecProfile

`func (o *PatchedBulkTunnelRequest) SetIpsecProfile(v BulkTunnelRequestIpsecProfile)`

SetIpsecProfile sets IpsecProfile field to given value.

### HasIpsecProfile

`func (o *PatchedBulkTunnelRequest) HasIpsecProfile() bool`

HasIpsecProfile returns a boolean if a field has been set.

### SetIpsecProfileNil

`func (o *PatchedBulkTunnelRequest) SetIpsecProfileNil(b bool)`

 SetIpsecProfileNil sets the value for IpsecProfile to be an explicit nil

### UnsetIpsecProfile
`func (o *PatchedBulkTunnelRequest) UnsetIpsecProfile()`

UnsetIpsecProfile ensures that no value is present for IpsecProfile, not even an explicit nil
### GetTenant

`func (o *PatchedBulkTunnelRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkTunnelRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkTunnelRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkTunnelRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkTunnelRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkTunnelRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTunnelId

`func (o *PatchedBulkTunnelRequest) GetTunnelId() int64`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *PatchedBulkTunnelRequest) GetTunnelIdOk() (*int64, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *PatchedBulkTunnelRequest) SetTunnelId(v int64)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *PatchedBulkTunnelRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### SetTunnelIdNil

`func (o *PatchedBulkTunnelRequest) SetTunnelIdNil(b bool)`

 SetTunnelIdNil sets the value for TunnelId to be an explicit nil

### UnsetTunnelId
`func (o *PatchedBulkTunnelRequest) UnsetTunnelId()`

UnsetTunnelId ensures that no value is present for TunnelId, not even an explicit nil
### GetDescription

`func (o *PatchedBulkTunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkTunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkTunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkTunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkTunnelRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkTunnelRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkTunnelRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkTunnelRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkTunnelRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkTunnelRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkTunnelRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkTunnelRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkTunnelRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkTunnelRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkTunnelRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkTunnelRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkTunnelRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkTunnelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkTunnelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkTunnelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkTunnelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkTunnelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


