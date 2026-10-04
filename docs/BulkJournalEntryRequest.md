# BulkJournalEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **int64** |  | 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] 
**Kind** | Pointer to [**BulkJournalEntryRequestKind**](BulkJournalEntryRequestKind.md) |  | [optional] 
**Comments** | **string** |  | 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkJournalEntryRequest

`func NewBulkJournalEntryRequest(id int32, assignedObjectType string, assignedObjectId int64, comments string, ) *BulkJournalEntryRequest`

NewBulkJournalEntryRequest instantiates a new BulkJournalEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJournalEntryRequestWithDefaults

`func NewBulkJournalEntryRequestWithDefaults() *BulkJournalEntryRequest`

NewBulkJournalEntryRequestWithDefaults instantiates a new BulkJournalEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkJournalEntryRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJournalEntryRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJournalEntryRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetAssignedObjectType

`func (o *BulkJournalEntryRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *BulkJournalEntryRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *BulkJournalEntryRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *BulkJournalEntryRequest) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *BulkJournalEntryRequest) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *BulkJournalEntryRequest) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetCreatedBy

`func (o *BulkJournalEntryRequest) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BulkJournalEntryRequest) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BulkJournalEntryRequest) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BulkJournalEntryRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *BulkJournalEntryRequest) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *BulkJournalEntryRequest) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetKind

`func (o *BulkJournalEntryRequest) GetKind() BulkJournalEntryRequestKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BulkJournalEntryRequest) GetKindOk() (*BulkJournalEntryRequestKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BulkJournalEntryRequest) SetKind(v BulkJournalEntryRequestKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BulkJournalEntryRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetComments

`func (o *BulkJournalEntryRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkJournalEntryRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkJournalEntryRequest) SetComments(v string)`

SetComments sets Comments field to given value.


### GetTags

`func (o *BulkJournalEntryRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkJournalEntryRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkJournalEntryRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkJournalEntryRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkJournalEntryRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkJournalEntryRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkJournalEntryRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkJournalEntryRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


