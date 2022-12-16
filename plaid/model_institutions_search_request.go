/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.210.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InstitutionsSearchRequest InstitutionsSearchRequest defines the request schema for `/institutions/search`
type InstitutionsSearchRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The search query. Institutions with names matching the query are returned
	Query string `json:"query"`
	// Filter the Institutions based on whether they support all products listed in `products`. Provide `null` to get institutions regardless of supported products. Note that when `auth` is specified as a product, if you are enabled for Instant Match or Automated Micro-deposits, institutions that support those products will be returned even if `auth` is not present in their product array.
	Products []Products `json:"products"`
	// Specify an array of Plaid-supported country codes this institution supports, using the ISO-3166-1 alpha-2 country code standard. In API versions 2019-05-29 and earlier, the `country_codes` parameter is an optional parameter within the `options` object and will default to `[US]` if it is not supplied. 
	CountryCodes []CountryCode `json:"country_codes"`
	Options *InstitutionsSearchRequestOptions `json:"options,omitempty"`
}

// NewInstitutionsSearchRequest instantiates a new InstitutionsSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsSearchRequest(query string, products []Products, countryCodes []CountryCode) *InstitutionsSearchRequest {
	this := InstitutionsSearchRequest{}
	this.Query = query
	this.Products = products
	this.CountryCodes = countryCodes
	return &this
}

// NewInstitutionsSearchRequestWithDefaults instantiates a new InstitutionsSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsSearchRequestWithDefaults() *InstitutionsSearchRequest {
	this := InstitutionsSearchRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InstitutionsSearchRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InstitutionsSearchRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InstitutionsSearchRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *InstitutionsSearchRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *InstitutionsSearchRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *InstitutionsSearchRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetQuery returns the Query field value
func (o *InstitutionsSearchRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequest) GetQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *InstitutionsSearchRequest) SetQuery(v string) {
	o.Query = v
}

// GetProducts returns the Products field value
// If the value is explicit nil, the zero value for []Products will be returned
func (o *InstitutionsSearchRequest) GetProducts() []Products {
	if o == nil {
		var ret []Products
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsSearchRequest) GetProductsOk() (*[]Products, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return &o.Products, true
}

// SetProducts sets field value
func (o *InstitutionsSearchRequest) SetProducts(v []Products) {
	o.Products = v
}

// GetCountryCodes returns the CountryCodes field value
func (o *InstitutionsSearchRequest) GetCountryCodes() []CountryCode {
	if o == nil {
		var ret []CountryCode
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequest) GetCountryCodesOk() (*[]CountryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *InstitutionsSearchRequest) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InstitutionsSearchRequest) GetOptions() InstitutionsSearchRequestOptions {
	if o == nil || o.Options == nil {
		var ret InstitutionsSearchRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequest) GetOptionsOk() (*InstitutionsSearchRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InstitutionsSearchRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given InstitutionsSearchRequestOptions and assigns it to the Options field.
func (o *InstitutionsSearchRequest) SetOptions(v InstitutionsSearchRequestOptions) {
	o.Options = &v
}

func (o InstitutionsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	if true {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionsSearchRequest struct {
	value *InstitutionsSearchRequest
	isSet bool
}

func (v NullableInstitutionsSearchRequest) Get() *InstitutionsSearchRequest {
	return v.value
}

func (v *NullableInstitutionsSearchRequest) Set(val *InstitutionsSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsSearchRequest(val *InstitutionsSearchRequest) *NullableInstitutionsSearchRequest {
	return &NullableInstitutionsSearchRequest{value: val, isSet: true}
}

func (v NullableInstitutionsSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


