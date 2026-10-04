# PatchedBulkObjectPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Actions** | Pointer to **[]string** | The list of actions granted by this permission | [optional] 
**Constraints** | Pointer to **interface{}** | Queryset filter matching the applicable objects of the selected type(s) | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedBulkObjectPermissionRequest

`func NewPatchedBulkObjectPermissionRequest(id int32, ) *PatchedBulkObjectPermissionRequest`

NewPatchedBulkObjectPermissionRequest instantiates a new PatchedBulkObjectPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkObjectPermissionRequestWithDefaults

`func NewPatchedBulkObjectPermissionRequestWithDefaults() *PatchedBulkObjectPermissionRequest`

NewPatchedBulkObjectPermissionRequestWithDefaults instantiates a new PatchedBulkObjectPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkObjectPermissionRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkObjectPermissionRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkObjectPermissionRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkObjectPermissionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkObjectPermissionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkObjectPermissionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkObjectPermissionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkObjectPermissionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkObjectPermissionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkObjectPermissionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkObjectPermissionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkObjectPermissionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkObjectPermissionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkObjectPermissionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkObjectPermissionRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectTypes

`func (o *PatchedBulkObjectPermissionRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedBulkObjectPermissionRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedBulkObjectPermissionRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedBulkObjectPermissionRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetActions

`func (o *PatchedBulkObjectPermissionRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PatchedBulkObjectPermissionRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PatchedBulkObjectPermissionRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PatchedBulkObjectPermissionRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConstraints

`func (o *PatchedBulkObjectPermissionRequest) GetConstraints() interface{}`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *PatchedBulkObjectPermissionRequest) GetConstraintsOk() (*interface{}, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *PatchedBulkObjectPermissionRequest) SetConstraints(v interface{})`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *PatchedBulkObjectPermissionRequest) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *PatchedBulkObjectPermissionRequest) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *PatchedBulkObjectPermissionRequest) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetGroups

`func (o *PatchedBulkObjectPermissionRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PatchedBulkObjectPermissionRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PatchedBulkObjectPermissionRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PatchedBulkObjectPermissionRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *PatchedBulkObjectPermissionRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PatchedBulkObjectPermissionRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PatchedBulkObjectPermissionRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PatchedBulkObjectPermissionRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


