# ModuleType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Profile** | Pointer to [**NullableBriefModuleTypeProfile**](BriefModuleTypeProfile.md) |  | [optional] 
**Manufacturer** | [**BriefManufacturer**](BriefManufacturer.md) |  | 
**Model** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**Airflow** | Pointer to [**NullableModuleTypeAirflow**](ModuleTypeAirflow.md) |  | [optional] 
**CoolingMethod** | Pointer to [**NullableDeviceCoolingMethod**](DeviceCoolingMethod.md) |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**WeightUnit** | Pointer to [**NullableDeviceTypeWeightUnit**](DeviceTypeWeightUnit.md) |  | [optional] 
**EndOfLife** | Pointer to **NullableString** | The date after which this module type is no longer supported by the manufacturer | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **interface{}** |  | [optional] 
**ModuleBayTypes** | Pointer to [**[]BriefModuleBayType**](BriefModuleBayType.md) |  | [optional] 
**Owner** | Pointer to [**NullableBriefOwner**](BriefOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**ModuleCount** | **int32** |  | [readonly] 
**ConsolePortTemplateCount** | **int32** |  | [readonly] 
**ConsoleServerPortTemplateCount** | **int32** |  | [readonly] 
**PowerPortTemplateCount** | **int32** |  | [readonly] 
**PowerOutletTemplateCount** | **int32** |  | [readonly] 
**CoolingIntakeTemplateCount** | **int32** |  | [readonly] 
**CoolingOutflowTemplateCount** | **int32** |  | [readonly] 
**InterfaceTemplateCount** | **int32** |  | [readonly] 
**FrontPortTemplateCount** | **int32** |  | [readonly] 
**RearPortTemplateCount** | **int32** |  | [readonly] 
**ModuleBayTemplateCount** | **int32** |  | [readonly] 

## Methods

### NewModuleType

`func NewModuleType(id int32, url string, displayUrl string, display string, manufacturer BriefManufacturer, model string, created NullableTime, lastUpdated NullableTime, moduleCount int32, consolePortTemplateCount int32, consoleServerPortTemplateCount int32, powerPortTemplateCount int32, powerOutletTemplateCount int32, coolingIntakeTemplateCount int32, coolingOutflowTemplateCount int32, interfaceTemplateCount int32, frontPortTemplateCount int32, rearPortTemplateCount int32, moduleBayTemplateCount int32, ) *ModuleType`

NewModuleType instantiates a new ModuleType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleTypeWithDefaults

`func NewModuleTypeWithDefaults() *ModuleType`

NewModuleTypeWithDefaults instantiates a new ModuleType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModuleType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModuleType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModuleType) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ModuleType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModuleType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModuleType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *ModuleType) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ModuleType) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ModuleType) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *ModuleType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ModuleType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ModuleType) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetProfile

`func (o *ModuleType) GetProfile() BriefModuleTypeProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ModuleType) GetProfileOk() (*BriefModuleTypeProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ModuleType) SetProfile(v BriefModuleTypeProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ModuleType) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *ModuleType) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ModuleType) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetManufacturer

`func (o *ModuleType) GetManufacturer() BriefManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ModuleType) GetManufacturerOk() (*BriefManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ModuleType) SetManufacturer(v BriefManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *ModuleType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModuleType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModuleType) SetModel(v string)`

SetModel sets Model field to given value.


### GetPartNumber

`func (o *ModuleType) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *ModuleType) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *ModuleType) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *ModuleType) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetAirflow

`func (o *ModuleType) GetAirflow() ModuleTypeAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *ModuleType) GetAirflowOk() (*ModuleTypeAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *ModuleType) SetAirflow(v ModuleTypeAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *ModuleType) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *ModuleType) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *ModuleType) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetCoolingMethod

`func (o *ModuleType) GetCoolingMethod() DeviceCoolingMethod`

GetCoolingMethod returns the CoolingMethod field if non-nil, zero value otherwise.

### GetCoolingMethodOk

`func (o *ModuleType) GetCoolingMethodOk() (*DeviceCoolingMethod, bool)`

GetCoolingMethodOk returns a tuple with the CoolingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingMethod

`func (o *ModuleType) SetCoolingMethod(v DeviceCoolingMethod)`

SetCoolingMethod sets CoolingMethod field to given value.

### HasCoolingMethod

`func (o *ModuleType) HasCoolingMethod() bool`

HasCoolingMethod returns a boolean if a field has been set.

### SetCoolingMethodNil

`func (o *ModuleType) SetCoolingMethodNil(b bool)`

 SetCoolingMethodNil sets the value for CoolingMethod to be an explicit nil

### UnsetCoolingMethod
`func (o *ModuleType) UnsetCoolingMethod()`

UnsetCoolingMethod ensures that no value is present for CoolingMethod, not even an explicit nil
### GetWeight

`func (o *ModuleType) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ModuleType) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ModuleType) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ModuleType) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *ModuleType) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *ModuleType) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *ModuleType) GetWeightUnit() DeviceTypeWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ModuleType) GetWeightUnitOk() (*DeviceTypeWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ModuleType) SetWeightUnit(v DeviceTypeWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ModuleType) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *ModuleType) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *ModuleType) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetEndOfLife

`func (o *ModuleType) GetEndOfLife() string`

GetEndOfLife returns the EndOfLife field if non-nil, zero value otherwise.

### GetEndOfLifeOk

`func (o *ModuleType) GetEndOfLifeOk() (*string, bool)`

GetEndOfLifeOk returns a tuple with the EndOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLife

`func (o *ModuleType) SetEndOfLife(v string)`

SetEndOfLife sets EndOfLife field to given value.

### HasEndOfLife

`func (o *ModuleType) HasEndOfLife() bool`

HasEndOfLife returns a boolean if a field has been set.

### SetEndOfLifeNil

`func (o *ModuleType) SetEndOfLifeNil(b bool)`

 SetEndOfLifeNil sets the value for EndOfLife to be an explicit nil

### UnsetEndOfLife
`func (o *ModuleType) UnsetEndOfLife()`

UnsetEndOfLife ensures that no value is present for EndOfLife, not even an explicit nil
### GetDescription

`func (o *ModuleType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributes

`func (o *ModuleType) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ModuleType) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ModuleType) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ModuleType) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ModuleType) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ModuleType) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetModuleBayTypes

`func (o *ModuleType) GetModuleBayTypes() []BriefModuleBayType`

GetModuleBayTypes returns the ModuleBayTypes field if non-nil, zero value otherwise.

### GetModuleBayTypesOk

`func (o *ModuleType) GetModuleBayTypesOk() (*[]BriefModuleBayType, bool)`

GetModuleBayTypesOk returns a tuple with the ModuleBayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTypes

`func (o *ModuleType) SetModuleBayTypes(v []BriefModuleBayType)`

SetModuleBayTypes sets ModuleBayTypes field to given value.

### HasModuleBayTypes

`func (o *ModuleType) HasModuleBayTypes() bool`

HasModuleBayTypes returns a boolean if a field has been set.

### GetOwner

`func (o *ModuleType) GetOwner() BriefOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModuleType) GetOwnerOk() (*BriefOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModuleType) SetOwner(v BriefOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModuleType) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ModuleType) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ModuleType) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *ModuleType) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ModuleType) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ModuleType) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ModuleType) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *ModuleType) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModuleType) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModuleType) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModuleType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ModuleType) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModuleType) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModuleType) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModuleType) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *ModuleType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModuleType) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModuleType) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ModuleType) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ModuleType) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ModuleType) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ModuleType) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ModuleType) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ModuleType) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ModuleType) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetModuleCount

`func (o *ModuleType) GetModuleCount() int32`

GetModuleCount returns the ModuleCount field if non-nil, zero value otherwise.

### GetModuleCountOk

`func (o *ModuleType) GetModuleCountOk() (*int32, bool)`

GetModuleCountOk returns a tuple with the ModuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleCount

`func (o *ModuleType) SetModuleCount(v int32)`

SetModuleCount sets ModuleCount field to given value.


### GetConsolePortTemplateCount

`func (o *ModuleType) GetConsolePortTemplateCount() int32`

GetConsolePortTemplateCount returns the ConsolePortTemplateCount field if non-nil, zero value otherwise.

### GetConsolePortTemplateCountOk

`func (o *ModuleType) GetConsolePortTemplateCountOk() (*int32, bool)`

GetConsolePortTemplateCountOk returns a tuple with the ConsolePortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePortTemplateCount

`func (o *ModuleType) SetConsolePortTemplateCount(v int32)`

SetConsolePortTemplateCount sets ConsolePortTemplateCount field to given value.


### GetConsoleServerPortTemplateCount

`func (o *ModuleType) GetConsoleServerPortTemplateCount() int32`

GetConsoleServerPortTemplateCount returns the ConsoleServerPortTemplateCount field if non-nil, zero value otherwise.

### GetConsoleServerPortTemplateCountOk

`func (o *ModuleType) GetConsoleServerPortTemplateCountOk() (*int32, bool)`

GetConsoleServerPortTemplateCountOk returns a tuple with the ConsoleServerPortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPortTemplateCount

`func (o *ModuleType) SetConsoleServerPortTemplateCount(v int32)`

SetConsoleServerPortTemplateCount sets ConsoleServerPortTemplateCount field to given value.


### GetPowerPortTemplateCount

`func (o *ModuleType) GetPowerPortTemplateCount() int32`

GetPowerPortTemplateCount returns the PowerPortTemplateCount field if non-nil, zero value otherwise.

### GetPowerPortTemplateCountOk

`func (o *ModuleType) GetPowerPortTemplateCountOk() (*int32, bool)`

GetPowerPortTemplateCountOk returns a tuple with the PowerPortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPortTemplateCount

`func (o *ModuleType) SetPowerPortTemplateCount(v int32)`

SetPowerPortTemplateCount sets PowerPortTemplateCount field to given value.


### GetPowerOutletTemplateCount

`func (o *ModuleType) GetPowerOutletTemplateCount() int32`

GetPowerOutletTemplateCount returns the PowerOutletTemplateCount field if non-nil, zero value otherwise.

### GetPowerOutletTemplateCountOk

`func (o *ModuleType) GetPowerOutletTemplateCountOk() (*int32, bool)`

GetPowerOutletTemplateCountOk returns a tuple with the PowerOutletTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutletTemplateCount

`func (o *ModuleType) SetPowerOutletTemplateCount(v int32)`

SetPowerOutletTemplateCount sets PowerOutletTemplateCount field to given value.


### GetCoolingIntakeTemplateCount

`func (o *ModuleType) GetCoolingIntakeTemplateCount() int32`

GetCoolingIntakeTemplateCount returns the CoolingIntakeTemplateCount field if non-nil, zero value otherwise.

### GetCoolingIntakeTemplateCountOk

`func (o *ModuleType) GetCoolingIntakeTemplateCountOk() (*int32, bool)`

GetCoolingIntakeTemplateCountOk returns a tuple with the CoolingIntakeTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingIntakeTemplateCount

`func (o *ModuleType) SetCoolingIntakeTemplateCount(v int32)`

SetCoolingIntakeTemplateCount sets CoolingIntakeTemplateCount field to given value.


### GetCoolingOutflowTemplateCount

`func (o *ModuleType) GetCoolingOutflowTemplateCount() int32`

GetCoolingOutflowTemplateCount returns the CoolingOutflowTemplateCount field if non-nil, zero value otherwise.

### GetCoolingOutflowTemplateCountOk

`func (o *ModuleType) GetCoolingOutflowTemplateCountOk() (*int32, bool)`

GetCoolingOutflowTemplateCountOk returns a tuple with the CoolingOutflowTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingOutflowTemplateCount

`func (o *ModuleType) SetCoolingOutflowTemplateCount(v int32)`

SetCoolingOutflowTemplateCount sets CoolingOutflowTemplateCount field to given value.


### GetInterfaceTemplateCount

`func (o *ModuleType) GetInterfaceTemplateCount() int32`

GetInterfaceTemplateCount returns the InterfaceTemplateCount field if non-nil, zero value otherwise.

### GetInterfaceTemplateCountOk

`func (o *ModuleType) GetInterfaceTemplateCountOk() (*int32, bool)`

GetInterfaceTemplateCountOk returns a tuple with the InterfaceTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceTemplateCount

`func (o *ModuleType) SetInterfaceTemplateCount(v int32)`

SetInterfaceTemplateCount sets InterfaceTemplateCount field to given value.


### GetFrontPortTemplateCount

`func (o *ModuleType) GetFrontPortTemplateCount() int32`

GetFrontPortTemplateCount returns the FrontPortTemplateCount field if non-nil, zero value otherwise.

### GetFrontPortTemplateCountOk

`func (o *ModuleType) GetFrontPortTemplateCountOk() (*int32, bool)`

GetFrontPortTemplateCountOk returns a tuple with the FrontPortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPortTemplateCount

`func (o *ModuleType) SetFrontPortTemplateCount(v int32)`

SetFrontPortTemplateCount sets FrontPortTemplateCount field to given value.


### GetRearPortTemplateCount

`func (o *ModuleType) GetRearPortTemplateCount() int32`

GetRearPortTemplateCount returns the RearPortTemplateCount field if non-nil, zero value otherwise.

### GetRearPortTemplateCountOk

`func (o *ModuleType) GetRearPortTemplateCountOk() (*int32, bool)`

GetRearPortTemplateCountOk returns a tuple with the RearPortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortTemplateCount

`func (o *ModuleType) SetRearPortTemplateCount(v int32)`

SetRearPortTemplateCount sets RearPortTemplateCount field to given value.


### GetModuleBayTemplateCount

`func (o *ModuleType) GetModuleBayTemplateCount() int32`

GetModuleBayTemplateCount returns the ModuleBayTemplateCount field if non-nil, zero value otherwise.

### GetModuleBayTemplateCountOk

`func (o *ModuleType) GetModuleBayTemplateCountOk() (*int32, bool)`

GetModuleBayTemplateCountOk returns a tuple with the ModuleBayTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTemplateCount

`func (o *ModuleType) SetModuleBayTemplateCount(v int32)`

SetModuleBayTemplateCount sets ModuleBayTemplateCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


