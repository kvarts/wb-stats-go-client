# GetOrders200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int64** | номер заказа | 
**Date** | **string** | дата заказа | 
**LastChangeDate** | **string** | дата время обновления информации в сервисе | 
**SupplierArticle** | **string** | ваш артикул | 
**TechSize** | **string** | размер | 
**Barcode** | **string** | штрих-код | 
**TotalPrice** | **float32** | цена до согласованной скидки/промо/спп | 
**DiscountPercent** | **float32** | согласованный итоговый дисконт | 
**WarehouseName** | **string** | склад отгрузки | 
**Oblast** | **string** | область | 
**IncomeID** | **int64** | номер поставки | 
**Odid** | **int64** | уникальный идентификатор позиции заказа | 
**NmId** | **int32** | Код WB | 
**Subject** | **string** | предмет | 
**Category** | **string** | категория | 
**Brand** | **string** | бренд | 
**IsCancel** | **bool** | признак отмены заказа (0 – отмены не было, 1 – отмена была | 
**CancelDt** | **string** | дата отмены заказа | 
**GNumber** | **string** | Номер заказа | 
**Srid** | **string** | Srid | 
**OrderType** | Pointer to **string** | Тип поступившего заказа | [optional] 

## Methods

### NewGetOrders200ResponseInner

`func NewGetOrders200ResponseInner(number int64, date string, lastChangeDate string, supplierArticle string, techSize string, barcode string, totalPrice float32, discountPercent float32, warehouseName string, oblast string, incomeID int64, odid int64, nmId int32, subject string, category string, brand string, isCancel bool, cancelDt string, gNumber string, srid string, ) *GetOrders200ResponseInner`

NewGetOrders200ResponseInner instantiates a new GetOrders200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrders200ResponseInnerWithDefaults

`func NewGetOrders200ResponseInnerWithDefaults() *GetOrders200ResponseInner`

NewGetOrders200ResponseInnerWithDefaults instantiates a new GetOrders200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GetOrders200ResponseInner) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetOrders200ResponseInner) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetOrders200ResponseInner) SetNumber(v int64)`

SetNumber sets Number field to given value.


### GetDate

`func (o *GetOrders200ResponseInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetOrders200ResponseInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetOrders200ResponseInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetLastChangeDate

`func (o *GetOrders200ResponseInner) GetLastChangeDate() string`

GetLastChangeDate returns the LastChangeDate field if non-nil, zero value otherwise.

### GetLastChangeDateOk

`func (o *GetOrders200ResponseInner) GetLastChangeDateOk() (*string, bool)`

GetLastChangeDateOk returns a tuple with the LastChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeDate

`func (o *GetOrders200ResponseInner) SetLastChangeDate(v string)`

SetLastChangeDate sets LastChangeDate field to given value.


### GetSupplierArticle

`func (o *GetOrders200ResponseInner) GetSupplierArticle() string`

GetSupplierArticle returns the SupplierArticle field if non-nil, zero value otherwise.

### GetSupplierArticleOk

`func (o *GetOrders200ResponseInner) GetSupplierArticleOk() (*string, bool)`

GetSupplierArticleOk returns a tuple with the SupplierArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierArticle

`func (o *GetOrders200ResponseInner) SetSupplierArticle(v string)`

SetSupplierArticle sets SupplierArticle field to given value.


### GetTechSize

`func (o *GetOrders200ResponseInner) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *GetOrders200ResponseInner) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *GetOrders200ResponseInner) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.


### GetBarcode

`func (o *GetOrders200ResponseInner) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *GetOrders200ResponseInner) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *GetOrders200ResponseInner) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetTotalPrice

`func (o *GetOrders200ResponseInner) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *GetOrders200ResponseInner) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *GetOrders200ResponseInner) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.


### GetDiscountPercent

`func (o *GetOrders200ResponseInner) GetDiscountPercent() float32`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *GetOrders200ResponseInner) GetDiscountPercentOk() (*float32, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *GetOrders200ResponseInner) SetDiscountPercent(v float32)`

SetDiscountPercent sets DiscountPercent field to given value.


### GetWarehouseName

`func (o *GetOrders200ResponseInner) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *GetOrders200ResponseInner) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *GetOrders200ResponseInner) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### GetOblast

`func (o *GetOrders200ResponseInner) GetOblast() string`

GetOblast returns the Oblast field if non-nil, zero value otherwise.

### GetOblastOk

`func (o *GetOrders200ResponseInner) GetOblastOk() (*string, bool)`

GetOblastOk returns a tuple with the Oblast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOblast

`func (o *GetOrders200ResponseInner) SetOblast(v string)`

SetOblast sets Oblast field to given value.


### GetIncomeID

`func (o *GetOrders200ResponseInner) GetIncomeID() int64`

GetIncomeID returns the IncomeID field if non-nil, zero value otherwise.

### GetIncomeIDOk

`func (o *GetOrders200ResponseInner) GetIncomeIDOk() (*int64, bool)`

GetIncomeIDOk returns a tuple with the IncomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeID

`func (o *GetOrders200ResponseInner) SetIncomeID(v int64)`

SetIncomeID sets IncomeID field to given value.


### GetOdid

`func (o *GetOrders200ResponseInner) GetOdid() int64`

GetOdid returns the Odid field if non-nil, zero value otherwise.

### GetOdidOk

`func (o *GetOrders200ResponseInner) GetOdidOk() (*int64, bool)`

GetOdidOk returns a tuple with the Odid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdid

`func (o *GetOrders200ResponseInner) SetOdid(v int64)`

SetOdid sets Odid field to given value.


### GetNmId

`func (o *GetOrders200ResponseInner) GetNmId() int32`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *GetOrders200ResponseInner) GetNmIdOk() (*int32, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *GetOrders200ResponseInner) SetNmId(v int32)`

SetNmId sets NmId field to given value.


### GetSubject

`func (o *GetOrders200ResponseInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetOrders200ResponseInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetOrders200ResponseInner) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCategory

`func (o *GetOrders200ResponseInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetOrders200ResponseInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetOrders200ResponseInner) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetBrand

`func (o *GetOrders200ResponseInner) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *GetOrders200ResponseInner) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *GetOrders200ResponseInner) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetIsCancel

`func (o *GetOrders200ResponseInner) GetIsCancel() bool`

GetIsCancel returns the IsCancel field if non-nil, zero value otherwise.

### GetIsCancelOk

`func (o *GetOrders200ResponseInner) GetIsCancelOk() (*bool, bool)`

GetIsCancelOk returns a tuple with the IsCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancel

`func (o *GetOrders200ResponseInner) SetIsCancel(v bool)`

SetIsCancel sets IsCancel field to given value.


### GetCancelDt

`func (o *GetOrders200ResponseInner) GetCancelDt() string`

GetCancelDt returns the CancelDt field if non-nil, zero value otherwise.

### GetCancelDtOk

`func (o *GetOrders200ResponseInner) GetCancelDtOk() (*string, bool)`

GetCancelDtOk returns a tuple with the CancelDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDt

`func (o *GetOrders200ResponseInner) SetCancelDt(v string)`

SetCancelDt sets CancelDt field to given value.


### GetGNumber

`func (o *GetOrders200ResponseInner) GetGNumber() string`

GetGNumber returns the GNumber field if non-nil, zero value otherwise.

### GetGNumberOk

`func (o *GetOrders200ResponseInner) GetGNumberOk() (*string, bool)`

GetGNumberOk returns a tuple with the GNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGNumber

`func (o *GetOrders200ResponseInner) SetGNumber(v string)`

SetGNumber sets GNumber field to given value.


### GetSrid

`func (o *GetOrders200ResponseInner) GetSrid() string`

GetSrid returns the Srid field if non-nil, zero value otherwise.

### GetSridOk

`func (o *GetOrders200ResponseInner) GetSridOk() (*string, bool)`

GetSridOk returns a tuple with the Srid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrid

`func (o *GetOrders200ResponseInner) SetSrid(v string)`

SetSrid sets Srid field to given value.


### GetOrderType

`func (o *GetOrders200ResponseInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *GetOrders200ResponseInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *GetOrders200ResponseInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *GetOrders200ResponseInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


