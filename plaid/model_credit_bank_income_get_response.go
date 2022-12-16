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

// CreditBankIncomeGetResponse CreditBankIncomeGetResponse defines the response schema for `/credit/bank_income/get`
type CreditBankIncomeGetResponse struct {
	BankIncome *[]CreditBankIncome `json:"bank_income,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditBankIncomeGetResponse CreditBankIncomeGetResponse

// NewCreditBankIncomeGetResponse instantiates a new CreditBankIncomeGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeGetResponse(requestId string) *CreditBankIncomeGetResponse {
	this := CreditBankIncomeGetResponse{}
	this.RequestId = requestId
	return &this
}

// NewCreditBankIncomeGetResponseWithDefaults instantiates a new CreditBankIncomeGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeGetResponseWithDefaults() *CreditBankIncomeGetResponse {
	this := CreditBankIncomeGetResponse{}
	return &this
}

// GetBankIncome returns the BankIncome field value if set, zero value otherwise.
func (o *CreditBankIncomeGetResponse) GetBankIncome() []CreditBankIncome {
	if o == nil || o.BankIncome == nil {
		var ret []CreditBankIncome
		return ret
	}
	return *o.BankIncome
}

// GetBankIncomeOk returns a tuple with the BankIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeGetResponse) GetBankIncomeOk() (*[]CreditBankIncome, bool) {
	if o == nil || o.BankIncome == nil {
		return nil, false
	}
	return o.BankIncome, true
}

// HasBankIncome returns a boolean if a field has been set.
func (o *CreditBankIncomeGetResponse) HasBankIncome() bool {
	if o != nil && o.BankIncome != nil {
		return true
	}

	return false
}

// SetBankIncome gets a reference to the given []CreditBankIncome and assigns it to the BankIncome field.
func (o *CreditBankIncomeGetResponse) SetBankIncome(v []CreditBankIncome) {
	o.BankIncome = &v
}

// GetRequestId returns the RequestId field value
func (o *CreditBankIncomeGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditBankIncomeGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditBankIncomeGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BankIncome != nil {
		toSerialize["bank_income"] = o.BankIncome
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditBankIncomeGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditBankIncomeGetResponse := _CreditBankIncomeGetResponse{}

	if err = json.Unmarshal(bytes, &varCreditBankIncomeGetResponse); err == nil {
		*o = CreditBankIncomeGetResponse(varCreditBankIncomeGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bank_income")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditBankIncomeGetResponse struct {
	value *CreditBankIncomeGetResponse
	isSet bool
}

func (v NullableCreditBankIncomeGetResponse) Get() *CreditBankIncomeGetResponse {
	return v.value
}

func (v *NullableCreditBankIncomeGetResponse) Set(val *CreditBankIncomeGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeGetResponse(val *CreditBankIncomeGetResponse) *NullableCreditBankIncomeGetResponse {
	return &NullableCreditBankIncomeGetResponse{value: val, isSet: true}
}

func (v NullableCreditBankIncomeGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


