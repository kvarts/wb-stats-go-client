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

// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	// Дата начала периода хранения
	DayBeg string `json:"dayBeg"`
	// Дата завершения периода хранения
	DayEnd string `json:"dayEnd"`
	// Код wildberries
	NmId int64 `json:"nmId"`
	// Размер
	TechSize string `json:"techSize"`
	// Количество дней на сайте
	DaysOnSite int32 `json:"daysOnSite"`
	// Остаток на начало периода
	Stock int64 `json:"stock"`
	// Количество продано
	SaleQty int64 `json:"saleQty"`
	// Сумма за хранение, руб.
	SumW float32 `json:"sumW"`
}

// NewInlineResponse2004 instantiates a new InlineResponse2004 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004(dayBeg string, dayEnd string, nmId int64, techSize string, daysOnSite int32, stock int64, saleQty int64, sumW float32) *InlineResponse2004 {
	this := InlineResponse2004{}
	this.DayBeg = dayBeg
	this.DayEnd = dayEnd
	this.NmId = nmId
	this.TechSize = techSize
	this.DaysOnSite = daysOnSite
	this.Stock = stock
	this.SaleQty = saleQty
	this.SumW = sumW
	return &this
}

// NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004WithDefaults() *InlineResponse2004 {
	this := InlineResponse2004{}
	return &this
}

// GetDayBeg returns the DayBeg field value
func (o *InlineResponse2004) GetDayBeg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayBeg
}

// GetDayBegOk returns a tuple with the DayBeg field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetDayBegOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DayBeg, true
}

// SetDayBeg sets field value
func (o *InlineResponse2004) SetDayBeg(v string) {
	o.DayBeg = v
}

// GetDayEnd returns the DayEnd field value
func (o *InlineResponse2004) GetDayEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayEnd
}

// GetDayEndOk returns a tuple with the DayEnd field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetDayEndOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DayEnd, true
}

// SetDayEnd sets field value
func (o *InlineResponse2004) SetDayEnd(v string) {
	o.DayEnd = v
}

// GetNmId returns the NmId field value
func (o *InlineResponse2004) GetNmId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NmId
}

// GetNmIdOk returns a tuple with the NmId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetNmIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NmId, true
}

// SetNmId sets field value
func (o *InlineResponse2004) SetNmId(v int64) {
	o.NmId = v
}

// GetTechSize returns the TechSize field value
func (o *InlineResponse2004) GetTechSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TechSize
}

// GetTechSizeOk returns a tuple with the TechSize field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetTechSizeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TechSize, true
}

// SetTechSize sets field value
func (o *InlineResponse2004) SetTechSize(v string) {
	o.TechSize = v
}

// GetDaysOnSite returns the DaysOnSite field value
func (o *InlineResponse2004) GetDaysOnSite() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysOnSite
}

// GetDaysOnSiteOk returns a tuple with the DaysOnSite field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetDaysOnSiteOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DaysOnSite, true
}

// SetDaysOnSite sets field value
func (o *InlineResponse2004) SetDaysOnSite(v int32) {
	o.DaysOnSite = v
}

// GetStock returns the Stock field value
func (o *InlineResponse2004) GetStock() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetStockOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *InlineResponse2004) SetStock(v int64) {
	o.Stock = v
}

// GetSaleQty returns the SaleQty field value
func (o *InlineResponse2004) GetSaleQty() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SaleQty
}

// GetSaleQtyOk returns a tuple with the SaleQty field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetSaleQtyOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SaleQty, true
}

// SetSaleQty sets field value
func (o *InlineResponse2004) SetSaleQty(v int64) {
	o.SaleQty = v
}

// GetSumW returns the SumW field value
func (o *InlineResponse2004) GetSumW() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SumW
}

// GetSumWOk returns a tuple with the SumW field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetSumWOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SumW, true
}

// SetSumW sets field value
func (o *InlineResponse2004) SetSumW(v float32) {
	o.SumW = v
}

func (o InlineResponse2004) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dayBeg"] = o.DayBeg
	}
	if true {
		toSerialize["dayEnd"] = o.DayEnd
	}
	if true {
		toSerialize["nmId"] = o.NmId
	}
	if true {
		toSerialize["techSize"] = o.TechSize
	}
	if true {
		toSerialize["daysOnSite"] = o.DaysOnSite
	}
	if true {
		toSerialize["stock"] = o.Stock
	}
	if true {
		toSerialize["saleQty"] = o.SaleQty
	}
	if true {
		toSerialize["sumW"] = o.SumW
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2004 struct {
	value *InlineResponse2004
	isSet bool
}

func (v NullableInlineResponse2004) Get() *InlineResponse2004 {
	return v.value
}

func (v *NullableInlineResponse2004) Set(val *InlineResponse2004) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2004) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2004) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2004(val *InlineResponse2004) *NullableInlineResponse2004 {
	return &NullableInlineResponse2004{value: val, isSet: true}
}

func (v NullableInlineResponse2004) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2004) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


