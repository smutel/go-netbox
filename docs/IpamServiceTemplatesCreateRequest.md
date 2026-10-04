# IpamServiceTemplatesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PortMappings** | Pointer to **[]string** |  | [optional] 
**Protocol** | Pointer to [**NullableBulkServiceRequestProtocol**](BulkServiceRequestProtocol.md) |  | [optional] 
**Ports** | Pointer to **[]int32** | Deprecated; use port_mappings. Reported only for single-protocol services. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewIpamServiceTemplatesCreateRequest

`func NewIpamServiceTemplatesCreateRequest(name string, ) *IpamServiceTemplatesCreateRequest`

NewIpamServiceTemplatesCreateRequest instantiates a new IpamServiceTemplatesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamServiceTemplatesCreateRequestWithDefaults

`func NewIpamServiceTemplatesCreateRequestWithDefaults() *IpamServiceTemplatesCreateRequest`

NewIpamServiceTemplatesCreateRequestWithDefaults instantiates a new IpamServiceTemplatesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IpamServiceTemplatesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamServiceTemplatesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamServiceTemplatesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPortMappings

`func (o *IpamServiceTemplatesCreateRequest) GetPortMappings() []string`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *IpamServiceTemplatesCreateRequest) GetPortMappingsOk() (*[]string, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *IpamServiceTemplatesCreateRequest) SetPortMappings(v []string)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *IpamServiceTemplatesCreateRequest) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetProtocol

`func (o *IpamServiceTemplatesCreateRequest) GetProtocol() BulkServiceRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpamServiceTemplatesCreateRequest) GetProtocolOk() (*BulkServiceRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpamServiceTemplatesCreateRequest) SetProtocol(v BulkServiceRequestProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpamServiceTemplatesCreateRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *IpamServiceTemplatesCreateRequest) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *IpamServiceTemplatesCreateRequest) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetPorts

`func (o *IpamServiceTemplatesCreateRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IpamServiceTemplatesCreateRequest) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IpamServiceTemplatesCreateRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IpamServiceTemplatesCreateRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *IpamServiceTemplatesCreateRequest) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *IpamServiceTemplatesCreateRequest) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetDescription

`func (o *IpamServiceTemplatesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpamServiceTemplatesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpamServiceTemplatesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpamServiceTemplatesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *IpamServiceTemplatesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IpamServiceTemplatesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IpamServiceTemplatesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IpamServiceTemplatesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *IpamServiceTemplatesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *IpamServiceTemplatesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *IpamServiceTemplatesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IpamServiceTemplatesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IpamServiceTemplatesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IpamServiceTemplatesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *IpamServiceTemplatesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamServiceTemplatesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamServiceTemplatesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamServiceTemplatesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IpamServiceTemplatesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IpamServiceTemplatesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IpamServiceTemplatesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IpamServiceTemplatesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


