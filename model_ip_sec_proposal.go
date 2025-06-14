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

// checks if the IPSecProposal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPSecProposal{}

// IPSecProposal Adds support for custom fields and tags.
type IPSecProposal struct {
	Id                      int32                              `json:"id"`
	Url                     string                             `json:"url"`
	Display                 string                             `json:"display"`
	Name                    string                             `json:"name"`
	Description             *string                            `json:"description,omitempty"`
	EncryptionAlgorithm     IKEProposalEncryptionAlgorithm     `json:"encryption_algorithm"`
	AuthenticationAlgorithm IKEProposalAuthenticationAlgorithm `json:"authentication_algorithm"`
	// Security association lifetime (seconds)
	SaLifetimeSeconds NullableInt32 `json:"sa_lifetime_seconds,omitempty"`
	// Security association lifetime (in kilobytes)
	SaLifetimeData       NullableInt32          `json:"sa_lifetime_data,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _IPSecProposal IPSecProposal

// NewIPSecProposal instantiates a new IPSecProposal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecProposal(id int32, url string, display string, name string, encryptionAlgorithm IKEProposalEncryptionAlgorithm, authenticationAlgorithm IKEProposalAuthenticationAlgorithm, created NullableTime, lastUpdated NullableTime) *IPSecProposal {
	this := IPSecProposal{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.EncryptionAlgorithm = encryptionAlgorithm
	this.AuthenticationAlgorithm = authenticationAlgorithm
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewIPSecProposalWithDefaults instantiates a new IPSecProposal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecProposalWithDefaults() *IPSecProposal {
	this := IPSecProposal{}
	return &this
}

// GetId returns the Id field value
func (o *IPSecProposal) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IPSecProposal) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *IPSecProposal) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IPSecProposal) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *IPSecProposal) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *IPSecProposal) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *IPSecProposal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IPSecProposal) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPSecProposal) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPSecProposal) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPSecProposal) SetDescription(v string) {
	o.Description = &v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value
func (o *IPSecProposal) GetEncryptionAlgorithm() IKEProposalEncryptionAlgorithm {
	if o == nil {
		var ret IKEProposalEncryptionAlgorithm
		return ret
	}

	return o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetEncryptionAlgorithmOk() (*IKEProposalEncryptionAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionAlgorithm, true
}

// SetEncryptionAlgorithm sets field value
func (o *IPSecProposal) SetEncryptionAlgorithm(v IKEProposalEncryptionAlgorithm) {
	o.EncryptionAlgorithm = v
}

// GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field value
func (o *IPSecProposal) GetAuthenticationAlgorithm() IKEProposalAuthenticationAlgorithm {
	if o == nil {
		var ret IKEProposalAuthenticationAlgorithm
		return ret
	}

	return o.AuthenticationAlgorithm
}

// GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field value
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetAuthenticationAlgorithmOk() (*IKEProposalAuthenticationAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationAlgorithm, true
}

// SetAuthenticationAlgorithm sets field value
func (o *IPSecProposal) SetAuthenticationAlgorithm(v IKEProposalAuthenticationAlgorithm) {
	o.AuthenticationAlgorithm = v
}

// GetSaLifetimeSeconds returns the SaLifetimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPSecProposal) GetSaLifetimeSeconds() int32 {
	if o == nil || IsNil(o.SaLifetimeSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetimeSeconds.Get()
}

// GetSaLifetimeSecondsOk returns a tuple with the SaLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecProposal) GetSaLifetimeSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetimeSeconds.Get(), o.SaLifetimeSeconds.IsSet()
}

// HasSaLifetimeSeconds returns a boolean if a field has been set.
func (o *IPSecProposal) HasSaLifetimeSeconds() bool {
	if o != nil && o.SaLifetimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetSaLifetimeSeconds gets a reference to the given NullableInt32 and assigns it to the SaLifetimeSeconds field.
func (o *IPSecProposal) SetSaLifetimeSeconds(v int32) {
	o.SaLifetimeSeconds.Set(&v)
}

// SetSaLifetimeSecondsNil sets the value for SaLifetimeSeconds to be an explicit nil
func (o *IPSecProposal) SetSaLifetimeSecondsNil() {
	o.SaLifetimeSeconds.Set(nil)
}

// UnsetSaLifetimeSeconds ensures that no value is present for SaLifetimeSeconds, not even an explicit nil
func (o *IPSecProposal) UnsetSaLifetimeSeconds() {
	o.SaLifetimeSeconds.Unset()
}

// GetSaLifetimeData returns the SaLifetimeData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPSecProposal) GetSaLifetimeData() int32 {
	if o == nil || IsNil(o.SaLifetimeData.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetimeData.Get()
}

// GetSaLifetimeDataOk returns a tuple with the SaLifetimeData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecProposal) GetSaLifetimeDataOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetimeData.Get(), o.SaLifetimeData.IsSet()
}

// HasSaLifetimeData returns a boolean if a field has been set.
func (o *IPSecProposal) HasSaLifetimeData() bool {
	if o != nil && o.SaLifetimeData.IsSet() {
		return true
	}

	return false
}

// SetSaLifetimeData gets a reference to the given NullableInt32 and assigns it to the SaLifetimeData field.
func (o *IPSecProposal) SetSaLifetimeData(v int32) {
	o.SaLifetimeData.Set(&v)
}

// SetSaLifetimeDataNil sets the value for SaLifetimeData to be an explicit nil
func (o *IPSecProposal) SetSaLifetimeDataNil() {
	o.SaLifetimeData.Set(nil)
}

// UnsetSaLifetimeData ensures that no value is present for SaLifetimeData, not even an explicit nil
func (o *IPSecProposal) UnsetSaLifetimeData() {
	o.SaLifetimeData.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPSecProposal) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPSecProposal) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPSecProposal) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPSecProposal) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPSecProposal) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *IPSecProposal) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPSecProposal) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProposal) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPSecProposal) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IPSecProposal) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IPSecProposal) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecProposal) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *IPSecProposal) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IPSecProposal) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecProposal) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *IPSecProposal) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o IPSecProposal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPSecProposal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["encryption_algorithm"] = o.EncryptionAlgorithm
	toSerialize["authentication_algorithm"] = o.AuthenticationAlgorithm
	if o.SaLifetimeSeconds.IsSet() {
		toSerialize["sa_lifetime_seconds"] = o.SaLifetimeSeconds.Get()
	}
	if o.SaLifetimeData.IsSet() {
		toSerialize["sa_lifetime_data"] = o.SaLifetimeData.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPSecProposal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"encryption_algorithm",
		"authentication_algorithm",
		"created",
		"last_updated",
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
	varIPSecProposal := _IPSecProposal{}

	err = json.Unmarshal(data, &varIPSecProposal)

	if err != nil {
		return err
	}

	*o = IPSecProposal(varIPSecProposal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "encryption_algorithm")
		delete(additionalProperties, "authentication_algorithm")
		delete(additionalProperties, "sa_lifetime_seconds")
		delete(additionalProperties, "sa_lifetime_data")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPSecProposal struct {
	value *IPSecProposal
	isSet bool
}

func (v NullableIPSecProposal) Get() *IPSecProposal {
	return v.value
}

func (v *NullableIPSecProposal) Set(val *IPSecProposal) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecProposal) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecProposal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecProposal(val *IPSecProposal) *NullableIPSecProposal {
	return &NullableIPSecProposal{value: val, isSet: true}
}

func (v NullableIPSecProposal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecProposal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
