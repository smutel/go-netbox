# PatchedWritableRackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Width** | Pointer to [**PatchedWritableRackRequestWidth**](PatchedWritableRackRequestWidth.md) |  | [optional] 
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
**Airflow** | Pointer to [**NullablePatchedWritableRackRequestAirflow**](PatchedWritableRackRequestAirflow.md) |  | [optional] 
**CoolingCapability** | Pointer to [**NullableBulkRackRequestCoolingCapability**](BulkRackRequestCoolingCapability.md) |  | [optional] 
**CoolingCapacity** | Pointer to **NullableFloat64** | Cooling capacity (kW) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableRackRequest

`func NewPatchedWritableRackRequest() *PatchedWritableRackRequest`

NewPatchedWritableRackRequest instantiates a new PatchedWritableRackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableRackRequestWithDefaults

`func NewPatchedWritableRackRequestWithDefaults() *PatchedWritableRackRequest`

NewPatchedWritableRackRequestWithDefaults instantiates a new PatchedWritableRackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableRackRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableRackRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableRackRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableRackRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFacilityId

`func (o *PatchedWritableRackRequest) GetFacilityId() string`

GetFacilityId returns the FacilityId field if non-nil, zero value otherwise.

### GetFacilityIdOk

`func (o *PatchedWritableRackRequest) GetFacilityIdOk() (*string, bool)`

GetFacilityIdOk returns a tuple with the FacilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityId

`func (o *PatchedWritableRackRequest) SetFacilityId(v string)`

SetFacilityId sets FacilityId field to given value.

### HasFacilityId

`func (o *PatchedWritableRackRequest) HasFacilityId() bool`

HasFacilityId returns a boolean if a field has been set.

### SetFacilityIdNil

`func (o *PatchedWritableRackRequest) SetFacilityIdNil(b bool)`

 SetFacilityIdNil sets the value for FacilityId to be an explicit nil

### UnsetFacilityId
`func (o *PatchedWritableRackRequest) UnsetFacilityId()`

UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
### GetSite

`func (o *PatchedWritableRackRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableRackRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableRackRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableRackRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedWritableRackRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWritableRackRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWritableRackRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWritableRackRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedWritableRackRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedWritableRackRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetGroup

`func (o *PatchedWritableRackRequest) GetGroup() BulkRackRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedWritableRackRequest) GetGroupOk() (*BulkRackRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedWritableRackRequest) SetGroup(v BulkRackRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedWritableRackRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedWritableRackRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedWritableRackRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetTenant

`func (o *PatchedWritableRackRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableRackRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableRackRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableRackRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableRackRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableRackRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *PatchedWritableRackRequest) GetStatus() BulkRackRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableRackRequest) GetStatusOk() (*BulkRackRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableRackRequest) SetStatus(v BulkRackRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableRackRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableRackRequest) GetRole() BulkRackRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableRackRequest) GetRoleOk() (*BulkRackRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableRackRequest) SetRole(v BulkRackRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableRackRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedWritableRackRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedWritableRackRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSerial

`func (o *PatchedWritableRackRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableRackRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableRackRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableRackRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedWritableRackRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedWritableRackRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedWritableRackRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedWritableRackRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedWritableRackRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedWritableRackRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetRackType

`func (o *PatchedWritableRackRequest) GetRackType() BulkRackRequestRackType`

GetRackType returns the RackType field if non-nil, zero value otherwise.

### GetRackTypeOk

`func (o *PatchedWritableRackRequest) GetRackTypeOk() (*BulkRackRequestRackType, bool)`

GetRackTypeOk returns a tuple with the RackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackType

`func (o *PatchedWritableRackRequest) SetRackType(v BulkRackRequestRackType)`

SetRackType sets RackType field to given value.

### HasRackType

`func (o *PatchedWritableRackRequest) HasRackType() bool`

HasRackType returns a boolean if a field has been set.

### SetRackTypeNil

`func (o *PatchedWritableRackRequest) SetRackTypeNil(b bool)`

 SetRackTypeNil sets the value for RackType to be an explicit nil

### UnsetRackType
`func (o *PatchedWritableRackRequest) UnsetRackType()`

UnsetRackType ensures that no value is present for RackType, not even an explicit nil
### GetFormFactor

`func (o *PatchedWritableRackRequest) GetFormFactor() BulkRackRequestFormFactor`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *PatchedWritableRackRequest) GetFormFactorOk() (*BulkRackRequestFormFactor, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *PatchedWritableRackRequest) SetFormFactor(v BulkRackRequestFormFactor)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *PatchedWritableRackRequest) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### SetFormFactorNil

`func (o *PatchedWritableRackRequest) SetFormFactorNil(b bool)`

 SetFormFactorNil sets the value for FormFactor to be an explicit nil

### UnsetFormFactor
`func (o *PatchedWritableRackRequest) UnsetFormFactor()`

UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
### GetWidth

`func (o *PatchedWritableRackRequest) GetWidth() PatchedWritableRackRequestWidth`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PatchedWritableRackRequest) GetWidthOk() (*PatchedWritableRackRequestWidth, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PatchedWritableRackRequest) SetWidth(v PatchedWritableRackRequestWidth)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PatchedWritableRackRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetUHeight

`func (o *PatchedWritableRackRequest) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *PatchedWritableRackRequest) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *PatchedWritableRackRequest) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *PatchedWritableRackRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetStartingUnit

`func (o *PatchedWritableRackRequest) GetStartingUnit() int32`

GetStartingUnit returns the StartingUnit field if non-nil, zero value otherwise.

### GetStartingUnitOk

`func (o *PatchedWritableRackRequest) GetStartingUnitOk() (*int32, bool)`

GetStartingUnitOk returns a tuple with the StartingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingUnit

`func (o *PatchedWritableRackRequest) SetStartingUnit(v int32)`

SetStartingUnit sets StartingUnit field to given value.

### HasStartingUnit

`func (o *PatchedWritableRackRequest) HasStartingUnit() bool`

HasStartingUnit returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedWritableRackRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedWritableRackRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedWritableRackRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedWritableRackRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *PatchedWritableRackRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *PatchedWritableRackRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMaxWeight

`func (o *PatchedWritableRackRequest) GetMaxWeight() int32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *PatchedWritableRackRequest) GetMaxWeightOk() (*int32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *PatchedWritableRackRequest) SetMaxWeight(v int32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *PatchedWritableRackRequest) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *PatchedWritableRackRequest) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *PatchedWritableRackRequest) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetWeightUnit

`func (o *PatchedWritableRackRequest) GetWeightUnit() BulkDeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *PatchedWritableRackRequest) GetWeightUnitOk() (*BulkDeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *PatchedWritableRackRequest) SetWeightUnit(v BulkDeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *PatchedWritableRackRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *PatchedWritableRackRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *PatchedWritableRackRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetDescUnits

`func (o *PatchedWritableRackRequest) GetDescUnits() bool`

GetDescUnits returns the DescUnits field if non-nil, zero value otherwise.

### GetDescUnitsOk

`func (o *PatchedWritableRackRequest) GetDescUnitsOk() (*bool, bool)`

GetDescUnitsOk returns a tuple with the DescUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescUnits

`func (o *PatchedWritableRackRequest) SetDescUnits(v bool)`

SetDescUnits sets DescUnits field to given value.

### HasDescUnits

`func (o *PatchedWritableRackRequest) HasDescUnits() bool`

HasDescUnits returns a boolean if a field has been set.

### GetOuterWidth

`func (o *PatchedWritableRackRequest) GetOuterWidth() int32`

GetOuterWidth returns the OuterWidth field if non-nil, zero value otherwise.

### GetOuterWidthOk

`func (o *PatchedWritableRackRequest) GetOuterWidthOk() (*int32, bool)`

GetOuterWidthOk returns a tuple with the OuterWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterWidth

`func (o *PatchedWritableRackRequest) SetOuterWidth(v int32)`

SetOuterWidth sets OuterWidth field to given value.

### HasOuterWidth

`func (o *PatchedWritableRackRequest) HasOuterWidth() bool`

HasOuterWidth returns a boolean if a field has been set.

### SetOuterWidthNil

`func (o *PatchedWritableRackRequest) SetOuterWidthNil(b bool)`

 SetOuterWidthNil sets the value for OuterWidth to be an explicit nil

### UnsetOuterWidth
`func (o *PatchedWritableRackRequest) UnsetOuterWidth()`

UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
### GetOuterHeight

`func (o *PatchedWritableRackRequest) GetOuterHeight() int32`

GetOuterHeight returns the OuterHeight field if non-nil, zero value otherwise.

### GetOuterHeightOk

`func (o *PatchedWritableRackRequest) GetOuterHeightOk() (*int32, bool)`

GetOuterHeightOk returns a tuple with the OuterHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterHeight

`func (o *PatchedWritableRackRequest) SetOuterHeight(v int32)`

SetOuterHeight sets OuterHeight field to given value.

### HasOuterHeight

`func (o *PatchedWritableRackRequest) HasOuterHeight() bool`

HasOuterHeight returns a boolean if a field has been set.

### SetOuterHeightNil

`func (o *PatchedWritableRackRequest) SetOuterHeightNil(b bool)`

 SetOuterHeightNil sets the value for OuterHeight to be an explicit nil

### UnsetOuterHeight
`func (o *PatchedWritableRackRequest) UnsetOuterHeight()`

UnsetOuterHeight ensures that no value is present for OuterHeight, not even an explicit nil
### GetOuterDepth

`func (o *PatchedWritableRackRequest) GetOuterDepth() int32`

GetOuterDepth returns the OuterDepth field if non-nil, zero value otherwise.

### GetOuterDepthOk

`func (o *PatchedWritableRackRequest) GetOuterDepthOk() (*int32, bool)`

GetOuterDepthOk returns a tuple with the OuterDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterDepth

`func (o *PatchedWritableRackRequest) SetOuterDepth(v int32)`

SetOuterDepth sets OuterDepth field to given value.

### HasOuterDepth

`func (o *PatchedWritableRackRequest) HasOuterDepth() bool`

HasOuterDepth returns a boolean if a field has been set.

### SetOuterDepthNil

`func (o *PatchedWritableRackRequest) SetOuterDepthNil(b bool)`

 SetOuterDepthNil sets the value for OuterDepth to be an explicit nil

### UnsetOuterDepth
`func (o *PatchedWritableRackRequest) UnsetOuterDepth()`

UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
### GetOuterUnit

`func (o *PatchedWritableRackRequest) GetOuterUnit() BulkRackRequestOuterUnit`

GetOuterUnit returns the OuterUnit field if non-nil, zero value otherwise.

### GetOuterUnitOk

`func (o *PatchedWritableRackRequest) GetOuterUnitOk() (*BulkRackRequestOuterUnit, bool)`

GetOuterUnitOk returns a tuple with the OuterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterUnit

`func (o *PatchedWritableRackRequest) SetOuterUnit(v BulkRackRequestOuterUnit)`

SetOuterUnit sets OuterUnit field to given value.

### HasOuterUnit

`func (o *PatchedWritableRackRequest) HasOuterUnit() bool`

HasOuterUnit returns a boolean if a field has been set.

### SetOuterUnitNil

`func (o *PatchedWritableRackRequest) SetOuterUnitNil(b bool)`

 SetOuterUnitNil sets the value for OuterUnit to be an explicit nil

### UnsetOuterUnit
`func (o *PatchedWritableRackRequest) UnsetOuterUnit()`

UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
### GetMountingDepth

`func (o *PatchedWritableRackRequest) GetMountingDepth() int32`

GetMountingDepth returns the MountingDepth field if non-nil, zero value otherwise.

### GetMountingDepthOk

`func (o *PatchedWritableRackRequest) GetMountingDepthOk() (*int32, bool)`

GetMountingDepthOk returns a tuple with the MountingDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountingDepth

`func (o *PatchedWritableRackRequest) SetMountingDepth(v int32)`

SetMountingDepth sets MountingDepth field to given value.

### HasMountingDepth

`func (o *PatchedWritableRackRequest) HasMountingDepth() bool`

HasMountingDepth returns a boolean if a field has been set.

### SetMountingDepthNil

`func (o *PatchedWritableRackRequest) SetMountingDepthNil(b bool)`

 SetMountingDepthNil sets the value for MountingDepth to be an explicit nil

### UnsetMountingDepth
`func (o *PatchedWritableRackRequest) UnsetMountingDepth()`

UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
### GetAirflow

`func (o *PatchedWritableRackRequest) GetAirflow() PatchedWritableRackRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *PatchedWritableRackRequest) GetAirflowOk() (*PatchedWritableRackRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *PatchedWritableRackRequest) SetAirflow(v PatchedWritableRackRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *PatchedWritableRackRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *PatchedWritableRackRequest) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *PatchedWritableRackRequest) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetCoolingCapability

`func (o *PatchedWritableRackRequest) GetCoolingCapability() BulkRackRequestCoolingCapability`

GetCoolingCapability returns the CoolingCapability field if non-nil, zero value otherwise.

### GetCoolingCapabilityOk

`func (o *PatchedWritableRackRequest) GetCoolingCapabilityOk() (*BulkRackRequestCoolingCapability, bool)`

GetCoolingCapabilityOk returns a tuple with the CoolingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapability

`func (o *PatchedWritableRackRequest) SetCoolingCapability(v BulkRackRequestCoolingCapability)`

SetCoolingCapability sets CoolingCapability field to given value.

### HasCoolingCapability

`func (o *PatchedWritableRackRequest) HasCoolingCapability() bool`

HasCoolingCapability returns a boolean if a field has been set.

### SetCoolingCapabilityNil

`func (o *PatchedWritableRackRequest) SetCoolingCapabilityNil(b bool)`

 SetCoolingCapabilityNil sets the value for CoolingCapability to be an explicit nil

### UnsetCoolingCapability
`func (o *PatchedWritableRackRequest) UnsetCoolingCapability()`

UnsetCoolingCapability ensures that no value is present for CoolingCapability, not even an explicit nil
### GetCoolingCapacity

`func (o *PatchedWritableRackRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *PatchedWritableRackRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *PatchedWritableRackRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *PatchedWritableRackRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *PatchedWritableRackRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *PatchedWritableRackRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetDescription

`func (o *PatchedWritableRackRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableRackRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableRackRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableRackRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedWritableRackRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedWritableRackRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedWritableRackRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedWritableRackRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedWritableRackRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedWritableRackRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedWritableRackRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableRackRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableRackRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableRackRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableRackRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableRackRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableRackRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableRackRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableRackRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableRackRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableRackRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableRackRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


