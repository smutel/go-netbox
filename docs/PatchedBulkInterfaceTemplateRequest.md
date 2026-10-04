# PatchedBulkInterfaceTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | Pointer to **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**BulkInterfaceRequestType**](BulkInterfaceRequestType.md) |  | [optional] 
**Channels** | Pointer to **NullableInt32** | The number of channels into which this interface is channelized | [optional] 
**ChannelId** | Pointer to **NullableInt32** | The channel on the parent interface to which this subinterface is bound | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MgmtOnly** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**NullableNestedInterfaceTemplateRequest**](NestedInterfaceTemplateRequest.md) |  | [optional] 
**Bridge** | Pointer to [**NullableNestedInterfaceTemplateRequest**](NestedInterfaceTemplateRequest.md) |  | [optional] 
**PoeMode** | Pointer to [**NullableBulkInterfaceTemplateRequestPoeMode**](BulkInterfaceTemplateRequestPoeMode.md) |  | [optional] 
**PoeType** | Pointer to [**NullableBulkInterfaceTemplateRequestPoeType**](BulkInterfaceTemplateRequestPoeType.md) |  | [optional] 
**RfRole** | Pointer to [**NullableBulkInterfaceTemplateRequestRfRole**](BulkInterfaceTemplateRequestRfRole.md) |  | [optional] 

## Methods

### NewPatchedBulkInterfaceTemplateRequest

`func NewPatchedBulkInterfaceTemplateRequest(id int32, ) *PatchedBulkInterfaceTemplateRequest`

NewPatchedBulkInterfaceTemplateRequest instantiates a new PatchedBulkInterfaceTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkInterfaceTemplateRequestWithDefaults

`func NewPatchedBulkInterfaceTemplateRequestWithDefaults() *PatchedBulkInterfaceTemplateRequest`

NewPatchedBulkInterfaceTemplateRequestWithDefaults instantiates a new PatchedBulkInterfaceTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkInterfaceTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkInterfaceTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *PatchedBulkInterfaceTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedBulkInterfaceTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedBulkInterfaceTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *PatchedBulkInterfaceTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedBulkInterfaceTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedBulkInterfaceTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *PatchedBulkInterfaceTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkInterfaceTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkInterfaceTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkInterfaceTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkInterfaceTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkInterfaceTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkInterfaceTemplateRequest) GetType() BulkInterfaceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetTypeOk() (*BulkInterfaceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkInterfaceTemplateRequest) SetType(v BulkInterfaceRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkInterfaceTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChannels

`func (o *PatchedBulkInterfaceTemplateRequest) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *PatchedBulkInterfaceTemplateRequest) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *PatchedBulkInterfaceTemplateRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetChannelId

`func (o *PatchedBulkInterfaceTemplateRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PatchedBulkInterfaceTemplateRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *PatchedBulkInterfaceTemplateRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetEnabled

`func (o *PatchedBulkInterfaceTemplateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkInterfaceTemplateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkInterfaceTemplateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMgmtOnly

`func (o *PatchedBulkInterfaceTemplateRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *PatchedBulkInterfaceTemplateRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *PatchedBulkInterfaceTemplateRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkInterfaceTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkInterfaceTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkInterfaceTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *PatchedBulkInterfaceTemplateRequest) GetParent() NestedInterfaceTemplateRequest`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetParentOk() (*NestedInterfaceTemplateRequest, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkInterfaceTemplateRequest) SetParent(v NestedInterfaceTemplateRequest)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkInterfaceTemplateRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *PatchedBulkInterfaceTemplateRequest) GetBridge() NestedInterfaceTemplateRequest`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetBridgeOk() (*NestedInterfaceTemplateRequest, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *PatchedBulkInterfaceTemplateRequest) SetBridge(v NestedInterfaceTemplateRequest)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *PatchedBulkInterfaceTemplateRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetPoeMode

`func (o *PatchedBulkInterfaceTemplateRequest) GetPoeMode() BulkInterfaceTemplateRequestPoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetPoeModeOk() (*BulkInterfaceTemplateRequestPoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *PatchedBulkInterfaceTemplateRequest) SetPoeMode(v BulkInterfaceTemplateRequestPoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *PatchedBulkInterfaceTemplateRequest) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### SetPoeModeNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetPoeModeNil(b bool)`

 SetPoeModeNil sets the value for PoeMode to be an explicit nil

### UnsetPoeMode
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetPoeMode()`

UnsetPoeMode ensures that no value is present for PoeMode, not even an explicit nil
### GetPoeType

`func (o *PatchedBulkInterfaceTemplateRequest) GetPoeType() BulkInterfaceTemplateRequestPoeType`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetPoeTypeOk() (*BulkInterfaceTemplateRequestPoeType, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *PatchedBulkInterfaceTemplateRequest) SetPoeType(v BulkInterfaceTemplateRequestPoeType)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *PatchedBulkInterfaceTemplateRequest) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### SetPoeTypeNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetPoeTypeNil(b bool)`

 SetPoeTypeNil sets the value for PoeType to be an explicit nil

### UnsetPoeType
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetPoeType()`

UnsetPoeType ensures that no value is present for PoeType, not even an explicit nil
### GetRfRole

`func (o *PatchedBulkInterfaceTemplateRequest) GetRfRole() BulkInterfaceTemplateRequestRfRole`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *PatchedBulkInterfaceTemplateRequest) GetRfRoleOk() (*BulkInterfaceTemplateRequestRfRole, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *PatchedBulkInterfaceTemplateRequest) SetRfRole(v BulkInterfaceTemplateRequestRfRole)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *PatchedBulkInterfaceTemplateRequest) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### SetRfRoleNil

`func (o *PatchedBulkInterfaceTemplateRequest) SetRfRoleNil(b bool)`

 SetRfRoleNil sets the value for RfRole to be an explicit nil

### UnsetRfRole
`func (o *PatchedBulkInterfaceTemplateRequest) UnsetRfRole()`

UnsetRfRole ensures that no value is present for RfRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


