# PatchedBulkCoolingIntakeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Device** | Pointer to [**BriefCoolingIntakeRequestDevice**](BriefCoolingIntakeRequestDevice.md) |  | [optional] 
**Module** | Pointer to [**NullableBulkConsolePortRequestModule**](BulkConsolePortRequestModule.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullableBulkCoolingIntakeRequestType**](BulkCoolingIntakeRequestType.md) |  | [optional] 
**Diameter** | Pointer to **NullableFloat64** |  | [optional] 
**DiameterUnit** | Pointer to [**NullableBulkCoolingIntakeRequestDiameterUnit**](BulkCoolingIntakeRequestDiameterUnit.md) |  | [optional] 
**MaxFlow** | Pointer to **NullableFloat64** |  | [optional] 
**MaxFlowUnit** | Pointer to [**NullableBulkCoolingFeedRequestMaxFlowUnit**](BulkCoolingFeedRequestMaxFlowUnit.md) |  | [optional] 
**CoolingOutflow** | Pointer to [**NullableNestedCoolingOutflowRequest**](NestedCoolingOutflowRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkCoolingIntakeRequest

`func NewPatchedBulkCoolingIntakeRequest(id int32, ) *PatchedBulkCoolingIntakeRequest`

NewPatchedBulkCoolingIntakeRequest instantiates a new PatchedBulkCoolingIntakeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCoolingIntakeRequestWithDefaults

`func NewPatchedBulkCoolingIntakeRequestWithDefaults() *PatchedBulkCoolingIntakeRequest`

NewPatchedBulkCoolingIntakeRequestWithDefaults instantiates a new PatchedBulkCoolingIntakeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCoolingIntakeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCoolingIntakeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCoolingIntakeRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *PatchedBulkCoolingIntakeRequest) GetDevice() BriefCoolingIntakeRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkCoolingIntakeRequest) GetDeviceOk() (*BriefCoolingIntakeRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkCoolingIntakeRequest) SetDevice(v BriefCoolingIntakeRequestDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkCoolingIntakeRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModule

`func (o *PatchedBulkCoolingIntakeRequest) GetModule() BulkConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PatchedBulkCoolingIntakeRequest) GetModuleOk() (*BulkConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PatchedBulkCoolingIntakeRequest) SetModule(v BulkConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *PatchedBulkCoolingIntakeRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *PatchedBulkCoolingIntakeRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *PatchedBulkCoolingIntakeRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *PatchedBulkCoolingIntakeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkCoolingIntakeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkCoolingIntakeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkCoolingIntakeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkCoolingIntakeRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkCoolingIntakeRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkCoolingIntakeRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkCoolingIntakeRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkCoolingIntakeRequest) GetType() BulkCoolingIntakeRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkCoolingIntakeRequest) GetTypeOk() (*BulkCoolingIntakeRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkCoolingIntakeRequest) SetType(v BulkCoolingIntakeRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkCoolingIntakeRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchedBulkCoolingIntakeRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchedBulkCoolingIntakeRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDiameter

`func (o *PatchedBulkCoolingIntakeRequest) GetDiameter() float64`

GetDiameter returns the Diameter field if non-nil, zero value otherwise.

### GetDiameterOk

`func (o *PatchedBulkCoolingIntakeRequest) GetDiameterOk() (*float64, bool)`

GetDiameterOk returns a tuple with the Diameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameter

`func (o *PatchedBulkCoolingIntakeRequest) SetDiameter(v float64)`

SetDiameter sets Diameter field to given value.

### HasDiameter

`func (o *PatchedBulkCoolingIntakeRequest) HasDiameter() bool`

HasDiameter returns a boolean if a field has been set.

### SetDiameterNil

`func (o *PatchedBulkCoolingIntakeRequest) SetDiameterNil(b bool)`

 SetDiameterNil sets the value for Diameter to be an explicit nil

### UnsetDiameter
`func (o *PatchedBulkCoolingIntakeRequest) UnsetDiameter()`

UnsetDiameter ensures that no value is present for Diameter, not even an explicit nil
### GetDiameterUnit

`func (o *PatchedBulkCoolingIntakeRequest) GetDiameterUnit() BulkCoolingIntakeRequestDiameterUnit`

GetDiameterUnit returns the DiameterUnit field if non-nil, zero value otherwise.

### GetDiameterUnitOk

`func (o *PatchedBulkCoolingIntakeRequest) GetDiameterUnitOk() (*BulkCoolingIntakeRequestDiameterUnit, bool)`

GetDiameterUnitOk returns a tuple with the DiameterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterUnit

`func (o *PatchedBulkCoolingIntakeRequest) SetDiameterUnit(v BulkCoolingIntakeRequestDiameterUnit)`

SetDiameterUnit sets DiameterUnit field to given value.

### HasDiameterUnit

`func (o *PatchedBulkCoolingIntakeRequest) HasDiameterUnit() bool`

HasDiameterUnit returns a boolean if a field has been set.

### SetDiameterUnitNil

`func (o *PatchedBulkCoolingIntakeRequest) SetDiameterUnitNil(b bool)`

 SetDiameterUnitNil sets the value for DiameterUnit to be an explicit nil

### UnsetDiameterUnit
`func (o *PatchedBulkCoolingIntakeRequest) UnsetDiameterUnit()`

UnsetDiameterUnit ensures that no value is present for DiameterUnit, not even an explicit nil
### GetMaxFlow

`func (o *PatchedBulkCoolingIntakeRequest) GetMaxFlow() float64`

GetMaxFlow returns the MaxFlow field if non-nil, zero value otherwise.

### GetMaxFlowOk

`func (o *PatchedBulkCoolingIntakeRequest) GetMaxFlowOk() (*float64, bool)`

GetMaxFlowOk returns a tuple with the MaxFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlow

`func (o *PatchedBulkCoolingIntakeRequest) SetMaxFlow(v float64)`

SetMaxFlow sets MaxFlow field to given value.

### HasMaxFlow

`func (o *PatchedBulkCoolingIntakeRequest) HasMaxFlow() bool`

HasMaxFlow returns a boolean if a field has been set.

### SetMaxFlowNil

`func (o *PatchedBulkCoolingIntakeRequest) SetMaxFlowNil(b bool)`

 SetMaxFlowNil sets the value for MaxFlow to be an explicit nil

### UnsetMaxFlow
`func (o *PatchedBulkCoolingIntakeRequest) UnsetMaxFlow()`

UnsetMaxFlow ensures that no value is present for MaxFlow, not even an explicit nil
### GetMaxFlowUnit

`func (o *PatchedBulkCoolingIntakeRequest) GetMaxFlowUnit() BulkCoolingFeedRequestMaxFlowUnit`

GetMaxFlowUnit returns the MaxFlowUnit field if non-nil, zero value otherwise.

### GetMaxFlowUnitOk

`func (o *PatchedBulkCoolingIntakeRequest) GetMaxFlowUnitOk() (*BulkCoolingFeedRequestMaxFlowUnit, bool)`

GetMaxFlowUnitOk returns a tuple with the MaxFlowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlowUnit

`func (o *PatchedBulkCoolingIntakeRequest) SetMaxFlowUnit(v BulkCoolingFeedRequestMaxFlowUnit)`

SetMaxFlowUnit sets MaxFlowUnit field to given value.

### HasMaxFlowUnit

`func (o *PatchedBulkCoolingIntakeRequest) HasMaxFlowUnit() bool`

HasMaxFlowUnit returns a boolean if a field has been set.

### SetMaxFlowUnitNil

`func (o *PatchedBulkCoolingIntakeRequest) SetMaxFlowUnitNil(b bool)`

 SetMaxFlowUnitNil sets the value for MaxFlowUnit to be an explicit nil

### UnsetMaxFlowUnit
`func (o *PatchedBulkCoolingIntakeRequest) UnsetMaxFlowUnit()`

UnsetMaxFlowUnit ensures that no value is present for MaxFlowUnit, not even an explicit nil
### GetCoolingOutflow

`func (o *PatchedBulkCoolingIntakeRequest) GetCoolingOutflow() NestedCoolingOutflowRequest`

GetCoolingOutflow returns the CoolingOutflow field if non-nil, zero value otherwise.

### GetCoolingOutflowOk

`func (o *PatchedBulkCoolingIntakeRequest) GetCoolingOutflowOk() (*NestedCoolingOutflowRequest, bool)`

GetCoolingOutflowOk returns a tuple with the CoolingOutflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingOutflow

`func (o *PatchedBulkCoolingIntakeRequest) SetCoolingOutflow(v NestedCoolingOutflowRequest)`

SetCoolingOutflow sets CoolingOutflow field to given value.

### HasCoolingOutflow

`func (o *PatchedBulkCoolingIntakeRequest) HasCoolingOutflow() bool`

HasCoolingOutflow returns a boolean if a field has been set.

### SetCoolingOutflowNil

`func (o *PatchedBulkCoolingIntakeRequest) SetCoolingOutflowNil(b bool)`

 SetCoolingOutflowNil sets the value for CoolingOutflow to be an explicit nil

### UnsetCoolingOutflow
`func (o *PatchedBulkCoolingIntakeRequest) UnsetCoolingOutflow()`

UnsetCoolingOutflow ensures that no value is present for CoolingOutflow, not even an explicit nil
### GetDescription

`func (o *PatchedBulkCoolingIntakeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkCoolingIntakeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkCoolingIntakeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkCoolingIntakeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkCoolingIntakeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkCoolingIntakeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkCoolingIntakeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkCoolingIntakeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkCoolingIntakeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkCoolingIntakeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkCoolingIntakeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkCoolingIntakeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkCoolingIntakeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkCoolingIntakeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkCoolingIntakeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkCoolingIntakeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkCoolingIntakeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkCoolingIntakeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


