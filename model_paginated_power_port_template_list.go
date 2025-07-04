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

// checks if the PaginatedPowerPortTemplateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedPowerPortTemplateList{}

// PaginatedPowerPortTemplateList struct for PaginatedPowerPortTemplateList
type PaginatedPowerPortTemplateList struct {
	Count                int32               `json:"count"`
	Next                 NullableString      `json:"next,omitempty"`
	Previous             NullableString      `json:"previous,omitempty"`
	Results              []PowerPortTemplate `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedPowerPortTemplateList PaginatedPowerPortTemplateList

// NewPaginatedPowerPortTemplateList instantiates a new PaginatedPowerPortTemplateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPowerPortTemplateList(count int32, results []PowerPortTemplate) *PaginatedPowerPortTemplateList {
	this := PaginatedPowerPortTemplateList{}
	this.Count = count
	this.Results = results
	return &this
}

// NewPaginatedPowerPortTemplateListWithDefaults instantiates a new PaginatedPowerPortTemplateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPowerPortTemplateListWithDefaults() *PaginatedPowerPortTemplateList {
	this := PaginatedPowerPortTemplateList{}
	return &this
}

// GetCount returns the Count field value
func (o *PaginatedPowerPortTemplateList) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedPowerPortTemplateList) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedPowerPortTemplateList) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedPowerPortTemplateList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedPowerPortTemplateList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedPowerPortTemplateList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedPowerPortTemplateList) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedPowerPortTemplateList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedPowerPortTemplateList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedPowerPortTemplateList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedPowerPortTemplateList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedPowerPortTemplateList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedPowerPortTemplateList) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedPowerPortTemplateList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedPowerPortTemplateList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PaginatedPowerPortTemplateList) GetResults() []PowerPortTemplate {
	if o == nil {
		var ret []PowerPortTemplate
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPowerPortTemplateList) GetResultsOk() ([]PowerPortTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPowerPortTemplateList) SetResults(v []PowerPortTemplate) {
	o.Results = v
}

func (o PaginatedPowerPortTemplateList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedPowerPortTemplateList) ToMap() (map[string]interface{}, error) {
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

func (o *PaginatedPowerPortTemplateList) UnmarshalJSON(data []byte) (err error) {
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
	varPaginatedPowerPortTemplateList := _PaginatedPowerPortTemplateList{}

	err = json.Unmarshal(data, &varPaginatedPowerPortTemplateList)

	if err != nil {
		return err
	}

	*o = PaginatedPowerPortTemplateList(varPaginatedPowerPortTemplateList)

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

type NullablePaginatedPowerPortTemplateList struct {
	value *PaginatedPowerPortTemplateList
	isSet bool
}

func (v NullablePaginatedPowerPortTemplateList) Get() *PaginatedPowerPortTemplateList {
	return v.value
}

func (v *NullablePaginatedPowerPortTemplateList) Set(val *PaginatedPowerPortTemplateList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPowerPortTemplateList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPowerPortTemplateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPowerPortTemplateList(val *PaginatedPowerPortTemplateList) *NullablePaginatedPowerPortTemplateList {
	return &NullablePaginatedPowerPortTemplateList{value: val, isSet: true}
}

func (v NullablePaginatedPowerPortTemplateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPowerPortTemplateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
