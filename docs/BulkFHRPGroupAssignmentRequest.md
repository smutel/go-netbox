# BulkFHRPGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Group** | [**BulkFHRPGroupAssignmentRequestGroup**](BulkFHRPGroupAssignmentRequestGroup.md) |  | 
**InterfaceType** | **string** |  | 
**InterfaceId** | **int64** |  | 
**Priority** | **int32** |  | 

## Methods

### NewBulkFHRPGroupAssignmentRequest

`func NewBulkFHRPGroupAssignmentRequest(id int32, group BulkFHRPGroupAssignmentRequestGroup, interfaceType string, interfaceId int64, priority int32, ) *BulkFHRPGroupAssignmentRequest`

NewBulkFHRPGroupAssignmentRequest instantiates a new BulkFHRPGroupAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkFHRPGroupAssignmentRequestWithDefaults

`func NewBulkFHRPGroupAssignmentRequestWithDefaults() *BulkFHRPGroupAssignmentRequest`

NewBulkFHRPGroupAssignmentRequestWithDefaults instantiates a new BulkFHRPGroupAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkFHRPGroupAssignmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkFHRPGroupAssignmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkFHRPGroupAssignmentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetGroup

`func (o *BulkFHRPGroupAssignmentRequest) GetGroup() BulkFHRPGroupAssignmentRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BulkFHRPGroupAssignmentRequest) GetGroupOk() (*BulkFHRPGroupAssignmentRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BulkFHRPGroupAssignmentRequest) SetGroup(v BulkFHRPGroupAssignmentRequestGroup)`

SetGroup sets Group field to given value.


### GetInterfaceType

`func (o *BulkFHRPGroupAssignmentRequest) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *BulkFHRPGroupAssignmentRequest) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *BulkFHRPGroupAssignmentRequest) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.


### GetInterfaceId

`func (o *BulkFHRPGroupAssignmentRequest) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *BulkFHRPGroupAssignmentRequest) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *BulkFHRPGroupAssignmentRequest) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.


### GetPriority

`func (o *BulkFHRPGroupAssignmentRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BulkFHRPGroupAssignmentRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BulkFHRPGroupAssignmentRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


