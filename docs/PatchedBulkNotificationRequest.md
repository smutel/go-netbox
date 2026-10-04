# PatchedBulkNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**BookmarkRequestUser**](BookmarkRequestUser.md) |  | [optional] 
**Read** | Pointer to **NullableTime** |  | [optional] 
**EventType** | Pointer to [**Event**](Event.md) |  | [optional] 

## Methods

### NewPatchedBulkNotificationRequest

`func NewPatchedBulkNotificationRequest(id int32, ) *PatchedBulkNotificationRequest`

NewPatchedBulkNotificationRequest instantiates a new PatchedBulkNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkNotificationRequestWithDefaults

`func NewPatchedBulkNotificationRequestWithDefaults() *PatchedBulkNotificationRequest`

NewPatchedBulkNotificationRequestWithDefaults instantiates a new PatchedBulkNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkNotificationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkNotificationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkNotificationRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PatchedBulkNotificationRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedBulkNotificationRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedBulkNotificationRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedBulkNotificationRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedBulkNotificationRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedBulkNotificationRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedBulkNotificationRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedBulkNotificationRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUser

`func (o *PatchedBulkNotificationRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkNotificationRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkNotificationRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkNotificationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRead

`func (o *PatchedBulkNotificationRequest) GetRead() time.Time`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *PatchedBulkNotificationRequest) GetReadOk() (*time.Time, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *PatchedBulkNotificationRequest) SetRead(v time.Time)`

SetRead sets Read field to given value.

### HasRead

`func (o *PatchedBulkNotificationRequest) HasRead() bool`

HasRead returns a boolean if a field has been set.

### SetReadNil

`func (o *PatchedBulkNotificationRequest) SetReadNil(b bool)`

 SetReadNil sets the value for Read to be an explicit nil

### UnsetRead
`func (o *PatchedBulkNotificationRequest) UnsetRead()`

UnsetRead ensures that no value is present for Read, not even an explicit nil
### GetEventType

`func (o *PatchedBulkNotificationRequest) GetEventType() Event`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PatchedBulkNotificationRequest) GetEventTypeOk() (*Event, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PatchedBulkNotificationRequest) SetEventType(v Event)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *PatchedBulkNotificationRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


