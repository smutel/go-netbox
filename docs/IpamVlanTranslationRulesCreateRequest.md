# IpamVlanTranslationRulesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **int32** |  | 
**LocalVid** | **int32** | Numeric VLAN ID (1-4094) | 
**RemoteVid** | **int32** | Numeric VLAN ID (1-4094) | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewIpamVlanTranslationRulesCreateRequest

`func NewIpamVlanTranslationRulesCreateRequest(policy int32, localVid int32, remoteVid int32, ) *IpamVlanTranslationRulesCreateRequest`

NewIpamVlanTranslationRulesCreateRequest instantiates a new IpamVlanTranslationRulesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamVlanTranslationRulesCreateRequestWithDefaults

`func NewIpamVlanTranslationRulesCreateRequestWithDefaults() *IpamVlanTranslationRulesCreateRequest`

NewIpamVlanTranslationRulesCreateRequestWithDefaults instantiates a new IpamVlanTranslationRulesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *IpamVlanTranslationRulesCreateRequest) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IpamVlanTranslationRulesCreateRequest) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IpamVlanTranslationRulesCreateRequest) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetLocalVid

`func (o *IpamVlanTranslationRulesCreateRequest) GetLocalVid() int32`

GetLocalVid returns the LocalVid field if non-nil, zero value otherwise.

### GetLocalVidOk

`func (o *IpamVlanTranslationRulesCreateRequest) GetLocalVidOk() (*int32, bool)`

GetLocalVidOk returns a tuple with the LocalVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVid

`func (o *IpamVlanTranslationRulesCreateRequest) SetLocalVid(v int32)`

SetLocalVid sets LocalVid field to given value.


### GetRemoteVid

`func (o *IpamVlanTranslationRulesCreateRequest) GetRemoteVid() int32`

GetRemoteVid returns the RemoteVid field if non-nil, zero value otherwise.

### GetRemoteVidOk

`func (o *IpamVlanTranslationRulesCreateRequest) GetRemoteVidOk() (*int32, bool)`

GetRemoteVidOk returns a tuple with the RemoteVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVid

`func (o *IpamVlanTranslationRulesCreateRequest) SetRemoteVid(v int32)`

SetRemoteVid sets RemoteVid field to given value.


### GetDescription

`func (o *IpamVlanTranslationRulesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpamVlanTranslationRulesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpamVlanTranslationRulesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpamVlanTranslationRulesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *IpamVlanTranslationRulesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamVlanTranslationRulesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamVlanTranslationRulesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamVlanTranslationRulesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IpamVlanTranslationRulesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IpamVlanTranslationRulesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IpamVlanTranslationRulesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IpamVlanTranslationRulesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


