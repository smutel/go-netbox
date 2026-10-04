# PatchedBulkPlatformRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Parent** | Pointer to [**NullableNestedPlatformRequest**](NestedPlatformRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to [**NullableBulkInventoryItemRequestManufacturer**](BulkInventoryItemRequestManufacturer.md) |  | [optional] 
**ConfigTemplate** | Pointer to [**NullableBulkDeviceRequestConfigTemplate**](BulkDeviceRequestConfigTemplate.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkPlatformRequest

`func NewPatchedBulkPlatformRequest(id int32, ) *PatchedBulkPlatformRequest`

NewPatchedBulkPlatformRequest instantiates a new PatchedBulkPlatformRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkPlatformRequestWithDefaults

`func NewPatchedBulkPlatformRequestWithDefaults() *PatchedBulkPlatformRequest`

NewPatchedBulkPlatformRequestWithDefaults instantiates a new PatchedBulkPlatformRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkPlatformRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkPlatformRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkPlatformRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetParent

`func (o *PatchedBulkPlatformRequest) GetParent() NestedPlatformRequest`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkPlatformRequest) GetParentOk() (*NestedPlatformRequest, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkPlatformRequest) SetParent(v NestedPlatformRequest)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkPlatformRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedBulkPlatformRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedBulkPlatformRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *PatchedBulkPlatformRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkPlatformRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkPlatformRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkPlatformRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedBulkPlatformRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedBulkPlatformRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedBulkPlatformRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedBulkPlatformRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetManufacturer

`func (o *PatchedBulkPlatformRequest) GetManufacturer() BulkInventoryItemRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *PatchedBulkPlatformRequest) GetManufacturerOk() (*BulkInventoryItemRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *PatchedBulkPlatformRequest) SetManufacturer(v BulkInventoryItemRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *PatchedBulkPlatformRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *PatchedBulkPlatformRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *PatchedBulkPlatformRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetConfigTemplate

`func (o *PatchedBulkPlatformRequest) GetConfigTemplate() BulkDeviceRequestConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *PatchedBulkPlatformRequest) GetConfigTemplateOk() (*BulkDeviceRequestConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *PatchedBulkPlatformRequest) SetConfigTemplate(v BulkDeviceRequestConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *PatchedBulkPlatformRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *PatchedBulkPlatformRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *PatchedBulkPlatformRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetDescription

`func (o *PatchedBulkPlatformRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkPlatformRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkPlatformRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkPlatformRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkPlatformRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkPlatformRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkPlatformRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkPlatformRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkPlatformRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkPlatformRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkPlatformRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkPlatformRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkPlatformRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkPlatformRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkPlatformRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkPlatformRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkPlatformRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkPlatformRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkPlatformRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkPlatformRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkPlatformRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkPlatformRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


