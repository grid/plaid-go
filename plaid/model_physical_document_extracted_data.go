/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.214.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PhysicalDocumentExtractedData Data extracted from a user-submitted document.
type PhysicalDocumentExtractedData struct {
	// Alpha-numeric ID number extracted via OCR from the user's document image.
	IdNumber NullableString `json:"id_number"`
	Category PhysicalDocumentCategory `json:"category"`
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	ExpirationDate NullableString `json:"expiration_date"`
	// Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.
	IssuingCountry string `json:"issuing_country"`
	// An ISO 3166-2 subdivision code. Related terms would be \"state\", \"province\", \"prefecture\", \"zone\", \"subdivision\", etc.
	IssuingRegion NullableString `json:"issuing_region"`
}

// NewPhysicalDocumentExtractedData instantiates a new PhysicalDocumentExtractedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalDocumentExtractedData(idNumber NullableString, category PhysicalDocumentCategory, expirationDate NullableString, issuingCountry string, issuingRegion NullableString) *PhysicalDocumentExtractedData {
	this := PhysicalDocumentExtractedData{}
	this.IdNumber = idNumber
	this.Category = category
	this.ExpirationDate = expirationDate
	this.IssuingCountry = issuingCountry
	this.IssuingRegion = issuingRegion
	return &this
}

// NewPhysicalDocumentExtractedDataWithDefaults instantiates a new PhysicalDocumentExtractedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalDocumentExtractedDataWithDefaults() *PhysicalDocumentExtractedData {
	this := PhysicalDocumentExtractedData{}
	return &this
}

// GetIdNumber returns the IdNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PhysicalDocumentExtractedData) GetIdNumber() string {
	if o == nil || o.IdNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.IdNumber.Get()
}

// GetIdNumberOk returns a tuple with the IdNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalDocumentExtractedData) GetIdNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdNumber.Get(), o.IdNumber.IsSet()
}

// SetIdNumber sets field value
func (o *PhysicalDocumentExtractedData) SetIdNumber(v string) {
	o.IdNumber.Set(&v)
}

// GetCategory returns the Category field value
func (o *PhysicalDocumentExtractedData) GetCategory() PhysicalDocumentCategory {
	if o == nil {
		var ret PhysicalDocumentCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PhysicalDocumentExtractedData) GetCategoryOk() (*PhysicalDocumentCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *PhysicalDocumentExtractedData) SetCategory(v PhysicalDocumentCategory) {
	o.Category = v
}

// GetExpirationDate returns the ExpirationDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PhysicalDocumentExtractedData) GetExpirationDate() string {
	if o == nil || o.ExpirationDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalDocumentExtractedData) GetExpirationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// SetExpirationDate sets field value
func (o *PhysicalDocumentExtractedData) SetExpirationDate(v string) {
	o.ExpirationDate.Set(&v)
}

// GetIssuingCountry returns the IssuingCountry field value
func (o *PhysicalDocumentExtractedData) GetIssuingCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuingCountry
}

// GetIssuingCountryOk returns a tuple with the IssuingCountry field value
// and a boolean to check if the value has been set.
func (o *PhysicalDocumentExtractedData) GetIssuingCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuingCountry, true
}

// SetIssuingCountry sets field value
func (o *PhysicalDocumentExtractedData) SetIssuingCountry(v string) {
	o.IssuingCountry = v
}

// GetIssuingRegion returns the IssuingRegion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PhysicalDocumentExtractedData) GetIssuingRegion() string {
	if o == nil || o.IssuingRegion.Get() == nil {
		var ret string
		return ret
	}

	return *o.IssuingRegion.Get()
}

// GetIssuingRegionOk returns a tuple with the IssuingRegion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalDocumentExtractedData) GetIssuingRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IssuingRegion.Get(), o.IssuingRegion.IsSet()
}

// SetIssuingRegion sets field value
func (o *PhysicalDocumentExtractedData) SetIssuingRegion(v string) {
	o.IssuingRegion.Set(&v)
}

func (o PhysicalDocumentExtractedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id_number"] = o.IdNumber.Get()
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["expiration_date"] = o.ExpirationDate.Get()
	}
	if true {
		toSerialize["issuing_country"] = o.IssuingCountry
	}
	if true {
		toSerialize["issuing_region"] = o.IssuingRegion.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalDocumentExtractedData struct {
	value *PhysicalDocumentExtractedData
	isSet bool
}

func (v NullablePhysicalDocumentExtractedData) Get() *PhysicalDocumentExtractedData {
	return v.value
}

func (v *NullablePhysicalDocumentExtractedData) Set(val *PhysicalDocumentExtractedData) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalDocumentExtractedData) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalDocumentExtractedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalDocumentExtractedData(val *PhysicalDocumentExtractedData) *NullablePhysicalDocumentExtractedData {
	return &NullablePhysicalDocumentExtractedData{value: val, isSet: true}
}

func (v NullablePhysicalDocumentExtractedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalDocumentExtractedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


