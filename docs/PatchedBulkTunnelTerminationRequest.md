# PatchedBulkTunnelTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Tunnel** | Pointer to [**BulkTunnelTerminationRequestTunnel**](BulkTunnelTerminationRequestTunnel.md) |  | [optional] 
**Role** | Pointer to [**BulkTunnelTerminationRequestRole**](BulkTunnelTerminationRequestRole.md) |  | [optional] 
**TerminationType** | Pointer to **string** |  | [optional] 
**TerminationId** | Pointer to **NullableInt64** |  | [optional] 
**OutsideIp** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkTunnelTerminationRequest

`func NewPatchedBulkTunnelTerminationRequest(id int32, ) *PatchedBulkTunnelTerminationRequest`

NewPatchedBulkTunnelTerminationRequest instantiates a new PatchedBulkTunnelTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkTunnelTerminationRequestWithDefaults

`func NewPatchedBulkTunnelTerminationRequestWithDefaults() *PatchedBulkTunnelTerminationRequest`

NewPatchedBulkTunnelTerminationRequestWithDefaults instantiates a new PatchedBulkTunnelTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkTunnelTerminationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkTunnelTerminationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkTunnelTerminationRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetTunnel

`func (o *PatchedBulkTunnelTerminationRequest) GetTunnel() BulkTunnelTerminationRequestTunnel`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *PatchedBulkTunnelTerminationRequest) GetTunnelOk() (*BulkTunnelTerminationRequestTunnel, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *PatchedBulkTunnelTerminationRequest) SetTunnel(v BulkTunnelTerminationRequestTunnel)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *PatchedBulkTunnelTerminationRequest) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkTunnelTerminationRequest) GetRole() BulkTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkTunnelTerminationRequest) GetRoleOk() (*BulkTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkTunnelTerminationRequest) SetRole(v BulkTunnelTerminationRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkTunnelTerminationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTerminationType

`func (o *PatchedBulkTunnelTerminationRequest) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *PatchedBulkTunnelTerminationRequest) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *PatchedBulkTunnelTerminationRequest) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.

### HasTerminationType

`func (o *PatchedBulkTunnelTerminationRequest) HasTerminationType() bool`

HasTerminationType returns a boolean if a field has been set.

### GetTerminationId

`func (o *PatchedBulkTunnelTerminationRequest) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *PatchedBulkTunnelTerminationRequest) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *PatchedBulkTunnelTerminationRequest) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.

### HasTerminationId

`func (o *PatchedBulkTunnelTerminationRequest) HasTerminationId() bool`

HasTerminationId returns a boolean if a field has been set.

### SetTerminationIdNil

`func (o *PatchedBulkTunnelTerminationRequest) SetTerminationIdNil(b bool)`

 SetTerminationIdNil sets the value for TerminationId to be an explicit nil

### UnsetTerminationId
`func (o *PatchedBulkTunnelTerminationRequest) UnsetTerminationId()`

UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
### GetOutsideIp

`func (o *PatchedBulkTunnelTerminationRequest) GetOutsideIp() BulkDeviceRequestPrimaryIp4`

GetOutsideIp returns the OutsideIp field if non-nil, zero value otherwise.

### GetOutsideIpOk

`func (o *PatchedBulkTunnelTerminationRequest) GetOutsideIpOk() (*BulkDeviceRequestPrimaryIp4, bool)`

GetOutsideIpOk returns a tuple with the OutsideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideIp

`func (o *PatchedBulkTunnelTerminationRequest) SetOutsideIp(v BulkDeviceRequestPrimaryIp4)`

SetOutsideIp sets OutsideIp field to given value.

### HasOutsideIp

`func (o *PatchedBulkTunnelTerminationRequest) HasOutsideIp() bool`

HasOutsideIp returns a boolean if a field has been set.

### SetOutsideIpNil

`func (o *PatchedBulkTunnelTerminationRequest) SetOutsideIpNil(b bool)`

 SetOutsideIpNil sets the value for OutsideIp to be an explicit nil

### UnsetOutsideIp
`func (o *PatchedBulkTunnelTerminationRequest) UnsetOutsideIp()`

UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
### GetTags

`func (o *PatchedBulkTunnelTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkTunnelTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkTunnelTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkTunnelTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkTunnelTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkTunnelTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkTunnelTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkTunnelTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


