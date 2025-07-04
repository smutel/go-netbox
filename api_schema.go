/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type SchemaAPI interface {

	/*
			SchemaRetrieve Method for SchemaRetrieve

			OpenApi3 schema for this API. Format can be selected via content negotiation.

		- YAML: application/vnd.oai.openapi
		- JSON: application/vnd.oai.openapi+json

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiSchemaRetrieveRequest
	*/
	SchemaRetrieve(ctx context.Context) ApiSchemaRetrieveRequest

	// SchemaRetrieveExecute executes the request
	//  @return map[string]interface{}
	SchemaRetrieveExecute(r ApiSchemaRetrieveRequest) (map[string]interface{}, *http.Response, error)
}

// SchemaAPIService SchemaAPI service
type SchemaAPIService service

type ApiSchemaRetrieveRequest struct {
	ctx        context.Context
	ApiService SchemaAPI
	format     *SchemaRetrieveFormatParameter
	lang       *SchemaRetrieveLangParameter
}

func (r ApiSchemaRetrieveRequest) Format(format SchemaRetrieveFormatParameter) ApiSchemaRetrieveRequest {
	r.format = &format
	return r
}

func (r ApiSchemaRetrieveRequest) Lang(lang SchemaRetrieveLangParameter) ApiSchemaRetrieveRequest {
	r.lang = &lang
	return r
}

func (r ApiSchemaRetrieveRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SchemaRetrieveExecute(r)
}

/*
SchemaRetrieve Method for SchemaRetrieve

OpenApi3 schema for this API. Format can be selected via content negotiation.

- YAML: application/vnd.oai.openapi
- JSON: application/vnd.oai.openapi+json

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSchemaRetrieveRequest
*/
func (a *SchemaAPIService) SchemaRetrieve(ctx context.Context) ApiSchemaRetrieveRequest {
	return ApiSchemaRetrieveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *SchemaAPIService) SchemaRetrieveExecute(r ApiSchemaRetrieveRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemaAPIService.SchemaRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/schema/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.lang != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lang", r.lang, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.oai.openapi", "application/yaml", "application/vnd.oai.openapi+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
