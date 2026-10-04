# BriefVLANTranslationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Policy** | **int32** |  | 
**LocalVid** | **int32** | Numeric VLAN ID (1-4094) | 
**RemoteVid** | **int32** | Numeric VLAN ID (1-4094) | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefVLANTranslationRule

`func NewBriefVLANTranslationRule(id int32, url string, display string, policy int32, localVid int32, remoteVid int32, ) *BriefVLANTranslationRule`

NewBriefVLANTranslationRule instantiates a new BriefVLANTranslationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefVLANTranslationRuleWithDefaults

`func NewBriefVLANTranslationRuleWithDefaults() *BriefVLANTranslationRule`

NewBriefVLANTranslationRuleWithDefaults instantiates a new BriefVLANTranslationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefVLANTranslationRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefVLANTranslationRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefVLANTranslationRule) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefVLANTranslationRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefVLANTranslationRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefVLANTranslationRule) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefVLANTranslationRule) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefVLANTranslationRule) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefVLANTranslationRule) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetPolicy

`func (o *BriefVLANTranslationRule) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BriefVLANTranslationRule) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BriefVLANTranslationRule) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetLocalVid

`func (o *BriefVLANTranslationRule) GetLocalVid() int32`

GetLocalVid returns the LocalVid field if non-nil, zero value otherwise.

### GetLocalVidOk

`func (o *BriefVLANTranslationRule) GetLocalVidOk() (*int32, bool)`

GetLocalVidOk returns a tuple with the LocalVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVid

`func (o *BriefVLANTranslationRule) SetLocalVid(v int32)`

SetLocalVid sets LocalVid field to given value.


### GetRemoteVid

`func (o *BriefVLANTranslationRule) GetRemoteVid() int32`

GetRemoteVid returns the RemoteVid field if non-nil, zero value otherwise.

### GetRemoteVidOk

`func (o *BriefVLANTranslationRule) GetRemoteVidOk() (*int32, bool)`

GetRemoteVidOk returns a tuple with the RemoteVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVid

`func (o *BriefVLANTranslationRule) SetRemoteVid(v int32)`

SetRemoteVid sets RemoteVid field to given value.


### GetDescription

`func (o *BriefVLANTranslationRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefVLANTranslationRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefVLANTranslationRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefVLANTranslationRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


