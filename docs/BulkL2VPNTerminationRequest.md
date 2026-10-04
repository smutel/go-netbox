# BulkL2VPNTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**L2vpn** | [**BriefL2VPNTerminationRequestL2vpn**](BriefL2VPNTerminationRequestL2vpn.md) |  | 
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **int64** |  | 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkL2VPNTerminationRequest

`func NewBulkL2VPNTerminationRequest(id int32, l2vpn BriefL2VPNTerminationRequestL2vpn, assignedObjectType string, assignedObjectId int64, ) *BulkL2VPNTerminationRequest`

NewBulkL2VPNTerminationRequest instantiates a new BulkL2VPNTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkL2VPNTerminationRequestWithDefaults

`func NewBulkL2VPNTerminationRequestWithDefaults() *BulkL2VPNTerminationRequest`

NewBulkL2VPNTerminationRequestWithDefaults instantiates a new BulkL2VPNTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkL2VPNTerminationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkL2VPNTerminationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkL2VPNTerminationRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetL2vpn

`func (o *BulkL2VPNTerminationRequest) GetL2vpn() BriefL2VPNTerminationRequestL2vpn`

GetL2vpn returns the L2vpn field if non-nil, zero value otherwise.

### GetL2vpnOk

`func (o *BulkL2VPNTerminationRequest) GetL2vpnOk() (*BriefL2VPNTerminationRequestL2vpn, bool)`

GetL2vpnOk returns a tuple with the L2vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2vpn

`func (o *BulkL2VPNTerminationRequest) SetL2vpn(v BriefL2VPNTerminationRequestL2vpn)`

SetL2vpn sets L2vpn field to given value.


### GetAssignedObjectType

`func (o *BulkL2VPNTerminationRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *BulkL2VPNTerminationRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *BulkL2VPNTerminationRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *BulkL2VPNTerminationRequest) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *BulkL2VPNTerminationRequest) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *BulkL2VPNTerminationRequest) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetTags

`func (o *BulkL2VPNTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkL2VPNTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkL2VPNTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkL2VPNTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkL2VPNTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkL2VPNTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkL2VPNTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkL2VPNTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


