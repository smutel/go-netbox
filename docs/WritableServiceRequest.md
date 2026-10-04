# WritableServiceRequest

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

### NewWritableServiceRequest

`func NewWritableServiceRequest(parentObjectType string, parentObjectId int64, name string, ) *WritableServiceRequest`

NewWritableServiceRequest instantiates a new WritableServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableServiceRequestWithDefaults

`func NewWritableServiceRequestWithDefaults() *WritableServiceRequest`

NewWritableServiceRequestWithDefaults instantiates a new WritableServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentObjectType

`func (o *WritableServiceRequest) GetParentObjectType() string`

GetParentObjectType returns the ParentObjectType field if non-nil, zero value otherwise.

### GetParentObjectTypeOk

`func (o *WritableServiceRequest) GetParentObjectTypeOk() (*string, bool)`

GetParentObjectTypeOk returns a tuple with the ParentObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectType

`func (o *WritableServiceRequest) SetParentObjectType(v string)`

SetParentObjectType sets ParentObjectType field to given value.


### GetParentObjectId

`func (o *WritableServiceRequest) GetParentObjectId() int64`

GetParentObjectId returns the ParentObjectId field if non-nil, zero value otherwise.

### GetParentObjectIdOk

`func (o *WritableServiceRequest) GetParentObjectIdOk() (*int64, bool)`

GetParentObjectIdOk returns a tuple with the ParentObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectId

`func (o *WritableServiceRequest) SetParentObjectId(v int64)`

SetParentObjectId sets ParentObjectId field to given value.


### GetName

`func (o *WritableServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPortMappings

`func (o *WritableServiceRequest) GetPortMappings() []string`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *WritableServiceRequest) GetPortMappingsOk() (*[]string, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *WritableServiceRequest) SetPortMappings(v []string)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *WritableServiceRequest) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetProtocol

`func (o *WritableServiceRequest) GetProtocol() BulkServiceRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WritableServiceRequest) GetProtocolOk() (*BulkServiceRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WritableServiceRequest) SetProtocol(v BulkServiceRequestProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *WritableServiceRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *WritableServiceRequest) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *WritableServiceRequest) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetPorts

`func (o *WritableServiceRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *WritableServiceRequest) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *WritableServiceRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *WritableServiceRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *WritableServiceRequest) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *WritableServiceRequest) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetIpaddresses

`func (o *WritableServiceRequest) GetIpaddresses() []int32`

GetIpaddresses returns the Ipaddresses field if non-nil, zero value otherwise.

### GetIpaddressesOk

`func (o *WritableServiceRequest) GetIpaddressesOk() (*[]int32, bool)`

GetIpaddressesOk returns a tuple with the Ipaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddresses

`func (o *WritableServiceRequest) SetIpaddresses(v []int32)`

SetIpaddresses sets Ipaddresses field to given value.

### HasIpaddresses

`func (o *WritableServiceRequest) HasIpaddresses() bool`

HasIpaddresses returns a boolean if a field has been set.

### GetDescription

`func (o *WritableServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *WritableServiceRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WritableServiceRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WritableServiceRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WritableServiceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WritableServiceRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WritableServiceRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *WritableServiceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableServiceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableServiceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableServiceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableServiceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableServiceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableServiceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableServiceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableServiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableServiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableServiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableServiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


