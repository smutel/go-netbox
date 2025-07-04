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

// checks if the CustomFieldChoiceSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldChoiceSetRequest{}

// CustomFieldChoiceSetRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CustomFieldChoiceSetRequest struct {
	Name         string                                `json:"name"`
	Description  *string                               `json:"description,omitempty"`
	BaseChoices  *CustomFieldChoiceSetBaseChoicesValue `json:"base_choices,omitempty"`
	ExtraChoices [][]interface{}                       `json:"extra_choices"`
	// Choices are automatically ordered alphabetically
	OrderAlphabetically  *bool `json:"order_alphabetically,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomFieldChoiceSetRequest CustomFieldChoiceSetRequest

// NewCustomFieldChoiceSetRequest instantiates a new CustomFieldChoiceSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldChoiceSetRequest(name string, extraChoices [][]interface{}) *CustomFieldChoiceSetRequest {
	this := CustomFieldChoiceSetRequest{}
	this.Name = name
	this.ExtraChoices = extraChoices
	return &this
}

// NewCustomFieldChoiceSetRequestWithDefaults instantiates a new CustomFieldChoiceSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldChoiceSetRequestWithDefaults() *CustomFieldChoiceSetRequest {
	this := CustomFieldChoiceSetRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CustomFieldChoiceSetRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSetRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomFieldChoiceSetRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomFieldChoiceSetRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSetRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomFieldChoiceSetRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomFieldChoiceSetRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBaseChoices returns the BaseChoices field value if set, zero value otherwise.
func (o *CustomFieldChoiceSetRequest) GetBaseChoices() CustomFieldChoiceSetBaseChoicesValue {
	if o == nil || IsNil(o.BaseChoices) {
		var ret CustomFieldChoiceSetBaseChoicesValue
		return ret
	}
	return *o.BaseChoices
}

// GetBaseChoicesOk returns a tuple with the BaseChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSetRequest) GetBaseChoicesOk() (*CustomFieldChoiceSetBaseChoicesValue, bool) {
	if o == nil || IsNil(o.BaseChoices) {
		return nil, false
	}
	return o.BaseChoices, true
}

// HasBaseChoices returns a boolean if a field has been set.
func (o *CustomFieldChoiceSetRequest) HasBaseChoices() bool {
	if o != nil && !IsNil(o.BaseChoices) {
		return true
	}

	return false
}

// SetBaseChoices gets a reference to the given CustomFieldChoiceSetBaseChoicesValue and assigns it to the BaseChoices field.
func (o *CustomFieldChoiceSetRequest) SetBaseChoices(v CustomFieldChoiceSetBaseChoicesValue) {
	o.BaseChoices = &v
}

// GetExtraChoices returns the ExtraChoices field value
func (o *CustomFieldChoiceSetRequest) GetExtraChoices() [][]interface{} {
	if o == nil {
		var ret [][]interface{}
		return ret
	}

	return o.ExtraChoices
}

// GetExtraChoicesOk returns a tuple with the ExtraChoices field value
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSetRequest) GetExtraChoicesOk() ([][]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraChoices, true
}

// SetExtraChoices sets field value
func (o *CustomFieldChoiceSetRequest) SetExtraChoices(v [][]interface{}) {
	o.ExtraChoices = v
}

// GetOrderAlphabetically returns the OrderAlphabetically field value if set, zero value otherwise.
func (o *CustomFieldChoiceSetRequest) GetOrderAlphabetically() bool {
	if o == nil || IsNil(o.OrderAlphabetically) {
		var ret bool
		return ret
	}
	return *o.OrderAlphabetically
}

// GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSetRequest) GetOrderAlphabeticallyOk() (*bool, bool) {
	if o == nil || IsNil(o.OrderAlphabetically) {
		return nil, false
	}
	return o.OrderAlphabetically, true
}

// HasOrderAlphabetically returns a boolean if a field has been set.
func (o *CustomFieldChoiceSetRequest) HasOrderAlphabetically() bool {
	if o != nil && !IsNil(o.OrderAlphabetically) {
		return true
	}

	return false
}

// SetOrderAlphabetically gets a reference to the given bool and assigns it to the OrderAlphabetically field.
func (o *CustomFieldChoiceSetRequest) SetOrderAlphabetically(v bool) {
	o.OrderAlphabetically = &v
}

func (o CustomFieldChoiceSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldChoiceSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BaseChoices) {
		toSerialize["base_choices"] = o.BaseChoices
	}
	toSerialize["extra_choices"] = o.ExtraChoices
	if !IsNil(o.OrderAlphabetically) {
		toSerialize["order_alphabetically"] = o.OrderAlphabetically
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomFieldChoiceSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"extra_choices",
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
	varCustomFieldChoiceSetRequest := _CustomFieldChoiceSetRequest{}

	err = json.Unmarshal(data, &varCustomFieldChoiceSetRequest)

	if err != nil {
		return err
	}

	*o = CustomFieldChoiceSetRequest(varCustomFieldChoiceSetRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "base_choices")
		delete(additionalProperties, "extra_choices")
		delete(additionalProperties, "order_alphabetically")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomFieldChoiceSetRequest struct {
	value *CustomFieldChoiceSetRequest
	isSet bool
}

func (v NullableCustomFieldChoiceSetRequest) Get() *CustomFieldChoiceSetRequest {
	return v.value
}

func (v *NullableCustomFieldChoiceSetRequest) Set(val *CustomFieldChoiceSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldChoiceSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldChoiceSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldChoiceSetRequest(val *CustomFieldChoiceSetRequest) *NullableCustomFieldChoiceSetRequest {
	return &NullableCustomFieldChoiceSetRequest{value: val, isSet: true}
}

func (v NullableCustomFieldChoiceSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldChoiceSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
