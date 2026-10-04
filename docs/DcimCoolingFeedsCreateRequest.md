# DcimCoolingFeedsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDcimCoolingFeedsCreateRequest

`func NewDcimCoolingFeedsCreateRequest(coolingSource BulkCoolingFeedRequestCoolingSource, name string, ) *DcimCoolingFeedsCreateRequest`

NewDcimCoolingFeedsCreateRequest instantiates a new DcimCoolingFeedsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimCoolingFeedsCreateRequestWithDefaults

`func NewDcimCoolingFeedsCreateRequestWithDefaults() *DcimCoolingFeedsCreateRequest`

NewDcimCoolingFeedsCreateRequestWithDefaults instantiates a new DcimCoolingFeedsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoolingSource

`func (o *DcimCoolingFeedsCreateRequest) GetCoolingSource() BulkCoolingFeedRequestCoolingSource`

GetCoolingSource returns the CoolingSource field if non-nil, zero value otherwise.

### GetCoolingSourceOk

`func (o *DcimCoolingFeedsCreateRequest) GetCoolingSourceOk() (*BulkCoolingFeedRequestCoolingSource, bool)`

GetCoolingSourceOk returns a tuple with the CoolingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingSource

`func (o *DcimCoolingFeedsCreateRequest) SetCoolingSource(v BulkCoolingFeedRequestCoolingSource)`

SetCoolingSource sets CoolingSource field to given value.


### GetRack

`func (o *DcimCoolingFeedsCreateRequest) GetRack() BulkCoolingFeedRequestRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *DcimCoolingFeedsCreateRequest) GetRackOk() (*BulkCoolingFeedRequestRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *DcimCoolingFeedsCreateRequest) SetRack(v BulkCoolingFeedRequestRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *DcimCoolingFeedsCreateRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *DcimCoolingFeedsCreateRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *DcimCoolingFeedsCreateRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetName

`func (o *DcimCoolingFeedsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimCoolingFeedsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimCoolingFeedsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DcimCoolingFeedsCreateRequest) GetStatus() BulkCoolingFeedRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DcimCoolingFeedsCreateRequest) GetStatusOk() (*BulkCoolingFeedRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DcimCoolingFeedsCreateRequest) SetStatus(v BulkCoolingFeedRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DcimCoolingFeedsCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCoolingCapacity

`func (o *DcimCoolingFeedsCreateRequest) GetCoolingCapacity() float64`

GetCoolingCapacity returns the CoolingCapacity field if non-nil, zero value otherwise.

### GetCoolingCapacityOk

`func (o *DcimCoolingFeedsCreateRequest) GetCoolingCapacityOk() (*float64, bool)`

GetCoolingCapacityOk returns a tuple with the CoolingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingCapacity

`func (o *DcimCoolingFeedsCreateRequest) SetCoolingCapacity(v float64)`

SetCoolingCapacity sets CoolingCapacity field to given value.

### HasCoolingCapacity

`func (o *DcimCoolingFeedsCreateRequest) HasCoolingCapacity() bool`

HasCoolingCapacity returns a boolean if a field has been set.

### SetCoolingCapacityNil

`func (o *DcimCoolingFeedsCreateRequest) SetCoolingCapacityNil(b bool)`

 SetCoolingCapacityNil sets the value for CoolingCapacity to be an explicit nil

### UnsetCoolingCapacity
`func (o *DcimCoolingFeedsCreateRequest) UnsetCoolingCapacity()`

UnsetCoolingCapacity ensures that no value is present for CoolingCapacity, not even an explicit nil
### GetMaxFlow

`func (o *DcimCoolingFeedsCreateRequest) GetMaxFlow() float64`

GetMaxFlow returns the MaxFlow field if non-nil, zero value otherwise.

### GetMaxFlowOk

`func (o *DcimCoolingFeedsCreateRequest) GetMaxFlowOk() (*float64, bool)`

GetMaxFlowOk returns a tuple with the MaxFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlow

`func (o *DcimCoolingFeedsCreateRequest) SetMaxFlow(v float64)`

SetMaxFlow sets MaxFlow field to given value.

### HasMaxFlow

`func (o *DcimCoolingFeedsCreateRequest) HasMaxFlow() bool`

HasMaxFlow returns a boolean if a field has been set.

### SetMaxFlowNil

`func (o *DcimCoolingFeedsCreateRequest) SetMaxFlowNil(b bool)`

 SetMaxFlowNil sets the value for MaxFlow to be an explicit nil

### UnsetMaxFlow
`func (o *DcimCoolingFeedsCreateRequest) UnsetMaxFlow()`

UnsetMaxFlow ensures that no value is present for MaxFlow, not even an explicit nil
### GetMaxFlowUnit

`func (o *DcimCoolingFeedsCreateRequest) GetMaxFlowUnit() BulkCoolingFeedRequestMaxFlowUnit`

GetMaxFlowUnit returns the MaxFlowUnit field if non-nil, zero value otherwise.

### GetMaxFlowUnitOk

`func (o *DcimCoolingFeedsCreateRequest) GetMaxFlowUnitOk() (*BulkCoolingFeedRequestMaxFlowUnit, bool)`

GetMaxFlowUnitOk returns a tuple with the MaxFlowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlowUnit

`func (o *DcimCoolingFeedsCreateRequest) SetMaxFlowUnit(v BulkCoolingFeedRequestMaxFlowUnit)`

SetMaxFlowUnit sets MaxFlowUnit field to given value.

### HasMaxFlowUnit

`func (o *DcimCoolingFeedsCreateRequest) HasMaxFlowUnit() bool`

HasMaxFlowUnit returns a boolean if a field has been set.

### SetMaxFlowUnitNil

`func (o *DcimCoolingFeedsCreateRequest) SetMaxFlowUnitNil(b bool)`

 SetMaxFlowUnitNil sets the value for MaxFlowUnit to be an explicit nil

### UnsetMaxFlowUnit
`func (o *DcimCoolingFeedsCreateRequest) UnsetMaxFlowUnit()`

UnsetMaxFlowUnit ensures that no value is present for MaxFlowUnit, not even an explicit nil
### GetDescription

`func (o *DcimCoolingFeedsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimCoolingFeedsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimCoolingFeedsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimCoolingFeedsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenant

`func (o *DcimCoolingFeedsCreateRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DcimCoolingFeedsCreateRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DcimCoolingFeedsCreateRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DcimCoolingFeedsCreateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *DcimCoolingFeedsCreateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *DcimCoolingFeedsCreateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetOwner

`func (o *DcimCoolingFeedsCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DcimCoolingFeedsCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DcimCoolingFeedsCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DcimCoolingFeedsCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DcimCoolingFeedsCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DcimCoolingFeedsCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *DcimCoolingFeedsCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DcimCoolingFeedsCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DcimCoolingFeedsCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DcimCoolingFeedsCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *DcimCoolingFeedsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DcimCoolingFeedsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DcimCoolingFeedsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DcimCoolingFeedsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DcimCoolingFeedsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DcimCoolingFeedsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DcimCoolingFeedsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DcimCoolingFeedsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


