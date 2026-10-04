# WritableCoolingIntakeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | 
**Module** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullablePatchedWritableCoolingIntakeRequestType**](PatchedWritableCoolingIntakeRequestType.md) |  | [optional] 
**Diameter** | Pointer to **NullableFloat64** |  | [optional] 
**DiameterUnit** | Pointer to [**NullableBulkCoolingIntakeRequestDiameterUnit**](BulkCoolingIntakeRequestDiameterUnit.md) |  | [optional] 
**MaxFlow** | Pointer to **NullableFloat64** |  | [optional] 
**MaxFlowUnit** | Pointer to [**NullableBulkCoolingFeedRequestMaxFlowUnit**](BulkCoolingFeedRequestMaxFlowUnit.md) |  | [optional] 
**CoolingOutflow** | Pointer to **NullableInt32** | The upstream cooling outflow supplying this intake | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableCoolingIntakeRequest

`func NewWritableCoolingIntakeRequest(device BriefCoolingIntakeRequestDevice, name string, ) *WritableCoolingIntakeRequest`

NewWritableCoolingIntakeRequest instantiates a new WritableCoolingIntakeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCoolingIntakeRequestWithDefaults

`func NewWritableCoolingIntakeRequestWithDefaults() *WritableCoolingIntakeRequest`

NewWritableCoolingIntakeRequestWithDefaults instantiates a new WritableCoolingIntakeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *WritableCoolingIntakeRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableCoolingIntakeRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableCoolingIntakeRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *WritableCoolingIntakeRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WritableCoolingIntakeRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WritableCoolingIntakeRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *WritableCoolingIntakeRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *WritableCoolingIntakeRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *WritableCoolingIntakeRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *WritableCoolingIntakeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableCoolingIntakeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableCoolingIntakeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableCoolingIntakeRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableCoolingIntakeRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableCoolingIntakeRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableCoolingIntakeRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritableCoolingIntakeRequest) GetType() PatchedWritableCoolingIntakeRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCoolingIntakeRequest) GetTypeOk() (*PatchedWritableCoolingIntakeRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCoolingIntakeRequest) SetType(v PatchedWritableCoolingIntakeRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *WritableCoolingIntakeRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *WritableCoolingIntakeRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WritableCoolingIntakeRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDiameter

`func (o *WritableCoolingIntakeRequest) GetDiameter() float64`

GetDiameter returns the Diameter field if non-nil, zero value otherwise.

### GetDiameterOk

`func (o *WritableCoolingIntakeRequest) GetDiameterOk() (*float64, bool)`

GetDiameterOk returns a tuple with the Diameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameter

`func (o *WritableCoolingIntakeRequest) SetDiameter(v float64)`

SetDiameter sets Diameter field to given value.

### HasDiameter

`func (o *WritableCoolingIntakeRequest) HasDiameter() bool`

HasDiameter returns a boolean if a field has been set.

### SetDiameterNil

`func (o *WritableCoolingIntakeRequest) SetDiameterNil(b bool)`

 SetDiameterNil sets the value for Diameter to be an explicit nil

### UnsetDiameter
`func (o *WritableCoolingIntakeRequest) UnsetDiameter()`

UnsetDiameter ensures that no value is present for Diameter, not even an explicit nil
### GetDiameterUnit

`func (o *WritableCoolingIntakeRequest) GetDiameterUnit() BulkCoolingIntakeRequestDiameterUnit`

GetDiameterUnit returns the DiameterUnit field if non-nil, zero value otherwise.

### GetDiameterUnitOk

`func (o *WritableCoolingIntakeRequest) GetDiameterUnitOk() (*BulkCoolingIntakeRequestDiameterUnit, bool)`

GetDiameterUnitOk returns a tuple with the DiameterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterUnit

`func (o *WritableCoolingIntakeRequest) SetDiameterUnit(v BulkCoolingIntakeRequestDiameterUnit)`

SetDiameterUnit sets DiameterUnit field to given value.

### HasDiameterUnit

`func (o *WritableCoolingIntakeRequest) HasDiameterUnit() bool`

HasDiameterUnit returns a boolean if a field has been set.

### SetDiameterUnitNil

`func (o *WritableCoolingIntakeRequest) SetDiameterUnitNil(b bool)`

 SetDiameterUnitNil sets the value for DiameterUnit to be an explicit nil

### UnsetDiameterUnit
`func (o *WritableCoolingIntakeRequest) UnsetDiameterUnit()`

UnsetDiameterUnit ensures that no value is present for DiameterUnit, not even an explicit nil
### GetMaxFlow

`func (o *WritableCoolingIntakeRequest) GetMaxFlow() float64`

GetMaxFlow returns the MaxFlow field if non-nil, zero value otherwise.

### GetMaxFlowOk

`func (o *WritableCoolingIntakeRequest) GetMaxFlowOk() (*float64, bool)`

GetMaxFlowOk returns a tuple with the MaxFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlow

`func (o *WritableCoolingIntakeRequest) SetMaxFlow(v float64)`

SetMaxFlow sets MaxFlow field to given value.

### HasMaxFlow

`func (o *WritableCoolingIntakeRequest) HasMaxFlow() bool`

HasMaxFlow returns a boolean if a field has been set.

### SetMaxFlowNil

`func (o *WritableCoolingIntakeRequest) SetMaxFlowNil(b bool)`

 SetMaxFlowNil sets the value for MaxFlow to be an explicit nil

### UnsetMaxFlow
`func (o *WritableCoolingIntakeRequest) UnsetMaxFlow()`

UnsetMaxFlow ensures that no value is present for MaxFlow, not even an explicit nil
### GetMaxFlowUnit

`func (o *WritableCoolingIntakeRequest) GetMaxFlowUnit() BulkCoolingFeedRequestMaxFlowUnit`

GetMaxFlowUnit returns the MaxFlowUnit field if non-nil, zero value otherwise.

### GetMaxFlowUnitOk

`func (o *WritableCoolingIntakeRequest) GetMaxFlowUnitOk() (*BulkCoolingFeedRequestMaxFlowUnit, bool)`

GetMaxFlowUnitOk returns a tuple with the MaxFlowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlowUnit

`func (o *WritableCoolingIntakeRequest) SetMaxFlowUnit(v BulkCoolingFeedRequestMaxFlowUnit)`

SetMaxFlowUnit sets MaxFlowUnit field to given value.

### HasMaxFlowUnit

`func (o *WritableCoolingIntakeRequest) HasMaxFlowUnit() bool`

HasMaxFlowUnit returns a boolean if a field has been set.

### SetMaxFlowUnitNil

`func (o *WritableCoolingIntakeRequest) SetMaxFlowUnitNil(b bool)`

 SetMaxFlowUnitNil sets the value for MaxFlowUnit to be an explicit nil

### UnsetMaxFlowUnit
`func (o *WritableCoolingIntakeRequest) UnsetMaxFlowUnit()`

UnsetMaxFlowUnit ensures that no value is present for MaxFlowUnit, not even an explicit nil
### GetCoolingOutflow

`func (o *WritableCoolingIntakeRequest) GetCoolingOutflow() int32`

GetCoolingOutflow returns the CoolingOutflow field if non-nil, zero value otherwise.

### GetCoolingOutflowOk

`func (o *WritableCoolingIntakeRequest) GetCoolingOutflowOk() (*int32, bool)`

GetCoolingOutflowOk returns a tuple with the CoolingOutflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingOutflow

`func (o *WritableCoolingIntakeRequest) SetCoolingOutflow(v int32)`

SetCoolingOutflow sets CoolingOutflow field to given value.

### HasCoolingOutflow

`func (o *WritableCoolingIntakeRequest) HasCoolingOutflow() bool`

HasCoolingOutflow returns a boolean if a field has been set.

### SetCoolingOutflowNil

`func (o *WritableCoolingIntakeRequest) SetCoolingOutflowNil(b bool)`

 SetCoolingOutflowNil sets the value for CoolingOutflow to be an explicit nil

### UnsetCoolingOutflow
`func (o *WritableCoolingIntakeRequest) UnsetCoolingOutflow()`

UnsetCoolingOutflow ensures that no value is present for CoolingOutflow, not even an explicit nil
### GetDescription

`func (o *WritableCoolingIntakeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableCoolingIntakeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableCoolingIntakeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableCoolingIntakeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *WritableCoolingIntakeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WritableCoolingIntakeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WritableCoolingIntakeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WritableCoolingIntakeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WritableCoolingIntakeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WritableCoolingIntakeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *WritableCoolingIntakeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCoolingIntakeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCoolingIntakeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCoolingIntakeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCoolingIntakeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCoolingIntakeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCoolingIntakeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCoolingIntakeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


