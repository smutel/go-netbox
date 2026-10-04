# PatchedVirtualMachineTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewPatchedVirtualMachineTypeRequest

`func NewPatchedVirtualMachineTypeRequest() *PatchedVirtualMachineTypeRequest`

NewPatchedVirtualMachineTypeRequest instantiates a new PatchedVirtualMachineTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVirtualMachineTypeRequestWithDefaults

`func NewPatchedVirtualMachineTypeRequestWithDefaults() *PatchedVirtualMachineTypeRequest`

NewPatchedVirtualMachineTypeRequestWithDefaults instantiates a new PatchedVirtualMachineTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedVirtualMachineTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedVirtualMachineTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedVirtualMachineTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedVirtualMachineTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedVirtualMachineTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedVirtualMachineTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedVirtualMachineTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedVirtualMachineTypeRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDefaultPlatform

`func (o *PatchedVirtualMachineTypeRequest) GetDefaultPlatform() BulkDeviceRequestPlatform`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *PatchedVirtualMachineTypeRequest) GetDefaultPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *PatchedVirtualMachineTypeRequest) SetDefaultPlatform(v BulkDeviceRequestPlatform)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *PatchedVirtualMachineTypeRequest) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *PatchedVirtualMachineTypeRequest) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *PatchedVirtualMachineTypeRequest) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetDefaultVcpus

`func (o *PatchedVirtualMachineTypeRequest) GetDefaultVcpus() float64`

GetDefaultVcpus returns the DefaultVcpus field if non-nil, zero value otherwise.

### GetDefaultVcpusOk

`func (o *PatchedVirtualMachineTypeRequest) GetDefaultVcpusOk() (*float64, bool)`

GetDefaultVcpusOk returns a tuple with the DefaultVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVcpus

`func (o *PatchedVirtualMachineTypeRequest) SetDefaultVcpus(v float64)`

SetDefaultVcpus sets DefaultVcpus field to given value.

### HasDefaultVcpus

`func (o *PatchedVirtualMachineTypeRequest) HasDefaultVcpus() bool`

HasDefaultVcpus returns a boolean if a field has been set.

### SetDefaultVcpusNil

`func (o *PatchedVirtualMachineTypeRequest) SetDefaultVcpusNil(b bool)`

 SetDefaultVcpusNil sets the value for DefaultVcpus to be an explicit nil

### UnsetDefaultVcpus
`func (o *PatchedVirtualMachineTypeRequest) UnsetDefaultVcpus()`

UnsetDefaultVcpus ensures that no value is present for DefaultVcpus, not even an explicit nil
### GetDefaultMemory

`func (o *PatchedVirtualMachineTypeRequest) GetDefaultMemory() int32`

GetDefaultMemory returns the DefaultMemory field if non-nil, zero value otherwise.

### GetDefaultMemoryOk

`func (o *PatchedVirtualMachineTypeRequest) GetDefaultMemoryOk() (*int32, bool)`

GetDefaultMemoryOk returns a tuple with the DefaultMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemory

`func (o *PatchedVirtualMachineTypeRequest) SetDefaultMemory(v int32)`

SetDefaultMemory sets DefaultMemory field to given value.

### HasDefaultMemory

`func (o *PatchedVirtualMachineTypeRequest) HasDefaultMemory() bool`

HasDefaultMemory returns a boolean if a field has been set.

### SetDefaultMemoryNil

`func (o *PatchedVirtualMachineTypeRequest) SetDefaultMemoryNil(b bool)`

 SetDefaultMemoryNil sets the value for DefaultMemory to be an explicit nil

### UnsetDefaultMemory
`func (o *PatchedVirtualMachineTypeRequest) UnsetDefaultMemory()`

UnsetDefaultMemory ensures that no value is present for DefaultMemory, not even an explicit nil
### GetDescription

`func (o *PatchedVirtualMachineTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedVirtualMachineTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedVirtualMachineTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedVirtualMachineTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedVirtualMachineTypeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedVirtualMachineTypeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedVirtualMachineTypeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedVirtualMachineTypeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedVirtualMachineTypeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedVirtualMachineTypeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedVirtualMachineTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedVirtualMachineTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedVirtualMachineTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedVirtualMachineTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedVirtualMachineTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedVirtualMachineTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedVirtualMachineTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedVirtualMachineTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedVirtualMachineTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVirtualMachineTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVirtualMachineTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVirtualMachineTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


