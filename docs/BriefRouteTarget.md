# BriefRouteTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** | Route target value (formatted in accordance with RFC 4360) | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefRouteTarget

`func NewBriefRouteTarget(id int32, url string, display string, name string, ) *BriefRouteTarget`

NewBriefRouteTarget instantiates a new BriefRouteTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefRouteTargetWithDefaults

`func NewBriefRouteTargetWithDefaults() *BriefRouteTarget`

NewBriefRouteTargetWithDefaults instantiates a new BriefRouteTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefRouteTarget) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefRouteTarget) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefRouteTarget) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefRouteTarget) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefRouteTarget) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefRouteTarget) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefRouteTarget) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefRouteTarget) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefRouteTarget) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefRouteTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefRouteTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefRouteTarget) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BriefRouteTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefRouteTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefRouteTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefRouteTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


