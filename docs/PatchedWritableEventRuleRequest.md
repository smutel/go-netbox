# PatchedWritableEventRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TypeCreate** | Pointer to **bool** | Triggers when a matching object is created. | [optional] 
**TypeUpdate** | Pointer to **bool** | Triggers when a matching object is updated. | [optional] 
**TypeDelete** | Pointer to **bool** | Triggers when a matching object is deleted. | [optional] 
**TypeJobStart** | Pointer to **bool** | Triggers when a job for a matching object is started. | [optional] 
**TypeJobEnd** | Pointer to **bool** | Triggers when a job for a matching object terminates. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Conditions** | Pointer to **interface{}** | A set of conditions which determine whether the event will be generated. | [optional] 
**ActionType** | Pointer to [**EventRuleActionTypeValue**](EventRuleActionTypeValue.md) |  | [optional] 
**ActionObjectType** | Pointer to **string** |  | [optional] 
**ActionObjectId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewPatchedWritableEventRuleRequest

`func NewPatchedWritableEventRuleRequest() *PatchedWritableEventRuleRequest`

NewPatchedWritableEventRuleRequest instantiates a new PatchedWritableEventRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableEventRuleRequestWithDefaults

`func NewPatchedWritableEventRuleRequestWithDefaults() *PatchedWritableEventRuleRequest`

NewPatchedWritableEventRuleRequestWithDefaults instantiates a new PatchedWritableEventRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *PatchedWritableEventRuleRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedWritableEventRuleRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedWritableEventRuleRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedWritableEventRuleRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableEventRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableEventRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableEventRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableEventRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeCreate

`func (o *PatchedWritableEventRuleRequest) GetTypeCreate() bool`

GetTypeCreate returns the TypeCreate field if non-nil, zero value otherwise.

### GetTypeCreateOk

`func (o *PatchedWritableEventRuleRequest) GetTypeCreateOk() (*bool, bool)`

GetTypeCreateOk returns a tuple with the TypeCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCreate

`func (o *PatchedWritableEventRuleRequest) SetTypeCreate(v bool)`

SetTypeCreate sets TypeCreate field to given value.

### HasTypeCreate

`func (o *PatchedWritableEventRuleRequest) HasTypeCreate() bool`

HasTypeCreate returns a boolean if a field has been set.

### GetTypeUpdate

`func (o *PatchedWritableEventRuleRequest) GetTypeUpdate() bool`

GetTypeUpdate returns the TypeUpdate field if non-nil, zero value otherwise.

### GetTypeUpdateOk

`func (o *PatchedWritableEventRuleRequest) GetTypeUpdateOk() (*bool, bool)`

GetTypeUpdateOk returns a tuple with the TypeUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUpdate

`func (o *PatchedWritableEventRuleRequest) SetTypeUpdate(v bool)`

SetTypeUpdate sets TypeUpdate field to given value.

### HasTypeUpdate

`func (o *PatchedWritableEventRuleRequest) HasTypeUpdate() bool`

HasTypeUpdate returns a boolean if a field has been set.

### GetTypeDelete

`func (o *PatchedWritableEventRuleRequest) GetTypeDelete() bool`

GetTypeDelete returns the TypeDelete field if non-nil, zero value otherwise.

### GetTypeDeleteOk

`func (o *PatchedWritableEventRuleRequest) GetTypeDeleteOk() (*bool, bool)`

GetTypeDeleteOk returns a tuple with the TypeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDelete

`func (o *PatchedWritableEventRuleRequest) SetTypeDelete(v bool)`

SetTypeDelete sets TypeDelete field to given value.

### HasTypeDelete

`func (o *PatchedWritableEventRuleRequest) HasTypeDelete() bool`

HasTypeDelete returns a boolean if a field has been set.

### GetTypeJobStart

`func (o *PatchedWritableEventRuleRequest) GetTypeJobStart() bool`

GetTypeJobStart returns the TypeJobStart field if non-nil, zero value otherwise.

### GetTypeJobStartOk

`func (o *PatchedWritableEventRuleRequest) GetTypeJobStartOk() (*bool, bool)`

GetTypeJobStartOk returns a tuple with the TypeJobStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeJobStart

`func (o *PatchedWritableEventRuleRequest) SetTypeJobStart(v bool)`

SetTypeJobStart sets TypeJobStart field to given value.

### HasTypeJobStart

`func (o *PatchedWritableEventRuleRequest) HasTypeJobStart() bool`

HasTypeJobStart returns a boolean if a field has been set.

### GetTypeJobEnd

`func (o *PatchedWritableEventRuleRequest) GetTypeJobEnd() bool`

GetTypeJobEnd returns the TypeJobEnd field if non-nil, zero value otherwise.

### GetTypeJobEndOk

`func (o *PatchedWritableEventRuleRequest) GetTypeJobEndOk() (*bool, bool)`

GetTypeJobEndOk returns a tuple with the TypeJobEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeJobEnd

`func (o *PatchedWritableEventRuleRequest) SetTypeJobEnd(v bool)`

SetTypeJobEnd sets TypeJobEnd field to given value.

### HasTypeJobEnd

`func (o *PatchedWritableEventRuleRequest) HasTypeJobEnd() bool`

HasTypeJobEnd returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedWritableEventRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedWritableEventRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedWritableEventRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedWritableEventRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConditions

`func (o *PatchedWritableEventRuleRequest) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PatchedWritableEventRuleRequest) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PatchedWritableEventRuleRequest) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PatchedWritableEventRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *PatchedWritableEventRuleRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PatchedWritableEventRuleRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetActionType

`func (o *PatchedWritableEventRuleRequest) GetActionType() EventRuleActionTypeValue`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PatchedWritableEventRuleRequest) GetActionTypeOk() (*EventRuleActionTypeValue, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PatchedWritableEventRuleRequest) SetActionType(v EventRuleActionTypeValue)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *PatchedWritableEventRuleRequest) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionObjectType

`func (o *PatchedWritableEventRuleRequest) GetActionObjectType() string`

GetActionObjectType returns the ActionObjectType field if non-nil, zero value otherwise.

### GetActionObjectTypeOk

`func (o *PatchedWritableEventRuleRequest) GetActionObjectTypeOk() (*string, bool)`

GetActionObjectTypeOk returns a tuple with the ActionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectType

`func (o *PatchedWritableEventRuleRequest) SetActionObjectType(v string)`

SetActionObjectType sets ActionObjectType field to given value.

### HasActionObjectType

`func (o *PatchedWritableEventRuleRequest) HasActionObjectType() bool`

HasActionObjectType returns a boolean if a field has been set.

### GetActionObjectId

`func (o *PatchedWritableEventRuleRequest) GetActionObjectId() int64`

GetActionObjectId returns the ActionObjectId field if non-nil, zero value otherwise.

### GetActionObjectIdOk

`func (o *PatchedWritableEventRuleRequest) GetActionObjectIdOk() (*int64, bool)`

GetActionObjectIdOk returns a tuple with the ActionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectId

`func (o *PatchedWritableEventRuleRequest) SetActionObjectId(v int64)`

SetActionObjectId sets ActionObjectId field to given value.

### HasActionObjectId

`func (o *PatchedWritableEventRuleRequest) HasActionObjectId() bool`

HasActionObjectId returns a boolean if a field has been set.

### SetActionObjectIdNil

`func (o *PatchedWritableEventRuleRequest) SetActionObjectIdNil(b bool)`

 SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil

### UnsetActionObjectId
`func (o *PatchedWritableEventRuleRequest) UnsetActionObjectId()`

UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
### GetDescription

`func (o *PatchedWritableEventRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableEventRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableEventRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableEventRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableEventRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableEventRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableEventRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableEventRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableEventRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableEventRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableEventRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableEventRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


