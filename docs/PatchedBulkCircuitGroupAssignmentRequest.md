# PatchedBulkCircuitGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Group** | Pointer to [**BriefCircuitGroupAssignmentSerializerRequestGroup**](BriefCircuitGroupAssignmentSerializerRequestGroup.md) |  | [optional] 
**MemberType** | Pointer to **string** |  | [optional] 
**MemberId** | Pointer to **int64** |  | [optional] 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewPatchedBulkCircuitGroupAssignmentRequest

`func NewPatchedBulkCircuitGroupAssignmentRequest(id int32, ) *PatchedBulkCircuitGroupAssignmentRequest`

NewPatchedBulkCircuitGroupAssignmentRequest instantiates a new PatchedBulkCircuitGroupAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCircuitGroupAssignmentRequestWithDefaults

`func NewPatchedBulkCircuitGroupAssignmentRequestWithDefaults() *PatchedBulkCircuitGroupAssignmentRequest`

NewPatchedBulkCircuitGroupAssignmentRequestWithDefaults instantiates a new PatchedBulkCircuitGroupAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCircuitGroupAssignmentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetGroup

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupAssignmentSerializerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupAssignmentSerializerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkCircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupAssignmentSerializerRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkCircuitGroupAssignmentRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMemberType

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *PatchedBulkCircuitGroupAssignmentRequest) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *PatchedBulkCircuitGroupAssignmentRequest) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetMemberId

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *PatchedBulkCircuitGroupAssignmentRequest) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *PatchedBulkCircuitGroupAssignmentRequest) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetPriority

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedBulkCircuitGroupAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedBulkCircuitGroupAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkCircuitGroupAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkCircuitGroupAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkCircuitGroupAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


