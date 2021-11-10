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

// TransferUserAddressInResponse The address associated with the account holder.
type TransferUserAddressInResponse struct {
	// The street number and name (i.e., \"100 Market St.\").
	Street NullableString `json:"street"`
	// Ex. \"San Francisco\"
	City NullableString `json:"city"`
	// The state or province (e.g., \"California\").
	Region NullableString `json:"region"`
	// The postal code (e.g., \"94103\").
	PostalCode NullableString `json:"postal_code"`
	// A two-letter country code (e.g., \"US\").
	Country NullableString `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _TransferUserAddressInResponse TransferUserAddressInResponse

// NewTransferUserAddressInResponse instantiates a new TransferUserAddressInResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferUserAddressInResponse(street NullableString, city NullableString, region NullableString, postalCode NullableString, country NullableString) *TransferUserAddressInResponse {
	this := TransferUserAddressInResponse{}
	this.Street = street
	this.City = city
	this.Region = region
	this.PostalCode = postalCode
	this.Country = country
	return &this
}

// NewTransferUserAddressInResponseWithDefaults instantiates a new TransferUserAddressInResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferUserAddressInResponseWithDefaults() *TransferUserAddressInResponse {
	this := TransferUserAddressInResponse{}
	return &this
}

// GetStreet returns the Street field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferUserAddressInResponse) GetStreet() string {
	if o == nil || o.Street.Get() == nil {
		var ret string
		return ret
	}

	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferUserAddressInResponse) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// SetStreet sets field value
func (o *TransferUserAddressInResponse) SetStreet(v string) {
	o.Street.Set(&v)
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferUserAddressInResponse) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferUserAddressInResponse) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *TransferUserAddressInResponse) SetCity(v string) {
	o.City.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferUserAddressInResponse) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferUserAddressInResponse) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *TransferUserAddressInResponse) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferUserAddressInResponse) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferUserAddressInResponse) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *TransferUserAddressInResponse) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferUserAddressInResponse) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferUserAddressInResponse) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *TransferUserAddressInResponse) SetCountry(v string) {
	o.Country.Set(&v)
}

func (o TransferUserAddressInResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["street"] = o.Street.Get()
	}
	if true {
		toSerialize["city"] = o.City.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferUserAddressInResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferUserAddressInResponse := _TransferUserAddressInResponse{}

	if err = json.Unmarshal(bytes, &varTransferUserAddressInResponse); err == nil {
		*o = TransferUserAddressInResponse(varTransferUserAddressInResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "street")
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferUserAddressInResponse struct {
	value *TransferUserAddressInResponse
	isSet bool
}

func (v NullableTransferUserAddressInResponse) Get() *TransferUserAddressInResponse {
	return v.value
}

func (v *NullableTransferUserAddressInResponse) Set(val *TransferUserAddressInResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferUserAddressInResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferUserAddressInResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferUserAddressInResponse(val *TransferUserAddressInResponse) *NullableTransferUserAddressInResponse {
	return &NullableTransferUserAddressInResponse{value: val, isSet: true}
}

func (v NullableTransferUserAddressInResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferUserAddressInResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


