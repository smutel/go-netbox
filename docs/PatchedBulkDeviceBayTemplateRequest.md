# PatchedBulkDeviceBayTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | Pointer to [**BulkDeviceBayTemplateRequestDeviceType**](BulkDeviceBayTemplateRequestDeviceType.md) |  | [optional] 
**Name** | Pointer to **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedBulkDeviceBayTemplateRequest

`func NewPatchedBulkDeviceBayTemplateRequest(id int32, ) *PatchedBulkDeviceBayTemplateRequest`

NewPatchedBulkDeviceBayTemplateRequest instantiates a new PatchedBulkDeviceBayTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkDeviceBayTemplateRequestWithDefaults

`func NewPatchedBulkDeviceBayTemplateRequestWithDefaults() *PatchedBulkDeviceBayTemplateRequest`

NewPatchedBulkDeviceBayTemplateRequestWithDefaults instantiates a new PatchedBulkDeviceBayTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkDeviceBayTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkDeviceBayTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkDeviceBayTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *PatchedBulkDeviceBayTemplateRequest) GetDeviceType() BulkDeviceBayTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedBulkDeviceBayTemplateRequest) GetDeviceTypeOk() (*BulkDeviceBayTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedBulkDeviceBayTemplateRequest) SetDeviceType(v BulkDeviceBayTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedBulkDeviceBayTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkDeviceBayTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkDeviceBayTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkDeviceBayTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkDeviceBayTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkDeviceBayTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkDeviceBayTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkDeviceBayTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkDeviceBayTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkDeviceBayTemplateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkDeviceBayTemplateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkDeviceBayTemplateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkDeviceBayTemplateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkDeviceBayTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkDeviceBayTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkDeviceBayTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkDeviceBayTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


