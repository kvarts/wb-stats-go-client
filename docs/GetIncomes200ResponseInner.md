# GetIncomes200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomeId** | **int64** | Номер поставки | 
**Number** | **string** | Номер УПД | 
**Date** | **string** | Дата поступления | 
**LastChangeDate** | **string** | Дата и время обновления данных в сервисе | 
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

### NewGetIncomes200ResponseInner

`func NewGetIncomes200ResponseInner(incomeId int64, number string, date string, lastChangeDate string, supplierArticle string, techSize string, barcode string, quantity int32, totalPrice float32, dateClose string, warehouseName string, nmId int64, status string, ) *GetIncomes200ResponseInner`

NewGetIncomes200ResponseInner instantiates a new GetIncomes200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIncomes200ResponseInnerWithDefaults

`func NewGetIncomes200ResponseInnerWithDefaults() *GetIncomes200ResponseInner`

NewGetIncomes200ResponseInnerWithDefaults instantiates a new GetIncomes200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomeId

`func (o *GetIncomes200ResponseInner) GetIncomeId() int64`

GetIncomeId returns the IncomeId field if non-nil, zero value otherwise.

### GetIncomeIdOk

`func (o *GetIncomes200ResponseInner) GetIncomeIdOk() (*int64, bool)`

GetIncomeIdOk returns a tuple with the IncomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeId

`func (o *GetIncomes200ResponseInner) SetIncomeId(v int64)`

SetIncomeId sets IncomeId field to given value.


### GetNumber

`func (o *GetIncomes200ResponseInner) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetIncomes200ResponseInner) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetIncomes200ResponseInner) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetDate

`func (o *GetIncomes200ResponseInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetIncomes200ResponseInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetIncomes200ResponseInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetLastChangeDate

`func (o *GetIncomes200ResponseInner) GetLastChangeDate() string`

GetLastChangeDate returns the LastChangeDate field if non-nil, zero value otherwise.

### GetLastChangeDateOk

`func (o *GetIncomes200ResponseInner) GetLastChangeDateOk() (*string, bool)`

GetLastChangeDateOk returns a tuple with the LastChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeDate

`func (o *GetIncomes200ResponseInner) SetLastChangeDate(v string)`

SetLastChangeDate sets LastChangeDate field to given value.


### GetSupplierArticle

`func (o *GetIncomes200ResponseInner) GetSupplierArticle() string`

GetSupplierArticle returns the SupplierArticle field if non-nil, zero value otherwise.

### GetSupplierArticleOk

`func (o *GetIncomes200ResponseInner) GetSupplierArticleOk() (*string, bool)`

GetSupplierArticleOk returns a tuple with the SupplierArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierArticle

`func (o *GetIncomes200ResponseInner) SetSupplierArticle(v string)`

SetSupplierArticle sets SupplierArticle field to given value.


### GetTechSize

`func (o *GetIncomes200ResponseInner) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *GetIncomes200ResponseInner) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *GetIncomes200ResponseInner) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.


### GetBarcode

`func (o *GetIncomes200ResponseInner) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *GetIncomes200ResponseInner) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *GetIncomes200ResponseInner) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetQuantity

`func (o *GetIncomes200ResponseInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetIncomes200ResponseInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetIncomes200ResponseInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTotalPrice

`func (o *GetIncomes200ResponseInner) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *GetIncomes200ResponseInner) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *GetIncomes200ResponseInner) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.


### GetDateClose

`func (o *GetIncomes200ResponseInner) GetDateClose() string`

GetDateClose returns the DateClose field if non-nil, zero value otherwise.

### GetDateCloseOk

`func (o *GetIncomes200ResponseInner) GetDateCloseOk() (*string, bool)`

GetDateCloseOk returns a tuple with the DateClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClose

`func (o *GetIncomes200ResponseInner) SetDateClose(v string)`

SetDateClose sets DateClose field to given value.


### GetWarehouseName

`func (o *GetIncomes200ResponseInner) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *GetIncomes200ResponseInner) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *GetIncomes200ResponseInner) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### GetNmId

`func (o *GetIncomes200ResponseInner) GetNmId() int64`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *GetIncomes200ResponseInner) GetNmIdOk() (*int64, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *GetIncomes200ResponseInner) SetNmId(v int64)`

SetNmId sets NmId field to given value.


### GetStatus

`func (o *GetIncomes200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIncomes200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIncomes200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


