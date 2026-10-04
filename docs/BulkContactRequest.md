# BulkContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Name** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to [**BulkContactRequestEmail**](BulkContactRequestEmail.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Link** | Pointer to [**BulkContactRequestLink**](BulkContactRequestLink.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkContactRequest

`func NewBulkContactRequest(id int32, name string, ) *BulkContactRequest`

NewBulkContactRequest instantiates a new BulkContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkContactRequestWithDefaults

`func NewBulkContactRequestWithDefaults() *BulkContactRequest`

NewBulkContactRequestWithDefaults instantiates a new BulkContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkContactRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkContactRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkContactRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetGroups

`func (o *BulkContactRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BulkContactRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BulkContactRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BulkContactRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetName

`func (o *BulkContactRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkContactRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkContactRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *BulkContactRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BulkContactRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BulkContactRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BulkContactRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPhone

`func (o *BulkContactRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BulkContactRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BulkContactRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BulkContactRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *BulkContactRequest) GetEmail() BulkContactRequestEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BulkContactRequest) GetEmailOk() (*BulkContactRequestEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BulkContactRequest) SetEmail(v BulkContactRequestEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BulkContactRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *BulkContactRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BulkContactRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BulkContactRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BulkContactRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLink

`func (o *BulkContactRequest) GetLink() BulkContactRequestLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *BulkContactRequest) GetLinkOk() (*BulkContactRequestLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *BulkContactRequest) SetLink(v BulkContactRequestLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *BulkContactRequest) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDescription

`func (o *BulkContactRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkContactRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkContactRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkContactRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkContactRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkContactRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkContactRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkContactRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkContactRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkContactRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkContactRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkContactRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkContactRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkContactRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkContactRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkContactRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkContactRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkContactRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkContactRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkContactRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkContactRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkContactRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


