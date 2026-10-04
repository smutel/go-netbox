# DcimInterfaceTemplatesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Bridge** | Pointer to **NullableInt32** |  | [optional] 
**PoeMode** | Pointer to [**NullableBulkInterfaceTemplateRequestPoeMode**](BulkInterfaceTemplateRequestPoeMode.md) |  | [optional] 
**PoeType** | Pointer to [**NullableBulkInterfaceTemplateRequestPoeType**](BulkInterfaceTemplateRequestPoeType.md) |  | [optional] 
**RfRole** | Pointer to [**NullableWirelessRole**](WirelessRole.md) |  | [optional] 

## Methods

### NewDcimInterfaceTemplatesCreateRequest

`func NewDcimInterfaceTemplatesCreateRequest(name string, type_ BulkInterfaceRequestType, ) *DcimInterfaceTemplatesCreateRequest`

NewDcimInterfaceTemplatesCreateRequest instantiates a new DcimInterfaceTemplatesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimInterfaceTemplatesCreateRequestWithDefaults

`func NewDcimInterfaceTemplatesCreateRequestWithDefaults() *DcimInterfaceTemplatesCreateRequest`

NewDcimInterfaceTemplatesCreateRequestWithDefaults instantiates a new DcimInterfaceTemplatesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *DcimInterfaceTemplatesCreateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DcimInterfaceTemplatesCreateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DcimInterfaceTemplatesCreateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *DcimInterfaceTemplatesCreateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *DcimInterfaceTemplatesCreateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *DcimInterfaceTemplatesCreateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *DcimInterfaceTemplatesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimInterfaceTemplatesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *DcimInterfaceTemplatesCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DcimInterfaceTemplatesCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DcimInterfaceTemplatesCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *DcimInterfaceTemplatesCreateRequest) GetType() BulkInterfaceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetTypeOk() (*BulkInterfaceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcimInterfaceTemplatesCreateRequest) SetType(v BulkInterfaceRequestType)`

SetType sets Type field to given value.


### GetChannels

`func (o *DcimInterfaceTemplatesCreateRequest) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *DcimInterfaceTemplatesCreateRequest) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *DcimInterfaceTemplatesCreateRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetChannelId

`func (o *DcimInterfaceTemplatesCreateRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *DcimInterfaceTemplatesCreateRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *DcimInterfaceTemplatesCreateRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetEnabled

`func (o *DcimInterfaceTemplatesCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DcimInterfaceTemplatesCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DcimInterfaceTemplatesCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMgmtOnly

`func (o *DcimInterfaceTemplatesCreateRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *DcimInterfaceTemplatesCreateRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *DcimInterfaceTemplatesCreateRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *DcimInterfaceTemplatesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimInterfaceTemplatesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimInterfaceTemplatesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *DcimInterfaceTemplatesCreateRequest) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *DcimInterfaceTemplatesCreateRequest) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *DcimInterfaceTemplatesCreateRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *DcimInterfaceTemplatesCreateRequest) GetBridge() int32`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetBridgeOk() (*int32, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *DcimInterfaceTemplatesCreateRequest) SetBridge(v int32)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *DcimInterfaceTemplatesCreateRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetPoeMode

`func (o *DcimInterfaceTemplatesCreateRequest) GetPoeMode() BulkInterfaceTemplateRequestPoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetPoeModeOk() (*BulkInterfaceTemplateRequestPoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *DcimInterfaceTemplatesCreateRequest) SetPoeMode(v BulkInterfaceTemplateRequestPoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *DcimInterfaceTemplatesCreateRequest) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### SetPoeModeNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetPoeModeNil(b bool)`

 SetPoeModeNil sets the value for PoeMode to be an explicit nil

### UnsetPoeMode
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetPoeMode()`

UnsetPoeMode ensures that no value is present for PoeMode, not even an explicit nil
### GetPoeType

`func (o *DcimInterfaceTemplatesCreateRequest) GetPoeType() BulkInterfaceTemplateRequestPoeType`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetPoeTypeOk() (*BulkInterfaceTemplateRequestPoeType, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *DcimInterfaceTemplatesCreateRequest) SetPoeType(v BulkInterfaceTemplateRequestPoeType)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *DcimInterfaceTemplatesCreateRequest) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### SetPoeTypeNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetPoeTypeNil(b bool)`

 SetPoeTypeNil sets the value for PoeType to be an explicit nil

### UnsetPoeType
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetPoeType()`

UnsetPoeType ensures that no value is present for PoeType, not even an explicit nil
### GetRfRole

`func (o *DcimInterfaceTemplatesCreateRequest) GetRfRole() WirelessRole`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *DcimInterfaceTemplatesCreateRequest) GetRfRoleOk() (*WirelessRole, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *DcimInterfaceTemplatesCreateRequest) SetRfRole(v WirelessRole)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *DcimInterfaceTemplatesCreateRequest) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### SetRfRoleNil

`func (o *DcimInterfaceTemplatesCreateRequest) SetRfRoleNil(b bool)`

 SetRfRoleNil sets the value for RfRole to be an explicit nil

### UnsetRfRole
`func (o *DcimInterfaceTemplatesCreateRequest) UnsetRfRole()`

UnsetRfRole ensures that no value is present for RfRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


