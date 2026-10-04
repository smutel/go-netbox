# PatchedBulkSiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** | Full name of the site | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkLocationRequestStatus**](BulkLocationRequestStatus.md) |  | [optional] 
**Region** | Pointer to [**NullableBulkSiteRequestRegion**](BulkSiteRequestRegion.md) |  | [optional] 
**Group** | Pointer to [**NullableBulkSiteRequestGroup**](BulkSiteRequestGroup.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Facility** | Pointer to **string** | Local facility ID or description | [optional] 
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PhysicalAddress** | Pointer to **string** | Physical location of the building | [optional] 
**ShippingAddress** | Pointer to **string** | If different from the physical address | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Asns** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkSiteRequest

`func NewPatchedBulkSiteRequest(id int32, ) *PatchedBulkSiteRequest`

NewPatchedBulkSiteRequest instantiates a new PatchedBulkSiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkSiteRequestWithDefaults

`func NewPatchedBulkSiteRequestWithDefaults() *PatchedBulkSiteRequest`

NewPatchedBulkSiteRequestWithDefaults instantiates a new PatchedBulkSiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkSiteRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkSiteRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkSiteRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkSiteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkSiteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkSiteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkSiteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedBulkSiteRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedBulkSiteRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedBulkSiteRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedBulkSiteRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkSiteRequest) GetStatus() BulkLocationRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkSiteRequest) GetStatusOk() (*BulkLocationRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkSiteRequest) SetStatus(v BulkLocationRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkSiteRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegion

`func (o *PatchedBulkSiteRequest) GetRegion() BulkSiteRequestRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PatchedBulkSiteRequest) GetRegionOk() (*BulkSiteRequestRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PatchedBulkSiteRequest) SetRegion(v BulkSiteRequestRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PatchedBulkSiteRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *PatchedBulkSiteRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *PatchedBulkSiteRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetGroup

`func (o *PatchedBulkSiteRequest) GetGroup() BulkSiteRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkSiteRequest) GetGroupOk() (*BulkSiteRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkSiteRequest) SetGroup(v BulkSiteRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkSiteRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedBulkSiteRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedBulkSiteRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetTenant

`func (o *PatchedBulkSiteRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkSiteRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkSiteRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkSiteRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkSiteRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkSiteRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetFacility

`func (o *PatchedBulkSiteRequest) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *PatchedBulkSiteRequest) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *PatchedBulkSiteRequest) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *PatchedBulkSiteRequest) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetTimeZone

`func (o *PatchedBulkSiteRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PatchedBulkSiteRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PatchedBulkSiteRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PatchedBulkSiteRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *PatchedBulkSiteRequest) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *PatchedBulkSiteRequest) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetDescription

`func (o *PatchedBulkSiteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkSiteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkSiteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkSiteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhysicalAddress

`func (o *PatchedBulkSiteRequest) GetPhysicalAddress() string`

GetPhysicalAddress returns the PhysicalAddress field if non-nil, zero value otherwise.

### GetPhysicalAddressOk

`func (o *PatchedBulkSiteRequest) GetPhysicalAddressOk() (*string, bool)`

GetPhysicalAddressOk returns a tuple with the PhysicalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAddress

`func (o *PatchedBulkSiteRequest) SetPhysicalAddress(v string)`

SetPhysicalAddress sets PhysicalAddress field to given value.

### HasPhysicalAddress

`func (o *PatchedBulkSiteRequest) HasPhysicalAddress() bool`

HasPhysicalAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *PatchedBulkSiteRequest) GetShippingAddress() string`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *PatchedBulkSiteRequest) GetShippingAddressOk() (*string, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *PatchedBulkSiteRequest) SetShippingAddress(v string)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *PatchedBulkSiteRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *PatchedBulkSiteRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PatchedBulkSiteRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PatchedBulkSiteRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PatchedBulkSiteRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *PatchedBulkSiteRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *PatchedBulkSiteRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *PatchedBulkSiteRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PatchedBulkSiteRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PatchedBulkSiteRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PatchedBulkSiteRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *PatchedBulkSiteRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *PatchedBulkSiteRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetOwner

`func (o *PatchedBulkSiteRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkSiteRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkSiteRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkSiteRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkSiteRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkSiteRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkSiteRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkSiteRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkSiteRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkSiteRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetAsns

`func (o *PatchedBulkSiteRequest) GetAsns() []int32`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *PatchedBulkSiteRequest) GetAsnsOk() (*[]int32, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *PatchedBulkSiteRequest) SetAsns(v []int32)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *PatchedBulkSiteRequest) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkSiteRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkSiteRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkSiteRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkSiteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkSiteRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkSiteRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkSiteRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkSiteRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


