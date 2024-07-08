/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	_context "context"
	"fmt"
	"io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// UserS3KeysApiService UserS3KeysApi service
type UserS3KeysApiService service

type ApiUmUsersS3keysDeleteRequest struct {
	ctx             _context.Context
	ApiService      *UserS3KeysApiService
	userId          string
	keyId           string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiUmUsersS3keysDeleteRequest) Pretty(pretty bool) ApiUmUsersS3keysDeleteRequest {
	r.pretty = &pretty
	return r
}
func (r ApiUmUsersS3keysDeleteRequest) Depth(depth int32) ApiUmUsersS3keysDeleteRequest {
	r.depth = &depth
	return r
}
func (r ApiUmUsersS3keysDeleteRequest) XContractNumber(xContractNumber int32) ApiUmUsersS3keysDeleteRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiUmUsersS3keysDeleteRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UmUsersS3keysDeleteExecute(r)
}

/*
 * UmUsersS3keysDelete Delete S3 keys
 * Delete the specified user S3 key.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId The unique ID of the user.
 * @param keyId The unique ID of the S3 key.
 * @return ApiUmUsersS3keysDeleteRequest
 */
func (a *UserS3KeysApiService) UmUsersS3keysDelete(ctx _context.Context, userId string, keyId string) ApiUmUsersS3keysDeleteRequest {
	return ApiUmUsersS3keysDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		keyId:      keyId,
	}
}

/*
 * Execute executes the request
 */
func (a *UserS3KeysApiService) UmUsersS3keysDeleteExecute(r ApiUmUsersS3keysDeleteRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserS3KeysApiService.UmUsersS3keysDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/um/users/{userId}/s3keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", _neturl.PathEscape(parameterToString(r.keyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "UmUsersS3keysDelete",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarAPIResponse, newErr
	}

	return localVarAPIResponse, nil
}

type ApiUmUsersS3keysFindByKeyIdRequest struct {
	ctx             _context.Context
	ApiService      *UserS3KeysApiService
	userId          string
	keyId           string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiUmUsersS3keysFindByKeyIdRequest) Pretty(pretty bool) ApiUmUsersS3keysFindByKeyIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiUmUsersS3keysFindByKeyIdRequest) Depth(depth int32) ApiUmUsersS3keysFindByKeyIdRequest {
	r.depth = &depth
	return r
}
func (r ApiUmUsersS3keysFindByKeyIdRequest) XContractNumber(xContractNumber int32) ApiUmUsersS3keysFindByKeyIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiUmUsersS3keysFindByKeyIdRequest) Execute() (S3Key, *APIResponse, error) {
	return r.ApiService.UmUsersS3keysFindByKeyIdExecute(r)
}

/*
 * UmUsersS3keysFindByKeyId Retrieve user S3 keys by key ID
 * Retrieve the specified user S3 key. The user ID is in the response body when the user is created, and in the list of the users, returned by GET. The key ID is in the response body when the S3 key is created, and in the list of all user S3 keys, returned by GET.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId The unique ID of the user.
 * @param keyId The unique ID of the S3 key.
 * @return ApiUmUsersS3keysFindByKeyIdRequest
 */
func (a *UserS3KeysApiService) UmUsersS3keysFindByKeyId(ctx _context.Context, userId string, keyId string) ApiUmUsersS3keysFindByKeyIdRequest {
	return ApiUmUsersS3keysFindByKeyIdRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		keyId:      keyId,
	}
}

/*
 * Execute executes the request
 * @return S3Key
 */
func (a *UserS3KeysApiService) UmUsersS3keysFindByKeyIdExecute(r ApiUmUsersS3keysFindByKeyIdRequest) (S3Key, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  S3Key
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserS3KeysApiService.UmUsersS3keysFindByKeyId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/um/users/{userId}/s3keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", _neturl.PathEscape(parameterToString(r.keyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "UmUsersS3keysFindByKeyId",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiUmUsersS3keysGetRequest struct {
	ctx             _context.Context
	ApiService      *UserS3KeysApiService
	filters         _neturl.Values
	orderBy         *string
	maxResults      *int32
	userId          string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiUmUsersS3keysGetRequest) Pretty(pretty bool) ApiUmUsersS3keysGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiUmUsersS3keysGetRequest) Depth(depth int32) ApiUmUsersS3keysGetRequest {
	r.depth = &depth
	return r
}
func (r ApiUmUsersS3keysGetRequest) XContractNumber(xContractNumber int32) ApiUmUsersS3keysGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiUmUsersS3keysGetRequest) Filter(key string, value string) ApiUmUsersS3keysGetRequest {
	filterKey := fmt.Sprintf(FilterQueryParam, key)
	r.filters[filterKey] = append(r.filters[filterKey], value)
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiUmUsersS3keysGetRequest) OrderBy(orderBy string) ApiUmUsersS3keysGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiUmUsersS3keysGetRequest) MaxResults(maxResults int32) ApiUmUsersS3keysGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiUmUsersS3keysGetRequest) Execute() (S3Keys, *APIResponse, error) {
	return r.ApiService.UmUsersS3keysGetExecute(r)
}

/*
 * UmUsersS3keysGet List user S3 keys
 * List S3 keys by user ID. The user ID is in the response body when the user is created, and in the list of the users, returned by GET.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId The unique ID of the user.
 * @return ApiUmUsersS3keysGetRequest
 */
func (a *UserS3KeysApiService) UmUsersS3keysGet(ctx _context.Context, userId string) ApiUmUsersS3keysGetRequest {
	return ApiUmUsersS3keysGetRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return S3Keys
 */
func (a *UserS3KeysApiService) UmUsersS3keysGetExecute(r ApiUmUsersS3keysGetRequest) (S3Keys, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  S3Keys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserS3KeysApiService.UmUsersS3keysGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/um/users/{userId}/s3keys"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("maxResults", parameterToString(*r.maxResults, ""))
	}
	if len(r.filters) > 0 {
		for k, v := range r.filters {
			for _, iv := range v {
				localVarQueryParams.Add(k, iv)
			}
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "UmUsersS3keysGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiUmUsersS3keysPostRequest struct {
	ctx             _context.Context
	ApiService      *UserS3KeysApiService
	userId          string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiUmUsersS3keysPostRequest) Pretty(pretty bool) ApiUmUsersS3keysPostRequest {
	r.pretty = &pretty
	return r
}
func (r ApiUmUsersS3keysPostRequest) Depth(depth int32) ApiUmUsersS3keysPostRequest {
	r.depth = &depth
	return r
}
func (r ApiUmUsersS3keysPostRequest) XContractNumber(xContractNumber int32) ApiUmUsersS3keysPostRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiUmUsersS3keysPostRequest) Execute() (S3Key, *APIResponse, error) {
	return r.ApiService.UmUsersS3keysPostExecute(r)
}

/*
 * UmUsersS3keysPost Create user S3 keys
 * Create an S3 key for the specified user. The user ID is in the response body when the user is created, and in the list of the users, returned by GET. A maximum of five keys per user can be generated.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId The unique ID of the user.
 * @return ApiUmUsersS3keysPostRequest
 */
func (a *UserS3KeysApiService) UmUsersS3keysPost(ctx _context.Context, userId string) ApiUmUsersS3keysPostRequest {
	return ApiUmUsersS3keysPostRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

/*
 * Execute executes the request
 * @return S3Key
 */
func (a *UserS3KeysApiService) UmUsersS3keysPostExecute(r ApiUmUsersS3keysPostRequest) (S3Key, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  S3Key
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserS3KeysApiService.UmUsersS3keysPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/um/users/{userId}/s3keys"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "UmUsersS3keysPost",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiUmUsersS3keysPutRequest struct {
	ctx             _context.Context
	ApiService      *UserS3KeysApiService
	userId          string
	keyId           string
	s3Key           *S3Key
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiUmUsersS3keysPutRequest) S3Key(s3Key S3Key) ApiUmUsersS3keysPutRequest {
	r.s3Key = &s3Key
	return r
}
func (r ApiUmUsersS3keysPutRequest) Pretty(pretty bool) ApiUmUsersS3keysPutRequest {
	r.pretty = &pretty
	return r
}
func (r ApiUmUsersS3keysPutRequest) Depth(depth int32) ApiUmUsersS3keysPutRequest {
	r.depth = &depth
	return r
}
func (r ApiUmUsersS3keysPutRequest) XContractNumber(xContractNumber int32) ApiUmUsersS3keysPutRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiUmUsersS3keysPutRequest) Execute() (S3Key, *APIResponse, error) {
	return r.ApiService.UmUsersS3keysPutExecute(r)
}

/*
 * UmUsersS3keysPut Modify a S3 Key by Key ID
 * Enables or disables the specified user S3 key.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId The unique ID of the user.
 * @param keyId The unique ID of the S3 key.
 * @return ApiUmUsersS3keysPutRequest
 */
func (a *UserS3KeysApiService) UmUsersS3keysPut(ctx _context.Context, userId string, keyId string) ApiUmUsersS3keysPutRequest {
	return ApiUmUsersS3keysPutRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		keyId:      keyId,
	}
}

/*
 * Execute executes the request
 * @return S3Key
 */
func (a *UserS3KeysApiService) UmUsersS3keysPutExecute(r ApiUmUsersS3keysPutRequest) (S3Key, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  S3Key
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserS3KeysApiService.UmUsersS3keysPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/um/users/{userId}/s3keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", _neturl.PathEscape(parameterToString(r.keyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.s3Key == nil {
		return localVarReturnValue, nil, reportError("s3Key is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("depth")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("depth", parameterToString(0, ""))
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.s3Key
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "UmUsersS3keysPut",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiUmUsersS3ssourlGetRequest struct {
	ctx             _context.Context
	ApiService      *UserS3KeysApiService
	filters         _neturl.Values
	orderBy         *string
	maxResults      *int32
	userId          string
	pretty          *bool
	xContractNumber *int32
}

func (r ApiUmUsersS3ssourlGetRequest) Pretty(pretty bool) ApiUmUsersS3ssourlGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiUmUsersS3ssourlGetRequest) XContractNumber(xContractNumber int32) ApiUmUsersS3ssourlGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiUmUsersS3ssourlGetRequest) Filter(key string, value string) ApiUmUsersS3ssourlGetRequest {
	filterKey := fmt.Sprintf(FilterQueryParam, key)
	r.filters[filterKey] = append(r.filters[filterKey], value)
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiUmUsersS3ssourlGetRequest) OrderBy(orderBy string) ApiUmUsersS3ssourlGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiUmUsersS3ssourlGetRequest) MaxResults(maxResults int32) ApiUmUsersS3ssourlGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiUmUsersS3ssourlGetRequest) Execute() (S3ObjectStorageSSO, *APIResponse, error) {
	return r.ApiService.UmUsersS3ssourlGetExecute(r)
}

/*
 * UmUsersS3ssourlGet Retrieve S3 single sign-on URLs
 * Retrieve S3 Object Storage single sign-on URLs for the the specified user. The user ID is in the response body when the user is created, and in the list of the users, returned by GET.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId The unique ID of the user.
 * @return ApiUmUsersS3ssourlGetRequest
 */
func (a *UserS3KeysApiService) UmUsersS3ssourlGet(ctx _context.Context, userId string) ApiUmUsersS3ssourlGetRequest {
	return ApiUmUsersS3ssourlGetRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return S3ObjectStorageSSO
 */
func (a *UserS3KeysApiService) UmUsersS3ssourlGetExecute(r ApiUmUsersS3ssourlGetRequest) (S3ObjectStorageSSO, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  S3ObjectStorageSSO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserS3KeysApiService.UmUsersS3ssourlGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/um/users/{userId}/s3ssourl"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	} else {
		defaultQueryParam := a.client.cfg.DefaultQueryParams.Get("pretty")
		if defaultQueryParam == "" {
			localVarQueryParams.Add("pretty", parameterToString(true, ""))
		}
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("maxResults", parameterToString(*r.maxResults, ""))
	}
	if len(r.filters) > 0 {
		for k, v := range r.filters {
			for _, iv := range v {
				localVarQueryParams.Add(k, iv)
			}
		}
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestURL:  localVarPath,
		RequestTime: httpRequestTime,
		Operation:   "UmUsersS3ssourlGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, string(localVarBody)),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = fmt.Sprintf(FormatStringErr, localVarHTTPResponse.Status, err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}