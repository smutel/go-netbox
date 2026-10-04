# PatchedBulkModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | Pointer to [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | [optional] 
**ModuleBay** | Pointer to [**NestedModuleBayRequest**](NestedModuleBayRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**BulkModuleRequestModuleType**](BulkModuleRequestModuleType.md) |  | [optional] 
**Status** | Pointer to [**BulkInventoryItemRequestStatus**](BulkInventoryItemRequestStatus.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ReplicateComponents** | Pointer to **bool** | Automatically populate components associated with this module type (default: true) | [optional] [default to true]
**AdoptComponents** | Pointer to **bool** | Adopt already existing components | [optional] [default to false]

## Methods

### NewPatchedBulkModuleRequest

`func NewPatchedBulkModuleRequest(id int32, ) *PatchedBulkModuleRequest`

NewPatchedBulkModuleRequest instantiates a new PatchedBulkModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkModuleRequestWithDefaults

`func NewPatchedBulkModuleRequestWithDefaults() *PatchedBulkModuleRequest`

NewPatchedBulkModuleRequestWithDefaults instantiates a new PatchedBulkModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkModuleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkModuleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkModuleRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *PatchedBulkModuleRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkModuleRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkModuleRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkModuleRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModuleBay

`func (o *PatchedBulkModuleRequest) GetModuleBay() NestedModuleBayRequest`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *PatchedBulkModuleRequest) GetModuleBayOk() (*NestedModuleBayRequest, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *PatchedBulkModuleRequest) SetModuleBay(v NestedModuleBayRequest)`

SetModuleBay sets ModuleBay field to given value.

### HasModuleBay

`func (o *PatchedBulkModuleRequest) HasModuleBay() bool`

HasModuleBay returns a boolean if a field has been set.

### GetModuleType

`func (o *PatchedBulkModuleRequest) GetModuleType() BulkModuleRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedBulkModuleRequest) GetModuleTypeOk() (*BulkModuleRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedBulkModuleRequest) SetModuleType(v BulkModuleRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedBulkModuleRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkModuleRequest) GetStatus() BulkInventoryItemRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkModuleRequest) GetStatusOk() (*BulkInventoryItemRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkModuleRequest) SetStatus(v BulkInventoryItemRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkModuleRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSerial

`func (o *PatchedBulkModuleRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedBulkModuleRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedBulkModuleRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedBulkModuleRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedBulkModuleRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedBulkModuleRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedBulkModuleRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedBulkModuleRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedBulkModuleRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedBulkModuleRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDescription

`func (o *PatchedBulkModuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkModuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkModuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkModuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkModuleRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkModuleRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkModuleRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkModuleRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkModuleRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkModuleRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkModuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkModuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkModuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkModuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkModuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkModuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkModuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkModuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkModuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkModuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkModuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkModuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetReplicateComponents

`func (o *PatchedBulkModuleRequest) GetReplicateComponents() bool`

GetReplicateComponents returns the ReplicateComponents field if non-nil, zero value otherwise.

### GetReplicateComponentsOk

`func (o *PatchedBulkModuleRequest) GetReplicateComponentsOk() (*bool, bool)`

GetReplicateComponentsOk returns a tuple with the ReplicateComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateComponents

`func (o *PatchedBulkModuleRequest) SetReplicateComponents(v bool)`

SetReplicateComponents sets ReplicateComponents field to given value.

### HasReplicateComponents

`func (o *PatchedBulkModuleRequest) HasReplicateComponents() bool`

HasReplicateComponents returns a boolean if a field has been set.

### GetAdoptComponents

`func (o *PatchedBulkModuleRequest) GetAdoptComponents() bool`

GetAdoptComponents returns the AdoptComponents field if non-nil, zero value otherwise.

### GetAdoptComponentsOk

`func (o *PatchedBulkModuleRequest) GetAdoptComponentsOk() (*bool, bool)`

GetAdoptComponentsOk returns a tuple with the AdoptComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptComponents

`func (o *PatchedBulkModuleRequest) SetAdoptComponents(v bool)`

SetAdoptComponents sets AdoptComponents field to given value.

### HasAdoptComponents

`func (o *PatchedBulkModuleRequest) HasAdoptComponents() bool`

HasAdoptComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


