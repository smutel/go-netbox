# BulkVLANTranslationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Policy** | **int32** |  | 
**LocalVid** | **int32** | Numeric VLAN ID (1-4094) | 
**RemoteVid** | **int32** | Numeric VLAN ID (1-4094) | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkVLANTranslationRuleRequest

`func NewBulkVLANTranslationRuleRequest(id int32, policy int32, localVid int32, remoteVid int32, ) *BulkVLANTranslationRuleRequest`

NewBulkVLANTranslationRuleRequest instantiates a new BulkVLANTranslationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVLANTranslationRuleRequestWithDefaults

`func NewBulkVLANTranslationRuleRequestWithDefaults() *BulkVLANTranslationRuleRequest`

NewBulkVLANTranslationRuleRequestWithDefaults instantiates a new BulkVLANTranslationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkVLANTranslationRuleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkVLANTranslationRuleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkVLANTranslationRuleRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetPolicy

`func (o *BulkVLANTranslationRuleRequest) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BulkVLANTranslationRuleRequest) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BulkVLANTranslationRuleRequest) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetLocalVid

`func (o *BulkVLANTranslationRuleRequest) GetLocalVid() int32`

GetLocalVid returns the LocalVid field if non-nil, zero value otherwise.

### GetLocalVidOk

`func (o *BulkVLANTranslationRuleRequest) GetLocalVidOk() (*int32, bool)`

GetLocalVidOk returns a tuple with the LocalVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVid

`func (o *BulkVLANTranslationRuleRequest) SetLocalVid(v int32)`

SetLocalVid sets LocalVid field to given value.


### GetRemoteVid

`func (o *BulkVLANTranslationRuleRequest) GetRemoteVid() int32`

GetRemoteVid returns the RemoteVid field if non-nil, zero value otherwise.

### GetRemoteVidOk

`func (o *BulkVLANTranslationRuleRequest) GetRemoteVidOk() (*int32, bool)`

GetRemoteVidOk returns a tuple with the RemoteVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVid

`func (o *BulkVLANTranslationRuleRequest) SetRemoteVid(v int32)`

SetRemoteVid sets RemoteVid field to given value.


### GetDescription

`func (o *BulkVLANTranslationRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkVLANTranslationRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkVLANTranslationRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkVLANTranslationRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *BulkVLANTranslationRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkVLANTranslationRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkVLANTranslationRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkVLANTranslationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkVLANTranslationRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkVLANTranslationRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkVLANTranslationRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkVLANTranslationRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


