# PatchedBulkCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Cid** | Pointer to **string** | Unique circuit ID | [optional] 
**Provider** | Pointer to [**BriefCircuitRequestProvider**](BriefCircuitRequestProvider.md) |  | [optional] 
**ProviderAccount** | Pointer to [**NullableBulkCircuitRequestProviderAccount**](BulkCircuitRequestProviderAccount.md) |  | [optional] 
**Type** | Pointer to [**BulkCircuitRequestType**](BulkCircuitRequestType.md) |  | [optional] 
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

### NewPatchedBulkCircuitRequest

`func NewPatchedBulkCircuitRequest(id int32, ) *PatchedBulkCircuitRequest`

NewPatchedBulkCircuitRequest instantiates a new PatchedBulkCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkCircuitRequestWithDefaults

`func NewPatchedBulkCircuitRequestWithDefaults() *PatchedBulkCircuitRequest`

NewPatchedBulkCircuitRequestWithDefaults instantiates a new PatchedBulkCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkCircuitRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkCircuitRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkCircuitRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCid

`func (o *PatchedBulkCircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *PatchedBulkCircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *PatchedBulkCircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *PatchedBulkCircuitRequest) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetProvider

`func (o *PatchedBulkCircuitRequest) GetProvider() BriefCircuitRequestProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PatchedBulkCircuitRequest) GetProviderOk() (*BriefCircuitRequestProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PatchedBulkCircuitRequest) SetProvider(v BriefCircuitRequestProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PatchedBulkCircuitRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderAccount

`func (o *PatchedBulkCircuitRequest) GetProviderAccount() BulkCircuitRequestProviderAccount`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *PatchedBulkCircuitRequest) GetProviderAccountOk() (*BulkCircuitRequestProviderAccount, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *PatchedBulkCircuitRequest) SetProviderAccount(v BulkCircuitRequestProviderAccount)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *PatchedBulkCircuitRequest) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *PatchedBulkCircuitRequest) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *PatchedBulkCircuitRequest) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *PatchedBulkCircuitRequest) GetType() BulkCircuitRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkCircuitRequest) GetTypeOk() (*BulkCircuitRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkCircuitRequest) SetType(v BulkCircuitRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkCircuitRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkCircuitRequest) GetStatus() BulkCircuitRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkCircuitRequest) GetStatusOk() (*BulkCircuitRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkCircuitRequest) SetStatus(v BulkCircuitRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkCircuitRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkCircuitRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkCircuitRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkCircuitRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkCircuitRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkCircuitRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkCircuitRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetInstallDate

`func (o *PatchedBulkCircuitRequest) GetInstallDate() string`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *PatchedBulkCircuitRequest) GetInstallDateOk() (*string, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *PatchedBulkCircuitRequest) SetInstallDate(v string)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *PatchedBulkCircuitRequest) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### SetInstallDateNil

`func (o *PatchedBulkCircuitRequest) SetInstallDateNil(b bool)`

 SetInstallDateNil sets the value for InstallDate to be an explicit nil

### UnsetInstallDate
`func (o *PatchedBulkCircuitRequest) UnsetInstallDate()`

UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
### GetTerminationDate

`func (o *PatchedBulkCircuitRequest) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *PatchedBulkCircuitRequest) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *PatchedBulkCircuitRequest) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *PatchedBulkCircuitRequest) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *PatchedBulkCircuitRequest) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *PatchedBulkCircuitRequest) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetCommitRate

`func (o *PatchedBulkCircuitRequest) GetCommitRate() int32`

GetCommitRate returns the CommitRate field if non-nil, zero value otherwise.

### GetCommitRateOk

`func (o *PatchedBulkCircuitRequest) GetCommitRateOk() (*int32, bool)`

GetCommitRateOk returns a tuple with the CommitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRate

`func (o *PatchedBulkCircuitRequest) SetCommitRate(v int32)`

SetCommitRate sets CommitRate field to given value.

### HasCommitRate

`func (o *PatchedBulkCircuitRequest) HasCommitRate() bool`

HasCommitRate returns a boolean if a field has been set.

### SetCommitRateNil

`func (o *PatchedBulkCircuitRequest) SetCommitRateNil(b bool)`

 SetCommitRateNil sets the value for CommitRate to be an explicit nil

### UnsetCommitRate
`func (o *PatchedBulkCircuitRequest) UnsetCommitRate()`

UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
### GetDescription

`func (o *PatchedBulkCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistance

`func (o *PatchedBulkCircuitRequest) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *PatchedBulkCircuitRequest) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *PatchedBulkCircuitRequest) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *PatchedBulkCircuitRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *PatchedBulkCircuitRequest) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *PatchedBulkCircuitRequest) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetDistanceUnit

`func (o *PatchedBulkCircuitRequest) GetDistanceUnit() BulkCircuitRequestDistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *PatchedBulkCircuitRequest) GetDistanceUnitOk() (*BulkCircuitRequestDistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *PatchedBulkCircuitRequest) SetDistanceUnit(v BulkCircuitRequestDistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *PatchedBulkCircuitRequest) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.

### SetDistanceUnitNil

`func (o *PatchedBulkCircuitRequest) SetDistanceUnitNil(b bool)`

 SetDistanceUnitNil sets the value for DistanceUnit to be an explicit nil

### UnsetDistanceUnit
`func (o *PatchedBulkCircuitRequest) UnsetDistanceUnit()`

UnsetDistanceUnit ensures that no value is present for DistanceUnit, not even an explicit nil
### GetOwner

`func (o *PatchedBulkCircuitRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkCircuitRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkCircuitRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkCircuitRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedBulkCircuitRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedBulkCircuitRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *PatchedBulkCircuitRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkCircuitRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkCircuitRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkCircuitRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkCircuitRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkCircuitRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkCircuitRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkCircuitRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkCircuitRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkCircuitRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkCircuitRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAssignments

`func (o *PatchedBulkCircuitRequest) GetAssignments() []BriefCircuitGroupAssignmentSerializerRequest`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *PatchedBulkCircuitRequest) GetAssignmentsOk() (*[]BriefCircuitGroupAssignmentSerializerRequest, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *PatchedBulkCircuitRequest) SetAssignments(v []BriefCircuitGroupAssignmentSerializerRequest)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *PatchedBulkCircuitRequest) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


