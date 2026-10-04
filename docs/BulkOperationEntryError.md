# BulkOperationEntryError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the object which failed. Present once the entry has been matched to an object; mutually exclusive with &#x60;index&#x60;. | [optional] 
**Index** | Pointer to **int32** | The zero-based position of the entry within the submitted list. Used where no object has been identified for the entry: always for creations, and for updates and deletions where the entry itself could not be interpreted (e.g. a missing or non-numeric &#x60;id&#x60;). Mutually exclusive with &#x60;id&#x60;. | [optional] 
**Errors** | **map[string]interface{}** | The errors for this entry, keyed by field name. Values are ordinarily arrays of messages. Errors which pertain to no particular field -- model validation, protection rules, restricted tags, object-level permissions, or the shape of the entry itself -- all appear under the single key &#x60;__all__&#x60;. | 

## Methods

### NewBulkOperationEntryError

`func NewBulkOperationEntryError(errors map[string]interface{}, ) *BulkOperationEntryError`

NewBulkOperationEntryError instantiates a new BulkOperationEntryError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationEntryErrorWithDefaults

`func NewBulkOperationEntryErrorWithDefaults() *BulkOperationEntryError`

NewBulkOperationEntryErrorWithDefaults instantiates a new BulkOperationEntryError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkOperationEntryError) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkOperationEntryError) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkOperationEntryError) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BulkOperationEntryError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *BulkOperationEntryError) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BulkOperationEntryError) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BulkOperationEntryError) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *BulkOperationEntryError) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetErrors

`func (o *BulkOperationEntryError) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkOperationEntryError) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkOperationEntryError) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


