# BulkDeviceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Manufacturer** | [**BriefDeviceTypeRequestManufacturer**](BriefDeviceTypeRequestManufacturer.md) |  | 
**DefaultPlatform** | Pointer to [**NullableBulkDeviceRequestPlatform**](BulkDeviceRequestPlatform.md) |  | [optional] 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**UHeight** | Pointer to **float64** |  | [optional] [default to 1.0]
**ExcludeFromUtilization** | Pointer to **bool** | Devices of this type are excluded when calculating rack utilization. | [optional] 
**IsFullDepth** | Pointer to **bool** | Device consumes both front and rear rack faces. | [optional] 
**SubdeviceRole** | Pointer to [**NullableBulkDeviceTypeRequestSubdeviceRole**](BulkDeviceTypeRequestSubdeviceRole.md) |  | [optional] 
**Airflow** | Pointer to [**NullableBulkDeviceTypeRequestAirflow**](BulkDeviceTypeRequestAirflow.md) |  | [optional] 
**CoolingMethod** | Pointer to [**NullableBulkDeviceRequestCoolingMethod**](BulkDeviceRequestCoolingMethod.md) |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**WeightUnit** | Pointer to [**NullableBulkDeviceTypeRequestWeightUnit**](BulkDeviceTypeRequestWeightUnit.md) |  | [optional] 
**EndOfLife** | Pointer to **NullableString** | The date after which this device type is no longer supported by the manufacturer | [optional] 
**FrontImage** | Pointer to ***os.File** |  | [optional] 
**RearImage** | Pointer to ***os.File** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkDeviceTypeRequest

`func NewBulkDeviceTypeRequest(id int32, manufacturer BriefDeviceTypeRequestManufacturer, model string, slug string, ) *BulkDeviceTypeRequest`

NewBulkDeviceTypeRequest instantiates a new BulkDeviceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeviceTypeRequestWithDefaults

`func NewBulkDeviceTypeRequestWithDefaults() *BulkDeviceTypeRequest`

NewBulkDeviceTypeRequestWithDefaults instantiates a new BulkDeviceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkDeviceTypeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkDeviceTypeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkDeviceTypeRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetManufacturer

`func (o *BulkDeviceTypeRequest) GetManufacturer() BriefDeviceTypeRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BulkDeviceTypeRequest) GetManufacturerOk() (*BriefDeviceTypeRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BulkDeviceTypeRequest) SetManufacturer(v BriefDeviceTypeRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetDefaultPlatform

`func (o *BulkDeviceTypeRequest) GetDefaultPlatform() BulkDeviceRequestPlatform`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *BulkDeviceTypeRequest) GetDefaultPlatformOk() (*BulkDeviceRequestPlatform, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *BulkDeviceTypeRequest) SetDefaultPlatform(v BulkDeviceRequestPlatform)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *BulkDeviceTypeRequest) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *BulkDeviceTypeRequest) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *BulkDeviceTypeRequest) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetModel

`func (o *BulkDeviceTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BulkDeviceTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BulkDeviceTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *BulkDeviceTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BulkDeviceTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BulkDeviceTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPartNumber

`func (o *BulkDeviceTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BulkDeviceTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BulkDeviceTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BulkDeviceTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetUHeight

`func (o *BulkDeviceTypeRequest) GetUHeight() float64`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *BulkDeviceTypeRequest) GetUHeightOk() (*float64, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *BulkDeviceTypeRequest) SetUHeight(v float64)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *BulkDeviceTypeRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetExcludeFromUtilization

`func (o *BulkDeviceTypeRequest) GetExcludeFromUtilization() bool`

GetExcludeFromUtilization returns the ExcludeFromUtilization field if non-nil, zero value otherwise.

### GetExcludeFromUtilizationOk

`func (o *BulkDeviceTypeRequest) GetExcludeFromUtilizationOk() (*bool, bool)`

GetExcludeFromUtilizationOk returns a tuple with the ExcludeFromUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromUtilization

`func (o *BulkDeviceTypeRequest) SetExcludeFromUtilization(v bool)`

SetExcludeFromUtilization sets ExcludeFromUtilization field to given value.

### HasExcludeFromUtilization

`func (o *BulkDeviceTypeRequest) HasExcludeFromUtilization() bool`

HasExcludeFromUtilization returns a boolean if a field has been set.

### GetIsFullDepth

`func (o *BulkDeviceTypeRequest) GetIsFullDepth() bool`

GetIsFullDepth returns the IsFullDepth field if non-nil, zero value otherwise.

### GetIsFullDepthOk

`func (o *BulkDeviceTypeRequest) GetIsFullDepthOk() (*bool, bool)`

GetIsFullDepthOk returns a tuple with the IsFullDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDepth

`func (o *BulkDeviceTypeRequest) SetIsFullDepth(v bool)`

SetIsFullDepth sets IsFullDepth field to given value.

### HasIsFullDepth

`func (o *BulkDeviceTypeRequest) HasIsFullDepth() bool`

HasIsFullDepth returns a boolean if a field has been set.

### GetSubdeviceRole

`func (o *BulkDeviceTypeRequest) GetSubdeviceRole() BulkDeviceTypeRequestSubdeviceRole`

GetSubdeviceRole returns the SubdeviceRole field if non-nil, zero value otherwise.

### GetSubdeviceRoleOk

`func (o *BulkDeviceTypeRequest) GetSubdeviceRoleOk() (*BulkDeviceTypeRequestSubdeviceRole, bool)`

GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdeviceRole

`func (o *BulkDeviceTypeRequest) SetSubdeviceRole(v BulkDeviceTypeRequestSubdeviceRole)`

SetSubdeviceRole sets SubdeviceRole field to given value.

### HasSubdeviceRole

`func (o *BulkDeviceTypeRequest) HasSubdeviceRole() bool`

HasSubdeviceRole returns a boolean if a field has been set.

### SetSubdeviceRoleNil

`func (o *BulkDeviceTypeRequest) SetSubdeviceRoleNil(b bool)`

 SetSubdeviceRoleNil sets the value for SubdeviceRole to be an explicit nil

### UnsetSubdeviceRole
`func (o *BulkDeviceTypeRequest) UnsetSubdeviceRole()`

UnsetSubdeviceRole ensures that no value is present for SubdeviceRole, not even an explicit nil
### GetAirflow

`func (o *BulkDeviceTypeRequest) GetAirflow() BulkDeviceTypeRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *BulkDeviceTypeRequest) GetAirflowOk() (*BulkDeviceTypeRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *BulkDeviceTypeRequest) SetAirflow(v BulkDeviceTypeRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *BulkDeviceTypeRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *BulkDeviceTypeRequest) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *BulkDeviceTypeRequest) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetCoolingMethod

`func (o *BulkDeviceTypeRequest) GetCoolingMethod() BulkDeviceRequestCoolingMethod`

GetCoolingMethod returns the CoolingMethod field if non-nil, zero value otherwise.

### GetCoolingMethodOk

`func (o *BulkDeviceTypeRequest) GetCoolingMethodOk() (*BulkDeviceRequestCoolingMethod, bool)`

GetCoolingMethodOk returns a tuple with the CoolingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingMethod

`func (o *BulkDeviceTypeRequest) SetCoolingMethod(v BulkDeviceRequestCoolingMethod)`

SetCoolingMethod sets CoolingMethod field to given value.

### HasCoolingMethod

`func (o *BulkDeviceTypeRequest) HasCoolingMethod() bool`

HasCoolingMethod returns a boolean if a field has been set.

### SetCoolingMethodNil

`func (o *BulkDeviceTypeRequest) SetCoolingMethodNil(b bool)`

 SetCoolingMethodNil sets the value for CoolingMethod to be an explicit nil

### UnsetCoolingMethod
`func (o *BulkDeviceTypeRequest) UnsetCoolingMethod()`

UnsetCoolingMethod ensures that no value is present for CoolingMethod, not even an explicit nil
### GetWeight

`func (o *BulkDeviceTypeRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkDeviceTypeRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkDeviceTypeRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkDeviceTypeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *BulkDeviceTypeRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *BulkDeviceTypeRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *BulkDeviceTypeRequest) GetWeightUnit() BulkDeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *BulkDeviceTypeRequest) GetWeightUnitOk() (*BulkDeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *BulkDeviceTypeRequest) SetWeightUnit(v BulkDeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *BulkDeviceTypeRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *BulkDeviceTypeRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *BulkDeviceTypeRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetEndOfLife

`func (o *BulkDeviceTypeRequest) GetEndOfLife() string`

GetEndOfLife returns the EndOfLife field if non-nil, zero value otherwise.

### GetEndOfLifeOk

`func (o *BulkDeviceTypeRequest) GetEndOfLifeOk() (*string, bool)`

GetEndOfLifeOk returns a tuple with the EndOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLife

`func (o *BulkDeviceTypeRequest) SetEndOfLife(v string)`

SetEndOfLife sets EndOfLife field to given value.

### HasEndOfLife

`func (o *BulkDeviceTypeRequest) HasEndOfLife() bool`

HasEndOfLife returns a boolean if a field has been set.

### SetEndOfLifeNil

`func (o *BulkDeviceTypeRequest) SetEndOfLifeNil(b bool)`

 SetEndOfLifeNil sets the value for EndOfLife to be an explicit nil

### UnsetEndOfLife
`func (o *BulkDeviceTypeRequest) UnsetEndOfLife()`

UnsetEndOfLife ensures that no value is present for EndOfLife, not even an explicit nil
### GetFrontImage

`func (o *BulkDeviceTypeRequest) GetFrontImage() *os.File`

GetFrontImage returns the FrontImage field if non-nil, zero value otherwise.

### GetFrontImageOk

`func (o *BulkDeviceTypeRequest) GetFrontImageOk() (**os.File, bool)`

GetFrontImageOk returns a tuple with the FrontImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImage

`func (o *BulkDeviceTypeRequest) SetFrontImage(v *os.File)`

SetFrontImage sets FrontImage field to given value.

### HasFrontImage

`func (o *BulkDeviceTypeRequest) HasFrontImage() bool`

HasFrontImage returns a boolean if a field has been set.

### GetRearImage

`func (o *BulkDeviceTypeRequest) GetRearImage() *os.File`

GetRearImage returns the RearImage field if non-nil, zero value otherwise.

### GetRearImageOk

`func (o *BulkDeviceTypeRequest) GetRearImageOk() (**os.File, bool)`

GetRearImageOk returns a tuple with the RearImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearImage

`func (o *BulkDeviceTypeRequest) SetRearImage(v *os.File)`

SetRearImage sets RearImage field to given value.

### HasRearImage

`func (o *BulkDeviceTypeRequest) HasRearImage() bool`

HasRearImage returns a boolean if a field has been set.

### GetDescription

`func (o *BulkDeviceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkDeviceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkDeviceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkDeviceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkDeviceTypeRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkDeviceTypeRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkDeviceTypeRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkDeviceTypeRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkDeviceTypeRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkDeviceTypeRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkDeviceTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkDeviceTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkDeviceTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkDeviceTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkDeviceTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkDeviceTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkDeviceTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkDeviceTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkDeviceTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkDeviceTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkDeviceTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkDeviceTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


