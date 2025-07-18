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
	"time"
)

// checks if the TokenProvisionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenProvisionRequest{}

// TokenProvisionRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type TokenProvisionRequest struct {
	Expires NullableTime `json:"expires,omitempty"`
	// Permit create/update/delete operations using this key
	WriteEnabled         *bool   `json:"write_enabled,omitempty"`
	Description          *string `json:"description,omitempty"`
	Username             string  `json:"username"`
	Password             string  `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _TokenProvisionRequest TokenProvisionRequest

// NewTokenProvisionRequest instantiates a new TokenProvisionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenProvisionRequest(username string, password string) *TokenProvisionRequest {
	this := TokenProvisionRequest{}
	this.Username = username
	this.Password = password
	return &this
}

// NewTokenProvisionRequestWithDefaults instantiates a new TokenProvisionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenProvisionRequestWithDefaults() *TokenProvisionRequest {
	this := TokenProvisionRequest{}
	return &this
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenProvisionRequest) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenProvisionRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *TokenProvisionRequest) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *TokenProvisionRequest) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *TokenProvisionRequest) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *TokenProvisionRequest) UnsetExpires() {
	o.Expires.Unset()
}

// GetWriteEnabled returns the WriteEnabled field value if set, zero value otherwise.
func (o *TokenProvisionRequest) GetWriteEnabled() bool {
	if o == nil || IsNil(o.WriteEnabled) {
		var ret bool
		return ret
	}
	return *o.WriteEnabled
}

// GetWriteEnabledOk returns a tuple with the WriteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProvisionRequest) GetWriteEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WriteEnabled) {
		return nil, false
	}
	return o.WriteEnabled, true
}

// HasWriteEnabled returns a boolean if a field has been set.
func (o *TokenProvisionRequest) HasWriteEnabled() bool {
	if o != nil && !IsNil(o.WriteEnabled) {
		return true
	}

	return false
}

// SetWriteEnabled gets a reference to the given bool and assigns it to the WriteEnabled field.
func (o *TokenProvisionRequest) SetWriteEnabled(v bool) {
	o.WriteEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TokenProvisionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProvisionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TokenProvisionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TokenProvisionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetUsername returns the Username field value
func (o *TokenProvisionRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TokenProvisionRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *TokenProvisionRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *TokenProvisionRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *TokenProvisionRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *TokenProvisionRequest) SetPassword(v string) {
	o.Password = v
}

func (o TokenProvisionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenProvisionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if !IsNil(o.WriteEnabled) {
		toSerialize["write_enabled"] = o.WriteEnabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenProvisionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
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
	varTokenProvisionRequest := _TokenProvisionRequest{}

	err = json.Unmarshal(data, &varTokenProvisionRequest)

	if err != nil {
		return err
	}

	*o = TokenProvisionRequest(varTokenProvisionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expires")
		delete(additionalProperties, "write_enabled")
		delete(additionalProperties, "description")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenProvisionRequest struct {
	value *TokenProvisionRequest
	isSet bool
}

func (v NullableTokenProvisionRequest) Get() *TokenProvisionRequest {
	return v.value
}

func (v *NullableTokenProvisionRequest) Set(val *TokenProvisionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenProvisionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenProvisionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenProvisionRequest(val *TokenProvisionRequest) *NullableTokenProvisionRequest {
	return &NullableTokenProvisionRequest{value: val, isSet: true}
}

func (v NullableTokenProvisionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenProvisionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
