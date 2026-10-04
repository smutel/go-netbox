# PatchedBulkRackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**FacilityId** | Pointer to **NullableString** |  | [optional] 
**Site** | Pointer to [**BulkCoolingSourceRequestSite**](BulkCoolingSourceRequestSite.md) |  | [optional] 
**Location** | Pointer to [**NullableBulkCoolingSourceRequestLocation**](BulkCoolingSourceRequestLocation.md) |  | [optional] 
**Group** | Pointer to [**NullableBulkRackRequestGroup**](BulkRackRequestGroup.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Status** | Pointer to [**BulkRackRequestStatus**](BulkRackRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableBulkRackRequestRole**](BulkRackRequestRole.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this rack | [optional] 
**RackType** | Pointer to [**NullableBulkRackRequestRackType**](BulkRackRequestRackType.md) |  | [optional] 
**FormFactor** | Pointer to [**NullableBulkRackRequestFormFactor**](BulkRackRequestFormFactor.md) |  | [optional] 
**Width** | Pointer to [**BulkRackRequestWidth**](BulkRackRequestWidth.md) |  | [optional] 
**UHeight** | Pointer to **int32** | Height in rack units | [optional] 
**StartingUnit** | Pointer to **int32** | Starting unit for rack | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**MaxWeight** | Pointer to **NullableInt32** | Maximum load capacity for the rack | [optional] 
**WeightUnit** | Pointer to [**NullableBulkDeviceTypeRequestWeightUnit**](BulkDeviceTypeRequestWeightUnit.md) |  | [optional] 
**DescUnits** | Pointer to **bool** | Units are numbered top-to-bottom | [optional] 
**OuterWidth** | Pointer to **NullableInt32** | Outer dimension of rack (width) | [optional] 
**OuterHeight** | Pointer to **NullableInt32** | Outer dimension of rack (height) | [optional] 
**OuterDepth** | Pointer to **NullableInt32** | Outer dimension of rack (depth) | [optional] 
**OuterUnit** | Pointer to [**NullableBulkRackRequestOuterUnit**](BulkRackRequestOuterUnit.md) |  | [optional] 
**MountingDepth** | Pointer to **NullableInt32** | Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails. | [optional] 
**Airflow** | Pointer to [**BulkRackRequestAirflow**](BulkRackRequestAirflow.md) |  | [optional] 
**CoolingCapability** | Pointer to [**NullableBulkRackRequestCoolingCapability**](BulkRackRequestCoolingCapability.md) |  | [optional] 
**CoolingCapacity** | Pointer to **NullableFloat64** | Cooling capacity (kW) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkRackRequest

`func NewPatchedBulkRackRequest(id int32, ) *PatchedBulkRackRequest`

NewPatchedBulkRackRequest instantiates a new PatchedBulkRackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkRackRequestWithDefaults

`func NewPatchedBulkRackRequestWithDefaults() *PatchedBulkRackRequest`

NewPatchedBulkRackRequestWithDefaults instantiates a new PatchedBulkRackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkRackRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkRackRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkRackRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkRackRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkRackRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkRackRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkRackRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFacilityId

`func (o *PatchedBulkRackRequest) GetFacilityId() string`

GetFacilityId returns the FacilityId field if non-nil, zero value otherwise.

### GetFacilityIdOk

`func (o *PatchedBulkRackRequest) GetFacilityIdOk() (*string, bool)`

GetFacilityIdOk returns a tuple with the FacilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityId

`func (o *PatchedBulkRackRequest) SetFacilityId(v string)`

SetFacilityId sets FacilityId field to given value.

### HasFacilityId

`func (o *PatchedBulkRackRequest) HasFacilityId() bool`

HasFacilityId returns a boolean if a field has been set.

### SetFacilityIdNil

`func (o *PatchedBulkRackRequest) SetFacilityIdNil(b bool)`

 SetFacilityIdNil sets the value for FacilityId to be an explicit nil

### UnsetFacilityId
`func (o *PatchedBulkRackRequest) UnsetFacilityId()`

UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
### GetSite

`func (o *PatchedBulkRackRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedBulkRackRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedBulkRackRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedBulkRackRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedBulkRackRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedBulkRackRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedBulkRackRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedBulkRackRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedBulkRackRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedBulkRackRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetGroup

`func (o *PatchedBulkRackRequest) GetGroup() BulkRackRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkRackRequest) GetGroupOk() (*BulkRackRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkRackRequest) SetGroup(v BulkRackRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkRackRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedBulkRackRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedBulkRackRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetTenant

`func (o *PatchedBulkRackRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkRackRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkRackRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkRackRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkRackRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkRackRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *PatchedBulkRackRequest) GetStatus() BulkRackRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkRackRequest) GetStatusOk() (*BulkRackRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkRackRequest) SetStatus(v BulkRackRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkRackRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkRackRequest) GetRole() BulkRackRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkRackRequest) GetRoleOk() (*BulkRackRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkRackRequest) SetRole(v BulkRackRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkRackRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkRackRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkRackRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSerial

`func (o *PatchedBulkRackRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedBulkRackRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedBulkRackRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedBulkRackRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedBulkRackRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedBulkRackRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedBulkRackRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedBulkRackRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedBulkRackRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedBulkRackRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetRackType

`func (o *PatchedBulkRackRequest) GetRackType() BulkRackRequestRackType`

GetRackType returns the RackType field if non-nil, zero value otherwise.

### GetRackTypeOk

`func (o *PatchedBulkRackRequest) GetRackTypeOk() (*BulkRackRequestRackType, bool)`

GetRackTypeOk returns a tuple with the RackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackType

`func (o *PatchedBulkRackRequest) SetRackType(v BulkRackRequestRackType)`

SetRackType sets RackType field to given value.

### HasRackType

`func (o *PatchedBulkRackRequest) HasRackType() bool`

HasRackType returns a boolean if a field has been set.

### SetRackTypeNil

`func (o *PatchedBulkRackRequest) SetRackTypeNil(b bool)`

 SetRackTypeNil sets the value for RackType to be an explicit nil

### UnsetRackType
`func (o *PatchedBulkRackRequest) UnsetRackType()`

UnsetRackType ensures that no value is present for RackType, not even an explicit nil
### GetFormFactor

`func (o *PatchedBulkRackRequest) GetFormFactor() BulkRackRequestFormFactor`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *PatchedBulkRackRequest) GetFormFactorOk() (*BulkRackRequestFormFactor, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *PatchedBulkRackRequest) SetFormFactor(v BulkRackRequestFormFactor)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *PatchedBulkRackRequest) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### SetFormFactorNil

`func (o *PatchedBulkRackRequest) SetFormFactorNil(b bool)`

 SetFormFactorNil sets the value for FormFactor to be an explicit nil

### UnsetFormFactor
`func (o *PatchedBulkRackRequest) UnsetFormFactor()`

UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
### GetWidth

`func (o *PatchedBulkRackRequest) GetWidth() BulkRackRequestWidth`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PatchedBulkRackRequest) GetWidthOk() (*BulkRackRequestWidth, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PatchedBulkRackRequest) SetWidth(v BulkRackRequestWidth)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PatchedBulkRackRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetUHeight

`func (o *PatchedBulkRackRequest) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *PatchedBulkRackRequest) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *PatchedBulkRackRequest) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *PatchedBulkRackRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetStartingUnit

`func (o *PatchedBulkRackRequest) GetStartingUnit() int32`

GetStartingUnit returns the StartingUnit field if non-nil, zero value otherwise.

### GetStartingUnitOk

`func (o *PatchedBulkRackRequest) GetStartingUnitOk() (*int32, bool)`

GetStartingUnitOk returns a tuple with the StartingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingUnit

`func (o *PatchedBulkRackRequest) SetStartingUnit(v int32)`

SetStartingUnit sets StartingUnit field to given value.

### HasStartingUnit

`func (o *PatchedBulkRackRequest) HasStartingUnit() bool`

HasStartingUnit returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkRackRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkRackRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkRackRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkRackRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *PatchedBulkRackRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *PatchedBulkRackRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMaxWeight

`func (o *PatchedBulkRackRequest) GetMaxWeight() int32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *PatchedBulkRackRequest) GetMaxWeightOk() (*int32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *PatchedBulkRackRequest) SetMaxWeight(v int32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *PatchedBulkRackRequest) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *PatchedBulkRackRequest) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *PatchedBulkRackRequest) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetWeightUnit

`func (o *PatchedBulkRackRequest) GetWeightUnit() BulkDeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *PatchedBulkRackRequest) GetWeightUnitOk() (*BulkDeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *PatchedBulkRackRequest) SetWeightUnit(v BulkDeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *PatchedBulkRackRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *PatchedBulkRackRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *PatchedBulkRackRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetDescUnits

`func (o *PatchedBulkRackRequest) GetDescUnits() bool`

GetDescUnits returns the DescUnits field if non-nil, zero value otherwise.

### GetDescUnitsOk

`func (o *PatchedBulkRackRequest) GetDescUnitsOk() (*bool, bool)`

GetDescUnitsOk returns a tuple with the DescUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescUnits

`func (o *PatchedBulkRackRequest) SetDescUnits(v bool)`

SetDescUnits sets DescUnits field to given value.

### HasDescUnits

`func (o *PatchedBulkRackRequest) HasDescUnits() bool`

HasDescUnits returns a boolean if a field has been set.

### GetOuterWidth

`func (o *PatchedBulkRackRequest) GetOuterWidth() int32`

GetOuterWidth returns the OuterWidth field if non-nil, zero value otherwise.

### GetOuterWidthOk

`func (o *PatchedBulkRackRequest) GetOuterWidthOk() (*int32, bool)`

GetOuterWidthOk returns a tuple with the OuterWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterWidth

`func (o *PatchedBulkRackRequest) SetOuterWidth(v int32)`

SetOuterWidth sets OuterWidth field to given value.

### HasOuterWidth

`func (o *PatchedBulkRackRequest) HasOuterWidth() bool`

HasOuterWidth returns a boolean if a field has been set.

### SetOuterWidthNil

`func (o *PatchedBulkRackRequest) SetOuterWidthNil(b bool)`

 SetOuterWidthNil sets the value for OuterWidth to be an explicit nil

### UnsetOuterWidth
`func (o *PatchedBulkRackRequest) UnsetOuterWidth()`

UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
### GetOuterHeight

`func (o *PatchedBulkRackRequest) GetOuterHeight() int32`

GetOuterHeight returns the OuterHeight field if non-nil, zero value otherwise.

### GetOuterHeightOk

`func (o *PatchedBulkRackRequest) GetOuterHeightOk() (*int32, bool)`

GetOuterHeightOk returns a tuple with the OuterHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterHeight

`func (o *PatchedBulkRackRequest) SetOuterHeight(v int32)`

SetOuterHeight sets OuterHeight field to given value.

### HasOuterHeight

`func (o *PatchedBulkRackRequest) HasOuterHeight() bool`

HasOuterHeight returns a boolean if a field has been set.

### SetOuterHeightNil

`func (o *PatchedBulkRackRequest) SetOuterHeightNil(b bool)`

 SetOuterHeightNil sets the value for OuterHeight to be an explicit nil

### UnsetOuterHeight
`func (o *PatchedBulkRackRequest) UnsetOuterHeight()`

UnsetOuterHeight ensures that no value is present for OuterHeight, not even an explicit nil
### GetOuterDepth

`func (o *PatchedBulkRackRequest) GetOuterDepth() int32`

GetOuterDepth returns the OuterDepth field if non-nil, zero value otherwise.

### GetOuterDepthOk

`func (o *PatchedBulkRackRequest) GetOuterDepthOk() (*int32, bool)`

GetOuterDepthOk returns a tuple with the OuterDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterDepth

`func (o *PatchedBulkRackRequest) SetOuterDepth(v int32)`

SetOuterDepth sets OuterDepth field to given value.

### HasOuterDepth

`func (o *PatchedBulkRackRequest) HasOuterDepth() bool`

HasOuterDepth returns a boolean if a field has been set.

### SetOuterDepthNil

`func (o *PatchedBulkRackRequest) SetOuterDepthNil(b bool)`

 SetOuterDepthNil sets the value for OuterDepth to be an explicit nil

### UnsetOuterDepth
`func (o *PatchedBulkRackRequest) UnsetOuterDepth()`

UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
### GetOuterUnit

`func (o *PatchedBulkRackRequest) GetOuterUnit() BulkRackRequestOuterUnit`

GetOuterUnit returns the OuterUnit field if non-nil, zero value otherwise.

### GetOuterUnitOk

`func (o *PatchedBulkRackRequest) GetOuterUnitOk() (*BulkRackRequestOuterUnit, bool)`

GetOuterUnitOk returns a tuple with the OuterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterUnit

`func (o *PatchedBulkRackRequest) SetOuterUnit(v BulkRackRequestOuterUnit)`

SetOuterUnit sets OuterUnit field to given value.

### HasOuterUnit

`func (o *PatchedBulkRackRequest) HasOuterUnit() bool`

HasOuterUnit returns a boolean if a field has been set.

### SetOuterUnitNil

`func (o *PatchedBulkRackRequest) SetOuterUnitNil(b bool)`

 SetOuterUnitNil sets the value for OuterUnit to be an explicit nil

### UnsetOuterUnit
`func (o *PatchedBulkRackRequest) UnsetOuterUnit()`

UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
### GetMountingDepth

`func (o *PatchedBulkRackRequest) GetMountingDepth() int32`

GetMountingDepth returns the MountingDepth field if non-nil, zero value otherwise.

### GetMountingDepthOk

`func (o *PatchedBulkRackRequest) GetMountingDepthOk() (*int32, bool)`

GetMountingDepthOk returns a tuple with the MountingDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountingDepth

`func (o *PatchedBulkRackRequest) SetMountingDepth(v int32)`

SetMountingDepth sets MountingDepth field to given value.

### HasMountingDepth

`func (o *PatchedBulkRackRequest) HasMountingDepth() bool`

HasMountingDepth returns a boolean if a field has been set.

### SetMountingDepthNil

`func (o *PatchedBulkRackRequest) SetMountingDepthNil(b bool)`

 SetMountingDepthNil sets the value for MountingDepth to be an explicit nil

### UnsetMountingDepth
`func (o *PatchedBulkRackRequest) UnsetMountingDepth()`

UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
### GetAirflow

`func (o *PatchedBulkRackRequest) GetAirflow() BulkRackRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *PatchedBulkRackRequest) GetAirflowOk() (*BulkRackRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *PatchedBulkRackRequest) SetAirflow(v BulkRackRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *PatchedBulkRackRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### GetCoolingCapability

`func (o *PatchedBulkRackRequest) GetCoolingCapability() BulkRackRequestCoolingCapability`

GetCoolingCapability returns the CoolingCapability field if non-nil, zero value otherwise.

### GetCoolingCapabilityOk

`func (o *PatchedBulkRackRequest) GetCoolingCapabilityOk() (*BulkRackRequestCoolingCapability, bool)`

GetCoolingCapabilityOk returns a tuple with the CoolingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapability

`func (o *PatchedBulkRackRequest) SetCoolingCapability(v BulkRackRequestCoolingCapability)`

SetCoolingCapability sets CoolingCapability field to given value.

### HasCoolingCapability

`func (o *PatchedBulkRackRequest) HasCoolingCapability() bool`

HasCoolingCapability returns a boolean if a field has been set.

### SetCoolingCapabilityNil

`func (o *PatchedBulkRackRequest) SetCoolingCapabilityNil(b bool)`

 SetCoolingCapabilityNil sets the value for CoolingCapability to be an explicit nil

### UnsetCoolingCapability
`func (o *PatchedBulkRackRequest) UnsetCoolingCapability()`

UnsetCoolingCapability ensures that no value is present for CoolingCapability, not even an explicit nil
### GetCoolingCapacity

`func (o *PatchedBulkRackRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *PatchedBulkRackRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *PatchedBulkRackRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *PatchedBulkRackRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *PatchedBulkRackRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *PatchedBulkRackRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetDescription

`func (o *PatchedBulkRackRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkRackRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkRackRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkRackRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkRackRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkRackRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkRackRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkRackRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkRackRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkRackRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkRackRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkRackRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkRackRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkRackRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkRackRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkRackRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkRackRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkRackRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkRackRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkRackRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkRackRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkRackRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


