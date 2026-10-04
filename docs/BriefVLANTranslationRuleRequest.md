# BriefVLANTranslationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **int32** |  | 
**LocalVid** | **int32** | Numeric VLAN ID (1-4094) | 
**RemoteVid** | **int32** | Numeric VLAN ID (1-4094) | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefVLANTranslationRuleRequest

`func NewBriefVLANTranslationRuleRequest(policy int32, localVid int32, remoteVid int32, ) *BriefVLANTranslationRuleRequest`

NewBriefVLANTranslationRuleRequest instantiates a new BriefVLANTranslationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefVLANTranslationRuleRequestWithDefaults

`func NewBriefVLANTranslationRuleRequestWithDefaults() *BriefVLANTranslationRuleRequest`

NewBriefVLANTranslationRuleRequestWithDefaults instantiates a new BriefVLANTranslationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *BriefVLANTranslationRuleRequest) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BriefVLANTranslationRuleRequest) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BriefVLANTranslationRuleRequest) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetLocalVid

`func (o *BriefVLANTranslationRuleRequest) GetLocalVid() int32`

GetLocalVid returns the LocalVid field if non-nil, zero value otherwise.

### GetLocalVidOk

`func (o *BriefVLANTranslationRuleRequest) GetLocalVidOk() (*int32, bool)`

GetLocalVidOk returns a tuple with the LocalVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVid

`func (o *BriefVLANTranslationRuleRequest) SetLocalVid(v int32)`

SetLocalVid sets LocalVid field to given value.


### GetRemoteVid

`func (o *BriefVLANTranslationRuleRequest) GetRemoteVid() int32`

GetRemoteVid returns the RemoteVid field if non-nil, zero value otherwise.

### GetRemoteVidOk

`func (o *BriefVLANTranslationRuleRequest) GetRemoteVidOk() (*int32, bool)`

GetRemoteVidOk returns a tuple with the RemoteVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVid

`func (o *BriefVLANTranslationRuleRequest) SetRemoteVid(v int32)`

SetRemoteVid sets RemoteVid field to given value.


### GetDescription

`func (o *BriefVLANTranslationRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefVLANTranslationRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefVLANTranslationRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefVLANTranslationRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


