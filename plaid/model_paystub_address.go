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

// PaystubAddress Address on the paystub
type PaystubAddress struct {
	// The full city name.
	City NullableString `json:"city,omitempty"`
	// The ISO 3166-1 alpha-2 country code.
	Country NullableString `json:"country,omitempty"`
	// The postal code of the address.
	PostalCode NullableString `json:"postal_code,omitempty"`
	// The region or state Example: `\"NC\"`
	Region NullableString `json:"region,omitempty"`
	// The full street address.
	Street NullableString `json:"street,omitempty"`
	// Street address line 1.
	Line1 NullableString `json:"line1,omitempty"`
	// Street address line 2.
	Line2 NullableString `json:"line2,omitempty"`
	// The region or state Example: `\"NC\"`
	StateCode NullableString `json:"state_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaystubAddress PaystubAddress

// NewPaystubAddress instantiates a new PaystubAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubAddress() *PaystubAddress {
	this := PaystubAddress{}
	return &this
}

// NewPaystubAddressWithDefaults instantiates a new PaystubAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubAddressWithDefaults() *PaystubAddress {
	this := PaystubAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubAddress) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubAddress) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *PaystubAddress) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *PaystubAddress) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *PaystubAddress) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *PaystubAddress) UnsetCity() {
	o.City.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubAddress) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubAddress) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *PaystubAddress) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *PaystubAddress) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *PaystubAddress) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *PaystubAddress) UnsetCountry() {
	o.Country.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubAddress) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PaystubAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *PaystubAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *PaystubAddress) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *PaystubAddress) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubAddress) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubAddress) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *PaystubAddress) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *PaystubAddress) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *PaystubAddress) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *PaystubAddress) UnsetRegion() {
	o.Region.Unset()
}

// GetStreet returns the Street field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubAddress) GetStreet() string {
	if o == nil || o.Street.Get() == nil {
		var ret string
		return ret
	}
	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubAddress) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// HasStreet returns a boolean if a field has been set.
func (o *PaystubAddress) HasStreet() bool {
	if o != nil && o.Street.IsSet() {
		return true
	}

	return false
}

// SetStreet gets a reference to the given NullableString and assigns it to the Street field.
func (o *PaystubAddress) SetStreet(v string) {
	o.Street.Set(&v)
}
// SetStreetNil sets the value for Street to be an explicit nil
func (o *PaystubAddress) SetStreetNil() {
	o.Street.Set(nil)
}

// UnsetStreet ensures that no value is present for Street, not even an explicit nil
func (o *PaystubAddress) UnsetStreet() {
	o.Street.Unset()
}

// GetLine1 returns the Line1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubAddress) GetLine1() string {
	if o == nil || o.Line1.Get() == nil {
		var ret string
		return ret
	}
	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubAddress) GetLine1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// HasLine1 returns a boolean if a field has been set.
func (o *PaystubAddress) HasLine1() bool {
	if o != nil && o.Line1.IsSet() {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given NullableString and assigns it to the Line1 field.
func (o *PaystubAddress) SetLine1(v string) {
	o.Line1.Set(&v)
}
// SetLine1Nil sets the value for Line1 to be an explicit nil
func (o *PaystubAddress) SetLine1Nil() {
	o.Line1.Set(nil)
}

// UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
func (o *PaystubAddress) UnsetLine1() {
	o.Line1.Unset()
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubAddress) GetLine2() string {
	if o == nil || o.Line2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubAddress) GetLine2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// HasLine2 returns a boolean if a field has been set.
func (o *PaystubAddress) HasLine2() bool {
	if o != nil && o.Line2.IsSet() {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given NullableString and assigns it to the Line2 field.
func (o *PaystubAddress) SetLine2(v string) {
	o.Line2.Set(&v)
}
// SetLine2Nil sets the value for Line2 to be an explicit nil
func (o *PaystubAddress) SetLine2Nil() {
	o.Line2.Set(nil)
}

// UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
func (o *PaystubAddress) UnsetLine2() {
	o.Line2.Unset()
}

// GetStateCode returns the StateCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaystubAddress) GetStateCode() string {
	if o == nil || o.StateCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateCode.Get()
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubAddress) GetStateCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StateCode.Get(), o.StateCode.IsSet()
}

// HasStateCode returns a boolean if a field has been set.
func (o *PaystubAddress) HasStateCode() bool {
	if o != nil && o.StateCode.IsSet() {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given NullableString and assigns it to the StateCode field.
func (o *PaystubAddress) SetStateCode(v string) {
	o.StateCode.Set(&v)
}
// SetStateCodeNil sets the value for StateCode to be an explicit nil
func (o *PaystubAddress) SetStateCodeNil() {
	o.StateCode.Set(nil)
}

// UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
func (o *PaystubAddress) UnsetStateCode() {
	o.StateCode.Unset()
}

func (o PaystubAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Street.IsSet() {
		toSerialize["street"] = o.Street.Get()
	}
	if o.Line1.IsSet() {
		toSerialize["line1"] = o.Line1.Get()
	}
	if o.Line2.IsSet() {
		toSerialize["line2"] = o.Line2.Get()
	}
	if o.StateCode.IsSet() {
		toSerialize["state_code"] = o.StateCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaystubAddress) UnmarshalJSON(bytes []byte) (err error) {
	varPaystubAddress := _PaystubAddress{}

	if err = json.Unmarshal(bytes, &varPaystubAddress); err == nil {
		*o = PaystubAddress(varPaystubAddress)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "country")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		delete(additionalProperties, "line1")
		delete(additionalProperties, "line2")
		delete(additionalProperties, "state_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaystubAddress struct {
	value *PaystubAddress
	isSet bool
}

func (v NullablePaystubAddress) Get() *PaystubAddress {
	return v.value
}

func (v *NullablePaystubAddress) Set(val *PaystubAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubAddress(val *PaystubAddress) *NullablePaystubAddress {
	return &NullablePaystubAddress{value: val, isSet: true}
}

func (v NullablePaystubAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


