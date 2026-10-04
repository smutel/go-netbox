# VirtualMachineStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**BulkVirtualMachineRequestStatus**](BulkVirtualMachineRequestStatus.md) |  | [optional] 
**Label** | Pointer to [**VirtualMachineStatusLabel**](VirtualMachineStatusLabel.md) |  | [optional] 

## Methods

### NewVirtualMachineStatus

`func NewVirtualMachineStatus() *VirtualMachineStatus`

NewVirtualMachineStatus instantiates a new VirtualMachineStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineStatusWithDefaults

`func NewVirtualMachineStatusWithDefaults() *VirtualMachineStatus`

NewVirtualMachineStatusWithDefaults instantiates a new VirtualMachineStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *VirtualMachineStatus) GetValue() BulkVirtualMachineRequestStatus`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VirtualMachineStatus) GetValueOk() (*BulkVirtualMachineRequestStatus, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VirtualMachineStatus) SetValue(v BulkVirtualMachineRequestStatus)`

SetValue sets Value field to given value.

### HasValue

`func (o *VirtualMachineStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *VirtualMachineStatus) GetLabel() VirtualMachineStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VirtualMachineStatus) GetLabelOk() (*VirtualMachineStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VirtualMachineStatus) SetLabel(v VirtualMachineStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VirtualMachineStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


