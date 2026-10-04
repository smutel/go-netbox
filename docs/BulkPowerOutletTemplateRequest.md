# BulkPowerOutletTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullableBulkPowerOutletRequestType**](BulkPowerOutletRequestType.md) |  | [optional] 
**Color** | Pointer to [**BriefModuleBayTypeColor**](BriefModuleBayTypeColor.md) |  | [optional] 
**PowerPort** | Pointer to [**NullableBulkPowerOutletTemplateRequestPowerPort**](BulkPowerOutletTemplateRequestPowerPort.md) |  | [optional] 
**FeedLeg** | Pointer to [**NullableBulkPowerOutletRequestFeedLeg**](BulkPowerOutletRequestFeedLeg.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkPowerOutletTemplateRequest

`func NewBulkPowerOutletTemplateRequest(id int32, name string, ) *BulkPowerOutletTemplateRequest`

NewBulkPowerOutletTemplateRequest instantiates a new BulkPowerOutletTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPowerOutletTemplateRequestWithDefaults

`func NewBulkPowerOutletTemplateRequestWithDefaults() *BulkPowerOutletTemplateRequest`

NewBulkPowerOutletTemplateRequestWithDefaults instantiates a new BulkPowerOutletTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkPowerOutletTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkPowerOutletTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkPowerOutletTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *BulkPowerOutletTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkPowerOutletTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkPowerOutletTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *BulkPowerOutletTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *BulkPowerOutletTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *BulkPowerOutletTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *BulkPowerOutletTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *BulkPowerOutletTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *BulkPowerOutletTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *BulkPowerOutletTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *BulkPowerOutletTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *BulkPowerOutletTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *BulkPowerOutletTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkPowerOutletTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkPowerOutletTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkPowerOutletTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkPowerOutletTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkPowerOutletTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkPowerOutletTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *BulkPowerOutletTemplateRequest) GetType() BulkPowerOutletRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkPowerOutletTemplateRequest) GetTypeOk() (*BulkPowerOutletRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkPowerOutletTemplateRequest) SetType(v BulkPowerOutletRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *BulkPowerOutletTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BulkPowerOutletTemplateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BulkPowerOutletTemplateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetColor

`func (o *BulkPowerOutletTemplateRequest) GetColor() BriefModuleBayTypeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BulkPowerOutletTemplateRequest) GetColorOk() (*BriefModuleBayTypeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BulkPowerOutletTemplateRequest) SetColor(v BriefModuleBayTypeColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *BulkPowerOutletTemplateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPowerPort

`func (o *BulkPowerOutletTemplateRequest) GetPowerPort() BulkPowerOutletTemplateRequestPowerPort`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *BulkPowerOutletTemplateRequest) GetPowerPortOk() (*BulkPowerOutletTemplateRequestPowerPort, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *BulkPowerOutletTemplateRequest) SetPowerPort(v BulkPowerOutletTemplateRequestPowerPort)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *BulkPowerOutletTemplateRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *BulkPowerOutletTemplateRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *BulkPowerOutletTemplateRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetFeedLeg

`func (o *BulkPowerOutletTemplateRequest) GetFeedLeg() BulkPowerOutletRequestFeedLeg`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *BulkPowerOutletTemplateRequest) GetFeedLegOk() (*BulkPowerOutletRequestFeedLeg, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *BulkPowerOutletTemplateRequest) SetFeedLeg(v BulkPowerOutletRequestFeedLeg)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *BulkPowerOutletTemplateRequest) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### SetFeedLegNil

`func (o *BulkPowerOutletTemplateRequest) SetFeedLegNil(b bool)`

 SetFeedLegNil sets the value for FeedLeg to be an explicit nil

### UnsetFeedLeg
`func (o *BulkPowerOutletTemplateRequest) UnsetFeedLeg()`

UnsetFeedLeg ensures that no value is present for FeedLeg, not even an explicit nil
### GetDescription

`func (o *BulkPowerOutletTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkPowerOutletTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkPowerOutletTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkPowerOutletTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


