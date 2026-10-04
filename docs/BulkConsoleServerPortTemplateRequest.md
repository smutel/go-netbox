# BulkConsoleServerPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**BulkConsolePortRequestType**](BulkConsolePortRequestType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkConsoleServerPortTemplateRequest

`func NewBulkConsoleServerPortTemplateRequest(id int32, name string, ) *BulkConsoleServerPortTemplateRequest`

NewBulkConsoleServerPortTemplateRequest instantiates a new BulkConsoleServerPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkConsoleServerPortTemplateRequestWithDefaults

`func NewBulkConsoleServerPortTemplateRequestWithDefaults() *BulkConsoleServerPortTemplateRequest`

NewBulkConsoleServerPortTemplateRequestWithDefaults instantiates a new BulkConsoleServerPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkConsoleServerPortTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkConsoleServerPortTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkConsoleServerPortTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *BulkConsoleServerPortTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkConsoleServerPortTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkConsoleServerPortTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *BulkConsoleServerPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *BulkConsoleServerPortTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *BulkConsoleServerPortTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *BulkConsoleServerPortTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *BulkConsoleServerPortTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *BulkConsoleServerPortTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *BulkConsoleServerPortTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *BulkConsoleServerPortTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *BulkConsoleServerPortTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *BulkConsoleServerPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkConsoleServerPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkConsoleServerPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkConsoleServerPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkConsoleServerPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkConsoleServerPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkConsoleServerPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *BulkConsoleServerPortTemplateRequest) GetType() BulkConsolePortRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkConsoleServerPortTemplateRequest) GetTypeOk() (*BulkConsolePortRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkConsoleServerPortTemplateRequest) SetType(v BulkConsolePortRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *BulkConsoleServerPortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *BulkConsoleServerPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkConsoleServerPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkConsoleServerPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkConsoleServerPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


