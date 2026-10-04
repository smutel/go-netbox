# BulkContactAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Contact** | [**BulkContactAssignmentRequestContact**](BulkContactAssignmentRequestContact.md) |  | 
**Role** | Pointer to [**NullableBulkContactAssignmentRequestRole**](BulkContactAssignmentRequestRole.md) |  | [optional] 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkContactAssignmentRequest

`func NewBulkContactAssignmentRequest(id int32, objectType string, objectId int64, contact BulkContactAssignmentRequestContact, ) *BulkContactAssignmentRequest`

NewBulkContactAssignmentRequest instantiates a new BulkContactAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkContactAssignmentRequestWithDefaults

`func NewBulkContactAssignmentRequestWithDefaults() *BulkContactAssignmentRequest`

NewBulkContactAssignmentRequestWithDefaults instantiates a new BulkContactAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkContactAssignmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkContactAssignmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkContactAssignmentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *BulkContactAssignmentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkContactAssignmentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkContactAssignmentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *BulkContactAssignmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BulkContactAssignmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BulkContactAssignmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetContact

`func (o *BulkContactAssignmentRequest) GetContact() BulkContactAssignmentRequestContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *BulkContactAssignmentRequest) GetContactOk() (*BulkContactAssignmentRequestContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *BulkContactAssignmentRequest) SetContact(v BulkContactAssignmentRequestContact)`

SetContact sets Contact field to given value.


### GetRole

`func (o *BulkContactAssignmentRequest) GetRole() BulkContactAssignmentRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkContactAssignmentRequest) GetRoleOk() (*BulkContactAssignmentRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkContactAssignmentRequest) SetRole(v BulkContactAssignmentRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkContactAssignmentRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkContactAssignmentRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkContactAssignmentRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPriority

`func (o *BulkContactAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BulkContactAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BulkContactAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BulkContactAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *BulkContactAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkContactAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkContactAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkContactAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkContactAssignmentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkContactAssignmentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkContactAssignmentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkContactAssignmentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


