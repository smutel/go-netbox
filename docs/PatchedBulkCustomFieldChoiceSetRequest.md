# PatchedBulkCustomFieldChoiceSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BaseChoices** | Pointer to [**NullableBulkCustomFieldChoiceSetRequestBaseChoices**](BulkCustomFieldChoiceSetRequestBaseChoices.md) |  | [optional] 
**ExtraChoices** | Pointer to **[][]interface{}** |  | [optional] 
**ChoiceColors** | Pointer to [**map[string]BulkCustomFieldChoiceSetRequestChoiceColorsValue**](BulkCustomFieldChoiceSetRequestChoiceColorsValue.md) |  | [optional] 
**OrderAlphabetically** | Pointer to **bool** | Choices are automatically ordered alphabetically | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewPatchedBulkCustomFieldChoiceSetRequest

`func NewPatchedBulkCustomFieldChoiceSetRequest(id int32, ) *PatchedBulkCustomFieldChoiceSetRequest`

NewPatchedBulkCustomFieldChoiceSetRequest instantiates a new PatchedBulkCustomFieldChoiceSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCustomFieldChoiceSetRequestWithDefaults

`func NewPatchedBulkCustomFieldChoiceSetRequestWithDefaults() *PatchedBulkCustomFieldChoiceSetRequest`

NewPatchedBulkCustomFieldChoiceSetRequestWithDefaults instantiates a new PatchedBulkCustomFieldChoiceSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkCustomFieldChoiceSetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkCustomFieldChoiceSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseChoices

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetBaseChoices() BulkCustomFieldChoiceSetRequestBaseChoices`

GetBaseChoices returns the BaseChoices field if non-nil, zero value otherwise.

### GetBaseChoicesOk

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetBaseChoicesOk() (*BulkCustomFieldChoiceSetRequestBaseChoices, bool)`

GetBaseChoicesOk returns a tuple with the BaseChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseChoices

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetBaseChoices(v BulkCustomFieldChoiceSetRequestBaseChoices)`

SetBaseChoices sets BaseChoices field to given value.

### HasBaseChoices

`func (o *PatchedBulkCustomFieldChoiceSetRequest) HasBaseChoices() bool`

HasBaseChoices returns a boolean if a field has been set.

### SetBaseChoicesNil

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetBaseChoicesNil(b bool)`

 SetBaseChoicesNil sets the value for BaseChoices to be an explicit nil

### UnsetBaseChoices
`func (o *PatchedBulkCustomFieldChoiceSetRequest) UnsetBaseChoices()`

UnsetBaseChoices ensures that no value is present for BaseChoices, not even an explicit nil
### GetExtraChoices

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetExtraChoices() [][]interface{}`

GetExtraChoices returns the ExtraChoices field if non-nil, zero value otherwise.

### GetExtraChoicesOk

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetExtraChoicesOk() (*[][]interface{}, bool)`

GetExtraChoicesOk returns a tuple with the ExtraChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraChoices

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetExtraChoices(v [][]interface{})`

SetExtraChoices sets ExtraChoices field to given value.

### HasExtraChoices

`func (o *PatchedBulkCustomFieldChoiceSetRequest) HasExtraChoices() bool`

HasExtraChoices returns a boolean if a field has been set.

### GetChoiceColors

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetChoiceColors() map[string]BulkCustomFieldChoiceSetRequestChoiceColorsValue`

GetChoiceColors returns the ChoiceColors field if non-nil, zero value otherwise.

### GetChoiceColorsOk

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetChoiceColorsOk() (*map[string]BulkCustomFieldChoiceSetRequestChoiceColorsValue, bool)`

GetChoiceColorsOk returns a tuple with the ChoiceColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceColors

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetChoiceColors(v map[string]BulkCustomFieldChoiceSetRequestChoiceColorsValue)`

SetChoiceColors sets ChoiceColors field to given value.

### HasChoiceColors

`func (o *PatchedBulkCustomFieldChoiceSetRequest) HasChoiceColors() bool`

HasChoiceColors returns a boolean if a field has been set.

### GetOrderAlphabetically

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetOrderAlphabetically() bool`

GetOrderAlphabetically returns the OrderAlphabetically field if non-nil, zero value otherwise.

### GetOrderAlphabeticallyOk

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetOrderAlphabeticallyOk() (*bool, bool)`

GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAlphabetically

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetOrderAlphabetically(v bool)`

SetOrderAlphabetically sets OrderAlphabetically field to given value.

### HasOrderAlphabetically

`func (o *PatchedBulkCustomFieldChoiceSetRequest) HasOrderAlphabetically() bool`

HasOrderAlphabetically returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkCustomFieldChoiceSetRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkCustomFieldChoiceSetRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkCustomFieldChoiceSetRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkCustomFieldChoiceSetRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


