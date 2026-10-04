# BulkBookmarkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**User** | [**BookmarkRequestUser**](BookmarkRequestUser.md) |  | 

## Methods

### NewBulkBookmarkRequest

`func NewBulkBookmarkRequest(id int32, objectType string, objectId int64, user BookmarkRequestUser, ) *BulkBookmarkRequest`

NewBulkBookmarkRequest instantiates a new BulkBookmarkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkBookmarkRequestWithDefaults

`func NewBulkBookmarkRequestWithDefaults() *BulkBookmarkRequest`

NewBulkBookmarkRequestWithDefaults instantiates a new BulkBookmarkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkBookmarkRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkBookmarkRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkBookmarkRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *BulkBookmarkRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkBookmarkRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkBookmarkRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *BulkBookmarkRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BulkBookmarkRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BulkBookmarkRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetUser

`func (o *BulkBookmarkRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BulkBookmarkRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BulkBookmarkRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


