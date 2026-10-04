# DcimCableBundlesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDcimCableBundlesCreateRequest

`func NewDcimCableBundlesCreateRequest(name string, ) *DcimCableBundlesCreateRequest`

NewDcimCableBundlesCreateRequest instantiates a new DcimCableBundlesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimCableBundlesCreateRequestWithDefaults

`func NewDcimCableBundlesCreateRequestWithDefaults() *DcimCableBundlesCreateRequest`

NewDcimCableBundlesCreateRequestWithDefaults instantiates a new DcimCableBundlesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DcimCableBundlesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimCableBundlesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimCableBundlesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DcimCableBundlesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimCableBundlesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimCableBundlesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimCableBundlesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *DcimCableBundlesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DcimCableBundlesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DcimCableBundlesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DcimCableBundlesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DcimCableBundlesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DcimCableBundlesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *DcimCableBundlesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DcimCableBundlesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DcimCableBundlesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DcimCableBundlesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *DcimCableBundlesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DcimCableBundlesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DcimCableBundlesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DcimCableBundlesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DcimCableBundlesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DcimCableBundlesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DcimCableBundlesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DcimCableBundlesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


