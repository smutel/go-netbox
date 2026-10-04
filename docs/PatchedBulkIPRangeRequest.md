# PatchedBulkIPRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**StartAddress** | Pointer to **string** |  | [optional] 
**EndAddress** | Pointer to **string** |  | [optional] 
**Vrf** | Pointer to [**NullableBulkIPAddressRequestVrf**](BulkIPAddressRequestVrf.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Status** | Pointer to [**BulkIPRangeRequestStatus**](BulkIPRangeRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableASNRequestRole**](ASNRequestRole.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**MarkPopulated** | Pointer to **bool** | Prevent the creation of IP addresses within this range | [optional] 
**MarkUtilized** | Pointer to **bool** | Report space as fully utilized | [optional] 

## Methods

### NewPatchedBulkIPRangeRequest

`func NewPatchedBulkIPRangeRequest(id int32, ) *PatchedBulkIPRangeRequest`

NewPatchedBulkIPRangeRequest instantiates a new PatchedBulkIPRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkIPRangeRequestWithDefaults

`func NewPatchedBulkIPRangeRequestWithDefaults() *PatchedBulkIPRangeRequest`

NewPatchedBulkIPRangeRequestWithDefaults instantiates a new PatchedBulkIPRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkIPRangeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkIPRangeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkIPRangeRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetStartAddress

`func (o *PatchedBulkIPRangeRequest) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *PatchedBulkIPRangeRequest) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *PatchedBulkIPRangeRequest) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *PatchedBulkIPRangeRequest) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *PatchedBulkIPRangeRequest) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *PatchedBulkIPRangeRequest) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *PatchedBulkIPRangeRequest) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *PatchedBulkIPRangeRequest) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetVrf

`func (o *PatchedBulkIPRangeRequest) GetVrf() BulkIPAddressRequestVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedBulkIPRangeRequest) GetVrfOk() (*BulkIPAddressRequestVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedBulkIPRangeRequest) SetVrf(v BulkIPAddressRequestVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedBulkIPRangeRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PatchedBulkIPRangeRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PatchedBulkIPRangeRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *PatchedBulkIPRangeRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkIPRangeRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkIPRangeRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkIPRangeRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkIPRangeRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkIPRangeRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *PatchedBulkIPRangeRequest) GetStatus() BulkIPRangeRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkIPRangeRequest) GetStatusOk() (*BulkIPRangeRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkIPRangeRequest) SetStatus(v BulkIPRangeRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkIPRangeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkIPRangeRequest) GetRole() ASNRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkIPRangeRequest) GetRoleOk() (*ASNRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkIPRangeRequest) SetRole(v ASNRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkIPRangeRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkIPRangeRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkIPRangeRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *PatchedBulkIPRangeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkIPRangeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkIPRangeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkIPRangeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkIPRangeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkIPRangeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkIPRangeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkIPRangeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkIPRangeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkIPRangeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkIPRangeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkIPRangeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkIPRangeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkIPRangeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkIPRangeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkIPRangeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkIPRangeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkIPRangeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkIPRangeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkIPRangeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkIPRangeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkIPRangeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMarkPopulated

`func (o *PatchedBulkIPRangeRequest) GetMarkPopulated() bool`

GetMarkPopulated returns the MarkPopulated field if non-nil, zero value otherwise.

### GetMarkPopulatedOk

`func (o *PatchedBulkIPRangeRequest) GetMarkPopulatedOk() (*bool, bool)`

GetMarkPopulatedOk returns a tuple with the MarkPopulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPopulated

`func (o *PatchedBulkIPRangeRequest) SetMarkPopulated(v bool)`

SetMarkPopulated sets MarkPopulated field to given value.

### HasMarkPopulated

`func (o *PatchedBulkIPRangeRequest) HasMarkPopulated() bool`

HasMarkPopulated returns a boolean if a field has been set.

### GetMarkUtilized

`func (o *PatchedBulkIPRangeRequest) GetMarkUtilized() bool`

GetMarkUtilized returns the MarkUtilized field if non-nil, zero value otherwise.

### GetMarkUtilizedOk

`func (o *PatchedBulkIPRangeRequest) GetMarkUtilizedOk() (*bool, bool)`

GetMarkUtilizedOk returns a tuple with the MarkUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkUtilized

`func (o *PatchedBulkIPRangeRequest) SetMarkUtilized(v bool)`

SetMarkUtilized sets MarkUtilized field to given value.

### HasMarkUtilized

`func (o *PatchedBulkIPRangeRequest) HasMarkUtilized() bool`

HasMarkUtilized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


