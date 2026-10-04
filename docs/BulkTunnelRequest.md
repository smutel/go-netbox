# BulkTunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Status** | [**BulkTunnelRequestStatus**](BulkTunnelRequestStatus.md) |  | 
**Group** | Pointer to [**NullableBulkTunnelRequestGroup**](BulkTunnelRequestGroup.md) |  | [optional] 
**Encapsulation** | [**BulkTunnelRequestEncapsulation**](BulkTunnelRequestEncapsulation.md) |  | 
**IpsecProfile** | Pointer to [**NullableBulkTunnelRequestIpsecProfile**](BulkTunnelRequestIpsecProfile.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**TunnelId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkTunnelRequest

`func NewBulkTunnelRequest(id int32, name string, status BulkTunnelRequestStatus, encapsulation BulkTunnelRequestEncapsulation, ) *BulkTunnelRequest`

NewBulkTunnelRequest instantiates a new BulkTunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkTunnelRequestWithDefaults

`func NewBulkTunnelRequestWithDefaults() *BulkTunnelRequest`

NewBulkTunnelRequestWithDefaults instantiates a new BulkTunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkTunnelRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkTunnelRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkTunnelRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkTunnelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkTunnelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkTunnelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *BulkTunnelRequest) GetStatus() BulkTunnelRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkTunnelRequest) GetStatusOk() (*BulkTunnelRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkTunnelRequest) SetStatus(v BulkTunnelRequestStatus)`

SetStatus sets Status field to given value.


### GetGroup

`func (o *BulkTunnelRequest) GetGroup() BulkTunnelRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkTunnelRequest) GetGroupOk() (*BulkTunnelRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkTunnelRequest) SetGroup(v BulkTunnelRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BulkTunnelRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BulkTunnelRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BulkTunnelRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetEncapsulation

`func (o *BulkTunnelRequest) GetEncapsulation() BulkTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *BulkTunnelRequest) GetEncapsulationOk() (*BulkTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *BulkTunnelRequest) SetEncapsulation(v BulkTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.


### GetIpsecProfile

`func (o *BulkTunnelRequest) GetIpsecProfile() BulkTunnelRequestIpsecProfile`

GetIpsecProfile returns the IpsecProfile field if non-nil, zero value otherwise.

### GetIpsecProfileOk

`func (o *BulkTunnelRequest) GetIpsecProfileOk() (*BulkTunnelRequestIpsecProfile, bool)`

GetIpsecProfileOk returns a tuple with the IpsecProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecProfile

`func (o *BulkTunnelRequest) SetIpsecProfile(v BulkTunnelRequestIpsecProfile)`

SetIpsecProfile sets IpsecProfile field to given value.

### HasIpsecProfile

`func (o *BulkTunnelRequest) HasIpsecProfile() bool`

HasIpsecProfile returns a boolean if a field has been set.

### SetIpsecProfileNil

`func (o *BulkTunnelRequest) SetIpsecProfileNil(b bool)`

 SetIpsecProfileNil sets the value for IpsecProfile to be an explicit nil

### UnsetIpsecProfile
`func (o *BulkTunnelRequest) UnsetIpsecProfile()`

UnsetIpsecProfile ensures that no value is present for IpsecProfile, not even an explicit nil
### GetTenant

`func (o *BulkTunnelRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkTunnelRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkTunnelRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkTunnelRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkTunnelRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkTunnelRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTunnelId

`func (o *BulkTunnelRequest) GetTunnelId() int64`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *BulkTunnelRequest) GetTunnelIdOk() (*int64, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *BulkTunnelRequest) SetTunnelId(v int64)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *BulkTunnelRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### SetTunnelIdNil

`func (o *BulkTunnelRequest) SetTunnelIdNil(b bool)`

 SetTunnelIdNil sets the value for TunnelId to be an explicit nil

### UnsetTunnelId
`func (o *BulkTunnelRequest) UnsetTunnelId()`

UnsetTunnelId ensures that no value is present for TunnelId, not even an explicit nil
### GetDescription

`func (o *BulkTunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkTunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkTunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkTunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkTunnelRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkTunnelRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkTunnelRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkTunnelRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkTunnelRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkTunnelRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkTunnelRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkTunnelRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkTunnelRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkTunnelRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkTunnelRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkTunnelRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkTunnelRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkTunnelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkTunnelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkTunnelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkTunnelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkTunnelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


