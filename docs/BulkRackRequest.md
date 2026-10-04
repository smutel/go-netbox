# BulkRackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**FacilityId** | Pointer to **NullableString** |  | [optional] 
**Site** | [**BulkCoolingSourceRequestSite**](BulkCoolingSourceRequestSite.md) |  | 
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

### NewBulkRackRequest

`func NewBulkRackRequest(id int32, name string, site BulkCoolingSourceRequestSite, ) *BulkRackRequest`

NewBulkRackRequest instantiates a new BulkRackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRackRequestWithDefaults

`func NewBulkRackRequestWithDefaults() *BulkRackRequest`

NewBulkRackRequestWithDefaults instantiates a new BulkRackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkRackRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkRackRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkRackRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkRackRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkRackRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkRackRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFacilityId

`func (o *BulkRackRequest) GetFacilityId() string`

GetFacilityId returns the FacilityId field if non-nil, zero value otherwise.

### GetFacilityIdOk

`func (o *BulkRackRequest) GetFacilityIdOk() (*string, bool)`

GetFacilityIdOk returns a tuple with the FacilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityId

`func (o *BulkRackRequest) SetFacilityId(v string)`

SetFacilityId sets FacilityId field to given value.

### HasFacilityId

`func (o *BulkRackRequest) HasFacilityId() bool`

HasFacilityId returns a boolean if a field has been set.

### SetFacilityIdNil

`func (o *BulkRackRequest) SetFacilityIdNil(b bool)`

 SetFacilityIdNil sets the value for FacilityId to be an explicit nil

### UnsetFacilityId
`func (o *BulkRackRequest) UnsetFacilityId()`

UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
### GetSite

`func (o *BulkRackRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *BulkRackRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *BulkRackRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *BulkRackRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BulkRackRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BulkRackRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BulkRackRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *BulkRackRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *BulkRackRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetGroup

`func (o *BulkRackRequest) GetGroup() BulkRackRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkRackRequest) GetGroupOk() (*BulkRackRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkRackRequest) SetGroup(v BulkRackRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BulkRackRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BulkRackRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BulkRackRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetTenant

`func (o *BulkRackRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkRackRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkRackRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkRackRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkRackRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkRackRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *BulkRackRequest) GetStatus() BulkRackRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkRackRequest) GetStatusOk() (*BulkRackRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkRackRequest) SetStatus(v BulkRackRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkRackRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *BulkRackRequest) GetRole() BulkRackRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkRackRequest) GetRoleOk() (*BulkRackRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkRackRequest) SetRole(v BulkRackRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkRackRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkRackRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkRackRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSerial

`func (o *BulkRackRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *BulkRackRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *BulkRackRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *BulkRackRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *BulkRackRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *BulkRackRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *BulkRackRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *BulkRackRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *BulkRackRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *BulkRackRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetRackType

`func (o *BulkRackRequest) GetRackType() BulkRackRequestRackType`

GetRackType returns the RackType field if non-nil, zero value otherwise.

### GetRackTypeOk

`func (o *BulkRackRequest) GetRackTypeOk() (*BulkRackRequestRackType, bool)`

GetRackTypeOk returns a tuple with the RackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackType

`func (o *BulkRackRequest) SetRackType(v BulkRackRequestRackType)`

SetRackType sets RackType field to given value.

### HasRackType

`func (o *BulkRackRequest) HasRackType() bool`

HasRackType returns a boolean if a field has been set.

### SetRackTypeNil

`func (o *BulkRackRequest) SetRackTypeNil(b bool)`

 SetRackTypeNil sets the value for RackType to be an explicit nil

### UnsetRackType
`func (o *BulkRackRequest) UnsetRackType()`

UnsetRackType ensures that no value is present for RackType, not even an explicit nil
### GetFormFactor

`func (o *BulkRackRequest) GetFormFactor() BulkRackRequestFormFactor`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *BulkRackRequest) GetFormFactorOk() (*BulkRackRequestFormFactor, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *BulkRackRequest) SetFormFactor(v BulkRackRequestFormFactor)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *BulkRackRequest) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### SetFormFactorNil

`func (o *BulkRackRequest) SetFormFactorNil(b bool)`

 SetFormFactorNil sets the value for FormFactor to be an explicit nil

### UnsetFormFactor
`func (o *BulkRackRequest) UnsetFormFactor()`

UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
### GetWidth

`func (o *BulkRackRequest) GetWidth() BulkRackRequestWidth`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BulkRackRequest) GetWidthOk() (*BulkRackRequestWidth, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BulkRackRequest) SetWidth(v BulkRackRequestWidth)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *BulkRackRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetUHeight

`func (o *BulkRackRequest) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *BulkRackRequest) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *BulkRackRequest) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *BulkRackRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetStartingUnit

`func (o *BulkRackRequest) GetStartingUnit() int32`

GetStartingUnit returns the StartingUnit field if non-nil, zero value otherwise.

### GetStartingUnitOk

`func (o *BulkRackRequest) GetStartingUnitOk() (*int32, bool)`

GetStartingUnitOk returns a tuple with the StartingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingUnit

`func (o *BulkRackRequest) SetStartingUnit(v int32)`

SetStartingUnit sets StartingUnit field to given value.

### HasStartingUnit

`func (o *BulkRackRequest) HasStartingUnit() bool`

HasStartingUnit returns a boolean if a field has been set.

### GetWeight

`func (o *BulkRackRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkRackRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkRackRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkRackRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *BulkRackRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *BulkRackRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMaxWeight

`func (o *BulkRackRequest) GetMaxWeight() int32`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *BulkRackRequest) GetMaxWeightOk() (*int32, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *BulkRackRequest) SetMaxWeight(v int32)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *BulkRackRequest) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### SetMaxWeightNil

`func (o *BulkRackRequest) SetMaxWeightNil(b bool)`

 SetMaxWeightNil sets the value for MaxWeight to be an explicit nil

### UnsetMaxWeight
`func (o *BulkRackRequest) UnsetMaxWeight()`

UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
### GetWeightUnit

`func (o *BulkRackRequest) GetWeightUnit() BulkDeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *BulkRackRequest) GetWeightUnitOk() (*BulkDeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *BulkRackRequest) SetWeightUnit(v BulkDeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *BulkRackRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *BulkRackRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *BulkRackRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetDescUnits

`func (o *BulkRackRequest) GetDescUnits() bool`

GetDescUnits returns the DescUnits field if non-nil, zero value otherwise.

### GetDescUnitsOk

`func (o *BulkRackRequest) GetDescUnitsOk() (*bool, bool)`

GetDescUnitsOk returns a tuple with the DescUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescUnits

`func (o *BulkRackRequest) SetDescUnits(v bool)`

SetDescUnits sets DescUnits field to given value.

### HasDescUnits

`func (o *BulkRackRequest) HasDescUnits() bool`

HasDescUnits returns a boolean if a field has been set.

### GetOuterWidth

`func (o *BulkRackRequest) GetOuterWidth() int32`

GetOuterWidth returns the OuterWidth field if non-nil, zero value otherwise.

### GetOuterWidthOk

`func (o *BulkRackRequest) GetOuterWidthOk() (*int32, bool)`

GetOuterWidthOk returns a tuple with the OuterWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterWidth

`func (o *BulkRackRequest) SetOuterWidth(v int32)`

SetOuterWidth sets OuterWidth field to given value.

### HasOuterWidth

`func (o *BulkRackRequest) HasOuterWidth() bool`

HasOuterWidth returns a boolean if a field has been set.

### SetOuterWidthNil

`func (o *BulkRackRequest) SetOuterWidthNil(b bool)`

 SetOuterWidthNil sets the value for OuterWidth to be an explicit nil

### UnsetOuterWidth
`func (o *BulkRackRequest) UnsetOuterWidth()`

UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
### GetOuterHeight

`func (o *BulkRackRequest) GetOuterHeight() int32`

GetOuterHeight returns the OuterHeight field if non-nil, zero value otherwise.

### GetOuterHeightOk

`func (o *BulkRackRequest) GetOuterHeightOk() (*int32, bool)`

GetOuterHeightOk returns a tuple with the OuterHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterHeight

`func (o *BulkRackRequest) SetOuterHeight(v int32)`

SetOuterHeight sets OuterHeight field to given value.

### HasOuterHeight

`func (o *BulkRackRequest) HasOuterHeight() bool`

HasOuterHeight returns a boolean if a field has been set.

### SetOuterHeightNil

`func (o *BulkRackRequest) SetOuterHeightNil(b bool)`

 SetOuterHeightNil sets the value for OuterHeight to be an explicit nil

### UnsetOuterHeight
`func (o *BulkRackRequest) UnsetOuterHeight()`

UnsetOuterHeight ensures that no value is present for OuterHeight, not even an explicit nil
### GetOuterDepth

`func (o *BulkRackRequest) GetOuterDepth() int32`

GetOuterDepth returns the OuterDepth field if non-nil, zero value otherwise.

### GetOuterDepthOk

`func (o *BulkRackRequest) GetOuterDepthOk() (*int32, bool)`

GetOuterDepthOk returns a tuple with the OuterDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterDepth

`func (o *BulkRackRequest) SetOuterDepth(v int32)`

SetOuterDepth sets OuterDepth field to given value.

### HasOuterDepth

`func (o *BulkRackRequest) HasOuterDepth() bool`

HasOuterDepth returns a boolean if a field has been set.

### SetOuterDepthNil

`func (o *BulkRackRequest) SetOuterDepthNil(b bool)`

 SetOuterDepthNil sets the value for OuterDepth to be an explicit nil

### UnsetOuterDepth
`func (o *BulkRackRequest) UnsetOuterDepth()`

UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
### GetOuterUnit

`func (o *BulkRackRequest) GetOuterUnit() BulkRackRequestOuterUnit`

GetOuterUnit returns the OuterUnit field if non-nil, zero value otherwise.

### GetOuterUnitOk

`func (o *BulkRackRequest) GetOuterUnitOk() (*BulkRackRequestOuterUnit, bool)`

GetOuterUnitOk returns a tuple with the OuterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterUnit

`func (o *BulkRackRequest) SetOuterUnit(v BulkRackRequestOuterUnit)`

SetOuterUnit sets OuterUnit field to given value.

### HasOuterUnit

`func (o *BulkRackRequest) HasOuterUnit() bool`

HasOuterUnit returns a boolean if a field has been set.

### SetOuterUnitNil

`func (o *BulkRackRequest) SetOuterUnitNil(b bool)`

 SetOuterUnitNil sets the value for OuterUnit to be an explicit nil

### UnsetOuterUnit
`func (o *BulkRackRequest) UnsetOuterUnit()`

UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
### GetMountingDepth

`func (o *BulkRackRequest) GetMountingDepth() int32`

GetMountingDepth returns the MountingDepth field if non-nil, zero value otherwise.

### GetMountingDepthOk

`func (o *BulkRackRequest) GetMountingDepthOk() (*int32, bool)`

GetMountingDepthOk returns a tuple with the MountingDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountingDepth

`func (o *BulkRackRequest) SetMountingDepth(v int32)`

SetMountingDepth sets MountingDepth field to given value.

### HasMountingDepth

`func (o *BulkRackRequest) HasMountingDepth() bool`

HasMountingDepth returns a boolean if a field has been set.

### SetMountingDepthNil

`func (o *BulkRackRequest) SetMountingDepthNil(b bool)`

 SetMountingDepthNil sets the value for MountingDepth to be an explicit nil

### UnsetMountingDepth
`func (o *BulkRackRequest) UnsetMountingDepth()`

UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
### GetAirflow

`func (o *BulkRackRequest) GetAirflow() BulkRackRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *BulkRackRequest) GetAirflowOk() (*BulkRackRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *BulkRackRequest) SetAirflow(v BulkRackRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *BulkRackRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### GetCoolingCapability

`func (o *BulkRackRequest) GetCoolingCapability() BulkRackRequestCoolingCapability`

GetCoolingCapability returns the CoolingCapability field if non-nil, zero value otherwise.

### GetCoolingCapabilityOk

`func (o *BulkRackRequest) GetCoolingCapabilityOk() (*BulkRackRequestCoolingCapability, bool)`

GetCoolingCapabilityOk returns a tuple with the CoolingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapability

`func (o *BulkRackRequest) SetCoolingCapability(v BulkRackRequestCoolingCapability)`

SetCoolingCapability sets CoolingCapability field to given value.

### HasCoolingCapability

`func (o *BulkRackRequest) HasCoolingCapability() bool`

HasCoolingCapability returns a boolean if a field has been set.

### SetCoolingCapabilityNil

`func (o *BulkRackRequest) SetCoolingCapabilityNil(b bool)`

 SetCoolingCapabilityNil sets the value for CoolingCapability to be an explicit nil

### UnsetCoolingCapability
`func (o *BulkRackRequest) UnsetCoolingCapability()`

UnsetCoolingCapability ensures that no value is present for CoolingCapability, not even an explicit nil
### GetCoolingCapacity

`func (o *BulkRackRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *BulkRackRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *BulkRackRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *BulkRackRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *BulkRackRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *BulkRackRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetDescription

`func (o *BulkRackRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkRackRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkRackRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkRackRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkRackRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkRackRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkRackRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkRackRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkRackRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkRackRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkRackRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkRackRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkRackRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkRackRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkRackRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkRackRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkRackRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkRackRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkRackRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkRackRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkRackRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkRackRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


