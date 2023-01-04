# GetStocks200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChangeDate** | **time.Time** | Дата и время обновления данных в сервисе | 
**SupplierArticle** | **string** | Артикул продавца | 
**TechSize** | **string** | Размер | 
**Barcode** | **string** | Штрих-код | 
**Quantity** | **int64** | Количество доступное для продажи | 
**IsSupply** | **bool** | Договор поставки | 
**IsRealization** | **bool** | Договор реализации | 
**QuantityFull** | **int64** | Полное количество | 
**WarehouseName** | **string** | Название склада | 
**NmId** | **int64** | Код wildberries | 
**Subject** | **string** | Предмет | 
**Category** | **string** | Категория | 
**DaysOnSite** | **int32** | Количество дней на сайте | 
**Brand** | **string** | Бренд | 
**SCCode** | **string** | Код контракта | 
**Price** | **float32** | Цена (?) | 
**Discount** | **float32** | Скидка (?) | 

## Methods

### NewGetStocks200ResponseInner

`func NewGetStocks200ResponseInner(lastChangeDate time.Time, supplierArticle string, techSize string, barcode string, quantity int64, isSupply bool, isRealization bool, quantityFull int64, warehouseName string, nmId int64, subject string, category string, daysOnSite int32, brand string, sCCode string, price float32, discount float32, ) *GetStocks200ResponseInner`

NewGetStocks200ResponseInner instantiates a new GetStocks200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStocks200ResponseInnerWithDefaults

`func NewGetStocks200ResponseInnerWithDefaults() *GetStocks200ResponseInner`

NewGetStocks200ResponseInnerWithDefaults instantiates a new GetStocks200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChangeDate

`func (o *GetStocks200ResponseInner) GetLastChangeDate() time.Time`

GetLastChangeDate returns the LastChangeDate field if non-nil, zero value otherwise.

### GetLastChangeDateOk

`func (o *GetStocks200ResponseInner) GetLastChangeDateOk() (*time.Time, bool)`

GetLastChangeDateOk returns a tuple with the LastChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeDate

`func (o *GetStocks200ResponseInner) SetLastChangeDate(v time.Time)`

SetLastChangeDate sets LastChangeDate field to given value.


### GetSupplierArticle

`func (o *GetStocks200ResponseInner) GetSupplierArticle() string`

GetSupplierArticle returns the SupplierArticle field if non-nil, zero value otherwise.

### GetSupplierArticleOk

`func (o *GetStocks200ResponseInner) GetSupplierArticleOk() (*string, bool)`

GetSupplierArticleOk returns a tuple with the SupplierArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierArticle

`func (o *GetStocks200ResponseInner) SetSupplierArticle(v string)`

SetSupplierArticle sets SupplierArticle field to given value.


### GetTechSize

`func (o *GetStocks200ResponseInner) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *GetStocks200ResponseInner) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *GetStocks200ResponseInner) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.


### GetBarcode

`func (o *GetStocks200ResponseInner) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *GetStocks200ResponseInner) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *GetStocks200ResponseInner) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetQuantity

`func (o *GetStocks200ResponseInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetStocks200ResponseInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetStocks200ResponseInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetIsSupply

`func (o *GetStocks200ResponseInner) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *GetStocks200ResponseInner) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *GetStocks200ResponseInner) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.


### GetIsRealization

`func (o *GetStocks200ResponseInner) GetIsRealization() bool`

GetIsRealization returns the IsRealization field if non-nil, zero value otherwise.

### GetIsRealizationOk

`func (o *GetStocks200ResponseInner) GetIsRealizationOk() (*bool, bool)`

GetIsRealizationOk returns a tuple with the IsRealization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealization

`func (o *GetStocks200ResponseInner) SetIsRealization(v bool)`

SetIsRealization sets IsRealization field to given value.


### GetQuantityFull

`func (o *GetStocks200ResponseInner) GetQuantityFull() int64`

GetQuantityFull returns the QuantityFull field if non-nil, zero value otherwise.

### GetQuantityFullOk

`func (o *GetStocks200ResponseInner) GetQuantityFullOk() (*int64, bool)`

GetQuantityFullOk returns a tuple with the QuantityFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityFull

`func (o *GetStocks200ResponseInner) SetQuantityFull(v int64)`

SetQuantityFull sets QuantityFull field to given value.


### GetWarehouseName

`func (o *GetStocks200ResponseInner) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *GetStocks200ResponseInner) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *GetStocks200ResponseInner) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### GetNmId

`func (o *GetStocks200ResponseInner) GetNmId() int64`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *GetStocks200ResponseInner) GetNmIdOk() (*int64, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *GetStocks200ResponseInner) SetNmId(v int64)`

SetNmId sets NmId field to given value.


### GetSubject

`func (o *GetStocks200ResponseInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetStocks200ResponseInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetStocks200ResponseInner) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCategory

`func (o *GetStocks200ResponseInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetStocks200ResponseInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetStocks200ResponseInner) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDaysOnSite

`func (o *GetStocks200ResponseInner) GetDaysOnSite() int32`

GetDaysOnSite returns the DaysOnSite field if non-nil, zero value otherwise.

### GetDaysOnSiteOk

`func (o *GetStocks200ResponseInner) GetDaysOnSiteOk() (*int32, bool)`

GetDaysOnSiteOk returns a tuple with the DaysOnSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOnSite

`func (o *GetStocks200ResponseInner) SetDaysOnSite(v int32)`

SetDaysOnSite sets DaysOnSite field to given value.


### GetBrand

`func (o *GetStocks200ResponseInner) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *GetStocks200ResponseInner) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *GetStocks200ResponseInner) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetSCCode

`func (o *GetStocks200ResponseInner) GetSCCode() string`

GetSCCode returns the SCCode field if non-nil, zero value otherwise.

### GetSCCodeOk

`func (o *GetStocks200ResponseInner) GetSCCodeOk() (*string, bool)`

GetSCCodeOk returns a tuple with the SCCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCCode

`func (o *GetStocks200ResponseInner) SetSCCode(v string)`

SetSCCode sets SCCode field to given value.


### GetPrice

`func (o *GetStocks200ResponseInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetStocks200ResponseInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetStocks200ResponseInner) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetDiscount

`func (o *GetStocks200ResponseInner) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *GetStocks200ResponseInner) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *GetStocks200ResponseInner) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


