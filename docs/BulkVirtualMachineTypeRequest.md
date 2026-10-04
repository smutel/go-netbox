# BulkVirtualMachineTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
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

### NewBulkVirtualMachineTypeRequest

`func NewBulkVirtualMachineTypeRequest(id int32, name string, slug string, ) *BulkVirtualMachineTypeRequest`

NewBulkVirtualMachineTypeRequest instantiates a new BulkVirtualMachineTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVirtualMachineTypeRequestWithDefaults

`func NewBulkVirtualMachineTypeRequestWithDefaults() *BulkVirtualMachineTypeRequest`

NewBulkVirtualMachineTypeRequestWithDefaults instantiates a new BulkVirtualMachineTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkVirtualMachineTypeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkVirtualMachineTypeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkVirtualMachineTypeRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkVirtualMachineTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkVirtualMachineTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkVirtualMachineTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BulkVirtualMachineTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BulkVirtualMachineTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BulkVirtualMachineTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDefaultPlatform

`func (o *BulkVirtualMachineTypeRequest) GetDefaultPlatform() BulkDeviceRequestPlatform`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *BulkVirtualMachineTypeRequest) GetDefaultPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *BulkVirtualMachineTypeRequest) SetDefaultPlatform(v BulkDeviceRequestPlatform)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *BulkVirtualMachineTypeRequest) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *BulkVirtualMachineTypeRequest) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *BulkVirtualMachineTypeRequest) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetDefaultVcpus

`func (o *BulkVirtualMachineTypeRequest) GetDefaultVcpus() float64`

GetDefaultVcpus returns the DefaultVcpus field if non-nil, zero value otherwise.

### GetDefaultVcpusOk

`func (o *BulkVirtualMachineTypeRequest) GetDefaultVcpusOk() (*float64, bool)`

GetDefaultVcpusOk returns a tuple with the DefaultVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVcpus

`func (o *BulkVirtualMachineTypeRequest) SetDefaultVcpus(v float64)`

SetDefaultVcpus sets DefaultVcpus field to given value.

### HasDefaultVcpus

`func (o *BulkVirtualMachineTypeRequest) HasDefaultVcpus() bool`

HasDefaultVcpus returns a boolean if a field has been set.

### SetDefaultVcpusNil

`func (o *BulkVirtualMachineTypeRequest) SetDefaultVcpusNil(b bool)`

 SetDefaultVcpusNil sets the value for DefaultVcpus to be an explicit nil

### UnsetDefaultVcpus
`func (o *BulkVirtualMachineTypeRequest) UnsetDefaultVcpus()`

UnsetDefaultVcpus ensures that no value is present for DefaultVcpus, not even an explicit nil
### GetDefaultMemory

`func (o *BulkVirtualMachineTypeRequest) GetDefaultMemory() int32`

GetDefaultMemory returns the DefaultMemory field if non-nil, zero value otherwise.

### GetDefaultMemoryOk

`func (o *BulkVirtualMachineTypeRequest) GetDefaultMemoryOk() (*int32, bool)`

GetDefaultMemoryOk returns a tuple with the DefaultMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemory

`func (o *BulkVirtualMachineTypeRequest) SetDefaultMemory(v int32)`

SetDefaultMemory sets DefaultMemory field to given value.

### HasDefaultMemory

`func (o *BulkVirtualMachineTypeRequest) HasDefaultMemory() bool`

HasDefaultMemory returns a boolean if a field has been set.

### SetDefaultMemoryNil

`func (o *BulkVirtualMachineTypeRequest) SetDefaultMemoryNil(b bool)`

 SetDefaultMemoryNil sets the value for DefaultMemory to be an explicit nil

### UnsetDefaultMemory
`func (o *BulkVirtualMachineTypeRequest) UnsetDefaultMemory()`

UnsetDefaultMemory ensures that no value is present for DefaultMemory, not even an explicit nil
### GetDescription

`func (o *BulkVirtualMachineTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkVirtualMachineTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkVirtualMachineTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkVirtualMachineTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkVirtualMachineTypeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkVirtualMachineTypeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkVirtualMachineTypeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkVirtualMachineTypeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkVirtualMachineTypeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkVirtualMachineTypeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkVirtualMachineTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkVirtualMachineTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkVirtualMachineTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkVirtualMachineTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkVirtualMachineTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkVirtualMachineTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkVirtualMachineTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkVirtualMachineTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkVirtualMachineTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkVirtualMachineTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkVirtualMachineTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkVirtualMachineTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


