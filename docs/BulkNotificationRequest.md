# BulkNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**User** | [**BookmarkRequestUser**](BookmarkRequestUser.md) |  | 
**Read** | Pointer to **NullableTime** |  | [optional] 
**EventType** | [**Event**](Event.md) |  | 

## Methods

### NewBulkNotificationRequest

`func NewBulkNotificationRequest(id int32, objectType string, objectId int64, user BookmarkRequestUser, eventType Event, ) *BulkNotificationRequest`

NewBulkNotificationRequest instantiates a new BulkNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkNotificationRequestWithDefaults

`func NewBulkNotificationRequestWithDefaults() *BulkNotificationRequest`

NewBulkNotificationRequestWithDefaults instantiates a new BulkNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkNotificationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkNotificationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkNotificationRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *BulkNotificationRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkNotificationRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkNotificationRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *BulkNotificationRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BulkNotificationRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BulkNotificationRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetUser

`func (o *BulkNotificationRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BulkNotificationRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BulkNotificationRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.


### GetRead

`func (o *BulkNotificationRequest) GetRead() time.Time`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *BulkNotificationRequest) GetReadOk() (*time.Time, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *BulkNotificationRequest) SetRead(v time.Time)`

SetRead sets Read field to given value.

### HasRead

`func (o *BulkNotificationRequest) HasRead() bool`

HasRead returns a boolean if a field has been set.

### SetReadNil

`func (o *BulkNotificationRequest) SetReadNil(b bool)`

 SetReadNil sets the value for Read to be an explicit nil

### UnsetRead
`func (o *BulkNotificationRequest) UnsetRead()`

UnsetRead ensures that no value is present for Read, not even an explicit nil
### GetEventType

`func (o *BulkNotificationRequest) GetEventType() Event`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BulkNotificationRequest) GetEventTypeOk() (*Event, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BulkNotificationRequest) SetEventType(v Event)`

SetEventType sets EventType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


