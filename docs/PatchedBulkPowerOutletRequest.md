# PatchedBulkPowerOutletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | Pointer to [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | [optional] 
**Module** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullableBulkPowerOutletRequestType**](BulkPowerOutletRequestType.md) |  | [optional] 
**Status** | Pointer to [**BulkPowerOutletRequestStatus**](BulkPowerOutletRequestStatus.md) |  | [optional] 
**Color** | Pointer to [**BriefModuleBayTypeColor**](BriefModuleBayTypeColor.md) |  | [optional] 
**PowerPort** | Pointer to [**NullableBulkPowerOutletRequestPowerPort**](BulkPowerOutletRequestPowerPort.md) |  | [optional] 
**FeedLeg** | Pointer to [**NullableBulkPowerOutletRequestFeedLeg**](BulkPowerOutletRequestFeedLeg.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkPowerOutletRequest

`func NewPatchedBulkPowerOutletRequest(id int32, ) *PatchedBulkPowerOutletRequest`

NewPatchedBulkPowerOutletRequest instantiates a new PatchedBulkPowerOutletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkPowerOutletRequestWithDefaults

`func NewPatchedBulkPowerOutletRequestWithDefaults() *PatchedBulkPowerOutletRequest`

NewPatchedBulkPowerOutletRequestWithDefaults instantiates a new PatchedBulkPowerOutletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkPowerOutletRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkPowerOutletRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkPowerOutletRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *PatchedBulkPowerOutletRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkPowerOutletRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkPowerOutletRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkPowerOutletRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModule

`func (o *PatchedBulkPowerOutletRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PatchedBulkPowerOutletRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PatchedBulkPowerOutletRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *PatchedBulkPowerOutletRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *PatchedBulkPowerOutletRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *PatchedBulkPowerOutletRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *PatchedBulkPowerOutletRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkPowerOutletRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkPowerOutletRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkPowerOutletRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkPowerOutletRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkPowerOutletRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkPowerOutletRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkPowerOutletRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkPowerOutletRequest) GetType() BulkPowerOutletRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkPowerOutletRequest) GetTypeOk() (*BulkPowerOutletRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkPowerOutletRequest) SetType(v BulkPowerOutletRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkPowerOutletRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchedBulkPowerOutletRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchedBulkPowerOutletRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStatus

`func (o *PatchedBulkPowerOutletRequest) GetStatus() BulkPowerOutletRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkPowerOutletRequest) GetStatusOk() (*BulkPowerOutletRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkPowerOutletRequest) SetStatus(v BulkPowerOutletRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkPowerOutletRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetColor

`func (o *PatchedBulkPowerOutletRequest) GetColor() BriefModuleBayTypeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedBulkPowerOutletRequest) GetColorOk() (*BriefModuleBayTypeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedBulkPowerOutletRequest) SetColor(v BriefModuleBayTypeColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedBulkPowerOutletRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPowerPort

`func (o *PatchedBulkPowerOutletRequest) GetPowerPort() BulkPowerOutletRequestPowerPort`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *PatchedBulkPowerOutletRequest) GetPowerPortOk() (*BulkPowerOutletRequestPowerPort, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *PatchedBulkPowerOutletRequest) SetPowerPort(v BulkPowerOutletRequestPowerPort)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *PatchedBulkPowerOutletRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *PatchedBulkPowerOutletRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *PatchedBulkPowerOutletRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetFeedLeg

`func (o *PatchedBulkPowerOutletRequest) GetFeedLeg() BulkPowerOutletRequestFeedLeg`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *PatchedBulkPowerOutletRequest) GetFeedLegOk() (*BulkPowerOutletRequestFeedLeg, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *PatchedBulkPowerOutletRequest) SetFeedLeg(v BulkPowerOutletRequestFeedLeg)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *PatchedBulkPowerOutletRequest) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### SetFeedLegNil

`func (o *PatchedBulkPowerOutletRequest) SetFeedLegNil(b bool)`

 SetFeedLegNil sets the value for FeedLeg to be an explicit nil

### UnsetFeedLeg
`func (o *PatchedBulkPowerOutletRequest) UnsetFeedLeg()`

UnsetFeedLeg ensures that no value is present for FeedLeg, not even an explicit nil
### GetDescription

`func (o *PatchedBulkPowerOutletRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkPowerOutletRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkPowerOutletRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkPowerOutletRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PatchedBulkPowerOutletRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PatchedBulkPowerOutletRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PatchedBulkPowerOutletRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PatchedBulkPowerOutletRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkPowerOutletRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkPowerOutletRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkPowerOutletRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkPowerOutletRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkPowerOutletRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkPowerOutletRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkPowerOutletRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkPowerOutletRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkPowerOutletRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkPowerOutletRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkPowerOutletRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkPowerOutletRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkPowerOutletRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkPowerOutletRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


