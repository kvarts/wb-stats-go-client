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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// Номер поставки
	IncomeId int64 `json:"incomeId"`
	// Номер УПД
	Number string `json:"number"`
	// Дата поступления
	Date string `json:"date"`
	// Дата и время обновления данных в сервисе
	LastChangeDate string `json:"lastChangeDate"`
	// Артикул продавца
	SupplierArticle string `json:"supplierArticle"`
	// Размер
	TechSize string `json:"techSize"`
	// Штрих-код
	Barcode string `json:"barcode"`
	// Количество
	Quantity int32 `json:"quantity"`
	// Цена из УПД
	TotalPrice float32 `json:"totalPrice"`
	// Дата принятия (закрытия) у wildberies
	DateClose string `json:"dateClose"`
	// Наименование склада
	WarehouseName string `json:"warehouseName"`
	// Код WB
	NmId int64 `json:"nmId"`
	// Текущий статус поставки
	Status string `json:"status"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001(incomeId int64, number string, date string, lastChangeDate string, supplierArticle string, techSize string, barcode string, quantity int32, totalPrice float32, dateClose string, warehouseName string, nmId int64, status string) *InlineResponse2001 {
	this := InlineResponse2001{}
	this.IncomeId = incomeId
	this.Number = number
	this.Date = date
	this.LastChangeDate = lastChangeDate
	this.SupplierArticle = supplierArticle
	this.TechSize = techSize
	this.Barcode = barcode
	this.Quantity = quantity
	this.TotalPrice = totalPrice
	this.DateClose = dateClose
	this.WarehouseName = warehouseName
	this.NmId = nmId
	this.Status = status
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetIncomeId returns the IncomeId field value
func (o *InlineResponse2001) GetIncomeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IncomeId
}

// GetIncomeIdOk returns a tuple with the IncomeId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetIncomeIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncomeId, true
}

// SetIncomeId sets field value
func (o *InlineResponse2001) SetIncomeId(v int64) {
	o.IncomeId = v
}

// GetNumber returns the Number field value
func (o *InlineResponse2001) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *InlineResponse2001) SetNumber(v string) {
	o.Number = v
}

// GetDate returns the Date field value
func (o *InlineResponse2001) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *InlineResponse2001) SetDate(v string) {
	o.Date = v
}

// GetLastChangeDate returns the LastChangeDate field value
func (o *InlineResponse2001) GetLastChangeDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastChangeDate
}

// GetLastChangeDateOk returns a tuple with the LastChangeDate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetLastChangeDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastChangeDate, true
}

// SetLastChangeDate sets field value
func (o *InlineResponse2001) SetLastChangeDate(v string) {
	o.LastChangeDate = v
}

// GetSupplierArticle returns the SupplierArticle field value
func (o *InlineResponse2001) GetSupplierArticle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupplierArticle
}

// GetSupplierArticleOk returns a tuple with the SupplierArticle field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSupplierArticleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupplierArticle, true
}

// SetSupplierArticle sets field value
func (o *InlineResponse2001) SetSupplierArticle(v string) {
	o.SupplierArticle = v
}

// GetTechSize returns the TechSize field value
func (o *InlineResponse2001) GetTechSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TechSize
}

// GetTechSizeOk returns a tuple with the TechSize field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTechSizeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TechSize, true
}

// SetTechSize sets field value
func (o *InlineResponse2001) SetTechSize(v string) {
	o.TechSize = v
}

// GetBarcode returns the Barcode field value
func (o *InlineResponse2001) GetBarcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetBarcodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Barcode, true
}

// SetBarcode sets field value
func (o *InlineResponse2001) SetBarcode(v string) {
	o.Barcode = v
}

// GetQuantity returns the Quantity field value
func (o *InlineResponse2001) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetQuantityOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InlineResponse2001) SetQuantity(v int32) {
	o.Quantity = v
}

// GetTotalPrice returns the TotalPrice field value
func (o *InlineResponse2001) GetTotalPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTotalPriceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *InlineResponse2001) SetTotalPrice(v float32) {
	o.TotalPrice = v
}

// GetDateClose returns the DateClose field value
func (o *InlineResponse2001) GetDateClose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateClose
}

// GetDateCloseOk returns a tuple with the DateClose field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDateCloseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateClose, true
}

// SetDateClose sets field value
func (o *InlineResponse2001) SetDateClose(v string) {
	o.DateClose = v
}

// GetWarehouseName returns the WarehouseName field value
func (o *InlineResponse2001) GetWarehouseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WarehouseName
}

// GetWarehouseNameOk returns a tuple with the WarehouseName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetWarehouseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WarehouseName, true
}

// SetWarehouseName sets field value
func (o *InlineResponse2001) SetWarehouseName(v string) {
	o.WarehouseName = v
}

// GetNmId returns the NmId field value
func (o *InlineResponse2001) GetNmId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NmId
}

// GetNmIdOk returns a tuple with the NmId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetNmIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NmId, true
}

// SetNmId sets field value
func (o *InlineResponse2001) SetNmId(v int64) {
	o.NmId = v
}

// GetStatus returns the Status field value
func (o *InlineResponse2001) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse2001) SetStatus(v string) {
	o.Status = v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["incomeId"] = o.IncomeId
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["lastChangeDate"] = o.LastChangeDate
	}
	if true {
		toSerialize["supplierArticle"] = o.SupplierArticle
	}
	if true {
		toSerialize["techSize"] = o.TechSize
	}
	if true {
		toSerialize["barcode"] = o.Barcode
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["totalPrice"] = o.TotalPrice
	}
	if true {
		toSerialize["dateClose"] = o.DateClose
	}
	if true {
		toSerialize["warehouseName"] = o.WarehouseName
	}
	if true {
		toSerialize["nmId"] = o.NmId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


