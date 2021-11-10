/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.46.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SecurityOverride Specify the security associated with the holding or investment transaction. When inputting custom security data to the Sandbox, Plaid will perform post-data-retrieval normalization and enrichment. These processes may cause the data returned by the Sandbox to be slightly different from the data you input. An ISO-4217 currency code and a security identifier (`ticker_symbol`, `cusip`, `isin`, or `sedol`) are required.
type SecurityOverride struct {
	// 12-character ISIN, a globally unique securities identifier.
	Isin *string `json:"isin,omitempty"`
	// 9-character CUSIP, an identifier assigned to North American securities.
	Cusip *string `json:"cusip,omitempty"`
	// 7-character SEDOL, an identifier assigned to securities in the UK.
	Sedol *string `json:"sedol,omitempty"`
	// A descriptive name for the security, suitable for display.
	Name *string `json:"name,omitempty"`
	// The security’s trading symbol for publicly traded securities, and otherwise a short identifier if available.
	TickerSymbol *string `json:"ticker_symbol,omitempty"`
	// Either a valid `iso_currency_code` or `unofficial_currency_code`
	Currency *string `json:"currency,omitempty"`
}

// NewSecurityOverride instantiates a new SecurityOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityOverride() *SecurityOverride {
	this := SecurityOverride{}
	return &this
}

// NewSecurityOverrideWithDefaults instantiates a new SecurityOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityOverrideWithDefaults() *SecurityOverride {
	this := SecurityOverride{}
	return &this
}

// GetIsin returns the Isin field value if set, zero value otherwise.
func (o *SecurityOverride) GetIsin() string {
	if o == nil || o.Isin == nil {
		var ret string
		return ret
	}
	return *o.Isin
}

// GetIsinOk returns a tuple with the Isin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOverride) GetIsinOk() (*string, bool) {
	if o == nil || o.Isin == nil {
		return nil, false
	}
	return o.Isin, true
}

// HasIsin returns a boolean if a field has been set.
func (o *SecurityOverride) HasIsin() bool {
	if o != nil && o.Isin != nil {
		return true
	}

	return false
}

// SetIsin gets a reference to the given string and assigns it to the Isin field.
func (o *SecurityOverride) SetIsin(v string) {
	o.Isin = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *SecurityOverride) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOverride) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *SecurityOverride) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *SecurityOverride) SetCusip(v string) {
	o.Cusip = &v
}

// GetSedol returns the Sedol field value if set, zero value otherwise.
func (o *SecurityOverride) GetSedol() string {
	if o == nil || o.Sedol == nil {
		var ret string
		return ret
	}
	return *o.Sedol
}

// GetSedolOk returns a tuple with the Sedol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOverride) GetSedolOk() (*string, bool) {
	if o == nil || o.Sedol == nil {
		return nil, false
	}
	return o.Sedol, true
}

// HasSedol returns a boolean if a field has been set.
func (o *SecurityOverride) HasSedol() bool {
	if o != nil && o.Sedol != nil {
		return true
	}

	return false
}

// SetSedol gets a reference to the given string and assigns it to the Sedol field.
func (o *SecurityOverride) SetSedol(v string) {
	o.Sedol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityOverride) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOverride) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityOverride) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityOverride) SetName(v string) {
	o.Name = &v
}

// GetTickerSymbol returns the TickerSymbol field value if set, zero value otherwise.
func (o *SecurityOverride) GetTickerSymbol() string {
	if o == nil || o.TickerSymbol == nil {
		var ret string
		return ret
	}
	return *o.TickerSymbol
}

// GetTickerSymbolOk returns a tuple with the TickerSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOverride) GetTickerSymbolOk() (*string, bool) {
	if o == nil || o.TickerSymbol == nil {
		return nil, false
	}
	return o.TickerSymbol, true
}

// HasTickerSymbol returns a boolean if a field has been set.
func (o *SecurityOverride) HasTickerSymbol() bool {
	if o != nil && o.TickerSymbol != nil {
		return true
	}

	return false
}

// SetTickerSymbol gets a reference to the given string and assigns it to the TickerSymbol field.
func (o *SecurityOverride) SetTickerSymbol(v string) {
	o.TickerSymbol = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SecurityOverride) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityOverride) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SecurityOverride) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SecurityOverride) SetCurrency(v string) {
	o.Currency = &v
}

func (o SecurityOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Isin != nil {
		toSerialize["isin"] = o.Isin
	}
	if o.Cusip != nil {
		toSerialize["cusip"] = o.Cusip
	}
	if o.Sedol != nil {
		toSerialize["sedol"] = o.Sedol
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TickerSymbol != nil {
		toSerialize["ticker_symbol"] = o.TickerSymbol
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityOverride struct {
	value *SecurityOverride
	isSet bool
}

func (v NullableSecurityOverride) Get() *SecurityOverride {
	return v.value
}

func (v *NullableSecurityOverride) Set(val *SecurityOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityOverride(val *SecurityOverride) *NullableSecurityOverride {
	return &NullableSecurityOverride{value: val, isSet: true}
}

func (v NullableSecurityOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


