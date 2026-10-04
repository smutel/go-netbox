# PatchedBulkVLANTranslationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Policy** | Pointer to **int32** |  | [optional] 
**LocalVid** | Pointer to **int32** | Numeric VLAN ID (1-4094) | [optional] 
**RemoteVid** | Pointer to **int32** | Numeric VLAN ID (1-4094) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBulkVLANTranslationRuleRequest

`func NewPatchedBulkVLANTranslationRuleRequest(id int32, ) *PatchedBulkVLANTranslationRuleRequest`

NewPatchedBulkVLANTranslationRuleRequest instantiates a new PatchedBulkVLANTranslationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkVLANTranslationRuleRequestWithDefaults

`func NewPatchedBulkVLANTranslationRuleRequestWithDefaults() *PatchedBulkVLANTranslationRuleRequest`

NewPatchedBulkVLANTranslationRuleRequestWithDefaults instantiates a new PatchedBulkVLANTranslationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkVLANTranslationRuleRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkVLANTranslationRuleRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkVLANTranslationRuleRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetPolicy

`func (o *PatchedBulkVLANTranslationRuleRequest) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PatchedBulkVLANTranslationRuleRequest) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PatchedBulkVLANTranslationRuleRequest) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PatchedBulkVLANTranslationRuleRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetLocalVid

`func (o *PatchedBulkVLANTranslationRuleRequest) GetLocalVid() int32`

GetLocalVid returns the LocalVid field if non-nil, zero value otherwise.

### GetLocalVidOk

`func (o *PatchedBulkVLANTranslationRuleRequest) GetLocalVidOk() (*int32, bool)`

GetLocalVidOk returns a tuple with the LocalVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVid

`func (o *PatchedBulkVLANTranslationRuleRequest) SetLocalVid(v int32)`

SetLocalVid sets LocalVid field to given value.

### HasLocalVid

`func (o *PatchedBulkVLANTranslationRuleRequest) HasLocalVid() bool`

HasLocalVid returns a boolean if a field has been set.

### GetRemoteVid

`func (o *PatchedBulkVLANTranslationRuleRequest) GetRemoteVid() int32`

GetRemoteVid returns the RemoteVid field if non-nil, zero value otherwise.

### GetRemoteVidOk

`func (o *PatchedBulkVLANTranslationRuleRequest) GetRemoteVidOk() (*int32, bool)`

GetRemoteVidOk returns a tuple with the RemoteVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVid

`func (o *PatchedBulkVLANTranslationRuleRequest) SetRemoteVid(v int32)`

SetRemoteVid sets RemoteVid field to given value.

### HasRemoteVid

`func (o *PatchedBulkVLANTranslationRuleRequest) HasRemoteVid() bool`

HasRemoteVid returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkVLANTranslationRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkVLANTranslationRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkVLANTranslationRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkVLANTranslationRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkVLANTranslationRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkVLANTranslationRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkVLANTranslationRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkVLANTranslationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkVLANTranslationRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkVLANTranslationRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkVLANTranslationRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkVLANTranslationRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


