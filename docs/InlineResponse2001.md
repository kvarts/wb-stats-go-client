# InlineResponse2001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomeId** | **int64** | Номер поставки | 
**Number** | **string** | Номер УПД | 
**Date** | **time.Time** | Дата поступления | 
**LastChangeDate** | **time.Time** | Дата и время обновления данных в сервисе | 
**SupplierArticle** | **string** | Артикул продавца | 
**TechSize** | **string** | Размер | 
**Barcode** | **string** | Штрих-код | 
**Quantity** | **int32** | Количество | 
**TotalPrice** | **float32** | Цена из УПД | 
**DateClose** | **string** | Дата принятия (закрытия) у wildberies | 
**WarehouseName** | **string** | Наименование склада | 
**NmId** | **int64** | Код WB | 
**Status** | **string** | Текущий статус поставки | 

## Methods

### NewInlineResponse2001

`func NewInlineResponse2001(incomeId int64, number string, date time.Time, lastChangeDate time.Time, supplierArticle string, techSize string, barcode string, quantity int32, totalPrice float32, dateClose string, warehouseName string, nmId int64, status string, ) *InlineResponse2001`

NewInlineResponse2001 instantiates a new InlineResponse2001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001WithDefaults

`func NewInlineResponse2001WithDefaults() *InlineResponse2001`

NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomeId

`func (o *InlineResponse2001) GetIncomeId() int64`

GetIncomeId returns the IncomeId field if non-nil, zero value otherwise.

### GetIncomeIdOk

`func (o *InlineResponse2001) GetIncomeIdOk() (*int64, bool)`

GetIncomeIdOk returns a tuple with the IncomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeId

`func (o *InlineResponse2001) SetIncomeId(v int64)`

SetIncomeId sets IncomeId field to given value.


### GetNumber

`func (o *InlineResponse2001) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse2001) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse2001) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetDate

`func (o *InlineResponse2001) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InlineResponse2001) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InlineResponse2001) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetLastChangeDate

`func (o *InlineResponse2001) GetLastChangeDate() time.Time`

GetLastChangeDate returns the LastChangeDate field if non-nil, zero value otherwise.

### GetLastChangeDateOk

`func (o *InlineResponse2001) GetLastChangeDateOk() (*time.Time, bool)`

GetLastChangeDateOk returns a tuple with the LastChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeDate

`func (o *InlineResponse2001) SetLastChangeDate(v time.Time)`

SetLastChangeDate sets LastChangeDate field to given value.


### GetSupplierArticle

`func (o *InlineResponse2001) GetSupplierArticle() string`

GetSupplierArticle returns the SupplierArticle field if non-nil, zero value otherwise.

### GetSupplierArticleOk

`func (o *InlineResponse2001) GetSupplierArticleOk() (*string, bool)`

GetSupplierArticleOk returns a tuple with the SupplierArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierArticle

`func (o *InlineResponse2001) SetSupplierArticle(v string)`

SetSupplierArticle sets SupplierArticle field to given value.


### GetTechSize

`func (o *InlineResponse2001) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *InlineResponse2001) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *InlineResponse2001) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.


### GetBarcode

`func (o *InlineResponse2001) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *InlineResponse2001) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *InlineResponse2001) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetQuantity

`func (o *InlineResponse2001) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InlineResponse2001) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InlineResponse2001) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTotalPrice

`func (o *InlineResponse2001) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *InlineResponse2001) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *InlineResponse2001) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.


### GetDateClose

`func (o *InlineResponse2001) GetDateClose() string`

GetDateClose returns the DateClose field if non-nil, zero value otherwise.

### GetDateCloseOk

`func (o *InlineResponse2001) GetDateCloseOk() (*string, bool)`

GetDateCloseOk returns a tuple with the DateClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClose

`func (o *InlineResponse2001) SetDateClose(v string)`

SetDateClose sets DateClose field to given value.


### GetWarehouseName

`func (o *InlineResponse2001) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *InlineResponse2001) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *InlineResponse2001) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### GetNmId

`func (o *InlineResponse2001) GetNmId() int64`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *InlineResponse2001) GetNmIdOk() (*int64, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *InlineResponse2001) SetNmId(v int64)`

SetNmId sets NmId field to given value.


### GetStatus

`func (o *InlineResponse2001) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2001) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2001) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


