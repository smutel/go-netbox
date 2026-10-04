# ModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | 
**ModuleBay** | [**NestedModuleBayRequest**](NestedModuleBayRequest.md) |  | 
**ModuleType** | [**BulkModuleRequestModuleType**](BulkModuleRequestModuleType.md) |  | 
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

### NewModuleRequest

`func NewModuleRequest(device BriefCoolingIntakeRequestDevice, moduleBay NestedModuleBayRequest, moduleType BulkModuleRequestModuleType, ) *ModuleRequest`

NewModuleRequest instantiates a new ModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleRequestWithDefaults

`func NewModuleRequestWithDefaults() *ModuleRequest`

NewModuleRequestWithDefaults instantiates a new ModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *ModuleRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ModuleRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ModuleRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetModuleBay

`func (o *ModuleRequest) GetModuleBay() NestedModuleBayRequest`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *ModuleRequest) GetModuleBayOk() (*NestedModuleBayRequest, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *ModuleRequest) SetModuleBay(v NestedModuleBayRequest)`

SetModuleBay sets ModuleBay field to given value.


### GetModuleType

`func (o *ModuleRequest) GetModuleType() BulkModuleRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *ModuleRequest) GetModuleTypeOk() (*BulkModuleRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *ModuleRequest) SetModuleType(v BulkModuleRequestModuleType)`

SetModuleType sets ModuleType field to given value.


### GetStatus

`func (o *ModuleRequest) GetStatus() BulkInventoryItemRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModuleRequest) GetStatusOk() (*BulkInventoryItemRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModuleRequest) SetStatus(v BulkInventoryItemRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModuleRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSerial

`func (o *ModuleRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ModuleRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ModuleRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ModuleRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *ModuleRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ModuleRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ModuleRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ModuleRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *ModuleRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *ModuleRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDescription

`func (o *ModuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *ModuleRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModuleRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModuleRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModuleRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ModuleRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ModuleRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *ModuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ModuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ModuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ModuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *ModuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ModuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetReplicateComponents

`func (o *ModuleRequest) GetReplicateComponents() bool`

GetReplicateComponents returns the ReplicateComponents field if non-nil, zero value otherwise.

### GetReplicateComponentsOk

`func (o *ModuleRequest) GetReplicateComponentsOk() (*bool, bool)`

GetReplicateComponentsOk returns a tuple with the ReplicateComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateComponents

`func (o *ModuleRequest) SetReplicateComponents(v bool)`

SetReplicateComponents sets ReplicateComponents field to given value.

### HasReplicateComponents

`func (o *ModuleRequest) HasReplicateComponents() bool`

HasReplicateComponents returns a boolean if a field has been set.

### GetAdoptComponents

`func (o *ModuleRequest) GetAdoptComponents() bool`

GetAdoptComponents returns the AdoptComponents field if non-nil, zero value otherwise.

### GetAdoptComponentsOk

`func (o *ModuleRequest) GetAdoptComponentsOk() (*bool, bool)`

GetAdoptComponentsOk returns a tuple with the AdoptComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptComponents

`func (o *ModuleRequest) SetAdoptComponents(v bool)`

SetAdoptComponents sets AdoptComponents field to given value.

### HasAdoptComponents

`func (o *ModuleRequest) HasAdoptComponents() bool`

HasAdoptComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


