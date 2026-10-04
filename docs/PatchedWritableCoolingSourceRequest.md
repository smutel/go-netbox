# PatchedWritableCoolingSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewPatchedWritableCoolingSourceRequest

`func NewPatchedWritableCoolingSourceRequest() *PatchedWritableCoolingSourceRequest`

NewPatchedWritableCoolingSourceRequest instantiates a new PatchedWritableCoolingSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCoolingSourceRequestWithDefaults

`func NewPatchedWritableCoolingSourceRequestWithDefaults() *PatchedWritableCoolingSourceRequest`

NewPatchedWritableCoolingSourceRequestWithDefaults instantiates a new PatchedWritableCoolingSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *PatchedWritableCoolingSourceRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableCoolingSourceRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableCoolingSourceRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableCoolingSourceRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedWritableCoolingSourceRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWritableCoolingSourceRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWritableCoolingSourceRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWritableCoolingSourceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedWritableCoolingSourceRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedWritableCoolingSourceRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *PatchedWritableCoolingSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableCoolingSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableCoolingSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableCoolingSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableCoolingSourceRequest) GetType() BulkCoolingSourceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableCoolingSourceRequest) GetTypeOk() (*BulkCoolingSourceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableCoolingSourceRequest) SetType(v BulkCoolingSourceRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableCoolingSourceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableCoolingSourceRequest) GetStatus() BulkCoolingFeedRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableCoolingSourceRequest) GetStatusOk() (*BulkCoolingFeedRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableCoolingSourceRequest) SetStatus(v BulkCoolingFeedRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableCoolingSourceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFluidType

`func (o *PatchedWritableCoolingSourceRequest) GetFluidType() BulkCoolingSourceRequestFluidType`

GetFluidType returns the FluidType field if non-nil, zero value otherwise.

### GetFluidTypeOk

`func (o *PatchedWritableCoolingSourceRequest) GetFluidTypeOk() (*BulkCoolingSourceRequestFluidType, bool)`

GetFluidTypeOk returns a tuple with the FluidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFluidType

`func (o *PatchedWritableCoolingSourceRequest) SetFluidType(v BulkCoolingSourceRequestFluidType)`

SetFluidType sets FluidType field to given value.

### HasFluidType

`func (o *PatchedWritableCoolingSourceRequest) HasFluidType() bool`

HasFluidType returns a boolean if a field has been set.

### SetFluidTypeNil

`func (o *PatchedWritableCoolingSourceRequest) SetFluidTypeNil(b bool)`

 SetFluidTypeNil sets the value for FluidType to be an explicit nil

### UnsetFluidType
`func (o *PatchedWritableCoolingSourceRequest) UnsetFluidType()`

UnsetFluidType ensures that no value is present for FluidType, not even an explicit nil
### GetCoolingCapacity

`func (o *PatchedWritableCoolingSourceRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *PatchedWritableCoolingSourceRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *PatchedWritableCoolingSourceRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *PatchedWritableCoolingSourceRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *PatchedWritableCoolingSourceRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *PatchedWritableCoolingSourceRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetDescription

`func (o *PatchedWritableCoolingSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableCoolingSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableCoolingSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableCoolingSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedWritableCoolingSourceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedWritableCoolingSourceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedWritableCoolingSourceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedWritableCoolingSourceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedWritableCoolingSourceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedWritableCoolingSourceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedWritableCoolingSourceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableCoolingSourceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableCoolingSourceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableCoolingSourceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableCoolingSourceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableCoolingSourceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableCoolingSourceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableCoolingSourceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableCoolingSourceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableCoolingSourceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableCoolingSourceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableCoolingSourceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


