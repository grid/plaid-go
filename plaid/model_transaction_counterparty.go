/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionCounterparty The counterparty, such as the merchant or financial institution, is extracted by Plaid from the raw description.
type TransactionCounterparty struct {
	// The name of the counterparty, such as the merchant or the financial institution, as extracted by Plaid from the raw description.
	Name string `json:"name"`
	Type CounterpartyType `json:"type"`
	// The website associated with the counterparty.
	Website NullableString `json:"website"`
	// The URL of a logo associated with the counterparty, if available. The logo is formatted as a 100x100 pixel PNG filepath.
	LogoUrl NullableString `json:"logo_url"`
	AdditionalProperties map[string]interface{}
}

type _TransactionCounterparty TransactionCounterparty

// NewTransactionCounterparty instantiates a new TransactionCounterparty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCounterparty(name string, type_ CounterpartyType, website NullableString, logoUrl NullableString) *TransactionCounterparty {
	this := TransactionCounterparty{}
	this.Name = name
	this.Type = type_
	this.Website = website
	this.LogoUrl = logoUrl
	return &this
}

// NewTransactionCounterpartyWithDefaults instantiates a new TransactionCounterparty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCounterpartyWithDefaults() *TransactionCounterparty {
	this := TransactionCounterparty{}
	return &this
}

// GetName returns the Name field value
func (o *TransactionCounterparty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TransactionCounterparty) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TransactionCounterparty) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TransactionCounterparty) GetType() CounterpartyType {
	if o == nil {
		var ret CounterpartyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionCounterparty) GetTypeOk() (*CounterpartyType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionCounterparty) SetType(v CounterpartyType) {
	o.Type = v
}

// GetWebsite returns the Website field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionCounterparty) GetWebsite() string {
	if o == nil || o.Website.Get() == nil {
		var ret string
		return ret
	}

	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionCounterparty) GetWebsiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// SetWebsite sets field value
func (o *TransactionCounterparty) SetWebsite(v string) {
	o.Website.Set(&v)
}

// GetLogoUrl returns the LogoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionCounterparty) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionCounterparty) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// SetLogoUrl sets field value
func (o *TransactionCounterparty) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

func (o TransactionCounterparty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["website"] = o.Website.Get()
	}
	if true {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionCounterparty) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionCounterparty := _TransactionCounterparty{}

	if err = json.Unmarshal(bytes, &varTransactionCounterparty); err == nil {
		*o = TransactionCounterparty(varTransactionCounterparty)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "website")
		delete(additionalProperties, "logo_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionCounterparty struct {
	value *TransactionCounterparty
	isSet bool
}

func (v NullableTransactionCounterparty) Get() *TransactionCounterparty {
	return v.value
}

func (v *NullableTransactionCounterparty) Set(val *TransactionCounterparty) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCounterparty) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCounterparty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCounterparty(val *TransactionCounterparty) *NullableTransactionCounterparty {
	return &NullableTransactionCounterparty{value: val, isSet: true}
}

func (v NullableTransactionCounterparty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCounterparty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

