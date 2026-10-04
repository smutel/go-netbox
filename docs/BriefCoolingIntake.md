# BriefCoolingIntake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**BriefDevice**](BriefDevice.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefCoolingIntake

`func NewBriefCoolingIntake(id int32, url string, display string, device BriefDevice, name string, ) *BriefCoolingIntake`

NewBriefCoolingIntake instantiates a new BriefCoolingIntake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefCoolingIntakeWithDefaults

`func NewBriefCoolingIntakeWithDefaults() *BriefCoolingIntake`

NewBriefCoolingIntakeWithDefaults instantiates a new BriefCoolingIntake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefCoolingIntake) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefCoolingIntake) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefCoolingIntake) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefCoolingIntake) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefCoolingIntake) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefCoolingIntake) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefCoolingIntake) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefCoolingIntake) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefCoolingIntake) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *BriefCoolingIntake) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BriefCoolingIntake) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BriefCoolingIntake) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetName

`func (o *BriefCoolingIntake) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefCoolingIntake) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefCoolingIntake) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BriefCoolingIntake) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefCoolingIntake) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefCoolingIntake) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefCoolingIntake) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


