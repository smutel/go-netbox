/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedRearPortTemplateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedRearPortTemplateList{}

// PaginatedRearPortTemplateList struct for PaginatedRearPortTemplateList
type PaginatedRearPortTemplateList struct {
	Count                int32              `json:"count"`
	Next                 NullableString     `json:"next,omitempty"`
	Previous             NullableString     `json:"previous,omitempty"`
	Results              []RearPortTemplate `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedRearPortTemplateList PaginatedRearPortTemplateList

// NewPaginatedRearPortTemplateList instantiates a new PaginatedRearPortTemplateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRearPortTemplateList(count int32, results []RearPortTemplate) *PaginatedRearPortTemplateList {
	this := PaginatedRearPortTemplateList{}
	this.Count = count
	this.Results = results
	return &this
}

// NewPaginatedRearPortTemplateListWithDefaults instantiates a new PaginatedRearPortTemplateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRearPortTemplateListWithDefaults() *PaginatedRearPortTemplateList {
	this := PaginatedRearPortTemplateList{}
	return &this
}

// GetCount returns the Count field value
func (o *PaginatedRearPortTemplateList) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedRearPortTemplateList) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedRearPortTemplateList) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedRearPortTemplateList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedRearPortTemplateList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedRearPortTemplateList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedRearPortTemplateList) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedRearPortTemplateList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedRearPortTemplateList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedRearPortTemplateList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedRearPortTemplateList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedRearPortTemplateList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedRearPortTemplateList) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedRearPortTemplateList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedRearPortTemplateList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PaginatedRearPortTemplateList) GetResults() []RearPortTemplate {
	if o == nil {
		var ret []RearPortTemplate
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRearPortTemplateList) GetResultsOk() ([]RearPortTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRearPortTemplateList) SetResults(v []RearPortTemplate) {
	o.Results = v
}

func (o PaginatedRearPortTemplateList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedRearPortTemplateList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedRearPortTemplateList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"results",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varPaginatedRearPortTemplateList := _PaginatedRearPortTemplateList{}

	err = json.Unmarshal(data, &varPaginatedRearPortTemplateList)

	if err != nil {
		return err
	}

	*o = PaginatedRearPortTemplateList(varPaginatedRearPortTemplateList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedRearPortTemplateList struct {
	value *PaginatedRearPortTemplateList
	isSet bool
}

func (v NullablePaginatedRearPortTemplateList) Get() *PaginatedRearPortTemplateList {
	return v.value
}

func (v *NullablePaginatedRearPortTemplateList) Set(val *PaginatedRearPortTemplateList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRearPortTemplateList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRearPortTemplateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRearPortTemplateList(val *PaginatedRearPortTemplateList) *NullablePaginatedRearPortTemplateList {
	return &NullablePaginatedRearPortTemplateList{value: val, isSet: true}
}

func (v NullablePaginatedRearPortTemplateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRearPortTemplateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
