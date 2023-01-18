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

// DistributionBreakdown Information about the accounts that the payment was distributed to.
type DistributionBreakdown struct {
	// Name of the account for the given distribution.
	AccountName NullableString `json:"account_name,omitempty"`
	// The name of the bank that the payment is being deposited to.
	BankName NullableString `json:"bank_name,omitempty"`
	// The amount distributed to this account.
	CurrentAmount NullableFloat64 `json:"current_amount,omitempty"`
	// The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code,omitempty"`
	// The last 2-4 alphanumeric characters of an account's official account number.
	Mask NullableString `json:"mask,omitempty"`
	// Type of the account that the paystub was sent to (e.g. 'checking').
	Type NullableString `json:"type,omitempty"`
	// The unofficial currency code associated with the net pay. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code,omitempty"`
	CurrentPay *Pay `json:"current_pay,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DistributionBreakdown DistributionBreakdown

// NewDistributionBreakdown instantiates a new DistributionBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionBreakdown() *DistributionBreakdown {
	this := DistributionBreakdown{}
	return &this
}

// NewDistributionBreakdownWithDefaults instantiates a new DistributionBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionBreakdownWithDefaults() *DistributionBreakdown {
	this := DistributionBreakdown{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DistributionBreakdown) GetAccountName() string {
	if o == nil || o.AccountName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountName.Get()
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionBreakdown) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountName.Get(), o.AccountName.IsSet()
}

// HasAccountName returns a boolean if a field has been set.
func (o *DistributionBreakdown) HasAccountName() bool {
	if o != nil && o.AccountName.IsSet() {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given NullableString and assigns it to the AccountName field.
func (o *DistributionBreakdown) SetAccountName(v string) {
	o.AccountName.Set(&v)
}
// SetAccountNameNil sets the value for AccountName to be an explicit nil
func (o *DistributionBreakdown) SetAccountNameNil() {
	o.AccountName.Set(nil)
}

// UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
func (o *DistributionBreakdown) UnsetAccountName() {
	o.AccountName.Unset()
}

// GetBankName returns the BankName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DistributionBreakdown) GetBankName() string {
	if o == nil || o.BankName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionBreakdown) GetBankNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// HasBankName returns a boolean if a field has been set.
func (o *DistributionBreakdown) HasBankName() bool {
	if o != nil && o.BankName.IsSet() {
		return true
	}

	return false
}

// SetBankName gets a reference to the given NullableString and assigns it to the BankName field.
func (o *DistributionBreakdown) SetBankName(v string) {
	o.BankName.Set(&v)
}
// SetBankNameNil sets the value for BankName to be an explicit nil
func (o *DistributionBreakdown) SetBankNameNil() {
	o.BankName.Set(nil)
}

// UnsetBankName ensures that no value is present for BankName, not even an explicit nil
func (o *DistributionBreakdown) UnsetBankName() {
	o.BankName.Unset()
}

// GetCurrentAmount returns the CurrentAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DistributionBreakdown) GetCurrentAmount() float64 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionBreakdown) GetCurrentAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// HasCurrentAmount returns a boolean if a field has been set.
func (o *DistributionBreakdown) HasCurrentAmount() bool {
	if o != nil && o.CurrentAmount.IsSet() {
		return true
	}

	return false
}

// SetCurrentAmount gets a reference to the given NullableFloat64 and assigns it to the CurrentAmount field.
func (o *DistributionBreakdown) SetCurrentAmount(v float64) {
	o.CurrentAmount.Set(&v)
}
// SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil
func (o *DistributionBreakdown) SetCurrentAmountNil() {
	o.CurrentAmount.Set(nil)
}

// UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
func (o *DistributionBreakdown) UnsetCurrentAmount() {
	o.CurrentAmount.Unset()
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DistributionBreakdown) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionBreakdown) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *DistributionBreakdown) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given NullableString and assigns it to the IsoCurrencyCode field.
func (o *DistributionBreakdown) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}
// SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil
func (o *DistributionBreakdown) SetIsoCurrencyCodeNil() {
	o.IsoCurrencyCode.Set(nil)
}

// UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
func (o *DistributionBreakdown) UnsetIsoCurrencyCode() {
	o.IsoCurrencyCode.Unset()
}

// GetMask returns the Mask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DistributionBreakdown) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionBreakdown) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// HasMask returns a boolean if a field has been set.
func (o *DistributionBreakdown) HasMask() bool {
	if o != nil && o.Mask.IsSet() {
		return true
	}

	return false
}

// SetMask gets a reference to the given NullableString and assigns it to the Mask field.
func (o *DistributionBreakdown) SetMask(v string) {
	o.Mask.Set(&v)
}
// SetMaskNil sets the value for Mask to be an explicit nil
func (o *DistributionBreakdown) SetMaskNil() {
	o.Mask.Set(nil)
}

// UnsetMask ensures that no value is present for Mask, not even an explicit nil
func (o *DistributionBreakdown) UnsetMask() {
	o.Mask.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DistributionBreakdown) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionBreakdown) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *DistributionBreakdown) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *DistributionBreakdown) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *DistributionBreakdown) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *DistributionBreakdown) UnsetType() {
	o.Type.Unset()
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DistributionBreakdown) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionBreakdown) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// HasUnofficialCurrencyCode returns a boolean if a field has been set.
func (o *DistributionBreakdown) HasUnofficialCurrencyCode() bool {
	if o != nil && o.UnofficialCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetUnofficialCurrencyCode gets a reference to the given NullableString and assigns it to the UnofficialCurrencyCode field.
func (o *DistributionBreakdown) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}
// SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil
func (o *DistributionBreakdown) SetUnofficialCurrencyCodeNil() {
	o.UnofficialCurrencyCode.Set(nil)
}

// UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
func (o *DistributionBreakdown) UnsetUnofficialCurrencyCode() {
	o.UnofficialCurrencyCode.Unset()
}

// GetCurrentPay returns the CurrentPay field value if set, zero value otherwise.
func (o *DistributionBreakdown) GetCurrentPay() Pay {
	if o == nil || o.CurrentPay == nil {
		var ret Pay
		return ret
	}
	return *o.CurrentPay
}

// GetCurrentPayOk returns a tuple with the CurrentPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBreakdown) GetCurrentPayOk() (*Pay, bool) {
	if o == nil || o.CurrentPay == nil {
		return nil, false
	}
	return o.CurrentPay, true
}

// HasCurrentPay returns a boolean if a field has been set.
func (o *DistributionBreakdown) HasCurrentPay() bool {
	if o != nil && o.CurrentPay != nil {
		return true
	}

	return false
}

// SetCurrentPay gets a reference to the given Pay and assigns it to the CurrentPay field.
func (o *DistributionBreakdown) SetCurrentPay(v Pay) {
	o.CurrentPay = &v
}

func (o DistributionBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountName.IsSet() {
		toSerialize["account_name"] = o.AccountName.Get()
	}
	if o.BankName.IsSet() {
		toSerialize["bank_name"] = o.BankName.Get()
	}
	if o.CurrentAmount.IsSet() {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if o.IsoCurrencyCode.IsSet() {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if o.Mask.IsSet() {
		toSerialize["mask"] = o.Mask.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.UnofficialCurrencyCode.IsSet() {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if o.CurrentPay != nil {
		toSerialize["current_pay"] = o.CurrentPay
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DistributionBreakdown) UnmarshalJSON(bytes []byte) (err error) {
	varDistributionBreakdown := _DistributionBreakdown{}

	if err = json.Unmarshal(bytes, &varDistributionBreakdown); err == nil {
		*o = DistributionBreakdown(varDistributionBreakdown)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "bank_name")
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "mask")
		delete(additionalProperties, "type")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "current_pay")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDistributionBreakdown struct {
	value *DistributionBreakdown
	isSet bool
}

func (v NullableDistributionBreakdown) Get() *DistributionBreakdown {
	return v.value
}

func (v *NullableDistributionBreakdown) Set(val *DistributionBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionBreakdown(val *DistributionBreakdown) *NullableDistributionBreakdown {
	return &NullableDistributionBreakdown{value: val, isSet: true}
}

func (v NullableDistributionBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


