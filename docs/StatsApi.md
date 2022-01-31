# \StatsApi

All URIs are relative to *https://suppliers-stats.wildberries.ru/api/v1/supplier*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncomes**](StatsApi.md#GetIncomes) | **Get** /incomes | Get Incomes
[**GetOrders**](StatsApi.md#GetOrders) | **Get** /orders | Get Orders
[**GetPaidStorage**](StatsApi.md#GetPaidStorage) | **Get** /stochrancost | Get Paid Storage
[**GetReportDetailByPeriod**](StatsApi.md#GetReportDetailByPeriod) | **Get** /reportDetailByPeriod | Ger Report Detail By Period
[**GetSales**](StatsApi.md#GetSales) | **Get** /sales | Get Sales
[**GetStocks**](StatsApi.md#GetStocks) | **Get** /stocks | Get Stocks



## GetIncomes

> []InlineResponse2001 GetIncomes(ctx).DateFrom(dateFrom).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.GetIncomes(context.Background()).DateFrom(dateFrom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetIncomes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncomes`: []InlineResponse2001
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

[**[]InlineResponse2001**](InlineResponse2001.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> []InlineResponse200 GetOrders(ctx).DateFrom(dateFrom).Flag(flag).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.GetOrders(context.Background()).DateFrom(dateFrom).Flag(flag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: []InlineResponse200
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

[**[]InlineResponse200**](InlineResponse200.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaidStorage

> []InlineResponse2004 GetPaidStorage(ctx).DateFrom(dateFrom).DateTo(dateTo).Execute()

Get Paid Storage



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
    dateTo := time.Now() // time.Time | Дата и время до которых выгружается информация

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.GetPaidStorage(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetPaidStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaidStorage`: []InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetPaidStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaidStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **time.Time** | Дата и время от которых выгружается информация | 
 **dateTo** | **time.Time** | Дата и время до которых выгружается информация | 

### Return type

[**[]InlineResponse2004**](InlineResponse2004.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportDetailByPeriod

> []InlineResponse2005 GetReportDetailByPeriod(ctx).DateFrom(dateFrom).DateTo(dateTo).Limit(limit).Rrdid(rrdid).Execute()

Ger Report Detail By Period



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dateFrom := "dateFrom_example" // string | Начальная дата периода (optional)
    dateTo := "dateTo_example" // string | Конечная дата периода (optional)
    limit := int32(56) // int32 | Максимальное количество записей, получаемых при запросе (optional)
    rrdid := int32(56) // int32 | Идентификатор записи, начиная с которой нужно получать данные при запросе (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.GetReportDetailByPeriod(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Limit(limit).Rrdid(rrdid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetReportDetailByPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportDetailByPeriod`: []InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetReportDetailByPeriod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportDetailByPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Начальная дата периода | 
 **dateTo** | **string** | Конечная дата периода | 
 **limit** | **int32** | Максимальное количество записей, получаемых при запросе | 
 **rrdid** | **int32** | Идентификатор записи, начиная с которой нужно получать данные при запросе | 

### Return type

[**[]InlineResponse2005**](InlineResponse2005.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSales

> []InlineResponse2003 GetSales(ctx).DateFrom(dateFrom).Flag(flag).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.GetSales(context.Background()).DateFrom(dateFrom).Flag(flag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetSales``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSales`: []InlineResponse2003
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

[**[]InlineResponse2003**](InlineResponse2003.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStocks

> []InlineResponse2002 GetStocks(ctx).DateFrom(dateFrom).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.GetStocks(context.Background()).DateFrom(dateFrom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetStocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStocks`: []InlineResponse2002
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

[**[]InlineResponse2002**](InlineResponse2002.md)

### Authorization

[Key](../README.md#Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

