# BulkPowerOutletRequestPowerPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkPowerOutletRequestPowerPort

`func NewBulkPowerOutletRequestPowerPort(device BriefCoolingIntakeRequestDevice, name string, ) *BulkPowerOutletRequestPowerPort`

NewBulkPowerOutletRequestPowerPort instantiates a new BulkPowerOutletRequestPowerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPowerOutletRequestPowerPortWithDefaults

`func NewBulkPowerOutletRequestPowerPortWithDefaults() *BulkPowerOutletRequestPowerPort`

NewBulkPowerOutletRequestPowerPortWithDefaults instantiates a new BulkPowerOutletRequestPowerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *BulkPowerOutletRequestPowerPort) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkPowerOutletRequestPowerPort) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkPowerOutletRequestPowerPort) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetName

`func (o *BulkPowerOutletRequestPowerPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkPowerOutletRequestPowerPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkPowerOutletRequestPowerPort) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkPowerOutletRequestPowerPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkPowerOutletRequestPowerPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkPowerOutletRequestPowerPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkPowerOutletRequestPowerPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


