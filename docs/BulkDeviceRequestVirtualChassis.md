# BulkDeviceRequestVirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Master** | Pointer to [**NullableNestedDeviceRequest**](NestedDeviceRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkDeviceRequestVirtualChassis

`func NewBulkDeviceRequestVirtualChassis(name string, ) *BulkDeviceRequestVirtualChassis`

NewBulkDeviceRequestVirtualChassis instantiates a new BulkDeviceRequestVirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeviceRequestVirtualChassisWithDefaults

`func NewBulkDeviceRequestVirtualChassisWithDefaults() *BulkDeviceRequestVirtualChassis`

NewBulkDeviceRequestVirtualChassisWithDefaults instantiates a new BulkDeviceRequestVirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BulkDeviceRequestVirtualChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkDeviceRequestVirtualChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkDeviceRequestVirtualChassis) SetName(v string)`

SetName sets Name field to given value.


### GetMaster

`func (o *BulkDeviceRequestVirtualChassis) GetMaster() NestedDeviceRequest`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *BulkDeviceRequestVirtualChassis) GetMasterOk() (*NestedDeviceRequest, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *BulkDeviceRequestVirtualChassis) SetMaster(v NestedDeviceRequest)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *BulkDeviceRequestVirtualChassis) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### SetMasterNil

`func (o *BulkDeviceRequestVirtualChassis) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *BulkDeviceRequestVirtualChassis) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
### GetDescription

`func (o *BulkDeviceRequestVirtualChassis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkDeviceRequestVirtualChassis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkDeviceRequestVirtualChassis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkDeviceRequestVirtualChassis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


