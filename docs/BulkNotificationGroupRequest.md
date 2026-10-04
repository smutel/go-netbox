# BulkNotificationGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBulkNotificationGroupRequest

`func NewBulkNotificationGroupRequest(id int32, name string, ) *BulkNotificationGroupRequest`

NewBulkNotificationGroupRequest instantiates a new BulkNotificationGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkNotificationGroupRequestWithDefaults

`func NewBulkNotificationGroupRequestWithDefaults() *BulkNotificationGroupRequest`

NewBulkNotificationGroupRequestWithDefaults instantiates a new BulkNotificationGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkNotificationGroupRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkNotificationGroupRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkNotificationGroupRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkNotificationGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkNotificationGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkNotificationGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkNotificationGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkNotificationGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkNotificationGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkNotificationGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroups

`func (o *BulkNotificationGroupRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BulkNotificationGroupRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BulkNotificationGroupRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BulkNotificationGroupRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *BulkNotificationGroupRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BulkNotificationGroupRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BulkNotificationGroupRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BulkNotificationGroupRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


