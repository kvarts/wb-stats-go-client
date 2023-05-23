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

// checks if the GetPaidStorage200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPaidStorage200ResponseInner{}

// GetPaidStorage200ResponseInner 
type GetPaidStorage200ResponseInner struct {
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

// NewGetPaidStorage200ResponseInner instantiates a new GetPaidStorage200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaidStorage200ResponseInner(dayBeg string, dayEnd string, nmId int64, techSize string, daysOnSite int32, stock int64, saleQty int64, sumW float32) *GetPaidStorage200ResponseInner {
	this := GetPaidStorage200ResponseInner{}
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

// NewGetPaidStorage200ResponseInnerWithDefaults instantiates a new GetPaidStorage200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaidStorage200ResponseInnerWithDefaults() *GetPaidStorage200ResponseInner {
	this := GetPaidStorage200ResponseInner{}
	return &this
}

// GetDayBeg returns the DayBeg field value
func (o *GetPaidStorage200ResponseInner) GetDayBeg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayBeg
}

// GetDayBegOk returns a tuple with the DayBeg field value
// and a boolean to check if the value has been set.
func (o *GetPaidStorage200ResponseInner) GetDayBegOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayBeg, true
}

// SetDayBeg sets field value
func (o *GetPaidStorage200ResponseInner) SetDayBeg(v string) {
	o.DayBeg = v
}

// GetDayEnd returns the DayEnd field value
func (o *GetPaidStorage200ResponseInner) GetDayEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayEnd
}

// GetDayEndOk returns a tuple with the DayEnd field value
// and a boolean to check if the value has been set.
func (o *GetPaidStorage200ResponseInner) GetDayEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayEnd, true
}

// SetDayEnd sets field value
func (o *GetPaidStorage200ResponseInner) SetDayEnd(v string) {
	o.DayEnd = v
}

// GetNmId returns the NmId field value
func (o *GetPaidStorage200ResponseInner) GetNmId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NmId
}

// GetNmIdOk returns a tuple with the NmId field value
// and a boolean to check if the value has been set.
func (o *GetPaidStorage200ResponseInner) GetNmIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NmId, true
}

// SetNmId sets field value
func (o *GetPaidStorage200ResponseInner) SetNmId(v int64) {
	o.NmId = v
}

// GetTechSize returns the TechSize field value
func (o *GetPaidStorage200ResponseInner) GetTechSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TechSize
}

// GetTechSizeOk returns a tuple with the TechSize field value
// and a boolean to check if the value has been set.
func (o *GetPaidStorage200ResponseInner) GetTechSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TechSize, true
}

// SetTechSize sets field value
func (o *GetPaidStorage200ResponseInner) SetTechSize(v string) {
	o.TechSize = v
}

// GetDaysOnSite returns the DaysOnSite field value
func (o *GetPaidStorage200ResponseInner) GetDaysOnSite() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysOnSite
}

// GetDaysOnSiteOk returns a tuple with the DaysOnSite field value
// and a boolean to check if the value has been set.
func (o *GetPaidStorage200ResponseInner) GetDaysOnSiteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysOnSite, true
}

// SetDaysOnSite sets field value
func (o *GetPaidStorage200ResponseInner) SetDaysOnSite(v int32) {
	o.DaysOnSite = v
}

// GetStock returns the Stock field value
func (o *GetPaidStorage200ResponseInner) GetStock() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *GetPaidStorage200ResponseInner) GetStockOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *GetPaidStorage200ResponseInner) SetStock(v int64) {
	o.Stock = v
}

// GetSaleQty returns the SaleQty field value
func (o *GetPaidStorage200ResponseInner) GetSaleQty() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SaleQty
}

// GetSaleQtyOk returns a tuple with the SaleQty field value
// and a boolean to check if the value has been set.
func (o *GetPaidStorage200ResponseInner) GetSaleQtyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaleQty, true
}

// SetSaleQty sets field value
func (o *GetPaidStorage200ResponseInner) SetSaleQty(v int64) {
	o.SaleQty = v
}

// GetSumW returns the SumW field value
func (o *GetPaidStorage200ResponseInner) GetSumW() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SumW
}

// GetSumWOk returns a tuple with the SumW field value
// and a boolean to check if the value has been set.
func (o *GetPaidStorage200ResponseInner) GetSumWOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumW, true
}

// SetSumW sets field value
func (o *GetPaidStorage200ResponseInner) SetSumW(v float32) {
	o.SumW = v
}

func (o GetPaidStorage200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPaidStorage200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dayBeg"] = o.DayBeg
	toSerialize["dayEnd"] = o.DayEnd
	toSerialize["nmId"] = o.NmId
	toSerialize["techSize"] = o.TechSize
	toSerialize["daysOnSite"] = o.DaysOnSite
	toSerialize["stock"] = o.Stock
	toSerialize["saleQty"] = o.SaleQty
	toSerialize["sumW"] = o.SumW
	return toSerialize, nil
}

type NullableGetPaidStorage200ResponseInner struct {
	value *GetPaidStorage200ResponseInner
	isSet bool
}

func (v NullableGetPaidStorage200ResponseInner) Get() *GetPaidStorage200ResponseInner {
	return v.value
}

func (v *NullableGetPaidStorage200ResponseInner) Set(val *GetPaidStorage200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaidStorage200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaidStorage200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaidStorage200ResponseInner(val *GetPaidStorage200ResponseInner) *NullableGetPaidStorage200ResponseInner {
	return &NullableGetPaidStorage200ResponseInner{value: val, isSet: true}
}

func (v NullableGetPaidStorage200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPaidStorage200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


