# BulkIPRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**StartAddress** | **string** |  | 
**EndAddress** | **string** |  | 
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

### NewBulkIPRangeRequest

`func NewBulkIPRangeRequest(id int32, startAddress string, endAddress string, ) *BulkIPRangeRequest`

NewBulkIPRangeRequest instantiates a new BulkIPRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIPRangeRequestWithDefaults

`func NewBulkIPRangeRequestWithDefaults() *BulkIPRangeRequest`

NewBulkIPRangeRequestWithDefaults instantiates a new BulkIPRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkIPRangeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkIPRangeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkIPRangeRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetStartAddress

`func (o *BulkIPRangeRequest) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *BulkIPRangeRequest) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *BulkIPRangeRequest) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.


### GetEndAddress

`func (o *BulkIPRangeRequest) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *BulkIPRangeRequest) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *BulkIPRangeRequest) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.


### GetVrf

`func (o *BulkIPRangeRequest) GetVrf() BulkIPAddressRequestVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *BulkIPRangeRequest) GetVrfOk() (*BulkIPAddressRequestVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *BulkIPRangeRequest) SetVrf(v BulkIPAddressRequestVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *BulkIPRangeRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *BulkIPRangeRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *BulkIPRangeRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *BulkIPRangeRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkIPRangeRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkIPRangeRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkIPRangeRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkIPRangeRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkIPRangeRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *BulkIPRangeRequest) GetStatus() BulkIPRangeRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkIPRangeRequest) GetStatusOk() (*BulkIPRangeRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkIPRangeRequest) SetStatus(v BulkIPRangeRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkIPRangeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *BulkIPRangeRequest) GetRole() ASNRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkIPRangeRequest) GetRoleOk() (*ASNRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkIPRangeRequest) SetRole(v ASNRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkIPRangeRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkIPRangeRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkIPRangeRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *BulkIPRangeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkIPRangeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkIPRangeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkIPRangeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkIPRangeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkIPRangeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkIPRangeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkIPRangeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkIPRangeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkIPRangeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkIPRangeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkIPRangeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkIPRangeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkIPRangeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkIPRangeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkIPRangeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkIPRangeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkIPRangeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkIPRangeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkIPRangeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkIPRangeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkIPRangeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMarkPopulated

`func (o *BulkIPRangeRequest) GetMarkPopulated() bool`

GetMarkPopulated returns the MarkPopulated field if non-nil, zero value otherwise.

### GetMarkPopulatedOk

`func (o *BulkIPRangeRequest) GetMarkPopulatedOk() (*bool, bool)`

GetMarkPopulatedOk returns a tuple with the MarkPopulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPopulated

`func (o *BulkIPRangeRequest) SetMarkPopulated(v bool)`

SetMarkPopulated sets MarkPopulated field to given value.

### HasMarkPopulated

`func (o *BulkIPRangeRequest) HasMarkPopulated() bool`

HasMarkPopulated returns a boolean if a field has been set.

### GetMarkUtilized

`func (o *BulkIPRangeRequest) GetMarkUtilized() bool`

GetMarkUtilized returns the MarkUtilized field if non-nil, zero value otherwise.

### GetMarkUtilizedOk

`func (o *BulkIPRangeRequest) GetMarkUtilizedOk() (*bool, bool)`

GetMarkUtilizedOk returns a tuple with the MarkUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkUtilized

`func (o *BulkIPRangeRequest) SetMarkUtilized(v bool)`

SetMarkUtilized sets MarkUtilized field to given value.

### HasMarkUtilized

`func (o *BulkIPRangeRequest) HasMarkUtilized() bool`

HasMarkUtilized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


