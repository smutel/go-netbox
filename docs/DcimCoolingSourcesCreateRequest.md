# DcimCoolingSourcesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDcimCoolingSourcesCreateRequest

`func NewDcimCoolingSourcesCreateRequest(site BulkCoolingSourceRequestSite, name string, type_ BulkCoolingSourceRequestType, ) *DcimCoolingSourcesCreateRequest`

NewDcimCoolingSourcesCreateRequest instantiates a new DcimCoolingSourcesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimCoolingSourcesCreateRequestWithDefaults

`func NewDcimCoolingSourcesCreateRequestWithDefaults() *DcimCoolingSourcesCreateRequest`

NewDcimCoolingSourcesCreateRequestWithDefaults instantiates a new DcimCoolingSourcesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *DcimCoolingSourcesCreateRequest) GetSite() BulkCoolingSourceRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DcimCoolingSourcesCreateRequest) GetSiteOk() (*BulkCoolingSourceRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DcimCoolingSourcesCreateRequest) SetSite(v BulkCoolingSourceRequestSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *DcimCoolingSourcesCreateRequest) GetLocation() BulkCoolingSourceRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DcimCoolingSourcesCreateRequest) GetLocationOk() (*BulkCoolingSourceRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DcimCoolingSourcesCreateRequest) SetLocation(v BulkCoolingSourceRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DcimCoolingSourcesCreateRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DcimCoolingSourcesCreateRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DcimCoolingSourcesCreateRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *DcimCoolingSourcesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimCoolingSourcesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimCoolingSourcesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DcimCoolingSourcesCreateRequest) GetType() BulkCoolingSourceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcimCoolingSourcesCreateRequest) GetTypeOk() (*BulkCoolingSourceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcimCoolingSourcesCreateRequest) SetType(v BulkCoolingSourceRequestType)`

SetType sets Type field to given value.


### GetStatus

`func (o *DcimCoolingSourcesCreateRequest) GetStatus() BulkCoolingFeedRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DcimCoolingSourcesCreateRequest) GetStatusOk() (*BulkCoolingFeedRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DcimCoolingSourcesCreateRequest) SetStatus(v BulkCoolingFeedRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DcimCoolingSourcesCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFluidType

`func (o *DcimCoolingSourcesCreateRequest) GetFluidType() BulkCoolingSourceRequestFluidType`

GetFluidType returns the FluidType field if non-nil, zero value otherwise.

### GetFluidTypeOk

`func (o *DcimCoolingSourcesCreateRequest) GetFluidTypeOk() (*BulkCoolingSourceRequestFluidType, bool)`

GetFluidTypeOk returns a tuple with the FluidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFluidType

`func (o *DcimCoolingSourcesCreateRequest) SetFluidType(v BulkCoolingSourceRequestFluidType)`

SetFluidType sets FluidType field to given value.

### HasFluidType

`func (o *DcimCoolingSourcesCreateRequest) HasFluidType() bool`

HasFluidType returns a boolean if a field has been set.

### SetFluidTypeNil

`func (o *DcimCoolingSourcesCreateRequest) SetFluidTypeNil(b bool)`

 SetFluidTypeNil sets the value for FluidType to be an explicit nil

### UnsetFluidType
`func (o *DcimCoolingSourcesCreateRequest) UnsetFluidType()`

UnsetFluidType ensures that no value is present for FluidType, not even an explicit nil
### GetCoolingCapacity

`func (o *DcimCoolingSourcesCreateRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *DcimCoolingSourcesCreateRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *DcimCoolingSourcesCreateRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *DcimCoolingSourcesCreateRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *DcimCoolingSourcesCreateRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *DcimCoolingSourcesCreateRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetDescription

`func (o *DcimCoolingSourcesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimCoolingSourcesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimCoolingSourcesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimCoolingSourcesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *DcimCoolingSourcesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DcimCoolingSourcesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DcimCoolingSourcesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DcimCoolingSourcesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DcimCoolingSourcesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DcimCoolingSourcesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *DcimCoolingSourcesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DcimCoolingSourcesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DcimCoolingSourcesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DcimCoolingSourcesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *DcimCoolingSourcesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DcimCoolingSourcesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DcimCoolingSourcesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DcimCoolingSourcesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DcimCoolingSourcesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DcimCoolingSourcesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DcimCoolingSourcesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DcimCoolingSourcesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


