# PatchedBulkVLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Site** | Pointer to [**NullableBulkVLANRequestSite**](BulkVLANRequestSite.md) |  | [optional] 
**Group** | Pointer to [**NullableBulkVLANRequestGroup**](BulkVLANRequestGroup.md) |  | [optional] 
**Vid** | Pointer to **int32** | Numeric VLAN ID (1-4094) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Status** | Pointer to [**BulkIPRangeRequestStatus**](BulkIPRangeRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableASNRequestRole**](ASNRequestRole.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**QinqRole** | Pointer to [**NullableBulkVLANRequestQinqRole**](BulkVLANRequestQinqRole.md) |  | [optional] 
**QinqSvlan** | Pointer to [**NullableNestedVLANRequest**](NestedVLANRequest.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkVLANRequest

`func NewPatchedBulkVLANRequest(id int32, ) *PatchedBulkVLANRequest`

NewPatchedBulkVLANRequest instantiates a new PatchedBulkVLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkVLANRequestWithDefaults

`func NewPatchedBulkVLANRequestWithDefaults() *PatchedBulkVLANRequest`

NewPatchedBulkVLANRequestWithDefaults instantiates a new PatchedBulkVLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkVLANRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkVLANRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkVLANRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSite

`func (o *PatchedBulkVLANRequest) GetSite() BulkVLANRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedBulkVLANRequest) GetSiteOk() (*BulkVLANRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedBulkVLANRequest) SetSite(v BulkVLANRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedBulkVLANRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *PatchedBulkVLANRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *PatchedBulkVLANRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetGroup

`func (o *PatchedBulkVLANRequest) GetGroup() BulkVLANRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkVLANRequest) GetGroupOk() (*BulkVLANRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkVLANRequest) SetGroup(v BulkVLANRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkVLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedBulkVLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedBulkVLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVid

`func (o *PatchedBulkVLANRequest) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *PatchedBulkVLANRequest) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *PatchedBulkVLANRequest) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *PatchedBulkVLANRequest) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkVLANRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkVLANRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkVLANRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkVLANRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkVLANRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkVLANRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkVLANRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkVLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkVLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkVLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *PatchedBulkVLANRequest) GetStatus() BulkIPRangeRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkVLANRequest) GetStatusOk() (*BulkIPRangeRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkVLANRequest) SetStatus(v BulkIPRangeRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkVLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkVLANRequest) GetRole() ASNRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkVLANRequest) GetRoleOk() (*ASNRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkVLANRequest) SetRole(v ASNRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkVLANRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkVLANRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkVLANRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *PatchedBulkVLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkVLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkVLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkVLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQinqRole

`func (o *PatchedBulkVLANRequest) GetQinqRole() BulkVLANRequestQinqRole`

GetQinqRole returns the QinqRole field if non-nil, zero value otherwise.

### GetQinqRoleOk

`func (o *PatchedBulkVLANRequest) GetQinqRoleOk() (*BulkVLANRequestQinqRole, bool)`

GetQinqRoleOk returns a tuple with the QinqRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqRole

`func (o *PatchedBulkVLANRequest) SetQinqRole(v BulkVLANRequestQinqRole)`

SetQinqRole sets QinqRole field to given value.

### HasQinqRole

`func (o *PatchedBulkVLANRequest) HasQinqRole() bool`

HasQinqRole returns a boolean if a field has been set.

### SetQinqRoleNil

`func (o *PatchedBulkVLANRequest) SetQinqRoleNil(b bool)`

 SetQinqRoleNil sets the value for QinqRole to be an explicit nil

### UnsetQinqRole
`func (o *PatchedBulkVLANRequest) UnsetQinqRole()`

UnsetQinqRole ensures that no value is present for QinqRole, not even an explicit nil
### GetQinqSvlan

`func (o *PatchedBulkVLANRequest) GetQinqSvlan() NestedVLANRequest`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *PatchedBulkVLANRequest) GetQinqSvlanOk() (*NestedVLANRequest, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *PatchedBulkVLANRequest) SetQinqSvlan(v NestedVLANRequest)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *PatchedBulkVLANRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *PatchedBulkVLANRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *PatchedBulkVLANRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetOwner

`func (o *PatchedBulkVLANRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkVLANRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkVLANRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkVLANRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkVLANRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkVLANRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkVLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkVLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkVLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkVLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkVLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkVLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkVLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkVLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkVLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkVLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkVLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkVLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


