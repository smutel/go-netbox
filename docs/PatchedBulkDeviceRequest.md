# PatchedBulkDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | Pointer to [**BulkDeviceBayTemplateRequestDeviceType**](BulkDeviceBayTemplateRequestDeviceType.md) |  | [optional] 
**Role** | Pointer to [**BulkDeviceRequestRole**](BulkDeviceRequestRole.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableBulkDeviceRequestPlatform**](BulkDeviceRequestPlatform.md) |  | [optional] 
**Serial** | Pointer to **string** | Chassis serial number, assigned by the manufacturer | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | Pointer to [**BulkCoolingSourceRequestSite**](BulkCoolingSourceRequestSite.md) |  | [optional] 
**Location** | Pointer to [**NullableBulkCoolingSourceRequestLocation**](BulkCoolingSourceRequestLocation.md) |  | [optional] 
**Rack** | Pointer to [**NullableBulkCoolingFeedRequestRack**](BulkCoolingFeedRequestRack.md) |  | [optional] 
**Position** | Pointer to **NullableFloat64** |  | [optional] 
**Face** | Pointer to [**BulkDeviceRequestFace**](BulkDeviceRequestFace.md) |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Status** | Pointer to [**BulkDeviceRequestStatus**](BulkDeviceRequestStatus.md) |  | [optional] 
**Airflow** | Pointer to [**BulkDeviceRequestAirflow**](BulkDeviceRequestAirflow.md) |  | [optional] 
**CoolingMethod** | Pointer to [**NullableBulkDeviceRequestCoolingMethod**](BulkDeviceRequestCoolingMethod.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**OobIp** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**Cluster** | Pointer to [**NullableBulkDeviceRequestCluster**](BulkDeviceRequestCluster.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableBulkDeviceRequestVirtualChassis**](BulkDeviceRequestVirtualChassis.md) |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** | Virtual chassis master election priority | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ConfigTemplate** | Pointer to [**NullableBulkDeviceRequestConfigTemplate**](BulkDeviceRequestConfigTemplate.md) |  | [optional] 
**LocalContextData** | Pointer to **interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkDeviceRequest

`func NewPatchedBulkDeviceRequest(id int32, ) *PatchedBulkDeviceRequest`

NewPatchedBulkDeviceRequest instantiates a new PatchedBulkDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkDeviceRequestWithDefaults

`func NewPatchedBulkDeviceRequestWithDefaults() *PatchedBulkDeviceRequest`

NewPatchedBulkDeviceRequestWithDefaults instantiates a new PatchedBulkDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkDeviceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkDeviceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkDeviceRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedBulkDeviceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedBulkDeviceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *PatchedBulkDeviceRequest) GetDeviceType() BulkDeviceBayTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedBulkDeviceRequest) GetDeviceTypeOk() (*BulkDeviceBayTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedBulkDeviceRequest) SetDeviceType(v BulkDeviceBayTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedBulkDeviceRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkDeviceRequest) GetRole() BulkDeviceRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkDeviceRequest) GetRoleOk() (*BulkDeviceRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkDeviceRequest) SetRole(v BulkDeviceRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkDeviceRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkDeviceRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkDeviceRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkDeviceRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkDeviceRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkDeviceRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkDeviceRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *PatchedBulkDeviceRequest) GetPlatform() BulkDeviceRequestPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedBulkDeviceRequest) GetPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedBulkDeviceRequest) SetPlatform(v BulkDeviceRequestPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedBulkDeviceRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedBulkDeviceRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedBulkDeviceRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *PatchedBulkDeviceRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedBulkDeviceRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedBulkDeviceRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedBulkDeviceRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedBulkDeviceRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedBulkDeviceRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedBulkDeviceRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedBulkDeviceRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedBulkDeviceRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedBulkDeviceRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *PatchedBulkDeviceRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedBulkDeviceRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedBulkDeviceRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedBulkDeviceRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedBulkDeviceRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedBulkDeviceRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedBulkDeviceRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedBulkDeviceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedBulkDeviceRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedBulkDeviceRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *PatchedBulkDeviceRequest) GetRack() BulkCoolingFeedRequestRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedBulkDeviceRequest) GetRackOk() (*BulkCoolingFeedRequestRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedBulkDeviceRequest) SetRack(v BulkCoolingFeedRequestRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedBulkDeviceRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedBulkDeviceRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedBulkDeviceRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *PatchedBulkDeviceRequest) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedBulkDeviceRequest) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedBulkDeviceRequest) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedBulkDeviceRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *PatchedBulkDeviceRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *PatchedBulkDeviceRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *PatchedBulkDeviceRequest) GetFace() BulkDeviceRequestFace`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *PatchedBulkDeviceRequest) GetFaceOk() (*BulkDeviceRequestFace, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *PatchedBulkDeviceRequest) SetFace(v BulkDeviceRequestFace)`

SetFace sets Face field to given value.

### HasFace

`func (o *PatchedBulkDeviceRequest) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetLatitude

`func (o *PatchedBulkDeviceRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PatchedBulkDeviceRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PatchedBulkDeviceRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PatchedBulkDeviceRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *PatchedBulkDeviceRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *PatchedBulkDeviceRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *PatchedBulkDeviceRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PatchedBulkDeviceRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PatchedBulkDeviceRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PatchedBulkDeviceRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *PatchedBulkDeviceRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *PatchedBulkDeviceRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetStatus

`func (o *PatchedBulkDeviceRequest) GetStatus() BulkDeviceRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkDeviceRequest) GetStatusOk() (*BulkDeviceRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkDeviceRequest) SetStatus(v BulkDeviceRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkDeviceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirflow

`func (o *PatchedBulkDeviceRequest) GetAirflow() BulkDeviceRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *PatchedBulkDeviceRequest) GetAirflowOk() (*BulkDeviceRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *PatchedBulkDeviceRequest) SetAirflow(v BulkDeviceRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *PatchedBulkDeviceRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### GetCoolingMethod

`func (o *PatchedBulkDeviceRequest) GetCoolingMethod() BulkDeviceRequestCoolingMethod`

GetCoolingMethod returns the CoolingMethod field if non-nil, zero value otherwise.

### GetCoolingMethodOk

`func (o *PatchedBulkDeviceRequest) GetCoolingMethodOk() (*BulkDeviceRequestCoolingMethod, bool)`

GetCoolingMethodOk returns a tuple with the CoolingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingMethod

`func (o *PatchedBulkDeviceRequest) SetCoolingMethod(v BulkDeviceRequestCoolingMethod)`

SetCoolingMethod sets CoolingMethod field to given value.

### HasCoolingMethod

`func (o *PatchedBulkDeviceRequest) HasCoolingMethod() bool`

HasCoolingMethod returns a boolean if a field has been set.

### SetCoolingMethodNil

`func (o *PatchedBulkDeviceRequest) SetCoolingMethodNil(b bool)`

 SetCoolingMethodNil sets the value for CoolingMethod to be an explicit nil

### UnsetCoolingMethod
`func (o *PatchedBulkDeviceRequest) UnsetCoolingMethod()`

UnsetCoolingMethod ensures that no value is present for CoolingMethod, not even an explicit nil
### GetPrimaryIp4

`func (o *PatchedBulkDeviceRequest) GetPrimaryIp4() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedBulkDeviceRequest) GetPrimaryIp4Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedBulkDeviceRequest) SetPrimaryIp4(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedBulkDeviceRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedBulkDeviceRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedBulkDeviceRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedBulkDeviceRequest) GetPrimaryIp6() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedBulkDeviceRequest) GetPrimaryIp6Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedBulkDeviceRequest) SetPrimaryIp6(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedBulkDeviceRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedBulkDeviceRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedBulkDeviceRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetOobIp

`func (o *PatchedBulkDeviceRequest) GetOobIp() BulkDeviceRequestPrimaryIp4`

GetOobIp returns the OobIp field if non-nil, zero value otherwise.

### GetOobIpOk

`func (o *PatchedBulkDeviceRequest) GetOobIpOk() (*BulkDeviceRequestPrimaryIp4, bool)`

GetOobIpOk returns a tuple with the OobIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIp

`func (o *PatchedBulkDeviceRequest) SetOobIp(v BulkDeviceRequestPrimaryIp4)`

SetOobIp sets OobIp field to given value.

### HasOobIp

`func (o *PatchedBulkDeviceRequest) HasOobIp() bool`

HasOobIp returns a boolean if a field has been set.

### SetOobIpNil

`func (o *PatchedBulkDeviceRequest) SetOobIpNil(b bool)`

 SetOobIpNil sets the value for OobIp to be an explicit nil

### UnsetOobIp
`func (o *PatchedBulkDeviceRequest) UnsetOobIp()`

UnsetOobIp ensures that no value is present for OobIp, not even an explicit nil
### GetCluster

`func (o *PatchedBulkDeviceRequest) GetCluster() BulkDeviceRequestCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedBulkDeviceRequest) GetClusterOk() (*BulkDeviceRequestCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedBulkDeviceRequest) SetCluster(v BulkDeviceRequestCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedBulkDeviceRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *PatchedBulkDeviceRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *PatchedBulkDeviceRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *PatchedBulkDeviceRequest) GetVirtualChassis() BulkDeviceRequestVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *PatchedBulkDeviceRequest) GetVirtualChassisOk() (*BulkDeviceRequestVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *PatchedBulkDeviceRequest) SetVirtualChassis(v BulkDeviceRequestVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *PatchedBulkDeviceRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *PatchedBulkDeviceRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *PatchedBulkDeviceRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *PatchedBulkDeviceRequest) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *PatchedBulkDeviceRequest) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *PatchedBulkDeviceRequest) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *PatchedBulkDeviceRequest) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *PatchedBulkDeviceRequest) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *PatchedBulkDeviceRequest) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *PatchedBulkDeviceRequest) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *PatchedBulkDeviceRequest) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *PatchedBulkDeviceRequest) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *PatchedBulkDeviceRequest) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *PatchedBulkDeviceRequest) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *PatchedBulkDeviceRequest) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetDescription

`func (o *PatchedBulkDeviceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkDeviceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkDeviceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkDeviceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkDeviceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkDeviceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkDeviceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkDeviceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkDeviceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkDeviceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkDeviceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkDeviceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkDeviceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkDeviceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *PatchedBulkDeviceRequest) GetConfigTemplate() BulkDeviceRequestConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *PatchedBulkDeviceRequest) GetConfigTemplateOk() (*BulkDeviceRequestConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *PatchedBulkDeviceRequest) SetConfigTemplate(v BulkDeviceRequestConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *PatchedBulkDeviceRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *PatchedBulkDeviceRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *PatchedBulkDeviceRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetLocalContextData

`func (o *PatchedBulkDeviceRequest) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *PatchedBulkDeviceRequest) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *PatchedBulkDeviceRequest) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *PatchedBulkDeviceRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *PatchedBulkDeviceRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *PatchedBulkDeviceRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *PatchedBulkDeviceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkDeviceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkDeviceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkDeviceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkDeviceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkDeviceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkDeviceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


