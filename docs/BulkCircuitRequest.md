# BulkCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Cid** | **string** | Unique circuit ID | 
**Provider** | [**BriefCircuitRequestProvider**](BriefCircuitRequestProvider.md) |  | 
**ProviderAccount** | Pointer to [**NullableBulkCircuitRequestProviderAccount**](BulkCircuitRequestProviderAccount.md) |  | [optional] 
**Type** | [**BulkCircuitRequestType**](BulkCircuitRequestType.md) |  | 
**Status** | Pointer to [**BulkCircuitRequestStatus**](BulkCircuitRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**InstallDate** | Pointer to **NullableString** |  | [optional] 
**TerminationDate** | Pointer to **NullableString** |  | [optional] 
**CommitRate** | Pointer to **NullableInt32** | Committed rate | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **NullableFloat64** |  | [optional] 
**DistanceUnit** | Pointer to [**NullableBulkCircuitRequestDistanceUnit**](BulkCircuitRequestDistanceUnit.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Assignments** | Pointer to [**[]BriefCircuitGroupAssignmentSerializerRequest**](BriefCircuitGroupAssignmentSerializerRequest.md) |  | [optional] 

## Methods

### NewBulkCircuitRequest

`func NewBulkCircuitRequest(id int32, cid string, provider BriefCircuitRequestProvider, type_ BulkCircuitRequestType, ) *BulkCircuitRequest`

NewBulkCircuitRequest instantiates a new BulkCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCircuitRequestWithDefaults

`func NewBulkCircuitRequestWithDefaults() *BulkCircuitRequest`

NewBulkCircuitRequestWithDefaults instantiates a new BulkCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkCircuitRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkCircuitRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkCircuitRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCid

`func (o *BulkCircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *BulkCircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *BulkCircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *BulkCircuitRequest) GetProvider() BriefCircuitRequestProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BulkCircuitRequest) GetProviderOk() (*BriefCircuitRequestProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BulkCircuitRequest) SetProvider(v BriefCircuitRequestProvider)`

SetProvider sets Provider field to given value.


### GetProviderAccount

`func (o *BulkCircuitRequest) GetProviderAccount() BulkCircuitRequestProviderAccount`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *BulkCircuitRequest) GetProviderAccountOk() (*BulkCircuitRequestProviderAccount, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *BulkCircuitRequest) SetProviderAccount(v BulkCircuitRequestProviderAccount)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *BulkCircuitRequest) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *BulkCircuitRequest) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *BulkCircuitRequest) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *BulkCircuitRequest) GetType() BulkCircuitRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkCircuitRequest) GetTypeOk() (*BulkCircuitRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkCircuitRequest) SetType(v BulkCircuitRequestType)`

SetType sets Type field to given value.


### GetStatus

`func (o *BulkCircuitRequest) GetStatus() BulkCircuitRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkCircuitRequest) GetStatusOk() (*BulkCircuitRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkCircuitRequest) SetStatus(v BulkCircuitRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkCircuitRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *BulkCircuitRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkCircuitRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkCircuitRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkCircuitRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkCircuitRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkCircuitRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetInstallDate

`func (o *BulkCircuitRequest) GetInstallDate() string`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *BulkCircuitRequest) GetInstallDateOk() (*string, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *BulkCircuitRequest) SetInstallDate(v string)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *BulkCircuitRequest) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### SetInstallDateNil

`func (o *BulkCircuitRequest) SetInstallDateNil(b bool)`

 SetInstallDateNil sets the value for InstallDate to be an explicit nil

### UnsetInstallDate
`func (o *BulkCircuitRequest) UnsetInstallDate()`

UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
### GetTerminationDate

`func (o *BulkCircuitRequest) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *BulkCircuitRequest) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *BulkCircuitRequest) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *BulkCircuitRequest) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *BulkCircuitRequest) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *BulkCircuitRequest) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetCommitRate

`func (o *BulkCircuitRequest) GetCommitRate() int32`

GetCommitRate returns the CommitRate field if non-nil, zero value otherwise.

### GetCommitRateOk

`func (o *BulkCircuitRequest) GetCommitRateOk() (*int32, bool)`

GetCommitRateOk returns a tuple with the CommitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRate

`func (o *BulkCircuitRequest) SetCommitRate(v int32)`

SetCommitRate sets CommitRate field to given value.

### HasCommitRate

`func (o *BulkCircuitRequest) HasCommitRate() bool`

HasCommitRate returns a boolean if a field has been set.

### SetCommitRateNil

`func (o *BulkCircuitRequest) SetCommitRateNil(b bool)`

 SetCommitRateNil sets the value for CommitRate to be an explicit nil

### UnsetCommitRate
`func (o *BulkCircuitRequest) UnsetCommitRate()`

UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
### GetDescription

`func (o *BulkCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistance

`func (o *BulkCircuitRequest) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *BulkCircuitRequest) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *BulkCircuitRequest) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *BulkCircuitRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *BulkCircuitRequest) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *BulkCircuitRequest) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetDistanceUnit

`func (o *BulkCircuitRequest) GetDistanceUnit() BulkCircuitRequestDistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *BulkCircuitRequest) GetDistanceUnitOk() (*BulkCircuitRequestDistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *BulkCircuitRequest) SetDistanceUnit(v BulkCircuitRequestDistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *BulkCircuitRequest) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.

### SetDistanceUnitNil

`func (o *BulkCircuitRequest) SetDistanceUnitNil(b bool)`

 SetDistanceUnitNil sets the value for DistanceUnit to be an explicit nil

### UnsetDistanceUnit
`func (o *BulkCircuitRequest) UnsetDistanceUnit()`

UnsetDistanceUnit ensures that no value is present for DistanceUnit, not even an explicit nil
### GetOwner

`func (o *BulkCircuitRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkCircuitRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkCircuitRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BulkCircuitRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BulkCircuitRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BulkCircuitRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *BulkCircuitRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkCircuitRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkCircuitRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkCircuitRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *BulkCircuitRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkCircuitRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkCircuitRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkCircuitRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkCircuitRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkCircuitRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkCircuitRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAssignments

`func (o *BulkCircuitRequest) GetAssignments() []BriefCircuitGroupAssignmentSerializerRequest`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *BulkCircuitRequest) GetAssignmentsOk() (*[]BriefCircuitGroupAssignmentSerializerRequest, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *BulkCircuitRequest) SetAssignments(v []BriefCircuitGroupAssignmentSerializerRequest)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *BulkCircuitRequest) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


