# PaginatedCoolingOutflowTemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]CoolingOutflowTemplate**](CoolingOutflowTemplate.md) |  | 

## Methods

### NewPaginatedCoolingOutflowTemplateList

`func NewPaginatedCoolingOutflowTemplateList(count int32, results []CoolingOutflowTemplate, ) *PaginatedCoolingOutflowTemplateList`

NewPaginatedCoolingOutflowTemplateList instantiates a new PaginatedCoolingOutflowTemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCoolingOutflowTemplateListWithDefaults

`func NewPaginatedCoolingOutflowTemplateListWithDefaults() *PaginatedCoolingOutflowTemplateList`

NewPaginatedCoolingOutflowTemplateListWithDefaults instantiates a new PaginatedCoolingOutflowTemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedCoolingOutflowTemplateList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedCoolingOutflowTemplateList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedCoolingOutflowTemplateList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedCoolingOutflowTemplateList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedCoolingOutflowTemplateList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedCoolingOutflowTemplateList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedCoolingOutflowTemplateList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedCoolingOutflowTemplateList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedCoolingOutflowTemplateList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedCoolingOutflowTemplateList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedCoolingOutflowTemplateList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedCoolingOutflowTemplateList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedCoolingOutflowTemplateList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedCoolingOutflowTemplateList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedCoolingOutflowTemplateList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedCoolingOutflowTemplateList) GetResults() []CoolingOutflowTemplate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCoolingOutflowTemplateList) GetResultsOk() (*[]CoolingOutflowTemplate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCoolingOutflowTemplateList) SetResults(v []CoolingOutflowTemplate)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


