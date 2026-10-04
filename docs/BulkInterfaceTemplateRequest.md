# BulkInterfaceTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**BulkInterfaceRequestType**](BulkInterfaceRequestType.md) |  | 
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

### NewBulkInterfaceTemplateRequest

`func NewBulkInterfaceTemplateRequest(id int32, name string, type_ BulkInterfaceRequestType, ) *BulkInterfaceTemplateRequest`

NewBulkInterfaceTemplateRequest instantiates a new BulkInterfaceTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInterfaceTemplateRequestWithDefaults

`func NewBulkInterfaceTemplateRequestWithDefaults() *BulkInterfaceTemplateRequest`

NewBulkInterfaceTemplateRequestWithDefaults instantiates a new BulkInterfaceTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkInterfaceTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkInterfaceTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkInterfaceTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *BulkInterfaceTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkInterfaceTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkInterfaceTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *BulkInterfaceTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *BulkInterfaceTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *BulkInterfaceTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *BulkInterfaceTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *BulkInterfaceTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *BulkInterfaceTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *BulkInterfaceTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *BulkInterfaceTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *BulkInterfaceTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *BulkInterfaceTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkInterfaceTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkInterfaceTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkInterfaceTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkInterfaceTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkInterfaceTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkInterfaceTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *BulkInterfaceTemplateRequest) GetType() BulkInterfaceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkInterfaceTemplateRequest) GetTypeOk() (*BulkInterfaceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkInterfaceTemplateRequest) SetType(v BulkInterfaceRequestType)`

SetType sets Type field to given value.


### GetChannels

`func (o *BulkInterfaceTemplateRequest) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *BulkInterfaceTemplateRequest) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *BulkInterfaceTemplateRequest) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *BulkInterfaceTemplateRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *BulkInterfaceTemplateRequest) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *BulkInterfaceTemplateRequest) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetChannelId

`func (o *BulkInterfaceTemplateRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *BulkInterfaceTemplateRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *BulkInterfaceTemplateRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *BulkInterfaceTemplateRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *BulkInterfaceTemplateRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *BulkInterfaceTemplateRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetEnabled

`func (o *BulkInterfaceTemplateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkInterfaceTemplateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkInterfaceTemplateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkInterfaceTemplateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMgmtOnly

`func (o *BulkInterfaceTemplateRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *BulkInterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *BulkInterfaceTemplateRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *BulkInterfaceTemplateRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *BulkInterfaceTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkInterfaceTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkInterfaceTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkInterfaceTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *BulkInterfaceTemplateRequest) GetParent() NestedInterfaceTemplateRequest`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BulkInterfaceTemplateRequest) GetParentOk() (*NestedInterfaceTemplateRequest, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BulkInterfaceTemplateRequest) SetParent(v NestedInterfaceTemplateRequest)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BulkInterfaceTemplateRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *BulkInterfaceTemplateRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *BulkInterfaceTemplateRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *BulkInterfaceTemplateRequest) GetBridge() NestedInterfaceTemplateRequest`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *BulkInterfaceTemplateRequest) GetBridgeOk() (*NestedInterfaceTemplateRequest, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *BulkInterfaceTemplateRequest) SetBridge(v NestedInterfaceTemplateRequest)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *BulkInterfaceTemplateRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *BulkInterfaceTemplateRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *BulkInterfaceTemplateRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetPoeMode

`func (o *BulkInterfaceTemplateRequest) GetPoeMode() BulkInterfaceTemplateRequestPoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *BulkInterfaceTemplateRequest) GetPoeModeOk() (*BulkInterfaceTemplateRequestPoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *BulkInterfaceTemplateRequest) SetPoeMode(v BulkInterfaceTemplateRequestPoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *BulkInterfaceTemplateRequest) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### SetPoeModeNil

`func (o *BulkInterfaceTemplateRequest) SetPoeModeNil(b bool)`

 SetPoeModeNil sets the value for PoeMode to be an explicit nil

### UnsetPoeMode
`func (o *BulkInterfaceTemplateRequest) UnsetPoeMode()`

UnsetPoeMode ensures that no value is present for PoeMode, not even an explicit nil
### GetPoeType

`func (o *BulkInterfaceTemplateRequest) GetPoeType() BulkInterfaceTemplateRequestPoeType`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *BulkInterfaceTemplateRequest) GetPoeTypeOk() (*BulkInterfaceTemplateRequestPoeType, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *BulkInterfaceTemplateRequest) SetPoeType(v BulkInterfaceTemplateRequestPoeType)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *BulkInterfaceTemplateRequest) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### SetPoeTypeNil

`func (o *BulkInterfaceTemplateRequest) SetPoeTypeNil(b bool)`

 SetPoeTypeNil sets the value for PoeType to be an explicit nil

### UnsetPoeType
`func (o *BulkInterfaceTemplateRequest) UnsetPoeType()`

UnsetPoeType ensures that no value is present for PoeType, not even an explicit nil
### GetRfRole

`func (o *BulkInterfaceTemplateRequest) GetRfRole() BulkInterfaceTemplateRequestRfRole`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *BulkInterfaceTemplateRequest) GetRfRoleOk() (*BulkInterfaceTemplateRequestRfRole, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *BulkInterfaceTemplateRequest) SetRfRole(v BulkInterfaceTemplateRequestRfRole)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *BulkInterfaceTemplateRequest) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### SetRfRoleNil

`func (o *BulkInterfaceTemplateRequest) SetRfRoleNil(b bool)`

 SetRfRoleNil sets the value for RfRole to be an explicit nil

### UnsetRfRole
`func (o *BulkInterfaceTemplateRequest) UnsetRfRole()`

UnsetRfRole ensures that no value is present for RfRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


