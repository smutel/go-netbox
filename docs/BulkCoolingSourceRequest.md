# BulkCoolingSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Site** | [**BulkCoolingSourceRequestSite**](BulkCoolingSourceRequestSite.md) |  | 
**Location** | Pointer to [**NullableBulkCoolingSourceRequestLocation**](BulkCoolingSourceRequestLocation.md) |  | [optional] 
**Name** | **string** |  | 
**Type** | [**BulkCoolingSourceRequestType**](BulkCoolingSourceRequestType.md) |  | 
**Status** | Pointer to [**BulkCoolingFeedRequestStatus**](BulkCoolingFeedRequestStatus.md) |  | [optional] 
**FluidType** | Pointer to [**NullableBulkCoolingSourceRequestFluidType**](BulkCoolingSourceRequestFluidType.md) |  | [optional] 
**CoolingCapacity** | Pointer to **NullableFloat64** | Total rated cooling capacity (kW) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkCoolingSourceRequest

`func NewBulkCoolingSourceRequest(id int32, site BulkCoolingSourceRequestSite, name string, type_ BulkCoolingSourceRequestType, ) *BulkCoolingSourceRequest`

NewBulkCoolingSourceRequest instantiates a new BulkCoolingSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCoolingSourceRequestWithDefaults

`func NewBulkCoolingSourceRequestWithDefaults() *BulkCoolingSourceRequest`

NewBulkCoolingSourceRequestWithDefaults instantiates a new BulkCoolingSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCoolingSourceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCoolingSourceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCoolingSourceRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSite

`func (o *BulkCoolingSourceRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *BulkCoolingSourceRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *BulkCoolingSourceRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *BulkCoolingSourceRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BulkCoolingSourceRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BulkCoolingSourceRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BulkCoolingSourceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *BulkCoolingSourceRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *BulkCoolingSourceRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *BulkCoolingSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkCoolingSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkCoolingSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BulkCoolingSourceRequest) GetType() BulkCoolingSourceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkCoolingSourceRequest) GetTypeOk() (*BulkCoolingSourceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkCoolingSourceRequest) SetType(v BulkCoolingSourceRequestType)`

SetType sets Type field to given value.


### GetStatus

`func (o *BulkCoolingSourceRequest) GetStatus() BulkCoolingFeedRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkCoolingSourceRequest) GetStatusOk() (*BulkCoolingFeedRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkCoolingSourceRequest) SetStatus(v BulkCoolingFeedRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkCoolingSourceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFluidType

`func (o *BulkCoolingSourceRequest) GetFluidType() BulkCoolingSourceRequestFluidType`

GetFluidType returns the FluidType field if non-nil, zero value otherwise.

### GetFluidTypeOk

`func (o *BulkCoolingSourceRequest) GetFluidTypeOk() (*BulkCoolingSourceRequestFluidType, bool)`

GetFluidTypeOk returns a tuple with the FluidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFluidType

`func (o *BulkCoolingSourceRequest) SetFluidType(v BulkCoolingSourceRequestFluidType)`

SetFluidType sets FluidType field to given value.

### HasFluidType

`func (o *BulkCoolingSourceRequest) HasFluidType() bool`

HasFluidType returns a boolean if a field has been set.

### SetFluidTypeNil

`func (o *BulkCoolingSourceRequest) SetFluidTypeNil(b bool)`

 SetFluidTypeNil sets the value for FluidType to be an explicit nil

### UnsetFluidType
`func (o *BulkCoolingSourceRequest) UnsetFluidType()`

UnsetFluidType ensures that no value is present for FluidType, not even an explicit nil
### GetCoolingCapacity

`func (o *BulkCoolingSourceRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *BulkCoolingSourceRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *BulkCoolingSourceRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *BulkCoolingSourceRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *BulkCoolingSourceRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *BulkCoolingSourceRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetDescription

`func (o *BulkCoolingSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkCoolingSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkCoolingSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkCoolingSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkCoolingSourceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkCoolingSourceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkCoolingSourceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkCoolingSourceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkCoolingSourceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkCoolingSourceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkCoolingSourceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkCoolingSourceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkCoolingSourceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkCoolingSourceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkCoolingSourceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkCoolingSourceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkCoolingSourceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkCoolingSourceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkCoolingSourceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkCoolingSourceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkCoolingSourceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkCoolingSourceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


