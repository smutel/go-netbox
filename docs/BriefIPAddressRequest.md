# BriefIPAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**NatInside** | Pointer to [**NullableNestedIPAddressRequest**](NestedIPAddressRequest.md) |  | [optional] 
**DnsName** | Pointer to [**BriefIPAddressDnsName**](BriefIPAddressDnsName.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefIPAddressRequest

`func NewBriefIPAddressRequest(address string, ) *BriefIPAddressRequest`

NewBriefIPAddressRequest instantiates a new BriefIPAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefIPAddressRequestWithDefaults

`func NewBriefIPAddressRequestWithDefaults() *BriefIPAddressRequest`

NewBriefIPAddressRequestWithDefaults instantiates a new BriefIPAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BriefIPAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BriefIPAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BriefIPAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNatInside

`func (o *BriefIPAddressRequest) GetNatInside() NestedIPAddressRequest`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *BriefIPAddressRequest) GetNatInsideOk() (*NestedIPAddressRequest, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *BriefIPAddressRequest) SetNatInside(v NestedIPAddressRequest)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *BriefIPAddressRequest) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *BriefIPAddressRequest) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *BriefIPAddressRequest) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetDnsName

`func (o *BriefIPAddressRequest) GetDnsName() BriefIPAddressDnsName`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *BriefIPAddressRequest) GetDnsNameOk() (*BriefIPAddressDnsName, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *BriefIPAddressRequest) SetDnsName(v BriefIPAddressDnsName)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *BriefIPAddressRequest) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *BriefIPAddressRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefIPAddressRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefIPAddressRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefIPAddressRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


