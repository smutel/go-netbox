# BulkVirtualDeviceContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Device** | [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | 
**Identifier** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**Status** | [**BulkVirtualDeviceContextRequestStatus**](BulkVirtualDeviceContextRequestStatus.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkVirtualDeviceContextRequest

`func NewBulkVirtualDeviceContextRequest(id int32, name string, device BriefCoolingIntakeRequestDevice, status BulkVirtualDeviceContextRequestStatus, ) *BulkVirtualDeviceContextRequest`

NewBulkVirtualDeviceContextRequest instantiates a new BulkVirtualDeviceContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVirtualDeviceContextRequestWithDefaults

`func NewBulkVirtualDeviceContextRequestWithDefaults() *BulkVirtualDeviceContextRequest`

NewBulkVirtualDeviceContextRequestWithDefaults instantiates a new BulkVirtualDeviceContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkVirtualDeviceContextRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkVirtualDeviceContextRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkVirtualDeviceContextRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkVirtualDeviceContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkVirtualDeviceContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkVirtualDeviceContextRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDevice

`func (o *BulkVirtualDeviceContextRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkVirtualDeviceContextRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkVirtualDeviceContextRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetIdentifier

`func (o *BulkVirtualDeviceContextRequest) GetIdentifier() int32`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BulkVirtualDeviceContextRequest) GetIdentifierOk() (*int32, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BulkVirtualDeviceContextRequest) SetIdentifier(v int32)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BulkVirtualDeviceContextRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *BulkVirtualDeviceContextRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *BulkVirtualDeviceContextRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetTenant

`func (o *BulkVirtualDeviceContextRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkVirtualDeviceContextRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkVirtualDeviceContextRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkVirtualDeviceContextRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkVirtualDeviceContextRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkVirtualDeviceContextRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPrimaryIp4

`func (o *BulkVirtualDeviceContextRequest) GetPrimaryIp4() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *BulkVirtualDeviceContextRequest) GetPrimaryIp4Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *BulkVirtualDeviceContextRequest) SetPrimaryIp4(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *BulkVirtualDeviceContextRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *BulkVirtualDeviceContextRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *BulkVirtualDeviceContextRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *BulkVirtualDeviceContextRequest) GetPrimaryIp6() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *BulkVirtualDeviceContextRequest) GetPrimaryIp6Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *BulkVirtualDeviceContextRequest) SetPrimaryIp6(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *BulkVirtualDeviceContextRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *BulkVirtualDeviceContextRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *BulkVirtualDeviceContextRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetStatus

`func (o *BulkVirtualDeviceContextRequest) GetStatus() BulkVirtualDeviceContextRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkVirtualDeviceContextRequest) GetStatusOk() (*BulkVirtualDeviceContextRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkVirtualDeviceContextRequest) SetStatus(v BulkVirtualDeviceContextRequestStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *BulkVirtualDeviceContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkVirtualDeviceContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkVirtualDeviceContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkVirtualDeviceContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkVirtualDeviceContextRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkVirtualDeviceContextRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkVirtualDeviceContextRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkVirtualDeviceContextRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkVirtualDeviceContextRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkVirtualDeviceContextRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkVirtualDeviceContextRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkVirtualDeviceContextRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkVirtualDeviceContextRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkVirtualDeviceContextRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkVirtualDeviceContextRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkVirtualDeviceContextRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkVirtualDeviceContextRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkVirtualDeviceContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkVirtualDeviceContextRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkVirtualDeviceContextRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkVirtualDeviceContextRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkVirtualDeviceContextRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


