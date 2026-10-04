# ServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentObjectType** | **string** |  | 
**ParentObjectId** | **int64** |  | 
**Name** | **string** |  | 
**PortMappings** | Pointer to **[]string** |  | [optional] 
**Protocol** | Pointer to [**NullableBulkServiceRequestProtocol**](BulkServiceRequestProtocol.md) |  | [optional] 
**Ports** | Pointer to **[]int32** | Deprecated; use port_mappings. Reported only for single-protocol services. | [optional] 
**Ipaddresses** | Pointer to **[]int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewServiceRequest

`func NewServiceRequest(parentObjectType string, parentObjectId int64, name string, ) *ServiceRequest`

NewServiceRequest instantiates a new ServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRequestWithDefaults

`func NewServiceRequestWithDefaults() *ServiceRequest`

NewServiceRequestWithDefaults instantiates a new ServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentObjectType

`func (o *ServiceRequest) GetParentObjectType() string`

GetParentObjectType returns the ParentObjectType field if non-nil, zero value otherwise.

### GetParentObjectTypeOk

`func (o *ServiceRequest) GetParentObjectTypeOk() (*string, bool)`

GetParentObjectTypeOk returns a tuple with the ParentObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectType

`func (o *ServiceRequest) SetParentObjectType(v string)`

SetParentObjectType sets ParentObjectType field to given value.


### GetParentObjectId

`func (o *ServiceRequest) GetParentObjectId() int64`

GetParentObjectId returns the ParentObjectId field if non-nil, zero value otherwise.

### GetParentObjectIdOk

`func (o *ServiceRequest) GetParentObjectIdOk() (*int64, bool)`

GetParentObjectIdOk returns a tuple with the ParentObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectId

`func (o *ServiceRequest) SetParentObjectId(v int64)`

SetParentObjectId sets ParentObjectId field to given value.


### GetName

`func (o *ServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPortMappings

`func (o *ServiceRequest) GetPortMappings() []string`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *ServiceRequest) GetPortMappingsOk() (*[]string, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *ServiceRequest) SetPortMappings(v []string)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *ServiceRequest) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetProtocol

`func (o *ServiceRequest) GetProtocol() BulkServiceRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServiceRequest) GetProtocolOk() (*BulkServiceRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServiceRequest) SetProtocol(v BulkServiceRequestProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ServiceRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *ServiceRequest) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *ServiceRequest) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetPorts

`func (o *ServiceRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ServiceRequest) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ServiceRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ServiceRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *ServiceRequest) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *ServiceRequest) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetIpaddresses

`func (o *ServiceRequest) GetIpaddresses() []int32`

GetIpaddresses returns the Ipaddresses field if non-nil, zero value otherwise.

### GetIpaddressesOk

`func (o *ServiceRequest) GetIpaddressesOk() (*[]int32, bool)`

GetIpaddressesOk returns a tuple with the Ipaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddresses

`func (o *ServiceRequest) SetIpaddresses(v []int32)`

SetIpaddresses sets Ipaddresses field to given value.

### HasIpaddresses

`func (o *ServiceRequest) HasIpaddresses() bool`

HasIpaddresses returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *ServiceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ServiceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ServiceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ServiceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ServiceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ServiceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *ServiceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ServiceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ServiceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ServiceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *ServiceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ServiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ServiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ServiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ServiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


