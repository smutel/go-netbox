# PatchedBulkSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**BookmarkRequestUser**](BookmarkRequestUser.md) |  | [optional] 

## Methods

### NewPatchedBulkSubscriptionRequest

`func NewPatchedBulkSubscriptionRequest(id int32, ) *PatchedBulkSubscriptionRequest`

NewPatchedBulkSubscriptionRequest instantiates a new PatchedBulkSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkSubscriptionRequestWithDefaults

`func NewPatchedBulkSubscriptionRequestWithDefaults() *PatchedBulkSubscriptionRequest`

NewPatchedBulkSubscriptionRequestWithDefaults instantiates a new PatchedBulkSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkSubscriptionRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkSubscriptionRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkSubscriptionRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PatchedBulkSubscriptionRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedBulkSubscriptionRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedBulkSubscriptionRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedBulkSubscriptionRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedBulkSubscriptionRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedBulkSubscriptionRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedBulkSubscriptionRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedBulkSubscriptionRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUser

`func (o *PatchedBulkSubscriptionRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkSubscriptionRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkSubscriptionRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkSubscriptionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


