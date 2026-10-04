# ModuleBay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**BriefDevice**](BriefDevice.md) |  | 
**Module** | Pointer to [**NullableBriefModule**](BriefModule.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Position** | Pointer to **string** | Identifier to reference when renaming installed components | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ModuleBayTypes** | Pointer to [**[]BriefModuleBayType**](BriefModuleBayType.md) |  | [optional] 
**InstalledModule** | Pointer to [**NullableBriefModule**](BriefModule.md) |  | [optional] 
**Owner** | Pointer to [**NullableBriefOwner**](BriefOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Occupied** | **bool** |  | [readonly] 
**IsModuleCompatible** | **bool** |  | [readonly] 

## Methods

### NewModuleBay

`func NewModuleBay(id int32, url string, displayUrl string, display string, device BriefDevice, name string, created NullableTime, lastUpdated NullableTime, occupied bool, isModuleCompatible bool, ) *ModuleBay`

NewModuleBay instantiates a new ModuleBay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleBayWithDefaults

`func NewModuleBayWithDefaults() *ModuleBay`

NewModuleBayWithDefaults instantiates a new ModuleBay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModuleBay) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModuleBay) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModuleBay) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ModuleBay) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModuleBay) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModuleBay) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *ModuleBay) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ModuleBay) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ModuleBay) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *ModuleBay) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ModuleBay) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ModuleBay) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *ModuleBay) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ModuleBay) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ModuleBay) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *ModuleBay) GetModule() BriefModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ModuleBay) GetModuleOk() (*BriefModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ModuleBay) SetModule(v BriefModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *ModuleBay) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *ModuleBay) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *ModuleBay) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *ModuleBay) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModuleBay) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModuleBay) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ModuleBay) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModuleBay) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModuleBay) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ModuleBay) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPosition

`func (o *ModuleBay) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModuleBay) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModuleBay) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModuleBay) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetEnabled

`func (o *ModuleBay) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModuleBay) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModuleBay) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModuleBay) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *ModuleBay) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleBay) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleBay) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleBay) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModuleBayTypes

`func (o *ModuleBay) GetModuleBayTypes() []BriefModuleBayType`

GetModuleBayTypes returns the ModuleBayTypes field if non-nil, zero value otherwise.

### GetModuleBayTypesOk

`func (o *ModuleBay) GetModuleBayTypesOk() (*[]BriefModuleBayType, bool)`

GetModuleBayTypesOk returns a tuple with the ModuleBayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTypes

`func (o *ModuleBay) SetModuleBayTypes(v []BriefModuleBayType)`

SetModuleBayTypes sets ModuleBayTypes field to given value.

### HasModuleBayTypes

`func (o *ModuleBay) HasModuleBayTypes() bool`

HasModuleBayTypes returns a boolean if a field has been set.

### GetInstalledModule

`func (o *ModuleBay) GetInstalledModule() BriefModule`

GetInstalledModule returns the InstalledModule field if non-nil, zero value otherwise.

### GetInstalledModuleOk

`func (o *ModuleBay) GetInstalledModuleOk() (*BriefModule, bool)`

GetInstalledModuleOk returns a tuple with the InstalledModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledModule

`func (o *ModuleBay) SetInstalledModule(v BriefModule)`

SetInstalledModule sets InstalledModule field to given value.

### HasInstalledModule

`func (o *ModuleBay) HasInstalledModule() bool`

HasInstalledModule returns a boolean if a field has been set.

### SetInstalledModuleNil

`func (o *ModuleBay) SetInstalledModuleNil(b bool)`

 SetInstalledModuleNil sets the value for InstalledModule to be an explicit nil

### UnsetInstalledModule
`func (o *ModuleBay) UnsetInstalledModule()`

UnsetInstalledModule ensures that no value is present for InstalledModule, not even an explicit nil
### GetOwner

`func (o *ModuleBay) GetOwner() BriefOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModuleBay) GetOwnerOk() (*BriefOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModuleBay) SetOwner(v BriefOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModuleBay) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ModuleBay) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ModuleBay) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *ModuleBay) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModuleBay) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModuleBay) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModuleBay) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ModuleBay) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModuleBay) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModuleBay) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModuleBay) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *ModuleBay) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModuleBay) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModuleBay) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ModuleBay) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ModuleBay) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ModuleBay) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ModuleBay) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ModuleBay) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ModuleBay) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ModuleBay) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetOccupied

`func (o *ModuleBay) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *ModuleBay) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *ModuleBay) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.


### GetIsModuleCompatible

`func (o *ModuleBay) GetIsModuleCompatible() bool`

GetIsModuleCompatible returns the IsModuleCompatible field if non-nil, zero value otherwise.

### GetIsModuleCompatibleOk

`func (o *ModuleBay) GetIsModuleCompatibleOk() (*bool, bool)`

GetIsModuleCompatibleOk returns a tuple with the IsModuleCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModuleCompatible

`func (o *ModuleBay) SetIsModuleCompatible(v bool)`

SetIsModuleCompatible sets IsModuleCompatible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


