# PatchedBulkFHRPGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Group** | Pointer to [**BulkFHRPGroupAssignmentRequestGroup**](BulkFHRPGroupAssignmentRequestGroup.md) |  | [optional] 
**InterfaceType** | Pointer to **string** |  | [optional] 
**InterfaceId** | Pointer to **int64** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedBulkFHRPGroupAssignmentRequest

`func NewPatchedBulkFHRPGroupAssignmentRequest(id int32, ) *PatchedBulkFHRPGroupAssignmentRequest`

NewPatchedBulkFHRPGroupAssignmentRequest instantiates a new PatchedBulkFHRPGroupAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkFHRPGroupAssignmentRequestWithDefaults

`func NewPatchedBulkFHRPGroupAssignmentRequestWithDefaults() *PatchedBulkFHRPGroupAssignmentRequest`

NewPatchedBulkFHRPGroupAssignmentRequestWithDefaults instantiates a new PatchedBulkFHRPGroupAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkFHRPGroupAssignmentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetGroup

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetGroup() BulkFHRPGroupAssignmentRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetGroupOk() (*BulkFHRPGroupAssignmentRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedBulkFHRPGroupAssignmentRequest) SetGroup(v BulkFHRPGroupAssignmentRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedBulkFHRPGroupAssignmentRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInterfaceType

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *PatchedBulkFHRPGroupAssignmentRequest) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *PatchedBulkFHRPGroupAssignmentRequest) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetInterfaceId

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *PatchedBulkFHRPGroupAssignmentRequest) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *PatchedBulkFHRPGroupAssignmentRequest) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetPriority

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedBulkFHRPGroupAssignmentRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedBulkFHRPGroupAssignmentRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedBulkFHRPGroupAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


