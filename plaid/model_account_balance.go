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
	"time"
)

// AccountBalance A set of fields describing the balance for an account. Balance information may be cached unless the balance object was returned by `/accounts/balance/get`.
type AccountBalance struct {
	// The amount of funds available to be withdrawn from the account, as determined by the financial institution.  For `credit`-type accounts, the `available` balance typically equals the `limit` less the `current` balance, less any pending outflows plus any pending inflows.  For `depository`-type accounts, the `available` balance typically equals the `current` balance less any pending outflows plus any pending inflows. For `depository`-type accounts, the `available` balance does not include the overdraft limit.  For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the `available` balance is the total cash available to withdraw as presented by the institution.  Note that not all institutions calculate the `available`  balance. In the event that `available` balance is unavailable, Plaid will return an `available` balance value of `null`.  Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by `/accounts/balance/get`.  If `current` is `null` this field is guaranteed not to be `null`.
	Available NullableFloat64 `json:"available"`
	// The total amount of funds in or owed by the account.  For `credit`-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  For `loan`-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (`ins_116944`). For Sallie Mae student loans, the account's balance includes both principal and any outstanding interest.  For `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier), the current balance is the total value of assets as presented by the institution.  Note that balance information may be cached unless the value was returned by `/accounts/balance/get`; if the Item is enabled for Transactions, the balance will be at least as recent as the most recent Transaction update. If you require realtime balance information, use the `available` balance as provided by `/accounts/balance/get`.  When returned by `/accounts/balance/get`, this field may be `null`. When this happens, `available` is guaranteed not to be `null`.
	Current NullableFloat64 `json:"current"`
	// For `credit`-type accounts, this represents the credit limit.  For `depository`-type accounts, this represents the pre-arranged overdraft limit, which is common for current (checking) accounts in Europe.  In North America, this field is typically only available for `credit`-type accounts.
	Limit NullableFloat64 `json:"limit"`
	// The ISO-4217 currency code of the balance. Always null if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the balance. Always null if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `unofficial_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the last time that the balance for the given account has been updated  This is currently only provided when the `min_last_updated_datetime` is passed when calling `/accounts/balance/get` for `ins_128026` (Capital One).
	LastUpdatedDatetime NullableTime `json:"last_updated_datetime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountBalance AccountBalance

// NewAccountBalance instantiates a new AccountBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBalance(available NullableFloat64, current NullableFloat64, limit NullableFloat64, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString) *AccountBalance {
	this := AccountBalance{}
	this.Available = available
	this.Current = current
	this.Limit = limit
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewAccountBalanceWithDefaults instantiates a new AccountBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBalanceWithDefaults() *AccountBalance {
	this := AccountBalance{}
	return &this
}

// GetAvailable returns the Available field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *AccountBalance) GetAvailable() float64 {
	if o == nil || o.Available.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Available.Get()
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBalance) GetAvailableOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Available.Get(), o.Available.IsSet()
}

// SetAvailable sets field value
func (o *AccountBalance) SetAvailable(v float64) {
	o.Available.Set(&v)
}

// GetCurrent returns the Current field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *AccountBalance) GetCurrent() float64 {
	if o == nil || o.Current.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Current.Get()
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBalance) GetCurrentOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Current.Get(), o.Current.IsSet()
}

// SetCurrent sets field value
func (o *AccountBalance) SetCurrent(v float64) {
	o.Current.Set(&v)
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *AccountBalance) GetLimit() float64 {
	if o == nil || o.Limit.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBalance) GetLimitOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// SetLimit sets field value
func (o *AccountBalance) SetLimit(v float64) {
	o.Limit.Set(&v)
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountBalance) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBalance) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *AccountBalance) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountBalance) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBalance) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *AccountBalance) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

// GetLastUpdatedDatetime returns the LastUpdatedDatetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountBalance) GetLastUpdatedDatetime() time.Time {
	if o == nil || o.LastUpdatedDatetime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDatetime.Get()
}

// GetLastUpdatedDatetimeOk returns a tuple with the LastUpdatedDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBalance) GetLastUpdatedDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdatedDatetime.Get(), o.LastUpdatedDatetime.IsSet()
}

// HasLastUpdatedDatetime returns a boolean if a field has been set.
func (o *AccountBalance) HasLastUpdatedDatetime() bool {
	if o != nil && o.LastUpdatedDatetime.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedDatetime gets a reference to the given NullableTime and assigns it to the LastUpdatedDatetime field.
func (o *AccountBalance) SetLastUpdatedDatetime(v time.Time) {
	o.LastUpdatedDatetime.Set(&v)
}
// SetLastUpdatedDatetimeNil sets the value for LastUpdatedDatetime to be an explicit nil
func (o *AccountBalance) SetLastUpdatedDatetimeNil() {
	o.LastUpdatedDatetime.Set(nil)
}

// UnsetLastUpdatedDatetime ensures that no value is present for LastUpdatedDatetime, not even an explicit nil
func (o *AccountBalance) UnsetLastUpdatedDatetime() {
	o.LastUpdatedDatetime.Unset()
}

func (o AccountBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["available"] = o.Available.Get()
	}
	if true {
		toSerialize["current"] = o.Current.Get()
	}
	if true {
		toSerialize["limit"] = o.Limit.Get()
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if o.LastUpdatedDatetime.IsSet() {
		toSerialize["last_updated_datetime"] = o.LastUpdatedDatetime.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountBalance) UnmarshalJSON(bytes []byte) (err error) {
	varAccountBalance := _AccountBalance{}

	if err = json.Unmarshal(bytes, &varAccountBalance); err == nil {
		*o = AccountBalance(varAccountBalance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "available")
		delete(additionalProperties, "current")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "last_updated_datetime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountBalance struct {
	value *AccountBalance
	isSet bool
}

func (v NullableAccountBalance) Get() *AccountBalance {
	return v.value
}

func (v *NullableAccountBalance) Set(val *AccountBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBalance(val *AccountBalance) *NullableAccountBalance {
	return &NullableAccountBalance{value: val, isSet: true}
}

func (v NullableAccountBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


