# BulkInventoryItemTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeviceType** | [**BulkDeviceBayTemplateRequestDeviceType**](BulkDeviceBayTemplateRequestDeviceType.md) |  | 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Role** | Pointer to [**NullableBulkInventoryItemRequestRole**](BulkInventoryItemRequestRole.md) |  | [optional] 
**Manufacturer** | Pointer to [**NullableBulkInventoryItemRequestManufacturer**](BulkInventoryItemRequestManufacturer.md) |  | [optional] 
**PartId** | Pointer to **string** | Manufacturer-assigned part identifier | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ComponentType** | Pointer to **NullableString** |  | [optional] 
**ComponentId** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewBulkInventoryItemTemplateRequest

`func NewBulkInventoryItemTemplateRequest(id int32, deviceType BulkDeviceBayTemplateRequestDeviceType, name string, ) *BulkInventoryItemTemplateRequest`

NewBulkInventoryItemTemplateRequest instantiates a new BulkInventoryItemTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInventoryItemTemplateRequestWithDefaults

`func NewBulkInventoryItemTemplateRequestWithDefaults() *BulkInventoryItemTemplateRequest`

NewBulkInventoryItemTemplateRequestWithDefaults instantiates a new BulkInventoryItemTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkInventoryItemTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkInventoryItemTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkInventoryItemTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *BulkInventoryItemTemplateRequest) GetDeviceType() BulkDeviceBayTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkInventoryItemTemplateRequest) GetDeviceTypeOk() (*BulkDeviceBayTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkInventoryItemTemplateRequest) SetDeviceType(v BulkDeviceBayTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetParent

`func (o *BulkInventoryItemTemplateRequest) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BulkInventoryItemTemplateRequest) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BulkInventoryItemTemplateRequest) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BulkInventoryItemTemplateRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *BulkInventoryItemTemplateRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *BulkInventoryItemTemplateRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *BulkInventoryItemTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkInventoryItemTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkInventoryItemTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkInventoryItemTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkInventoryItemTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkInventoryItemTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkInventoryItemTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRole

`func (o *BulkInventoryItemTemplateRequest) GetRole() BulkInventoryItemRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkInventoryItemTemplateRequest) GetRoleOk() (*BulkInventoryItemRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkInventoryItemTemplateRequest) SetRole(v BulkInventoryItemRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkInventoryItemTemplateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkInventoryItemTemplateRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkInventoryItemTemplateRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetManufacturer

`func (o *BulkInventoryItemTemplateRequest) GetManufacturer() BulkInventoryItemRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BulkInventoryItemTemplateRequest) GetManufacturerOk() (*BulkInventoryItemRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BulkInventoryItemTemplateRequest) SetManufacturer(v BulkInventoryItemRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *BulkInventoryItemTemplateRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *BulkInventoryItemTemplateRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *BulkInventoryItemTemplateRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetPartId

`func (o *BulkInventoryItemTemplateRequest) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BulkInventoryItemTemplateRequest) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BulkInventoryItemTemplateRequest) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BulkInventoryItemTemplateRequest) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetDescription

`func (o *BulkInventoryItemTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkInventoryItemTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkInventoryItemTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkInventoryItemTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComponentType

`func (o *BulkInventoryItemTemplateRequest) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *BulkInventoryItemTemplateRequest) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *BulkInventoryItemTemplateRequest) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *BulkInventoryItemTemplateRequest) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### SetComponentTypeNil

`func (o *BulkInventoryItemTemplateRequest) SetComponentTypeNil(b bool)`

 SetComponentTypeNil sets the value for ComponentType to be an explicit nil

### UnsetComponentType
`func (o *BulkInventoryItemTemplateRequest) UnsetComponentType()`

UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
### GetComponentId

`func (o *BulkInventoryItemTemplateRequest) GetComponentId() int64`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *BulkInventoryItemTemplateRequest) GetComponentIdOk() (*int64, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *BulkInventoryItemTemplateRequest) SetComponentId(v int64)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *BulkInventoryItemTemplateRequest) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### SetComponentIdNil

`func (o *BulkInventoryItemTemplateRequest) SetComponentIdNil(b bool)`

 SetComponentIdNil sets the value for ComponentId to be an explicit nil

### UnsetComponentId
`func (o *BulkInventoryItemTemplateRequest) UnsetComponentId()`

UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


