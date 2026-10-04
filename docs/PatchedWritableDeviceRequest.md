# PatchedWritableDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Face** | Pointer to [**NullableRackFace1**](RackFace1.md) |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Status** | Pointer to [**BulkDeviceRequestStatus**](BulkDeviceRequestStatus.md) |  | [optional] 
**Airflow** | Pointer to [**NullableBulkDeviceTypeRequestAirflow**](BulkDeviceTypeRequestAirflow.md) |  | [optional] 
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

### NewPatchedWritableDeviceRequest

`func NewPatchedWritableDeviceRequest() *PatchedWritableDeviceRequest`

NewPatchedWritableDeviceRequest instantiates a new PatchedWritableDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableDeviceRequestWithDefaults

`func NewPatchedWritableDeviceRequestWithDefaults() *PatchedWritableDeviceRequest`

NewPatchedWritableDeviceRequestWithDefaults instantiates a new PatchedWritableDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedWritableDeviceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedWritableDeviceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *PatchedWritableDeviceRequest) GetDeviceType() BulkDeviceBayTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableDeviceRequest) GetDeviceTypeOk() (*BulkDeviceBayTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableDeviceRequest) SetDeviceType(v BulkDeviceBayTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableDeviceRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableDeviceRequest) GetRole() BulkDeviceRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableDeviceRequest) GetRoleOk() (*BulkDeviceRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableDeviceRequest) SetRole(v BulkDeviceRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableDeviceRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableDeviceRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableDeviceRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableDeviceRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableDeviceRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableDeviceRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableDeviceRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *PatchedWritableDeviceRequest) GetPlatform() BulkDeviceRequestPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedWritableDeviceRequest) GetPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedWritableDeviceRequest) SetPlatform(v BulkDeviceRequestPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedWritableDeviceRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedWritableDeviceRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedWritableDeviceRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *PatchedWritableDeviceRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableDeviceRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableDeviceRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableDeviceRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedWritableDeviceRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedWritableDeviceRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedWritableDeviceRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedWritableDeviceRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedWritableDeviceRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedWritableDeviceRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *PatchedWritableDeviceRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableDeviceRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableDeviceRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableDeviceRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedWritableDeviceRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWritableDeviceRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWritableDeviceRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWritableDeviceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedWritableDeviceRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedWritableDeviceRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *PatchedWritableDeviceRequest) GetRack() BulkCoolingFeedRequestRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedWritableDeviceRequest) GetRackOk() (*BulkCoolingFeedRequestRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedWritableDeviceRequest) SetRack(v BulkCoolingFeedRequestRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedWritableDeviceRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedWritableDeviceRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedWritableDeviceRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *PatchedWritableDeviceRequest) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedWritableDeviceRequest) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedWritableDeviceRequest) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedWritableDeviceRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *PatchedWritableDeviceRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *PatchedWritableDeviceRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *PatchedWritableDeviceRequest) GetFace() RackFace1`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *PatchedWritableDeviceRequest) GetFaceOk() (*RackFace1, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *PatchedWritableDeviceRequest) SetFace(v RackFace1)`

SetFace sets Face field to given value.

### HasFace

`func (o *PatchedWritableDeviceRequest) HasFace() bool`

HasFace returns a boolean if a field has been set.

### SetFaceNil

`func (o *PatchedWritableDeviceRequest) SetFaceNil(b bool)`

 SetFaceNil sets the value for Face to be an explicit nil

### UnsetFace
`func (o *PatchedWritableDeviceRequest) UnsetFace()`

UnsetFace ensures that no value is present for Face, not even an explicit nil
### GetLatitude

`func (o *PatchedWritableDeviceRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PatchedWritableDeviceRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PatchedWritableDeviceRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PatchedWritableDeviceRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *PatchedWritableDeviceRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *PatchedWritableDeviceRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *PatchedWritableDeviceRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PatchedWritableDeviceRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PatchedWritableDeviceRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PatchedWritableDeviceRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *PatchedWritableDeviceRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *PatchedWritableDeviceRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetStatus

`func (o *PatchedWritableDeviceRequest) GetStatus() BulkDeviceRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableDeviceRequest) GetStatusOk() (*BulkDeviceRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableDeviceRequest) SetStatus(v BulkDeviceRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableDeviceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirflow

`func (o *PatchedWritableDeviceRequest) GetAirflow() BulkDeviceTypeRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *PatchedWritableDeviceRequest) GetAirflowOk() (*BulkDeviceTypeRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *PatchedWritableDeviceRequest) SetAirflow(v BulkDeviceTypeRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *PatchedWritableDeviceRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *PatchedWritableDeviceRequest) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *PatchedWritableDeviceRequest) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetCoolingMethod

`func (o *PatchedWritableDeviceRequest) GetCoolingMethod() BulkDeviceRequestCoolingMethod`

GetCoolingMethod returns the CoolingMethod field if non-nil, zero value otherwise.

### GetCoolingMethodOk

`func (o *PatchedWritableDeviceRequest) GetCoolingMethodOk() (*BulkDeviceRequestCoolingMethod, bool)`

GetCoolingMethodOk returns a tuple with the CoolingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingMethod

`func (o *PatchedWritableDeviceRequest) SetCoolingMethod(v BulkDeviceRequestCoolingMethod)`

SetCoolingMethod sets CoolingMethod field to given value.

### HasCoolingMethod

`func (o *PatchedWritableDeviceRequest) HasCoolingMethod() bool`

HasCoolingMethod returns a boolean if a field has been set.

### SetCoolingMethodNil

`func (o *PatchedWritableDeviceRequest) SetCoolingMethodNil(b bool)`

 SetCoolingMethodNil sets the value for CoolingMethod to be an explicit nil

### UnsetCoolingMethod
`func (o *PatchedWritableDeviceRequest) UnsetCoolingMethod()`

UnsetCoolingMethod ensures that no value is present for CoolingMethod, not even an explicit nil
### GetPrimaryIp4

`func (o *PatchedWritableDeviceRequest) GetPrimaryIp4() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedWritableDeviceRequest) GetPrimaryIp4Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedWritableDeviceRequest) SetPrimaryIp4(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedWritableDeviceRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedWritableDeviceRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedWritableDeviceRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedWritableDeviceRequest) GetPrimaryIp6() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedWritableDeviceRequest) GetPrimaryIp6Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedWritableDeviceRequest) SetPrimaryIp6(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedWritableDeviceRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedWritableDeviceRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedWritableDeviceRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetOobIp

`func (o *PatchedWritableDeviceRequest) GetOobIp() BulkDeviceRequestPrimaryIp4`

GetOobIp returns the OobIp field if non-nil, zero value otherwise.

### GetOobIpOk

`func (o *PatchedWritableDeviceRequest) GetOobIpOk() (*BulkDeviceRequestPrimaryIp4, bool)`

GetOobIpOk returns a tuple with the OobIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIp

`func (o *PatchedWritableDeviceRequest) SetOobIp(v BulkDeviceRequestPrimaryIp4)`

SetOobIp sets OobIp field to given value.

### HasOobIp

`func (o *PatchedWritableDeviceRequest) HasOobIp() bool`

HasOobIp returns a boolean if a field has been set.

### SetOobIpNil

`func (o *PatchedWritableDeviceRequest) SetOobIpNil(b bool)`

 SetOobIpNil sets the value for OobIp to be an explicit nil

### UnsetOobIp
`func (o *PatchedWritableDeviceRequest) UnsetOobIp()`

UnsetOobIp ensures that no value is present for OobIp, not even an explicit nil
### GetCluster

`func (o *PatchedWritableDeviceRequest) GetCluster() BulkDeviceRequestCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedWritableDeviceRequest) GetClusterOk() (*BulkDeviceRequestCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedWritableDeviceRequest) SetCluster(v BulkDeviceRequestCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedWritableDeviceRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *PatchedWritableDeviceRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *PatchedWritableDeviceRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *PatchedWritableDeviceRequest) GetVirtualChassis() BulkDeviceRequestVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *PatchedWritableDeviceRequest) GetVirtualChassisOk() (*BulkDeviceRequestVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *PatchedWritableDeviceRequest) SetVirtualChassis(v BulkDeviceRequestVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *PatchedWritableDeviceRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *PatchedWritableDeviceRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *PatchedWritableDeviceRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *PatchedWritableDeviceRequest) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *PatchedWritableDeviceRequest) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *PatchedWritableDeviceRequest) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *PatchedWritableDeviceRequest) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *PatchedWritableDeviceRequest) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *PatchedWritableDeviceRequest) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *PatchedWritableDeviceRequest) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *PatchedWritableDeviceRequest) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *PatchedWritableDeviceRequest) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *PatchedWritableDeviceRequest) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *PatchedWritableDeviceRequest) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *PatchedWritableDeviceRequest) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetDescription

`func (o *PatchedWritableDeviceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableDeviceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableDeviceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableDeviceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedWritableDeviceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedWritableDeviceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedWritableDeviceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedWritableDeviceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedWritableDeviceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedWritableDeviceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedWritableDeviceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableDeviceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableDeviceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableDeviceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *PatchedWritableDeviceRequest) GetConfigTemplate() BulkDeviceRequestConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *PatchedWritableDeviceRequest) GetConfigTemplateOk() (*BulkDeviceRequestConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *PatchedWritableDeviceRequest) SetConfigTemplate(v BulkDeviceRequestConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *PatchedWritableDeviceRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *PatchedWritableDeviceRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *PatchedWritableDeviceRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetLocalContextData

`func (o *PatchedWritableDeviceRequest) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *PatchedWritableDeviceRequest) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *PatchedWritableDeviceRequest) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *PatchedWritableDeviceRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *PatchedWritableDeviceRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *PatchedWritableDeviceRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *PatchedWritableDeviceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableDeviceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableDeviceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableDeviceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableDeviceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableDeviceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableDeviceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


