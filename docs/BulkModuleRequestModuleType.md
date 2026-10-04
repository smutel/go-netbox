# BulkModuleRequestModuleType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**NullableBriefModuleTypeRequestProfile**](BriefModuleTypeRequestProfile.md) |  | [optional] 
**Manufacturer** | [**BriefDeviceTypeRequestManufacturer**](BriefDeviceTypeRequestManufacturer.md) |  | 
**Model** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkModuleRequestModuleType

`func NewBulkModuleRequestModuleType(manufacturer BriefDeviceTypeRequestManufacturer, model string, ) *BulkModuleRequestModuleType`

NewBulkModuleRequestModuleType instantiates a new BulkModuleRequestModuleType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkModuleRequestModuleTypeWithDefaults

`func NewBulkModuleRequestModuleTypeWithDefaults() *BulkModuleRequestModuleType`

NewBulkModuleRequestModuleTypeWithDefaults instantiates a new BulkModuleRequestModuleType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *BulkModuleRequestModuleType) GetProfile() BriefModuleTypeRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *BulkModuleRequestModuleType) GetProfileOk() (*BriefModuleTypeRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *BulkModuleRequestModuleType) SetProfile(v BriefModuleTypeRequestProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *BulkModuleRequestModuleType) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *BulkModuleRequestModuleType) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *BulkModuleRequestModuleType) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetManufacturer

`func (o *BulkModuleRequestModuleType) GetManufacturer() BriefDeviceTypeRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BulkModuleRequestModuleType) GetManufacturerOk() (*BriefDeviceTypeRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BulkModuleRequestModuleType) SetManufacturer(v BriefDeviceTypeRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *BulkModuleRequestModuleType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BulkModuleRequestModuleType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BulkModuleRequestModuleType) SetModel(v string)`

SetModel sets Model field to given value.


### GetDescription

`func (o *BulkModuleRequestModuleType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkModuleRequestModuleType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkModuleRequestModuleType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkModuleRequestModuleType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


