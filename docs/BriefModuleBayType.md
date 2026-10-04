# BriefModuleBayType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Manufacturer** | Pointer to [**NullableBriefManufacturer**](BriefManufacturer.md) |  | [optional] 
**Color** | Pointer to [**BriefModuleBayTypeColor**](BriefModuleBayTypeColor.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefModuleBayType

`func NewBriefModuleBayType(id int32, url string, display string, name string, slug string, ) *BriefModuleBayType`

NewBriefModuleBayType instantiates a new BriefModuleBayType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefModuleBayTypeWithDefaults

`func NewBriefModuleBayTypeWithDefaults() *BriefModuleBayType`

NewBriefModuleBayTypeWithDefaults instantiates a new BriefModuleBayType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefModuleBayType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefModuleBayType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefModuleBayType) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefModuleBayType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefModuleBayType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefModuleBayType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefModuleBayType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefModuleBayType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefModuleBayType) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefModuleBayType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefModuleBayType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefModuleBayType) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefModuleBayType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefModuleBayType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefModuleBayType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetManufacturer

`func (o *BriefModuleBayType) GetManufacturer() BriefManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BriefModuleBayType) GetManufacturerOk() (*BriefManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BriefModuleBayType) SetManufacturer(v BriefManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *BriefModuleBayType) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *BriefModuleBayType) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *BriefModuleBayType) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetColor

`func (o *BriefModuleBayType) GetColor() BriefModuleBayTypeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BriefModuleBayType) GetColorOk() (*BriefModuleBayTypeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BriefModuleBayType) SetColor(v BriefModuleBayTypeColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *BriefModuleBayType) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *BriefModuleBayType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefModuleBayType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefModuleBayType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefModuleBayType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


