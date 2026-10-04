# BulkCircuitTerminationRequestCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** | Unique circuit ID | 
**Provider** | [**BriefCircuitRequestProvider**](BriefCircuitRequestProvider.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkCircuitTerminationRequestCircuit

`func NewBulkCircuitTerminationRequestCircuit(cid string, provider BriefCircuitRequestProvider, ) *BulkCircuitTerminationRequestCircuit`

NewBulkCircuitTerminationRequestCircuit instantiates a new BulkCircuitTerminationRequestCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCircuitTerminationRequestCircuitWithDefaults

`func NewBulkCircuitTerminationRequestCircuitWithDefaults() *BulkCircuitTerminationRequestCircuit`

NewBulkCircuitTerminationRequestCircuitWithDefaults instantiates a new BulkCircuitTerminationRequestCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *BulkCircuitTerminationRequestCircuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *BulkCircuitTerminationRequestCircuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *BulkCircuitTerminationRequestCircuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *BulkCircuitTerminationRequestCircuit) GetProvider() BriefCircuitRequestProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BulkCircuitTerminationRequestCircuit) GetProviderOk() (*BriefCircuitRequestProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BulkCircuitTerminationRequestCircuit) SetProvider(v BriefCircuitRequestProvider)`

SetProvider sets Provider field to given value.


### GetDescription

`func (o *BulkCircuitTerminationRequestCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkCircuitTerminationRequestCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkCircuitTerminationRequestCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkCircuitTerminationRequestCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


