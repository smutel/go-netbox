# VirtualMachineTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Slug** | **string** |  | 
**DefaultPlatform** | Pointer to [**NullableBulkDeviceRequestPlatform**](BulkDeviceRequestPlatform.md) |  | [optional] 
**DefaultVcpus** | Pointer to **NullableFloat64** |  | [optional] 
**DefaultMemory** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVirtualMachineTypeRequest

`func NewVirtualMachineTypeRequest(name string, slug string, ) *VirtualMachineTypeRequest`

NewVirtualMachineTypeRequest instantiates a new VirtualMachineTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineTypeRequestWithDefaults

`func NewVirtualMachineTypeRequestWithDefaults() *VirtualMachineTypeRequest`

NewVirtualMachineTypeRequestWithDefaults instantiates a new VirtualMachineTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualMachineTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachineTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachineTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *VirtualMachineTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VirtualMachineTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VirtualMachineTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDefaultPlatform

`func (o *VirtualMachineTypeRequest) GetDefaultPlatform() BulkDeviceRequestPlatform`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *VirtualMachineTypeRequest) GetDefaultPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *VirtualMachineTypeRequest) SetDefaultPlatform(v BulkDeviceRequestPlatform)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *VirtualMachineTypeRequest) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *VirtualMachineTypeRequest) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *VirtualMachineTypeRequest) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetDefaultVcpus

`func (o *VirtualMachineTypeRequest) GetDefaultVcpus() float64`

GetDefaultVcpus returns the DefaultVcpus field if non-nil, zero value otherwise.

### GetDefaultVcpusOk

`func (o *VirtualMachineTypeRequest) GetDefaultVcpusOk() (*float64, bool)`

GetDefaultVcpusOk returns a tuple with the DefaultVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVcpus

`func (o *VirtualMachineTypeRequest) SetDefaultVcpus(v float64)`

SetDefaultVcpus sets DefaultVcpus field to given value.

### HasDefaultVcpus

`func (o *VirtualMachineTypeRequest) HasDefaultVcpus() bool`

HasDefaultVcpus returns a boolean if a field has been set.

### SetDefaultVcpusNil

`func (o *VirtualMachineTypeRequest) SetDefaultVcpusNil(b bool)`

 SetDefaultVcpusNil sets the value for DefaultVcpus to be an explicit nil

### UnsetDefaultVcpus
`func (o *VirtualMachineTypeRequest) UnsetDefaultVcpus()`

UnsetDefaultVcpus ensures that no value is present for DefaultVcpus, not even an explicit nil
### GetDefaultMemory

`func (o *VirtualMachineTypeRequest) GetDefaultMemory() int32`

GetDefaultMemory returns the DefaultMemory field if non-nil, zero value otherwise.

### GetDefaultMemoryOk

`func (o *VirtualMachineTypeRequest) GetDefaultMemoryOk() (*int32, bool)`

GetDefaultMemoryOk returns a tuple with the DefaultMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemory

`func (o *VirtualMachineTypeRequest) SetDefaultMemory(v int32)`

SetDefaultMemory sets DefaultMemory field to given value.

### HasDefaultMemory

`func (o *VirtualMachineTypeRequest) HasDefaultMemory() bool`

HasDefaultMemory returns a boolean if a field has been set.

### SetDefaultMemoryNil

`func (o *VirtualMachineTypeRequest) SetDefaultMemoryNil(b bool)`

 SetDefaultMemoryNil sets the value for DefaultMemory to be an explicit nil

### UnsetDefaultMemory
`func (o *VirtualMachineTypeRequest) UnsetDefaultMemory()`

UnsetDefaultMemory ensures that no value is present for DefaultMemory, not even an explicit nil
### GetDescription

`func (o *VirtualMachineTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualMachineTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualMachineTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualMachineTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *VirtualMachineTypeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VirtualMachineTypeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VirtualMachineTypeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VirtualMachineTypeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VirtualMachineTypeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VirtualMachineTypeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VirtualMachineTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VirtualMachineTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VirtualMachineTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VirtualMachineTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VirtualMachineTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualMachineTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualMachineTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualMachineTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualMachineTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualMachineTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualMachineTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualMachineTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


