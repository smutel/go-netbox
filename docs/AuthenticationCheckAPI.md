# \AuthenticationCheckAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationCheckRetrieve**](AuthenticationCheckAPI.md#AuthenticationCheckRetrieve) | **Get** /api/authentication-check/ | 



## AuthenticationCheckRetrieve

> map[string]interface{} AuthenticationCheckRetrieve(ctx).Brief(brief).Fields(fields).Omit(omit).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smutel/go-netbox/v4"
)

func main() {
	brief := true // bool | Return only brief fields for each object. (optional)
	fields := "fields_example" // string | Comma-separated list of fields to include in the response. Example: `fields=id,name`. (optional)
	omit := "omit_example" // string | Comma-separated list of fields to exclude from the response. Example: `omit=description,tags`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationCheckAPI.AuthenticationCheckRetrieve(context.Background()).Brief(brief).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationCheckAPI.AuthenticationCheckRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationCheckRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationCheckAPI.AuthenticationCheckRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationCheckRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brief** | **bool** | Return only brief fields for each object. | 
 **fields** | **string** | Comma-separated list of fields to include in the response. Example: &#x60;fields&#x3D;id,name&#x60;. | 
 **omit** | **string** | Comma-separated list of fields to exclude from the response. Example: &#x60;omit&#x3D;description,tags&#x60;. | 

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

