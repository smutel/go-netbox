# BulkOperationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** | A summary of the failure, e.g. \&quot;1 of 3 objects could not be updated.\&quot; | 
**Errors** | Pointer to [**[]BulkOperationEntryError**](BulkOperationEntryError.md) | One entry per object which failed; objects which would have succeeded are omitted, as a bulk operation is all-or-none. Absent where the request could not be attributed to individual entries at all (e.g. a request body which is not a list). | [optional] 

## Methods

### NewBulkOperationError

`func NewBulkOperationError(detail string, ) *BulkOperationError`

NewBulkOperationError instantiates a new BulkOperationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationErrorWithDefaults

`func NewBulkOperationErrorWithDefaults() *BulkOperationError`

NewBulkOperationErrorWithDefaults instantiates a new BulkOperationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *BulkOperationError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BulkOperationError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BulkOperationError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrors

`func (o *BulkOperationError) GetErrors() []BulkOperationEntryError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkOperationError) GetErrorsOk() (*[]BulkOperationEntryError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkOperationError) SetErrors(v []BulkOperationEntryError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkOperationError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


