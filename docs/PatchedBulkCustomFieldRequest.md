# PatchedBulkCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**BulkCustomFieldRequestType**](BulkCustomFieldRequestType.md) |  | [optional] 
**RelatedObjectType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** | Internal field name | [optional] 
**Label** | Pointer to **string** | Name of the field as displayed to users (if not provided, &#39;the field&#39;s name will be used) | [optional] 
**GroupName** | Pointer to **string** | Custom fields within the same group will be displayed together | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** | This field is required when creating new objects or editing an existing object. | [optional] 
**Unique** | Pointer to **bool** | The value of this field must be unique for the assigned object | [optional] 
**SearchWeight** | Pointer to **int32** | Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored. | [optional] 
**FilterLogic** | Pointer to [**BulkCustomFieldRequestFilterLogic**](BulkCustomFieldRequestFilterLogic.md) |  | [optional] 
**UiVisible** | Pointer to [**BulkCustomFieldRequestUiVisible**](BulkCustomFieldRequestUiVisible.md) |  | [optional] 
**UiEditable** | Pointer to [**BulkCustomFieldRequestUiEditable**](BulkCustomFieldRequestUiEditable.md) |  | [optional] 
**IsCloneable** | Pointer to **bool** | Replicate this value when cloning objects | [optional] 
**NullsFirst** | Pointer to **bool** | Sort null values before non-null values when ordering by this field | [optional] 
**Default** | Pointer to **interface{}** | Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] 
**RelatedObjectFilter** | Pointer to **interface{}** | Filter the object selection choices using a query_params dict (must be a JSON value).Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] 
**Weight** | Pointer to **int32** | Fields with higher weights appear lower in a form. | [optional] 
**ValidationMinimum** | Pointer to **NullableFloat64** | Minimum allowed value (for numeric fields) | [optional] 
**ValidationMaximum** | Pointer to **NullableFloat64** | Maximum allowed value (for numeric fields) | [optional] 
**ValidationRegex** | Pointer to **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. | [optional] 
**ValidationSchema** | Pointer to **interface{}** | A JSON schema definition for validating the custom field value | [optional] 
**ChoiceSet** | Pointer to [**NullableBulkCustomFieldRequestChoiceSet**](BulkCustomFieldRequestChoiceSet.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedBulkCustomFieldRequest

`func NewPatchedBulkCustomFieldRequest(id int32, ) *PatchedBulkCustomFieldRequest`

NewPatchedBulkCustomFieldRequest instantiates a new PatchedBulkCustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCustomFieldRequestWithDefaults

`func NewPatchedBulkCustomFieldRequestWithDefaults() *PatchedBulkCustomFieldRequest`

NewPatchedBulkCustomFieldRequestWithDefaults instantiates a new PatchedBulkCustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCustomFieldRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCustomFieldRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCustomFieldRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *PatchedBulkCustomFieldRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedBulkCustomFieldRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedBulkCustomFieldRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedBulkCustomFieldRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkCustomFieldRequest) GetType() BulkCustomFieldRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkCustomFieldRequest) GetTypeOk() (*BulkCustomFieldRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkCustomFieldRequest) SetType(v BulkCustomFieldRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkCustomFieldRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelatedObjectType

`func (o *PatchedBulkCustomFieldRequest) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *PatchedBulkCustomFieldRequest) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *PatchedBulkCustomFieldRequest) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.

### HasRelatedObjectType

`func (o *PatchedBulkCustomFieldRequest) HasRelatedObjectType() bool`

HasRelatedObjectType returns a boolean if a field has been set.

### SetRelatedObjectTypeNil

`func (o *PatchedBulkCustomFieldRequest) SetRelatedObjectTypeNil(b bool)`

 SetRelatedObjectTypeNil sets the value for RelatedObjectType to be an explicit nil

### UnsetRelatedObjectType
`func (o *PatchedBulkCustomFieldRequest) UnsetRelatedObjectType()`

UnsetRelatedObjectType ensures that no value is present for RelatedObjectType, not even an explicit nil
### GetName

`func (o *PatchedBulkCustomFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkCustomFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkCustomFieldRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkCustomFieldRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkCustomFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkCustomFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkCustomFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkCustomFieldRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupName

`func (o *PatchedBulkCustomFieldRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PatchedBulkCustomFieldRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PatchedBulkCustomFieldRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PatchedBulkCustomFieldRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkCustomFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkCustomFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkCustomFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkCustomFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *PatchedBulkCustomFieldRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PatchedBulkCustomFieldRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PatchedBulkCustomFieldRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PatchedBulkCustomFieldRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUnique

`func (o *PatchedBulkCustomFieldRequest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *PatchedBulkCustomFieldRequest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *PatchedBulkCustomFieldRequest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *PatchedBulkCustomFieldRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetSearchWeight

`func (o *PatchedBulkCustomFieldRequest) GetSearchWeight() int32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *PatchedBulkCustomFieldRequest) GetSearchWeightOk() (*int32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *PatchedBulkCustomFieldRequest) SetSearchWeight(v int32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *PatchedBulkCustomFieldRequest) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.

### GetFilterLogic

`func (o *PatchedBulkCustomFieldRequest) GetFilterLogic() BulkCustomFieldRequestFilterLogic`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *PatchedBulkCustomFieldRequest) GetFilterLogicOk() (*BulkCustomFieldRequestFilterLogic, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *PatchedBulkCustomFieldRequest) SetFilterLogic(v BulkCustomFieldRequestFilterLogic)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *PatchedBulkCustomFieldRequest) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetUiVisible

`func (o *PatchedBulkCustomFieldRequest) GetUiVisible() BulkCustomFieldRequestUiVisible`

GetUiVisible returns the UiVisible field if non-nil, zero value otherwise.

### GetUiVisibleOk

`func (o *PatchedBulkCustomFieldRequest) GetUiVisibleOk() (*BulkCustomFieldRequestUiVisible, bool)`

GetUiVisibleOk returns a tuple with the UiVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiVisible

`func (o *PatchedBulkCustomFieldRequest) SetUiVisible(v BulkCustomFieldRequestUiVisible)`

SetUiVisible sets UiVisible field to given value.

### HasUiVisible

`func (o *PatchedBulkCustomFieldRequest) HasUiVisible() bool`

HasUiVisible returns a boolean if a field has been set.

### GetUiEditable

`func (o *PatchedBulkCustomFieldRequest) GetUiEditable() BulkCustomFieldRequestUiEditable`

GetUiEditable returns the UiEditable field if non-nil, zero value otherwise.

### GetUiEditableOk

`func (o *PatchedBulkCustomFieldRequest) GetUiEditableOk() (*BulkCustomFieldRequestUiEditable, bool)`

GetUiEditableOk returns a tuple with the UiEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiEditable

`func (o *PatchedBulkCustomFieldRequest) SetUiEditable(v BulkCustomFieldRequestUiEditable)`

SetUiEditable sets UiEditable field to given value.

### HasUiEditable

`func (o *PatchedBulkCustomFieldRequest) HasUiEditable() bool`

HasUiEditable returns a boolean if a field has been set.

### GetIsCloneable

`func (o *PatchedBulkCustomFieldRequest) GetIsCloneable() bool`

GetIsCloneable returns the IsCloneable field if non-nil, zero value otherwise.

### GetIsCloneableOk

`func (o *PatchedBulkCustomFieldRequest) GetIsCloneableOk() (*bool, bool)`

GetIsCloneableOk returns a tuple with the IsCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloneable

`func (o *PatchedBulkCustomFieldRequest) SetIsCloneable(v bool)`

SetIsCloneable sets IsCloneable field to given value.

### HasIsCloneable

`func (o *PatchedBulkCustomFieldRequest) HasIsCloneable() bool`

HasIsCloneable returns a boolean if a field has been set.

### GetNullsFirst

`func (o *PatchedBulkCustomFieldRequest) GetNullsFirst() bool`

GetNullsFirst returns the NullsFirst field if non-nil, zero value otherwise.

### GetNullsFirstOk

`func (o *PatchedBulkCustomFieldRequest) GetNullsFirstOk() (*bool, bool)`

GetNullsFirstOk returns a tuple with the NullsFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullsFirst

`func (o *PatchedBulkCustomFieldRequest) SetNullsFirst(v bool)`

SetNullsFirst sets NullsFirst field to given value.

### HasNullsFirst

`func (o *PatchedBulkCustomFieldRequest) HasNullsFirst() bool`

HasNullsFirst returns a boolean if a field has been set.

### GetDefault

`func (o *PatchedBulkCustomFieldRequest) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PatchedBulkCustomFieldRequest) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PatchedBulkCustomFieldRequest) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PatchedBulkCustomFieldRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *PatchedBulkCustomFieldRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *PatchedBulkCustomFieldRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetRelatedObjectFilter

`func (o *PatchedBulkCustomFieldRequest) GetRelatedObjectFilter() interface{}`

GetRelatedObjectFilter returns the RelatedObjectFilter field if non-nil, zero value otherwise.

### GetRelatedObjectFilterOk

`func (o *PatchedBulkCustomFieldRequest) GetRelatedObjectFilterOk() (*interface{}, bool)`

GetRelatedObjectFilterOk returns a tuple with the RelatedObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectFilter

`func (o *PatchedBulkCustomFieldRequest) SetRelatedObjectFilter(v interface{})`

SetRelatedObjectFilter sets RelatedObjectFilter field to given value.

### HasRelatedObjectFilter

`func (o *PatchedBulkCustomFieldRequest) HasRelatedObjectFilter() bool`

HasRelatedObjectFilter returns a boolean if a field has been set.

### SetRelatedObjectFilterNil

`func (o *PatchedBulkCustomFieldRequest) SetRelatedObjectFilterNil(b bool)`

 SetRelatedObjectFilterNil sets the value for RelatedObjectFilter to be an explicit nil

### UnsetRelatedObjectFilter
`func (o *PatchedBulkCustomFieldRequest) UnsetRelatedObjectFilter()`

UnsetRelatedObjectFilter ensures that no value is present for RelatedObjectFilter, not even an explicit nil
### GetWeight

`func (o *PatchedBulkCustomFieldRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkCustomFieldRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkCustomFieldRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkCustomFieldRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetValidationMinimum

`func (o *PatchedBulkCustomFieldRequest) GetValidationMinimum() float64`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *PatchedBulkCustomFieldRequest) GetValidationMinimumOk() (*float64, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *PatchedBulkCustomFieldRequest) SetValidationMinimum(v float64)`

SetValidationMinimum sets ValidationMinimum field to given value.

### HasValidationMinimum

`func (o *PatchedBulkCustomFieldRequest) HasValidationMinimum() bool`

HasValidationMinimum returns a boolean if a field has been set.

### SetValidationMinimumNil

`func (o *PatchedBulkCustomFieldRequest) SetValidationMinimumNil(b bool)`

 SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil

### UnsetValidationMinimum
`func (o *PatchedBulkCustomFieldRequest) UnsetValidationMinimum()`

UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
### GetValidationMaximum

`func (o *PatchedBulkCustomFieldRequest) GetValidationMaximum() float64`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *PatchedBulkCustomFieldRequest) GetValidationMaximumOk() (*float64, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *PatchedBulkCustomFieldRequest) SetValidationMaximum(v float64)`

SetValidationMaximum sets ValidationMaximum field to given value.

### HasValidationMaximum

`func (o *PatchedBulkCustomFieldRequest) HasValidationMaximum() bool`

HasValidationMaximum returns a boolean if a field has been set.

### SetValidationMaximumNil

`func (o *PatchedBulkCustomFieldRequest) SetValidationMaximumNil(b bool)`

 SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil

### UnsetValidationMaximum
`func (o *PatchedBulkCustomFieldRequest) UnsetValidationMaximum()`

UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
### GetValidationRegex

`func (o *PatchedBulkCustomFieldRequest) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *PatchedBulkCustomFieldRequest) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *PatchedBulkCustomFieldRequest) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *PatchedBulkCustomFieldRequest) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetValidationSchema

`func (o *PatchedBulkCustomFieldRequest) GetValidationSchema() interface{}`

GetValidationSchema returns the ValidationSchema field if non-nil, zero value otherwise.

### GetValidationSchemaOk

`func (o *PatchedBulkCustomFieldRequest) GetValidationSchemaOk() (*interface{}, bool)`

GetValidationSchemaOk returns a tuple with the ValidationSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationSchema

`func (o *PatchedBulkCustomFieldRequest) SetValidationSchema(v interface{})`

SetValidationSchema sets ValidationSchema field to given value.

### HasValidationSchema

`func (o *PatchedBulkCustomFieldRequest) HasValidationSchema() bool`

HasValidationSchema returns a boolean if a field has been set.

### SetValidationSchemaNil

`func (o *PatchedBulkCustomFieldRequest) SetValidationSchemaNil(b bool)`

 SetValidationSchemaNil sets the value for ValidationSchema to be an explicit nil

### UnsetValidationSchema
`func (o *PatchedBulkCustomFieldRequest) UnsetValidationSchema()`

UnsetValidationSchema ensures that no value is present for ValidationSchema, not even an explicit nil
### GetChoiceSet

`func (o *PatchedBulkCustomFieldRequest) GetChoiceSet() BulkCustomFieldRequestChoiceSet`

GetChoiceSet returns the ChoiceSet field if non-nil, zero value otherwise.

### GetChoiceSetOk

`func (o *PatchedBulkCustomFieldRequest) GetChoiceSetOk() (*BulkCustomFieldRequestChoiceSet, bool)`

GetChoiceSetOk returns a tuple with the ChoiceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceSet

`func (o *PatchedBulkCustomFieldRequest) SetChoiceSet(v BulkCustomFieldRequestChoiceSet)`

SetChoiceSet sets ChoiceSet field to given value.

### HasChoiceSet

`func (o *PatchedBulkCustomFieldRequest) HasChoiceSet() bool`

HasChoiceSet returns a boolean if a field has been set.

### SetChoiceSetNil

`func (o *PatchedBulkCustomFieldRequest) SetChoiceSetNil(b bool)`

 SetChoiceSetNil sets the value for ChoiceSet to be an explicit nil

### UnsetChoiceSet
`func (o *PatchedBulkCustomFieldRequest) UnsetChoiceSet()`

UnsetChoiceSet ensures that no value is present for ChoiceSet, not even an explicit nil
### GetOwner

`func (o *PatchedBulkCustomFieldRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkCustomFieldRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkCustomFieldRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkCustomFieldRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkCustomFieldRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkCustomFieldRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkCustomFieldRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkCustomFieldRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkCustomFieldRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkCustomFieldRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


