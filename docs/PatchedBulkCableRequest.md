# PatchedBulkCableRequest

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

### NewPatchedBulkCableRequest

`func NewPatchedBulkCableRequest(id int32, ) *PatchedBulkCableRequest`

NewPatchedBulkCableRequest instantiates a new PatchedBulkCableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCableRequestWithDefaults

`func NewPatchedBulkCableRequestWithDefaults() *PatchedBulkCableRequest`

NewPatchedBulkCableRequestWithDefaults instantiates a new PatchedBulkCableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCableRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCableRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCableRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *PatchedBulkCableRequest) GetType() BulkCableRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkCableRequest) GetTypeOk() (*BulkCableRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkCableRequest) SetType(v BulkCableRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkCableRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchedBulkCableRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchedBulkCableRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetATerminations

`func (o *PatchedBulkCableRequest) GetATerminations() []GenericObjectRequest`

GetATerminations returns the ATerminations field if non-nil, zero value otherwise.

### GetATerminationsOk

`func (o *PatchedBulkCableRequest) GetATerminationsOk() (*[]GenericObjectRequest, bool)`

GetATerminationsOk returns a tuple with the ATerminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetATerminations

`func (o *PatchedBulkCableRequest) SetATerminations(v []GenericObjectRequest)`

SetATerminations sets ATerminations field to given value.

### HasATerminations

`func (o *PatchedBulkCableRequest) HasATerminations() bool`

HasATerminations returns a boolean if a field has been set.

### GetBTerminations

`func (o *PatchedBulkCableRequest) GetBTerminations() []GenericObjectRequest`

GetBTerminations returns the BTerminations field if non-nil, zero value otherwise.

### GetBTerminationsOk

`func (o *PatchedBulkCableRequest) GetBTerminationsOk() (*[]GenericObjectRequest, bool)`

GetBTerminationsOk returns a tuple with the BTerminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTerminations

`func (o *PatchedBulkCableRequest) SetBTerminations(v []GenericObjectRequest)`

SetBTerminations sets BTerminations field to given value.

### HasBTerminations

`func (o *PatchedBulkCableRequest) HasBTerminations() bool`

HasBTerminations returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkCableRequest) GetStatus() BulkCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkCableRequest) GetStatusOk() (*BulkCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkCableRequest) SetStatus(v BulkCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkCableRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProfile

`func (o *PatchedBulkCableRequest) GetProfile() BulkCableRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PatchedBulkCableRequest) GetProfileOk() (*BulkCableRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PatchedBulkCableRequest) SetProfile(v BulkCableRequestProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PatchedBulkCableRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkCableRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkCableRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkCableRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkCableRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkCableRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkCableRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetBundle

`func (o *PatchedBulkCableRequest) GetBundle() BulkCableRequestBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *PatchedBulkCableRequest) GetBundleOk() (*BulkCableRequestBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *PatchedBulkCableRequest) SetBundle(v BulkCableRequestBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *PatchedBulkCableRequest) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### SetBundleNil

`func (o *PatchedBulkCableRequest) SetBundleNil(b bool)`

 SetBundleNil sets the value for Bundle to be an explicit nil

### UnsetBundle
`func (o *PatchedBulkCableRequest) UnsetBundle()`

UnsetBundle ensures that no value is present for Bundle, not even an explicit nil
### GetLabel

`func (o *PatchedBulkCableRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkCableRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkCableRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkCableRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *PatchedBulkCableRequest) GetColor() BriefModuleBayTypeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedBulkCableRequest) GetColorOk() (*BriefModuleBayTypeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedBulkCableRequest) SetColor(v BriefModuleBayTypeColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedBulkCableRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *PatchedBulkCableRequest) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PatchedBulkCableRequest) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PatchedBulkCableRequest) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *PatchedBulkCableRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *PatchedBulkCableRequest) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *PatchedBulkCableRequest) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *PatchedBulkCableRequest) GetLengthUnit() BulkCableRequestLengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *PatchedBulkCableRequest) GetLengthUnitOk() (*BulkCableRequestLengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *PatchedBulkCableRequest) SetLengthUnit(v BulkCableRequestLengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *PatchedBulkCableRequest) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### SetLengthUnitNil

`func (o *PatchedBulkCableRequest) SetLengthUnitNil(b bool)`

 SetLengthUnitNil sets the value for LengthUnit to be an explicit nil

### UnsetLengthUnit
`func (o *PatchedBulkCableRequest) UnsetLengthUnit()`

UnsetLengthUnit ensures that no value is present for LengthUnit, not even an explicit nil
### GetDescription

`func (o *PatchedBulkCableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkCableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkCableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkCableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkCableRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkCableRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkCableRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkCableRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkCableRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkCableRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkCableRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkCableRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkCableRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkCableRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkCableRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkCableRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkCableRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkCableRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkCableRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkCableRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkCableRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkCableRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


