# BriefASN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Asn** | **int64** | 16- or 32-bit autonomous system number | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefASN

`func NewBriefASN(id int32, url string, display string, asn int64, ) *BriefASN`

NewBriefASN instantiates a new BriefASN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefASNWithDefaults

`func NewBriefASNWithDefaults() *BriefASN`

NewBriefASNWithDefaults instantiates a new BriefASN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefASN) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefASN) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefASN) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefASN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefASN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefASN) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefASN) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefASN) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefASN) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetAsn

`func (o *BriefASN) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *BriefASN) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *BriefASN) SetAsn(v int64)`

SetAsn sets Asn field to given value.


### GetDescription

`func (o *BriefASN) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefASN) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefASN) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefASN) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


