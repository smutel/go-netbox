# VirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**VirtualMachineType** | Pointer to [**NullableBriefVirtualMachineType**](BriefVirtualMachineType.md) |  | [optional] 
**Role** | Pointer to [**NullableBriefDeviceRole**](BriefDeviceRole.md) |  | [optional] 
**Status** | Pointer to [**VirtualMachineStatus**](VirtualMachineStatus.md) |  | [optional] 
**StartOnBoot** | Pointer to [**VirtualMachineStartOnBoot**](VirtualMachineStartOnBoot.md) |  | [optional] 
**Site** | Pointer to [**NullableBriefSite**](BriefSite.md) |  | [optional] 
**Cluster** | Pointer to [**NullableBriefCluster**](BriefCluster.md) |  | [optional] 
**Device** | Pointer to [**NullableBriefDevice**](BriefDevice.md) |  | [optional] 
**Platform** | Pointer to [**NullableBriefPlatform**](BriefPlatform.md) |  | [optional] 
**PrimaryIp** | [**NullableBriefIPAddress**](BriefIPAddress.md) |  | [readonly] 
**PrimaryIp4** | Pointer to [**NullableBriefIPAddress**](BriefIPAddress.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableBriefIPAddress**](BriefIPAddress.md) |  | [optional] 
**Vcpus** | Pointer to **NullableFloat64** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Disk** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenant**](BriefTenant.md) |  | [optional] 
**Owner** | Pointer to [**NullableBriefOwner**](BriefOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**LocalContextData** | Pointer to **interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**ConfigTemplate** | Pointer to [**NullableBriefConfigTemplate**](BriefConfigTemplate.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**InterfaceCount** | **int32** |  | [readonly] 
**VirtualDiskCount** | **int32** |  | [readonly] 
**ConfigContext** | **interface{}** |  | [readonly] 

## Methods

### NewVirtualMachine

`func NewVirtualMachine(id int32, url string, displayUrl string, display string, name string, primaryIp NullableBriefIPAddress, created NullableTime, lastUpdated NullableTime, interfaceCount int32, virtualDiskCount int32, configContext interface{}, ) *VirtualMachine`

NewVirtualMachine instantiates a new VirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineWithDefaults

`func NewVirtualMachineWithDefaults() *VirtualMachine`

NewVirtualMachineWithDefaults instantiates a new VirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualMachine) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualMachine) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualMachine) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *VirtualMachine) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualMachine) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualMachine) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *VirtualMachine) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *VirtualMachine) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *VirtualMachine) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *VirtualMachine) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualMachine) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualMachine) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *VirtualMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachine) SetName(v string)`

SetName sets Name field to given value.


### GetVirtualMachineType

`func (o *VirtualMachine) GetVirtualMachineType() BriefVirtualMachineType`

GetVirtualMachineType returns the VirtualMachineType field if non-nil, zero value otherwise.

### GetVirtualMachineTypeOk

`func (o *VirtualMachine) GetVirtualMachineTypeOk() (*BriefVirtualMachineType, bool)`

GetVirtualMachineTypeOk returns a tuple with the VirtualMachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineType

`func (o *VirtualMachine) SetVirtualMachineType(v BriefVirtualMachineType)`

SetVirtualMachineType sets VirtualMachineType field to given value.

### HasVirtualMachineType

`func (o *VirtualMachine) HasVirtualMachineType() bool`

HasVirtualMachineType returns a boolean if a field has been set.

### SetVirtualMachineTypeNil

`func (o *VirtualMachine) SetVirtualMachineTypeNil(b bool)`

 SetVirtualMachineTypeNil sets the value for VirtualMachineType to be an explicit nil

### UnsetVirtualMachineType
`func (o *VirtualMachine) UnsetVirtualMachineType()`

UnsetVirtualMachineType ensures that no value is present for VirtualMachineType, not even an explicit nil
### GetRole

`func (o *VirtualMachine) GetRole() BriefDeviceRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VirtualMachine) GetRoleOk() (*BriefDeviceRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VirtualMachine) SetRole(v BriefDeviceRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *VirtualMachine) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VirtualMachine) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VirtualMachine) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetStatus

`func (o *VirtualMachine) GetStatus() VirtualMachineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualMachine) GetStatusOk() (*VirtualMachineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualMachine) SetStatus(v VirtualMachineStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualMachine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartOnBoot

`func (o *VirtualMachine) GetStartOnBoot() VirtualMachineStartOnBoot`

GetStartOnBoot returns the StartOnBoot field if non-nil, zero value otherwise.

### GetStartOnBootOk

`func (o *VirtualMachine) GetStartOnBootOk() (*VirtualMachineStartOnBoot, bool)`

GetStartOnBootOk returns a tuple with the StartOnBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOnBoot

`func (o *VirtualMachine) SetStartOnBoot(v VirtualMachineStartOnBoot)`

SetStartOnBoot sets StartOnBoot field to given value.

### HasStartOnBoot

`func (o *VirtualMachine) HasStartOnBoot() bool`

HasStartOnBoot returns a boolean if a field has been set.

### GetSite

`func (o *VirtualMachine) GetSite() BriefSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *VirtualMachine) GetSiteOk() (*BriefSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *VirtualMachine) SetSite(v BriefSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *VirtualMachine) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *VirtualMachine) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *VirtualMachine) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetCluster

`func (o *VirtualMachine) GetCluster() BriefCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualMachine) GetClusterOk() (*BriefCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualMachine) SetCluster(v BriefCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualMachine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *VirtualMachine) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *VirtualMachine) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetDevice

`func (o *VirtualMachine) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VirtualMachine) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VirtualMachine) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VirtualMachine) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *VirtualMachine) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *VirtualMachine) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetPlatform

`func (o *VirtualMachine) GetPlatform() BriefPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VirtualMachine) GetPlatformOk() (*BriefPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VirtualMachine) SetPlatform(v BriefPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VirtualMachine) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *VirtualMachine) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *VirtualMachine) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPrimaryIp

`func (o *VirtualMachine) GetPrimaryIp() BriefIPAddress`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *VirtualMachine) GetPrimaryIpOk() (*BriefIPAddress, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *VirtualMachine) SetPrimaryIp(v BriefIPAddress)`

SetPrimaryIp sets PrimaryIp field to given value.


### SetPrimaryIpNil

`func (o *VirtualMachine) SetPrimaryIpNil(b bool)`

 SetPrimaryIpNil sets the value for PrimaryIp to be an explicit nil

### UnsetPrimaryIp
`func (o *VirtualMachine) UnsetPrimaryIp()`

UnsetPrimaryIp ensures that no value is present for PrimaryIp, not even an explicit nil
### GetPrimaryIp4

`func (o *VirtualMachine) GetPrimaryIp4() BriefIPAddress`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *VirtualMachine) GetPrimaryIp4Ok() (*BriefIPAddress, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *VirtualMachine) SetPrimaryIp4(v BriefIPAddress)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *VirtualMachine) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *VirtualMachine) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *VirtualMachine) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *VirtualMachine) GetPrimaryIp6() BriefIPAddress`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *VirtualMachine) GetPrimaryIp6Ok() (*BriefIPAddress, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *VirtualMachine) SetPrimaryIp6(v BriefIPAddress)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *VirtualMachine) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *VirtualMachine) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *VirtualMachine) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetVcpus

`func (o *VirtualMachine) GetVcpus() float64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *VirtualMachine) GetVcpusOk() (*float64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *VirtualMachine) SetVcpus(v float64)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *VirtualMachine) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *VirtualMachine) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *VirtualMachine) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *VirtualMachine) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VirtualMachine) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VirtualMachine) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VirtualMachine) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *VirtualMachine) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *VirtualMachine) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *VirtualMachine) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VirtualMachine) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VirtualMachine) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *VirtualMachine) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *VirtualMachine) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *VirtualMachine) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetDescription

`func (o *VirtualMachine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualMachine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualMachine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualMachine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSerial

`func (o *VirtualMachine) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VirtualMachine) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VirtualMachine) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VirtualMachine) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTenant

`func (o *VirtualMachine) GetTenant() BriefTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VirtualMachine) GetTenantOk() (*BriefTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VirtualMachine) SetTenant(v BriefTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VirtualMachine) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VirtualMachine) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VirtualMachine) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetOwner

`func (o *VirtualMachine) GetOwner() BriefOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VirtualMachine) GetOwnerOk() (*BriefOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VirtualMachine) SetOwner(v BriefOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VirtualMachine) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VirtualMachine) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VirtualMachine) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VirtualMachine) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VirtualMachine) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VirtualMachine) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VirtualMachine) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VirtualMachine) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualMachine) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualMachine) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualMachine) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLocalContextData

`func (o *VirtualMachine) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *VirtualMachine) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *VirtualMachine) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *VirtualMachine) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *VirtualMachine) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *VirtualMachine) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetConfigTemplate

`func (o *VirtualMachine) GetConfigTemplate() BriefConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *VirtualMachine) GetConfigTemplateOk() (*BriefConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *VirtualMachine) SetConfigTemplate(v BriefConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *VirtualMachine) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *VirtualMachine) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *VirtualMachine) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetCustomFields

`func (o *VirtualMachine) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualMachine) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualMachine) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualMachine) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *VirtualMachine) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VirtualMachine) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VirtualMachine) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VirtualMachine) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VirtualMachine) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VirtualMachine) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VirtualMachine) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VirtualMachine) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VirtualMachine) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VirtualMachine) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetInterfaceCount

`func (o *VirtualMachine) GetInterfaceCount() int32`

GetInterfaceCount returns the InterfaceCount field if non-nil, zero value otherwise.

### GetInterfaceCountOk

`func (o *VirtualMachine) GetInterfaceCountOk() (*int32, bool)`

GetInterfaceCountOk returns a tuple with the InterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceCount

`func (o *VirtualMachine) SetInterfaceCount(v int32)`

SetInterfaceCount sets InterfaceCount field to given value.


### GetVirtualDiskCount

`func (o *VirtualMachine) GetVirtualDiskCount() int32`

GetVirtualDiskCount returns the VirtualDiskCount field if non-nil, zero value otherwise.

### GetVirtualDiskCountOk

`func (o *VirtualMachine) GetVirtualDiskCountOk() (*int32, bool)`

GetVirtualDiskCountOk returns a tuple with the VirtualDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskCount

`func (o *VirtualMachine) SetVirtualDiskCount(v int32)`

SetVirtualDiskCount sets VirtualDiskCount field to given value.


### GetConfigContext

`func (o *VirtualMachine) GetConfigContext() interface{}`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *VirtualMachine) GetConfigContextOk() (*interface{}, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *VirtualMachine) SetConfigContext(v interface{})`

SetConfigContext sets ConfigContext field to given value.


### SetConfigContextNil

`func (o *VirtualMachine) SetConfigContextNil(b bool)`

 SetConfigContextNil sets the value for ConfigContext to be an explicit nil

### UnsetConfigContext
`func (o *VirtualMachine) UnsetConfigContext()`

UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


