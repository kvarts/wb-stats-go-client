/*
WB Supplier Stats

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetOrders200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrders200ResponseInner{}

// GetOrders200ResponseInner struct for GetOrders200ResponseInner
type GetOrders200ResponseInner struct {
	// номер заказа
	Number int64 `json:"number"`
	// дата заказа
	Date string `json:"date"`
	// дата время обновления информации в сервисе
	LastChangeDate string `json:"lastChangeDate"`
	// ваш артикул
	SupplierArticle string `json:"supplierArticle"`
	// размер
	TechSize string `json:"techSize"`
	// штрих-код
	Barcode string `json:"barcode"`
	// цена до согласованной скидки/промо/спп
	TotalPrice float32 `json:"totalPrice"`
	// согласованный итоговый дисконт
	DiscountPercent float32 `json:"discountPercent"`
	// склад отгрузки
	WarehouseName string `json:"warehouseName"`
	// номер поставки
	IncomeID int64 `json:"incomeID"`
	// Код WB
	NmId int32 `json:"nmId"`
	// предмет
	Subject string `json:"subject"`
	// категория
	Category string `json:"category"`
	// бренд
	Brand string `json:"brand"`
	// признак отмены заказа (0 – отмены не было, 1 – отмена была
	IsCancel bool `json:"isCancel"`
	// Номер заказа
	GNumber string `json:"gNumber"`
	// Srid
	Srid string `json:"srid"`
	// Тип поступившего заказа
	OrderType *string `json:"orderType,omitempty"`
	// область
	RegionName string `json:"regionName"`
	// дата отмены заказа
	CancelDate string `json:"cancelDate"`
	// Скидка WB
	Spp float32 `json:"spp"`
	// Фактическая цена с учетом всех скидок
	FinishedPrice float32 `json:"finishedPrice"`
	// Цена со скидкой продавца
	PriceWithDisc float32 `json:"priceWithDisc"`
	// Страна
	CountryName string `json:"countryName"`
	// Округ
	OblastOkrugName string `json:"oblastOkrugName"`
	// Договор поставки
	IsSupply bool `json:"isSupply"`
	// Договор реализации
	IsRealization bool `json:"isRealization"`
}

// NewGetOrders200ResponseInner instantiates a new GetOrders200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrders200ResponseInner(number int64, date string, lastChangeDate string, supplierArticle string, techSize string, barcode string, totalPrice float32, discountPercent float32, warehouseName string, incomeID int64, nmId int32, subject string, category string, brand string, isCancel bool, gNumber string, srid string, regionName string, cancelDate string, spp float32, finishedPrice float32, priceWithDisc float32, countryName string, oblastOkrugName string, isSupply bool, isRealization bool) *GetOrders200ResponseInner {
	this := GetOrders200ResponseInner{}
	this.Number = number
	this.Date = date
	this.LastChangeDate = lastChangeDate
	this.SupplierArticle = supplierArticle
	this.TechSize = techSize
	this.Barcode = barcode
	this.TotalPrice = totalPrice
	this.DiscountPercent = discountPercent
	this.WarehouseName = warehouseName
	this.IncomeID = incomeID
	this.NmId = nmId
	this.Subject = subject
	this.Category = category
	this.Brand = brand
	this.IsCancel = isCancel
	this.GNumber = gNumber
	this.Srid = srid
	this.RegionName = regionName
	this.CancelDate = cancelDate
	this.Spp = spp
	this.FinishedPrice = finishedPrice
	this.PriceWithDisc = priceWithDisc
	this.CountryName = countryName
	this.OblastOkrugName = oblastOkrugName
	this.IsSupply = isSupply
	this.IsRealization = isRealization
	return &this
}

// NewGetOrders200ResponseInnerWithDefaults instantiates a new GetOrders200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrders200ResponseInnerWithDefaults() *GetOrders200ResponseInner {
	this := GetOrders200ResponseInner{}
	return &this
}

// GetNumber returns the Number field value
func (o *GetOrders200ResponseInner) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *GetOrders200ResponseInner) SetNumber(v int64) {
	o.Number = v
}

// GetDate returns the Date field value
func (o *GetOrders200ResponseInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetOrders200ResponseInner) SetDate(v string) {
	o.Date = v
}

// GetLastChangeDate returns the LastChangeDate field value
func (o *GetOrders200ResponseInner) GetLastChangeDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastChangeDate
}

// GetLastChangeDateOk returns a tuple with the LastChangeDate field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetLastChangeDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastChangeDate, true
}

// SetLastChangeDate sets field value
func (o *GetOrders200ResponseInner) SetLastChangeDate(v string) {
	o.LastChangeDate = v
}

// GetSupplierArticle returns the SupplierArticle field value
func (o *GetOrders200ResponseInner) GetSupplierArticle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupplierArticle
}

// GetSupplierArticleOk returns a tuple with the SupplierArticle field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetSupplierArticleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupplierArticle, true
}

// SetSupplierArticle sets field value
func (o *GetOrders200ResponseInner) SetSupplierArticle(v string) {
	o.SupplierArticle = v
}

// GetTechSize returns the TechSize field value
func (o *GetOrders200ResponseInner) GetTechSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TechSize
}

// GetTechSizeOk returns a tuple with the TechSize field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetTechSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TechSize, true
}

// SetTechSize sets field value
func (o *GetOrders200ResponseInner) SetTechSize(v string) {
	o.TechSize = v
}

// GetBarcode returns the Barcode field value
func (o *GetOrders200ResponseInner) GetBarcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetBarcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Barcode, true
}

// SetBarcode sets field value
func (o *GetOrders200ResponseInner) SetBarcode(v string) {
	o.Barcode = v
}

// GetTotalPrice returns the TotalPrice field value
func (o *GetOrders200ResponseInner) GetTotalPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetTotalPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *GetOrders200ResponseInner) SetTotalPrice(v float32) {
	o.TotalPrice = v
}

// GetDiscountPercent returns the DiscountPercent field value
func (o *GetOrders200ResponseInner) GetDiscountPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiscountPercent
}

// GetDiscountPercentOk returns a tuple with the DiscountPercent field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetDiscountPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountPercent, true
}

// SetDiscountPercent sets field value
func (o *GetOrders200ResponseInner) SetDiscountPercent(v float32) {
	o.DiscountPercent = v
}

// GetWarehouseName returns the WarehouseName field value
func (o *GetOrders200ResponseInner) GetWarehouseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WarehouseName
}

// GetWarehouseNameOk returns a tuple with the WarehouseName field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetWarehouseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseName, true
}

// SetWarehouseName sets field value
func (o *GetOrders200ResponseInner) SetWarehouseName(v string) {
	o.WarehouseName = v
}

// GetIncomeID returns the IncomeID field value
func (o *GetOrders200ResponseInner) GetIncomeID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IncomeID
}

// GetIncomeIDOk returns a tuple with the IncomeID field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetIncomeIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncomeID, true
}

// SetIncomeID sets field value
func (o *GetOrders200ResponseInner) SetIncomeID(v int64) {
	o.IncomeID = v
}

// GetNmId returns the NmId field value
func (o *GetOrders200ResponseInner) GetNmId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NmId
}

// GetNmIdOk returns a tuple with the NmId field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetNmIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NmId, true
}

// SetNmId sets field value
func (o *GetOrders200ResponseInner) SetNmId(v int32) {
	o.NmId = v
}

// GetSubject returns the Subject field value
func (o *GetOrders200ResponseInner) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *GetOrders200ResponseInner) SetSubject(v string) {
	o.Subject = v
}

// GetCategory returns the Category field value
func (o *GetOrders200ResponseInner) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *GetOrders200ResponseInner) SetCategory(v string) {
	o.Category = v
}

// GetBrand returns the Brand field value
func (o *GetOrders200ResponseInner) GetBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Brand
}

// GetBrandOk returns a tuple with the Brand field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brand, true
}

// SetBrand sets field value
func (o *GetOrders200ResponseInner) SetBrand(v string) {
	o.Brand = v
}

// GetIsCancel returns the IsCancel field value
func (o *GetOrders200ResponseInner) GetIsCancel() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCancel
}

// GetIsCancelOk returns a tuple with the IsCancel field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetIsCancelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCancel, true
}

// SetIsCancel sets field value
func (o *GetOrders200ResponseInner) SetIsCancel(v bool) {
	o.IsCancel = v
}

// GetGNumber returns the GNumber field value
func (o *GetOrders200ResponseInner) GetGNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GNumber
}

// GetGNumberOk returns a tuple with the GNumber field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetGNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GNumber, true
}

// SetGNumber sets field value
func (o *GetOrders200ResponseInner) SetGNumber(v string) {
	o.GNumber = v
}

// GetSrid returns the Srid field value
func (o *GetOrders200ResponseInner) GetSrid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Srid
}

// GetSridOk returns a tuple with the Srid field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetSridOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Srid, true
}

// SetSrid sets field value
func (o *GetOrders200ResponseInner) SetSrid(v string) {
	o.Srid = v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *GetOrders200ResponseInner) GetOrderType() string {
	if o == nil || isNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || isNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *GetOrders200ResponseInner) HasOrderType() bool {
	if o != nil && !isNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *GetOrders200ResponseInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetRegionName returns the RegionName field value
func (o *GetOrders200ResponseInner) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *GetOrders200ResponseInner) SetRegionName(v string) {
	o.RegionName = v
}

// GetCancelDate returns the CancelDate field value
func (o *GetOrders200ResponseInner) GetCancelDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelDate
}

// GetCancelDateOk returns a tuple with the CancelDate field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetCancelDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelDate, true
}

// SetCancelDate sets field value
func (o *GetOrders200ResponseInner) SetCancelDate(v string) {
	o.CancelDate = v
}

// GetSpp returns the Spp field value
func (o *GetOrders200ResponseInner) GetSpp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Spp
}

// GetSppOk returns a tuple with the Spp field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetSppOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spp, true
}

// SetSpp sets field value
func (o *GetOrders200ResponseInner) SetSpp(v float32) {
	o.Spp = v
}

// GetFinishedPrice returns the FinishedPrice field value
func (o *GetOrders200ResponseInner) GetFinishedPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FinishedPrice
}

// GetFinishedPriceOk returns a tuple with the FinishedPrice field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetFinishedPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedPrice, true
}

// SetFinishedPrice sets field value
func (o *GetOrders200ResponseInner) SetFinishedPrice(v float32) {
	o.FinishedPrice = v
}

// GetPriceWithDisc returns the PriceWithDisc field value
func (o *GetOrders200ResponseInner) GetPriceWithDisc() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceWithDisc
}

// GetPriceWithDiscOk returns a tuple with the PriceWithDisc field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetPriceWithDiscOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceWithDisc, true
}

// SetPriceWithDisc sets field value
func (o *GetOrders200ResponseInner) SetPriceWithDisc(v float32) {
	o.PriceWithDisc = v
}

// GetCountryName returns the CountryName field value
func (o *GetOrders200ResponseInner) GetCountryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetCountryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryName, true
}

// SetCountryName sets field value
func (o *GetOrders200ResponseInner) SetCountryName(v string) {
	o.CountryName = v
}

// GetOblastOkrugName returns the OblastOkrugName field value
func (o *GetOrders200ResponseInner) GetOblastOkrugName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OblastOkrugName
}

// GetOblastOkrugNameOk returns a tuple with the OblastOkrugName field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetOblastOkrugNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OblastOkrugName, true
}

// SetOblastOkrugName sets field value
func (o *GetOrders200ResponseInner) SetOblastOkrugName(v string) {
	o.OblastOkrugName = v
}

// GetIsSupply returns the IsSupply field value
func (o *GetOrders200ResponseInner) GetIsSupply() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSupply
}

// GetIsSupplyOk returns a tuple with the IsSupply field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetIsSupplyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSupply, true
}

// SetIsSupply sets field value
func (o *GetOrders200ResponseInner) SetIsSupply(v bool) {
	o.IsSupply = v
}

// GetIsRealization returns the IsRealization field value
func (o *GetOrders200ResponseInner) GetIsRealization() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRealization
}

// GetIsRealizationOk returns a tuple with the IsRealization field value
// and a boolean to check if the value has been set.
func (o *GetOrders200ResponseInner) GetIsRealizationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRealization, true
}

// SetIsRealization sets field value
func (o *GetOrders200ResponseInner) SetIsRealization(v bool) {
	o.IsRealization = v
}

func (o GetOrders200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrders200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	toSerialize["date"] = o.Date
	toSerialize["lastChangeDate"] = o.LastChangeDate
	toSerialize["supplierArticle"] = o.SupplierArticle
	toSerialize["techSize"] = o.TechSize
	toSerialize["barcode"] = o.Barcode
	toSerialize["totalPrice"] = o.TotalPrice
	toSerialize["discountPercent"] = o.DiscountPercent
	toSerialize["warehouseName"] = o.WarehouseName
	toSerialize["incomeID"] = o.IncomeID
	toSerialize["nmId"] = o.NmId
	toSerialize["subject"] = o.Subject
	toSerialize["category"] = o.Category
	toSerialize["brand"] = o.Brand
	toSerialize["isCancel"] = o.IsCancel
	toSerialize["gNumber"] = o.GNumber
	toSerialize["srid"] = o.Srid
	if !isNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	toSerialize["regionName"] = o.RegionName
	toSerialize["cancelDate"] = o.CancelDate
	toSerialize["spp"] = o.Spp
	toSerialize["finishedPrice"] = o.FinishedPrice
	toSerialize["priceWithDisc"] = o.PriceWithDisc
	toSerialize["countryName"] = o.CountryName
	toSerialize["oblastOkrugName"] = o.OblastOkrugName
	toSerialize["isSupply"] = o.IsSupply
	toSerialize["isRealization"] = o.IsRealization
	return toSerialize, nil
}

type NullableGetOrders200ResponseInner struct {
	value *GetOrders200ResponseInner
	isSet bool
}

func (v NullableGetOrders200ResponseInner) Get() *GetOrders200ResponseInner {
	return v.value
}

func (v *NullableGetOrders200ResponseInner) Set(val *GetOrders200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrders200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrders200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrders200ResponseInner(val *GetOrders200ResponseInner) *NullableGetOrders200ResponseInner {
	return &NullableGetOrders200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrders200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrders200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


