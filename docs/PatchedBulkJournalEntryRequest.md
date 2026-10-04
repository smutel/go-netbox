# PatchedBulkJournalEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AssignedObjectType** | Pointer to **string** |  | [optional] 
**AssignedObjectId** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] 
**Kind** | Pointer to [**BulkJournalEntryRequestKind**](BulkJournalEntryRequestKind.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkJournalEntryRequest

`func NewPatchedBulkJournalEntryRequest(id int32, ) *PatchedBulkJournalEntryRequest`

NewPatchedBulkJournalEntryRequest instantiates a new PatchedBulkJournalEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkJournalEntryRequestWithDefaults

`func NewPatchedBulkJournalEntryRequestWithDefaults() *PatchedBulkJournalEntryRequest`

NewPatchedBulkJournalEntryRequestWithDefaults instantiates a new PatchedBulkJournalEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkJournalEntryRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkJournalEntryRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkJournalEntryRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetAssignedObjectType

`func (o *PatchedBulkJournalEntryRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *PatchedBulkJournalEntryRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *PatchedBulkJournalEntryRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *PatchedBulkJournalEntryRequest) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### GetAssignedObjectId

`func (o *PatchedBulkJournalEntryRequest) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *PatchedBulkJournalEntryRequest) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *PatchedBulkJournalEntryRequest) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *PatchedBulkJournalEntryRequest) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedBulkJournalEntryRequest) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedBulkJournalEntryRequest) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedBulkJournalEntryRequest) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedBulkJournalEntryRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PatchedBulkJournalEntryRequest) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PatchedBulkJournalEntryRequest) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetKind

`func (o *PatchedBulkJournalEntryRequest) GetKind() BulkJournalEntryRequestKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatchedBulkJournalEntryRequest) GetKindOk() (*BulkJournalEntryRequestKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatchedBulkJournalEntryRequest) SetKind(v BulkJournalEntryRequestKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatchedBulkJournalEntryRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetComments

`func (o *PatchedBulkJournalEntryRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkJournalEntryRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkJournalEntryRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkJournalEntryRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkJournalEntryRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkJournalEntryRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkJournalEntryRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkJournalEntryRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkJournalEntryRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkJournalEntryRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkJournalEntryRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkJournalEntryRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


