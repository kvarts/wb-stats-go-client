# \StatsApi

All URIs are relative to *https://suppliers-stats.wildberries.ru/api/v1/supplier*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncomes**](StatsApi.md#GetIncomes) | **Get** /incomes | Get Incomes
[**GetOrders**](StatsApi.md#GetOrders) | **Get** /orders | Get Orders
[**GetReportDetailByPeriod**](StatsApi.md#GetReportDetailByPeriod) | **Get** /reportDetailByPeriod | Get Report Detail By Period
[**GetSales**](StatsApi.md#GetSales) | **Get** /sales | Get Sales
[**GetStocks**](StatsApi.md#GetStocks) | **Get** /stocks | Get Stocks



## GetIncomes

> []GetIncomes200ResponseInner GetIncomes(ctx).DateFrom(dateFrom).Execute()

Get Incomes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dateFrom := time.Now() // time.Time | Дата и время от которых выгружается информация

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetIncomes(context.Background()).DateFrom(dateFrom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetIncomes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncomes`: []GetIncomes200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetIncomes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **time.Time** | Дата и время от которых выгружается информация | 

### Return type

[**[]GetIncomes200ResponseInner**](GetIncomes200ResponseInner.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> []GetOrders200ResponseInner GetOrders(ctx).DateFrom(dateFrom).Flag(flag).Execute()

Get Orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dateFrom := time.Now() // time.Time | 
    flag := int32(56) // int32 | 1 - то за одну дату, 0 - за все что больше переданной даты (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetOrders(context.Background()).DateFrom(dateFrom).Flag(flag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: []GetOrders200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **time.Time** |  | 
 **flag** | **int32** | 1 - то за одну дату, 0 - за все что больше переданной даты | 

### Return type

[**[]GetOrders200ResponseInner**](GetOrders200ResponseInner.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportDetailByPeriod

> []GetReportDetailByPeriod200ResponseInner GetReportDetailByPeriod(ctx).DateFrom(dateFrom).DateTo(dateTo).Limit(limit).Rrdid(rrdid).Execute()

Get Report Detail By Period



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dateFrom := time.Now() // time.Time | Начальная дата периода
    dateTo := time.Now() // time.Time | Конечная дата периода
    limit := int64(789) // int64 | Максимальное количество записей, получаемых при запросе
    rrdid := int64(789) // int64 | Идентификатор записи, начиная с которой нужно получать данные при запросе

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetReportDetailByPeriod(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Limit(limit).Rrdid(rrdid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetReportDetailByPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportDetailByPeriod`: []GetReportDetailByPeriod200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetReportDetailByPeriod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportDetailByPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **time.Time** | Начальная дата периода | 
 **dateTo** | **time.Time** | Конечная дата периода | 
 **limit** | **int64** | Максимальное количество записей, получаемых при запросе | 
 **rrdid** | **int64** | Идентификатор записи, начиная с которой нужно получать данные при запросе | 

### Return type

[**[]GetReportDetailByPeriod200ResponseInner**](GetReportDetailByPeriod200ResponseInner.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSales

> []GetSales200ResponseInner GetSales(ctx).DateFrom(dateFrom).Flag(flag).Execute()

Get Sales



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dateFrom := time.Now() // time.Time | Дата и время от которых выгружается информация
    flag := int32(56) // int32 | 1 - то за одну дату, 0 - за все что больше переданной даты (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetSales(context.Background()).DateFrom(dateFrom).Flag(flag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetSales``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSales`: []GetSales200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetSales`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **time.Time** | Дата и время от которых выгружается информация | 
 **flag** | **int32** | 1 - то за одну дату, 0 - за все что больше переданной даты | 

### Return type

[**[]GetSales200ResponseInner**](GetSales200ResponseInner.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStocks

> []GetStocks200ResponseInner GetStocks(ctx).DateFrom(dateFrom).Execute()

Get Stocks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dateFrom := time.Now() // time.Time | Дата и время от которых выгружается информация

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetStocks(context.Background()).DateFrom(dateFrom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetStocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStocks`: []GetStocks200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetStocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **time.Time** | Дата и время от которых выгружается информация | 

### Return type

[**[]GetStocks200ResponseInner**](GetStocks200ResponseInner.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

