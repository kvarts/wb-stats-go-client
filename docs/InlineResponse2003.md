# InlineResponse2003

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
**Odid** | **int64** | Уникальный идентификатор позиции заказа | 
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

## Methods

### NewInlineResponse2003

`func NewInlineResponse2003(number string, date string, lastChangeDate string, supplierArticle string, techSize string, barcode string, quantity int64, totalPrice float32, discountPercent float32, isSupply bool, isRealization bool, orderId int64, promoCodeDiscount float32, warehouseName string, countryName string, oblastOkrugName string, regionName string, incomeID int64, saleID string, odid int64, spp float32, forPay float32, finishedPrice float32, priceWithDisc float32, nmId int64, subject string, category string, brand string, isStorno float32, gNumber string, ) *InlineResponse2003`

NewInlineResponse2003 instantiates a new InlineResponse2003 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003WithDefaults

`func NewInlineResponse2003WithDefaults() *InlineResponse2003`

NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse2003) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse2003) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse2003) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetDate

`func (o *InlineResponse2003) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InlineResponse2003) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InlineResponse2003) SetDate(v string)`

SetDate sets Date field to given value.


### GetLastChangeDate

`func (o *InlineResponse2003) GetLastChangeDate() string`

GetLastChangeDate returns the LastChangeDate field if non-nil, zero value otherwise.

### GetLastChangeDateOk

`func (o *InlineResponse2003) GetLastChangeDateOk() (*string, bool)`

GetLastChangeDateOk returns a tuple with the LastChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeDate

`func (o *InlineResponse2003) SetLastChangeDate(v string)`

SetLastChangeDate sets LastChangeDate field to given value.


### GetSupplierArticle

`func (o *InlineResponse2003) GetSupplierArticle() string`

GetSupplierArticle returns the SupplierArticle field if non-nil, zero value otherwise.

### GetSupplierArticleOk

`func (o *InlineResponse2003) GetSupplierArticleOk() (*string, bool)`

GetSupplierArticleOk returns a tuple with the SupplierArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierArticle

`func (o *InlineResponse2003) SetSupplierArticle(v string)`

SetSupplierArticle sets SupplierArticle field to given value.


### GetTechSize

`func (o *InlineResponse2003) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *InlineResponse2003) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *InlineResponse2003) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.


### GetBarcode

`func (o *InlineResponse2003) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *InlineResponse2003) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *InlineResponse2003) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetQuantity

`func (o *InlineResponse2003) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InlineResponse2003) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InlineResponse2003) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetTotalPrice

`func (o *InlineResponse2003) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *InlineResponse2003) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *InlineResponse2003) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.


### GetDiscountPercent

`func (o *InlineResponse2003) GetDiscountPercent() float32`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *InlineResponse2003) GetDiscountPercentOk() (*float32, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *InlineResponse2003) SetDiscountPercent(v float32)`

SetDiscountPercent sets DiscountPercent field to given value.


### GetIsSupply

`func (o *InlineResponse2003) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *InlineResponse2003) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *InlineResponse2003) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.


### GetIsRealization

`func (o *InlineResponse2003) GetIsRealization() bool`

GetIsRealization returns the IsRealization field if non-nil, zero value otherwise.

### GetIsRealizationOk

`func (o *InlineResponse2003) GetIsRealizationOk() (*bool, bool)`

GetIsRealizationOk returns a tuple with the IsRealization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealization

`func (o *InlineResponse2003) SetIsRealization(v bool)`

SetIsRealization sets IsRealization field to given value.


### GetOrderId

`func (o *InlineResponse2003) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *InlineResponse2003) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *InlineResponse2003) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetPromoCodeDiscount

`func (o *InlineResponse2003) GetPromoCodeDiscount() float32`

GetPromoCodeDiscount returns the PromoCodeDiscount field if non-nil, zero value otherwise.

### GetPromoCodeDiscountOk

`func (o *InlineResponse2003) GetPromoCodeDiscountOk() (*float32, bool)`

GetPromoCodeDiscountOk returns a tuple with the PromoCodeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCodeDiscount

`func (o *InlineResponse2003) SetPromoCodeDiscount(v float32)`

SetPromoCodeDiscount sets PromoCodeDiscount field to given value.


### GetWarehouseName

`func (o *InlineResponse2003) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *InlineResponse2003) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *InlineResponse2003) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### GetCountryName

`func (o *InlineResponse2003) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *InlineResponse2003) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *InlineResponse2003) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetOblastOkrugName

`func (o *InlineResponse2003) GetOblastOkrugName() string`

GetOblastOkrugName returns the OblastOkrugName field if non-nil, zero value otherwise.

### GetOblastOkrugNameOk

`func (o *InlineResponse2003) GetOblastOkrugNameOk() (*string, bool)`

GetOblastOkrugNameOk returns a tuple with the OblastOkrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOblastOkrugName

`func (o *InlineResponse2003) SetOblastOkrugName(v string)`

SetOblastOkrugName sets OblastOkrugName field to given value.


### GetRegionName

`func (o *InlineResponse2003) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *InlineResponse2003) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *InlineResponse2003) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetIncomeID

`func (o *InlineResponse2003) GetIncomeID() int64`

GetIncomeID returns the IncomeID field if non-nil, zero value otherwise.

### GetIncomeIDOk

`func (o *InlineResponse2003) GetIncomeIDOk() (*int64, bool)`

GetIncomeIDOk returns a tuple with the IncomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeID

`func (o *InlineResponse2003) SetIncomeID(v int64)`

SetIncomeID sets IncomeID field to given value.


### GetSaleID

`func (o *InlineResponse2003) GetSaleID() string`

GetSaleID returns the SaleID field if non-nil, zero value otherwise.

### GetSaleIDOk

`func (o *InlineResponse2003) GetSaleIDOk() (*string, bool)`

GetSaleIDOk returns a tuple with the SaleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleID

`func (o *InlineResponse2003) SetSaleID(v string)`

SetSaleID sets SaleID field to given value.


### GetOdid

`func (o *InlineResponse2003) GetOdid() int64`

GetOdid returns the Odid field if non-nil, zero value otherwise.

### GetOdidOk

`func (o *InlineResponse2003) GetOdidOk() (*int64, bool)`

GetOdidOk returns a tuple with the Odid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdid

`func (o *InlineResponse2003) SetOdid(v int64)`

SetOdid sets Odid field to given value.


### GetSpp

`func (o *InlineResponse2003) GetSpp() float32`

GetSpp returns the Spp field if non-nil, zero value otherwise.

### GetSppOk

`func (o *InlineResponse2003) GetSppOk() (*float32, bool)`

GetSppOk returns a tuple with the Spp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpp

`func (o *InlineResponse2003) SetSpp(v float32)`

SetSpp sets Spp field to given value.


### GetForPay

`func (o *InlineResponse2003) GetForPay() float32`

GetForPay returns the ForPay field if non-nil, zero value otherwise.

### GetForPayOk

`func (o *InlineResponse2003) GetForPayOk() (*float32, bool)`

GetForPayOk returns a tuple with the ForPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForPay

`func (o *InlineResponse2003) SetForPay(v float32)`

SetForPay sets ForPay field to given value.


### GetFinishedPrice

`func (o *InlineResponse2003) GetFinishedPrice() float32`

GetFinishedPrice returns the FinishedPrice field if non-nil, zero value otherwise.

### GetFinishedPriceOk

`func (o *InlineResponse2003) GetFinishedPriceOk() (*float32, bool)`

GetFinishedPriceOk returns a tuple with the FinishedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedPrice

`func (o *InlineResponse2003) SetFinishedPrice(v float32)`

SetFinishedPrice sets FinishedPrice field to given value.


### GetPriceWithDisc

`func (o *InlineResponse2003) GetPriceWithDisc() float32`

GetPriceWithDisc returns the PriceWithDisc field if non-nil, zero value otherwise.

### GetPriceWithDiscOk

`func (o *InlineResponse2003) GetPriceWithDiscOk() (*float32, bool)`

GetPriceWithDiscOk returns a tuple with the PriceWithDisc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceWithDisc

`func (o *InlineResponse2003) SetPriceWithDisc(v float32)`

SetPriceWithDisc sets PriceWithDisc field to given value.


### GetNmId

`func (o *InlineResponse2003) GetNmId() int64`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *InlineResponse2003) GetNmIdOk() (*int64, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *InlineResponse2003) SetNmId(v int64)`

SetNmId sets NmId field to given value.


### GetSubject

`func (o *InlineResponse2003) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InlineResponse2003) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InlineResponse2003) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCategory

`func (o *InlineResponse2003) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse2003) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse2003) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetBrand

`func (o *InlineResponse2003) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *InlineResponse2003) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *InlineResponse2003) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetIsStorno

`func (o *InlineResponse2003) GetIsStorno() float32`

GetIsStorno returns the IsStorno field if non-nil, zero value otherwise.

### GetIsStornoOk

`func (o *InlineResponse2003) GetIsStornoOk() (*float32, bool)`

GetIsStornoOk returns a tuple with the IsStorno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStorno

`func (o *InlineResponse2003) SetIsStorno(v float32)`

SetIsStorno sets IsStorno field to given value.


### GetGNumber

`func (o *InlineResponse2003) GetGNumber() string`

GetGNumber returns the GNumber field if non-nil, zero value otherwise.

### GetGNumberOk

`func (o *InlineResponse2003) GetGNumberOk() (*string, bool)`

GetGNumberOk returns a tuple with the GNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGNumber

`func (o *InlineResponse2003) SetGNumber(v string)`

SetGNumber sets GNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


