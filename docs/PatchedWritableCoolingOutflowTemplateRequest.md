# PatchedWritableCoolingOutflowTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBulkConsolePortTemplateRequestDeviceType**](BulkConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBulkConsolePortTemplateRequestModuleType**](BulkConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | Pointer to **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullableBulkCoolingIntakeRequestType**](BulkCoolingIntakeRequestType.md) |  | [optional] 
**Diameter** | Pointer to **NullableFloat64** |  | [optional] 
**DiameterUnit** | Pointer to [**NullableBulkCoolingIntakeRequestDiameterUnit**](BulkCoolingIntakeRequestDiameterUnit.md) |  | [optional] 
**CoolingIntake** | Pointer to [**NullableBulkCoolingOutflowTemplateRequestCoolingIntake**](BulkCoolingOutflowTemplateRequestCoolingIntake.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedWritableCoolingOutflowTemplateRequest

`func NewPatchedWritableCoolingOutflowTemplateRequest() *PatchedWritableCoolingOutflowTemplateRequest`

NewPatchedWritableCoolingOutflowTemplateRequest instantiates a new PatchedWritableCoolingOutflowTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCoolingOutflowTemplateRequestWithDefaults

`func NewPatchedWritableCoolingOutflowTemplateRequestWithDefaults() *PatchedWritableCoolingOutflowTemplateRequest`

NewPatchedWritableCoolingOutflowTemplateRequestWithDefaults instantiates a new PatchedWritableCoolingOutflowTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *PatchedWritableCoolingOutflowTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *PatchedWritableCoolingOutflowTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetType() BulkCoolingIntakeRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetTypeOk() (*BulkCoolingIntakeRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetType(v BulkCoolingIntakeRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchedWritableCoolingOutflowTemplateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDiameter

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetDiameter() float64`

GetDiameter returns the Diameter field if non-nil, zero value otherwise.

### GetDiameterOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetDiameterOk() (*float64, bool)`

GetDiameterOk returns a tuple with the Diameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameter

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetDiameter(v float64)`

SetDiameter sets Diameter field to given value.

### HasDiameter

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasDiameter() bool`

HasDiameter returns a boolean if a field has been set.

### SetDiameterNil

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetDiameterNil(b bool)`

 SetDiameterNil sets the value for Diameter to be an explicit nil

### UnsetDiameter
`func (o *PatchedWritableCoolingOutflowTemplateRequest) UnsetDiameter()`

UnsetDiameter ensures that no value is present for Diameter, not even an explicit nil
### GetDiameterUnit

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetDiameterUnit() BulkCoolingIntakeRequestDiameterUnit`

GetDiameterUnit returns the DiameterUnit field if non-nil, zero value otherwise.

### GetDiameterUnitOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetDiameterUnitOk() (*BulkCoolingIntakeRequestDiameterUnit, bool)`

GetDiameterUnitOk returns a tuple with the DiameterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterUnit

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetDiameterUnit(v BulkCoolingIntakeRequestDiameterUnit)`

SetDiameterUnit sets DiameterUnit field to given value.

### HasDiameterUnit

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasDiameterUnit() bool`

HasDiameterUnit returns a boolean if a field has been set.

### SetDiameterUnitNil

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetDiameterUnitNil(b bool)`

 SetDiameterUnitNil sets the value for DiameterUnit to be an explicit nil

### UnsetDiameterUnit
`func (o *PatchedWritableCoolingOutflowTemplateRequest) UnsetDiameterUnit()`

UnsetDiameterUnit ensures that no value is present for DiameterUnit, not even an explicit nil
### GetCoolingIntake

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetCoolingIntake() BulkCoolingOutflowTemplateRequestCoolingIntake`

GetCoolingIntake returns the CoolingIntake field if non-nil, zero value otherwise.

### GetCoolingIntakeOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetCoolingIntakeOk() (*BulkCoolingOutflowTemplateRequestCoolingIntake, bool)`

GetCoolingIntakeOk returns a tuple with the CoolingIntake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingIntake

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetCoolingIntake(v BulkCoolingOutflowTemplateRequestCoolingIntake)`

SetCoolingIntake sets CoolingIntake field to given value.

### HasCoolingIntake

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasCoolingIntake() bool`

HasCoolingIntake returns a boolean if a field has been set.

### SetCoolingIntakeNil

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetCoolingIntakeNil(b bool)`

 SetCoolingIntakeNil sets the value for CoolingIntake to be an explicit nil

### UnsetCoolingIntake
`func (o *PatchedWritableCoolingOutflowTemplateRequest) UnsetCoolingIntake()`

UnsetCoolingIntake ensures that no value is present for CoolingIntake, not even an explicit nil
### GetDescription

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableCoolingOutflowTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableCoolingOutflowTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableCoolingOutflowTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


