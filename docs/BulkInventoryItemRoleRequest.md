# BulkInventoryItemRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkInventoryItemRoleRequest

`func NewBulkInventoryItemRoleRequest(id int32, name string, slug string, ) *BulkInventoryItemRoleRequest`

NewBulkInventoryItemRoleRequest instantiates a new BulkInventoryItemRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInventoryItemRoleRequestWithDefaults

`func NewBulkInventoryItemRoleRequestWithDefaults() *BulkInventoryItemRoleRequest`

NewBulkInventoryItemRoleRequestWithDefaults instantiates a new BulkInventoryItemRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkInventoryItemRoleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkInventoryItemRoleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkInventoryItemRoleRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkInventoryItemRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkInventoryItemRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkInventoryItemRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BulkInventoryItemRoleRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BulkInventoryItemRoleRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BulkInventoryItemRoleRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetColor

`func (o *BulkInventoryItemRoleRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BulkInventoryItemRoleRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BulkInventoryItemRoleRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *BulkInventoryItemRoleRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *BulkInventoryItemRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkInventoryItemRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkInventoryItemRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkInventoryItemRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkInventoryItemRoleRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkInventoryItemRoleRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkInventoryItemRoleRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkInventoryItemRoleRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkInventoryItemRoleRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkInventoryItemRoleRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkInventoryItemRoleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkInventoryItemRoleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkInventoryItemRoleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkInventoryItemRoleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkInventoryItemRoleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkInventoryItemRoleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkInventoryItemRoleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkInventoryItemRoleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkInventoryItemRoleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkInventoryItemRoleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkInventoryItemRoleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkInventoryItemRoleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


