# CoolingFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**CoolingSource** | [**BriefCoolingSource**](BriefCoolingSource.md) |  | 
**Rack** | Pointer to [**NullableBriefRack**](BriefRack.md) |  | [optional] 
**Name** | **string** |  | 
**Status** | Pointer to [**CoolingFeedStatus**](CoolingFeedStatus.md) |  | [optional] 
**CoolingCapacity** | Pointer to **NullableFloat64** | Rated cooling capacity (kW) | [optional] 
**MaxFlow** | Pointer to **NullableFloat64** |  | [optional] 
**MaxFlowUnit** | Pointer to [**NullableCoolingFeedMaxFlowUnit**](CoolingFeedMaxFlowUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenant**](BriefTenant.md) |  | [optional] 
**Owner** | Pointer to [**NullableBriefOwner**](BriefOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewCoolingFeed

`func NewCoolingFeed(id int32, url string, displayUrl string, display string, coolingSource BriefCoolingSource, name string, created NullableTime, lastUpdated NullableTime, ) *CoolingFeed`

NewCoolingFeed instantiates a new CoolingFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoolingFeedWithDefaults

`func NewCoolingFeedWithDefaults() *CoolingFeed`

NewCoolingFeedWithDefaults instantiates a new CoolingFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoolingFeed) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoolingFeed) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoolingFeed) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CoolingFeed) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CoolingFeed) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CoolingFeed) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *CoolingFeed) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *CoolingFeed) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *CoolingFeed) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *CoolingFeed) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CoolingFeed) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CoolingFeed) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetCoolingSource

`func (o *CoolingFeed) GetCoolingSource() BriefCoolingSource`

GetCoolingSource returns the CoolingSource field if non-nil, zero value otherwise.

### GetCoolingSourceOk

`func (o *CoolingFeed) GetCoolingSourceOk() (*BriefCoolingSource, bool)`

GetCoolingSourceOk returns a tuple with the CoolingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingSource

`func (o *CoolingFeed) SetCoolingSource(v BriefCoolingSource)`

SetCoolingSource sets CoolingSource field to given value.


### GetRack

`func (o *CoolingFeed) GetRack() BriefRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *CoolingFeed) GetRackOk() (*BriefRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *CoolingFeed) SetRack(v BriefRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *CoolingFeed) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *CoolingFeed) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *CoolingFeed) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetName

`func (o *CoolingFeed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoolingFeed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoolingFeed) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *CoolingFeed) GetStatus() CoolingFeedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoolingFeed) GetStatusOk() (*CoolingFeedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoolingFeed) SetStatus(v CoolingFeedStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoolingFeed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCoolingCapacity

`func (o *CoolingFeed) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *CoolingFeed) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *CoolingFeed) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *CoolingFeed) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *CoolingFeed) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *CoolingFeed) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetMaxFlow

`func (o *CoolingFeed) GetMaxFlow() float64`

GetMaxFlow returns the MaxFlow field if non-nil, zero value otherwise.

### GetMaxFlowOk

`func (o *CoolingFeed) GetMaxFlowOk() (*float64, bool)`

GetMaxFlowOk returns a tuple with the MaxFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlow

`func (o *CoolingFeed) SetMaxFlow(v float64)`

SetMaxFlow sets MaxFlow field to given value.

### HasMaxFlow

`func (o *CoolingFeed) HasMaxFlow() bool`

HasMaxFlow returns a boolean if a field has been set.

### SetMaxFlowNil

`func (o *CoolingFeed) SetMaxFlowNil(b bool)`

 SetMaxFlowNil sets the value for MaxFlow to be an explicit nil

### UnsetMaxFlow
`func (o *CoolingFeed) UnsetMaxFlow()`

UnsetMaxFlow ensures that no value is present for MaxFlow, not even an explicit nil
### GetMaxFlowUnit

`func (o *CoolingFeed) GetMaxFlowUnit() CoolingFeedMaxFlowUnit`

GetMaxFlowUnit returns the MaxFlowUnit field if non-nil, zero value otherwise.

### GetMaxFlowUnitOk

`func (o *CoolingFeed) GetMaxFlowUnitOk() (*CoolingFeedMaxFlowUnit, bool)`

GetMaxFlowUnitOk returns a tuple with the MaxFlowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlowUnit

`func (o *CoolingFeed) SetMaxFlowUnit(v CoolingFeedMaxFlowUnit)`

SetMaxFlowUnit sets MaxFlowUnit field to given value.

### HasMaxFlowUnit

`func (o *CoolingFeed) HasMaxFlowUnit() bool`

HasMaxFlowUnit returns a boolean if a field has been set.

### SetMaxFlowUnitNil

`func (o *CoolingFeed) SetMaxFlowUnitNil(b bool)`

 SetMaxFlowUnitNil sets the value for MaxFlowUnit to be an explicit nil

### UnsetMaxFlowUnit
`func (o *CoolingFeed) UnsetMaxFlowUnit()`

UnsetMaxFlowUnit ensures that no value is present for MaxFlowUnit, not even an explicit nil
### GetDescription

`func (o *CoolingFeed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoolingFeed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoolingFeed) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoolingFeed) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenant

`func (o *CoolingFeed) GetTenant() BriefTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoolingFeed) GetTenantOk() (*BriefTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoolingFeed) SetTenant(v BriefTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoolingFeed) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *CoolingFeed) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *CoolingFeed) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetOwner

`func (o *CoolingFeed) GetOwner() BriefOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CoolingFeed) GetOwnerOk() (*BriefOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CoolingFeed) SetOwner(v BriefOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CoolingFeed) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *CoolingFeed) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *CoolingFeed) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *CoolingFeed) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CoolingFeed) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CoolingFeed) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CoolingFeed) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *CoolingFeed) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CoolingFeed) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CoolingFeed) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CoolingFeed) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CoolingFeed) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CoolingFeed) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CoolingFeed) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CoolingFeed) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *CoolingFeed) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CoolingFeed) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CoolingFeed) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CoolingFeed) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CoolingFeed) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CoolingFeed) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CoolingFeed) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CoolingFeed) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CoolingFeed) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CoolingFeed) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


