# BriefIPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Family** | [**AggregateFamily**](AggregateFamily.md) |  | 
**Address** | **string** |  | 
**NatInside** | Pointer to [**NullableNestedIPAddress**](NestedIPAddress.md) |  | [optional] 
**NatOutside** | [**[]NestedIPAddress**](NestedIPAddress.md) |  | [readonly] 
**DnsName** | Pointer to [**BriefIPAddressDnsName**](BriefIPAddressDnsName.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefIPAddress

`func NewBriefIPAddress(id int32, url string, display string, family AggregateFamily, address string, natOutside []NestedIPAddress, ) *BriefIPAddress`

NewBriefIPAddress instantiates a new BriefIPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefIPAddressWithDefaults

`func NewBriefIPAddressWithDefaults() *BriefIPAddress`

NewBriefIPAddressWithDefaults instantiates a new BriefIPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefIPAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefIPAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefIPAddress) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefIPAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefIPAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefIPAddress) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefIPAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefIPAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefIPAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetFamily

`func (o *BriefIPAddress) GetFamily() AggregateFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *BriefIPAddress) GetFamilyOk() (*AggregateFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *BriefIPAddress) SetFamily(v AggregateFamily)`

SetFamily sets Family field to given value.


### GetAddress

`func (o *BriefIPAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BriefIPAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BriefIPAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNatInside

`func (o *BriefIPAddress) GetNatInside() NestedIPAddress`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *BriefIPAddress) GetNatInsideOk() (*NestedIPAddress, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *BriefIPAddress) SetNatInside(v NestedIPAddress)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *BriefIPAddress) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *BriefIPAddress) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *BriefIPAddress) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetNatOutside

`func (o *BriefIPAddress) GetNatOutside() []NestedIPAddress`

GetNatOutside returns the NatOutside field if non-nil, zero value otherwise.

### GetNatOutsideOk

`func (o *BriefIPAddress) GetNatOutsideOk() (*[]NestedIPAddress, bool)`

GetNatOutsideOk returns a tuple with the NatOutside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOutside

`func (o *BriefIPAddress) SetNatOutside(v []NestedIPAddress)`

SetNatOutside sets NatOutside field to given value.


### GetDnsName

`func (o *BriefIPAddress) GetDnsName() BriefIPAddressDnsName`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *BriefIPAddress) GetDnsNameOk() (*BriefIPAddressDnsName, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *BriefIPAddress) SetDnsName(v BriefIPAddressDnsName)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *BriefIPAddress) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *BriefIPAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefIPAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefIPAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefIPAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


