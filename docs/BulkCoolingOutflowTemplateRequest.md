# BulkCoolingOutflowTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
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

### NewBulkCoolingOutflowTemplateRequest

`func NewBulkCoolingOutflowTemplateRequest(id int32, name string, ) *BulkCoolingOutflowTemplateRequest`

NewBulkCoolingOutflowTemplateRequest instantiates a new BulkCoolingOutflowTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCoolingOutflowTemplateRequestWithDefaults

`func NewBulkCoolingOutflowTemplateRequestWithDefaults() *BulkCoolingOutflowTemplateRequest`

NewBulkCoolingOutflowTemplateRequestWithDefaults instantiates a new BulkCoolingOutflowTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCoolingOutflowTemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCoolingOutflowTemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCoolingOutflowTemplateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *BulkCoolingOutflowTemplateRequest) GetDeviceType() BulkConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkCoolingOutflowTemplateRequest) GetDeviceTypeOk() (*BulkConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkCoolingOutflowTemplateRequest) SetDeviceType(v BulkConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *BulkCoolingOutflowTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *BulkCoolingOutflowTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *BulkCoolingOutflowTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *BulkCoolingOutflowTemplateRequest) GetModuleType() BulkConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *BulkCoolingOutflowTemplateRequest) GetModuleTypeOk() (*BulkConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *BulkCoolingOutflowTemplateRequest) SetModuleType(v BulkConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *BulkCoolingOutflowTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *BulkCoolingOutflowTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *BulkCoolingOutflowTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *BulkCoolingOutflowTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkCoolingOutflowTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkCoolingOutflowTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkCoolingOutflowTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkCoolingOutflowTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkCoolingOutflowTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkCoolingOutflowTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *BulkCoolingOutflowTemplateRequest) GetType() BulkCoolingIntakeRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkCoolingOutflowTemplateRequest) GetTypeOk() (*BulkCoolingIntakeRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkCoolingOutflowTemplateRequest) SetType(v BulkCoolingIntakeRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *BulkCoolingOutflowTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BulkCoolingOutflowTemplateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BulkCoolingOutflowTemplateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDiameter

`func (o *BulkCoolingOutflowTemplateRequest) GetDiameter() float64`

GetDiameter returns the Diameter field if non-nil, zero value otherwise.

### GetDiameterOk

`func (o *BulkCoolingOutflowTemplateRequest) GetDiameterOk() (*float64, bool)`

GetDiameterOk returns a tuple with the Diameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameter

`func (o *BulkCoolingOutflowTemplateRequest) SetDiameter(v float64)`

SetDiameter sets Diameter field to given value.

### HasDiameter

`func (o *BulkCoolingOutflowTemplateRequest) HasDiameter() bool`

HasDiameter returns a boolean if a field has been set.

### SetDiameterNil

`func (o *BulkCoolingOutflowTemplateRequest) SetDiameterNil(b bool)`

 SetDiameterNil sets the value for Diameter to be an explicit nil

### UnsetDiameter
`func (o *BulkCoolingOutflowTemplateRequest) UnsetDiameter()`

UnsetDiameter ensures that no value is present for Diameter, not even an explicit nil
### GetDiameterUnit

`func (o *BulkCoolingOutflowTemplateRequest) GetDiameterUnit() BulkCoolingIntakeRequestDiameterUnit`

GetDiameterUnit returns the DiameterUnit field if non-nil, zero value otherwise.

### GetDiameterUnitOk

`func (o *BulkCoolingOutflowTemplateRequest) GetDiameterUnitOk() (*BulkCoolingIntakeRequestDiameterUnit, bool)`

GetDiameterUnitOk returns a tuple with the DiameterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiameterUnit

`func (o *BulkCoolingOutflowTemplateRequest) SetDiameterUnit(v BulkCoolingIntakeRequestDiameterUnit)`

SetDiameterUnit sets DiameterUnit field to given value.

### HasDiameterUnit

`func (o *BulkCoolingOutflowTemplateRequest) HasDiameterUnit() bool`

HasDiameterUnit returns a boolean if a field has been set.

### SetDiameterUnitNil

`func (o *BulkCoolingOutflowTemplateRequest) SetDiameterUnitNil(b bool)`

 SetDiameterUnitNil sets the value for DiameterUnit to be an explicit nil

### UnsetDiameterUnit
`func (o *BulkCoolingOutflowTemplateRequest) UnsetDiameterUnit()`

UnsetDiameterUnit ensures that no value is present for DiameterUnit, not even an explicit nil
### GetCoolingIntake

`func (o *BulkCoolingOutflowTemplateRequest) GetCoolingIntake() BulkCoolingOutflowTemplateRequestCoolingIntake`

GetCoolingIntake returns the CoolingIntake field if non-nil, zero value otherwise.

### GetCoolingIntakeOk

`func (o *BulkCoolingOutflowTemplateRequest) GetCoolingIntakeOk() (*BulkCoolingOutflowTemplateRequestCoolingIntake, bool)`

GetCoolingIntakeOk returns a tuple with the CoolingIntake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingIntake

`func (o *BulkCoolingOutflowTemplateRequest) SetCoolingIntake(v BulkCoolingOutflowTemplateRequestCoolingIntake)`

SetCoolingIntake sets CoolingIntake field to given value.

### HasCoolingIntake

`func (o *BulkCoolingOutflowTemplateRequest) HasCoolingIntake() bool`

HasCoolingIntake returns a boolean if a field has been set.

### SetCoolingIntakeNil

`func (o *BulkCoolingOutflowTemplateRequest) SetCoolingIntakeNil(b bool)`

 SetCoolingIntakeNil sets the value for CoolingIntake to be an explicit nil

### UnsetCoolingIntake
`func (o *BulkCoolingOutflowTemplateRequest) UnsetCoolingIntake()`

UnsetCoolingIntake ensures that no value is present for CoolingIntake, not even an explicit nil
### GetDescription

`func (o *BulkCoolingOutflowTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkCoolingOutflowTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkCoolingOutflowTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkCoolingOutflowTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


