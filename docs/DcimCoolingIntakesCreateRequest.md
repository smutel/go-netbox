# DcimCoolingIntakesCreateRequest

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
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDcimCoolingIntakesCreateRequest

`func NewDcimCoolingIntakesCreateRequest(device BriefCoolingIntakeRequestDevice, name string, ) *DcimCoolingIntakesCreateRequest`

NewDcimCoolingIntakesCreateRequest instantiates a new DcimCoolingIntakesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimCoolingIntakesCreateRequestWithDefaults

`func NewDcimCoolingIntakesCreateRequestWithDefaults() *DcimCoolingIntakesCreateRequest`

NewDcimCoolingIntakesCreateRequestWithDefaults instantiates a new DcimCoolingIntakesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *DcimCoolingIntakesCreateRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DcimCoolingIntakesCreateRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DcimCoolingIntakesCreateRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *DcimCoolingIntakesCreateRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *DcimCoolingIntakesCreateRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *DcimCoolingIntakesCreateRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *DcimCoolingIntakesCreateRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *DcimCoolingIntakesCreateRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *DcimCoolingIntakesCreateRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *DcimCoolingIntakesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimCoolingIntakesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimCoolingIntakesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *DcimCoolingIntakesCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DcimCoolingIntakesCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DcimCoolingIntakesCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DcimCoolingIntakesCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *DcimCoolingIntakesCreateRequest) GetType() PatchedWritableCoolingIntakeRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcimCoolingIntakesCreateRequest) GetTypeOk() (*PatchedWritableCoolingIntakeRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcimCoolingIntakesCreateRequest) SetType(v PatchedWritableCoolingIntakeRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *DcimCoolingIntakesCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DcimCoolingIntakesCreateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DcimCoolingIntakesCreateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDiameter

`func (o *DcimCoolingIntakesCreateRequest) GetDiameter() float64`

GetDiameter returns the Diameter field if non-nil, zero value otherwise.

### GetDiameterOk

`func (o *DcimCoolingIntakesCreateRequest) GetDiameterOk() (*float64, bool)`

GetDiameterOk returns a tuple with the Diameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameter

`func (o *DcimCoolingIntakesCreateRequest) SetDiameter(v float64)`

SetDiameter sets Diameter field to given value.

### HasDiameter

`func (o *DcimCoolingIntakesCreateRequest) HasDiameter() bool`

HasDiameter returns a boolean if a field has been set.

### SetDiameterNil

`func (o *DcimCoolingIntakesCreateRequest) SetDiameterNil(b bool)`

 SetDiameterNil sets the value for Diameter to be an explicit nil

### UnsetDiameter
`func (o *DcimCoolingIntakesCreateRequest) UnsetDiameter()`

UnsetDiameter ensures that no value is present for Diameter, not even an explicit nil
### GetDiameterUnit

`func (o *DcimCoolingIntakesCreateRequest) GetDiameterUnit() BulkCoolingIntakeRequestDiameterUnit`

GetDiameterUnit returns the DiameterUnit field if non-nil, zero value otherwise.

### GetDiameterUnitOk

`func (o *DcimCoolingIntakesCreateRequest) GetDiameterUnitOk() (*BulkCoolingIntakeRequestDiameterUnit, bool)`

GetDiameterUnitOk returns a tuple with the DiameterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterUnit

`func (o *DcimCoolingIntakesCreateRequest) SetDiameterUnit(v BulkCoolingIntakeRequestDiameterUnit)`

SetDiameterUnit sets DiameterUnit field to given value.

### HasDiameterUnit

`func (o *DcimCoolingIntakesCreateRequest) HasDiameterUnit() bool`

HasDiameterUnit returns a boolean if a field has been set.

### SetDiameterUnitNil

`func (o *DcimCoolingIntakesCreateRequest) SetDiameterUnitNil(b bool)`

 SetDiameterUnitNil sets the value for DiameterUnit to be an explicit nil

### UnsetDiameterUnit
`func (o *DcimCoolingIntakesCreateRequest) UnsetDiameterUnit()`

UnsetDiameterUnit ensures that no value is present for DiameterUnit, not even an explicit nil
### GetMaxFlow

`func (o *DcimCoolingIntakesCreateRequest) GetMaxFlow() float64`

GetMaxFlow returns the MaxFlow field if non-nil, zero value otherwise.

### GetMaxFlowOk

`func (o *DcimCoolingIntakesCreateRequest) GetMaxFlowOk() (*float64, bool)`

GetMaxFlowOk returns a tuple with the MaxFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlow

`func (o *DcimCoolingIntakesCreateRequest) SetMaxFlow(v float64)`

SetMaxFlow sets MaxFlow field to given value.

### HasMaxFlow

`func (o *DcimCoolingIntakesCreateRequest) HasMaxFlow() bool`

HasMaxFlow returns a boolean if a field has been set.

### SetMaxFlowNil

`func (o *DcimCoolingIntakesCreateRequest) SetMaxFlowNil(b bool)`

 SetMaxFlowNil sets the value for MaxFlow to be an explicit nil

### UnsetMaxFlow
`func (o *DcimCoolingIntakesCreateRequest) UnsetMaxFlow()`

UnsetMaxFlow ensures that no value is present for MaxFlow, not even an explicit nil
### GetMaxFlowUnit

`func (o *DcimCoolingIntakesCreateRequest) GetMaxFlowUnit() BulkCoolingFeedRequestMaxFlowUnit`

GetMaxFlowUnit returns the MaxFlowUnit field if non-nil, zero value otherwise.

### GetMaxFlowUnitOk

`func (o *DcimCoolingIntakesCreateRequest) GetMaxFlowUnitOk() (*BulkCoolingFeedRequestMaxFlowUnit, bool)`

GetMaxFlowUnitOk returns a tuple with the MaxFlowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlowUnit

`func (o *DcimCoolingIntakesCreateRequest) SetMaxFlowUnit(v BulkCoolingFeedRequestMaxFlowUnit)`

SetMaxFlowUnit sets MaxFlowUnit field to given value.

### HasMaxFlowUnit

`func (o *DcimCoolingIntakesCreateRequest) HasMaxFlowUnit() bool`

HasMaxFlowUnit returns a boolean if a field has been set.

### SetMaxFlowUnitNil

`func (o *DcimCoolingIntakesCreateRequest) SetMaxFlowUnitNil(b bool)`

 SetMaxFlowUnitNil sets the value for MaxFlowUnit to be an explicit nil

### UnsetMaxFlowUnit
`func (o *DcimCoolingIntakesCreateRequest) UnsetMaxFlowUnit()`

UnsetMaxFlowUnit ensures that no value is present for MaxFlowUnit, not even an explicit nil
### GetCoolingOutflow

`func (o *DcimCoolingIntakesCreateRequest) GetCoolingOutflow() int32`

GetCoolingOutflow returns the CoolingOutflow field if non-nil, zero value otherwise.

### GetCoolingOutflowOk

`func (o *DcimCoolingIntakesCreateRequest) GetCoolingOutflowOk() (*int32, bool)`

GetCoolingOutflowOk returns a tuple with the CoolingOutflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingOutflow

`func (o *DcimCoolingIntakesCreateRequest) SetCoolingOutflow(v int32)`

SetCoolingOutflow sets CoolingOutflow field to given value.

### HasCoolingOutflow

`func (o *DcimCoolingIntakesCreateRequest) HasCoolingOutflow() bool`

HasCoolingOutflow returns a boolean if a field has been set.

### SetCoolingOutflowNil

`func (o *DcimCoolingIntakesCreateRequest) SetCoolingOutflowNil(b bool)`

 SetCoolingOutflowNil sets the value for CoolingOutflow to be an explicit nil

### UnsetCoolingOutflow
`func (o *DcimCoolingIntakesCreateRequest) UnsetCoolingOutflow()`

UnsetCoolingOutflow ensures that no value is present for CoolingOutflow, not even an explicit nil
### GetDescription

`func (o *DcimCoolingIntakesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimCoolingIntakesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimCoolingIntakesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimCoolingIntakesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *DcimCoolingIntakesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DcimCoolingIntakesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DcimCoolingIntakesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DcimCoolingIntakesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DcimCoolingIntakesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DcimCoolingIntakesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *DcimCoolingIntakesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DcimCoolingIntakesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DcimCoolingIntakesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DcimCoolingIntakesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DcimCoolingIntakesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DcimCoolingIntakesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DcimCoolingIntakesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DcimCoolingIntakesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


