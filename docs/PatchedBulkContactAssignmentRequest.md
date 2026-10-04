# PatchedBulkContactAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**Contact** | Pointer to [**BulkContactAssignmentRequestContact**](BulkContactAssignmentRequestContact.md) |  | [optional] 
**Role** | Pointer to [**NullableBulkContactAssignmentRequestRole**](BulkContactAssignmentRequestRole.md) |  | [optional] 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkContactAssignmentRequest

`func NewPatchedBulkContactAssignmentRequest(id int32, ) *PatchedBulkContactAssignmentRequest`

NewPatchedBulkContactAssignmentRequest instantiates a new PatchedBulkContactAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkContactAssignmentRequestWithDefaults

`func NewPatchedBulkContactAssignmentRequestWithDefaults() *PatchedBulkContactAssignmentRequest`

NewPatchedBulkContactAssignmentRequestWithDefaults instantiates a new PatchedBulkContactAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkContactAssignmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkContactAssignmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkContactAssignmentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PatchedBulkContactAssignmentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedBulkContactAssignmentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedBulkContactAssignmentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedBulkContactAssignmentRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedBulkContactAssignmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedBulkContactAssignmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedBulkContactAssignmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedBulkContactAssignmentRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetContact

`func (o *PatchedBulkContactAssignmentRequest) GetContact() BulkContactAssignmentRequestContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatchedBulkContactAssignmentRequest) GetContactOk() (*BulkContactAssignmentRequestContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatchedBulkContactAssignmentRequest) SetContact(v BulkContactAssignmentRequestContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatchedBulkContactAssignmentRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkContactAssignmentRequest) GetRole() BulkContactAssignmentRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkContactAssignmentRequest) GetRoleOk() (*BulkContactAssignmentRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkContactAssignmentRequest) SetRole(v BulkContactAssignmentRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkContactAssignmentRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkContactAssignmentRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkContactAssignmentRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPriority

`func (o *PatchedBulkContactAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedBulkContactAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedBulkContactAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedBulkContactAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkContactAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkContactAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkContactAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkContactAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkContactAssignmentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkContactAssignmentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkContactAssignmentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkContactAssignmentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


