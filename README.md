# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/kvarts/wb-stats-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://suppliers-stats.wildberries.ru/api/v1/supplier*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*StatsApi* | [**GetIncomes**](docs/StatsApi.md#getincomes) | **Get** /incomes | Get Incomes
*StatsApi* | [**GetOrders**](docs/StatsApi.md#getorders) | **Get** /orders | Get Orders
*StatsApi* | [**GetPaidStorage**](docs/StatsApi.md#getpaidstorage) | **Get** /stochrancost | Get Paid Storage
*StatsApi* | [**GetReportDetailByPeriod**](docs/StatsApi.md#getreportdetailbyperiod) | **Get** /reportDetailByPeriod | Get Report Detail By Period
*StatsApi* | [**GetSales**](docs/StatsApi.md#getsales) | **Get** /sales | Get Sales
*StatsApi* | [**GetStocks**](docs/StatsApi.md#getstocks) | **Get** /stocks | Get Stocks


## Documentation For Models

 - [GetIncomes200ResponseInner](docs/GetIncomes200ResponseInner.md)
 - [GetOrders200ResponseInner](docs/GetOrders200ResponseInner.md)
 - [GetPaidStorage200ResponseInner](docs/GetPaidStorage200ResponseInner.md)
 - [GetReportDetailByPeriod200ResponseInner](docs/GetReportDetailByPeriod200ResponseInner.md)
 - [GetSales200ResponseInner](docs/GetSales200ResponseInner.md)
 - [GetStocks200ResponseInner](docs/GetStocks200ResponseInner.md)


## Documentation For Authorization



### Key

- **Type**: API key
- **API key parameter name**: key
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: key and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



