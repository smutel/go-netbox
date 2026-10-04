# PatchedBulkVMInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**VirtualMachine** | Pointer to [**BulkVMInterfaceRequestVirtualMachine**](BulkVMInterfaceRequestVirtualMachine.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**NullableNestedVMInterfaceRequest**](NestedVMInterfaceRequest.md) |  | [optional] 
**Bridge** | Pointer to [**NullableNestedVMInterfaceRequest**](NestedVMInterfaceRequest.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**PrimaryMacAddress** | Pointer to [**NullableBulkInterfaceRequestPrimaryMacAddress**](BulkInterfaceRequestPrimaryMacAddress.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**BulkInterfaceRequestMode**](BulkInterfaceRequestMode.md) |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableBulkInterfaceRequestUntaggedVlan**](BulkInterfaceRequestUntaggedVlan.md) |  | [optional] 
**TaggedVlans** | Pointer to **[]int32** |  | [optional] 
**QinqSvlan** | Pointer to [**NullableBulkInterfaceRequestUntaggedVlan**](BulkInterfaceRequestUntaggedVlan.md) |  | [optional] 
**VlanTranslationPolicy** | Pointer to [**NullableBulkInterfaceRequestVlanTranslationPolicy**](BulkInterfaceRequestVlanTranslationPolicy.md) |  | [optional] 
**Vrf** | Pointer to [**NullableBulkIPAddressRequestVrf**](BulkIPAddressRequestVrf.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkVMInterfaceRequest

`func NewPatchedBulkVMInterfaceRequest(id int32, ) *PatchedBulkVMInterfaceRequest`

NewPatchedBulkVMInterfaceRequest instantiates a new PatchedBulkVMInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkVMInterfaceRequestWithDefaults

`func NewPatchedBulkVMInterfaceRequestWithDefaults() *PatchedBulkVMInterfaceRequest`

NewPatchedBulkVMInterfaceRequestWithDefaults instantiates a new PatchedBulkVMInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkVMInterfaceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkVMInterfaceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkVMInterfaceRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetVirtualMachine

`func (o *PatchedBulkVMInterfaceRequest) GetVirtualMachine() BulkVMInterfaceRequestVirtualMachine`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *PatchedBulkVMInterfaceRequest) GetVirtualMachineOk() (*BulkVMInterfaceRequestVirtualMachine, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *PatchedBulkVMInterfaceRequest) SetVirtualMachine(v BulkVMInterfaceRequestVirtualMachine)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *PatchedBulkVMInterfaceRequest) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkVMInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkVMInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkVMInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkVMInterfaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkVMInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkVMInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkVMInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkVMInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *PatchedBulkVMInterfaceRequest) GetParent() NestedVMInterfaceRequest`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkVMInterfaceRequest) GetParentOk() (*NestedVMInterfaceRequest, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkVMInterfaceRequest) SetParent(v NestedVMInterfaceRequest)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkVMInterfaceRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedBulkVMInterfaceRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedBulkVMInterfaceRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *PatchedBulkVMInterfaceRequest) GetBridge() NestedVMInterfaceRequest`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *PatchedBulkVMInterfaceRequest) GetBridgeOk() (*NestedVMInterfaceRequest, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *PatchedBulkVMInterfaceRequest) SetBridge(v NestedVMInterfaceRequest)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *PatchedBulkVMInterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *PatchedBulkVMInterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *PatchedBulkVMInterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetMtu

`func (o *PatchedBulkVMInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PatchedBulkVMInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PatchedBulkVMInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PatchedBulkVMInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *PatchedBulkVMInterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *PatchedBulkVMInterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetMacAddress

`func (o *PatchedBulkVMInterfaceRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PatchedBulkVMInterfaceRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PatchedBulkVMInterfaceRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PatchedBulkVMInterfaceRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *PatchedBulkVMInterfaceRequest) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *PatchedBulkVMInterfaceRequest) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetPrimaryMacAddress

`func (o *PatchedBulkVMInterfaceRequest) GetPrimaryMacAddress() BulkInterfaceRequestPrimaryMacAddress`

GetPrimaryMacAddress returns the PrimaryMacAddress field if non-nil, zero value otherwise.

### GetPrimaryMacAddressOk

`func (o *PatchedBulkVMInterfaceRequest) GetPrimaryMacAddressOk() (*BulkInterfaceRequestPrimaryMacAddress, bool)`

GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMacAddress

`func (o *PatchedBulkVMInterfaceRequest) SetPrimaryMacAddress(v BulkInterfaceRequestPrimaryMacAddress)`

SetPrimaryMacAddress sets PrimaryMacAddress field to given value.

### HasPrimaryMacAddress

`func (o *PatchedBulkVMInterfaceRequest) HasPrimaryMacAddress() bool`

HasPrimaryMacAddress returns a boolean if a field has been set.

### SetPrimaryMacAddressNil

`func (o *PatchedBulkVMInterfaceRequest) SetPrimaryMacAddressNil(b bool)`

 SetPrimaryMacAddressNil sets the value for PrimaryMacAddress to be an explicit nil

### UnsetPrimaryMacAddress
`func (o *PatchedBulkVMInterfaceRequest) UnsetPrimaryMacAddress()`

UnsetPrimaryMacAddress ensures that no value is present for PrimaryMacAddress, not even an explicit nil
### GetDescription

`func (o *PatchedBulkVMInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkVMInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkVMInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkVMInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *PatchedBulkVMInterfaceRequest) GetMode() BulkInterfaceRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedBulkVMInterfaceRequest) GetModeOk() (*BulkInterfaceRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedBulkVMInterfaceRequest) SetMode(v BulkInterfaceRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedBulkVMInterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetUntaggedVlan

`func (o *PatchedBulkVMInterfaceRequest) GetUntaggedVlan() BulkInterfaceRequestUntaggedVlan`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *PatchedBulkVMInterfaceRequest) GetUntaggedVlanOk() (*BulkInterfaceRequestUntaggedVlan, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *PatchedBulkVMInterfaceRequest) SetUntaggedVlan(v BulkInterfaceRequestUntaggedVlan)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *PatchedBulkVMInterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *PatchedBulkVMInterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *PatchedBulkVMInterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *PatchedBulkVMInterfaceRequest) GetTaggedVlans() []int32`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *PatchedBulkVMInterfaceRequest) GetTaggedVlansOk() (*[]int32, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *PatchedBulkVMInterfaceRequest) SetTaggedVlans(v []int32)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *PatchedBulkVMInterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetQinqSvlan

`func (o *PatchedBulkVMInterfaceRequest) GetQinqSvlan() BulkInterfaceRequestUntaggedVlan`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *PatchedBulkVMInterfaceRequest) GetQinqSvlanOk() (*BulkInterfaceRequestUntaggedVlan, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *PatchedBulkVMInterfaceRequest) SetQinqSvlan(v BulkInterfaceRequestUntaggedVlan)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *PatchedBulkVMInterfaceRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *PatchedBulkVMInterfaceRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *PatchedBulkVMInterfaceRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetVlanTranslationPolicy

`func (o *PatchedBulkVMInterfaceRequest) GetVlanTranslationPolicy() BulkInterfaceRequestVlanTranslationPolicy`

GetVlanTranslationPolicy returns the VlanTranslationPolicy field if non-nil, zero value otherwise.

### GetVlanTranslationPolicyOk

`func (o *PatchedBulkVMInterfaceRequest) GetVlanTranslationPolicyOk() (*BulkInterfaceRequestVlanTranslationPolicy, bool)`

GetVlanTranslationPolicyOk returns a tuple with the VlanTranslationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTranslationPolicy

`func (o *PatchedBulkVMInterfaceRequest) SetVlanTranslationPolicy(v BulkInterfaceRequestVlanTranslationPolicy)`

SetVlanTranslationPolicy sets VlanTranslationPolicy field to given value.

### HasVlanTranslationPolicy

`func (o *PatchedBulkVMInterfaceRequest) HasVlanTranslationPolicy() bool`

HasVlanTranslationPolicy returns a boolean if a field has been set.

### SetVlanTranslationPolicyNil

`func (o *PatchedBulkVMInterfaceRequest) SetVlanTranslationPolicyNil(b bool)`

 SetVlanTranslationPolicyNil sets the value for VlanTranslationPolicy to be an explicit nil

### UnsetVlanTranslationPolicy
`func (o *PatchedBulkVMInterfaceRequest) UnsetVlanTranslationPolicy()`

UnsetVlanTranslationPolicy ensures that no value is present for VlanTranslationPolicy, not even an explicit nil
### GetVrf

`func (o *PatchedBulkVMInterfaceRequest) GetVrf() BulkIPAddressRequestVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedBulkVMInterfaceRequest) GetVrfOk() (*BulkIPAddressRequestVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedBulkVMInterfaceRequest) SetVrf(v BulkIPAddressRequestVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedBulkVMInterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PatchedBulkVMInterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PatchedBulkVMInterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetOwner

`func (o *PatchedBulkVMInterfaceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkVMInterfaceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkVMInterfaceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkVMInterfaceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkVMInterfaceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkVMInterfaceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkVMInterfaceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkVMInterfaceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkVMInterfaceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkVMInterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkVMInterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkVMInterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkVMInterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkVMInterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


