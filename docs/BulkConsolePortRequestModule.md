# BulkConsolePortRequestModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | 
**ModuleBay** | [**NestedModuleBayRequest**](NestedModuleBayRequest.md) |  | 

## Methods

### NewBulkConsolePortRequestModule

`func NewBulkConsolePortRequestModule(device BriefCoolingIntakeRequestDevice, moduleBay NestedModuleBayRequest, ) *BulkConsolePortRequestModule`

NewBulkConsolePortRequestModule instantiates a new BulkConsolePortRequestModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkConsolePortRequestModuleWithDefaults

`func NewBulkConsolePortRequestModuleWithDefaults() *BulkConsolePortRequestModule`

NewBulkConsolePortRequestModuleWithDefaults instantiates a new BulkConsolePortRequestModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *BulkConsolePortRequestModule) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkConsolePortRequestModule) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkConsolePortRequestModule) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetModuleBay

`func (o *BulkConsolePortRequestModule) GetModuleBay() NestedModuleBayRequest`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *BulkConsolePortRequestModule) GetModuleBayOk() (*NestedModuleBayRequest, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *BulkConsolePortRequestModule) SetModuleBay(v NestedModuleBayRequest)`

SetModuleBay sets ModuleBay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


