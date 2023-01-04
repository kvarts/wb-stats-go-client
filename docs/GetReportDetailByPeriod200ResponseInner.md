# GetReportDetailByPeriod200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RealizationreportId** | **int32** | Идентификатор отчёта | 
**SuppliercontractCode** | Pointer to **int32** | Договор | [optional] 
**RrdId** | **int64** | Идентификатор записи | 
**GiId** | **int64** | Идентификатор поставки | 
**SubjectName** | **string** | Наименование товара | 
**NmId** | Pointer to **int64** | SKU товара | [optional] 
**BrandName** | **string** | Бренд товара | 
**SaName** | **string** | Артикул поставщика | 
**TsName** | **string** | Размер | 
**Barcode** | **string** | Баркод | 
**DocTypeName** | **string** | Тип документа | 
**Quantity** | **int32** | Количество | 
**RetailPrice** | **float32** | Розничная цена | 
**RetailAmount** | **float32** | Сумма продаж (возвратов) | 
**SalePercent** | **float32** | Согласованная скидка | 
**CommissionPercent** | **float32** | Процент комиссии | 
**OfficeName** | **string** | Наименование склада | 
**SupplierOperName** | **string** | Обоснование для оплаты (Тип операции?) | 
**OrderDt** | **string** | Дата заказа | 
**SaleDt** | **string** | Дата продажи | 
**RrDt** | **string** | Дата операции | 
**ShkId** | **int64** | Штрих код (короба?) | 
**RetailPriceWithdiscRub** | **float32** | Розничная цена с учётом согласованной скидки | 
**DeliveryAmount** | **int32** | Количество доставок | 
**ReturnAmount** | **int32** | Количество возвратов | 
**DeliveryRub** | **float32** | Стоимость логистики | 
**GiBoxTypeName** | **string** | Тип коробов | 
**ProductDiscountForReport** | **float32** | Согласованный продуктовый дисконт | 
**SupplierPromo** | **float32** | Промокод | 
**Rid** | **int64** | Идентификатор заказа | 
**PpvzSppPrc** | **float32** | Скидка постоянного покупателя | 
**PpvzKvwPrcBase** | **float32** | Базовый размер коэффициента вознаграждения вайлдберриз без НДС, % | 
**PpvzKvwPrc** | **float32** | Итоговый размер коэффициента вознаграждения вайлдберриз без НДС, % | 
**PpvzSalesCommission** | **float32** | Вознаграждение с продаж до вычета услуг поверенного, без НДС | 
**PpvzForPay** | **float32** | К перечислению продавцу за реализованный Товар | 
**PpvzReward** | **float32** | Возмещение расходов услуг поверенного | 
**PpvzVw** | **float32** | Вознаграждение вайлдберриз, без НДС | 
**PpvzVwNds** | **float32** | НДС с вознаграждения вайлдберриз | 
**PpvzOfficeId** | **int64** | Идентификатор ПВЗ? | 
**PpvzOfficeName** | Pointer to **string** | Наименование ПВЗ? (в случае если ppvz_office_id &#x3D; 0) | [optional] 
**PpvzSupplierId** | **int64** | Идентификатор владельца ПВЗ? | 
**PpvzSupplierName** | Pointer to **string** | Наименование владельца ПВЗ | [optional] 
**PpvzSupplierInn** | Pointer to **string** | ИНН владельца ПВЗ | [optional] 
**DeclarationNumber** | Pointer to **string** | Номер таможенной декларации | [optional] 
**StickerId** | Pointer to **string** | Цифровое значение стикера, который клеится на товар в процессе сборки заказа по системе Маркетплейс. | [optional] 
**SiteCountry** | **string** | Страна продажи | 
**Penalty** | **float32** | Штрафы | 
**AdditionalPayment** | **float32** | Доплаты | 
**Srid** | **string** | Новый идентификатор заказа | 

## Methods

### NewGetReportDetailByPeriod200ResponseInner

`func NewGetReportDetailByPeriod200ResponseInner(realizationreportId int32, rrdId int64, giId int64, subjectName string, brandName string, saName string, tsName string, barcode string, docTypeName string, quantity int32, retailPrice float32, retailAmount float32, salePercent float32, commissionPercent float32, officeName string, supplierOperName string, orderDt string, saleDt string, rrDt string, shkId int64, retailPriceWithdiscRub float32, deliveryAmount int32, returnAmount int32, deliveryRub float32, giBoxTypeName string, productDiscountForReport float32, supplierPromo float32, rid int64, ppvzSppPrc float32, ppvzKvwPrcBase float32, ppvzKvwPrc float32, ppvzSalesCommission float32, ppvzForPay float32, ppvzReward float32, ppvzVw float32, ppvzVwNds float32, ppvzOfficeId int64, ppvzSupplierId int64, siteCountry string, penalty float32, additionalPayment float32, srid string, ) *GetReportDetailByPeriod200ResponseInner`

NewGetReportDetailByPeriod200ResponseInner instantiates a new GetReportDetailByPeriod200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReportDetailByPeriod200ResponseInnerWithDefaults

`func NewGetReportDetailByPeriod200ResponseInnerWithDefaults() *GetReportDetailByPeriod200ResponseInner`

NewGetReportDetailByPeriod200ResponseInnerWithDefaults instantiates a new GetReportDetailByPeriod200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealizationreportId

`func (o *GetReportDetailByPeriod200ResponseInner) GetRealizationreportId() int32`

GetRealizationreportId returns the RealizationreportId field if non-nil, zero value otherwise.

### GetRealizationreportIdOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetRealizationreportIdOk() (*int32, bool)`

GetRealizationreportIdOk returns a tuple with the RealizationreportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizationreportId

`func (o *GetReportDetailByPeriod200ResponseInner) SetRealizationreportId(v int32)`

SetRealizationreportId sets RealizationreportId field to given value.


### GetSuppliercontractCode

`func (o *GetReportDetailByPeriod200ResponseInner) GetSuppliercontractCode() int32`

GetSuppliercontractCode returns the SuppliercontractCode field if non-nil, zero value otherwise.

### GetSuppliercontractCodeOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSuppliercontractCodeOk() (*int32, bool)`

GetSuppliercontractCodeOk returns a tuple with the SuppliercontractCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliercontractCode

`func (o *GetReportDetailByPeriod200ResponseInner) SetSuppliercontractCode(v int32)`

SetSuppliercontractCode sets SuppliercontractCode field to given value.

### HasSuppliercontractCode

`func (o *GetReportDetailByPeriod200ResponseInner) HasSuppliercontractCode() bool`

HasSuppliercontractCode returns a boolean if a field has been set.

### GetRrdId

`func (o *GetReportDetailByPeriod200ResponseInner) GetRrdId() int64`

GetRrdId returns the RrdId field if non-nil, zero value otherwise.

### GetRrdIdOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetRrdIdOk() (*int64, bool)`

GetRrdIdOk returns a tuple with the RrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrdId

`func (o *GetReportDetailByPeriod200ResponseInner) SetRrdId(v int64)`

SetRrdId sets RrdId field to given value.


### GetGiId

`func (o *GetReportDetailByPeriod200ResponseInner) GetGiId() int64`

GetGiId returns the GiId field if non-nil, zero value otherwise.

### GetGiIdOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetGiIdOk() (*int64, bool)`

GetGiIdOk returns a tuple with the GiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiId

`func (o *GetReportDetailByPeriod200ResponseInner) SetGiId(v int64)`

SetGiId sets GiId field to given value.


### GetSubjectName

`func (o *GetReportDetailByPeriod200ResponseInner) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *GetReportDetailByPeriod200ResponseInner) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.


### GetNmId

`func (o *GetReportDetailByPeriod200ResponseInner) GetNmId() int64`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetNmIdOk() (*int64, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *GetReportDetailByPeriod200ResponseInner) SetNmId(v int64)`

SetNmId sets NmId field to given value.

### HasNmId

`func (o *GetReportDetailByPeriod200ResponseInner) HasNmId() bool`

HasNmId returns a boolean if a field has been set.

### GetBrandName

`func (o *GetReportDetailByPeriod200ResponseInner) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *GetReportDetailByPeriod200ResponseInner) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.


### GetSaName

`func (o *GetReportDetailByPeriod200ResponseInner) GetSaName() string`

GetSaName returns the SaName field if non-nil, zero value otherwise.

### GetSaNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSaNameOk() (*string, bool)`

GetSaNameOk returns a tuple with the SaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaName

`func (o *GetReportDetailByPeriod200ResponseInner) SetSaName(v string)`

SetSaName sets SaName field to given value.


### GetTsName

`func (o *GetReportDetailByPeriod200ResponseInner) GetTsName() string`

GetTsName returns the TsName field if non-nil, zero value otherwise.

### GetTsNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetTsNameOk() (*string, bool)`

GetTsNameOk returns a tuple with the TsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsName

`func (o *GetReportDetailByPeriod200ResponseInner) SetTsName(v string)`

SetTsName sets TsName field to given value.


### GetBarcode

`func (o *GetReportDetailByPeriod200ResponseInner) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *GetReportDetailByPeriod200ResponseInner) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetDocTypeName

`func (o *GetReportDetailByPeriod200ResponseInner) GetDocTypeName() string`

GetDocTypeName returns the DocTypeName field if non-nil, zero value otherwise.

### GetDocTypeNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetDocTypeNameOk() (*string, bool)`

GetDocTypeNameOk returns a tuple with the DocTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocTypeName

`func (o *GetReportDetailByPeriod200ResponseInner) SetDocTypeName(v string)`

SetDocTypeName sets DocTypeName field to given value.


### GetQuantity

`func (o *GetReportDetailByPeriod200ResponseInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetReportDetailByPeriod200ResponseInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetRetailPrice

`func (o *GetReportDetailByPeriod200ResponseInner) GetRetailPrice() float32`

GetRetailPrice returns the RetailPrice field if non-nil, zero value otherwise.

### GetRetailPriceOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetRetailPriceOk() (*float32, bool)`

GetRetailPriceOk returns a tuple with the RetailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPrice

`func (o *GetReportDetailByPeriod200ResponseInner) SetRetailPrice(v float32)`

SetRetailPrice sets RetailPrice field to given value.


### GetRetailAmount

`func (o *GetReportDetailByPeriod200ResponseInner) GetRetailAmount() float32`

GetRetailAmount returns the RetailAmount field if non-nil, zero value otherwise.

### GetRetailAmountOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetRetailAmountOk() (*float32, bool)`

GetRetailAmountOk returns a tuple with the RetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailAmount

`func (o *GetReportDetailByPeriod200ResponseInner) SetRetailAmount(v float32)`

SetRetailAmount sets RetailAmount field to given value.


### GetSalePercent

`func (o *GetReportDetailByPeriod200ResponseInner) GetSalePercent() float32`

GetSalePercent returns the SalePercent field if non-nil, zero value otherwise.

### GetSalePercentOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSalePercentOk() (*float32, bool)`

GetSalePercentOk returns a tuple with the SalePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalePercent

`func (o *GetReportDetailByPeriod200ResponseInner) SetSalePercent(v float32)`

SetSalePercent sets SalePercent field to given value.


### GetCommissionPercent

`func (o *GetReportDetailByPeriod200ResponseInner) GetCommissionPercent() float32`

GetCommissionPercent returns the CommissionPercent field if non-nil, zero value otherwise.

### GetCommissionPercentOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetCommissionPercentOk() (*float32, bool)`

GetCommissionPercentOk returns a tuple with the CommissionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionPercent

`func (o *GetReportDetailByPeriod200ResponseInner) SetCommissionPercent(v float32)`

SetCommissionPercent sets CommissionPercent field to given value.


### GetOfficeName

`func (o *GetReportDetailByPeriod200ResponseInner) GetOfficeName() string`

GetOfficeName returns the OfficeName field if non-nil, zero value otherwise.

### GetOfficeNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetOfficeNameOk() (*string, bool)`

GetOfficeNameOk returns a tuple with the OfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeName

`func (o *GetReportDetailByPeriod200ResponseInner) SetOfficeName(v string)`

SetOfficeName sets OfficeName field to given value.


### GetSupplierOperName

`func (o *GetReportDetailByPeriod200ResponseInner) GetSupplierOperName() string`

GetSupplierOperName returns the SupplierOperName field if non-nil, zero value otherwise.

### GetSupplierOperNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSupplierOperNameOk() (*string, bool)`

GetSupplierOperNameOk returns a tuple with the SupplierOperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierOperName

`func (o *GetReportDetailByPeriod200ResponseInner) SetSupplierOperName(v string)`

SetSupplierOperName sets SupplierOperName field to given value.


### GetOrderDt

`func (o *GetReportDetailByPeriod200ResponseInner) GetOrderDt() string`

GetOrderDt returns the OrderDt field if non-nil, zero value otherwise.

### GetOrderDtOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetOrderDtOk() (*string, bool)`

GetOrderDtOk returns a tuple with the OrderDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDt

`func (o *GetReportDetailByPeriod200ResponseInner) SetOrderDt(v string)`

SetOrderDt sets OrderDt field to given value.


### GetSaleDt

`func (o *GetReportDetailByPeriod200ResponseInner) GetSaleDt() string`

GetSaleDt returns the SaleDt field if non-nil, zero value otherwise.

### GetSaleDtOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSaleDtOk() (*string, bool)`

GetSaleDtOk returns a tuple with the SaleDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDt

`func (o *GetReportDetailByPeriod200ResponseInner) SetSaleDt(v string)`

SetSaleDt sets SaleDt field to given value.


### GetRrDt

`func (o *GetReportDetailByPeriod200ResponseInner) GetRrDt() string`

GetRrDt returns the RrDt field if non-nil, zero value otherwise.

### GetRrDtOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetRrDtOk() (*string, bool)`

GetRrDtOk returns a tuple with the RrDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrDt

`func (o *GetReportDetailByPeriod200ResponseInner) SetRrDt(v string)`

SetRrDt sets RrDt field to given value.


### GetShkId

`func (o *GetReportDetailByPeriod200ResponseInner) GetShkId() int64`

GetShkId returns the ShkId field if non-nil, zero value otherwise.

### GetShkIdOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetShkIdOk() (*int64, bool)`

GetShkIdOk returns a tuple with the ShkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShkId

`func (o *GetReportDetailByPeriod200ResponseInner) SetShkId(v int64)`

SetShkId sets ShkId field to given value.


### GetRetailPriceWithdiscRub

`func (o *GetReportDetailByPeriod200ResponseInner) GetRetailPriceWithdiscRub() float32`

GetRetailPriceWithdiscRub returns the RetailPriceWithdiscRub field if non-nil, zero value otherwise.

### GetRetailPriceWithdiscRubOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetRetailPriceWithdiscRubOk() (*float32, bool)`

GetRetailPriceWithdiscRubOk returns a tuple with the RetailPriceWithdiscRub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPriceWithdiscRub

`func (o *GetReportDetailByPeriod200ResponseInner) SetRetailPriceWithdiscRub(v float32)`

SetRetailPriceWithdiscRub sets RetailPriceWithdiscRub field to given value.


### GetDeliveryAmount

`func (o *GetReportDetailByPeriod200ResponseInner) GetDeliveryAmount() int32`

GetDeliveryAmount returns the DeliveryAmount field if non-nil, zero value otherwise.

### GetDeliveryAmountOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetDeliveryAmountOk() (*int32, bool)`

GetDeliveryAmountOk returns a tuple with the DeliveryAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAmount

`func (o *GetReportDetailByPeriod200ResponseInner) SetDeliveryAmount(v int32)`

SetDeliveryAmount sets DeliveryAmount field to given value.


### GetReturnAmount

`func (o *GetReportDetailByPeriod200ResponseInner) GetReturnAmount() int32`

GetReturnAmount returns the ReturnAmount field if non-nil, zero value otherwise.

### GetReturnAmountOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetReturnAmountOk() (*int32, bool)`

GetReturnAmountOk returns a tuple with the ReturnAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAmount

`func (o *GetReportDetailByPeriod200ResponseInner) SetReturnAmount(v int32)`

SetReturnAmount sets ReturnAmount field to given value.


### GetDeliveryRub

`func (o *GetReportDetailByPeriod200ResponseInner) GetDeliveryRub() float32`

GetDeliveryRub returns the DeliveryRub field if non-nil, zero value otherwise.

### GetDeliveryRubOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetDeliveryRubOk() (*float32, bool)`

GetDeliveryRubOk returns a tuple with the DeliveryRub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryRub

`func (o *GetReportDetailByPeriod200ResponseInner) SetDeliveryRub(v float32)`

SetDeliveryRub sets DeliveryRub field to given value.


### GetGiBoxTypeName

`func (o *GetReportDetailByPeriod200ResponseInner) GetGiBoxTypeName() string`

GetGiBoxTypeName returns the GiBoxTypeName field if non-nil, zero value otherwise.

### GetGiBoxTypeNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetGiBoxTypeNameOk() (*string, bool)`

GetGiBoxTypeNameOk returns a tuple with the GiBoxTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiBoxTypeName

`func (o *GetReportDetailByPeriod200ResponseInner) SetGiBoxTypeName(v string)`

SetGiBoxTypeName sets GiBoxTypeName field to given value.


### GetProductDiscountForReport

`func (o *GetReportDetailByPeriod200ResponseInner) GetProductDiscountForReport() float32`

GetProductDiscountForReport returns the ProductDiscountForReport field if non-nil, zero value otherwise.

### GetProductDiscountForReportOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetProductDiscountForReportOk() (*float32, bool)`

GetProductDiscountForReportOk returns a tuple with the ProductDiscountForReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDiscountForReport

`func (o *GetReportDetailByPeriod200ResponseInner) SetProductDiscountForReport(v float32)`

SetProductDiscountForReport sets ProductDiscountForReport field to given value.


### GetSupplierPromo

`func (o *GetReportDetailByPeriod200ResponseInner) GetSupplierPromo() float32`

GetSupplierPromo returns the SupplierPromo field if non-nil, zero value otherwise.

### GetSupplierPromoOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSupplierPromoOk() (*float32, bool)`

GetSupplierPromoOk returns a tuple with the SupplierPromo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierPromo

`func (o *GetReportDetailByPeriod200ResponseInner) SetSupplierPromo(v float32)`

SetSupplierPromo sets SupplierPromo field to given value.


### GetRid

`func (o *GetReportDetailByPeriod200ResponseInner) GetRid() int64`

GetRid returns the Rid field if non-nil, zero value otherwise.

### GetRidOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetRidOk() (*int64, bool)`

GetRidOk returns a tuple with the Rid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRid

`func (o *GetReportDetailByPeriod200ResponseInner) SetRid(v int64)`

SetRid sets Rid field to given value.


### GetPpvzSppPrc

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSppPrc() float32`

GetPpvzSppPrc returns the PpvzSppPrc field if non-nil, zero value otherwise.

### GetPpvzSppPrcOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSppPrcOk() (*float32, bool)`

GetPpvzSppPrcOk returns a tuple with the PpvzSppPrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzSppPrc

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzSppPrc(v float32)`

SetPpvzSppPrc sets PpvzSppPrc field to given value.


### GetPpvzKvwPrcBase

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzKvwPrcBase() float32`

GetPpvzKvwPrcBase returns the PpvzKvwPrcBase field if non-nil, zero value otherwise.

### GetPpvzKvwPrcBaseOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzKvwPrcBaseOk() (*float32, bool)`

GetPpvzKvwPrcBaseOk returns a tuple with the PpvzKvwPrcBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzKvwPrcBase

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzKvwPrcBase(v float32)`

SetPpvzKvwPrcBase sets PpvzKvwPrcBase field to given value.


### GetPpvzKvwPrc

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzKvwPrc() float32`

GetPpvzKvwPrc returns the PpvzKvwPrc field if non-nil, zero value otherwise.

### GetPpvzKvwPrcOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzKvwPrcOk() (*float32, bool)`

GetPpvzKvwPrcOk returns a tuple with the PpvzKvwPrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzKvwPrc

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzKvwPrc(v float32)`

SetPpvzKvwPrc sets PpvzKvwPrc field to given value.


### GetPpvzSalesCommission

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSalesCommission() float32`

GetPpvzSalesCommission returns the PpvzSalesCommission field if non-nil, zero value otherwise.

### GetPpvzSalesCommissionOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSalesCommissionOk() (*float32, bool)`

GetPpvzSalesCommissionOk returns a tuple with the PpvzSalesCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzSalesCommission

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzSalesCommission(v float32)`

SetPpvzSalesCommission sets PpvzSalesCommission field to given value.


### GetPpvzForPay

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzForPay() float32`

GetPpvzForPay returns the PpvzForPay field if non-nil, zero value otherwise.

### GetPpvzForPayOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzForPayOk() (*float32, bool)`

GetPpvzForPayOk returns a tuple with the PpvzForPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzForPay

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzForPay(v float32)`

SetPpvzForPay sets PpvzForPay field to given value.


### GetPpvzReward

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzReward() float32`

GetPpvzReward returns the PpvzReward field if non-nil, zero value otherwise.

### GetPpvzRewardOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzRewardOk() (*float32, bool)`

GetPpvzRewardOk returns a tuple with the PpvzReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzReward

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzReward(v float32)`

SetPpvzReward sets PpvzReward field to given value.


### GetPpvzVw

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzVw() float32`

GetPpvzVw returns the PpvzVw field if non-nil, zero value otherwise.

### GetPpvzVwOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzVwOk() (*float32, bool)`

GetPpvzVwOk returns a tuple with the PpvzVw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzVw

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzVw(v float32)`

SetPpvzVw sets PpvzVw field to given value.


### GetPpvzVwNds

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzVwNds() float32`

GetPpvzVwNds returns the PpvzVwNds field if non-nil, zero value otherwise.

### GetPpvzVwNdsOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzVwNdsOk() (*float32, bool)`

GetPpvzVwNdsOk returns a tuple with the PpvzVwNds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzVwNds

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzVwNds(v float32)`

SetPpvzVwNds sets PpvzVwNds field to given value.


### GetPpvzOfficeId

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzOfficeId() int64`

GetPpvzOfficeId returns the PpvzOfficeId field if non-nil, zero value otherwise.

### GetPpvzOfficeIdOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzOfficeIdOk() (*int64, bool)`

GetPpvzOfficeIdOk returns a tuple with the PpvzOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzOfficeId

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzOfficeId(v int64)`

SetPpvzOfficeId sets PpvzOfficeId field to given value.


### GetPpvzOfficeName

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzOfficeName() string`

GetPpvzOfficeName returns the PpvzOfficeName field if non-nil, zero value otherwise.

### GetPpvzOfficeNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzOfficeNameOk() (*string, bool)`

GetPpvzOfficeNameOk returns a tuple with the PpvzOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzOfficeName

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzOfficeName(v string)`

SetPpvzOfficeName sets PpvzOfficeName field to given value.

### HasPpvzOfficeName

`func (o *GetReportDetailByPeriod200ResponseInner) HasPpvzOfficeName() bool`

HasPpvzOfficeName returns a boolean if a field has been set.

### GetPpvzSupplierId

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSupplierId() int64`

GetPpvzSupplierId returns the PpvzSupplierId field if non-nil, zero value otherwise.

### GetPpvzSupplierIdOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSupplierIdOk() (*int64, bool)`

GetPpvzSupplierIdOk returns a tuple with the PpvzSupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzSupplierId

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzSupplierId(v int64)`

SetPpvzSupplierId sets PpvzSupplierId field to given value.


### GetPpvzSupplierName

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSupplierName() string`

GetPpvzSupplierName returns the PpvzSupplierName field if non-nil, zero value otherwise.

### GetPpvzSupplierNameOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSupplierNameOk() (*string, bool)`

GetPpvzSupplierNameOk returns a tuple with the PpvzSupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzSupplierName

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzSupplierName(v string)`

SetPpvzSupplierName sets PpvzSupplierName field to given value.

### HasPpvzSupplierName

`func (o *GetReportDetailByPeriod200ResponseInner) HasPpvzSupplierName() bool`

HasPpvzSupplierName returns a boolean if a field has been set.

### GetPpvzSupplierInn

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSupplierInn() string`

GetPpvzSupplierInn returns the PpvzSupplierInn field if non-nil, zero value otherwise.

### GetPpvzSupplierInnOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPpvzSupplierInnOk() (*string, bool)`

GetPpvzSupplierInnOk returns a tuple with the PpvzSupplierInn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpvzSupplierInn

`func (o *GetReportDetailByPeriod200ResponseInner) SetPpvzSupplierInn(v string)`

SetPpvzSupplierInn sets PpvzSupplierInn field to given value.

### HasPpvzSupplierInn

`func (o *GetReportDetailByPeriod200ResponseInner) HasPpvzSupplierInn() bool`

HasPpvzSupplierInn returns a boolean if a field has been set.

### GetDeclarationNumber

`func (o *GetReportDetailByPeriod200ResponseInner) GetDeclarationNumber() string`

GetDeclarationNumber returns the DeclarationNumber field if non-nil, zero value otherwise.

### GetDeclarationNumberOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetDeclarationNumberOk() (*string, bool)`

GetDeclarationNumberOk returns a tuple with the DeclarationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationNumber

`func (o *GetReportDetailByPeriod200ResponseInner) SetDeclarationNumber(v string)`

SetDeclarationNumber sets DeclarationNumber field to given value.

### HasDeclarationNumber

`func (o *GetReportDetailByPeriod200ResponseInner) HasDeclarationNumber() bool`

HasDeclarationNumber returns a boolean if a field has been set.

### GetStickerId

`func (o *GetReportDetailByPeriod200ResponseInner) GetStickerId() string`

GetStickerId returns the StickerId field if non-nil, zero value otherwise.

### GetStickerIdOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetStickerIdOk() (*string, bool)`

GetStickerIdOk returns a tuple with the StickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerId

`func (o *GetReportDetailByPeriod200ResponseInner) SetStickerId(v string)`

SetStickerId sets StickerId field to given value.

### HasStickerId

`func (o *GetReportDetailByPeriod200ResponseInner) HasStickerId() bool`

HasStickerId returns a boolean if a field has been set.

### GetSiteCountry

`func (o *GetReportDetailByPeriod200ResponseInner) GetSiteCountry() string`

GetSiteCountry returns the SiteCountry field if non-nil, zero value otherwise.

### GetSiteCountryOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSiteCountryOk() (*string, bool)`

GetSiteCountryOk returns a tuple with the SiteCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCountry

`func (o *GetReportDetailByPeriod200ResponseInner) SetSiteCountry(v string)`

SetSiteCountry sets SiteCountry field to given value.


### GetPenalty

`func (o *GetReportDetailByPeriod200ResponseInner) GetPenalty() float32`

GetPenalty returns the Penalty field if non-nil, zero value otherwise.

### GetPenaltyOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetPenaltyOk() (*float32, bool)`

GetPenaltyOk returns a tuple with the Penalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenalty

`func (o *GetReportDetailByPeriod200ResponseInner) SetPenalty(v float32)`

SetPenalty sets Penalty field to given value.


### GetAdditionalPayment

`func (o *GetReportDetailByPeriod200ResponseInner) GetAdditionalPayment() float32`

GetAdditionalPayment returns the AdditionalPayment field if non-nil, zero value otherwise.

### GetAdditionalPaymentOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetAdditionalPaymentOk() (*float32, bool)`

GetAdditionalPaymentOk returns a tuple with the AdditionalPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPayment

`func (o *GetReportDetailByPeriod200ResponseInner) SetAdditionalPayment(v float32)`

SetAdditionalPayment sets AdditionalPayment field to given value.


### GetSrid

`func (o *GetReportDetailByPeriod200ResponseInner) GetSrid() string`

GetSrid returns the Srid field if non-nil, zero value otherwise.

### GetSridOk

`func (o *GetReportDetailByPeriod200ResponseInner) GetSridOk() (*string, bool)`

GetSridOk returns a tuple with the Srid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrid

`func (o *GetReportDetailByPeriod200ResponseInner) SetSrid(v string)`

SetSrid sets Srid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


