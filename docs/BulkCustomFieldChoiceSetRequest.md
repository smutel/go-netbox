# BulkCustomFieldChoiceSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**BaseChoices** | Pointer to [**NullableBulkCustomFieldChoiceSetRequestBaseChoices**](BulkCustomFieldChoiceSetRequestBaseChoices.md) |  | [optional] 
**ExtraChoices** | **[][]interface{}** |  | 
**ChoiceColors** | Pointer to [**map[string]BulkCustomFieldChoiceSetRequestChoiceColorsValue**](BulkCustomFieldChoiceSetRequestChoiceColorsValue.md) |  | [optional] 
**OrderAlphabetically** | Pointer to **bool** | Choices are automatically ordered alphabetically | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewBulkCustomFieldChoiceSetRequest

`func NewBulkCustomFieldChoiceSetRequest(id int32, name string, extraChoices [][]interface{}, ) *BulkCustomFieldChoiceSetRequest`

NewBulkCustomFieldChoiceSetRequest instantiates a new BulkCustomFieldChoiceSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCustomFieldChoiceSetRequestWithDefaults

`func NewBulkCustomFieldChoiceSetRequestWithDefaults() *BulkCustomFieldChoiceSetRequest`

NewBulkCustomFieldChoiceSetRequestWithDefaults instantiates a new BulkCustomFieldChoiceSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCustomFieldChoiceSetRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCustomFieldChoiceSetRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCustomFieldChoiceSetRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BulkCustomFieldChoiceSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkCustomFieldChoiceSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkCustomFieldChoiceSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkCustomFieldChoiceSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkCustomFieldChoiceSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkCustomFieldChoiceSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkCustomFieldChoiceSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseChoices

`func (o *BulkCustomFieldChoiceSetRequest) GetBaseChoices() BulkCustomFieldChoiceSetRequestBaseChoices`

GetBaseChoices returns the BaseChoices field if non-nil, zero value otherwise.

### GetBaseChoicesOk

`func (o *BulkCustomFieldChoiceSetRequest) GetBaseChoicesOk() (*BulkCustomFieldChoiceSetRequestBaseChoices, bool)`

GetBaseChoicesOk returns a tuple with the BaseChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseChoices

`func (o *BulkCustomFieldChoiceSetRequest) SetBaseChoices(v BulkCustomFieldChoiceSetRequestBaseChoices)`

SetBaseChoices sets BaseChoices field to given value.

### HasBaseChoices

`func (o *BulkCustomFieldChoiceSetRequest) HasBaseChoices() bool`

HasBaseChoices returns a boolean if a field has been set.

### SetBaseChoicesNil

`func (o *BulkCustomFieldChoiceSetRequest) SetBaseChoicesNil(b bool)`

 SetBaseChoicesNil sets the value for BaseChoices to be an explicit nil

### UnsetBaseChoices
`func (o *BulkCustomFieldChoiceSetRequest) UnsetBaseChoices()`

UnsetBaseChoices ensures that no value is present for BaseChoices, not even an explicit nil
### GetExtraChoices

`func (o *BulkCustomFieldChoiceSetRequest) GetExtraChoices() [][]interface{}`

GetExtraChoices returns the ExtraChoices field if non-nil, zero value otherwise.

### GetExtraChoicesOk

`func (o *BulkCustomFieldChoiceSetRequest) GetExtraChoicesOk() (*[][]interface{}, bool)`

GetExtraChoicesOk returns a tuple with the ExtraChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraChoices

`func (o *BulkCustomFieldChoiceSetRequest) SetExtraChoices(v [][]interface{})`

SetExtraChoices sets ExtraChoices field to given value.


### GetChoiceColors

`func (o *BulkCustomFieldChoiceSetRequest) GetChoiceColors() map[string]BulkCustomFieldChoiceSetRequestChoiceColorsValue`

GetChoiceColors returns the ChoiceColors field if non-nil, zero value otherwise.

### GetChoiceColorsOk

`func (o *BulkCustomFieldChoiceSetRequest) GetChoiceColorsOk() (*map[string]BulkCustomFieldChoiceSetRequestChoiceColorsValue, bool)`

GetChoiceColorsOk returns a tuple with the ChoiceColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceColors

`func (o *BulkCustomFieldChoiceSetRequest) SetChoiceColors(v map[string]BulkCustomFieldChoiceSetRequestChoiceColorsValue)`

SetChoiceColors sets ChoiceColors field to given value.

### HasChoiceColors

`func (o *BulkCustomFieldChoiceSetRequest) HasChoiceColors() bool`

HasChoiceColors returns a boolean if a field has been set.

### GetOrderAlphabetically

`func (o *BulkCustomFieldChoiceSetRequest) GetOrderAlphabetically() bool`

GetOrderAlphabetically returns the OrderAlphabetically field if non-nil, zero value otherwise.

### GetOrderAlphabeticallyOk

`func (o *BulkCustomFieldChoiceSetRequest) GetOrderAlphabeticallyOk() (*bool, bool)`

GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAlphabetically

`func (o *BulkCustomFieldChoiceSetRequest) SetOrderAlphabetically(v bool)`

SetOrderAlphabetically sets OrderAlphabetically field to given value.

### HasOrderAlphabetically

`func (o *BulkCustomFieldChoiceSetRequest) HasOrderAlphabetically() bool`

HasOrderAlphabetically returns a boolean if a field has been set.

### GetOwner

`func (o *BulkCustomFieldChoiceSetRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkCustomFieldChoiceSetRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkCustomFieldChoiceSetRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkCustomFieldChoiceSetRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkCustomFieldChoiceSetRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkCustomFieldChoiceSetRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


