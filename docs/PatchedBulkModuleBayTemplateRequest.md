# PatchedBulkModuleBayTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | Pointer to **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Position** | Pointer to **string** | Identifier to reference when renaming installed components | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ModuleBayTypes** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedBulkModuleBayTemplateRequest

`func NewPatchedBulkModuleBayTemplateRequest(id int32, ) *PatchedBulkModuleBayTemplateRequest`

NewPatchedBulkModuleBayTemplateRequest instantiates a new PatchedBulkModuleBayTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkModuleBayTemplateRequestWithDefaults

`func NewPatchedBulkModuleBayTemplateRequestWithDefaults() *PatchedBulkModuleBayTemplateRequest`

NewPatchedBulkModuleBayTemplateRequestWithDefaults instantiates a new PatchedBulkModuleBayTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkModuleBayTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkModuleBayTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *PatchedBulkModuleBayTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedBulkModuleBayTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedBulkModuleBayTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *PatchedBulkModuleBayTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *PatchedBulkModuleBayTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *PatchedBulkModuleBayTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedBulkModuleBayTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedBulkModuleBayTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *PatchedBulkModuleBayTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *PatchedBulkModuleBayTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *PatchedBulkModuleBayTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkModuleBayTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkModuleBayTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkModuleBayTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkModuleBayTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkModuleBayTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPosition

`func (o *PatchedBulkModuleBayTemplateRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedBulkModuleBayTemplateRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedBulkModuleBayTemplateRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkModuleBayTemplateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkModuleBayTemplateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkModuleBayTemplateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkModuleBayTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkModuleBayTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkModuleBayTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModuleBayTypes

`func (o *PatchedBulkModuleBayTemplateRequest) GetModuleBayTypes() []int32`

GetModuleBayTypes returns the ModuleBayTypes field if non-nil, zero value otherwise.

### GetModuleBayTypesOk

`func (o *PatchedBulkModuleBayTemplateRequest) GetModuleBayTypesOk() (*[]int32, bool)`

GetModuleBayTypesOk returns a tuple with the ModuleBayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTypes

`func (o *PatchedBulkModuleBayTemplateRequest) SetModuleBayTypes(v []int32)`

SetModuleBayTypes sets ModuleBayTypes field to given value.

### HasModuleBayTypes

`func (o *PatchedBulkModuleBayTemplateRequest) HasModuleBayTypes() bool`

HasModuleBayTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


