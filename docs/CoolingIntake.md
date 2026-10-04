# CoolingIntake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**BriefDevice**](BriefDevice.md) |  | 
**Module** | Pointer to [**NullableBriefModule**](BriefModule.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullableCoolingIntakeType**](CoolingIntakeType.md) |  | [optional] 
**Diameter** | Pointer to **NullableFloat64** |  | [optional] 
**DiameterUnit** | Pointer to [**NullableCoolingIntakeDiameterUnit**](CoolingIntakeDiameterUnit.md) |  | [optional] 
**MaxFlow** | Pointer to **NullableFloat64** |  | [optional] 
**MaxFlowUnit** | Pointer to [**NullableCoolingFeedMaxFlowUnit**](CoolingFeedMaxFlowUnit.md) |  | [optional] 
**CoolingOutflow** | Pointer to [**NullableNestedCoolingOutflow**](NestedCoolingOutflow.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableBriefOwner**](BriefOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewCoolingIntake

`func NewCoolingIntake(id int32, url string, displayUrl string, display string, device BriefDevice, name string, created NullableTime, lastUpdated NullableTime, ) *CoolingIntake`

NewCoolingIntake instantiates a new CoolingIntake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoolingIntakeWithDefaults

`func NewCoolingIntakeWithDefaults() *CoolingIntake`

NewCoolingIntakeWithDefaults instantiates a new CoolingIntake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoolingIntake) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoolingIntake) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoolingIntake) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CoolingIntake) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CoolingIntake) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CoolingIntake) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *CoolingIntake) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *CoolingIntake) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *CoolingIntake) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *CoolingIntake) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CoolingIntake) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CoolingIntake) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *CoolingIntake) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CoolingIntake) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CoolingIntake) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *CoolingIntake) GetModule() BriefModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CoolingIntake) GetModuleOk() (*BriefModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CoolingIntake) SetModule(v BriefModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *CoolingIntake) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *CoolingIntake) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *CoolingIntake) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *CoolingIntake) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoolingIntake) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoolingIntake) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CoolingIntake) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CoolingIntake) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CoolingIntake) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CoolingIntake) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *CoolingIntake) GetType() CoolingIntakeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoolingIntake) GetTypeOk() (*CoolingIntakeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoolingIntake) SetType(v CoolingIntakeType)`

SetType sets Type field to given value.

### HasType

`func (o *CoolingIntake) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CoolingIntake) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CoolingIntake) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDiameter

`func (o *CoolingIntake) GetDiameter() float64`

GetDiameter returns the Diameter field if non-nil, zero value otherwise.

### GetDiameterOk

`func (o *CoolingIntake) GetDiameterOk() (*float64, bool)`

GetDiameterOk returns a tuple with the Diameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameter

`func (o *CoolingIntake) SetDiameter(v float64)`

SetDiameter sets Diameter field to given value.

### HasDiameter

`func (o *CoolingIntake) HasDiameter() bool`

HasDiameter returns a boolean if a field has been set.

### SetDiameterNil

`func (o *CoolingIntake) SetDiameterNil(b bool)`

 SetDiameterNil sets the value for Diameter to be an explicit nil

### UnsetDiameter
`func (o *CoolingIntake) UnsetDiameter()`

UnsetDiameter ensures that no value is present for Diameter, not even an explicit nil
### GetDiameterUnit

`func (o *CoolingIntake) GetDiameterUnit() CoolingIntakeDiameterUnit`

GetDiameterUnit returns the DiameterUnit field if non-nil, zero value otherwise.

### GetDiameterUnitOk

`func (o *CoolingIntake) GetDiameterUnitOk() (*CoolingIntakeDiameterUnit, bool)`

GetDiameterUnitOk returns a tuple with the DiameterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterUnit

`func (o *CoolingIntake) SetDiameterUnit(v CoolingIntakeDiameterUnit)`

SetDiameterUnit sets DiameterUnit field to given value.

### HasDiameterUnit

`func (o *CoolingIntake) HasDiameterUnit() bool`

HasDiameterUnit returns a boolean if a field has been set.

### SetDiameterUnitNil

`func (o *CoolingIntake) SetDiameterUnitNil(b bool)`

 SetDiameterUnitNil sets the value for DiameterUnit to be an explicit nil

### UnsetDiameterUnit
`func (o *CoolingIntake) UnsetDiameterUnit()`

UnsetDiameterUnit ensures that no value is present for DiameterUnit, not even an explicit nil
### GetMaxFlow

`func (o *CoolingIntake) GetMaxFlow() float64`

GetMaxFlow returns the MaxFlow field if non-nil, zero value otherwise.

### GetMaxFlowOk

`func (o *CoolingIntake) GetMaxFlowOk() (*float64, bool)`

GetMaxFlowOk returns a tuple with the MaxFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlow

`func (o *CoolingIntake) SetMaxFlow(v float64)`

SetMaxFlow sets MaxFlow field to given value.

### HasMaxFlow

`func (o *CoolingIntake) HasMaxFlow() bool`

HasMaxFlow returns a boolean if a field has been set.

### SetMaxFlowNil

`func (o *CoolingIntake) SetMaxFlowNil(b bool)`

 SetMaxFlowNil sets the value for MaxFlow to be an explicit nil

### UnsetMaxFlow
`func (o *CoolingIntake) UnsetMaxFlow()`

UnsetMaxFlow ensures that no value is present for MaxFlow, not even an explicit nil
### GetMaxFlowUnit

`func (o *CoolingIntake) GetMaxFlowUnit() CoolingFeedMaxFlowUnit`

GetMaxFlowUnit returns the MaxFlowUnit field if non-nil, zero value otherwise.

### GetMaxFlowUnitOk

`func (o *CoolingIntake) GetMaxFlowUnitOk() (*CoolingFeedMaxFlowUnit, bool)`

GetMaxFlowUnitOk returns a tuple with the MaxFlowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlowUnit

`func (o *CoolingIntake) SetMaxFlowUnit(v CoolingFeedMaxFlowUnit)`

SetMaxFlowUnit sets MaxFlowUnit field to given value.

### HasMaxFlowUnit

`func (o *CoolingIntake) HasMaxFlowUnit() bool`

HasMaxFlowUnit returns a boolean if a field has been set.

### SetMaxFlowUnitNil

`func (o *CoolingIntake) SetMaxFlowUnitNil(b bool)`

 SetMaxFlowUnitNil sets the value for MaxFlowUnit to be an explicit nil

### UnsetMaxFlowUnit
`func (o *CoolingIntake) UnsetMaxFlowUnit()`

UnsetMaxFlowUnit ensures that no value is present for MaxFlowUnit, not even an explicit nil
### GetCoolingOutflow

`func (o *CoolingIntake) GetCoolingOutflow() NestedCoolingOutflow`

GetCoolingOutflow returns the CoolingOutflow field if non-nil, zero value otherwise.

### GetCoolingOutflowOk

`func (o *CoolingIntake) GetCoolingOutflowOk() (*NestedCoolingOutflow, bool)`

GetCoolingOutflowOk returns a tuple with the CoolingOutflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingOutflow

`func (o *CoolingIntake) SetCoolingOutflow(v NestedCoolingOutflow)`

SetCoolingOutflow sets CoolingOutflow field to given value.

### HasCoolingOutflow

`func (o *CoolingIntake) HasCoolingOutflow() bool`

HasCoolingOutflow returns a boolean if a field has been set.

### SetCoolingOutflowNil

`func (o *CoolingIntake) SetCoolingOutflowNil(b bool)`

 SetCoolingOutflowNil sets the value for CoolingOutflow to be an explicit nil

### UnsetCoolingOutflow
`func (o *CoolingIntake) UnsetCoolingOutflow()`

UnsetCoolingOutflow ensures that no value is present for CoolingOutflow, not even an explicit nil
### GetDescription

`func (o *CoolingIntake) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoolingIntake) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoolingIntake) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoolingIntake) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *CoolingIntake) GetOwner() BriefOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CoolingIntake) GetOwnerOk() (*BriefOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CoolingIntake) SetOwner(v BriefOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CoolingIntake) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *CoolingIntake) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *CoolingIntake) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *CoolingIntake) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CoolingIntake) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CoolingIntake) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CoolingIntake) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CoolingIntake) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CoolingIntake) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CoolingIntake) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CoolingIntake) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *CoolingIntake) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CoolingIntake) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CoolingIntake) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CoolingIntake) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CoolingIntake) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CoolingIntake) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CoolingIntake) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CoolingIntake) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CoolingIntake) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CoolingIntake) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


