# PatchedBulkDeviceBayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | Pointer to [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstalledDevice** | Pointer to [**NullableBulkDeviceBayRequestInstalledDevice**](BulkDeviceBayRequestInstalledDevice.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkDeviceBayRequest

`func NewPatchedBulkDeviceBayRequest(id int32, ) *PatchedBulkDeviceBayRequest`

NewPatchedBulkDeviceBayRequest instantiates a new PatchedBulkDeviceBayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkDeviceBayRequestWithDefaults

`func NewPatchedBulkDeviceBayRequestWithDefaults() *PatchedBulkDeviceBayRequest`

NewPatchedBulkDeviceBayRequestWithDefaults instantiates a new PatchedBulkDeviceBayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkDeviceBayRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkDeviceBayRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkDeviceBayRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *PatchedBulkDeviceBayRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkDeviceBayRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkDeviceBayRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkDeviceBayRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkDeviceBayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkDeviceBayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkDeviceBayRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkDeviceBayRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkDeviceBayRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkDeviceBayRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkDeviceBayRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkDeviceBayRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkDeviceBayRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkDeviceBayRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkDeviceBayRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkDeviceBayRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkDeviceBayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkDeviceBayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkDeviceBayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkDeviceBayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstalledDevice

`func (o *PatchedBulkDeviceBayRequest) GetInstalledDevice() BulkDeviceBayRequestInstalledDevice`

GetInstalledDevice returns the InstalledDevice field if non-nil, zero value otherwise.

### GetInstalledDeviceOk

`func (o *PatchedBulkDeviceBayRequest) GetInstalledDeviceOk() (*BulkDeviceBayRequestInstalledDevice, bool)`

GetInstalledDeviceOk returns a tuple with the InstalledDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledDevice

`func (o *PatchedBulkDeviceBayRequest) SetInstalledDevice(v BulkDeviceBayRequestInstalledDevice)`

SetInstalledDevice sets InstalledDevice field to given value.

### HasInstalledDevice

`func (o *PatchedBulkDeviceBayRequest) HasInstalledDevice() bool`

HasInstalledDevice returns a boolean if a field has been set.

### SetInstalledDeviceNil

`func (o *PatchedBulkDeviceBayRequest) SetInstalledDeviceNil(b bool)`

 SetInstalledDeviceNil sets the value for InstalledDevice to be an explicit nil

### UnsetInstalledDevice
`func (o *PatchedBulkDeviceBayRequest) UnsetInstalledDevice()`

UnsetInstalledDevice ensures that no value is present for InstalledDevice, not even an explicit nil
### GetOwner

`func (o *PatchedBulkDeviceBayRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkDeviceBayRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkDeviceBayRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkDeviceBayRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkDeviceBayRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkDeviceBayRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkDeviceBayRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkDeviceBayRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkDeviceBayRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkDeviceBayRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkDeviceBayRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkDeviceBayRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkDeviceBayRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkDeviceBayRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


