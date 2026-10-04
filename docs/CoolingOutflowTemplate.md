# CoolingOutflowTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**DeviceType** | Pointer to [**NullableBriefDeviceType**](BriefDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleType**](BriefModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullableCoolingIntakeType**](CoolingIntakeType.md) |  | [optional] 
**Diameter** | Pointer to **NullableFloat64** |  | [optional] 
**DiameterUnit** | Pointer to [**NullableCoolingIntakeDiameterUnit**](CoolingIntakeDiameterUnit.md) |  | [optional] 
**CoolingIntake** | Pointer to [**NullableBriefCoolingIntakeTemplate**](BriefCoolingIntakeTemplate.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewCoolingOutflowTemplate

`func NewCoolingOutflowTemplate(id int32, url string, display string, name string, created NullableTime, lastUpdated NullableTime, ) *CoolingOutflowTemplate`

NewCoolingOutflowTemplate instantiates a new CoolingOutflowTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoolingOutflowTemplateWithDefaults

`func NewCoolingOutflowTemplateWithDefaults() *CoolingOutflowTemplate`

NewCoolingOutflowTemplateWithDefaults instantiates a new CoolingOutflowTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoolingOutflowTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoolingOutflowTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoolingOutflowTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CoolingOutflowTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CoolingOutflowTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CoolingOutflowTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *CoolingOutflowTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CoolingOutflowTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CoolingOutflowTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDeviceType

`func (o *CoolingOutflowTemplate) GetDeviceType() BriefDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CoolingOutflowTemplate) GetDeviceTypeOk() (*BriefDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CoolingOutflowTemplate) SetDeviceType(v BriefDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *CoolingOutflowTemplate) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *CoolingOutflowTemplate) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *CoolingOutflowTemplate) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *CoolingOutflowTemplate) GetModuleType() BriefModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *CoolingOutflowTemplate) GetModuleTypeOk() (*BriefModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *CoolingOutflowTemplate) SetModuleType(v BriefModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *CoolingOutflowTemplate) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *CoolingOutflowTemplate) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *CoolingOutflowTemplate) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *CoolingOutflowTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoolingOutflowTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoolingOutflowTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CoolingOutflowTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CoolingOutflowTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CoolingOutflowTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CoolingOutflowTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *CoolingOutflowTemplate) GetType() CoolingIntakeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoolingOutflowTemplate) GetTypeOk() (*CoolingIntakeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoolingOutflowTemplate) SetType(v CoolingIntakeType)`

SetType sets Type field to given value.

### HasType

`func (o *CoolingOutflowTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CoolingOutflowTemplate) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CoolingOutflowTemplate) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDiameter

`func (o *CoolingOutflowTemplate) GetDiameter() float64`

GetDiameter returns the Diameter field if non-nil, zero value otherwise.

### GetDiameterOk

`func (o *CoolingOutflowTemplate) GetDiameterOk() (*float64, bool)`

GetDiameterOk returns a tuple with the Diameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameter

`func (o *CoolingOutflowTemplate) SetDiameter(v float64)`

SetDiameter sets Diameter field to given value.

### HasDiameter

`func (o *CoolingOutflowTemplate) HasDiameter() bool`

HasDiameter returns a boolean if a field has been set.

### SetDiameterNil

`func (o *CoolingOutflowTemplate) SetDiameterNil(b bool)`

 SetDiameterNil sets the value for Diameter to be an explicit nil

### UnsetDiameter
`func (o *CoolingOutflowTemplate) UnsetDiameter()`

UnsetDiameter ensures that no value is present for Diameter, not even an explicit nil
### GetDiameterUnit

`func (o *CoolingOutflowTemplate) GetDiameterUnit() CoolingIntakeDiameterUnit`

GetDiameterUnit returns the DiameterUnit field if non-nil, zero value otherwise.

### GetDiameterUnitOk

`func (o *CoolingOutflowTemplate) GetDiameterUnitOk() (*CoolingIntakeDiameterUnit, bool)`

GetDiameterUnitOk returns a tuple with the DiameterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterUnit

`func (o *CoolingOutflowTemplate) SetDiameterUnit(v CoolingIntakeDiameterUnit)`

SetDiameterUnit sets DiameterUnit field to given value.

### HasDiameterUnit

`func (o *CoolingOutflowTemplate) HasDiameterUnit() bool`

HasDiameterUnit returns a boolean if a field has been set.

### SetDiameterUnitNil

`func (o *CoolingOutflowTemplate) SetDiameterUnitNil(b bool)`

 SetDiameterUnitNil sets the value for DiameterUnit to be an explicit nil

### UnsetDiameterUnit
`func (o *CoolingOutflowTemplate) UnsetDiameterUnit()`

UnsetDiameterUnit ensures that no value is present for DiameterUnit, not even an explicit nil
### GetCoolingIntake

`func (o *CoolingOutflowTemplate) GetCoolingIntake() BriefCoolingIntakeTemplate`

GetCoolingIntake returns the CoolingIntake field if non-nil, zero value otherwise.

### GetCoolingIntakeOk

`func (o *CoolingOutflowTemplate) GetCoolingIntakeOk() (*BriefCoolingIntakeTemplate, bool)`

GetCoolingIntakeOk returns a tuple with the CoolingIntake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingIntake

`func (o *CoolingOutflowTemplate) SetCoolingIntake(v BriefCoolingIntakeTemplate)`

SetCoolingIntake sets CoolingIntake field to given value.

### HasCoolingIntake

`func (o *CoolingOutflowTemplate) HasCoolingIntake() bool`

HasCoolingIntake returns a boolean if a field has been set.

### SetCoolingIntakeNil

`func (o *CoolingOutflowTemplate) SetCoolingIntakeNil(b bool)`

 SetCoolingIntakeNil sets the value for CoolingIntake to be an explicit nil

### UnsetCoolingIntake
`func (o *CoolingOutflowTemplate) UnsetCoolingIntake()`

UnsetCoolingIntake ensures that no value is present for CoolingIntake, not even an explicit nil
### GetDescription

`func (o *CoolingOutflowTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoolingOutflowTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoolingOutflowTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoolingOutflowTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *CoolingOutflowTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CoolingOutflowTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CoolingOutflowTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CoolingOutflowTemplate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CoolingOutflowTemplate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CoolingOutflowTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CoolingOutflowTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CoolingOutflowTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CoolingOutflowTemplate) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CoolingOutflowTemplate) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


