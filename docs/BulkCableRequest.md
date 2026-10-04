# BulkCableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Type** | Pointer to [**NullableBulkCableRequestType**](BulkCableRequestType.md) |  | [optional] 
**ATerminations** | Pointer to [**[]GenericObjectRequest**](GenericObjectRequest.md) |  | [optional] 
**BTerminations** | Pointer to [**[]GenericObjectRequest**](GenericObjectRequest.md) |  | [optional] 
**Status** | Pointer to [**BulkCableRequestStatus**](BulkCableRequestStatus.md) |  | [optional] 
**Profile** | Pointer to [**BulkCableRequestProfile**](BulkCableRequestProfile.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Bundle** | Pointer to [**NullableBulkCableRequestBundle**](BulkCableRequestBundle.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to [**BriefModuleBayTypeColor**](BriefModuleBayTypeColor.md) |  | [optional] 
**Length** | Pointer to **NullableFloat64** |  | [optional] 
**LengthUnit** | Pointer to [**NullableBulkCableRequestLengthUnit**](BulkCableRequestLengthUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkCableRequest

`func NewBulkCableRequest(id int32, ) *BulkCableRequest`

NewBulkCableRequest instantiates a new BulkCableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCableRequestWithDefaults

`func NewBulkCableRequestWithDefaults() *BulkCableRequest`

NewBulkCableRequestWithDefaults instantiates a new BulkCableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCableRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCableRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCableRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *BulkCableRequest) GetType() BulkCableRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkCableRequest) GetTypeOk() (*BulkCableRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkCableRequest) SetType(v BulkCableRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *BulkCableRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BulkCableRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BulkCableRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetATerminations

`func (o *BulkCableRequest) GetATerminations() []GenericObjectRequest`

GetATerminations returns the ATerminations field if non-nil, zero value otherwise.

### GetATerminationsOk

`func (o *BulkCableRequest) GetATerminationsOk() (*[]GenericObjectRequest, bool)`

GetATerminationsOk returns a tuple with the ATerminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetATerminations

`func (o *BulkCableRequest) SetATerminations(v []GenericObjectRequest)`

SetATerminations sets ATerminations field to given value.

### HasATerminations

`func (o *BulkCableRequest) HasATerminations() bool`

HasATerminations returns a boolean if a field has been set.

### GetBTerminations

`func (o *BulkCableRequest) GetBTerminations() []GenericObjectRequest`

GetBTerminations returns the BTerminations field if non-nil, zero value otherwise.

### GetBTerminationsOk

`func (o *BulkCableRequest) GetBTerminationsOk() (*[]GenericObjectRequest, bool)`

GetBTerminationsOk returns a tuple with the BTerminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTerminations

`func (o *BulkCableRequest) SetBTerminations(v []GenericObjectRequest)`

SetBTerminations sets BTerminations field to given value.

### HasBTerminations

`func (o *BulkCableRequest) HasBTerminations() bool`

HasBTerminations returns a boolean if a field has been set.

### GetStatus

`func (o *BulkCableRequest) GetStatus() BulkCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkCableRequest) GetStatusOk() (*BulkCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkCableRequest) SetStatus(v BulkCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkCableRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProfile

`func (o *BulkCableRequest) GetProfile() BulkCableRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *BulkCableRequest) GetProfileOk() (*BulkCableRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *BulkCableRequest) SetProfile(v BulkCableRequestProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *BulkCableRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTenant

`func (o *BulkCableRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkCableRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkCableRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkCableRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkCableRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkCableRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetBundle

`func (o *BulkCableRequest) GetBundle() BulkCableRequestBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *BulkCableRequest) GetBundleOk() (*BulkCableRequestBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *BulkCableRequest) SetBundle(v BulkCableRequestBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *BulkCableRequest) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### SetBundleNil

`func (o *BulkCableRequest) SetBundleNil(b bool)`

 SetBundleNil sets the value for Bundle to be an explicit nil

### UnsetBundle
`func (o *BulkCableRequest) UnsetBundle()`

UnsetBundle ensures that no value is present for Bundle, not even an explicit nil
### GetLabel

`func (o *BulkCableRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkCableRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkCableRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkCableRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *BulkCableRequest) GetColor() BriefModuleBayTypeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BulkCableRequest) GetColorOk() (*BriefModuleBayTypeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BulkCableRequest) SetColor(v BriefModuleBayTypeColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *BulkCableRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *BulkCableRequest) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *BulkCableRequest) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *BulkCableRequest) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *BulkCableRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *BulkCableRequest) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *BulkCableRequest) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *BulkCableRequest) GetLengthUnit() BulkCableRequestLengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *BulkCableRequest) GetLengthUnitOk() (*BulkCableRequestLengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *BulkCableRequest) SetLengthUnit(v BulkCableRequestLengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *BulkCableRequest) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### SetLengthUnitNil

`func (o *BulkCableRequest) SetLengthUnitNil(b bool)`

 SetLengthUnitNil sets the value for LengthUnit to be an explicit nil

### UnsetLengthUnit
`func (o *BulkCableRequest) UnsetLengthUnit()`

UnsetLengthUnit ensures that no value is present for LengthUnit, not even an explicit nil
### GetDescription

`func (o *BulkCableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkCableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkCableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkCableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *BulkCableRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkCableRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkCableRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkCableRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkCableRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkCableRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkCableRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkCableRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkCableRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkCableRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkCableRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkCableRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkCableRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkCableRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkCableRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkCableRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkCableRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkCableRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


