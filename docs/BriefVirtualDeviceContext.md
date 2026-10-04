# BriefVirtualDeviceContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Device** | [**BriefDevice**](BriefDevice.md) |  | 
**Identifier** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefVirtualDeviceContext

`func NewBriefVirtualDeviceContext(id int32, url string, display string, name string, device BriefDevice, ) *BriefVirtualDeviceContext`

NewBriefVirtualDeviceContext instantiates a new BriefVirtualDeviceContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefVirtualDeviceContextWithDefaults

`func NewBriefVirtualDeviceContextWithDefaults() *BriefVirtualDeviceContext`

NewBriefVirtualDeviceContextWithDefaults instantiates a new BriefVirtualDeviceContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefVirtualDeviceContext) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefVirtualDeviceContext) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefVirtualDeviceContext) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefVirtualDeviceContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefVirtualDeviceContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefVirtualDeviceContext) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefVirtualDeviceContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefVirtualDeviceContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefVirtualDeviceContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefVirtualDeviceContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefVirtualDeviceContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefVirtualDeviceContext) SetName(v string)`

SetName sets Name field to given value.


### GetDevice

`func (o *BriefVirtualDeviceContext) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BriefVirtualDeviceContext) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BriefVirtualDeviceContext) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetIdentifier

`func (o *BriefVirtualDeviceContext) GetIdentifier() int32`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BriefVirtualDeviceContext) GetIdentifierOk() (*int32, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BriefVirtualDeviceContext) SetIdentifier(v int32)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BriefVirtualDeviceContext) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *BriefVirtualDeviceContext) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *BriefVirtualDeviceContext) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetDescription

`func (o *BriefVirtualDeviceContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefVirtualDeviceContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefVirtualDeviceContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefVirtualDeviceContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


