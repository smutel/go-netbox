# BulkGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBulkGroupRequest

`func NewBulkGroupRequest(id int32, name string, ) *BulkGroupRequest`

NewBulkGroupRequest instantiates a new BulkGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkGroupRequestWithDefaults

`func NewBulkGroupRequestWithDefaults() *BulkGroupRequest`

NewBulkGroupRequestWithDefaults instantiates a new BulkGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkGroupRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkGroupRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkGroupRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *BulkGroupRequest) GetPermissions() []int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *BulkGroupRequest) GetPermissionsOk() (*[]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *BulkGroupRequest) SetPermissions(v []int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *BulkGroupRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


