# BriefObjectPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ObjectTypes** | **[]string** |  | 
**Actions** | **[]string** | The list of actions granted by this permission | 

## Methods

### NewBriefObjectPermission

`func NewBriefObjectPermission(id int32, url string, display string, name string, objectTypes []string, actions []string, ) *BriefObjectPermission`

NewBriefObjectPermission instantiates a new BriefObjectPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefObjectPermissionWithDefaults

`func NewBriefObjectPermissionWithDefaults() *BriefObjectPermission`

NewBriefObjectPermissionWithDefaults instantiates a new BriefObjectPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefObjectPermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefObjectPermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefObjectPermission) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefObjectPermission) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefObjectPermission) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefObjectPermission) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefObjectPermission) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefObjectPermission) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefObjectPermission) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefObjectPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefObjectPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefObjectPermission) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BriefObjectPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefObjectPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefObjectPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefObjectPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BriefObjectPermission) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BriefObjectPermission) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BriefObjectPermission) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BriefObjectPermission) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectTypes

`func (o *BriefObjectPermission) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *BriefObjectPermission) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *BriefObjectPermission) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetActions

`func (o *BriefObjectPermission) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BriefObjectPermission) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BriefObjectPermission) SetActions(v []string)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


