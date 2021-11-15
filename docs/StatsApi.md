# \StatsApi

All URIs are relative to *https://suppliers-stats.wildberries.ru/api/v1/supplier*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrders**](StatsApi.md#GetOrders) | **Get** /orders | Get Orders



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

