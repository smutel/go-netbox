# BulkEventRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**EventTypes** | [**[]BulkEventRuleRequestEventTypesInner**](BulkEventRuleRequestEventTypesInner.md) | The types of event which will trigger this rule. | 
**Conditions** | Pointer to **interface{}** | A set of conditions which determine whether the event will be generated. | [optional] 
**ActionType** | [**BulkEventRuleRequestActionType**](BulkEventRuleRequestActionType.md) |  | 
**ActionObjectType** | Pointer to **NullableString** |  | [optional] 
**ActionObjectId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewBulkEventRuleRequest

`func NewBulkEventRuleRequest(id int32, objectTypes []string, name string, eventTypes []BulkEventRuleRequestEventTypesInner, actionType BulkEventRuleRequestActionType, ) *BulkEventRuleRequest`

NewBulkEventRuleRequest instantiates a new BulkEventRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEventRuleRequestWithDefaults

`func NewBulkEventRuleRequestWithDefaults() *BulkEventRuleRequest`

NewBulkEventRuleRequestWithDefaults instantiates a new BulkEventRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkEventRuleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkEventRuleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkEventRuleRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *BulkEventRuleRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *BulkEventRuleRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *BulkEventRuleRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetName

`func (o *BulkEventRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkEventRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkEventRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *BulkEventRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkEventRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkEventRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkEventRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypes

`func (o *BulkEventRuleRequest) GetEventTypes() []BulkEventRuleRequestEventTypesInner`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *BulkEventRuleRequest) GetEventTypesOk() (*[]BulkEventRuleRequestEventTypesInner, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *BulkEventRuleRequest) SetEventTypes(v []BulkEventRuleRequestEventTypesInner)`

SetEventTypes sets EventTypes field to given value.


### GetConditions

`func (o *BulkEventRuleRequest) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *BulkEventRuleRequest) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *BulkEventRuleRequest) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *BulkEventRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *BulkEventRuleRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *BulkEventRuleRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetActionType

`func (o *BulkEventRuleRequest) GetActionType() BulkEventRuleRequestActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *BulkEventRuleRequest) GetActionTypeOk() (*BulkEventRuleRequestActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *BulkEventRuleRequest) SetActionType(v BulkEventRuleRequestActionType)`

SetActionType sets ActionType field to given value.


### GetActionObjectType

`func (o *BulkEventRuleRequest) GetActionObjectType() string`

GetActionObjectType returns the ActionObjectType field if non-nil, zero value otherwise.

### GetActionObjectTypeOk

`func (o *BulkEventRuleRequest) GetActionObjectTypeOk() (*string, bool)`

GetActionObjectTypeOk returns a tuple with the ActionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectType

`func (o *BulkEventRuleRequest) SetActionObjectType(v string)`

SetActionObjectType sets ActionObjectType field to given value.

### HasActionObjectType

`func (o *BulkEventRuleRequest) HasActionObjectType() bool`

HasActionObjectType returns a boolean if a field has been set.

### SetActionObjectTypeNil

`func (o *BulkEventRuleRequest) SetActionObjectTypeNil(b bool)`

 SetActionObjectTypeNil sets the value for ActionObjectType to be an explicit nil

### UnsetActionObjectType
`func (o *BulkEventRuleRequest) UnsetActionObjectType()`

UnsetActionObjectType ensures that no value is present for ActionObjectType, not even an explicit nil
### GetActionObjectId

`func (o *BulkEventRuleRequest) GetActionObjectId() int64`

GetActionObjectId returns the ActionObjectId field if non-nil, zero value otherwise.

### GetActionObjectIdOk

`func (o *BulkEventRuleRequest) GetActionObjectIdOk() (*int64, bool)`

GetActionObjectIdOk returns a tuple with the ActionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectId

`func (o *BulkEventRuleRequest) SetActionObjectId(v int64)`

SetActionObjectId sets ActionObjectId field to given value.

### HasActionObjectId

`func (o *BulkEventRuleRequest) HasActionObjectId() bool`

HasActionObjectId returns a boolean if a field has been set.

### SetActionObjectIdNil

`func (o *BulkEventRuleRequest) SetActionObjectIdNil(b bool)`

 SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil

### UnsetActionObjectId
`func (o *BulkEventRuleRequest) UnsetActionObjectId()`

UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
### GetDescription

`func (o *BulkEventRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkEventRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkEventRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkEventRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkEventRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkEventRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkEventRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkEventRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetOwner

`func (o *BulkEventRuleRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkEventRuleRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkEventRuleRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkEventRuleRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkEventRuleRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkEventRuleRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *BulkEventRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkEventRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkEventRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkEventRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


