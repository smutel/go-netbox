# PatchedBulkCustomLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**LinkText** | Pointer to **string** | Jinja2 template code for link text | [optional] 
**LinkUrl** | Pointer to **string** | Jinja2 template code for link URL | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Links with the same group will appear as a dropdown menu | [optional] 
**ButtonClass** | Pointer to [**BulkCustomLinkRequestButtonClass**](BulkCustomLinkRequestButtonClass.md) |  | [optional] 
**NewWindow** | Pointer to **bool** | Force link to open in a new window | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewPatchedBulkCustomLinkRequest

`func NewPatchedBulkCustomLinkRequest(id int32, ) *PatchedBulkCustomLinkRequest`

NewPatchedBulkCustomLinkRequest instantiates a new PatchedBulkCustomLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCustomLinkRequestWithDefaults

`func NewPatchedBulkCustomLinkRequestWithDefaults() *PatchedBulkCustomLinkRequest`

NewPatchedBulkCustomLinkRequestWithDefaults instantiates a new PatchedBulkCustomLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCustomLinkRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCustomLinkRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCustomLinkRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *PatchedBulkCustomLinkRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedBulkCustomLinkRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedBulkCustomLinkRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedBulkCustomLinkRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkCustomLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkCustomLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkCustomLinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkCustomLinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkCustomLinkRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkCustomLinkRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkCustomLinkRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkCustomLinkRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLinkText

`func (o *PatchedBulkCustomLinkRequest) GetLinkText() string`

GetLinkText returns the LinkText field if non-nil, zero value otherwise.

### GetLinkTextOk

`func (o *PatchedBulkCustomLinkRequest) GetLinkTextOk() (*string, bool)`

GetLinkTextOk returns a tuple with the LinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkText

`func (o *PatchedBulkCustomLinkRequest) SetLinkText(v string)`

SetLinkText sets LinkText field to given value.

### HasLinkText

`func (o *PatchedBulkCustomLinkRequest) HasLinkText() bool`

HasLinkText returns a boolean if a field has been set.

### GetLinkUrl

`func (o *PatchedBulkCustomLinkRequest) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *PatchedBulkCustomLinkRequest) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *PatchedBulkCustomLinkRequest) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *PatchedBulkCustomLinkRequest) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkCustomLinkRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkCustomLinkRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkCustomLinkRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkCustomLinkRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *PatchedBulkCustomLinkRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PatchedBulkCustomLinkRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PatchedBulkCustomLinkRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PatchedBulkCustomLinkRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *PatchedBulkCustomLinkRequest) GetButtonClass() BulkCustomLinkRequestButtonClass`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *PatchedBulkCustomLinkRequest) GetButtonClassOk() (*BulkCustomLinkRequestButtonClass, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *PatchedBulkCustomLinkRequest) SetButtonClass(v BulkCustomLinkRequestButtonClass)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *PatchedBulkCustomLinkRequest) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetNewWindow

`func (o *PatchedBulkCustomLinkRequest) GetNewWindow() bool`

GetNewWindow returns the NewWindow field if non-nil, zero value otherwise.

### GetNewWindowOk

`func (o *PatchedBulkCustomLinkRequest) GetNewWindowOk() (*bool, bool)`

GetNewWindowOk returns a tuple with the NewWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindow

`func (o *PatchedBulkCustomLinkRequest) SetNewWindow(v bool)`

SetNewWindow sets NewWindow field to given value.

### HasNewWindow

`func (o *PatchedBulkCustomLinkRequest) HasNewWindow() bool`

HasNewWindow returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkCustomLinkRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkCustomLinkRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkCustomLinkRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkCustomLinkRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkCustomLinkRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkCustomLinkRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


