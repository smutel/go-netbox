# PatchedBulkCoolingSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Site** | Pointer to [**BulkCoolingSourceRequestSite**](BulkCoolingSourceRequestSite.md) |  | [optional] 
**Location** | Pointer to [**NullableBulkCoolingSourceRequestLocation**](BulkCoolingSourceRequestLocation.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**BulkCoolingSourceRequestType**](BulkCoolingSourceRequestType.md) |  | [optional] 
**Status** | Pointer to [**BulkCoolingFeedRequestStatus**](BulkCoolingFeedRequestStatus.md) |  | [optional] 
**FluidType** | Pointer to [**NullableBulkCoolingSourceRequestFluidType**](BulkCoolingSourceRequestFluidType.md) |  | [optional] 
**CoolingCapacity** | Pointer to **NullableFloat64** | Total rated cooling capacity (kW) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkCoolingSourceRequest

`func NewPatchedBulkCoolingSourceRequest(id int32, ) *PatchedBulkCoolingSourceRequest`

NewPatchedBulkCoolingSourceRequest instantiates a new PatchedBulkCoolingSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCoolingSourceRequestWithDefaults

`func NewPatchedBulkCoolingSourceRequestWithDefaults() *PatchedBulkCoolingSourceRequest`

NewPatchedBulkCoolingSourceRequestWithDefaults instantiates a new PatchedBulkCoolingSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCoolingSourceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCoolingSourceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCoolingSourceRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSite

`func (o *PatchedBulkCoolingSourceRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedBulkCoolingSourceRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedBulkCoolingSourceRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedBulkCoolingSourceRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedBulkCoolingSourceRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedBulkCoolingSourceRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedBulkCoolingSourceRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedBulkCoolingSourceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedBulkCoolingSourceRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedBulkCoolingSourceRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *PatchedBulkCoolingSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkCoolingSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkCoolingSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkCoolingSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkCoolingSourceRequest) GetType() BulkCoolingSourceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkCoolingSourceRequest) GetTypeOk() (*BulkCoolingSourceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkCoolingSourceRequest) SetType(v BulkCoolingSourceRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkCoolingSourceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkCoolingSourceRequest) GetStatus() BulkCoolingFeedRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkCoolingSourceRequest) GetStatusOk() (*BulkCoolingFeedRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkCoolingSourceRequest) SetStatus(v BulkCoolingFeedRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkCoolingSourceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFluidType

`func (o *PatchedBulkCoolingSourceRequest) GetFluidType() BulkCoolingSourceRequestFluidType`

GetFluidType returns the FluidType field if non-nil, zero value otherwise.

### GetFluidTypeOk

`func (o *PatchedBulkCoolingSourceRequest) GetFluidTypeOk() (*BulkCoolingSourceRequestFluidType, bool)`

GetFluidTypeOk returns a tuple with the FluidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFluidType

`func (o *PatchedBulkCoolingSourceRequest) SetFluidType(v BulkCoolingSourceRequestFluidType)`

SetFluidType sets FluidType field to given value.

### HasFluidType

`func (o *PatchedBulkCoolingSourceRequest) HasFluidType() bool`

HasFluidType returns a boolean if a field has been set.

### SetFluidTypeNil

`func (o *PatchedBulkCoolingSourceRequest) SetFluidTypeNil(b bool)`

 SetFluidTypeNil sets the value for FluidType to be an explicit nil

### UnsetFluidType
`func (o *PatchedBulkCoolingSourceRequest) UnsetFluidType()`

UnsetFluidType ensures that no value is present for FluidType, not even an explicit nil
### GetCoolingCapacity

`func (o *PatchedBulkCoolingSourceRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *PatchedBulkCoolingSourceRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *PatchedBulkCoolingSourceRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *PatchedBulkCoolingSourceRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *PatchedBulkCoolingSourceRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *PatchedBulkCoolingSourceRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetDescription

`func (o *PatchedBulkCoolingSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkCoolingSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkCoolingSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkCoolingSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkCoolingSourceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkCoolingSourceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkCoolingSourceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkCoolingSourceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkCoolingSourceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkCoolingSourceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkCoolingSourceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkCoolingSourceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkCoolingSourceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkCoolingSourceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkCoolingSourceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkCoolingSourceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkCoolingSourceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkCoolingSourceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkCoolingSourceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkCoolingSourceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkCoolingSourceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkCoolingSourceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


