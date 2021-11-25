# InlineResponse2002

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
**QuantityNotInOrders** | **int64** | Количество не в заказе | 
**WarehouseName** | **string** | Название склада | 
**InWayToClient** | **int32** | Количество штук в пути к клиенту | 
**InWayFromClient** | **int32** | Количество штук в пути от клиента | 
**NmId** | **int64** | Код wildberries | 
**Subject** | **string** | Предмет | 
**Category** | **string** | Категория | 
**DaysOnSite** | **int32** | Количество дней на сайте | 
**Brand** | **string** | Бренд | 
**SCCode** | **string** | Код контракта | 
**Price** | **float32** | Цена (?) | 
**Discount** | **float32** | Скидка (?) | 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002(lastChangeDate time.Time, supplierArticle string, techSize string, barcode string, quantity int64, isSupply bool, isRealization bool, quantityFull int64, quantityNotInOrders int64, warehouseName string, inWayToClient int32, inWayFromClient int32, nmId int64, subject string, category string, daysOnSite int32, brand string, sCCode string, price float32, discount float32, ) *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChangeDate

`func (o *InlineResponse2002) GetLastChangeDate() time.Time`

GetLastChangeDate returns the LastChangeDate field if non-nil, zero value otherwise.

### GetLastChangeDateOk

`func (o *InlineResponse2002) GetLastChangeDateOk() (*time.Time, bool)`

GetLastChangeDateOk returns a tuple with the LastChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeDate

`func (o *InlineResponse2002) SetLastChangeDate(v time.Time)`

SetLastChangeDate sets LastChangeDate field to given value.


### GetSupplierArticle

`func (o *InlineResponse2002) GetSupplierArticle() string`

GetSupplierArticle returns the SupplierArticle field if non-nil, zero value otherwise.

### GetSupplierArticleOk

`func (o *InlineResponse2002) GetSupplierArticleOk() (*string, bool)`

GetSupplierArticleOk returns a tuple with the SupplierArticle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierArticle

`func (o *InlineResponse2002) SetSupplierArticle(v string)`

SetSupplierArticle sets SupplierArticle field to given value.


### GetTechSize

`func (o *InlineResponse2002) GetTechSize() string`

GetTechSize returns the TechSize field if non-nil, zero value otherwise.

### GetTechSizeOk

`func (o *InlineResponse2002) GetTechSizeOk() (*string, bool)`

GetTechSizeOk returns a tuple with the TechSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSize

`func (o *InlineResponse2002) SetTechSize(v string)`

SetTechSize sets TechSize field to given value.


### GetBarcode

`func (o *InlineResponse2002) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *InlineResponse2002) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *InlineResponse2002) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetQuantity

`func (o *InlineResponse2002) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InlineResponse2002) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InlineResponse2002) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetIsSupply

`func (o *InlineResponse2002) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *InlineResponse2002) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *InlineResponse2002) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.


### GetIsRealization

`func (o *InlineResponse2002) GetIsRealization() bool`

GetIsRealization returns the IsRealization field if non-nil, zero value otherwise.

### GetIsRealizationOk

`func (o *InlineResponse2002) GetIsRealizationOk() (*bool, bool)`

GetIsRealizationOk returns a tuple with the IsRealization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealization

`func (o *InlineResponse2002) SetIsRealization(v bool)`

SetIsRealization sets IsRealization field to given value.


### GetQuantityFull

`func (o *InlineResponse2002) GetQuantityFull() int64`

GetQuantityFull returns the QuantityFull field if non-nil, zero value otherwise.

### GetQuantityFullOk

`func (o *InlineResponse2002) GetQuantityFullOk() (*int64, bool)`

GetQuantityFullOk returns a tuple with the QuantityFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityFull

`func (o *InlineResponse2002) SetQuantityFull(v int64)`

SetQuantityFull sets QuantityFull field to given value.


### GetQuantityNotInOrders

`func (o *InlineResponse2002) GetQuantityNotInOrders() int64`

GetQuantityNotInOrders returns the QuantityNotInOrders field if non-nil, zero value otherwise.

### GetQuantityNotInOrdersOk

`func (o *InlineResponse2002) GetQuantityNotInOrdersOk() (*int64, bool)`

GetQuantityNotInOrdersOk returns a tuple with the QuantityNotInOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityNotInOrders

`func (o *InlineResponse2002) SetQuantityNotInOrders(v int64)`

SetQuantityNotInOrders sets QuantityNotInOrders field to given value.


### GetWarehouseName

`func (o *InlineResponse2002) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *InlineResponse2002) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *InlineResponse2002) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### GetInWayToClient

`func (o *InlineResponse2002) GetInWayToClient() int32`

GetInWayToClient returns the InWayToClient field if non-nil, zero value otherwise.

### GetInWayToClientOk

`func (o *InlineResponse2002) GetInWayToClientOk() (*int32, bool)`

GetInWayToClientOk returns a tuple with the InWayToClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWayToClient

`func (o *InlineResponse2002) SetInWayToClient(v int32)`

SetInWayToClient sets InWayToClient field to given value.


### GetInWayFromClient

`func (o *InlineResponse2002) GetInWayFromClient() int32`

GetInWayFromClient returns the InWayFromClient field if non-nil, zero value otherwise.

### GetInWayFromClientOk

`func (o *InlineResponse2002) GetInWayFromClientOk() (*int32, bool)`

GetInWayFromClientOk returns a tuple with the InWayFromClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWayFromClient

`func (o *InlineResponse2002) SetInWayFromClient(v int32)`

SetInWayFromClient sets InWayFromClient field to given value.


### GetNmId

`func (o *InlineResponse2002) GetNmId() int64`

GetNmId returns the NmId field if non-nil, zero value otherwise.

### GetNmIdOk

`func (o *InlineResponse2002) GetNmIdOk() (*int64, bool)`

GetNmIdOk returns a tuple with the NmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmId

`func (o *InlineResponse2002) SetNmId(v int64)`

SetNmId sets NmId field to given value.


### GetSubject

`func (o *InlineResponse2002) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InlineResponse2002) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InlineResponse2002) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCategory

`func (o *InlineResponse2002) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse2002) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse2002) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDaysOnSite

`func (o *InlineResponse2002) GetDaysOnSite() int32`

GetDaysOnSite returns the DaysOnSite field if non-nil, zero value otherwise.

### GetDaysOnSiteOk

`func (o *InlineResponse2002) GetDaysOnSiteOk() (*int32, bool)`

GetDaysOnSiteOk returns a tuple with the DaysOnSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOnSite

`func (o *InlineResponse2002) SetDaysOnSite(v int32)`

SetDaysOnSite sets DaysOnSite field to given value.


### GetBrand

`func (o *InlineResponse2002) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *InlineResponse2002) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *InlineResponse2002) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetSCCode

`func (o *InlineResponse2002) GetSCCode() string`

GetSCCode returns the SCCode field if non-nil, zero value otherwise.

### GetSCCodeOk

`func (o *InlineResponse2002) GetSCCodeOk() (*string, bool)`

GetSCCodeOk returns a tuple with the SCCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCCode

`func (o *InlineResponse2002) SetSCCode(v string)`

SetSCCode sets SCCode field to given value.


### GetPrice

`func (o *InlineResponse2002) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse2002) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse2002) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetDiscount

`func (o *InlineResponse2002) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *InlineResponse2002) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *InlineResponse2002) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


