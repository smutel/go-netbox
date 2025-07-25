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

// checks if the WritableSiteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableSiteRequest{}

// WritableSiteRequest Adds support for custom fields and tags.
type WritableSiteRequest struct {
	// Full name of the site
	Name   string                        `json:"name"`
	Slug   string                        `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Status *LocationStatusValue          `json:"status,omitempty"`
	Region NullableBriefRegionRequest    `json:"region,omitempty"`
	Group  NullableBriefSiteGroupRequest `json:"group,omitempty"`
	Tenant NullableBriefTenantRequest    `json:"tenant,omitempty"`
	// Local facility ID or description
	Facility    *string        `json:"facility,omitempty"`
	TimeZone    NullableString `json:"time_zone,omitempty"`
	Description *string        `json:"description,omitempty"`
	// Physical location of the building
	PhysicalAddress *string `json:"physical_address,omitempty"`
	// If different from the physical address
	ShippingAddress *string `json:"shipping_address,omitempty"`
	// GPS coordinate in decimal format (xx.yyyyyy)
	Latitude NullableFloat64 `json:"latitude,omitempty"`
	// GPS coordinate in decimal format (xx.yyyyyy)
	Longitude            NullableFloat64        `json:"longitude,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Asns                 []int32                `json:"asns,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableSiteRequest WritableSiteRequest

// NewWritableSiteRequest instantiates a new WritableSiteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableSiteRequest(name string, slug string) *WritableSiteRequest {
	this := WritableSiteRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewWritableSiteRequestWithDefaults instantiates a new WritableSiteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableSiteRequestWithDefaults() *WritableSiteRequest {
	this := WritableSiteRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableSiteRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableSiteRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *WritableSiteRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WritableSiteRequest) SetSlug(v string) {
	o.Slug = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetStatus() LocationStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret LocationStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetStatusOk() (*LocationStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given LocationStatusValue and assigns it to the Status field.
func (o *WritableSiteRequest) SetStatus(v LocationStatusValue) {
	o.Status = &v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableSiteRequest) GetRegion() BriefRegionRequest {
	if o == nil || IsNil(o.Region.Get()) {
		var ret BriefRegionRequest
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableSiteRequest) GetRegionOk() (*BriefRegionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableBriefRegionRequest and assigns it to the Region field.
func (o *WritableSiteRequest) SetRegion(v BriefRegionRequest) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *WritableSiteRequest) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *WritableSiteRequest) UnsetRegion() {
	o.Region.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableSiteRequest) GetGroup() BriefSiteGroupRequest {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefSiteGroupRequest
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableSiteRequest) GetGroupOk() (*BriefSiteGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefSiteGroupRequest and assigns it to the Group field.
func (o *WritableSiteRequest) SetGroup(v BriefSiteGroupRequest) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *WritableSiteRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *WritableSiteRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableSiteRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableSiteRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *WritableSiteRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableSiteRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableSiteRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetFacility() string {
	if o == nil || IsNil(o.Facility) {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetFacilityOk() (*string, bool) {
	if o == nil || IsNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasFacility() bool {
	if o != nil && !IsNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *WritableSiteRequest) SetFacility(v string) {
	o.Facility = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableSiteRequest) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone.Get()) {
		var ret string
		return ret
	}
	return *o.TimeZone.Get()
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableSiteRequest) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZone.Get(), o.TimeZone.IsSet()
}

// HasTimeZone returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasTimeZone() bool {
	if o != nil && o.TimeZone.IsSet() {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given NullableString and assigns it to the TimeZone field.
func (o *WritableSiteRequest) SetTimeZone(v string) {
	o.TimeZone.Set(&v)
}

// SetTimeZoneNil sets the value for TimeZone to be an explicit nil
func (o *WritableSiteRequest) SetTimeZoneNil() {
	o.TimeZone.Set(nil)
}

// UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
func (o *WritableSiteRequest) UnsetTimeZone() {
	o.TimeZone.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableSiteRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPhysicalAddress returns the PhysicalAddress field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetPhysicalAddress() string {
	if o == nil || IsNil(o.PhysicalAddress) {
		var ret string
		return ret
	}
	return *o.PhysicalAddress
}

// GetPhysicalAddressOk returns a tuple with the PhysicalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetPhysicalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalAddress) {
		return nil, false
	}
	return o.PhysicalAddress, true
}

// HasPhysicalAddress returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasPhysicalAddress() bool {
	if o != nil && !IsNil(o.PhysicalAddress) {
		return true
	}

	return false
}

// SetPhysicalAddress gets a reference to the given string and assigns it to the PhysicalAddress field.
func (o *WritableSiteRequest) SetPhysicalAddress(v string) {
	o.PhysicalAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetShippingAddress() string {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret string
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetShippingAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given string and assigns it to the ShippingAddress field.
func (o *WritableSiteRequest) SetShippingAddress(v string) {
	o.ShippingAddress = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableSiteRequest) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableSiteRequest) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *WritableSiteRequest) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}

// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *WritableSiteRequest) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *WritableSiteRequest) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableSiteRequest) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableSiteRequest) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *WritableSiteRequest) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}

// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *WritableSiteRequest) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *WritableSiteRequest) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableSiteRequest) SetComments(v string) {
	o.Comments = &v
}

// GetAsns returns the Asns field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetAsns() []int32 {
	if o == nil || IsNil(o.Asns) {
		var ret []int32
		return ret
	}
	return o.Asns
}

// GetAsnsOk returns a tuple with the Asns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetAsnsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Asns) {
		return nil, false
	}
	return o.Asns, true
}

// HasAsns returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasAsns() bool {
	if o != nil && !IsNil(o.Asns) {
		return true
	}

	return false
}

// SetAsns gets a reference to the given []int32 and assigns it to the Asns field.
func (o *WritableSiteRequest) SetAsns(v []int32) {
	o.Asns = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableSiteRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableSiteRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableSiteRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableSiteRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableSiteRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableSiteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableSiteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if o.TimeZone.IsSet() {
		toSerialize["time_zone"] = o.TimeZone.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PhysicalAddress) {
		toSerialize["physical_address"] = o.PhysicalAddress
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Asns) {
		toSerialize["asns"] = o.Asns
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableSiteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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
	varWritableSiteRequest := _WritableSiteRequest{}

	err = json.Unmarshal(data, &varWritableSiteRequest)

	if err != nil {
		return err
	}

	*o = WritableSiteRequest(varWritableSiteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "status")
		delete(additionalProperties, "region")
		delete(additionalProperties, "group")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "facility")
		delete(additionalProperties, "time_zone")
		delete(additionalProperties, "description")
		delete(additionalProperties, "physical_address")
		delete(additionalProperties, "shipping_address")
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "asns")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableSiteRequest struct {
	value *WritableSiteRequest
	isSet bool
}

func (v NullableWritableSiteRequest) Get() *WritableSiteRequest {
	return v.value
}

func (v *NullableWritableSiteRequest) Set(val *WritableSiteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableSiteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableSiteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableSiteRequest(val *WritableSiteRequest) *NullableWritableSiteRequest {
	return &NullableWritableSiteRequest{value: val, isSet: true}
}

func (v NullableWritableSiteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableSiteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
