# PatchedBulkEventRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EventTypes** | Pointer to [**[]BulkEventRuleRequestEventTypesInner**](BulkEventRuleRequestEventTypesInner.md) | The types of event which will trigger this rule. | [optional] 
**Conditions** | Pointer to **interface{}** | A set of conditions which determine whether the event will be generated. | [optional] 
**ActionType** | Pointer to [**BulkEventRuleRequestActionType**](BulkEventRuleRequestActionType.md) |  | [optional] 
**ActionObjectType** | Pointer to **NullableString** |  | [optional] 
**ActionObjectId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewPatchedBulkEventRuleRequest

`func NewPatchedBulkEventRuleRequest(id int32, ) *PatchedBulkEventRuleRequest`

NewPatchedBulkEventRuleRequest instantiates a new PatchedBulkEventRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkEventRuleRequestWithDefaults

`func NewPatchedBulkEventRuleRequestWithDefaults() *PatchedBulkEventRuleRequest`

NewPatchedBulkEventRuleRequestWithDefaults instantiates a new PatchedBulkEventRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkEventRuleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkEventRuleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkEventRuleRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *PatchedBulkEventRuleRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedBulkEventRuleRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedBulkEventRuleRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedBulkEventRuleRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkEventRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkEventRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkEventRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkEventRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkEventRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkEventRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkEventRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkEventRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypes

`func (o *PatchedBulkEventRuleRequest) GetEventTypes() []BulkEventRuleRequestEventTypesInner`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *PatchedBulkEventRuleRequest) GetEventTypesOk() (*[]BulkEventRuleRequestEventTypesInner, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *PatchedBulkEventRuleRequest) SetEventTypes(v []BulkEventRuleRequestEventTypesInner)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *PatchedBulkEventRuleRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetConditions

`func (o *PatchedBulkEventRuleRequest) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PatchedBulkEventRuleRequest) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PatchedBulkEventRuleRequest) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PatchedBulkEventRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *PatchedBulkEventRuleRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PatchedBulkEventRuleRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetActionType

`func (o *PatchedBulkEventRuleRequest) GetActionType() BulkEventRuleRequestActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PatchedBulkEventRuleRequest) GetActionTypeOk() (*BulkEventRuleRequestActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PatchedBulkEventRuleRequest) SetActionType(v BulkEventRuleRequestActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *PatchedBulkEventRuleRequest) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionObjectType

`func (o *PatchedBulkEventRuleRequest) GetActionObjectType() string`

GetActionObjectType returns the ActionObjectType field if non-nil, zero value otherwise.

### GetActionObjectTypeOk

`func (o *PatchedBulkEventRuleRequest) GetActionObjectTypeOk() (*string, bool)`

GetActionObjectTypeOk returns a tuple with the ActionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectType

`func (o *PatchedBulkEventRuleRequest) SetActionObjectType(v string)`

SetActionObjectType sets ActionObjectType field to given value.

### HasActionObjectType

`func (o *PatchedBulkEventRuleRequest) HasActionObjectType() bool`

HasActionObjectType returns a boolean if a field has been set.

### SetActionObjectTypeNil

`func (o *PatchedBulkEventRuleRequest) SetActionObjectTypeNil(b bool)`

 SetActionObjectTypeNil sets the value for ActionObjectType to be an explicit nil

### UnsetActionObjectType
`func (o *PatchedBulkEventRuleRequest) UnsetActionObjectType()`

UnsetActionObjectType ensures that no value is present for ActionObjectType, not even an explicit nil
### GetActionObjectId

`func (o *PatchedBulkEventRuleRequest) GetActionObjectId() int64`

GetActionObjectId returns the ActionObjectId field if non-nil, zero value otherwise.

### GetActionObjectIdOk

`func (o *PatchedBulkEventRuleRequest) GetActionObjectIdOk() (*int64, bool)`

GetActionObjectIdOk returns a tuple with the ActionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectId

`func (o *PatchedBulkEventRuleRequest) SetActionObjectId(v int64)`

SetActionObjectId sets ActionObjectId field to given value.

### HasActionObjectId

`func (o *PatchedBulkEventRuleRequest) HasActionObjectId() bool`

HasActionObjectId returns a boolean if a field has been set.

### SetActionObjectIdNil

`func (o *PatchedBulkEventRuleRequest) SetActionObjectIdNil(b bool)`

 SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil

### UnsetActionObjectId
`func (o *PatchedBulkEventRuleRequest) UnsetActionObjectId()`

UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
### GetDescription

`func (o *PatchedBulkEventRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkEventRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkEventRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkEventRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkEventRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkEventRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkEventRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkEventRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkEventRuleRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkEventRuleRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkEventRuleRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkEventRuleRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkEventRuleRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkEventRuleRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *PatchedBulkEventRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkEventRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkEventRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkEventRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


