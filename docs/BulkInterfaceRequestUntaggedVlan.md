# BulkInterfaceRequestUntaggedVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vid** | **int32** | Numeric VLAN ID (1-4094) | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkInterfaceRequestUntaggedVlan

`func NewBulkInterfaceRequestUntaggedVlan(vid int32, name string, ) *BulkInterfaceRequestUntaggedVlan`

NewBulkInterfaceRequestUntaggedVlan instantiates a new BulkInterfaceRequestUntaggedVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInterfaceRequestUntaggedVlanWithDefaults

`func NewBulkInterfaceRequestUntaggedVlanWithDefaults() *BulkInterfaceRequestUntaggedVlan`

NewBulkInterfaceRequestUntaggedVlanWithDefaults instantiates a new BulkInterfaceRequestUntaggedVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVid

`func (o *BulkInterfaceRequestUntaggedVlan) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *BulkInterfaceRequestUntaggedVlan) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *BulkInterfaceRequestUntaggedVlan) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *BulkInterfaceRequestUntaggedVlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkInterfaceRequestUntaggedVlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkInterfaceRequestUntaggedVlan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkInterfaceRequestUntaggedVlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkInterfaceRequestUntaggedVlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkInterfaceRequestUntaggedVlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkInterfaceRequestUntaggedVlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


