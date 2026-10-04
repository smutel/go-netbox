# BulkVirtualMachineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**VirtualMachineType** | Pointer to [**NullableBulkVirtualMachineRequestVirtualMachineType**](BulkVirtualMachineRequestVirtualMachineType.md) |  | [optional] 
**Role** | Pointer to [**NullableBulkVirtualMachineRequestRole**](BulkVirtualMachineRequestRole.md) |  | [optional] 
**Status** | Pointer to [**BulkVirtualMachineRequestStatus**](BulkVirtualMachineRequestStatus.md) |  | [optional] 
**StartOnBoot** | Pointer to [**BulkVirtualMachineRequestStartOnBoot**](BulkVirtualMachineRequestStartOnBoot.md) |  | [optional] 
**Site** | Pointer to [**NullableBulkVLANRequestSite**](BulkVLANRequestSite.md) |  | [optional] 
**Cluster** | Pointer to [**NullableBulkDeviceRequestCluster**](BulkDeviceRequestCluster.md) |  | [optional] 
**Device** | Pointer to [**NullableBulkDeviceBayRequestInstalledDevice**](BulkDeviceBayRequestInstalledDevice.md) |  | [optional] 
**Platform** | Pointer to [**NullableBulkDeviceRequestPlatform**](BulkDeviceRequestPlatform.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**Vcpus** | Pointer to **NullableFloat64** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Disk** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**LocalContextData** | Pointer to **interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**ConfigTemplate** | Pointer to [**NullableBulkDeviceRequestConfigTemplate**](BulkDeviceRequestConfigTemplate.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkVirtualMachineRequest

`func NewBulkVirtualMachineRequest(id int32, name string, ) *BulkVirtualMachineRequest`

NewBulkVirtualMachineRequest instantiates a new BulkVirtualMachineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVirtualMachineRequestWithDefaults

`func NewBulkVirtualMachineRequestWithDefaults() *BulkVirtualMachineRequest`

NewBulkVirtualMachineRequestWithDefaults instantiates a new BulkVirtualMachineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkVirtualMachineRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkVirtualMachineRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkVirtualMachineRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkVirtualMachineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkVirtualMachineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkVirtualMachineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVirtualMachineType

`func (o *BulkVirtualMachineRequest) GetVirtualMachineType() BulkVirtualMachineRequestVirtualMachineType`

GetVirtualMachineType returns the VirtualMachineType field if non-nil, zero value otherwise.

### GetVirtualMachineTypeOk

`func (o *BulkVirtualMachineRequest) GetVirtualMachineTypeOk() (*BulkVirtualMachineRequestVirtualMachineType, bool)`

GetVirtualMachineTypeOk returns a tuple with the VirtualMachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineType

`func (o *BulkVirtualMachineRequest) SetVirtualMachineType(v BulkVirtualMachineRequestVirtualMachineType)`

SetVirtualMachineType sets VirtualMachineType field to given value.

### HasVirtualMachineType

`func (o *BulkVirtualMachineRequest) HasVirtualMachineType() bool`

HasVirtualMachineType returns a boolean if a field has been set.

### SetVirtualMachineTypeNil

`func (o *BulkVirtualMachineRequest) SetVirtualMachineTypeNil(b bool)`

 SetVirtualMachineTypeNil sets the value for VirtualMachineType to be an explicit nil

### UnsetVirtualMachineType
`func (o *BulkVirtualMachineRequest) UnsetVirtualMachineType()`

UnsetVirtualMachineType ensures that no value is present for VirtualMachineType, not even an explicit nil
### GetRole

`func (o *BulkVirtualMachineRequest) GetRole() BulkVirtualMachineRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkVirtualMachineRequest) GetRoleOk() (*BulkVirtualMachineRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkVirtualMachineRequest) SetRole(v BulkVirtualMachineRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkVirtualMachineRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkVirtualMachineRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkVirtualMachineRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetStatus

`func (o *BulkVirtualMachineRequest) GetStatus() BulkVirtualMachineRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkVirtualMachineRequest) GetStatusOk() (*BulkVirtualMachineRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkVirtualMachineRequest) SetStatus(v BulkVirtualMachineRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkVirtualMachineRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartOnBoot

`func (o *BulkVirtualMachineRequest) GetStartOnBoot() BulkVirtualMachineRequestStartOnBoot`

GetStartOnBoot returns the StartOnBoot field if non-nil, zero value otherwise.

### GetStartOnBootOk

`func (o *BulkVirtualMachineRequest) GetStartOnBootOk() (*BulkVirtualMachineRequestStartOnBoot, bool)`

GetStartOnBootOk returns a tuple with the StartOnBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOnBoot

`func (o *BulkVirtualMachineRequest) SetStartOnBoot(v BulkVirtualMachineRequestStartOnBoot)`

SetStartOnBoot sets StartOnBoot field to given value.

### HasStartOnBoot

`func (o *BulkVirtualMachineRequest) HasStartOnBoot() bool`

HasStartOnBoot returns a boolean if a field has been set.

### GetSite

`func (o *BulkVirtualMachineRequest) GetSite() BulkVLANRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *BulkVirtualMachineRequest) GetSiteOk() (*BulkVLANRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *BulkVirtualMachineRequest) SetSite(v BulkVLANRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *BulkVirtualMachineRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *BulkVirtualMachineRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *BulkVirtualMachineRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetCluster

`func (o *BulkVirtualMachineRequest) GetCluster() BulkDeviceRequestCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *BulkVirtualMachineRequest) GetClusterOk() (*BulkDeviceRequestCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *BulkVirtualMachineRequest) SetCluster(v BulkDeviceRequestCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *BulkVirtualMachineRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *BulkVirtualMachineRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *BulkVirtualMachineRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetDevice

`func (o *BulkVirtualMachineRequest) GetDevice() BulkDeviceBayRequestInstalledDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkVirtualMachineRequest) GetDeviceOk() (*BulkDeviceBayRequestInstalledDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkVirtualMachineRequest) SetDevice(v BulkDeviceBayRequestInstalledDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BulkVirtualMachineRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *BulkVirtualMachineRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *BulkVirtualMachineRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetPlatform

`func (o *BulkVirtualMachineRequest) GetPlatform() BulkDeviceRequestPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BulkVirtualMachineRequest) GetPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BulkVirtualMachineRequest) SetPlatform(v BulkDeviceRequestPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *BulkVirtualMachineRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *BulkVirtualMachineRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *BulkVirtualMachineRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPrimaryIp4

`func (o *BulkVirtualMachineRequest) GetPrimaryIp4() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *BulkVirtualMachineRequest) GetPrimaryIp4Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *BulkVirtualMachineRequest) SetPrimaryIp4(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *BulkVirtualMachineRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *BulkVirtualMachineRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *BulkVirtualMachineRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *BulkVirtualMachineRequest) GetPrimaryIp6() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *BulkVirtualMachineRequest) GetPrimaryIp6Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *BulkVirtualMachineRequest) SetPrimaryIp6(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *BulkVirtualMachineRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *BulkVirtualMachineRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *BulkVirtualMachineRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetVcpus

`func (o *BulkVirtualMachineRequest) GetVcpus() float64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *BulkVirtualMachineRequest) GetVcpusOk() (*float64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *BulkVirtualMachineRequest) SetVcpus(v float64)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *BulkVirtualMachineRequest) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *BulkVirtualMachineRequest) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *BulkVirtualMachineRequest) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *BulkVirtualMachineRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *BulkVirtualMachineRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *BulkVirtualMachineRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *BulkVirtualMachineRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *BulkVirtualMachineRequest) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *BulkVirtualMachineRequest) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *BulkVirtualMachineRequest) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *BulkVirtualMachineRequest) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *BulkVirtualMachineRequest) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *BulkVirtualMachineRequest) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *BulkVirtualMachineRequest) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *BulkVirtualMachineRequest) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetDescription

`func (o *BulkVirtualMachineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkVirtualMachineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkVirtualMachineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkVirtualMachineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSerial

`func (o *BulkVirtualMachineRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *BulkVirtualMachineRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *BulkVirtualMachineRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *BulkVirtualMachineRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTenant

`func (o *BulkVirtualMachineRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkVirtualMachineRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkVirtualMachineRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkVirtualMachineRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkVirtualMachineRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkVirtualMachineRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetOwner

`func (o *BulkVirtualMachineRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkVirtualMachineRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkVirtualMachineRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkVirtualMachineRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkVirtualMachineRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkVirtualMachineRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkVirtualMachineRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkVirtualMachineRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkVirtualMachineRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkVirtualMachineRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkVirtualMachineRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkVirtualMachineRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkVirtualMachineRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkVirtualMachineRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLocalContextData

`func (o *BulkVirtualMachineRequest) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *BulkVirtualMachineRequest) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *BulkVirtualMachineRequest) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *BulkVirtualMachineRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *BulkVirtualMachineRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *BulkVirtualMachineRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetConfigTemplate

`func (o *BulkVirtualMachineRequest) GetConfigTemplate() BulkDeviceRequestConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *BulkVirtualMachineRequest) GetConfigTemplateOk() (*BulkDeviceRequestConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *BulkVirtualMachineRequest) SetConfigTemplate(v BulkDeviceRequestConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *BulkVirtualMachineRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *BulkVirtualMachineRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *BulkVirtualMachineRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetCustomFields

`func (o *BulkVirtualMachineRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkVirtualMachineRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkVirtualMachineRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkVirtualMachineRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


