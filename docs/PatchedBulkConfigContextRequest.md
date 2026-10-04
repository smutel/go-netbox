# PatchedBulkConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Profile** | Pointer to [**NullableBulkConfigContextRequestProfile**](BulkConfigContextRequestProfile.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Regions** | Pointer to **[]int32** |  | [optional] 
**SiteGroups** | Pointer to **[]int32** |  | [optional] 
**Sites** | Pointer to **[]int32** |  | [optional] 
**Locations** | Pointer to **[]int32** |  | [optional] 
**DeviceTypes** | Pointer to **[]int32** |  | [optional] 
**Roles** | Pointer to **[]int32** |  | [optional] 
**Platforms** | Pointer to **[]int32** |  | [optional] 
**ClusterTypes** | Pointer to **[]int32** |  | [optional] 
**ClusterGroups** | Pointer to **[]int32** |  | [optional] 
**Clusters** | Pointer to **[]int32** |  | [optional] 
**TenantGroups** | Pointer to **[]int32** |  | [optional] 
**Tenants** | Pointer to **[]int32** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DataSource** | Pointer to [**BulkConfigContextProfileRequestDataSource**](BulkConfigContextProfileRequestDataSource.md) |  | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedBulkConfigContextRequest

`func NewPatchedBulkConfigContextRequest(id int32, ) *PatchedBulkConfigContextRequest`

NewPatchedBulkConfigContextRequest instantiates a new PatchedBulkConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkConfigContextRequestWithDefaults

`func NewPatchedBulkConfigContextRequestWithDefaults() *PatchedBulkConfigContextRequest`

NewPatchedBulkConfigContextRequestWithDefaults instantiates a new PatchedBulkConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkConfigContextRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkConfigContextRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkConfigContextRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkConfigContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkConfigContextRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkConfigContextRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkConfigContextRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkConfigContextRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetProfile

`func (o *PatchedBulkConfigContextRequest) GetProfile() BulkConfigContextRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PatchedBulkConfigContextRequest) GetProfileOk() (*BulkConfigContextRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PatchedBulkConfigContextRequest) SetProfile(v BulkConfigContextRequestProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PatchedBulkConfigContextRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *PatchedBulkConfigContextRequest) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *PatchedBulkConfigContextRequest) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetDescription

`func (o *PatchedBulkConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchedBulkConfigContextRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchedBulkConfigContextRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchedBulkConfigContextRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchedBulkConfigContextRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRegions

`func (o *PatchedBulkConfigContextRequest) GetRegions() []int32`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PatchedBulkConfigContextRequest) GetRegionsOk() (*[]int32, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PatchedBulkConfigContextRequest) SetRegions(v []int32)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PatchedBulkConfigContextRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSiteGroups

`func (o *PatchedBulkConfigContextRequest) GetSiteGroups() []int32`

GetSiteGroups returns the SiteGroups field if non-nil, zero value otherwise.

### GetSiteGroupsOk

`func (o *PatchedBulkConfigContextRequest) GetSiteGroupsOk() (*[]int32, bool)`

GetSiteGroupsOk returns a tuple with the SiteGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroups

`func (o *PatchedBulkConfigContextRequest) SetSiteGroups(v []int32)`

SetSiteGroups sets SiteGroups field to given value.

### HasSiteGroups

`func (o *PatchedBulkConfigContextRequest) HasSiteGroups() bool`

HasSiteGroups returns a boolean if a field has been set.

### GetSites

`func (o *PatchedBulkConfigContextRequest) GetSites() []int32`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PatchedBulkConfigContextRequest) GetSitesOk() (*[]int32, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PatchedBulkConfigContextRequest) SetSites(v []int32)`

SetSites sets Sites field to given value.

### HasSites

`func (o *PatchedBulkConfigContextRequest) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetLocations

`func (o *PatchedBulkConfigContextRequest) GetLocations() []int32`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *PatchedBulkConfigContextRequest) GetLocationsOk() (*[]int32, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *PatchedBulkConfigContextRequest) SetLocations(v []int32)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *PatchedBulkConfigContextRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *PatchedBulkConfigContextRequest) GetDeviceTypes() []int32`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *PatchedBulkConfigContextRequest) GetDeviceTypesOk() (*[]int32, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *PatchedBulkConfigContextRequest) SetDeviceTypes(v []int32)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *PatchedBulkConfigContextRequest) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetRoles

`func (o *PatchedBulkConfigContextRequest) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PatchedBulkConfigContextRequest) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PatchedBulkConfigContextRequest) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PatchedBulkConfigContextRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPlatforms

`func (o *PatchedBulkConfigContextRequest) GetPlatforms() []int32`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *PatchedBulkConfigContextRequest) GetPlatformsOk() (*[]int32, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *PatchedBulkConfigContextRequest) SetPlatforms(v []int32)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *PatchedBulkConfigContextRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterTypes

`func (o *PatchedBulkConfigContextRequest) GetClusterTypes() []int32`

GetClusterTypes returns the ClusterTypes field if non-nil, zero value otherwise.

### GetClusterTypesOk

`func (o *PatchedBulkConfigContextRequest) GetClusterTypesOk() (*[]int32, bool)`

GetClusterTypesOk returns a tuple with the ClusterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypes

`func (o *PatchedBulkConfigContextRequest) SetClusterTypes(v []int32)`

SetClusterTypes sets ClusterTypes field to given value.

### HasClusterTypes

`func (o *PatchedBulkConfigContextRequest) HasClusterTypes() bool`

HasClusterTypes returns a boolean if a field has been set.

### GetClusterGroups

`func (o *PatchedBulkConfigContextRequest) GetClusterGroups() []int32`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *PatchedBulkConfigContextRequest) GetClusterGroupsOk() (*[]int32, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *PatchedBulkConfigContextRequest) SetClusterGroups(v []int32)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *PatchedBulkConfigContextRequest) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *PatchedBulkConfigContextRequest) GetClusters() []int32`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *PatchedBulkConfigContextRequest) GetClustersOk() (*[]int32, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *PatchedBulkConfigContextRequest) SetClusters(v []int32)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *PatchedBulkConfigContextRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *PatchedBulkConfigContextRequest) GetTenantGroups() []int32`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *PatchedBulkConfigContextRequest) GetTenantGroupsOk() (*[]int32, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *PatchedBulkConfigContextRequest) SetTenantGroups(v []int32)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *PatchedBulkConfigContextRequest) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *PatchedBulkConfigContextRequest) GetTenants() []int32`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *PatchedBulkConfigContextRequest) GetTenantsOk() (*[]int32, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *PatchedBulkConfigContextRequest) SetTenants(v []int32)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *PatchedBulkConfigContextRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkConfigContextRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkConfigContextRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkConfigContextRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkConfigContextRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkConfigContextRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkConfigContextRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkConfigContextRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkConfigContextRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkConfigContextRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDataSource

`func (o *PatchedBulkConfigContextRequest) GetDataSource() BulkConfigContextProfileRequestDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *PatchedBulkConfigContextRequest) GetDataSourceOk() (*BulkConfigContextProfileRequestDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *PatchedBulkConfigContextRequest) SetDataSource(v BulkConfigContextProfileRequestDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *PatchedBulkConfigContextRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetData

`func (o *PatchedBulkConfigContextRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatchedBulkConfigContextRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatchedBulkConfigContextRequest) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PatchedBulkConfigContextRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PatchedBulkConfigContextRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PatchedBulkConfigContextRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


