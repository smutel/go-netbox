# PatchedBulkNotificationGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedBulkNotificationGroupRequest

`func NewPatchedBulkNotificationGroupRequest(id int32, ) *PatchedBulkNotificationGroupRequest`

NewPatchedBulkNotificationGroupRequest instantiates a new PatchedBulkNotificationGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkNotificationGroupRequestWithDefaults

`func NewPatchedBulkNotificationGroupRequestWithDefaults() *PatchedBulkNotificationGroupRequest`

NewPatchedBulkNotificationGroupRequestWithDefaults instantiates a new PatchedBulkNotificationGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkNotificationGroupRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkNotificationGroupRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkNotificationGroupRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkNotificationGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkNotificationGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkNotificationGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkNotificationGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkNotificationGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkNotificationGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkNotificationGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkNotificationGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroups

`func (o *PatchedBulkNotificationGroupRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PatchedBulkNotificationGroupRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PatchedBulkNotificationGroupRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PatchedBulkNotificationGroupRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *PatchedBulkNotificationGroupRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PatchedBulkNotificationGroupRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PatchedBulkNotificationGroupRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PatchedBulkNotificationGroupRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


