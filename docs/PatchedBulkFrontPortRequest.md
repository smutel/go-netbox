# PatchedBulkFrontPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | Pointer to [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | [optional] 
**Module** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**BulkFrontPortRequestType**](BulkFrontPortRequestType.md) |  | [optional] 
**Color** | Pointer to [**BriefModuleBayTypeColor**](BriefModuleBayTypeColor.md) |  | [optional] 
**Positions** | Pointer to **int32** |  | [optional] 
**RearPorts** | Pointer to [**[]FrontPortMappingRequest**](FrontPortMappingRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkFrontPortRequest

`func NewPatchedBulkFrontPortRequest(id int32, ) *PatchedBulkFrontPortRequest`

NewPatchedBulkFrontPortRequest instantiates a new PatchedBulkFrontPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkFrontPortRequestWithDefaults

`func NewPatchedBulkFrontPortRequestWithDefaults() *PatchedBulkFrontPortRequest`

NewPatchedBulkFrontPortRequestWithDefaults instantiates a new PatchedBulkFrontPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkFrontPortRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkFrontPortRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkFrontPortRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *PatchedBulkFrontPortRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkFrontPortRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkFrontPortRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkFrontPortRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModule

`func (o *PatchedBulkFrontPortRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PatchedBulkFrontPortRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PatchedBulkFrontPortRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *PatchedBulkFrontPortRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *PatchedBulkFrontPortRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *PatchedBulkFrontPortRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *PatchedBulkFrontPortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkFrontPortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkFrontPortRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkFrontPortRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkFrontPortRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkFrontPortRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkFrontPortRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkFrontPortRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkFrontPortRequest) GetType() BulkFrontPortRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkFrontPortRequest) GetTypeOk() (*BulkFrontPortRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkFrontPortRequest) SetType(v BulkFrontPortRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkFrontPortRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetColor

`func (o *PatchedBulkFrontPortRequest) GetColor() BriefModuleBayTypeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedBulkFrontPortRequest) GetColorOk() (*BriefModuleBayTypeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedBulkFrontPortRequest) SetColor(v BriefModuleBayTypeColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedBulkFrontPortRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPositions

`func (o *PatchedBulkFrontPortRequest) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *PatchedBulkFrontPortRequest) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *PatchedBulkFrontPortRequest) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *PatchedBulkFrontPortRequest) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetRearPorts

`func (o *PatchedBulkFrontPortRequest) GetRearPorts() []FrontPortMappingRequest`

GetRearPorts returns the RearPorts field if non-nil, zero value otherwise.

### GetRearPortsOk

`func (o *PatchedBulkFrontPortRequest) GetRearPortsOk() (*[]FrontPortMappingRequest, bool)`

GetRearPortsOk returns a tuple with the RearPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPorts

`func (o *PatchedBulkFrontPortRequest) SetRearPorts(v []FrontPortMappingRequest)`

SetRearPorts sets RearPorts field to given value.

### HasRearPorts

`func (o *PatchedBulkFrontPortRequest) HasRearPorts() bool`

HasRearPorts returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkFrontPortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkFrontPortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkFrontPortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkFrontPortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PatchedBulkFrontPortRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PatchedBulkFrontPortRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PatchedBulkFrontPortRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PatchedBulkFrontPortRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkFrontPortRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkFrontPortRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkFrontPortRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkFrontPortRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkFrontPortRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkFrontPortRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkFrontPortRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkFrontPortRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkFrontPortRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkFrontPortRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkFrontPortRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkFrontPortRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkFrontPortRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkFrontPortRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


