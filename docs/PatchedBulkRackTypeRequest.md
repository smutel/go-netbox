# PatchedBulkRackTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Manufacturer** | Pointer to [**BriefDeviceTypeRequestManufacturer**](BriefDeviceTypeRequestManufacturer.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FormFactor** | Pointer to [**BulkRackTypeRequestFormFactor**](BulkRackTypeRequestFormFactor.md) |  | [optional] 
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

### NewPatchedBulkRackTypeRequest

`func NewPatchedBulkRackTypeRequest(id int32, ) *PatchedBulkRackTypeRequest`

NewPatchedBulkRackTypeRequest instantiates a new PatchedBulkRackTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkRackTypeRequestWithDefaults

`func NewPatchedBulkRackTypeRequestWithDefaults() *PatchedBulkRackTypeRequest`

NewPatchedBulkRackTypeRequestWithDefaults instantiates a new PatchedBulkRackTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkRackTypeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkRackTypeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkRackTypeRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetManufacturer

`func (o *PatchedBulkRackTypeRequest) GetManufacturer() BriefDeviceTypeRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *PatchedBulkRackTypeRequest) GetManufacturerOk() (*BriefDeviceTypeRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *PatchedBulkRackTypeRequest) SetManufacturer(v BriefDeviceTypeRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *PatchedBulkRackTypeRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *PatchedBulkRackTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PatchedBulkRackTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PatchedBulkRackTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PatchedBulkRackTypeRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedBulkRackTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedBulkRackTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedBulkRackTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedBulkRackTypeRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkRackTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkRackTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkRackTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkRackTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormFactor

`func (o *PatchedBulkRackTypeRequest) GetFormFactor() BulkRackTypeRequestFormFactor`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *PatchedBulkRackTypeRequest) GetFormFactorOk() (*BulkRackTypeRequestFormFactor, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *PatchedBulkRackTypeRequest) SetFormFactor(v BulkRackTypeRequestFormFactor)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *PatchedBulkRackTypeRequest) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### GetWidth

`func (o *PatchedBulkRackTypeRequest) GetWidth() BulkRackRequestWidth`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PatchedBulkRackTypeRequest) GetWidthOk() (*BulkRackRequestWidth, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PatchedBulkRackTypeRequest) SetWidth(v BulkRackRequestWidth)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PatchedBulkRackTypeRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetUHeight

`func (o *PatchedBulkRackTypeRequest) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *PatchedBulkRackTypeRequest) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *PatchedBulkRackTypeRequest) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *PatchedBulkRackTypeRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetStartingUnit

`func (o *PatchedBulkRackTypeRequest) GetStartingUnit() int32`

GetStartingUnit returns the StartingUnit field if non-nil, zero value otherwise.

### GetStartingUnitOk

`func (o *PatchedBulkRackTypeRequest) GetStartingUnitOk() (*int32, bool)`

GetStartingUnitOk returns a tuple with the StartingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingUnit

`func (o *PatchedBulkRackTypeRequest) SetStartingUnit(v int32)`

SetStartingUnit sets StartingUnit field to given value.

### HasStartingUnit

`func (o *PatchedBulkRackTypeRequest) HasStartingUnit() bool`

HasStartingUnit returns a boolean if a field has been set.

### GetDescUnits

`func (o *PatchedBulkRackTypeRequest) GetDescUnits() bool`

GetDescUnits returns the DescUnits field if non-nil, zero value otherwise.

### GetDescUnitsOk

`func (o *PatchedBulkRackTypeRequest) GetDescUnitsOk() (*bool, bool)`

GetDescUnitsOk returns a tuple with the DescUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescUnits

`func (o *PatchedBulkRackTypeRequest) SetDescUnits(v bool)`

SetDescUnits sets DescUnits field to given value.

### HasDescUnits

`func (o *PatchedBulkRackTypeRequest) HasDescUnits() bool`

HasDescUnits returns a boolean if a field has been set.

### GetOuterWidth

`func (o *PatchedBulkRackTypeRequest) GetOuterWidth() int32`

GetOuterWidth returns the OuterWidth field if non-nil, zero value otherwise.

### GetOuterWidthOk

`func (o *PatchedBulkRackTypeRequest) GetOuterWidthOk() (*int32, bool)`

GetOuterWidthOk returns a tuple with the OuterWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterWidth

`func (o *PatchedBulkRackTypeRequest) SetOuterWidth(v int32)`

SetOuterWidth sets OuterWidth field to given value.

### HasOuterWidth

`func (o *PatchedBulkRackTypeRequest) HasOuterWidth() bool`

HasOuterWidth returns a boolean if a field has been set.

### SetOuterWidthNil

`func (o *PatchedBulkRackTypeRequest) SetOuterWidthNil(b bool)`

 SetOuterWidthNil sets the value for OuterWidth to be an explicit nil

### UnsetOuterWidth
`func (o *PatchedBulkRackTypeRequest) UnsetOuterWidth()`

UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
### GetOuterHeight

`func (o *PatchedBulkRackTypeRequest) GetOuterHeight() int32`

GetOuterHeight returns the OuterHeight field if non-nil, zero value otherwise.

### GetOuterHeightOk

`func (o *PatchedBulkRackTypeRequest) GetOuterHeightOk() (*int32, bool)`

GetOuterHeightOk returns a tuple with the OuterHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterHeight

`func (o *PatchedBulkRackTypeRequest) SetOuterHeight(v int32)`

SetOuterHeight sets OuterHeight field to given value.

### HasOuterHeight

`func (o *PatchedBulkRackTypeRequest) HasOuterHeight() bool`

HasOuterHeight returns a boolean if a field has been set.

### SetOuterHeightNil

`func (o *PatchedBulkRackTypeRequest) SetOuterHeightNil(b bool)`

 SetOuterHeightNil sets the value for OuterHeight to be an explicit nil

### UnsetOuterHeight
`func (o *PatchedBulkRackTypeRequest) UnsetOuterHeight()`

UnsetOuterHeight ensures that no value is present for OuterHeight, not even an explicit nil
### GetOuterDepth

`func (o *PatchedBulkRackTypeRequest) GetOuterDepth() int32`

GetOuterDepth returns the OuterDepth field if non-nil, zero value otherwise.

### GetOuterDepthOk

`func (o *PatchedBulkRackTypeRequest) GetOuterDepthOk() (*int32, bool)`

GetOuterDepthOk returns a tuple with the OuterDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterDepth

`func (o *PatchedBulkRackTypeRequest) SetOuterDepth(v int32)`

SetOuterDepth sets OuterDepth field to given value.

### HasOuterDepth

`func (o *PatchedBulkRackTypeRequest) HasOuterDepth() bool`

HasOuterDepth returns a boolean if a field has been set.

### SetOuterDepthNil

`func (o *PatchedBulkRackTypeRequest) SetOuterDepthNil(b bool)`

 SetOuterDepthNil sets the value for OuterDepth to be an explicit nil

### UnsetOuterDepth
`func (o *PatchedBulkRackTypeRequest) UnsetOuterDepth()`

UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
### GetOuterUnit

`func (o *PatchedBulkRackTypeRequest) GetOuterUnit() BulkRackRequestOuterUnit`

GetOuterUnit returns the OuterUnit field if non-nil, zero value otherwise.

### GetOuterUnitOk

`func (o *PatchedBulkRackTypeRequest) GetOuterUnitOk() (*BulkRackRequestOuterUnit, bool)`

GetOuterUnitOk returns a tuple with the OuterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterUnit

`func (o *PatchedBulkRackTypeRequest) SetOuterUnit(v BulkRackRequestOuterUnit)`

SetOuterUnit sets OuterUnit field to given value.

### HasOuterUnit

`func (o *PatchedBulkRackTypeRequest) HasOuterUnit() bool`

HasOuterUnit returns a boolean if a field has been set.

### SetOuterUnitNil

`func (o *PatchedBulkRackTypeRequest) SetOuterUnitNil(b bool)`

 SetOuterUnitNil sets the value for OuterUnit to be an explicit nil

### UnsetOuterUnit
`func (o *PatchedBulkRackTypeRequest) UnsetOuterUnit()`

UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
### GetWeight

`func (o *PatchedBulkRackTypeRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkRackTypeRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkRackTypeRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkRackTypeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *PatchedBulkRackTypeRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *PatchedBulkRackTypeRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMaxWeight

`func (o *PatchedBulkRackTypeRequest) GetMaxWeight() int32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *PatchedBulkRackTypeRequest) GetMaxWeightOk() (*int32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *PatchedBulkRackTypeRequest) SetMaxWeight(v int32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *PatchedBulkRackTypeRequest) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *PatchedBulkRackTypeRequest) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *PatchedBulkRackTypeRequest) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetWeightUnit

`func (o *PatchedBulkRackTypeRequest) GetWeightUnit() BulkDeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *PatchedBulkRackTypeRequest) GetWeightUnitOk() (*BulkDeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *PatchedBulkRackTypeRequest) SetWeightUnit(v BulkDeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *PatchedBulkRackTypeRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *PatchedBulkRackTypeRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *PatchedBulkRackTypeRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetMountingDepth

`func (o *PatchedBulkRackTypeRequest) GetMountingDepth() int32`

GetMountingDepth returns the MountingDepth field if non-nil, zero value otherwise.

### GetMountingDepthOk

`func (o *PatchedBulkRackTypeRequest) GetMountingDepthOk() (*int32, bool)`

GetMountingDepthOk returns a tuple with the MountingDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountingDepth

`func (o *PatchedBulkRackTypeRequest) SetMountingDepth(v int32)`

SetMountingDepth sets MountingDepth field to given value.

### HasMountingDepth

`func (o *PatchedBulkRackTypeRequest) HasMountingDepth() bool`

HasMountingDepth returns a boolean if a field has been set.

### SetMountingDepthNil

`func (o *PatchedBulkRackTypeRequest) SetMountingDepthNil(b bool)`

 SetMountingDepthNil sets the value for MountingDepth to be an explicit nil

### UnsetMountingDepth
`func (o *PatchedBulkRackTypeRequest) UnsetMountingDepth()`

UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
### GetCoolingCapability

`func (o *PatchedBulkRackTypeRequest) GetCoolingCapability() BulkRackRequestCoolingCapability`

GetCoolingCapability returns the CoolingCapability field if non-nil, zero value otherwise.

### GetCoolingCapabilityOk

`func (o *PatchedBulkRackTypeRequest) GetCoolingCapabilityOk() (*BulkRackRequestCoolingCapability, bool)`

GetCoolingCapabilityOk returns a tuple with the CoolingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapability

`func (o *PatchedBulkRackTypeRequest) SetCoolingCapability(v BulkRackRequestCoolingCapability)`

SetCoolingCapability sets CoolingCapability field to given value.

### HasCoolingCapability

`func (o *PatchedBulkRackTypeRequest) HasCoolingCapability() bool`

HasCoolingCapability returns a boolean if a field has been set.

### SetCoolingCapabilityNil

`func (o *PatchedBulkRackTypeRequest) SetCoolingCapabilityNil(b bool)`

 SetCoolingCapabilityNil sets the value for CoolingCapability to be an explicit nil

### UnsetCoolingCapability
`func (o *PatchedBulkRackTypeRequest) UnsetCoolingCapability()`

UnsetCoolingCapability ensures that no value is present for CoolingCapability, not even an explicit nil
### GetCoolingCapacity

`func (o *PatchedBulkRackTypeRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *PatchedBulkRackTypeRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *PatchedBulkRackTypeRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *PatchedBulkRackTypeRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *PatchedBulkRackTypeRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *PatchedBulkRackTypeRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetOwner

`func (o *PatchedBulkRackTypeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkRackTypeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkRackTypeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkRackTypeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkRackTypeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkRackTypeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkRackTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkRackTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkRackTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkRackTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkRackTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkRackTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkRackTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkRackTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkRackTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkRackTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkRackTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkRackTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


