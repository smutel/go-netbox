# BulkModuleBayTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Position** | Pointer to **string** | Identifier to reference when renaming installed components | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ModuleBayTypes** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBulkModuleBayTemplateRequest

`func NewBulkModuleBayTemplateRequest(id int32, name string, ) *BulkModuleBayTemplateRequest`

NewBulkModuleBayTemplateRequest instantiates a new BulkModuleBayTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkModuleBayTemplateRequestWithDefaults

`func NewBulkModuleBayTemplateRequestWithDefaults() *BulkModuleBayTemplateRequest`

NewBulkModuleBayTemplateRequestWithDefaults instantiates a new BulkModuleBayTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkModuleBayTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkModuleBayTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkModuleBayTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *BulkModuleBayTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkModuleBayTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkModuleBayTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *BulkModuleBayTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *BulkModuleBayTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *BulkModuleBayTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *BulkModuleBayTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *BulkModuleBayTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *BulkModuleBayTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *BulkModuleBayTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *BulkModuleBayTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *BulkModuleBayTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *BulkModuleBayTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkModuleBayTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkModuleBayTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkModuleBayTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkModuleBayTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkModuleBayTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkModuleBayTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPosition

`func (o *BulkModuleBayTemplateRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BulkModuleBayTemplateRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BulkModuleBayTemplateRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BulkModuleBayTemplateRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetEnabled

`func (o *BulkModuleBayTemplateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkModuleBayTemplateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkModuleBayTemplateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkModuleBayTemplateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *BulkModuleBayTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkModuleBayTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkModuleBayTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkModuleBayTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModuleBayTypes

`func (o *BulkModuleBayTemplateRequest) GetModuleBayTypes() []int32`

GetModuleBayTypes returns the ModuleBayTypes field if non-nil, zero value otherwise.

### GetModuleBayTypesOk

`func (o *BulkModuleBayTemplateRequest) GetModuleBayTypesOk() (*[]int32, bool)`

GetModuleBayTypesOk returns a tuple with the ModuleBayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTypes

`func (o *BulkModuleBayTemplateRequest) SetModuleBayTypes(v []int32)`

SetModuleBayTypes sets ModuleBayTypes field to given value.

### HasModuleBayTypes

`func (o *BulkModuleBayTemplateRequest) HasModuleBayTypes() bool`

HasModuleBayTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


