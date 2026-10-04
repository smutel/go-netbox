# WritableDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | [**BulkDeviceBayTemplateRequestDeviceType**](BulkDeviceBayTemplateRequestDeviceType.md) |  | 
**Role** | [**BulkDeviceRequestRole**](BulkDeviceRequestRole.md) |  | 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableBulkDeviceRequestPlatform**](BulkDeviceRequestPlatform.md) |  | [optional] 
**Serial** | Pointer to **string** | Chassis serial number, assigned by the manufacturer | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | [**BulkCoolingSourceRequestSite**](BulkCoolingSourceRequestSite.md) |  | 
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

### NewWritableDeviceRequest

`func NewWritableDeviceRequest(deviceType BulkDeviceBayTemplateRequestDeviceType, role BulkDeviceRequestRole, site BulkCoolingSourceRequestSite, ) *WritableDeviceRequest`

NewWritableDeviceRequest instantiates a new WritableDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableDeviceRequestWithDefaults

`func NewWritableDeviceRequestWithDefaults() *WritableDeviceRequest`

NewWritableDeviceRequestWithDefaults instantiates a new WritableDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WritableDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WritableDeviceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WritableDeviceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *WritableDeviceRequest) GetDeviceType() BulkDeviceBayTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WritableDeviceRequest) GetDeviceTypeOk() (*BulkDeviceBayTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WritableDeviceRequest) SetDeviceType(v BulkDeviceBayTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetRole

`func (o *WritableDeviceRequest) GetRole() BulkDeviceRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritableDeviceRequest) GetRoleOk() (*BulkDeviceRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritableDeviceRequest) SetRole(v BulkDeviceRequestRole)`

SetRole sets Role field to given value.


### GetTenant

`func (o *WritableDeviceRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableDeviceRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableDeviceRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableDeviceRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableDeviceRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableDeviceRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *WritableDeviceRequest) GetPlatform() BulkDeviceRequestPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WritableDeviceRequest) GetPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WritableDeviceRequest) SetPlatform(v BulkDeviceRequestPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *WritableDeviceRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *WritableDeviceRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *WritableDeviceRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *WritableDeviceRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *WritableDeviceRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *WritableDeviceRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *WritableDeviceRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *WritableDeviceRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *WritableDeviceRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *WritableDeviceRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *WritableDeviceRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *WritableDeviceRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *WritableDeviceRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *WritableDeviceRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritableDeviceRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritableDeviceRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *WritableDeviceRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WritableDeviceRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WritableDeviceRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WritableDeviceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *WritableDeviceRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *WritableDeviceRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *WritableDeviceRequest) GetRack() BulkCoolingFeedRequestRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *WritableDeviceRequest) GetRackOk() (*BulkCoolingFeedRequestRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *WritableDeviceRequest) SetRack(v BulkCoolingFeedRequestRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *WritableDeviceRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *WritableDeviceRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *WritableDeviceRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *WritableDeviceRequest) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WritableDeviceRequest) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WritableDeviceRequest) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WritableDeviceRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *WritableDeviceRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *WritableDeviceRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *WritableDeviceRequest) GetFace() RackFace1`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *WritableDeviceRequest) GetFaceOk() (*RackFace1, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *WritableDeviceRequest) SetFace(v RackFace1)`

SetFace sets Face field to given value.

### HasFace

`func (o *WritableDeviceRequest) HasFace() bool`

HasFace returns a boolean if a field has been set.

### SetFaceNil

`func (o *WritableDeviceRequest) SetFaceNil(b bool)`

 SetFaceNil sets the value for Face to be an explicit nil

### UnsetFace
`func (o *WritableDeviceRequest) UnsetFace()`

UnsetFace ensures that no value is present for Face, not even an explicit nil
### GetLatitude

`func (o *WritableDeviceRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *WritableDeviceRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *WritableDeviceRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *WritableDeviceRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *WritableDeviceRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *WritableDeviceRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *WritableDeviceRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *WritableDeviceRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *WritableDeviceRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *WritableDeviceRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *WritableDeviceRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *WritableDeviceRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetStatus

`func (o *WritableDeviceRequest) GetStatus() BulkDeviceRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableDeviceRequest) GetStatusOk() (*BulkDeviceRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableDeviceRequest) SetStatus(v BulkDeviceRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableDeviceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirflow

`func (o *WritableDeviceRequest) GetAirflow() BulkDeviceTypeRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *WritableDeviceRequest) GetAirflowOk() (*BulkDeviceTypeRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *WritableDeviceRequest) SetAirflow(v BulkDeviceTypeRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *WritableDeviceRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *WritableDeviceRequest) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *WritableDeviceRequest) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetCoolingMethod

`func (o *WritableDeviceRequest) GetCoolingMethod() BulkDeviceRequestCoolingMethod`

GetCoolingMethod returns the CoolingMethod field if non-nil, zero value otherwise.

### GetCoolingMethodOk

`func (o *WritableDeviceRequest) GetCoolingMethodOk() (*BulkDeviceRequestCoolingMethod, bool)`

GetCoolingMethodOk returns a tuple with the CoolingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingMethod

`func (o *WritableDeviceRequest) SetCoolingMethod(v BulkDeviceRequestCoolingMethod)`

SetCoolingMethod sets CoolingMethod field to given value.

### HasCoolingMethod

`func (o *WritableDeviceRequest) HasCoolingMethod() bool`

HasCoolingMethod returns a boolean if a field has been set.

### SetCoolingMethodNil

`func (o *WritableDeviceRequest) SetCoolingMethodNil(b bool)`

 SetCoolingMethodNil sets the value for CoolingMethod to be an explicit nil

### UnsetCoolingMethod
`func (o *WritableDeviceRequest) UnsetCoolingMethod()`

UnsetCoolingMethod ensures that no value is present for CoolingMethod, not even an explicit nil
### GetPrimaryIp4

`func (o *WritableDeviceRequest) GetPrimaryIp4() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *WritableDeviceRequest) GetPrimaryIp4Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *WritableDeviceRequest) SetPrimaryIp4(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *WritableDeviceRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *WritableDeviceRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *WritableDeviceRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *WritableDeviceRequest) GetPrimaryIp6() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *WritableDeviceRequest) GetPrimaryIp6Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *WritableDeviceRequest) SetPrimaryIp6(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *WritableDeviceRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *WritableDeviceRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *WritableDeviceRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetOobIp

`func (o *WritableDeviceRequest) GetOobIp() BulkDeviceRequestPrimaryIp4`

GetOobIp returns the OobIp field if non-nil, zero value otherwise.

### GetOobIpOk

`func (o *WritableDeviceRequest) GetOobIpOk() (*BulkDeviceRequestPrimaryIp4, bool)`

GetOobIpOk returns a tuple with the OobIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIp

`func (o *WritableDeviceRequest) SetOobIp(v BulkDeviceRequestPrimaryIp4)`

SetOobIp sets OobIp field to given value.

### HasOobIp

`func (o *WritableDeviceRequest) HasOobIp() bool`

HasOobIp returns a boolean if a field has been set.

### SetOobIpNil

`func (o *WritableDeviceRequest) SetOobIpNil(b bool)`

 SetOobIpNil sets the value for OobIp to be an explicit nil

### UnsetOobIp
`func (o *WritableDeviceRequest) UnsetOobIp()`

UnsetOobIp ensures that no value is present for OobIp, not even an explicit nil
### GetCluster

`func (o *WritableDeviceRequest) GetCluster() BulkDeviceRequestCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *WritableDeviceRequest) GetClusterOk() (*BulkDeviceRequestCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *WritableDeviceRequest) SetCluster(v BulkDeviceRequestCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *WritableDeviceRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *WritableDeviceRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *WritableDeviceRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *WritableDeviceRequest) GetVirtualChassis() BulkDeviceRequestVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *WritableDeviceRequest) GetVirtualChassisOk() (*BulkDeviceRequestVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *WritableDeviceRequest) SetVirtualChassis(v BulkDeviceRequestVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *WritableDeviceRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *WritableDeviceRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *WritableDeviceRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *WritableDeviceRequest) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *WritableDeviceRequest) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *WritableDeviceRequest) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *WritableDeviceRequest) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *WritableDeviceRequest) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *WritableDeviceRequest) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *WritableDeviceRequest) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *WritableDeviceRequest) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *WritableDeviceRequest) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *WritableDeviceRequest) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *WritableDeviceRequest) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *WritableDeviceRequest) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetDescription

`func (o *WritableDeviceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableDeviceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableDeviceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableDeviceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *WritableDeviceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WritableDeviceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WritableDeviceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WritableDeviceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WritableDeviceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WritableDeviceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *WritableDeviceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableDeviceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableDeviceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableDeviceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *WritableDeviceRequest) GetConfigTemplate() BulkDeviceRequestConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *WritableDeviceRequest) GetConfigTemplateOk() (*BulkDeviceRequestConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *WritableDeviceRequest) SetConfigTemplate(v BulkDeviceRequestConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *WritableDeviceRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *WritableDeviceRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *WritableDeviceRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetLocalContextData

`func (o *WritableDeviceRequest) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *WritableDeviceRequest) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *WritableDeviceRequest) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *WritableDeviceRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *WritableDeviceRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *WritableDeviceRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *WritableDeviceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableDeviceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableDeviceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableDeviceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableDeviceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableDeviceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableDeviceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


