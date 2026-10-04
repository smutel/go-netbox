# BulkRackTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Manufacturer** | [**BriefDeviceTypeRequestManufacturer**](BriefDeviceTypeRequestManufacturer.md) |  | 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FormFactor** | [**BulkRackTypeRequestFormFactor**](BulkRackTypeRequestFormFactor.md) |  | 
**Width** | Pointer to [**BulkRackRequestWidth**](BulkRackRequestWidth.md) |  | [optional] 
**UHeight** | Pointer to **int32** | Height in rack units | [optional] 
**StartingUnit** | Pointer to **int32** | Starting unit for rack | [optional] 
**DescUnits** | Pointer to **bool** | Units are numbered top-to-bottom | [optional] 
**OuterWidth** | Pointer to **NullableInt32** | Outer dimension of rack (width) | [optional] 
**OuterHeight** | Pointer to **NullableInt32** | Outer dimension of rack (height) | [optional] 
**OuterDepth** | Pointer to **NullableInt32** | Outer dimension of rack (depth) | [optional] 
**OuterUnit** | Pointer to [**NullableBulkRackRequestOuterUnit**](BulkRackRequestOuterUnit.md) |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**MaxWeight** | Pointer to **NullableInt32** | Maximum load capacity for the rack | [optional] 
**WeightUnit** | Pointer to [**NullableBulkDeviceTypeRequestWeightUnit**](BulkDeviceTypeRequestWeightUnit.md) |  | [optional] 
**MountingDepth** | Pointer to **NullableInt32** | Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails. | [optional] 
**CoolingCapability** | Pointer to [**NullableBulkRackRequestCoolingCapability**](BulkRackRequestCoolingCapability.md) |  | [optional] 
**CoolingCapacity** | Pointer to **NullableFloat64** | Cooling capacity (kW) | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkRackTypeRequest

`func NewBulkRackTypeRequest(id int32, manufacturer BriefDeviceTypeRequestManufacturer, model string, slug string, formFactor BulkRackTypeRequestFormFactor, ) *BulkRackTypeRequest`

NewBulkRackTypeRequest instantiates a new BulkRackTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRackTypeRequestWithDefaults

`func NewBulkRackTypeRequestWithDefaults() *BulkRackTypeRequest`

NewBulkRackTypeRequestWithDefaults instantiates a new BulkRackTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkRackTypeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkRackTypeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkRackTypeRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetManufacturer

`func (o *BulkRackTypeRequest) GetManufacturer() BriefDeviceTypeRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BulkRackTypeRequest) GetManufacturerOk() (*BriefDeviceTypeRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BulkRackTypeRequest) SetManufacturer(v BriefDeviceTypeRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *BulkRackTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BulkRackTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BulkRackTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *BulkRackTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BulkRackTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BulkRackTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BulkRackTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkRackTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkRackTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkRackTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormFactor

`func (o *BulkRackTypeRequest) GetFormFactor() BulkRackTypeRequestFormFactor`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *BulkRackTypeRequest) GetFormFactorOk() (*BulkRackTypeRequestFormFactor, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *BulkRackTypeRequest) SetFormFactor(v BulkRackTypeRequestFormFactor)`

SetFormFactor sets FormFactor field to given value.


### GetWidth

`func (o *BulkRackTypeRequest) GetWidth() BulkRackRequestWidth`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BulkRackTypeRequest) GetWidthOk() (*BulkRackRequestWidth, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BulkRackTypeRequest) SetWidth(v BulkRackRequestWidth)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *BulkRackTypeRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetUHeight

`func (o *BulkRackTypeRequest) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *BulkRackTypeRequest) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *BulkRackTypeRequest) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *BulkRackTypeRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetStartingUnit

`func (o *BulkRackTypeRequest) GetStartingUnit() int32`

GetStartingUnit returns the StartingUnit field if non-nil, zero value otherwise.

### GetStartingUnitOk

`func (o *BulkRackTypeRequest) GetStartingUnitOk() (*int32, bool)`

GetStartingUnitOk returns a tuple with the StartingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingUnit

`func (o *BulkRackTypeRequest) SetStartingUnit(v int32)`

SetStartingUnit sets StartingUnit field to given value.

### HasStartingUnit

`func (o *BulkRackTypeRequest) HasStartingUnit() bool`

HasStartingUnit returns a boolean if a field has been set.

### GetDescUnits

`func (o *BulkRackTypeRequest) GetDescUnits() bool`

GetDescUnits returns the DescUnits field if non-nil, zero value otherwise.

### GetDescUnitsOk

`func (o *BulkRackTypeRequest) GetDescUnitsOk() (*bool, bool)`

GetDescUnitsOk returns a tuple with the DescUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescUnits

`func (o *BulkRackTypeRequest) SetDescUnits(v bool)`

SetDescUnits sets DescUnits field to given value.

### HasDescUnits

`func (o *BulkRackTypeRequest) HasDescUnits() bool`

HasDescUnits returns a boolean if a field has been set.

### GetOuterWidth

`func (o *BulkRackTypeRequest) GetOuterWidth() int32`

GetOuterWidth returns the OuterWidth field if non-nil, zero value otherwise.

### GetOuterWidthOk

`func (o *BulkRackTypeRequest) GetOuterWidthOk() (*int32, bool)`

GetOuterWidthOk returns a tuple with the OuterWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterWidth

`func (o *BulkRackTypeRequest) SetOuterWidth(v int32)`

SetOuterWidth sets OuterWidth field to given value.

### HasOuterWidth

`func (o *BulkRackTypeRequest) HasOuterWidth() bool`

HasOuterWidth returns a boolean if a field has been set.

### SetOuterWidthNil

`func (o *BulkRackTypeRequest) SetOuterWidthNil(b bool)`

 SetOuterWidthNil sets the value for OuterWidth to be an explicit nil

### UnsetOuterWidth
`func (o *BulkRackTypeRequest) UnsetOuterWidth()`

UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
### GetOuterHeight

`func (o *BulkRackTypeRequest) GetOuterHeight() int32`

GetOuterHeight returns the OuterHeight field if non-nil, zero value otherwise.

### GetOuterHeightOk

`func (o *BulkRackTypeRequest) GetOuterHeightOk() (*int32, bool)`

GetOuterHeightOk returns a tuple with the OuterHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterHeight

`func (o *BulkRackTypeRequest) SetOuterHeight(v int32)`

SetOuterHeight sets OuterHeight field to given value.

### HasOuterHeight

`func (o *BulkRackTypeRequest) HasOuterHeight() bool`

HasOuterHeight returns a boolean if a field has been set.

### SetOuterHeightNil

`func (o *BulkRackTypeRequest) SetOuterHeightNil(b bool)`

 SetOuterHeightNil sets the value for OuterHeight to be an explicit nil

### UnsetOuterHeight
`func (o *BulkRackTypeRequest) UnsetOuterHeight()`

UnsetOuterHeight ensures that no value is present for OuterHeight, not even an explicit nil
### GetOuterDepth

`func (o *BulkRackTypeRequest) GetOuterDepth() int32`

GetOuterDepth returns the OuterDepth field if non-nil, zero value otherwise.

### GetOuterDepthOk

`func (o *BulkRackTypeRequest) GetOuterDepthOk() (*int32, bool)`

GetOuterDepthOk returns a tuple with the OuterDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterDepth

`func (o *BulkRackTypeRequest) SetOuterDepth(v int32)`

SetOuterDepth sets OuterDepth field to given value.

### HasOuterDepth

`func (o *BulkRackTypeRequest) HasOuterDepth() bool`

HasOuterDepth returns a boolean if a field has been set.

### SetOuterDepthNil

`func (o *BulkRackTypeRequest) SetOuterDepthNil(b bool)`

 SetOuterDepthNil sets the value for OuterDepth to be an explicit nil

### UnsetOuterDepth
`func (o *BulkRackTypeRequest) UnsetOuterDepth()`

UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
### GetOuterUnit

`func (o *BulkRackTypeRequest) GetOuterUnit() BulkRackRequestOuterUnit`

GetOuterUnit returns the OuterUnit field if non-nil, zero value otherwise.

### GetOuterUnitOk

`func (o *BulkRackTypeRequest) GetOuterUnitOk() (*BulkRackRequestOuterUnit, bool)`

GetOuterUnitOk returns a tuple with the OuterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterUnit

`func (o *BulkRackTypeRequest) SetOuterUnit(v BulkRackRequestOuterUnit)`

SetOuterUnit sets OuterUnit field to given value.

### HasOuterUnit

`func (o *BulkRackTypeRequest) HasOuterUnit() bool`

HasOuterUnit returns a boolean if a field has been set.

### SetOuterUnitNil

`func (o *BulkRackTypeRequest) SetOuterUnitNil(b bool)`

 SetOuterUnitNil sets the value for OuterUnit to be an explicit nil

### UnsetOuterUnit
`func (o *BulkRackTypeRequest) UnsetOuterUnit()`

UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
### GetWeight

`func (o *BulkRackTypeRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkRackTypeRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkRackTypeRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkRackTypeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *BulkRackTypeRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *BulkRackTypeRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMaxWeight

`func (o *BulkRackTypeRequest) GetMaxWeight() int32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *BulkRackTypeRequest) GetMaxWeightOk() (*int32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *BulkRackTypeRequest) SetMaxWeight(v int32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *BulkRackTypeRequest) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *BulkRackTypeRequest) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *BulkRackTypeRequest) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetWeightUnit

`func (o *BulkRackTypeRequest) GetWeightUnit() BulkDeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *BulkRackTypeRequest) GetWeightUnitOk() (*BulkDeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *BulkRackTypeRequest) SetWeightUnit(v BulkDeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *BulkRackTypeRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *BulkRackTypeRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *BulkRackTypeRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetMountingDepth

`func (o *BulkRackTypeRequest) GetMountingDepth() int32`

GetMountingDepth returns the MountingDepth field if non-nil, zero value otherwise.

### GetMountingDepthOk

`func (o *BulkRackTypeRequest) GetMountingDepthOk() (*int32, bool)`

GetMountingDepthOk returns a tuple with the MountingDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountingDepth

`func (o *BulkRackTypeRequest) SetMountingDepth(v int32)`

SetMountingDepth sets MountingDepth field to given value.

### HasMountingDepth

`func (o *BulkRackTypeRequest) HasMountingDepth() bool`

HasMountingDepth returns a boolean if a field has been set.

### SetMountingDepthNil

`func (o *BulkRackTypeRequest) SetMountingDepthNil(b bool)`

 SetMountingDepthNil sets the value for MountingDepth to be an explicit nil

### UnsetMountingDepth
`func (o *BulkRackTypeRequest) UnsetMountingDepth()`

UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
### GetCoolingCapability

`func (o *BulkRackTypeRequest) GetCoolingCapability() BulkRackRequestCoolingCapability`

GetCoolingCapability returns the CoolingCapability field if non-nil, zero value otherwise.

### GetCoolingCapabilityOk

`func (o *BulkRackTypeRequest) GetCoolingCapabilityOk() (*BulkRackRequestCoolingCapability, bool)`

GetCoolingCapabilityOk returns a tuple with the CoolingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapability

`func (o *BulkRackTypeRequest) SetCoolingCapability(v BulkRackRequestCoolingCapability)`

SetCoolingCapability sets CoolingCapability field to given value.

### HasCoolingCapability

`func (o *BulkRackTypeRequest) HasCoolingCapability() bool`

HasCoolingCapability returns a boolean if a field has been set.

### SetCoolingCapabilityNil

`func (o *BulkRackTypeRequest) SetCoolingCapabilityNil(b bool)`

 SetCoolingCapabilityNil sets the value for CoolingCapability to be an explicit nil

### UnsetCoolingCapability
`func (o *BulkRackTypeRequest) UnsetCoolingCapability()`

UnsetCoolingCapability ensures that no value is present for CoolingCapability, not even an explicit nil
### GetCoolingCapacity

`func (o *BulkRackTypeRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *BulkRackTypeRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *BulkRackTypeRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *BulkRackTypeRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *BulkRackTypeRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *BulkRackTypeRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetOwner

`func (o *BulkRackTypeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkRackTypeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkRackTypeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkRackTypeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkRackTypeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkRackTypeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkRackTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkRackTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkRackTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkRackTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkRackTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkRackTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkRackTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkRackTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkRackTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkRackTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkRackTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkRackTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


