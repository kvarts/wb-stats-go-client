# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | номер заказа | 
**Date** | **string** | дата заказа | 
**LastChangeDate** | **string** | дата время обновления информации в сервисе | 
**SupplierArticle** | **string** | ваш артикул | 
**TechSize** | **string** | размер | 
**Barcode** | **string** | штрих-код | 
**Quantity** | **int32** | кол-во | 
**TotalPrice** | **int32** | цена до согласованной скидки/промо/спп | 
**DiscountPercent** | **int32** | согласованный итоговый дисконт | 
**WarehouseName** | **string** | склад отгрузки | 
**Oblast** | **string** | область | 
**IncomeID** | **int32** | номер поставки | 
**Odid** | **int32** | уникальный идентификатор позиции заказа | 
**NmId** | **int32** | Код WB | 
**Subject** | **string** | предмет | 
**Category** | **string** | категория | 
**Brand** | **string** | бренд | 
**IsCancel** | **bool** | признак отмены заказа (0 – отмены не было, 1 – отмена была | 
**CancelDt** | **string** | дата отмены заказа | 
**GNumber** | **string** |  | 

## Methods

### NewInlineResponse200

`func NewInlineResponse200(number int32, date string, lastChangeDate string, supplierArticle string, techSize string, barcode string, quantity int32, totalPrice int32, discountPercent int32, warehouseName string, oblast string, incomeID int32, odid int32, nmId int32, subject string, category string, brand string, isCancel bool, cancelDt string, gNumber string, ) *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse200) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetDate

`func (o *InlineResponse200) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InlineResponse200) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InlineResponse200) SetDate(v string)`

SetDate sets Date field to given value.


### GetLastChangeDate

`func (o *InlineResponse200) GetLastChangeDate() string`

GetLastChangeDate returns the LastChangeDate field if non-nil, zero value otherwise.

### GetLastChangeDateOk

`func (o *InlineResponse200) GetLastChangeDateOk() (*string, bool)`

GetLastChangeDateOk returns a tuple with the LastChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeDate

`func (o *InlineResponse200) SetLastChangeDate(v string)`

SetLastChangeDate sets LastChangeDate field to given value.


### GetSupplierArticle

`func (o *InlineResponse200) GetSupplierArticle() string`

GetSupplierArticle returns the SupplierArticle field if non-nil, zero value otherwise.

### GetSupplierArticleOk

`func (o *InlineResponse200) GetSupplierArticleOk() (*string, bool)`

GetSupplierArticleOk returns a tuple with the SupplierArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierArticle

`func (o *InlineResponse200) SetSupplierArticle(v string)`

SetSupplierArticle sets SupplierArticle field to given value.


### GetTechSize

`func (o *InlineResponse200) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *InlineResponse200) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *InlineResponse200) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.


### GetBarcode

`func (o *InlineResponse200) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *InlineResponse200) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *InlineResponse200) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetQuantity

`func (o *InlineResponse200) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InlineResponse200) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InlineResponse200) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTotalPrice

`func (o *InlineResponse200) GetTotalPrice() int32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *InlineResponse200) GetTotalPriceOk() (*int32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *InlineResponse200) SetTotalPrice(v int32)`

SetTotalPrice sets TotalPrice field to given value.


### GetDiscountPercent

`func (o *InlineResponse200) GetDiscountPercent() int32`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *InlineResponse200) GetDiscountPercentOk() (*int32, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *InlineResponse200) SetDiscountPercent(v int32)`

SetDiscountPercent sets DiscountPercent field to given value.


### GetWarehouseName

`func (o *InlineResponse200) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *InlineResponse200) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *InlineResponse200) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### GetOblast

`func (o *InlineResponse200) GetOblast() string`

GetOblast returns the Oblast field if non-nil, zero value otherwise.

### GetOblastOk

`func (o *InlineResponse200) GetOblastOk() (*string, bool)`

GetOblastOk returns a tuple with the Oblast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOblast

`func (o *InlineResponse200) SetOblast(v string)`

SetOblast sets Oblast field to given value.


### GetIncomeID

`func (o *InlineResponse200) GetIncomeID() int32`

GetIncomeID returns the IncomeID field if non-nil, zero value otherwise.

### GetIncomeIDOk

`func (o *InlineResponse200) GetIncomeIDOk() (*int32, bool)`

GetIncomeIDOk returns a tuple with the IncomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeID

`func (o *InlineResponse200) SetIncomeID(v int32)`

SetIncomeID sets IncomeID field to given value.


### GetOdid

`func (o *InlineResponse200) GetOdid() int32`

GetOdid returns the Odid field if non-nil, zero value otherwise.

### GetOdidOk

`func (o *InlineResponse200) GetOdidOk() (*int32, bool)`

GetOdidOk returns a tuple with the Odid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdid

`func (o *InlineResponse200) SetOdid(v int32)`

SetOdid sets Odid field to given value.


### GetNmId

`func (o *InlineResponse200) GetNmId() int32`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *InlineResponse200) GetNmIdOk() (*int32, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *InlineResponse200) SetNmId(v int32)`

SetNmId sets NmId field to given value.


### GetSubject

`func (o *InlineResponse200) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InlineResponse200) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InlineResponse200) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCategory

`func (o *InlineResponse200) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse200) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse200) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetBrand

`func (o *InlineResponse200) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *InlineResponse200) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *InlineResponse200) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetIsCancel

`func (o *InlineResponse200) GetIsCancel() bool`

GetIsCancel returns the IsCancel field if non-nil, zero value otherwise.

### GetIsCancelOk

`func (o *InlineResponse200) GetIsCancelOk() (*bool, bool)`

GetIsCancelOk returns a tuple with the IsCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancel

`func (o *InlineResponse200) SetIsCancel(v bool)`

SetIsCancel sets IsCancel field to given value.


### GetCancelDt

`func (o *InlineResponse200) GetCancelDt() string`

GetCancelDt returns the CancelDt field if non-nil, zero value otherwise.

### GetCancelDtOk

`func (o *InlineResponse200) GetCancelDtOk() (*string, bool)`

GetCancelDtOk returns a tuple with the CancelDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDt

`func (o *InlineResponse200) SetCancelDt(v string)`

SetCancelDt sets CancelDt field to given value.


### GetGNumber

`func (o *InlineResponse200) GetGNumber() string`

GetGNumber returns the GNumber field if non-nil, zero value otherwise.

### GetGNumberOk

`func (o *InlineResponse200) GetGNumberOk() (*string, bool)`

GetGNumberOk returns a tuple with the GNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGNumber

`func (o *InlineResponse200) SetGNumber(v string)`

SetGNumber sets GNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


