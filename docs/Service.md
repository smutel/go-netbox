# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**ParentObjectType** | **string** |  | 
**ParentObjectId** | **int64** |  | 
**Parent** | **interface{}** |  | [readonly] 
**Name** | **string** |  | 
**PortMappings** | Pointer to **[]string** |  | [optional] 
**Protocol** | Pointer to [**NullableServiceProtocol**](ServiceProtocol.md) |  | [optional] 
**Ports** | Pointer to **[]int32** | Deprecated; use port_mappings. Reported only for single-protocol services. | [optional] 
**Ipaddresses** | Pointer to [**[]BriefIPAddress**](BriefIPAddress.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableBriefOwner**](BriefOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewService

`func NewService(id int32, url string, displayUrl string, display string, parentObjectType string, parentObjectId int64, parent interface{}, name string, created NullableTime, lastUpdated NullableTime, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Service) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Service) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Service) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *Service) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *Service) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *Service) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *Service) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Service) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Service) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetParentObjectType

`func (o *Service) GetParentObjectType() string`

GetParentObjectType returns the ParentObjectType field if non-nil, zero value otherwise.

### GetParentObjectTypeOk

`func (o *Service) GetParentObjectTypeOk() (*string, bool)`

GetParentObjectTypeOk returns a tuple with the ParentObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectType

`func (o *Service) SetParentObjectType(v string)`

SetParentObjectType sets ParentObjectType field to given value.


### GetParentObjectId

`func (o *Service) GetParentObjectId() int64`

GetParentObjectId returns the ParentObjectId field if non-nil, zero value otherwise.

### GetParentObjectIdOk

`func (o *Service) GetParentObjectIdOk() (*int64, bool)`

GetParentObjectIdOk returns a tuple with the ParentObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectId

`func (o *Service) SetParentObjectId(v int64)`

SetParentObjectId sets ParentObjectId field to given value.


### GetParent

`func (o *Service) GetParent() interface{}`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Service) GetParentOk() (*interface{}, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Service) SetParent(v interface{})`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *Service) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Service) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetPortMappings

`func (o *Service) GetPortMappings() []string`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *Service) GetPortMappingsOk() (*[]string, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *Service) SetPortMappings(v []string)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *Service) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetProtocol

`func (o *Service) GetProtocol() ServiceProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Service) GetProtocolOk() (*ServiceProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Service) SetProtocol(v ServiceProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Service) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *Service) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *Service) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetPorts

`func (o *Service) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Service) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Service) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Service) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *Service) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *Service) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetIpaddresses

`func (o *Service) GetIpaddresses() []BriefIPAddress`

GetIpaddresses returns the Ipaddresses field if non-nil, zero value otherwise.

### GetIpaddressesOk

`func (o *Service) GetIpaddressesOk() (*[]BriefIPAddress, bool)`

GetIpaddressesOk returns a tuple with the Ipaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddresses

`func (o *Service) SetIpaddresses(v []BriefIPAddress)`

SetIpaddresses sets Ipaddresses field to given value.

### HasIpaddresses

`func (o *Service) HasIpaddresses() bool`

HasIpaddresses returns a boolean if a field has been set.

### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Service) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *Service) GetOwner() BriefOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Service) GetOwnerOk() (*BriefOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Service) SetOwner(v BriefOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Service) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Service) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Service) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *Service) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Service) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Service) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Service) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Service) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Service) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Service) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Service) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Service) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Service) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Service) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Service) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Service) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Service) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Service) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Service) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Service) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Service) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Service) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Service) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Service) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Service) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


