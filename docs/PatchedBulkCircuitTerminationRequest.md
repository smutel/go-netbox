# PatchedBulkCircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Circuit** | Pointer to [**BulkCircuitTerminationRequestCircuit**](BulkCircuitTerminationRequestCircuit.md) |  | [optional] 
**TermSide** | Pointer to [**TerminationSide1**](TerminationSide1.md) |  | [optional] 
**TerminationType** | Pointer to **NullableString** |  | [optional] 
**TerminationId** | Pointer to **NullableInt32** |  | [optional] 
**PortSpeed** | Pointer to **NullableInt32** | Physical circuit speed | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** | ID of the local cross-connect | [optional] 
**PpInfo** | Pointer to **string** | Patch panel ID and port number(s) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkCircuitTerminationRequest

`func NewPatchedBulkCircuitTerminationRequest(id int32, ) *PatchedBulkCircuitTerminationRequest`

NewPatchedBulkCircuitTerminationRequest instantiates a new PatchedBulkCircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCircuitTerminationRequestWithDefaults

`func NewPatchedBulkCircuitTerminationRequestWithDefaults() *PatchedBulkCircuitTerminationRequest`

NewPatchedBulkCircuitTerminationRequestWithDefaults instantiates a new PatchedBulkCircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCircuitTerminationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCircuitTerminationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCircuitTerminationRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCircuit

`func (o *PatchedBulkCircuitTerminationRequest) GetCircuit() BulkCircuitTerminationRequestCircuit`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *PatchedBulkCircuitTerminationRequest) GetCircuitOk() (*BulkCircuitTerminationRequestCircuit, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *PatchedBulkCircuitTerminationRequest) SetCircuit(v BulkCircuitTerminationRequestCircuit)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *PatchedBulkCircuitTerminationRequest) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetTermSide

`func (o *PatchedBulkCircuitTerminationRequest) GetTermSide() TerminationSide1`

GetTermSide returns the TermSide field if non-nil, zero value otherwise.

### GetTermSideOk

`func (o *PatchedBulkCircuitTerminationRequest) GetTermSideOk() (*TerminationSide1, bool)`

GetTermSideOk returns a tuple with the TermSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSide

`func (o *PatchedBulkCircuitTerminationRequest) SetTermSide(v TerminationSide1)`

SetTermSide sets TermSide field to given value.

### HasTermSide

`func (o *PatchedBulkCircuitTerminationRequest) HasTermSide() bool`

HasTermSide returns a boolean if a field has been set.

### GetTerminationType

`func (o *PatchedBulkCircuitTerminationRequest) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *PatchedBulkCircuitTerminationRequest) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *PatchedBulkCircuitTerminationRequest) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.

### HasTerminationType

`func (o *PatchedBulkCircuitTerminationRequest) HasTerminationType() bool`

HasTerminationType returns a boolean if a field has been set.

### SetTerminationTypeNil

`func (o *PatchedBulkCircuitTerminationRequest) SetTerminationTypeNil(b bool)`

 SetTerminationTypeNil sets the value for TerminationType to be an explicit nil

### UnsetTerminationType
`func (o *PatchedBulkCircuitTerminationRequest) UnsetTerminationType()`

UnsetTerminationType ensures that no value is present for TerminationType, not even an explicit nil
### GetTerminationId

`func (o *PatchedBulkCircuitTerminationRequest) GetTerminationId() int32`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *PatchedBulkCircuitTerminationRequest) GetTerminationIdOk() (*int32, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *PatchedBulkCircuitTerminationRequest) SetTerminationId(v int32)`

SetTerminationId sets TerminationId field to given value.

### HasTerminationId

`func (o *PatchedBulkCircuitTerminationRequest) HasTerminationId() bool`

HasTerminationId returns a boolean if a field has been set.

### SetTerminationIdNil

`func (o *PatchedBulkCircuitTerminationRequest) SetTerminationIdNil(b bool)`

 SetTerminationIdNil sets the value for TerminationId to be an explicit nil

### UnsetTerminationId
`func (o *PatchedBulkCircuitTerminationRequest) UnsetTerminationId()`

UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
### GetPortSpeed

`func (o *PatchedBulkCircuitTerminationRequest) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *PatchedBulkCircuitTerminationRequest) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *PatchedBulkCircuitTerminationRequest) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *PatchedBulkCircuitTerminationRequest) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *PatchedBulkCircuitTerminationRequest) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *PatchedBulkCircuitTerminationRequest) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *PatchedBulkCircuitTerminationRequest) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *PatchedBulkCircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *PatchedBulkCircuitTerminationRequest) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *PatchedBulkCircuitTerminationRequest) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *PatchedBulkCircuitTerminationRequest) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *PatchedBulkCircuitTerminationRequest) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *PatchedBulkCircuitTerminationRequest) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *PatchedBulkCircuitTerminationRequest) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *PatchedBulkCircuitTerminationRequest) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *PatchedBulkCircuitTerminationRequest) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetPpInfo

`func (o *PatchedBulkCircuitTerminationRequest) GetPpInfo() string`

GetPpInfo returns the PpInfo field if non-nil, zero value otherwise.

### GetPpInfoOk

`func (o *PatchedBulkCircuitTerminationRequest) GetPpInfoOk() (*string, bool)`

GetPpInfoOk returns a tuple with the PpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpInfo

`func (o *PatchedBulkCircuitTerminationRequest) SetPpInfo(v string)`

SetPpInfo sets PpInfo field to given value.

### HasPpInfo

`func (o *PatchedBulkCircuitTerminationRequest) HasPpInfo() bool`

HasPpInfo returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkCircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkCircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkCircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkCircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PatchedBulkCircuitTerminationRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PatchedBulkCircuitTerminationRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PatchedBulkCircuitTerminationRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PatchedBulkCircuitTerminationRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkCircuitTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkCircuitTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkCircuitTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkCircuitTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkCircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkCircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkCircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkCircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


