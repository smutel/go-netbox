# PatchedBulkVirtualMachineTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**DefaultPlatform** | Pointer to [**NullableBulkDeviceRequestPlatform**](BulkDeviceRequestPlatform.md) |  | [optional] 
**DefaultVcpus** | Pointer to **NullableFloat64** |  | [optional] 
**DefaultMemory** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkVirtualMachineTypeRequest

`func NewPatchedBulkVirtualMachineTypeRequest(id int32, ) *PatchedBulkVirtualMachineTypeRequest`

NewPatchedBulkVirtualMachineTypeRequest instantiates a new PatchedBulkVirtualMachineTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkVirtualMachineTypeRequestWithDefaults

`func NewPatchedBulkVirtualMachineTypeRequestWithDefaults() *PatchedBulkVirtualMachineTypeRequest`

NewPatchedBulkVirtualMachineTypeRequestWithDefaults instantiates a new PatchedBulkVirtualMachineTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkVirtualMachineTypeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkVirtualMachineTypeRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkVirtualMachineTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkVirtualMachineTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkVirtualMachineTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedBulkVirtualMachineTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedBulkVirtualMachineTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedBulkVirtualMachineTypeRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDefaultPlatform

`func (o *PatchedBulkVirtualMachineTypeRequest) GetDefaultPlatform() BulkDeviceRequestPlatform`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetDefaultPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *PatchedBulkVirtualMachineTypeRequest) SetDefaultPlatform(v BulkDeviceRequestPlatform)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *PatchedBulkVirtualMachineTypeRequest) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *PatchedBulkVirtualMachineTypeRequest) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *PatchedBulkVirtualMachineTypeRequest) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetDefaultVcpus

`func (o *PatchedBulkVirtualMachineTypeRequest) GetDefaultVcpus() float64`

GetDefaultVcpus returns the DefaultVcpus field if non-nil, zero value otherwise.

### GetDefaultVcpusOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetDefaultVcpusOk() (*float64, bool)`

GetDefaultVcpusOk returns a tuple with the DefaultVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVcpus

`func (o *PatchedBulkVirtualMachineTypeRequest) SetDefaultVcpus(v float64)`

SetDefaultVcpus sets DefaultVcpus field to given value.

### HasDefaultVcpus

`func (o *PatchedBulkVirtualMachineTypeRequest) HasDefaultVcpus() bool`

HasDefaultVcpus returns a boolean if a field has been set.

### SetDefaultVcpusNil

`func (o *PatchedBulkVirtualMachineTypeRequest) SetDefaultVcpusNil(b bool)`

 SetDefaultVcpusNil sets the value for DefaultVcpus to be an explicit nil

### UnsetDefaultVcpus
`func (o *PatchedBulkVirtualMachineTypeRequest) UnsetDefaultVcpus()`

UnsetDefaultVcpus ensures that no value is present for DefaultVcpus, not even an explicit nil
### GetDefaultMemory

`func (o *PatchedBulkVirtualMachineTypeRequest) GetDefaultMemory() int32`

GetDefaultMemory returns the DefaultMemory field if non-nil, zero value otherwise.

### GetDefaultMemoryOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetDefaultMemoryOk() (*int32, bool)`

GetDefaultMemoryOk returns a tuple with the DefaultMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemory

`func (o *PatchedBulkVirtualMachineTypeRequest) SetDefaultMemory(v int32)`

SetDefaultMemory sets DefaultMemory field to given value.

### HasDefaultMemory

`func (o *PatchedBulkVirtualMachineTypeRequest) HasDefaultMemory() bool`

HasDefaultMemory returns a boolean if a field has been set.

### SetDefaultMemoryNil

`func (o *PatchedBulkVirtualMachineTypeRequest) SetDefaultMemoryNil(b bool)`

 SetDefaultMemoryNil sets the value for DefaultMemory to be an explicit nil

### UnsetDefaultMemory
`func (o *PatchedBulkVirtualMachineTypeRequest) UnsetDefaultMemory()`

UnsetDefaultMemory ensures that no value is present for DefaultMemory, not even an explicit nil
### GetDescription

`func (o *PatchedBulkVirtualMachineTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkVirtualMachineTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkVirtualMachineTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkVirtualMachineTypeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkVirtualMachineTypeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkVirtualMachineTypeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkVirtualMachineTypeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkVirtualMachineTypeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkVirtualMachineTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkVirtualMachineTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkVirtualMachineTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkVirtualMachineTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkVirtualMachineTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkVirtualMachineTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkVirtualMachineTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkVirtualMachineTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkVirtualMachineTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkVirtualMachineTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


