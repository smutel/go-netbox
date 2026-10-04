# BulkModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
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

### NewBulkModuleRequest

`func NewBulkModuleRequest(id int32, device BriefCoolingIntakeRequestDevice, moduleBay NestedModuleBayRequest, moduleType BulkModuleRequestModuleType, ) *BulkModuleRequest`

NewBulkModuleRequest instantiates a new BulkModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkModuleRequestWithDefaults

`func NewBulkModuleRequestWithDefaults() *BulkModuleRequest`

NewBulkModuleRequestWithDefaults instantiates a new BulkModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkModuleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkModuleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkModuleRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *BulkModuleRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkModuleRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkModuleRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetModuleBay

`func (o *BulkModuleRequest) GetModuleBay() NestedModuleBayRequest`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *BulkModuleRequest) GetModuleBayOk() (*NestedModuleBayRequest, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *BulkModuleRequest) SetModuleBay(v NestedModuleBayRequest)`

SetModuleBay sets ModuleBay field to given value.


### GetModuleType

`func (o *BulkModuleRequest) GetModuleType() BulkModuleRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *BulkModuleRequest) GetModuleTypeOk() (*BulkModuleRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *BulkModuleRequest) SetModuleType(v BulkModuleRequestModuleType)`

SetModuleType sets ModuleType field to given value.


### GetStatus

`func (o *BulkModuleRequest) GetStatus() BulkInventoryItemRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkModuleRequest) GetStatusOk() (*BulkInventoryItemRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkModuleRequest) SetStatus(v BulkInventoryItemRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkModuleRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSerial

`func (o *BulkModuleRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *BulkModuleRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *BulkModuleRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *BulkModuleRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *BulkModuleRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *BulkModuleRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *BulkModuleRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *BulkModuleRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *BulkModuleRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *BulkModuleRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDescription

`func (o *BulkModuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkModuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkModuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkModuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkModuleRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkModuleRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkModuleRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkModuleRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkModuleRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkModuleRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkModuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkModuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkModuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkModuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkModuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkModuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkModuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkModuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkModuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkModuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkModuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkModuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetReplicateComponents

`func (o *BulkModuleRequest) GetReplicateComponents() bool`

GetReplicateComponents returns the ReplicateComponents field if non-nil, zero value otherwise.

### GetReplicateComponentsOk

`func (o *BulkModuleRequest) GetReplicateComponentsOk() (*bool, bool)`

GetReplicateComponentsOk returns a tuple with the ReplicateComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateComponents

`func (o *BulkModuleRequest) SetReplicateComponents(v bool)`

SetReplicateComponents sets ReplicateComponents field to given value.

### HasReplicateComponents

`func (o *BulkModuleRequest) HasReplicateComponents() bool`

HasReplicateComponents returns a boolean if a field has been set.

### GetAdoptComponents

`func (o *BulkModuleRequest) GetAdoptComponents() bool`

GetAdoptComponents returns the AdoptComponents field if non-nil, zero value otherwise.

### GetAdoptComponentsOk

`func (o *BulkModuleRequest) GetAdoptComponentsOk() (*bool, bool)`

GetAdoptComponentsOk returns a tuple with the AdoptComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptComponents

`func (o *BulkModuleRequest) SetAdoptComponents(v bool)`

SetAdoptComponents sets AdoptComponents field to given value.

### HasAdoptComponents

`func (o *BulkModuleRequest) HasAdoptComponents() bool`

HasAdoptComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


