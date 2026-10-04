# PatchedBulkBookmarkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**BookmarkRequestUser**](BookmarkRequestUser.md) |  | [optional] 

## Methods

### NewPatchedBulkBookmarkRequest

`func NewPatchedBulkBookmarkRequest(id int32, ) *PatchedBulkBookmarkRequest`

NewPatchedBulkBookmarkRequest instantiates a new PatchedBulkBookmarkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkBookmarkRequestWithDefaults

`func NewPatchedBulkBookmarkRequestWithDefaults() *PatchedBulkBookmarkRequest`

NewPatchedBulkBookmarkRequestWithDefaults instantiates a new PatchedBulkBookmarkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkBookmarkRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkBookmarkRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkBookmarkRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PatchedBulkBookmarkRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedBulkBookmarkRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedBulkBookmarkRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedBulkBookmarkRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedBulkBookmarkRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedBulkBookmarkRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedBulkBookmarkRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedBulkBookmarkRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUser

`func (o *PatchedBulkBookmarkRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkBookmarkRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkBookmarkRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkBookmarkRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


