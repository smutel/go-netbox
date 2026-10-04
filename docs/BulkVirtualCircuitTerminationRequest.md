# BulkVirtualCircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**VirtualCircuit** | [**BulkVirtualCircuitTerminationRequestVirtualCircuit**](BulkVirtualCircuitTerminationRequestVirtualCircuit.md) |  | 
**Role** | Pointer to [**BulkTunnelTerminationRequestRole**](BulkTunnelTerminationRequestRole.md) |  | [optional] 
**Interface** | [**BulkVirtualCircuitTerminationRequestInterface**](BulkVirtualCircuitTerminationRequestInterface.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkVirtualCircuitTerminationRequest

`func NewBulkVirtualCircuitTerminationRequest(id int32, virtualCircuit BulkVirtualCircuitTerminationRequestVirtualCircuit, interface_ BulkVirtualCircuitTerminationRequestInterface, ) *BulkVirtualCircuitTerminationRequest`

NewBulkVirtualCircuitTerminationRequest instantiates a new BulkVirtualCircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVirtualCircuitTerminationRequestWithDefaults

`func NewBulkVirtualCircuitTerminationRequestWithDefaults() *BulkVirtualCircuitTerminationRequest`

NewBulkVirtualCircuitTerminationRequestWithDefaults instantiates a new BulkVirtualCircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkVirtualCircuitTerminationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkVirtualCircuitTerminationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkVirtualCircuitTerminationRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetVirtualCircuit

`func (o *BulkVirtualCircuitTerminationRequest) GetVirtualCircuit() BulkVirtualCircuitTerminationRequestVirtualCircuit`

GetVirtualCircuit returns the VirtualCircuit field if non-nil, zero value otherwise.

### GetVirtualCircuitOk

`func (o *BulkVirtualCircuitTerminationRequest) GetVirtualCircuitOk() (*BulkVirtualCircuitTerminationRequestVirtualCircuit, bool)`

GetVirtualCircuitOk returns a tuple with the VirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuit

`func (o *BulkVirtualCircuitTerminationRequest) SetVirtualCircuit(v BulkVirtualCircuitTerminationRequestVirtualCircuit)`

SetVirtualCircuit sets VirtualCircuit field to given value.


### GetRole

`func (o *BulkVirtualCircuitTerminationRequest) GetRole() BulkTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkVirtualCircuitTerminationRequest) GetRoleOk() (*BulkTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkVirtualCircuitTerminationRequest) SetRole(v BulkTunnelTerminationRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkVirtualCircuitTerminationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetInterface

`func (o *BulkVirtualCircuitTerminationRequest) GetInterface() BulkVirtualCircuitTerminationRequestInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BulkVirtualCircuitTerminationRequest) GetInterfaceOk() (*BulkVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BulkVirtualCircuitTerminationRequest) SetInterface(v BulkVirtualCircuitTerminationRequestInterface)`

SetInterface sets Interface field to given value.


### GetDescription

`func (o *BulkVirtualCircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkVirtualCircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkVirtualCircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkVirtualCircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *BulkVirtualCircuitTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkVirtualCircuitTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkVirtualCircuitTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkVirtualCircuitTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkVirtualCircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkVirtualCircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkVirtualCircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkVirtualCircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


