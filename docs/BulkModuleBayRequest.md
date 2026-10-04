# BulkModuleBayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | 
**Module** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Name** | **string** |  | 
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

### NewBulkModuleBayRequest

`func NewBulkModuleBayRequest(id int32, device BriefCoolingIntakeRequestDevice, name string, ) *BulkModuleBayRequest`

NewBulkModuleBayRequest instantiates a new BulkModuleBayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkModuleBayRequestWithDefaults

`func NewBulkModuleBayRequestWithDefaults() *BulkModuleBayRequest`

NewBulkModuleBayRequestWithDefaults instantiates a new BulkModuleBayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkModuleBayRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkModuleBayRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkModuleBayRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *BulkModuleBayRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkModuleBayRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkModuleBayRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *BulkModuleBayRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *BulkModuleBayRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *BulkModuleBayRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *BulkModuleBayRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *BulkModuleBayRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *BulkModuleBayRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *BulkModuleBayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkModuleBayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkModuleBayRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkModuleBayRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkModuleBayRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkModuleBayRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkModuleBayRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPosition

`func (o *BulkModuleBayRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BulkModuleBayRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BulkModuleBayRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BulkModuleBayRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetEnabled

`func (o *BulkModuleBayRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkModuleBayRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkModuleBayRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkModuleBayRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *BulkModuleBayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkModuleBayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkModuleBayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkModuleBayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModuleBayTypes

`func (o *BulkModuleBayRequest) GetModuleBayTypes() []int32`

GetModuleBayTypes returns the ModuleBayTypes field if non-nil, zero value otherwise.

### GetModuleBayTypesOk

`func (o *BulkModuleBayRequest) GetModuleBayTypesOk() (*[]int32, bool)`

GetModuleBayTypesOk returns a tuple with the ModuleBayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTypes

`func (o *BulkModuleBayRequest) SetModuleBayTypes(v []int32)`

SetModuleBayTypes sets ModuleBayTypes field to given value.

### HasModuleBayTypes

`func (o *BulkModuleBayRequest) HasModuleBayTypes() bool`

HasModuleBayTypes returns a boolean if a field has been set.

### GetInstalledModule

`func (o *BulkModuleBayRequest) GetInstalledModule() BulkConsolePortRequestModule`

GetInstalledModule returns the InstalledModule field if non-nil, zero value otherwise.

### GetInstalledModuleOk

`func (o *BulkModuleBayRequest) GetInstalledModuleOk() (*BulkConsolePortRequestModule, bool)`

GetInstalledModuleOk returns a tuple with the InstalledModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledModule

`func (o *BulkModuleBayRequest) SetInstalledModule(v BulkConsolePortRequestModule)`

SetInstalledModule sets InstalledModule field to given value.

### HasInstalledModule

`func (o *BulkModuleBayRequest) HasInstalledModule() bool`

HasInstalledModule returns a boolean if a field has been set.

### SetInstalledModuleNil

`func (o *BulkModuleBayRequest) SetInstalledModuleNil(b bool)`

 SetInstalledModuleNil sets the value for InstalledModule to be an explicit nil

### UnsetInstalledModule
`func (o *BulkModuleBayRequest) UnsetInstalledModule()`

UnsetInstalledModule ensures that no value is present for InstalledModule, not even an explicit nil
### GetOwner

`func (o *BulkModuleBayRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkModuleBayRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkModuleBayRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkModuleBayRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkModuleBayRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkModuleBayRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *BulkModuleBayRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkModuleBayRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkModuleBayRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkModuleBayRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkModuleBayRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkModuleBayRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkModuleBayRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkModuleBayRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


