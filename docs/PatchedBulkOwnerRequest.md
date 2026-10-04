# PatchedBulkOwnerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**NullableBulkOwnerRequestGroup**](BulkOwnerRequestGroup.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedBulkOwnerRequest

`func NewPatchedBulkOwnerRequest(id int32, ) *PatchedBulkOwnerRequest`

NewPatchedBulkOwnerRequest instantiates a new PatchedBulkOwnerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkOwnerRequestWithDefaults

`func NewPatchedBulkOwnerRequestWithDefaults() *PatchedBulkOwnerRequest`

NewPatchedBulkOwnerRequestWithDefaults instantiates a new PatchedBulkOwnerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkOwnerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkOwnerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkOwnerRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkOwnerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkOwnerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkOwnerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkOwnerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedBulkOwnerRequest) GetGroup() BulkOwnerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkOwnerRequest) GetGroupOk() (*BulkOwnerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkOwnerRequest) SetGroup(v BulkOwnerRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkOwnerRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedBulkOwnerRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedBulkOwnerRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetDescription

`func (o *PatchedBulkOwnerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkOwnerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkOwnerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkOwnerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserGroups

`func (o *PatchedBulkOwnerRequest) GetUserGroups() []int32`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *PatchedBulkOwnerRequest) GetUserGroupsOk() (*[]int32, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *PatchedBulkOwnerRequest) SetUserGroups(v []int32)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *PatchedBulkOwnerRequest) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsers

`func (o *PatchedBulkOwnerRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PatchedBulkOwnerRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PatchedBulkOwnerRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PatchedBulkOwnerRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


