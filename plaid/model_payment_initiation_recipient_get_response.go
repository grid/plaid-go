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

// PaymentInitiationRecipientGetResponse PaymentInitiationRecipientGetResponse defines the response schema for `/payment_initiation/recipient/get`
type PaymentInitiationRecipientGetResponse struct {
	// The ID of the recipient.
	RecipientId string `json:"recipient_id"`
	// The name of the recipient.
	Name string `json:"name"`
	Address NullablePaymentInitiationAddress `json:"address,omitempty"`
	// The International Bank Account Number (IBAN) for the recipient.
	Iban NullableString `json:"iban,omitempty"`
	Bacs NullableRecipientBACSNullable `json:"bacs,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationRecipientGetResponse PaymentInitiationRecipientGetResponse

// NewPaymentInitiationRecipientGetResponse instantiates a new PaymentInitiationRecipientGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipientGetResponse(recipientId string, name string, requestId string) *PaymentInitiationRecipientGetResponse {
	this := PaymentInitiationRecipientGetResponse{}
	this.RecipientId = recipientId
	this.Name = name
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationRecipientGetResponseWithDefaults instantiates a new PaymentInitiationRecipientGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientGetResponseWithDefaults() *PaymentInitiationRecipientGetResponse {
	this := PaymentInitiationRecipientGetResponse{}
	return &this
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationRecipientGetResponse) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientGetResponse) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationRecipientGetResponse) SetRecipientId(v string) {
	o.RecipientId = v
}

// GetName returns the Name field value
func (o *PaymentInitiationRecipientGetResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientGetResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentInitiationRecipientGetResponse) SetName(v string) {
	o.Name = v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipientGetResponse) GetAddress() PaymentInitiationAddress {
	if o == nil || o.Address.Get() == nil {
		var ret PaymentInitiationAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipientGetResponse) GetAddressOk() (*PaymentInitiationAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientGetResponse) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullablePaymentInitiationAddress and assigns it to the Address field.
func (o *PaymentInitiationRecipientGetResponse) SetAddress(v PaymentInitiationAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *PaymentInitiationRecipientGetResponse) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *PaymentInitiationRecipientGetResponse) UnsetAddress() {
	o.Address.Unset()
}

// GetIban returns the Iban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipientGetResponse) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}
	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipientGetResponse) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// HasIban returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientGetResponse) HasIban() bool {
	if o != nil && o.Iban.IsSet() {
		return true
	}

	return false
}

// SetIban gets a reference to the given NullableString and assigns it to the Iban field.
func (o *PaymentInitiationRecipientGetResponse) SetIban(v string) {
	o.Iban.Set(&v)
}
// SetIbanNil sets the value for Iban to be an explicit nil
func (o *PaymentInitiationRecipientGetResponse) SetIbanNil() {
	o.Iban.Set(nil)
}

// UnsetIban ensures that no value is present for Iban, not even an explicit nil
func (o *PaymentInitiationRecipientGetResponse) UnsetIban() {
	o.Iban.Unset()
}

// GetBacs returns the Bacs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipientGetResponse) GetBacs() RecipientBACSNullable {
	if o == nil || o.Bacs.Get() == nil {
		var ret RecipientBACSNullable
		return ret
	}
	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipientGetResponse) GetBacsOk() (*RecipientBACSNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// HasBacs returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientGetResponse) HasBacs() bool {
	if o != nil && o.Bacs.IsSet() {
		return true
	}

	return false
}

// SetBacs gets a reference to the given NullableRecipientBACSNullable and assigns it to the Bacs field.
func (o *PaymentInitiationRecipientGetResponse) SetBacs(v RecipientBACSNullable) {
	o.Bacs.Set(&v)
}
// SetBacsNil sets the value for Bacs to be an explicit nil
func (o *PaymentInitiationRecipientGetResponse) SetBacsNil() {
	o.Bacs.Set(nil)
}

// UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
func (o *PaymentInitiationRecipientGetResponse) UnsetBacs() {
	o.Bacs.Unset()
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationRecipientGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationRecipientGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationRecipientGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Iban.IsSet() {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.Bacs.IsSet() {
		toSerialize["bacs"] = o.Bacs.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationRecipientGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationRecipientGetResponse := _PaymentInitiationRecipientGetResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationRecipientGetResponse); err == nil {
		*o = PaymentInitiationRecipientGetResponse(varPaymentInitiationRecipientGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recipient_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "address")
		delete(additionalProperties, "iban")
		delete(additionalProperties, "bacs")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationRecipientGetResponse struct {
	value *PaymentInitiationRecipientGetResponse
	isSet bool
}

func (v NullablePaymentInitiationRecipientGetResponse) Get() *PaymentInitiationRecipientGetResponse {
	return v.value
}

func (v *NullablePaymentInitiationRecipientGetResponse) Set(val *PaymentInitiationRecipientGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipientGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipientGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipientGetResponse(val *PaymentInitiationRecipientGetResponse) *NullablePaymentInitiationRecipientGetResponse {
	return &NullablePaymentInitiationRecipientGetResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipientGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipientGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


