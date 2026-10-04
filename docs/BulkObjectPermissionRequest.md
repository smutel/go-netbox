# BulkObjectPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ObjectTypes** | **[]string** |  | 
**Actions** | **[]string** | The list of actions granted by this permission | 
**Constraints** | Pointer to **interface{}** | Queryset filter matching the applicable objects of the selected type(s) | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBulkObjectPermissionRequest

`func NewBulkObjectPermissionRequest(id int32, name string, objectTypes []string, actions []string, ) *BulkObjectPermissionRequest`

NewBulkObjectPermissionRequest instantiates a new BulkObjectPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkObjectPermissionRequestWithDefaults

`func NewBulkObjectPermissionRequestWithDefaults() *BulkObjectPermissionRequest`

NewBulkObjectPermissionRequestWithDefaults instantiates a new BulkObjectPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkObjectPermissionRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkObjectPermissionRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkObjectPermissionRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkObjectPermissionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkObjectPermissionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkObjectPermissionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkObjectPermissionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkObjectPermissionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkObjectPermissionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkObjectPermissionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BulkObjectPermissionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkObjectPermissionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkObjectPermissionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkObjectPermissionRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectTypes

`func (o *BulkObjectPermissionRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *BulkObjectPermissionRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *BulkObjectPermissionRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetActions

`func (o *BulkObjectPermissionRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BulkObjectPermissionRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BulkObjectPermissionRequest) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetConstraints

`func (o *BulkObjectPermissionRequest) GetConstraints() interface{}`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *BulkObjectPermissionRequest) GetConstraintsOk() (*interface{}, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *BulkObjectPermissionRequest) SetConstraints(v interface{})`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *BulkObjectPermissionRequest) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *BulkObjectPermissionRequest) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *BulkObjectPermissionRequest) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetGroups

`func (o *BulkObjectPermissionRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BulkObjectPermissionRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BulkObjectPermissionRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BulkObjectPermissionRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *BulkObjectPermissionRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BulkObjectPermissionRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BulkObjectPermissionRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BulkObjectPermissionRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


