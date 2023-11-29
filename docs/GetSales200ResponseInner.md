# GetSales200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Номер документа | 
**Date** | **string** | Дата продажи | 
**LastChangeDate** | **string** | Дата и время обновления информации в сервисе | 
**SupplierArticle** | **string** | Артикул продавца | 
**TechSize** | **string** | Размер | 
**Barcode** | **string** | Штрих-код | 
**Quantity** | **int64** | Количество | 
**TotalPrice** | **float32** | Начальная розничная цена товара | 
**DiscountPercent** | **float32** | Согласованная скидка на товар | 
**IsSupply** | **bool** | Договор поставки | 
**IsRealization** | **bool** | Договор реализации | 
**OrderId** | **int64** | Идентификатор заказа | 
**PromoCodeDiscount** | **float32** | Согласованный промокод | 
**WarehouseName** | **string** | Склад отгрузки | 
**CountryName** | **string** | Страна | 
**OblastOkrugName** | **string** | Округ | 
**RegionName** | **string** | Регион | 
**IncomeID** | **int64** | Идентификатор поставки | 
**SaleID** | **string** | Уникальный идентификатор продажи/возврата | 
**Spp** | **float32** | Согласованная скидка постоянного покупателя | 
**ForPay** | **float32** | К перечислению поставщику | 
**FinishedPrice** | **float32** | Фактическая цена из заказа с учетом всех скидок | 
**PriceWithDisc** | **float32** | Цена, от которой считается вознаграждение поставщика forpay (с учетом всех согласованных скидок) | 
**NmId** | **int64** | код wildberries | 
**Subject** | **string** | Предмет | 
**Category** | **string** | Категория | 
**Brand** | **string** | Бренд | 
**IsStorno** | **float32** | (?) | 
**GNumber** | **string** | Номер заказа | 
**Srid** | **string** | Srid | 

## Methods

### NewGetSales200ResponseInner

`func NewGetSales200ResponseInner(number string, date string, lastChangeDate string, supplierArticle string, techSize string, barcode string, quantity int64, totalPrice float32, discountPercent float32, isSupply bool, isRealization bool, orderId int64, promoCodeDiscount float32, warehouseName string, countryName string, oblastOkrugName string, regionName string, incomeID int64, saleID string, spp float32, forPay float32, finishedPrice float32, priceWithDisc float32, nmId int64, subject string, category string, brand string, isStorno float32, gNumber string, srid string, ) *GetSales200ResponseInner`

NewGetSales200ResponseInner instantiates a new GetSales200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSales200ResponseInnerWithDefaults

`func NewGetSales200ResponseInnerWithDefaults() *GetSales200ResponseInner`

NewGetSales200ResponseInnerWithDefaults instantiates a new GetSales200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GetSales200ResponseInner) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetSales200ResponseInner) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetSales200ResponseInner) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetDate

`func (o *GetSales200ResponseInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetSales200ResponseInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetSales200ResponseInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetLastChangeDate

`func (o *GetSales200ResponseInner) GetLastChangeDate() string`

GetLastChangeDate returns the LastChangeDate field if non-nil, zero value otherwise.

### GetLastChangeDateOk

`func (o *GetSales200ResponseInner) GetLastChangeDateOk() (*string, bool)`

GetLastChangeDateOk returns a tuple with the LastChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeDate

`func (o *GetSales200ResponseInner) SetLastChangeDate(v string)`

SetLastChangeDate sets LastChangeDate field to given value.


### GetSupplierArticle

`func (o *GetSales200ResponseInner) GetSupplierArticle() string`

GetSupplierArticle returns the SupplierArticle field if non-nil, zero value otherwise.

### GetSupplierArticleOk

`func (o *GetSales200ResponseInner) GetSupplierArticleOk() (*string, bool)`

GetSupplierArticleOk returns a tuple with the SupplierArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierArticle

`func (o *GetSales200ResponseInner) SetSupplierArticle(v string)`

SetSupplierArticle sets SupplierArticle field to given value.


### GetTechSize

`func (o *GetSales200ResponseInner) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *GetSales200ResponseInner) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *GetSales200ResponseInner) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.


### GetBarcode

`func (o *GetSales200ResponseInner) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *GetSales200ResponseInner) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *GetSales200ResponseInner) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetQuantity

`func (o *GetSales200ResponseInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetSales200ResponseInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetSales200ResponseInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetTotalPrice

`func (o *GetSales200ResponseInner) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *GetSales200ResponseInner) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *GetSales200ResponseInner) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.


### GetDiscountPercent

`func (o *GetSales200ResponseInner) GetDiscountPercent() float32`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *GetSales200ResponseInner) GetDiscountPercentOk() (*float32, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *GetSales200ResponseInner) SetDiscountPercent(v float32)`

SetDiscountPercent sets DiscountPercent field to given value.


### GetIsSupply

`func (o *GetSales200ResponseInner) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *GetSales200ResponseInner) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *GetSales200ResponseInner) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.


### GetIsRealization

`func (o *GetSales200ResponseInner) GetIsRealization() bool`

GetIsRealization returns the IsRealization field if non-nil, zero value otherwise.

### GetIsRealizationOk

`func (o *GetSales200ResponseInner) GetIsRealizationOk() (*bool, bool)`

GetIsRealizationOk returns a tuple with the IsRealization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealization

`func (o *GetSales200ResponseInner) SetIsRealization(v bool)`

SetIsRealization sets IsRealization field to given value.


### GetOrderId

`func (o *GetSales200ResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetSales200ResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetSales200ResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetPromoCodeDiscount

`func (o *GetSales200ResponseInner) GetPromoCodeDiscount() float32`

GetPromoCodeDiscount returns the PromoCodeDiscount field if non-nil, zero value otherwise.

### GetPromoCodeDiscountOk

`func (o *GetSales200ResponseInner) GetPromoCodeDiscountOk() (*float32, bool)`

GetPromoCodeDiscountOk returns a tuple with the PromoCodeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCodeDiscount

`func (o *GetSales200ResponseInner) SetPromoCodeDiscount(v float32)`

SetPromoCodeDiscount sets PromoCodeDiscount field to given value.


### GetWarehouseName

`func (o *GetSales200ResponseInner) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *GetSales200ResponseInner) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *GetSales200ResponseInner) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### GetCountryName

`func (o *GetSales200ResponseInner) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *GetSales200ResponseInner) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *GetSales200ResponseInner) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetOblastOkrugName

`func (o *GetSales200ResponseInner) GetOblastOkrugName() string`

GetOblastOkrugName returns the OblastOkrugName field if non-nil, zero value otherwise.

### GetOblastOkrugNameOk

`func (o *GetSales200ResponseInner) GetOblastOkrugNameOk() (*string, bool)`

GetOblastOkrugNameOk returns a tuple with the OblastOkrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOblastOkrugName

`func (o *GetSales200ResponseInner) SetOblastOkrugName(v string)`

SetOblastOkrugName sets OblastOkrugName field to given value.


### GetRegionName

`func (o *GetSales200ResponseInner) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *GetSales200ResponseInner) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *GetSales200ResponseInner) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetIncomeID

`func (o *GetSales200ResponseInner) GetIncomeID() int64`

GetIncomeID returns the IncomeID field if non-nil, zero value otherwise.

### GetIncomeIDOk

`func (o *GetSales200ResponseInner) GetIncomeIDOk() (*int64, bool)`

GetIncomeIDOk returns a tuple with the IncomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeID

`func (o *GetSales200ResponseInner) SetIncomeID(v int64)`

SetIncomeID sets IncomeID field to given value.


### GetSaleID

`func (o *GetSales200ResponseInner) GetSaleID() string`

GetSaleID returns the SaleID field if non-nil, zero value otherwise.

### GetSaleIDOk

`func (o *GetSales200ResponseInner) GetSaleIDOk() (*string, bool)`

GetSaleIDOk returns a tuple with the SaleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleID

`func (o *GetSales200ResponseInner) SetSaleID(v string)`

SetSaleID sets SaleID field to given value.


### GetSpp

`func (o *GetSales200ResponseInner) GetSpp() float32`

GetSpp returns the Spp field if non-nil, zero value otherwise.

### GetSppOk

`func (o *GetSales200ResponseInner) GetSppOk() (*float32, bool)`

GetSppOk returns a tuple with the Spp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpp

`func (o *GetSales200ResponseInner) SetSpp(v float32)`

SetSpp sets Spp field to given value.


### GetForPay

`func (o *GetSales200ResponseInner) GetForPay() float32`

GetForPay returns the ForPay field if non-nil, zero value otherwise.

### GetForPayOk

`func (o *GetSales200ResponseInner) GetForPayOk() (*float32, bool)`

GetForPayOk returns a tuple with the ForPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForPay

`func (o *GetSales200ResponseInner) SetForPay(v float32)`

SetForPay sets ForPay field to given value.


### GetFinishedPrice

`func (o *GetSales200ResponseInner) GetFinishedPrice() float32`

GetFinishedPrice returns the FinishedPrice field if non-nil, zero value otherwise.

### GetFinishedPriceOk

`func (o *GetSales200ResponseInner) GetFinishedPriceOk() (*float32, bool)`

GetFinishedPriceOk returns a tuple with the FinishedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedPrice

`func (o *GetSales200ResponseInner) SetFinishedPrice(v float32)`

SetFinishedPrice sets FinishedPrice field to given value.


### GetPriceWithDisc

`func (o *GetSales200ResponseInner) GetPriceWithDisc() float32`

GetPriceWithDisc returns the PriceWithDisc field if non-nil, zero value otherwise.

### GetPriceWithDiscOk

`func (o *GetSales200ResponseInner) GetPriceWithDiscOk() (*float32, bool)`

GetPriceWithDiscOk returns a tuple with the PriceWithDisc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceWithDisc

`func (o *GetSales200ResponseInner) SetPriceWithDisc(v float32)`

SetPriceWithDisc sets PriceWithDisc field to given value.


### GetNmId

`func (o *GetSales200ResponseInner) GetNmId() int64`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *GetSales200ResponseInner) GetNmIdOk() (*int64, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *GetSales200ResponseInner) SetNmId(v int64)`

SetNmId sets NmId field to given value.


### GetSubject

`func (o *GetSales200ResponseInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetSales200ResponseInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetSales200ResponseInner) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCategory

`func (o *GetSales200ResponseInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetSales200ResponseInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetSales200ResponseInner) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetBrand

`func (o *GetSales200ResponseInner) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *GetSales200ResponseInner) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *GetSales200ResponseInner) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetIsStorno

`func (o *GetSales200ResponseInner) GetIsStorno() float32`

GetIsStorno returns the IsStorno field if non-nil, zero value otherwise.

### GetIsStornoOk

`func (o *GetSales200ResponseInner) GetIsStornoOk() (*float32, bool)`

GetIsStornoOk returns a tuple with the IsStorno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStorno

`func (o *GetSales200ResponseInner) SetIsStorno(v float32)`

SetIsStorno sets IsStorno field to given value.


### GetGNumber

`func (o *GetSales200ResponseInner) GetGNumber() string`

GetGNumber returns the GNumber field if non-nil, zero value otherwise.

### GetGNumberOk

`func (o *GetSales200ResponseInner) GetGNumberOk() (*string, bool)`

GetGNumberOk returns a tuple with the GNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGNumber

`func (o *GetSales200ResponseInner) SetGNumber(v string)`

SetGNumber sets GNumber field to given value.


### GetSrid

`func (o *GetSales200ResponseInner) GetSrid() string`

GetSrid returns the Srid field if non-nil, zero value otherwise.

### GetSridOk

`func (o *GetSales200ResponseInner) GetSridOk() (*string, bool)`

GetSridOk returns a tuple with the Srid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrid

`func (o *GetSales200ResponseInner) SetSrid(v string)`

SetSrid sets Srid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


