# BulkCustomLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ObjectTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**LinkText** | **string** | Jinja2 template code for link text | 
**LinkUrl** | **string** | Jinja2 template code for link URL | 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Links with the same group will appear as a dropdown menu | [optional] 
**ButtonClass** | Pointer to [**BulkCustomLinkRequestButtonClass**](BulkCustomLinkRequestButtonClass.md) |  | [optional] 
**NewWindow** | Pointer to **bool** | Force link to open in a new window | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewBulkCustomLinkRequest

`func NewBulkCustomLinkRequest(id int32, objectTypes []string, name string, linkText string, linkUrl string, ) *BulkCustomLinkRequest`

NewBulkCustomLinkRequest instantiates a new BulkCustomLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCustomLinkRequestWithDefaults

`func NewBulkCustomLinkRequestWithDefaults() *BulkCustomLinkRequest`

NewBulkCustomLinkRequestWithDefaults instantiates a new BulkCustomLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCustomLinkRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCustomLinkRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCustomLinkRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectTypes

`func (o *BulkCustomLinkRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *BulkCustomLinkRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *BulkCustomLinkRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetName

`func (o *BulkCustomLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkCustomLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkCustomLinkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *BulkCustomLinkRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkCustomLinkRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkCustomLinkRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkCustomLinkRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLinkText

`func (o *BulkCustomLinkRequest) GetLinkText() string`

GetLinkText returns the LinkText field if non-nil, zero value otherwise.

### GetLinkTextOk

`func (o *BulkCustomLinkRequest) GetLinkTextOk() (*string, bool)`

GetLinkTextOk returns a tuple with the LinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkText

`func (o *BulkCustomLinkRequest) SetLinkText(v string)`

SetLinkText sets LinkText field to given value.


### GetLinkUrl

`func (o *BulkCustomLinkRequest) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *BulkCustomLinkRequest) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *BulkCustomLinkRequest) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.


### GetWeight

`func (o *BulkCustomLinkRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkCustomLinkRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkCustomLinkRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkCustomLinkRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *BulkCustomLinkRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *BulkCustomLinkRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *BulkCustomLinkRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *BulkCustomLinkRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *BulkCustomLinkRequest) GetButtonClass() BulkCustomLinkRequestButtonClass`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *BulkCustomLinkRequest) GetButtonClassOk() (*BulkCustomLinkRequestButtonClass, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *BulkCustomLinkRequest) SetButtonClass(v BulkCustomLinkRequestButtonClass)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *BulkCustomLinkRequest) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetNewWindow

`func (o *BulkCustomLinkRequest) GetNewWindow() bool`

GetNewWindow returns the NewWindow field if non-nil, zero value otherwise.

### GetNewWindowOk

`func (o *BulkCustomLinkRequest) GetNewWindowOk() (*bool, bool)`

GetNewWindowOk returns a tuple with the NewWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindow

`func (o *BulkCustomLinkRequest) SetNewWindow(v bool)`

SetNewWindow sets NewWindow field to given value.

### HasNewWindow

`func (o *BulkCustomLinkRequest) HasNewWindow() bool`

HasNewWindow returns a boolean if a field has been set.

### GetOwner

`func (o *BulkCustomLinkRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkCustomLinkRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkCustomLinkRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkCustomLinkRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkCustomLinkRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkCustomLinkRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


