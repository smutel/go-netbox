# PatchedBulkRearPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | Pointer to **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**BulkFrontPortRequestType**](BulkFrontPortRequestType.md) |  | [optional] 
**Color** | Pointer to [**BriefModuleBayTypeColor**](BriefModuleBayTypeColor.md) |  | [optional] 
**Positions** | Pointer to **int32** |  | [optional] 
**FrontPorts** | Pointer to [**[]RearPortTemplateMappingRequest**](RearPortTemplateMappingRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedBulkRearPortTemplateRequest

`func NewPatchedBulkRearPortTemplateRequest(id int32, ) *PatchedBulkRearPortTemplateRequest`

NewPatchedBulkRearPortTemplateRequest instantiates a new PatchedBulkRearPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkRearPortTemplateRequestWithDefaults

`func NewPatchedBulkRearPortTemplateRequestWithDefaults() *PatchedBulkRearPortTemplateRequest`

NewPatchedBulkRearPortTemplateRequestWithDefaults instantiates a new PatchedBulkRearPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkRearPortTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkRearPortTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkRearPortTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *PatchedBulkRearPortTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedBulkRearPortTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedBulkRearPortTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedBulkRearPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *PatchedBulkRearPortTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *PatchedBulkRearPortTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *PatchedBulkRearPortTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedBulkRearPortTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedBulkRearPortTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedBulkRearPortTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *PatchedBulkRearPortTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *PatchedBulkRearPortTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *PatchedBulkRearPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkRearPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkRearPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkRearPortTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkRearPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkRearPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkRearPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkRearPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkRearPortTemplateRequest) GetType() BulkFrontPortRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkRearPortTemplateRequest) GetTypeOk() (*BulkFrontPortRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkRearPortTemplateRequest) SetType(v BulkFrontPortRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkRearPortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetColor

`func (o *PatchedBulkRearPortTemplateRequest) GetColor() BriefModuleBayTypeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedBulkRearPortTemplateRequest) GetColorOk() (*BriefModuleBayTypeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedBulkRearPortTemplateRequest) SetColor(v BriefModuleBayTypeColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedBulkRearPortTemplateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPositions

`func (o *PatchedBulkRearPortTemplateRequest) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *PatchedBulkRearPortTemplateRequest) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *PatchedBulkRearPortTemplateRequest) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *PatchedBulkRearPortTemplateRequest) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetFrontPorts

`func (o *PatchedBulkRearPortTemplateRequest) GetFrontPorts() []RearPortTemplateMappingRequest`

GetFrontPorts returns the FrontPorts field if non-nil, zero value otherwise.

### GetFrontPortsOk

`func (o *PatchedBulkRearPortTemplateRequest) GetFrontPortsOk() (*[]RearPortTemplateMappingRequest, bool)`

GetFrontPortsOk returns a tuple with the FrontPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPorts

`func (o *PatchedBulkRearPortTemplateRequest) SetFrontPorts(v []RearPortTemplateMappingRequest)`

SetFrontPorts sets FrontPorts field to given value.

### HasFrontPorts

`func (o *PatchedBulkRearPortTemplateRequest) HasFrontPorts() bool`

HasFrontPorts returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkRearPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkRearPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkRearPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkRearPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


