# BulkTunnelTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Tunnel** | [**BulkTunnelTerminationRequestTunnel**](BulkTunnelTerminationRequestTunnel.md) |  | 
**Role** | [**BulkTunnelTerminationRequestRole**](BulkTunnelTerminationRequestRole.md) |  | 
**TerminationType** | **string** |  | 
**TerminationId** | Pointer to **NullableInt64** |  | [optional] 
**OutsideIp** | Pointer to [**NullableBulkDeviceRequestPrimaryIp4**](BulkDeviceRequestPrimaryIp4.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkTunnelTerminationRequest

`func NewBulkTunnelTerminationRequest(id int32, tunnel BulkTunnelTerminationRequestTunnel, role BulkTunnelTerminationRequestRole, terminationType string, ) *BulkTunnelTerminationRequest`

NewBulkTunnelTerminationRequest instantiates a new BulkTunnelTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkTunnelTerminationRequestWithDefaults

`func NewBulkTunnelTerminationRequestWithDefaults() *BulkTunnelTerminationRequest`

NewBulkTunnelTerminationRequestWithDefaults instantiates a new BulkTunnelTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkTunnelTerminationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkTunnelTerminationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkTunnelTerminationRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetTunnel

`func (o *BulkTunnelTerminationRequest) GetTunnel() BulkTunnelTerminationRequestTunnel`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *BulkTunnelTerminationRequest) GetTunnelOk() (*BulkTunnelTerminationRequestTunnel, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *BulkTunnelTerminationRequest) SetTunnel(v BulkTunnelTerminationRequestTunnel)`

SetTunnel sets Tunnel field to given value.


### GetRole

`func (o *BulkTunnelTerminationRequest) GetRole() BulkTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkTunnelTerminationRequest) GetRoleOk() (*BulkTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkTunnelTerminationRequest) SetRole(v BulkTunnelTerminationRequestRole)`

SetRole sets Role field to given value.


### GetTerminationType

`func (o *BulkTunnelTerminationRequest) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *BulkTunnelTerminationRequest) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *BulkTunnelTerminationRequest) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.


### GetTerminationId

`func (o *BulkTunnelTerminationRequest) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *BulkTunnelTerminationRequest) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *BulkTunnelTerminationRequest) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.

### HasTerminationId

`func (o *BulkTunnelTerminationRequest) HasTerminationId() bool`

HasTerminationId returns a boolean if a field has been set.

### SetTerminationIdNil

`func (o *BulkTunnelTerminationRequest) SetTerminationIdNil(b bool)`

 SetTerminationIdNil sets the value for TerminationId to be an explicit nil

### UnsetTerminationId
`func (o *BulkTunnelTerminationRequest) UnsetTerminationId()`

UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
### GetOutsideIp

`func (o *BulkTunnelTerminationRequest) GetOutsideIp() BulkDeviceRequestPrimaryIp4`

GetOutsideIp returns the OutsideIp field if non-nil, zero value otherwise.

### GetOutsideIpOk

`func (o *BulkTunnelTerminationRequest) GetOutsideIpOk() (*BulkDeviceRequestPrimaryIp4, bool)`

GetOutsideIpOk returns a tuple with the OutsideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideIp

`func (o *BulkTunnelTerminationRequest) SetOutsideIp(v BulkDeviceRequestPrimaryIp4)`

SetOutsideIp sets OutsideIp field to given value.

### HasOutsideIp

`func (o *BulkTunnelTerminationRequest) HasOutsideIp() bool`

HasOutsideIp returns a boolean if a field has been set.

### SetOutsideIpNil

`func (o *BulkTunnelTerminationRequest) SetOutsideIpNil(b bool)`

 SetOutsideIpNil sets the value for OutsideIp to be an explicit nil

### UnsetOutsideIp
`func (o *BulkTunnelTerminationRequest) UnsetOutsideIp()`

UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
### GetTags

`func (o *BulkTunnelTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkTunnelTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkTunnelTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkTunnelTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkTunnelTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkTunnelTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkTunnelTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkTunnelTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


