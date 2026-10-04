# BulkOwnerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Group** | [**NullableBulkOwnerRequestGroup**](BulkOwnerRequestGroup.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBulkOwnerRequest

`func NewBulkOwnerRequest(id int32, name string, group NullableBulkOwnerRequestGroup, ) *BulkOwnerRequest`

NewBulkOwnerRequest instantiates a new BulkOwnerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOwnerRequestWithDefaults

`func NewBulkOwnerRequestWithDefaults() *BulkOwnerRequest`

NewBulkOwnerRequestWithDefaults instantiates a new BulkOwnerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkOwnerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkOwnerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkOwnerRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkOwnerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkOwnerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkOwnerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetGroup

`func (o *BulkOwnerRequest) GetGroup() BulkOwnerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkOwnerRequest) GetGroupOk() (*BulkOwnerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkOwnerRequest) SetGroup(v BulkOwnerRequestGroup)`

SetGroup sets Group field to given value.


### SetGroupNil

`func (o *BulkOwnerRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BulkOwnerRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetDescription

`func (o *BulkOwnerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkOwnerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkOwnerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkOwnerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserGroups

`func (o *BulkOwnerRequest) GetUserGroups() []int32`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *BulkOwnerRequest) GetUserGroupsOk() (*[]int32, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *BulkOwnerRequest) SetUserGroups(v []int32)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *BulkOwnerRequest) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsers

`func (o *BulkOwnerRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BulkOwnerRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BulkOwnerRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BulkOwnerRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


