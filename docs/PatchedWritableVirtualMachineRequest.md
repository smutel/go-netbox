# PatchedWritableVirtualMachineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
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

### NewPatchedWritableVirtualMachineRequest

`func NewPatchedWritableVirtualMachineRequest() *PatchedWritableVirtualMachineRequest`

NewPatchedWritableVirtualMachineRequest instantiates a new PatchedWritableVirtualMachineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableVirtualMachineRequestWithDefaults

`func NewPatchedWritableVirtualMachineRequestWithDefaults() *PatchedWritableVirtualMachineRequest`

NewPatchedWritableVirtualMachineRequestWithDefaults instantiates a new PatchedWritableVirtualMachineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableVirtualMachineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableVirtualMachineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableVirtualMachineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableVirtualMachineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVirtualMachineType

`func (o *PatchedWritableVirtualMachineRequest) GetVirtualMachineType() BulkVirtualMachineRequestVirtualMachineType`

GetVirtualMachineType returns the VirtualMachineType field if non-nil, zero value otherwise.

### GetVirtualMachineTypeOk

`func (o *PatchedWritableVirtualMachineRequest) GetVirtualMachineTypeOk() (*BulkVirtualMachineRequestVirtualMachineType, bool)`

GetVirtualMachineTypeOk returns a tuple with the VirtualMachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineType

`func (o *PatchedWritableVirtualMachineRequest) SetVirtualMachineType(v BulkVirtualMachineRequestVirtualMachineType)`

SetVirtualMachineType sets VirtualMachineType field to given value.

### HasVirtualMachineType

`func (o *PatchedWritableVirtualMachineRequest) HasVirtualMachineType() bool`

HasVirtualMachineType returns a boolean if a field has been set.

### SetVirtualMachineTypeNil

`func (o *PatchedWritableVirtualMachineRequest) SetVirtualMachineTypeNil(b bool)`

 SetVirtualMachineTypeNil sets the value for VirtualMachineType to be an explicit nil

### UnsetVirtualMachineType
`func (o *PatchedWritableVirtualMachineRequest) UnsetVirtualMachineType()`

UnsetVirtualMachineType ensures that no value is present for VirtualMachineType, not even an explicit nil
### GetRole

`func (o *PatchedWritableVirtualMachineRequest) GetRole() BulkVirtualMachineRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableVirtualMachineRequest) GetRoleOk() (*BulkVirtualMachineRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableVirtualMachineRequest) SetRole(v BulkVirtualMachineRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableVirtualMachineRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedWritableVirtualMachineRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedWritableVirtualMachineRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetStatus

`func (o *PatchedWritableVirtualMachineRequest) GetStatus() BulkVirtualMachineRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableVirtualMachineRequest) GetStatusOk() (*BulkVirtualMachineRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableVirtualMachineRequest) SetStatus(v BulkVirtualMachineRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableVirtualMachineRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartOnBoot

`func (o *PatchedWritableVirtualMachineRequest) GetStartOnBoot() BulkVirtualMachineRequestStartOnBoot`

GetStartOnBoot returns the StartOnBoot field if non-nil, zero value otherwise.

### GetStartOnBootOk

`func (o *PatchedWritableVirtualMachineRequest) GetStartOnBootOk() (*BulkVirtualMachineRequestStartOnBoot, bool)`

GetStartOnBootOk returns a tuple with the StartOnBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOnBoot

`func (o *PatchedWritableVirtualMachineRequest) SetStartOnBoot(v BulkVirtualMachineRequestStartOnBoot)`

SetStartOnBoot sets StartOnBoot field to given value.

### HasStartOnBoot

`func (o *PatchedWritableVirtualMachineRequest) HasStartOnBoot() bool`

HasStartOnBoot returns a boolean if a field has been set.

### GetSite

`func (o *PatchedWritableVirtualMachineRequest) GetSite() BulkVLANRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableVirtualMachineRequest) GetSiteOk() (*BulkVLANRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableVirtualMachineRequest) SetSite(v BulkVLANRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableVirtualMachineRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *PatchedWritableVirtualMachineRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *PatchedWritableVirtualMachineRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetCluster

`func (o *PatchedWritableVirtualMachineRequest) GetCluster() BulkDeviceRequestCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedWritableVirtualMachineRequest) GetClusterOk() (*BulkDeviceRequestCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedWritableVirtualMachineRequest) SetCluster(v BulkDeviceRequestCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedWritableVirtualMachineRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *PatchedWritableVirtualMachineRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *PatchedWritableVirtualMachineRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetDevice

`func (o *PatchedWritableVirtualMachineRequest) GetDevice() BulkDeviceBayRequestInstalledDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableVirtualMachineRequest) GetDeviceOk() (*BulkDeviceBayRequestInstalledDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableVirtualMachineRequest) SetDevice(v BulkDeviceBayRequestInstalledDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableVirtualMachineRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *PatchedWritableVirtualMachineRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *PatchedWritableVirtualMachineRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetPlatform

`func (o *PatchedWritableVirtualMachineRequest) GetPlatform() BulkDeviceRequestPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedWritableVirtualMachineRequest) GetPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedWritableVirtualMachineRequest) SetPlatform(v BulkDeviceRequestPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedWritableVirtualMachineRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedWritableVirtualMachineRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedWritableVirtualMachineRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPrimaryIp4

`func (o *PatchedWritableVirtualMachineRequest) GetPrimaryIp4() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedWritableVirtualMachineRequest) GetPrimaryIp4Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedWritableVirtualMachineRequest) SetPrimaryIp4(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedWritableVirtualMachineRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedWritableVirtualMachineRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedWritableVirtualMachineRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedWritableVirtualMachineRequest) GetPrimaryIp6() BulkDeviceRequestPrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedWritableVirtualMachineRequest) GetPrimaryIp6Ok() (*BulkDeviceRequestPrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedWritableVirtualMachineRequest) SetPrimaryIp6(v BulkDeviceRequestPrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedWritableVirtualMachineRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedWritableVirtualMachineRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedWritableVirtualMachineRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetVcpus

`func (o *PatchedWritableVirtualMachineRequest) GetVcpus() float64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *PatchedWritableVirtualMachineRequest) GetVcpusOk() (*float64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *PatchedWritableVirtualMachineRequest) SetVcpus(v float64)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *PatchedWritableVirtualMachineRequest) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *PatchedWritableVirtualMachineRequest) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *PatchedWritableVirtualMachineRequest) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *PatchedWritableVirtualMachineRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *PatchedWritableVirtualMachineRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *PatchedWritableVirtualMachineRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *PatchedWritableVirtualMachineRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *PatchedWritableVirtualMachineRequest) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *PatchedWritableVirtualMachineRequest) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *PatchedWritableVirtualMachineRequest) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *PatchedWritableVirtualMachineRequest) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *PatchedWritableVirtualMachineRequest) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *PatchedWritableVirtualMachineRequest) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *PatchedWritableVirtualMachineRequest) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *PatchedWritableVirtualMachineRequest) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetDescription

`func (o *PatchedWritableVirtualMachineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableVirtualMachineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableVirtualMachineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableVirtualMachineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSerial

`func (o *PatchedWritableVirtualMachineRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableVirtualMachineRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableVirtualMachineRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableVirtualMachineRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableVirtualMachineRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableVirtualMachineRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableVirtualMachineRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableVirtualMachineRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableVirtualMachineRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableVirtualMachineRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetOwner

`func (o *PatchedWritableVirtualMachineRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedWritableVirtualMachineRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedWritableVirtualMachineRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedWritableVirtualMachineRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedWritableVirtualMachineRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedWritableVirtualMachineRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedWritableVirtualMachineRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableVirtualMachineRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableVirtualMachineRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableVirtualMachineRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableVirtualMachineRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableVirtualMachineRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableVirtualMachineRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableVirtualMachineRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLocalContextData

`func (o *PatchedWritableVirtualMachineRequest) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *PatchedWritableVirtualMachineRequest) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *PatchedWritableVirtualMachineRequest) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *PatchedWritableVirtualMachineRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *PatchedWritableVirtualMachineRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *PatchedWritableVirtualMachineRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetConfigTemplate

`func (o *PatchedWritableVirtualMachineRequest) GetConfigTemplate() BulkDeviceRequestConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *PatchedWritableVirtualMachineRequest) GetConfigTemplateOk() (*BulkDeviceRequestConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *PatchedWritableVirtualMachineRequest) SetConfigTemplate(v BulkDeviceRequestConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *PatchedWritableVirtualMachineRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *PatchedWritableVirtualMachineRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *PatchedWritableVirtualMachineRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetCustomFields

`func (o *PatchedWritableVirtualMachineRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableVirtualMachineRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableVirtualMachineRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableVirtualMachineRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


