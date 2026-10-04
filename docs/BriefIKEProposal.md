# BriefIKEProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefIKEProposal

`func NewBriefIKEProposal(id int32, url string, display string, name string, ) *BriefIKEProposal`

NewBriefIKEProposal instantiates a new BriefIKEProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefIKEProposalWithDefaults

`func NewBriefIKEProposalWithDefaults() *BriefIKEProposal`

NewBriefIKEProposalWithDefaults instantiates a new BriefIKEProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefIKEProposal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefIKEProposal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefIKEProposal) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefIKEProposal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefIKEProposal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefIKEProposal) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefIKEProposal) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefIKEProposal) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefIKEProposal) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefIKEProposal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefIKEProposal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefIKEProposal) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BriefIKEProposal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefIKEProposal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefIKEProposal) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefIKEProposal) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


