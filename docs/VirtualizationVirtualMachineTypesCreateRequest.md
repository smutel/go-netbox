# VirtualizationVirtualMachineTypesCreateRequest

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
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVirtualizationVirtualMachineTypesCreateRequest

`func NewVirtualizationVirtualMachineTypesCreateRequest(name string, slug string, ) *VirtualizationVirtualMachineTypesCreateRequest`

NewVirtualizationVirtualMachineTypesCreateRequest instantiates a new VirtualizationVirtualMachineTypesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualMachineTypesCreateRequestWithDefaults

`func NewVirtualizationVirtualMachineTypesCreateRequestWithDefaults() *VirtualizationVirtualMachineTypesCreateRequest`

NewVirtualizationVirtualMachineTypesCreateRequestWithDefaults instantiates a new VirtualizationVirtualMachineTypesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDefaultPlatform

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetDefaultPlatform() BulkDeviceRequestPlatform`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetDefaultPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetDefaultPlatform(v BulkDeviceRequestPlatform)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *VirtualizationVirtualMachineTypesCreateRequest) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *VirtualizationVirtualMachineTypesCreateRequest) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetDefaultVcpus

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetDefaultVcpus() float64`

GetDefaultVcpus returns the DefaultVcpus field if non-nil, zero value otherwise.

### GetDefaultVcpusOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetDefaultVcpusOk() (*float64, bool)`

GetDefaultVcpusOk returns a tuple with the DefaultVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVcpus

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetDefaultVcpus(v float64)`

SetDefaultVcpus sets DefaultVcpus field to given value.

### HasDefaultVcpus

`func (o *VirtualizationVirtualMachineTypesCreateRequest) HasDefaultVcpus() bool`

HasDefaultVcpus returns a boolean if a field has been set.

### SetDefaultVcpusNil

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetDefaultVcpusNil(b bool)`

 SetDefaultVcpusNil sets the value for DefaultVcpus to be an explicit nil

### UnsetDefaultVcpus
`func (o *VirtualizationVirtualMachineTypesCreateRequest) UnsetDefaultVcpus()`

UnsetDefaultVcpus ensures that no value is present for DefaultVcpus, not even an explicit nil
### GetDefaultMemory

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetDefaultMemory() int32`

GetDefaultMemory returns the DefaultMemory field if non-nil, zero value otherwise.

### GetDefaultMemoryOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetDefaultMemoryOk() (*int32, bool)`

GetDefaultMemoryOk returns a tuple with the DefaultMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemory

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetDefaultMemory(v int32)`

SetDefaultMemory sets DefaultMemory field to given value.

### HasDefaultMemory

`func (o *VirtualizationVirtualMachineTypesCreateRequest) HasDefaultMemory() bool`

HasDefaultMemory returns a boolean if a field has been set.

### SetDefaultMemoryNil

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetDefaultMemoryNil(b bool)`

 SetDefaultMemoryNil sets the value for DefaultMemory to be an explicit nil

### UnsetDefaultMemory
`func (o *VirtualizationVirtualMachineTypesCreateRequest) UnsetDefaultMemory()`

UnsetDefaultMemory ensures that no value is present for DefaultMemory, not even an explicit nil
### GetDescription

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationVirtualMachineTypesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VirtualizationVirtualMachineTypesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VirtualizationVirtualMachineTypesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VirtualizationVirtualMachineTypesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualizationVirtualMachineTypesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualizationVirtualMachineTypesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualizationVirtualMachineTypesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualizationVirtualMachineTypesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


