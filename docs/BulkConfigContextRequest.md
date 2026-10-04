# BulkConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
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
**Data** | **interface{}** |  | 

## Methods

### NewBulkConfigContextRequest

`func NewBulkConfigContextRequest(id int32, name string, data interface{}, ) *BulkConfigContextRequest`

NewBulkConfigContextRequest instantiates a new BulkConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkConfigContextRequestWithDefaults

`func NewBulkConfigContextRequestWithDefaults() *BulkConfigContextRequest`

NewBulkConfigContextRequestWithDefaults instantiates a new BulkConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkConfigContextRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkConfigContextRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkConfigContextRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *BulkConfigContextRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkConfigContextRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkConfigContextRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkConfigContextRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetProfile

`func (o *BulkConfigContextRequest) GetProfile() BulkConfigContextRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *BulkConfigContextRequest) GetProfileOk() (*BulkConfigContextRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *BulkConfigContextRequest) SetProfile(v BulkConfigContextRequestProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *BulkConfigContextRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *BulkConfigContextRequest) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *BulkConfigContextRequest) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetDescription

`func (o *BulkConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *BulkConfigContextRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BulkConfigContextRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BulkConfigContextRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BulkConfigContextRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRegions

`func (o *BulkConfigContextRequest) GetRegions() []int32`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *BulkConfigContextRequest) GetRegionsOk() (*[]int32, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *BulkConfigContextRequest) SetRegions(v []int32)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *BulkConfigContextRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSiteGroups

`func (o *BulkConfigContextRequest) GetSiteGroups() []int32`

GetSiteGroups returns the SiteGroups field if non-nil, zero value otherwise.

### GetSiteGroupsOk

`func (o *BulkConfigContextRequest) GetSiteGroupsOk() (*[]int32, bool)`

GetSiteGroupsOk returns a tuple with the SiteGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroups

`func (o *BulkConfigContextRequest) SetSiteGroups(v []int32)`

SetSiteGroups sets SiteGroups field to given value.

### HasSiteGroups

`func (o *BulkConfigContextRequest) HasSiteGroups() bool`

HasSiteGroups returns a boolean if a field has been set.

### GetSites

`func (o *BulkConfigContextRequest) GetSites() []int32`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *BulkConfigContextRequest) GetSitesOk() (*[]int32, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *BulkConfigContextRequest) SetSites(v []int32)`

SetSites sets Sites field to given value.

### HasSites

`func (o *BulkConfigContextRequest) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetLocations

`func (o *BulkConfigContextRequest) GetLocations() []int32`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *BulkConfigContextRequest) GetLocationsOk() (*[]int32, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *BulkConfigContextRequest) SetLocations(v []int32)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *BulkConfigContextRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *BulkConfigContextRequest) GetDeviceTypes() []int32`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *BulkConfigContextRequest) GetDeviceTypesOk() (*[]int32, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *BulkConfigContextRequest) SetDeviceTypes(v []int32)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *BulkConfigContextRequest) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetRoles

`func (o *BulkConfigContextRequest) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BulkConfigContextRequest) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BulkConfigContextRequest) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BulkConfigContextRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPlatforms

`func (o *BulkConfigContextRequest) GetPlatforms() []int32`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *BulkConfigContextRequest) GetPlatformsOk() (*[]int32, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *BulkConfigContextRequest) SetPlatforms(v []int32)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *BulkConfigContextRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterTypes

`func (o *BulkConfigContextRequest) GetClusterTypes() []int32`

GetClusterTypes returns the ClusterTypes field if non-nil, zero value otherwise.

### GetClusterTypesOk

`func (o *BulkConfigContextRequest) GetClusterTypesOk() (*[]int32, bool)`

GetClusterTypesOk returns a tuple with the ClusterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypes

`func (o *BulkConfigContextRequest) SetClusterTypes(v []int32)`

SetClusterTypes sets ClusterTypes field to given value.

### HasClusterTypes

`func (o *BulkConfigContextRequest) HasClusterTypes() bool`

HasClusterTypes returns a boolean if a field has been set.

### GetClusterGroups

`func (o *BulkConfigContextRequest) GetClusterGroups() []int32`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *BulkConfigContextRequest) GetClusterGroupsOk() (*[]int32, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *BulkConfigContextRequest) SetClusterGroups(v []int32)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *BulkConfigContextRequest) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *BulkConfigContextRequest) GetClusters() []int32`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *BulkConfigContextRequest) GetClustersOk() (*[]int32, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *BulkConfigContextRequest) SetClusters(v []int32)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *BulkConfigContextRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *BulkConfigContextRequest) GetTenantGroups() []int32`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *BulkConfigContextRequest) GetTenantGroupsOk() (*[]int32, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *BulkConfigContextRequest) SetTenantGroups(v []int32)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *BulkConfigContextRequest) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *BulkConfigContextRequest) GetTenants() []int32`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *BulkConfigContextRequest) GetTenantsOk() (*[]int32, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *BulkConfigContextRequest) SetTenants(v []int32)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *BulkConfigContextRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetOwner

`func (o *BulkConfigContextRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkConfigContextRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkConfigContextRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkConfigContextRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkConfigContextRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkConfigContextRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *BulkConfigContextRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkConfigContextRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkConfigContextRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDataSource

`func (o *BulkConfigContextRequest) GetDataSource() BulkConfigContextProfileRequestDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *BulkConfigContextRequest) GetDataSourceOk() (*BulkConfigContextProfileRequestDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *BulkConfigContextRequest) SetDataSource(v BulkConfigContextProfileRequestDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *BulkConfigContextRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetData

`func (o *BulkConfigContextRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkConfigContextRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkConfigContextRequest) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *BulkConfigContextRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BulkConfigContextRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


