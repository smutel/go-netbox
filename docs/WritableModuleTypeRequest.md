# WritableModuleTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**NullableBriefModuleTypeRequestProfile**](BriefModuleTypeRequestProfile.md) |  | [optional] 
**Manufacturer** | [**BriefDeviceTypeRequestManufacturer**](BriefDeviceTypeRequestManufacturer.md) |  | 
**Model** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**Airflow** | Pointer to [**NullableBulkModuleTypeRequestAirflow**](BulkModuleTypeRequestAirflow.md) |  | [optional] 
**CoolingMethod** | Pointer to [**NullableBulkDeviceRequestCoolingMethod**](BulkDeviceRequestCoolingMethod.md) |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**WeightUnit** | Pointer to [**NullableBulkDeviceTypeRequestWeightUnit**](BulkDeviceTypeRequestWeightUnit.md) |  | [optional] 
**EndOfLife** | Pointer to **NullableString** | The date after which this module type is no longer supported by the manufacturer | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **interface{}** |  | [optional] 
**ModuleBayTypes** | Pointer to **[]int32** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableModuleTypeRequest

`func NewWritableModuleTypeRequest(manufacturer BriefDeviceTypeRequestManufacturer, model string, ) *WritableModuleTypeRequest`

NewWritableModuleTypeRequest instantiates a new WritableModuleTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableModuleTypeRequestWithDefaults

`func NewWritableModuleTypeRequestWithDefaults() *WritableModuleTypeRequest`

NewWritableModuleTypeRequestWithDefaults instantiates a new WritableModuleTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *WritableModuleTypeRequest) GetProfile() BriefModuleTypeRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WritableModuleTypeRequest) GetProfileOk() (*BriefModuleTypeRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WritableModuleTypeRequest) SetProfile(v BriefModuleTypeRequestProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *WritableModuleTypeRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *WritableModuleTypeRequest) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *WritableModuleTypeRequest) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetManufacturer

`func (o *WritableModuleTypeRequest) GetManufacturer() BriefDeviceTypeRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *WritableModuleTypeRequest) GetManufacturerOk() (*BriefDeviceTypeRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *WritableModuleTypeRequest) SetManufacturer(v BriefDeviceTypeRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *WritableModuleTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *WritableModuleTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *WritableModuleTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPartNumber

`func (o *WritableModuleTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *WritableModuleTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *WritableModuleTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *WritableModuleTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetAirflow

`func (o *WritableModuleTypeRequest) GetAirflow() BulkModuleTypeRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *WritableModuleTypeRequest) GetAirflowOk() (*BulkModuleTypeRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *WritableModuleTypeRequest) SetAirflow(v BulkModuleTypeRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *WritableModuleTypeRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *WritableModuleTypeRequest) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *WritableModuleTypeRequest) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetCoolingMethod

`func (o *WritableModuleTypeRequest) GetCoolingMethod() BulkDeviceRequestCoolingMethod`

GetCoolingMethod returns the CoolingMethod field if non-nil, zero value otherwise.

### GetCoolingMethodOk

`func (o *WritableModuleTypeRequest) GetCoolingMethodOk() (*BulkDeviceRequestCoolingMethod, bool)`

GetCoolingMethodOk returns a tuple with the CoolingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingMethod

`func (o *WritableModuleTypeRequest) SetCoolingMethod(v BulkDeviceRequestCoolingMethod)`

SetCoolingMethod sets CoolingMethod field to given value.

### HasCoolingMethod

`func (o *WritableModuleTypeRequest) HasCoolingMethod() bool`

HasCoolingMethod returns a boolean if a field has been set.

### SetCoolingMethodNil

`func (o *WritableModuleTypeRequest) SetCoolingMethodNil(b bool)`

 SetCoolingMethodNil sets the value for CoolingMethod to be an explicit nil

### UnsetCoolingMethod
`func (o *WritableModuleTypeRequest) UnsetCoolingMethod()`

UnsetCoolingMethod ensures that no value is present for CoolingMethod, not even an explicit nil
### GetWeight

`func (o *WritableModuleTypeRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WritableModuleTypeRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WritableModuleTypeRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WritableModuleTypeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *WritableModuleTypeRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *WritableModuleTypeRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *WritableModuleTypeRequest) GetWeightUnit() BulkDeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *WritableModuleTypeRequest) GetWeightUnitOk() (*BulkDeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *WritableModuleTypeRequest) SetWeightUnit(v BulkDeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *WritableModuleTypeRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *WritableModuleTypeRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *WritableModuleTypeRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetEndOfLife

`func (o *WritableModuleTypeRequest) GetEndOfLife() string`

GetEndOfLife returns the EndOfLife field if non-nil, zero value otherwise.

### GetEndOfLifeOk

`func (o *WritableModuleTypeRequest) GetEndOfLifeOk() (*string, bool)`

GetEndOfLifeOk returns a tuple with the EndOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLife

`func (o *WritableModuleTypeRequest) SetEndOfLife(v string)`

SetEndOfLife sets EndOfLife field to given value.

### HasEndOfLife

`func (o *WritableModuleTypeRequest) HasEndOfLife() bool`

HasEndOfLife returns a boolean if a field has been set.

### SetEndOfLifeNil

`func (o *WritableModuleTypeRequest) SetEndOfLifeNil(b bool)`

 SetEndOfLifeNil sets the value for EndOfLife to be an explicit nil

### UnsetEndOfLife
`func (o *WritableModuleTypeRequest) UnsetEndOfLife()`

UnsetEndOfLife ensures that no value is present for EndOfLife, not even an explicit nil
### GetDescription

`func (o *WritableModuleTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableModuleTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableModuleTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableModuleTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributes

`func (o *WritableModuleTypeRequest) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WritableModuleTypeRequest) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WritableModuleTypeRequest) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WritableModuleTypeRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *WritableModuleTypeRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *WritableModuleTypeRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetModuleBayTypes

`func (o *WritableModuleTypeRequest) GetModuleBayTypes() []int32`

GetModuleBayTypes returns the ModuleBayTypes field if non-nil, zero value otherwise.

### GetModuleBayTypesOk

`func (o *WritableModuleTypeRequest) GetModuleBayTypesOk() (*[]int32, bool)`

GetModuleBayTypesOk returns a tuple with the ModuleBayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTypes

`func (o *WritableModuleTypeRequest) SetModuleBayTypes(v []int32)`

SetModuleBayTypes sets ModuleBayTypes field to given value.

### HasModuleBayTypes

`func (o *WritableModuleTypeRequest) HasModuleBayTypes() bool`

HasModuleBayTypes returns a boolean if a field has been set.

### GetOwner

`func (o *WritableModuleTypeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WritableModuleTypeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WritableModuleTypeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WritableModuleTypeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WritableModuleTypeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WritableModuleTypeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *WritableModuleTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableModuleTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableModuleTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableModuleTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableModuleTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableModuleTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableModuleTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableModuleTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableModuleTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableModuleTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableModuleTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableModuleTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


