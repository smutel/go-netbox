# BulkVLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Site** | Pointer to [**NullableBulkVLANRequestSite**](BulkVLANRequestSite.md) |  | [optional] 
**Group** | Pointer to [**NullableBulkVLANRequestGroup**](BulkVLANRequestGroup.md) |  | [optional] 
**Vid** | **int32** | Numeric VLAN ID (1-4094) | 
**Name** | **string** |  | 
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

### NewBulkVLANRequest

`func NewBulkVLANRequest(id int32, vid int32, name string, ) *BulkVLANRequest`

NewBulkVLANRequest instantiates a new BulkVLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVLANRequestWithDefaults

`func NewBulkVLANRequestWithDefaults() *BulkVLANRequest`

NewBulkVLANRequestWithDefaults instantiates a new BulkVLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkVLANRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkVLANRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkVLANRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSite

`func (o *BulkVLANRequest) GetSite() BulkVLANRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *BulkVLANRequest) GetSiteOk() (*BulkVLANRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *BulkVLANRequest) SetSite(v BulkVLANRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *BulkVLANRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *BulkVLANRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *BulkVLANRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetGroup

`func (o *BulkVLANRequest) GetGroup() BulkVLANRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkVLANRequest) GetGroupOk() (*BulkVLANRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkVLANRequest) SetGroup(v BulkVLANRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BulkVLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BulkVLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BulkVLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVid

`func (o *BulkVLANRequest) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *BulkVLANRequest) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *BulkVLANRequest) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *BulkVLANRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkVLANRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkVLANRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTenant

`func (o *BulkVLANRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkVLANRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkVLANRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkVLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkVLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkVLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *BulkVLANRequest) GetStatus() BulkIPRangeRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkVLANRequest) GetStatusOk() (*BulkIPRangeRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkVLANRequest) SetStatus(v BulkIPRangeRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkVLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *BulkVLANRequest) GetRole() ASNRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkVLANRequest) GetRoleOk() (*ASNRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkVLANRequest) SetRole(v ASNRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkVLANRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkVLANRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkVLANRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *BulkVLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkVLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkVLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkVLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQinqRole

`func (o *BulkVLANRequest) GetQinqRole() BulkVLANRequestQinqRole`

GetQinqRole returns the QinqRole field if non-nil, zero value otherwise.

### GetQinqRoleOk

`func (o *BulkVLANRequest) GetQinqRoleOk() (*BulkVLANRequestQinqRole, bool)`

GetQinqRoleOk returns a tuple with the QinqRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqRole

`func (o *BulkVLANRequest) SetQinqRole(v BulkVLANRequestQinqRole)`

SetQinqRole sets QinqRole field to given value.

### HasQinqRole

`func (o *BulkVLANRequest) HasQinqRole() bool`

HasQinqRole returns a boolean if a field has been set.

### SetQinqRoleNil

`func (o *BulkVLANRequest) SetQinqRoleNil(b bool)`

 SetQinqRoleNil sets the value for QinqRole to be an explicit nil

### UnsetQinqRole
`func (o *BulkVLANRequest) UnsetQinqRole()`

UnsetQinqRole ensures that no value is present for QinqRole, not even an explicit nil
### GetQinqSvlan

`func (o *BulkVLANRequest) GetQinqSvlan() NestedVLANRequest`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *BulkVLANRequest) GetQinqSvlanOk() (*NestedVLANRequest, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *BulkVLANRequest) SetQinqSvlan(v NestedVLANRequest)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *BulkVLANRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *BulkVLANRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *BulkVLANRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetOwner

`func (o *BulkVLANRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkVLANRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkVLANRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkVLANRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkVLANRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkVLANRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkVLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkVLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkVLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkVLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkVLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkVLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkVLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkVLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkVLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkVLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkVLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkVLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


