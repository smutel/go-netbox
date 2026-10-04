# CustomFieldStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**CustomFieldStatusValue**](CustomFieldStatusValue.md) |  | [optional] 
**Label** | Pointer to [**CustomFieldStatusLabel**](CustomFieldStatusLabel.md) |  | [optional] 

## Methods

### NewCustomFieldStatus

`func NewCustomFieldStatus() *CustomFieldStatus`

NewCustomFieldStatus instantiates a new CustomFieldStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldStatusWithDefaults

`func NewCustomFieldStatusWithDefaults() *CustomFieldStatus`

NewCustomFieldStatusWithDefaults instantiates a new CustomFieldStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CustomFieldStatus) GetValue() CustomFieldStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldStatus) GetValueOk() (*CustomFieldStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldStatus) SetValue(v CustomFieldStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CustomFieldStatus) GetLabel() CustomFieldStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomFieldStatus) GetLabelOk() (*CustomFieldStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomFieldStatus) SetLabel(v CustomFieldStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomFieldStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


