# CoolingOutflowTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullableBulkCoolingIntakeRequestType**](BulkCoolingIntakeRequestType.md) |  | [optional] 
**Diameter** | Pointer to **NullableFloat64** |  | [optional] 
**DiameterUnit** | Pointer to [**NullableBulkCoolingIntakeRequestDiameterUnit**](BulkCoolingIntakeRequestDiameterUnit.md) |  | [optional] 
**CoolingIntake** | Pointer to [**NullableBulkCoolingOutflowTemplateRequestCoolingIntake**](BulkCoolingOutflowTemplateRequestCoolingIntake.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCoolingOutflowTemplateRequest

`func NewCoolingOutflowTemplateRequest(name string, ) *CoolingOutflowTemplateRequest`

NewCoolingOutflowTemplateRequest instantiates a new CoolingOutflowTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoolingOutflowTemplateRequestWithDefaults

`func NewCoolingOutflowTemplateRequestWithDefaults() *CoolingOutflowTemplateRequest`

NewCoolingOutflowTemplateRequestWithDefaults instantiates a new CoolingOutflowTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *CoolingOutflowTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CoolingOutflowTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CoolingOutflowTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *CoolingOutflowTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *CoolingOutflowTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *CoolingOutflowTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *CoolingOutflowTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *CoolingOutflowTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *CoolingOutflowTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *CoolingOutflowTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *CoolingOutflowTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *CoolingOutflowTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *CoolingOutflowTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoolingOutflowTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoolingOutflowTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CoolingOutflowTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CoolingOutflowTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CoolingOutflowTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CoolingOutflowTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *CoolingOutflowTemplateRequest) GetType() BulkCoolingIntakeRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoolingOutflowTemplateRequest) GetTypeOk() (*BulkCoolingIntakeRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoolingOutflowTemplateRequest) SetType(v BulkCoolingIntakeRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *CoolingOutflowTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CoolingOutflowTemplateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CoolingOutflowTemplateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDiameter

`func (o *CoolingOutflowTemplateRequest) GetDiameter() float64`

GetDiameter returns the Diameter field if non-nil, zero value otherwise.

### GetDiameterOk

`func (o *CoolingOutflowTemplateRequest) GetDiameterOk() (*float64, bool)`

GetDiameterOk returns a tuple with the Diameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameter

`func (o *CoolingOutflowTemplateRequest) SetDiameter(v float64)`

SetDiameter sets Diameter field to given value.

### HasDiameter

`func (o *CoolingOutflowTemplateRequest) HasDiameter() bool`

HasDiameter returns a boolean if a field has been set.

### SetDiameterNil

`func (o *CoolingOutflowTemplateRequest) SetDiameterNil(b bool)`

 SetDiameterNil sets the value for Diameter to be an explicit nil

### UnsetDiameter
`func (o *CoolingOutflowTemplateRequest) UnsetDiameter()`

UnsetDiameter ensures that no value is present for Diameter, not even an explicit nil
### GetDiameterUnit

`func (o *CoolingOutflowTemplateRequest) GetDiameterUnit() BulkCoolingIntakeRequestDiameterUnit`

GetDiameterUnit returns the DiameterUnit field if non-nil, zero value otherwise.

### GetDiameterUnitOk

`func (o *CoolingOutflowTemplateRequest) GetDiameterUnitOk() (*BulkCoolingIntakeRequestDiameterUnit, bool)`

GetDiameterUnitOk returns a tuple with the DiameterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterUnit

`func (o *CoolingOutflowTemplateRequest) SetDiameterUnit(v BulkCoolingIntakeRequestDiameterUnit)`

SetDiameterUnit sets DiameterUnit field to given value.

### HasDiameterUnit

`func (o *CoolingOutflowTemplateRequest) HasDiameterUnit() bool`

HasDiameterUnit returns a boolean if a field has been set.

### SetDiameterUnitNil

`func (o *CoolingOutflowTemplateRequest) SetDiameterUnitNil(b bool)`

 SetDiameterUnitNil sets the value for DiameterUnit to be an explicit nil

### UnsetDiameterUnit
`func (o *CoolingOutflowTemplateRequest) UnsetDiameterUnit()`

UnsetDiameterUnit ensures that no value is present for DiameterUnit, not even an explicit nil
### GetCoolingIntake

`func (o *CoolingOutflowTemplateRequest) GetCoolingIntake() BulkCoolingOutflowTemplateRequestCoolingIntake`

GetCoolingIntake returns the CoolingIntake field if non-nil, zero value otherwise.

### GetCoolingIntakeOk

`func (o *CoolingOutflowTemplateRequest) GetCoolingIntakeOk() (*BulkCoolingOutflowTemplateRequestCoolingIntake, bool)`

GetCoolingIntakeOk returns a tuple with the CoolingIntake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingIntake

`func (o *CoolingOutflowTemplateRequest) SetCoolingIntake(v BulkCoolingOutflowTemplateRequestCoolingIntake)`

SetCoolingIntake sets CoolingIntake field to given value.

### HasCoolingIntake

`func (o *CoolingOutflowTemplateRequest) HasCoolingIntake() bool`

HasCoolingIntake returns a boolean if a field has been set.

### SetCoolingIntakeNil

`func (o *CoolingOutflowTemplateRequest) SetCoolingIntakeNil(b bool)`

 SetCoolingIntakeNil sets the value for CoolingIntake to be an explicit nil

### UnsetCoolingIntake
`func (o *CoolingOutflowTemplateRequest) UnsetCoolingIntake()`

UnsetCoolingIntake ensures that no value is present for CoolingIntake, not even an explicit nil
### GetDescription

`func (o *CoolingOutflowTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoolingOutflowTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoolingOutflowTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoolingOutflowTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


