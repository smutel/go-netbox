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

// checks if the CircuitCircuitTermination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitCircuitTermination{}

// CircuitCircuitTermination Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type CircuitCircuitTermination struct {
	Id              int32                        `json:"id"`
	Url             string                       `json:"url"`
	Display         string                       `json:"display"`
	Site            NullableBriefSite            `json:"site"`
	ProviderNetwork NullableBriefProviderNetwork `json:"provider_network"`
	// Physical circuit speed
	PortSpeed NullableInt32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt32 `json:"upstream_speed,omitempty"`
	// ID of the local cross-connect
	XconnectId           *string `json:"xconnect_id,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CircuitCircuitTermination CircuitCircuitTermination

// NewCircuitCircuitTermination instantiates a new CircuitCircuitTermination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitCircuitTermination(id int32, url string, display string, site NullableBriefSite, providerNetwork NullableBriefProviderNetwork) *CircuitCircuitTermination {
	this := CircuitCircuitTermination{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Site = site
	this.ProviderNetwork = providerNetwork
	return &this
}

// NewCircuitCircuitTerminationWithDefaults instantiates a new CircuitCircuitTermination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitCircuitTerminationWithDefaults() *CircuitCircuitTermination {
	this := CircuitCircuitTermination{}
	return &this
}

// GetId returns the Id field value
func (o *CircuitCircuitTermination) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTermination) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CircuitCircuitTermination) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CircuitCircuitTermination) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTermination) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CircuitCircuitTermination) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *CircuitCircuitTermination) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTermination) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CircuitCircuitTermination) SetDisplay(v string) {
	o.Display = v
}

// GetSite returns the Site field value
// If the value is explicit nil, the zero value for BriefSite will be returned
func (o *CircuitCircuitTermination) GetSite() BriefSite {
	if o == nil || o.Site.Get() == nil {
		var ret BriefSite
		return ret
	}

	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTermination) GetSiteOk() (*BriefSite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// SetSite sets field value
func (o *CircuitCircuitTermination) SetSite(v BriefSite) {
	o.Site.Set(&v)
}

// GetProviderNetwork returns the ProviderNetwork field value
// If the value is explicit nil, the zero value for BriefProviderNetwork will be returned
func (o *CircuitCircuitTermination) GetProviderNetwork() BriefProviderNetwork {
	if o == nil || o.ProviderNetwork.Get() == nil {
		var ret BriefProviderNetwork
		return ret
	}

	return *o.ProviderNetwork.Get()
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTermination) GetProviderNetworkOk() (*BriefProviderNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderNetwork.Get(), o.ProviderNetwork.IsSet()
}

// SetProviderNetwork sets field value
func (o *CircuitCircuitTermination) SetProviderNetwork(v BriefProviderNetwork) {
	o.ProviderNetwork.Set(&v)
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitCircuitTermination) GetPortSpeed() int32 {
	if o == nil || IsNil(o.PortSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTermination) GetPortSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *CircuitCircuitTermination) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt32 and assigns it to the PortSpeed field.
func (o *CircuitCircuitTermination) SetPortSpeed(v int32) {
	o.PortSpeed.Set(&v)
}

// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *CircuitCircuitTermination) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *CircuitCircuitTermination) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitCircuitTermination) GetUpstreamSpeed() int32 {
	if o == nil || IsNil(o.UpstreamSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitCircuitTermination) GetUpstreamSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *CircuitCircuitTermination) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt32 and assigns it to the UpstreamSpeed field.
func (o *CircuitCircuitTermination) SetUpstreamSpeed(v int32) {
	o.UpstreamSpeed.Set(&v)
}

// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *CircuitCircuitTermination) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *CircuitCircuitTermination) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *CircuitCircuitTermination) GetXconnectId() string {
	if o == nil || IsNil(o.XconnectId) {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTermination) GetXconnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.XconnectId) {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *CircuitCircuitTermination) HasXconnectId() bool {
	if o != nil && !IsNil(o.XconnectId) {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *CircuitCircuitTermination) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CircuitCircuitTermination) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTermination) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CircuitCircuitTermination) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CircuitCircuitTermination) SetDescription(v string) {
	o.Description = &v
}

func (o CircuitCircuitTermination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitCircuitTermination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["site"] = o.Site.Get()
	toSerialize["provider_network"] = o.ProviderNetwork.Get()
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if !IsNil(o.XconnectId) {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CircuitCircuitTermination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"site",
		"provider_network",
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
	varCircuitCircuitTermination := _CircuitCircuitTermination{}

	err = json.Unmarshal(data, &varCircuitCircuitTermination)

	if err != nil {
		return err
	}

	*o = CircuitCircuitTermination(varCircuitCircuitTermination)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "site")
		delete(additionalProperties, "provider_network")
		delete(additionalProperties, "port_speed")
		delete(additionalProperties, "upstream_speed")
		delete(additionalProperties, "xconnect_id")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCircuitCircuitTermination struct {
	value *CircuitCircuitTermination
	isSet bool
}

func (v NullableCircuitCircuitTermination) Get() *CircuitCircuitTermination {
	return v.value
}

func (v *NullableCircuitCircuitTermination) Set(val *CircuitCircuitTermination) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitCircuitTermination) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitCircuitTermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitCircuitTermination(val *CircuitCircuitTermination) *NullableCircuitCircuitTermination {
	return &NullableCircuitCircuitTermination{value: val, isSet: true}
}

func (v NullableCircuitCircuitTermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitCircuitTermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
