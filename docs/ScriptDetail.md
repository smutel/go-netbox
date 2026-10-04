# ScriptDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Module** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Description** | **NullableString** |  | [readonly] 
**Vars** | **interface{}** |  | [readonly] 
**Result** | [**Job**](Job.md) |  | [readonly] 
**Display** | **string** |  | [readonly] 
**IsExecutable** | **bool** |  | [readonly] 

## Methods

### NewScriptDetail

`func NewScriptDetail(id int32, url string, displayUrl string, module int32, name string, description NullableString, vars interface{}, result Job, display string, isExecutable bool, ) *ScriptDetail`

NewScriptDetail instantiates a new ScriptDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptDetailWithDefaults

`func NewScriptDetailWithDefaults() *ScriptDetail`

NewScriptDetailWithDefaults instantiates a new ScriptDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScriptDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ScriptDetail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScriptDetail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScriptDetail) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *ScriptDetail) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ScriptDetail) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ScriptDetail) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetModule

`func (o *ScriptDetail) GetModule() int32`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ScriptDetail) GetModuleOk() (*int32, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ScriptDetail) SetModule(v int32)`

SetModule sets Module field to given value.


### GetName

`func (o *ScriptDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptDetail) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScriptDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScriptDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScriptDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ScriptDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScriptDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVars

`func (o *ScriptDetail) GetVars() interface{}`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *ScriptDetail) GetVarsOk() (*interface{}, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *ScriptDetail) SetVars(v interface{})`

SetVars sets Vars field to given value.


### SetVarsNil

`func (o *ScriptDetail) SetVarsNil(b bool)`

 SetVarsNil sets the value for Vars to be an explicit nil

### UnsetVars
`func (o *ScriptDetail) UnsetVars()`

UnsetVars ensures that no value is present for Vars, not even an explicit nil
### GetResult

`func (o *ScriptDetail) GetResult() Job`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ScriptDetail) GetResultOk() (*Job, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ScriptDetail) SetResult(v Job)`

SetResult sets Result field to given value.


### GetDisplay

`func (o *ScriptDetail) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ScriptDetail) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ScriptDetail) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetIsExecutable

`func (o *ScriptDetail) GetIsExecutable() bool`

GetIsExecutable returns the IsExecutable field if non-nil, zero value otherwise.

### GetIsExecutableOk

`func (o *ScriptDetail) GetIsExecutableOk() (*bool, bool)`

GetIsExecutableOk returns a tuple with the IsExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExecutable

`func (o *ScriptDetail) SetIsExecutable(v bool)`

SetIsExecutable sets IsExecutable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


