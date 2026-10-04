# BulkCircuitGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Group** | [**BriefCircuitGroupAssignmentSerializerRequestGroup**](BriefCircuitGroupAssignmentSerializerRequestGroup.md) |  | 
**MemberType** | **string** |  | 
**MemberId** | **int64** |  | 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewBulkCircuitGroupAssignmentRequest

`func NewBulkCircuitGroupAssignmentRequest(id int32, group BriefCircuitGroupAssignmentSerializerRequestGroup, memberType string, memberId int64, ) *BulkCircuitGroupAssignmentRequest`

NewBulkCircuitGroupAssignmentRequest instantiates a new BulkCircuitGroupAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCircuitGroupAssignmentRequestWithDefaults

`func NewBulkCircuitGroupAssignmentRequestWithDefaults() *BulkCircuitGroupAssignmentRequest`

NewBulkCircuitGroupAssignmentRequestWithDefaults instantiates a new BulkCircuitGroupAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCircuitGroupAssignmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCircuitGroupAssignmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCircuitGroupAssignmentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetGroup

`func (o *BulkCircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupAssignmentSerializerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkCircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupAssignmentSerializerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkCircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupAssignmentSerializerRequestGroup)`

SetGroup sets Group field to given value.


### GetMemberType

`func (o *BulkCircuitGroupAssignmentRequest) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *BulkCircuitGroupAssignmentRequest) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *BulkCircuitGroupAssignmentRequest) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.


### GetMemberId

`func (o *BulkCircuitGroupAssignmentRequest) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *BulkCircuitGroupAssignmentRequest) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *BulkCircuitGroupAssignmentRequest) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.


### GetPriority

`func (o *BulkCircuitGroupAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BulkCircuitGroupAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BulkCircuitGroupAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BulkCircuitGroupAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *BulkCircuitGroupAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkCircuitGroupAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkCircuitGroupAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkCircuitGroupAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


