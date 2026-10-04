# BulkCoolingFeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CoolingSource** | [**BulkCoolingFeedRequestCoolingSource**](BulkCoolingFeedRequestCoolingSource.md) |  | 
**Rack** | Pointer to [**NullableBulkCoolingFeedRequestRack**](BulkCoolingFeedRequestRack.md) |  | [optional] 
**Name** | **string** |  | 
**Status** | Pointer to [**BulkCoolingFeedRequestStatus**](BulkCoolingFeedRequestStatus.md) |  | [optional] 
**CoolingCapacity** | Pointer to **NullableFloat64** | Rated cooling capacity (kW) | [optional] 
**MaxFlow** | Pointer to **NullableFloat64** |  | [optional] 
**MaxFlowUnit** | Pointer to [**NullableBulkCoolingFeedRequestMaxFlowUnit**](BulkCoolingFeedRequestMaxFlowUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkCoolingFeedRequest

`func NewBulkCoolingFeedRequest(id int32, coolingSource BulkCoolingFeedRequestCoolingSource, name string, ) *BulkCoolingFeedRequest`

NewBulkCoolingFeedRequest instantiates a new BulkCoolingFeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCoolingFeedRequestWithDefaults

`func NewBulkCoolingFeedRequestWithDefaults() *BulkCoolingFeedRequest`

NewBulkCoolingFeedRequestWithDefaults instantiates a new BulkCoolingFeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCoolingFeedRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCoolingFeedRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCoolingFeedRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCoolingSource

`func (o *BulkCoolingFeedRequest) GetCoolingSource() BulkCoolingFeedRequestCoolingSource`

GetCoolingSource returns the CoolingSource field if non-nil, zero value otherwise.

### GetCoolingSourceOk

`func (o *BulkCoolingFeedRequest) GetCoolingSourceOk() (*BulkCoolingFeedRequestCoolingSource, bool)`

GetCoolingSourceOk returns a tuple with the CoolingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingSource

`func (o *BulkCoolingFeedRequest) SetCoolingSource(v BulkCoolingFeedRequestCoolingSource)`

SetCoolingSource sets CoolingSource field to given value.


### GetRack

`func (o *BulkCoolingFeedRequest) GetRack() BulkCoolingFeedRequestRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *BulkCoolingFeedRequest) GetRackOk() (*BulkCoolingFeedRequestRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *BulkCoolingFeedRequest) SetRack(v BulkCoolingFeedRequestRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *BulkCoolingFeedRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *BulkCoolingFeedRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *BulkCoolingFeedRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetName

`func (o *BulkCoolingFeedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkCoolingFeedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkCoolingFeedRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *BulkCoolingFeedRequest) GetStatus() BulkCoolingFeedRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkCoolingFeedRequest) GetStatusOk() (*BulkCoolingFeedRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkCoolingFeedRequest) SetStatus(v BulkCoolingFeedRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkCoolingFeedRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCoolingCapacity

`func (o *BulkCoolingFeedRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *BulkCoolingFeedRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *BulkCoolingFeedRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *BulkCoolingFeedRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *BulkCoolingFeedRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *BulkCoolingFeedRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetMaxFlow

`func (o *BulkCoolingFeedRequest) GetMaxFlow() float64`

GetMaxFlow returns the MaxFlow field if non-nil, zero value otherwise.

### GetMaxFlowOk

`func (o *BulkCoolingFeedRequest) GetMaxFlowOk() (*float64, bool)`

GetMaxFlowOk returns a tuple with the MaxFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlow

`func (o *BulkCoolingFeedRequest) SetMaxFlow(v float64)`

SetMaxFlow sets MaxFlow field to given value.

### HasMaxFlow

`func (o *BulkCoolingFeedRequest) HasMaxFlow() bool`

HasMaxFlow returns a boolean if a field has been set.

### SetMaxFlowNil

`func (o *BulkCoolingFeedRequest) SetMaxFlowNil(b bool)`

 SetMaxFlowNil sets the value for MaxFlow to be an explicit nil

### UnsetMaxFlow
`func (o *BulkCoolingFeedRequest) UnsetMaxFlow()`

UnsetMaxFlow ensures that no value is present for MaxFlow, not even an explicit nil
### GetMaxFlowUnit

`func (o *BulkCoolingFeedRequest) GetMaxFlowUnit() BulkCoolingFeedRequestMaxFlowUnit`

GetMaxFlowUnit returns the MaxFlowUnit field if non-nil, zero value otherwise.

### GetMaxFlowUnitOk

`func (o *BulkCoolingFeedRequest) GetMaxFlowUnitOk() (*BulkCoolingFeedRequestMaxFlowUnit, bool)`

GetMaxFlowUnitOk returns a tuple with the MaxFlowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlowUnit

`func (o *BulkCoolingFeedRequest) SetMaxFlowUnit(v BulkCoolingFeedRequestMaxFlowUnit)`

SetMaxFlowUnit sets MaxFlowUnit field to given value.

### HasMaxFlowUnit

`func (o *BulkCoolingFeedRequest) HasMaxFlowUnit() bool`

HasMaxFlowUnit returns a boolean if a field has been set.

### SetMaxFlowUnitNil

`func (o *BulkCoolingFeedRequest) SetMaxFlowUnitNil(b bool)`

 SetMaxFlowUnitNil sets the value for MaxFlowUnit to be an explicit nil

### UnsetMaxFlowUnit
`func (o *BulkCoolingFeedRequest) UnsetMaxFlowUnit()`

UnsetMaxFlowUnit ensures that no value is present for MaxFlowUnit, not even an explicit nil
### GetDescription

`func (o *BulkCoolingFeedRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkCoolingFeedRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkCoolingFeedRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkCoolingFeedRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenant

`func (o *BulkCoolingFeedRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkCoolingFeedRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkCoolingFeedRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkCoolingFeedRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkCoolingFeedRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkCoolingFeedRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetOwner

`func (o *BulkCoolingFeedRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkCoolingFeedRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkCoolingFeedRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkCoolingFeedRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkCoolingFeedRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkCoolingFeedRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkCoolingFeedRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkCoolingFeedRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkCoolingFeedRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkCoolingFeedRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkCoolingFeedRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkCoolingFeedRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkCoolingFeedRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkCoolingFeedRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkCoolingFeedRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkCoolingFeedRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkCoolingFeedRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkCoolingFeedRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


