# PatchedBulkInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | Pointer to [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | [optional] 
**Vdcs** | Pointer to **[]int32** |  | [optional] 
**Module** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**BulkInterfaceRequestType**](BulkInterfaceRequestType.md) |  | [optional] 
**Channels** | Pointer to **NullableInt32** | The number of channels into which this interface is channelized | [optional] 
**ChannelId** | Pointer to **NullableInt32** | The channel on the parent interface to which this subinterface is bound | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**NullableNestedInterfaceRequest**](NestedInterfaceRequest.md) |  | [optional] 
**Bridge** | Pointer to [**NullableNestedInterfaceRequest**](NestedInterfaceRequest.md) |  | [optional] 
**Lag** | Pointer to [**NullableNestedInterfaceRequest**](NestedInterfaceRequest.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**PrimaryMacAddress** | Pointer to [**NullableBulkInterfaceRequestPrimaryMacAddress**](BulkInterfaceRequestPrimaryMacAddress.md) |  | [optional] 
**Speed** | Pointer to **NullableInt64** |  | [optional] 
**Duplex** | Pointer to [**NullableBulkInterfaceRequestDuplex**](BulkInterfaceRequestDuplex.md) |  | [optional] 
**Wwn** | Pointer to **NullableString** |  | [optional] 
**MgmtOnly** | Pointer to **bool** | This interface is used only for out-of-band management | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**BulkInterfaceRequestMode**](BulkInterfaceRequestMode.md) |  | [optional] 
**RfRole** | Pointer to [**BulkInterfaceRequestRfRole**](BulkInterfaceRequestRfRole.md) |  | [optional] 
**RfChannel** | Pointer to [**BulkInterfaceRequestRfChannel**](BulkInterfaceRequestRfChannel.md) |  | [optional] 
**PoeMode** | Pointer to [**BulkInterfaceRequestPoeMode**](BulkInterfaceRequestPoeMode.md) |  | [optional] 
**PoeType** | Pointer to [**BulkInterfaceRequestPoeType**](BulkInterfaceRequestPoeType.md) |  | [optional] 
**RfChannelFrequency** | Pointer to **NullableFloat64** | Populated by selected channel (if set) | [optional] 
**RfChannelWidth** | Pointer to **NullableFloat64** | Populated by selected channel (if set) | [optional] 
**TxPower** | Pointer to **NullableInt32** |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableBulkInterfaceRequestUntaggedVlan**](BulkInterfaceRequestUntaggedVlan.md) |  | [optional] 
**TaggedVlans** | Pointer to **[]int32** |  | [optional] 
**QinqSvlan** | Pointer to [**NullableBulkInterfaceRequestUntaggedVlan**](BulkInterfaceRequestUntaggedVlan.md) |  | [optional] 
**VlanTranslationPolicy** | Pointer to [**NullableBulkInterfaceRequestVlanTranslationPolicy**](BulkInterfaceRequestVlanTranslationPolicy.md) |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**WirelessLans** | Pointer to **[]int32** |  | [optional] 
**Vrf** | Pointer to [**NullableBulkIPAddressRequestVrf**](BulkIPAddressRequestVrf.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkInterfaceRequest

`func NewPatchedBulkInterfaceRequest(id int32, ) *PatchedBulkInterfaceRequest`

NewPatchedBulkInterfaceRequest instantiates a new PatchedBulkInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkInterfaceRequestWithDefaults

`func NewPatchedBulkInterfaceRequestWithDefaults() *PatchedBulkInterfaceRequest`

NewPatchedBulkInterfaceRequestWithDefaults instantiates a new PatchedBulkInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkInterfaceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkInterfaceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkInterfaceRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *PatchedBulkInterfaceRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkInterfaceRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkInterfaceRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkInterfaceRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetVdcs

`func (o *PatchedBulkInterfaceRequest) GetVdcs() []int32`

GetVdcs returns the Vdcs field if non-nil, zero value otherwise.

### GetVdcsOk

`func (o *PatchedBulkInterfaceRequest) GetVdcsOk() (*[]int32, bool)`

GetVdcsOk returns a tuple with the Vdcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcs

`func (o *PatchedBulkInterfaceRequest) SetVdcs(v []int32)`

SetVdcs sets Vdcs field to given value.

### HasVdcs

`func (o *PatchedBulkInterfaceRequest) HasVdcs() bool`

HasVdcs returns a boolean if a field has been set.

### GetModule

`func (o *PatchedBulkInterfaceRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PatchedBulkInterfaceRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PatchedBulkInterfaceRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *PatchedBulkInterfaceRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *PatchedBulkInterfaceRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *PatchedBulkInterfaceRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *PatchedBulkInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkInterfaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkInterfaceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkInterfaceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkInterfaceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkInterfaceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkInterfaceRequest) GetType() BulkInterfaceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkInterfaceRequest) GetTypeOk() (*BulkInterfaceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkInterfaceRequest) SetType(v BulkInterfaceRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkInterfaceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChannels

`func (o *PatchedBulkInterfaceRequest) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *PatchedBulkInterfaceRequest) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *PatchedBulkInterfaceRequest) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *PatchedBulkInterfaceRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *PatchedBulkInterfaceRequest) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *PatchedBulkInterfaceRequest) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetChannelId

`func (o *PatchedBulkInterfaceRequest) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PatchedBulkInterfaceRequest) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PatchedBulkInterfaceRequest) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *PatchedBulkInterfaceRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *PatchedBulkInterfaceRequest) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *PatchedBulkInterfaceRequest) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetEnabled

`func (o *PatchedBulkInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *PatchedBulkInterfaceRequest) GetParent() NestedInterfaceRequest`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkInterfaceRequest) GetParentOk() (*NestedInterfaceRequest, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkInterfaceRequest) SetParent(v NestedInterfaceRequest)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkInterfaceRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedBulkInterfaceRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedBulkInterfaceRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *PatchedBulkInterfaceRequest) GetBridge() NestedInterfaceRequest`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *PatchedBulkInterfaceRequest) GetBridgeOk() (*NestedInterfaceRequest, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *PatchedBulkInterfaceRequest) SetBridge(v NestedInterfaceRequest)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *PatchedBulkInterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *PatchedBulkInterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *PatchedBulkInterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetLag

`func (o *PatchedBulkInterfaceRequest) GetLag() NestedInterfaceRequest`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *PatchedBulkInterfaceRequest) GetLagOk() (*NestedInterfaceRequest, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *PatchedBulkInterfaceRequest) SetLag(v NestedInterfaceRequest)`

SetLag sets Lag field to given value.

### HasLag

`func (o *PatchedBulkInterfaceRequest) HasLag() bool`

HasLag returns a boolean if a field has been set.

### SetLagNil

`func (o *PatchedBulkInterfaceRequest) SetLagNil(b bool)`

 SetLagNil sets the value for Lag to be an explicit nil

### UnsetLag
`func (o *PatchedBulkInterfaceRequest) UnsetLag()`

UnsetLag ensures that no value is present for Lag, not even an explicit nil
### GetMtu

`func (o *PatchedBulkInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PatchedBulkInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PatchedBulkInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PatchedBulkInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *PatchedBulkInterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *PatchedBulkInterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetMacAddress

`func (o *PatchedBulkInterfaceRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PatchedBulkInterfaceRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PatchedBulkInterfaceRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PatchedBulkInterfaceRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *PatchedBulkInterfaceRequest) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *PatchedBulkInterfaceRequest) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetPrimaryMacAddress

`func (o *PatchedBulkInterfaceRequest) GetPrimaryMacAddress() BulkInterfaceRequestPrimaryMacAddress`

GetPrimaryMacAddress returns the PrimaryMacAddress field if non-nil, zero value otherwise.

### GetPrimaryMacAddressOk

`func (o *PatchedBulkInterfaceRequest) GetPrimaryMacAddressOk() (*BulkInterfaceRequestPrimaryMacAddress, bool)`

GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMacAddress

`func (o *PatchedBulkInterfaceRequest) SetPrimaryMacAddress(v BulkInterfaceRequestPrimaryMacAddress)`

SetPrimaryMacAddress sets PrimaryMacAddress field to given value.

### HasPrimaryMacAddress

`func (o *PatchedBulkInterfaceRequest) HasPrimaryMacAddress() bool`

HasPrimaryMacAddress returns a boolean if a field has been set.

### SetPrimaryMacAddressNil

`func (o *PatchedBulkInterfaceRequest) SetPrimaryMacAddressNil(b bool)`

 SetPrimaryMacAddressNil sets the value for PrimaryMacAddress to be an explicit nil

### UnsetPrimaryMacAddress
`func (o *PatchedBulkInterfaceRequest) UnsetPrimaryMacAddress()`

UnsetPrimaryMacAddress ensures that no value is present for PrimaryMacAddress, not even an explicit nil
### GetSpeed

`func (o *PatchedBulkInterfaceRequest) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *PatchedBulkInterfaceRequest) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *PatchedBulkInterfaceRequest) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *PatchedBulkInterfaceRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *PatchedBulkInterfaceRequest) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *PatchedBulkInterfaceRequest) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetDuplex

`func (o *PatchedBulkInterfaceRequest) GetDuplex() BulkInterfaceRequestDuplex`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *PatchedBulkInterfaceRequest) GetDuplexOk() (*BulkInterfaceRequestDuplex, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *PatchedBulkInterfaceRequest) SetDuplex(v BulkInterfaceRequestDuplex)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *PatchedBulkInterfaceRequest) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### SetDuplexNil

`func (o *PatchedBulkInterfaceRequest) SetDuplexNil(b bool)`

 SetDuplexNil sets the value for Duplex to be an explicit nil

### UnsetDuplex
`func (o *PatchedBulkInterfaceRequest) UnsetDuplex()`

UnsetDuplex ensures that no value is present for Duplex, not even an explicit nil
### GetWwn

`func (o *PatchedBulkInterfaceRequest) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *PatchedBulkInterfaceRequest) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *PatchedBulkInterfaceRequest) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *PatchedBulkInterfaceRequest) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### SetWwnNil

`func (o *PatchedBulkInterfaceRequest) SetWwnNil(b bool)`

 SetWwnNil sets the value for Wwn to be an explicit nil

### UnsetWwn
`func (o *PatchedBulkInterfaceRequest) UnsetWwn()`

UnsetWwn ensures that no value is present for Wwn, not even an explicit nil
### GetMgmtOnly

`func (o *PatchedBulkInterfaceRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *PatchedBulkInterfaceRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *PatchedBulkInterfaceRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *PatchedBulkInterfaceRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *PatchedBulkInterfaceRequest) GetMode() BulkInterfaceRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedBulkInterfaceRequest) GetModeOk() (*BulkInterfaceRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedBulkInterfaceRequest) SetMode(v BulkInterfaceRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedBulkInterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRfRole

`func (o *PatchedBulkInterfaceRequest) GetRfRole() BulkInterfaceRequestRfRole`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *PatchedBulkInterfaceRequest) GetRfRoleOk() (*BulkInterfaceRequestRfRole, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *PatchedBulkInterfaceRequest) SetRfRole(v BulkInterfaceRequestRfRole)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *PatchedBulkInterfaceRequest) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### GetRfChannel

`func (o *PatchedBulkInterfaceRequest) GetRfChannel() BulkInterfaceRequestRfChannel`

GetRfChannel returns the RfChannel field if non-nil, zero value otherwise.

### GetRfChannelOk

`func (o *PatchedBulkInterfaceRequest) GetRfChannelOk() (*BulkInterfaceRequestRfChannel, bool)`

GetRfChannelOk returns a tuple with the RfChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannel

`func (o *PatchedBulkInterfaceRequest) SetRfChannel(v BulkInterfaceRequestRfChannel)`

SetRfChannel sets RfChannel field to given value.

### HasRfChannel

`func (o *PatchedBulkInterfaceRequest) HasRfChannel() bool`

HasRfChannel returns a boolean if a field has been set.

### GetPoeMode

`func (o *PatchedBulkInterfaceRequest) GetPoeMode() BulkInterfaceRequestPoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *PatchedBulkInterfaceRequest) GetPoeModeOk() (*BulkInterfaceRequestPoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *PatchedBulkInterfaceRequest) SetPoeMode(v BulkInterfaceRequestPoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *PatchedBulkInterfaceRequest) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### GetPoeType

`func (o *PatchedBulkInterfaceRequest) GetPoeType() BulkInterfaceRequestPoeType`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *PatchedBulkInterfaceRequest) GetPoeTypeOk() (*BulkInterfaceRequestPoeType, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *PatchedBulkInterfaceRequest) SetPoeType(v BulkInterfaceRequestPoeType)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *PatchedBulkInterfaceRequest) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### GetRfChannelFrequency

`func (o *PatchedBulkInterfaceRequest) GetRfChannelFrequency() float64`

GetRfChannelFrequency returns the RfChannelFrequency field if non-nil, zero value otherwise.

### GetRfChannelFrequencyOk

`func (o *PatchedBulkInterfaceRequest) GetRfChannelFrequencyOk() (*float64, bool)`

GetRfChannelFrequencyOk returns a tuple with the RfChannelFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannelFrequency

`func (o *PatchedBulkInterfaceRequest) SetRfChannelFrequency(v float64)`

SetRfChannelFrequency sets RfChannelFrequency field to given value.

### HasRfChannelFrequency

`func (o *PatchedBulkInterfaceRequest) HasRfChannelFrequency() bool`

HasRfChannelFrequency returns a boolean if a field has been set.

### SetRfChannelFrequencyNil

`func (o *PatchedBulkInterfaceRequest) SetRfChannelFrequencyNil(b bool)`

 SetRfChannelFrequencyNil sets the value for RfChannelFrequency to be an explicit nil

### UnsetRfChannelFrequency
`func (o *PatchedBulkInterfaceRequest) UnsetRfChannelFrequency()`

UnsetRfChannelFrequency ensures that no value is present for RfChannelFrequency, not even an explicit nil
### GetRfChannelWidth

`func (o *PatchedBulkInterfaceRequest) GetRfChannelWidth() float64`

GetRfChannelWidth returns the RfChannelWidth field if non-nil, zero value otherwise.

### GetRfChannelWidthOk

`func (o *PatchedBulkInterfaceRequest) GetRfChannelWidthOk() (*float64, bool)`

GetRfChannelWidthOk returns a tuple with the RfChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannelWidth

`func (o *PatchedBulkInterfaceRequest) SetRfChannelWidth(v float64)`

SetRfChannelWidth sets RfChannelWidth field to given value.

### HasRfChannelWidth

`func (o *PatchedBulkInterfaceRequest) HasRfChannelWidth() bool`

HasRfChannelWidth returns a boolean if a field has been set.

### SetRfChannelWidthNil

`func (o *PatchedBulkInterfaceRequest) SetRfChannelWidthNil(b bool)`

 SetRfChannelWidthNil sets the value for RfChannelWidth to be an explicit nil

### UnsetRfChannelWidth
`func (o *PatchedBulkInterfaceRequest) UnsetRfChannelWidth()`

UnsetRfChannelWidth ensures that no value is present for RfChannelWidth, not even an explicit nil
### GetTxPower

`func (o *PatchedBulkInterfaceRequest) GetTxPower() int32`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *PatchedBulkInterfaceRequest) GetTxPowerOk() (*int32, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *PatchedBulkInterfaceRequest) SetTxPower(v int32)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *PatchedBulkInterfaceRequest) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### SetTxPowerNil

`func (o *PatchedBulkInterfaceRequest) SetTxPowerNil(b bool)`

 SetTxPowerNil sets the value for TxPower to be an explicit nil

### UnsetTxPower
`func (o *PatchedBulkInterfaceRequest) UnsetTxPower()`

UnsetTxPower ensures that no value is present for TxPower, not even an explicit nil
### GetUntaggedVlan

`func (o *PatchedBulkInterfaceRequest) GetUntaggedVlan() BulkInterfaceRequestUntaggedVlan`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *PatchedBulkInterfaceRequest) GetUntaggedVlanOk() (*BulkInterfaceRequestUntaggedVlan, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *PatchedBulkInterfaceRequest) SetUntaggedVlan(v BulkInterfaceRequestUntaggedVlan)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *PatchedBulkInterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *PatchedBulkInterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *PatchedBulkInterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *PatchedBulkInterfaceRequest) GetTaggedVlans() []int32`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *PatchedBulkInterfaceRequest) GetTaggedVlansOk() (*[]int32, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *PatchedBulkInterfaceRequest) SetTaggedVlans(v []int32)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *PatchedBulkInterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetQinqSvlan

`func (o *PatchedBulkInterfaceRequest) GetQinqSvlan() BulkInterfaceRequestUntaggedVlan`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *PatchedBulkInterfaceRequest) GetQinqSvlanOk() (*BulkInterfaceRequestUntaggedVlan, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *PatchedBulkInterfaceRequest) SetQinqSvlan(v BulkInterfaceRequestUntaggedVlan)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *PatchedBulkInterfaceRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *PatchedBulkInterfaceRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *PatchedBulkInterfaceRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetVlanTranslationPolicy

`func (o *PatchedBulkInterfaceRequest) GetVlanTranslationPolicy() BulkInterfaceRequestVlanTranslationPolicy`

GetVlanTranslationPolicy returns the VlanTranslationPolicy field if non-nil, zero value otherwise.

### GetVlanTranslationPolicyOk

`func (o *PatchedBulkInterfaceRequest) GetVlanTranslationPolicyOk() (*BulkInterfaceRequestVlanTranslationPolicy, bool)`

GetVlanTranslationPolicyOk returns a tuple with the VlanTranslationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTranslationPolicy

`func (o *PatchedBulkInterfaceRequest) SetVlanTranslationPolicy(v BulkInterfaceRequestVlanTranslationPolicy)`

SetVlanTranslationPolicy sets VlanTranslationPolicy field to given value.

### HasVlanTranslationPolicy

`func (o *PatchedBulkInterfaceRequest) HasVlanTranslationPolicy() bool`

HasVlanTranslationPolicy returns a boolean if a field has been set.

### SetVlanTranslationPolicyNil

`func (o *PatchedBulkInterfaceRequest) SetVlanTranslationPolicyNil(b bool)`

 SetVlanTranslationPolicyNil sets the value for VlanTranslationPolicy to be an explicit nil

### UnsetVlanTranslationPolicy
`func (o *PatchedBulkInterfaceRequest) UnsetVlanTranslationPolicy()`

UnsetVlanTranslationPolicy ensures that no value is present for VlanTranslationPolicy, not even an explicit nil
### GetMarkConnected

`func (o *PatchedBulkInterfaceRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PatchedBulkInterfaceRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PatchedBulkInterfaceRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PatchedBulkInterfaceRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetWirelessLans

`func (o *PatchedBulkInterfaceRequest) GetWirelessLans() []int32`

GetWirelessLans returns the WirelessLans field if non-nil, zero value otherwise.

### GetWirelessLansOk

`func (o *PatchedBulkInterfaceRequest) GetWirelessLansOk() (*[]int32, bool)`

GetWirelessLansOk returns a tuple with the WirelessLans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessLans

`func (o *PatchedBulkInterfaceRequest) SetWirelessLans(v []int32)`

SetWirelessLans sets WirelessLans field to given value.

### HasWirelessLans

`func (o *PatchedBulkInterfaceRequest) HasWirelessLans() bool`

HasWirelessLans returns a boolean if a field has been set.

### GetVrf

`func (o *PatchedBulkInterfaceRequest) GetVrf() BulkIPAddressRequestVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedBulkInterfaceRequest) GetVrfOk() (*BulkIPAddressRequestVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedBulkInterfaceRequest) SetVrf(v BulkIPAddressRequestVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedBulkInterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PatchedBulkInterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PatchedBulkInterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetOwner

`func (o *PatchedBulkInterfaceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkInterfaceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkInterfaceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkInterfaceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkInterfaceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkInterfaceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkInterfaceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkInterfaceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkInterfaceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkInterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkInterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkInterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkInterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkInterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


