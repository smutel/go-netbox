# PatchedBulkVirtualCircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**VirtualCircuit** | Pointer to [**BulkVirtualCircuitTerminationRequestVirtualCircuit**](BulkVirtualCircuitTerminationRequestVirtualCircuit.md) |  | [optional] 
**Role** | Pointer to [**BulkTunnelTerminationRequestRole**](BulkTunnelTerminationRequestRole.md) |  | [optional] 
**Interface** | Pointer to [**BulkVirtualCircuitTerminationRequestInterface**](BulkVirtualCircuitTerminationRequestInterface.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkVirtualCircuitTerminationRequest

`func NewPatchedBulkVirtualCircuitTerminationRequest(id int32, ) *PatchedBulkVirtualCircuitTerminationRequest`

NewPatchedBulkVirtualCircuitTerminationRequest instantiates a new PatchedBulkVirtualCircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkVirtualCircuitTerminationRequestWithDefaults

`func NewPatchedBulkVirtualCircuitTerminationRequestWithDefaults() *PatchedBulkVirtualCircuitTerminationRequest`

NewPatchedBulkVirtualCircuitTerminationRequestWithDefaults instantiates a new PatchedBulkVirtualCircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkVirtualCircuitTerminationRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetVirtualCircuit

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetVirtualCircuit() BulkVirtualCircuitTerminationRequestVirtualCircuit`

GetVirtualCircuit returns the VirtualCircuit field if non-nil, zero value otherwise.

### GetVirtualCircuitOk

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetVirtualCircuitOk() (*BulkVirtualCircuitTerminationRequestVirtualCircuit, bool)`

GetVirtualCircuitOk returns a tuple with the VirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuit

`func (o *PatchedBulkVirtualCircuitTerminationRequest) SetVirtualCircuit(v BulkVirtualCircuitTerminationRequestVirtualCircuit)`

SetVirtualCircuit sets VirtualCircuit field to given value.

### HasVirtualCircuit

`func (o *PatchedBulkVirtualCircuitTerminationRequest) HasVirtualCircuit() bool`

HasVirtualCircuit returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetRole() BulkTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetRoleOk() (*BulkTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkVirtualCircuitTerminationRequest) SetRole(v BulkTunnelTerminationRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkVirtualCircuitTerminationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetInterface

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetInterface() BulkVirtualCircuitTerminationRequestInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetInterfaceOk() (*BulkVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedBulkVirtualCircuitTerminationRequest) SetInterface(v BulkVirtualCircuitTerminationRequestInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedBulkVirtualCircuitTerminationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkVirtualCircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkVirtualCircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkVirtualCircuitTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkVirtualCircuitTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkVirtualCircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkVirtualCircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkVirtualCircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


