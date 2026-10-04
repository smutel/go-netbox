# PatchedModuleBayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | [optional] 
**Module** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Position** | Pointer to **string** | Identifier to reference when renaming installed components | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ModuleBayTypes** | Pointer to **[]int32** |  | [optional] 
**InstalledModule** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedModuleBayRequest

`func NewPatchedModuleBayRequest() *PatchedModuleBayRequest`

NewPatchedModuleBayRequest instantiates a new PatchedModuleBayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedModuleBayRequestWithDefaults

`func NewPatchedModuleBayRequestWithDefaults() *PatchedModuleBayRequest`

NewPatchedModuleBayRequestWithDefaults instantiates a new PatchedModuleBayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PatchedModuleBayRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedModuleBayRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedModuleBayRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedModuleBayRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModule

`func (o *PatchedModuleBayRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PatchedModuleBayRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PatchedModuleBayRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *PatchedModuleBayRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *PatchedModuleBayRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *PatchedModuleBayRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *PatchedModuleBayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedModuleBayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedModuleBayRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedModuleBayRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedModuleBayRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedModuleBayRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedModuleBayRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedModuleBayRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPosition

`func (o *PatchedModuleBayRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedModuleBayRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedModuleBayRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedModuleBayRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedModuleBayRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedModuleBayRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedModuleBayRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedModuleBayRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedModuleBayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedModuleBayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedModuleBayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedModuleBayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModuleBayTypes

`func (o *PatchedModuleBayRequest) GetModuleBayTypes() []int32`

GetModuleBayTypes returns the ModuleBayTypes field if non-nil, zero value otherwise.

### GetModuleBayTypesOk

`func (o *PatchedModuleBayRequest) GetModuleBayTypesOk() (*[]int32, bool)`

GetModuleBayTypesOk returns a tuple with the ModuleBayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTypes

`func (o *PatchedModuleBayRequest) SetModuleBayTypes(v []int32)`

SetModuleBayTypes sets ModuleBayTypes field to given value.

### HasModuleBayTypes

`func (o *PatchedModuleBayRequest) HasModuleBayTypes() bool`

HasModuleBayTypes returns a boolean if a field has been set.

### GetInstalledModule

`func (o *PatchedModuleBayRequest) GetInstalledModule() BulkConsolePortRequestModule`

GetInstalledModule returns the InstalledModule field if non-nil, zero value otherwise.

### GetInstalledModuleOk

`func (o *PatchedModuleBayRequest) GetInstalledModuleOk() (*BulkConsolePortRequestModule, bool)`

GetInstalledModuleOk returns a tuple with the InstalledModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledModule

`func (o *PatchedModuleBayRequest) SetInstalledModule(v BulkConsolePortRequestModule)`

SetInstalledModule sets InstalledModule field to given value.

### HasInstalledModule

`func (o *PatchedModuleBayRequest) HasInstalledModule() bool`

HasInstalledModule returns a boolean if a field has been set.

### SetInstalledModuleNil

`func (o *PatchedModuleBayRequest) SetInstalledModuleNil(b bool)`

 SetInstalledModuleNil sets the value for InstalledModule to be an explicit nil

### UnsetInstalledModule
`func (o *PatchedModuleBayRequest) UnsetInstalledModule()`

UnsetInstalledModule ensures that no value is present for InstalledModule, not even an explicit nil
### GetOwner

`func (o *PatchedModuleBayRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedModuleBayRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedModuleBayRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedModuleBayRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedModuleBayRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedModuleBayRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedModuleBayRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedModuleBayRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedModuleBayRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedModuleBayRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedModuleBayRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedModuleBayRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedModuleBayRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedModuleBayRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


