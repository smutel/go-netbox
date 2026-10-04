# BulkVirtualCircuitTerminationRequestVirtualCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** | Unique circuit ID | 
**ProviderNetwork** | [**BriefVirtualCircuitRequestProviderNetwork**](BriefVirtualCircuitRequestProviderNetwork.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkVirtualCircuitTerminationRequestVirtualCircuit

`func NewBulkVirtualCircuitTerminationRequestVirtualCircuit(cid string, providerNetwork BriefVirtualCircuitRequestProviderNetwork, ) *BulkVirtualCircuitTerminationRequestVirtualCircuit`

NewBulkVirtualCircuitTerminationRequestVirtualCircuit instantiates a new BulkVirtualCircuitTerminationRequestVirtualCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVirtualCircuitTerminationRequestVirtualCircuitWithDefaults

`func NewBulkVirtualCircuitTerminationRequestVirtualCircuitWithDefaults() *BulkVirtualCircuitTerminationRequestVirtualCircuit`

NewBulkVirtualCircuitTerminationRequestVirtualCircuitWithDefaults instantiates a new BulkVirtualCircuitTerminationRequestVirtualCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProviderNetwork

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) GetProviderNetwork() BriefVirtualCircuitRequestProviderNetwork`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) GetProviderNetworkOk() (*BriefVirtualCircuitRequestProviderNetwork, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) SetProviderNetwork(v BriefVirtualCircuitRequestProviderNetwork)`

SetProviderNetwork sets ProviderNetwork field to given value.


### GetDescription

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkVirtualCircuitTerminationRequestVirtualCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


