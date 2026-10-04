# BulkInventoryItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Status** | Pointer to [**BulkInventoryItemRequestStatus**](BulkInventoryItemRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableBulkInventoryItemRequestRole**](BulkInventoryItemRequestRole.md) |  | [optional] 
**Manufacturer** | Pointer to [**NullableBulkInventoryItemRequestManufacturer**](BulkInventoryItemRequestManufacturer.md) |  | [optional] 
**PartId** | Pointer to **string** | Manufacturer-assigned part identifier | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this item | [optional] 
**Discovered** | Pointer to **bool** | This item was automatically discovered | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ComponentType** | Pointer to **NullableString** |  | [optional] 
**ComponentId** | Pointer to **NullableInt64** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkInventoryItemRequest

`func NewBulkInventoryItemRequest(id int32, device BriefCoolingIntakeRequestDevice, name string, ) *BulkInventoryItemRequest`

NewBulkInventoryItemRequest instantiates a new BulkInventoryItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInventoryItemRequestWithDefaults

`func NewBulkInventoryItemRequestWithDefaults() *BulkInventoryItemRequest`

NewBulkInventoryItemRequestWithDefaults instantiates a new BulkInventoryItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkInventoryItemRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkInventoryItemRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkInventoryItemRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *BulkInventoryItemRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkInventoryItemRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkInventoryItemRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetParent

`func (o *BulkInventoryItemRequest) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BulkInventoryItemRequest) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BulkInventoryItemRequest) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BulkInventoryItemRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *BulkInventoryItemRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *BulkInventoryItemRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *BulkInventoryItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkInventoryItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkInventoryItemRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkInventoryItemRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkInventoryItemRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkInventoryItemRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkInventoryItemRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *BulkInventoryItemRequest) GetStatus() BulkInventoryItemRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkInventoryItemRequest) GetStatusOk() (*BulkInventoryItemRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkInventoryItemRequest) SetStatus(v BulkInventoryItemRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkInventoryItemRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *BulkInventoryItemRequest) GetRole() BulkInventoryItemRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkInventoryItemRequest) GetRoleOk() (*BulkInventoryItemRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkInventoryItemRequest) SetRole(v BulkInventoryItemRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkInventoryItemRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkInventoryItemRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkInventoryItemRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetManufacturer

`func (o *BulkInventoryItemRequest) GetManufacturer() BulkInventoryItemRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BulkInventoryItemRequest) GetManufacturerOk() (*BulkInventoryItemRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BulkInventoryItemRequest) SetManufacturer(v BulkInventoryItemRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *BulkInventoryItemRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *BulkInventoryItemRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *BulkInventoryItemRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetPartId

`func (o *BulkInventoryItemRequest) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BulkInventoryItemRequest) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BulkInventoryItemRequest) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BulkInventoryItemRequest) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetSerial

`func (o *BulkInventoryItemRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *BulkInventoryItemRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *BulkInventoryItemRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *BulkInventoryItemRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *BulkInventoryItemRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *BulkInventoryItemRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *BulkInventoryItemRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *BulkInventoryItemRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *BulkInventoryItemRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *BulkInventoryItemRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDiscovered

`func (o *BulkInventoryItemRequest) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *BulkInventoryItemRequest) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *BulkInventoryItemRequest) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *BulkInventoryItemRequest) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDescription

`func (o *BulkInventoryItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkInventoryItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkInventoryItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkInventoryItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComponentType

`func (o *BulkInventoryItemRequest) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *BulkInventoryItemRequest) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *BulkInventoryItemRequest) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *BulkInventoryItemRequest) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### SetComponentTypeNil

`func (o *BulkInventoryItemRequest) SetComponentTypeNil(b bool)`

 SetComponentTypeNil sets the value for ComponentType to be an explicit nil

### UnsetComponentType
`func (o *BulkInventoryItemRequest) UnsetComponentType()`

UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
### GetComponentId

`func (o *BulkInventoryItemRequest) GetComponentId() int64`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *BulkInventoryItemRequest) GetComponentIdOk() (*int64, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *BulkInventoryItemRequest) SetComponentId(v int64)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *BulkInventoryItemRequest) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### SetComponentIdNil

`func (o *BulkInventoryItemRequest) SetComponentIdNil(b bool)`

 SetComponentIdNil sets the value for ComponentId to be an explicit nil

### UnsetComponentId
`func (o *BulkInventoryItemRequest) UnsetComponentId()`

UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
### GetOwner

`func (o *BulkInventoryItemRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkInventoryItemRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkInventoryItemRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkInventoryItemRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkInventoryItemRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkInventoryItemRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *BulkInventoryItemRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkInventoryItemRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkInventoryItemRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkInventoryItemRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkInventoryItemRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkInventoryItemRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkInventoryItemRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkInventoryItemRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


