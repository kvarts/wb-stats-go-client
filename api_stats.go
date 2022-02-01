/*
WB Supplier Stats

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// StatsApiService StatsApi service
type StatsApiService service

type ApiGetIncomesRequest struct {
	ctx _context.Context
	ApiService *StatsApiService
	dateFrom *time.Time
}

// Дата и время от которых выгружается информация
func (r ApiGetIncomesRequest) DateFrom(dateFrom time.Time) ApiGetIncomesRequest {
	r.dateFrom = &dateFrom
	return r
}

func (r ApiGetIncomesRequest) Execute() ([]InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.GetIncomesExecute(r)
}

/*
GetIncomes Get Incomes

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIncomesRequest
*/
func (a *StatsApiService) GetIncomes(ctx _context.Context) ApiGetIncomesRequest {
	return ApiGetIncomesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InlineResponse2001
func (a *StatsApiService) GetIncomesExecute(r ApiGetIncomesRequest) ([]InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsApiService.GetIncomes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/incomes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.dateFrom == nil {
		return localVarReturnValue, nil, reportError("dateFrom is required and must be specified")
	}

	localVarQueryParams.Add("dateFrom", parameterToString(*r.dateFrom, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrdersRequest struct {
	ctx _context.Context
	ApiService *StatsApiService
	dateFrom *time.Time
	flag *int32
}

func (r ApiGetOrdersRequest) DateFrom(dateFrom time.Time) ApiGetOrdersRequest {
	r.dateFrom = &dateFrom
	return r
}
// 1 - то за одну дату, 0 - за все что больше переданной даты
func (r ApiGetOrdersRequest) Flag(flag int32) ApiGetOrdersRequest {
	r.flag = &flag
	return r
}

func (r ApiGetOrdersRequest) Execute() ([]InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.GetOrdersExecute(r)
}

/*
GetOrders Get Orders

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOrdersRequest
*/
func (a *StatsApiService) GetOrders(ctx _context.Context) ApiGetOrdersRequest {
	return ApiGetOrdersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InlineResponse200
func (a *StatsApiService) GetOrdersExecute(r ApiGetOrdersRequest) ([]InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsApiService.GetOrders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.dateFrom == nil {
		return localVarReturnValue, nil, reportError("dateFrom is required and must be specified")
	}

	if r.flag != nil {
		localVarQueryParams.Add("flag", parameterToString(*r.flag, ""))
	}
	localVarQueryParams.Add("dateFrom", parameterToString(*r.dateFrom, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaidStorageRequest struct {
	ctx _context.Context
	ApiService *StatsApiService
	dateFrom *time.Time
	dateTo *time.Time
}

// Дата и время от которых выгружается информация
func (r ApiGetPaidStorageRequest) DateFrom(dateFrom time.Time) ApiGetPaidStorageRequest {
	r.dateFrom = &dateFrom
	return r
}
// Дата и время до которых выгружается информация
func (r ApiGetPaidStorageRequest) DateTo(dateTo time.Time) ApiGetPaidStorageRequest {
	r.dateTo = &dateTo
	return r
}

func (r ApiGetPaidStorageRequest) Execute() ([]InlineResponse2004, *_nethttp.Response, error) {
	return r.ApiService.GetPaidStorageExecute(r)
}

/*
GetPaidStorage Get Paid Storage

Получить информацию о платном хранилище - факт

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPaidStorageRequest
*/
func (a *StatsApiService) GetPaidStorage(ctx _context.Context) ApiGetPaidStorageRequest {
	return ApiGetPaidStorageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InlineResponse2004
func (a *StatsApiService) GetPaidStorageExecute(r ApiGetPaidStorageRequest) ([]InlineResponse2004, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2004
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsApiService.GetPaidStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stochrancost"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.dateFrom == nil {
		return localVarReturnValue, nil, reportError("dateFrom is required and must be specified")
	}
	if r.dateTo == nil {
		return localVarReturnValue, nil, reportError("dateTo is required and must be specified")
	}

	localVarQueryParams.Add("dateFrom", parameterToString(*r.dateFrom, ""))
	localVarQueryParams.Add("dateTo", parameterToString(*r.dateTo, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetReportDetailByPeriodRequest struct {
	ctx _context.Context
	ApiService *StatsApiService
	dateFrom *time.Time
	dateTo *time.Time
	limit *int64
	rrdid *int64
}

// Начальная дата периода
func (r ApiGetReportDetailByPeriodRequest) DateFrom(dateFrom time.Time) ApiGetReportDetailByPeriodRequest {
	r.dateFrom = &dateFrom
	return r
}
// Конечная дата периода
func (r ApiGetReportDetailByPeriodRequest) DateTo(dateTo time.Time) ApiGetReportDetailByPeriodRequest {
	r.dateTo = &dateTo
	return r
}
// Максимальное количество записей, получаемых при запросе
func (r ApiGetReportDetailByPeriodRequest) Limit(limit int64) ApiGetReportDetailByPeriodRequest {
	r.limit = &limit
	return r
}
// Идентификатор записи, начиная с которой нужно получать данные при запросе
func (r ApiGetReportDetailByPeriodRequest) Rrdid(rrdid int64) ApiGetReportDetailByPeriodRequest {
	r.rrdid = &rrdid
	return r
}

func (r ApiGetReportDetailByPeriodRequest) Execute() ([]InlineResponse2005, *_nethttp.Response, error) {
	return r.ApiService.GetReportDetailByPeriodExecute(r)
}

/*
GetReportDetailByPeriod Get Report Detail By Period

Получить информацию об отчёте о продажах по реализации

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetReportDetailByPeriodRequest
*/
func (a *StatsApiService) GetReportDetailByPeriod(ctx _context.Context) ApiGetReportDetailByPeriodRequest {
	return ApiGetReportDetailByPeriodRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InlineResponse2005
func (a *StatsApiService) GetReportDetailByPeriodExecute(r ApiGetReportDetailByPeriodRequest) ([]InlineResponse2005, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2005
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsApiService.GetReportDetailByPeriod")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reportDetailByPeriod"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.dateFrom == nil {
		return localVarReturnValue, nil, reportError("dateFrom is required and must be specified")
	}
	if r.dateTo == nil {
		return localVarReturnValue, nil, reportError("dateTo is required and must be specified")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if *r.limit < 0 {
		return localVarReturnValue, nil, reportError("limit must be greater than 0")
	}
	if r.rrdid == nil {
		return localVarReturnValue, nil, reportError("rrdid is required and must be specified")
	}
	if *r.rrdid < 0 {
		return localVarReturnValue, nil, reportError("rrdid must be greater than 0")
	}

	localVarQueryParams.Add("dateFrom", parameterToString(*r.dateFrom, ""))
	localVarQueryParams.Add("dateTo", parameterToString(*r.dateTo, ""))
	localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	localVarQueryParams.Add("rrdid", parameterToString(*r.rrdid, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSalesRequest struct {
	ctx _context.Context
	ApiService *StatsApiService
	dateFrom *time.Time
	flag *int32
}

// Дата и время от которых выгружается информация
func (r ApiGetSalesRequest) DateFrom(dateFrom time.Time) ApiGetSalesRequest {
	r.dateFrom = &dateFrom
	return r
}
// 1 - то за одну дату, 0 - за все что больше переданной даты
func (r ApiGetSalesRequest) Flag(flag int32) ApiGetSalesRequest {
	r.flag = &flag
	return r
}

func (r ApiGetSalesRequest) Execute() ([]InlineResponse2003, *_nethttp.Response, error) {
	return r.ApiService.GetSalesExecute(r)
}

/*
GetSales Get Sales

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSalesRequest
*/
func (a *StatsApiService) GetSales(ctx _context.Context) ApiGetSalesRequest {
	return ApiGetSalesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InlineResponse2003
func (a *StatsApiService) GetSalesExecute(r ApiGetSalesRequest) ([]InlineResponse2003, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2003
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsApiService.GetSales")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.dateFrom == nil {
		return localVarReturnValue, nil, reportError("dateFrom is required and must be specified")
	}

	localVarQueryParams.Add("dateFrom", parameterToString(*r.dateFrom, ""))
	if r.flag != nil {
		localVarQueryParams.Add("flag", parameterToString(*r.flag, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetStocksRequest struct {
	ctx _context.Context
	ApiService *StatsApiService
	dateFrom *time.Time
}

// Дата и время от которых выгружается информация
func (r ApiGetStocksRequest) DateFrom(dateFrom time.Time) ApiGetStocksRequest {
	r.dateFrom = &dateFrom
	return r
}

func (r ApiGetStocksRequest) Execute() ([]InlineResponse2002, *_nethttp.Response, error) {
	return r.ApiService.GetStocksExecute(r)
}

/*
GetStocks Get Stocks

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStocksRequest
*/
func (a *StatsApiService) GetStocks(ctx _context.Context) ApiGetStocksRequest {
	return ApiGetStocksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InlineResponse2002
func (a *StatsApiService) GetStocksExecute(r ApiGetStocksRequest) ([]InlineResponse2002, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2002
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsApiService.GetStocks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stocks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.dateFrom == nil {
		return localVarReturnValue, nil, reportError("dateFrom is required and must be specified")
	}

	localVarQueryParams.Add("dateFrom", parameterToString(*r.dateFrom, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
