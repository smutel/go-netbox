# BulkFrontPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | 
**Module** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**BulkFrontPortRequestType**](BulkFrontPortRequestType.md) |  | 
**Color** | Pointer to [**BriefModuleBayTypeColor**](BriefModuleBayTypeColor.md) |  | [optional] 
**Positions** | Pointer to **int32** |  | [optional] 
**RearPorts** | Pointer to [**[]FrontPortMappingRequest**](FrontPortMappingRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkFrontPortRequest

`func NewBulkFrontPortRequest(id int32, device BriefCoolingIntakeRequestDevice, name string, type_ BulkFrontPortRequestType, ) *BulkFrontPortRequest`

NewBulkFrontPortRequest instantiates a new BulkFrontPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkFrontPortRequestWithDefaults

`func NewBulkFrontPortRequestWithDefaults() *BulkFrontPortRequest`

NewBulkFrontPortRequestWithDefaults instantiates a new BulkFrontPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkFrontPortRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkFrontPortRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkFrontPortRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *BulkFrontPortRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkFrontPortRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkFrontPortRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *BulkFrontPortRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *BulkFrontPortRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *BulkFrontPortRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *BulkFrontPortRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *BulkFrontPortRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *BulkFrontPortRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *BulkFrontPortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkFrontPortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkFrontPortRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkFrontPortRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkFrontPortRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkFrontPortRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkFrontPortRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *BulkFrontPortRequest) GetType() BulkFrontPortRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkFrontPortRequest) GetTypeOk() (*BulkFrontPortRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkFrontPortRequest) SetType(v BulkFrontPortRequestType)`

SetType sets Type field to given value.


### GetColor

`func (o *BulkFrontPortRequest) GetColor() BriefModuleBayTypeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BulkFrontPortRequest) GetColorOk() (*BriefModuleBayTypeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BulkFrontPortRequest) SetColor(v BriefModuleBayTypeColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *BulkFrontPortRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPositions

`func (o *BulkFrontPortRequest) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *BulkFrontPortRequest) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *BulkFrontPortRequest) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *BulkFrontPortRequest) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetRearPorts

`func (o *BulkFrontPortRequest) GetRearPorts() []FrontPortMappingRequest`

GetRearPorts returns the RearPorts field if non-nil, zero value otherwise.

### GetRearPortsOk

`func (o *BulkFrontPortRequest) GetRearPortsOk() (*[]FrontPortMappingRequest, bool)`

GetRearPortsOk returns a tuple with the RearPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPorts

`func (o *BulkFrontPortRequest) SetRearPorts(v []FrontPortMappingRequest)`

SetRearPorts sets RearPorts field to given value.

### HasRearPorts

`func (o *BulkFrontPortRequest) HasRearPorts() bool`

HasRearPorts returns a boolean if a field has been set.

### GetDescription

`func (o *BulkFrontPortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkFrontPortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkFrontPortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkFrontPortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *BulkFrontPortRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *BulkFrontPortRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *BulkFrontPortRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *BulkFrontPortRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetOwner

`func (o *BulkFrontPortRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkFrontPortRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkFrontPortRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkFrontPortRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkFrontPortRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkFrontPortRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *BulkFrontPortRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkFrontPortRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkFrontPortRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkFrontPortRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkFrontPortRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkFrontPortRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkFrontPortRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkFrontPortRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


