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
	"time"
)

// AccountsBalanceGetRequestOptions An optional object to filter `/accounts/balance/get` results.
type AccountsBalanceGetRequestOptions struct {
	// A list of `account_ids` to retrieve for the Item. The default value is `null`.  Note: An error will be returned if a provided `account_id` is not associated with the Item.
	AccountIds *[]string `json:"account_ids,omitempty"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the oldest acceptable balance when making a request to `/accounts/balance/get`.  If the balance that is pulled for `ins_128026` (Capital One) is older than the given timestamp, an `INVALID_REQUEST` error with the code of `LAST_UPDATED_DATETIME_OUT_OF_RANGE` will be returned with the most recent timestamp for the requested account contained in the response.  This field is only used when the institution is `ins_128026` (Capital One), in which case a value must be provided or an `INVALID_REQUEST` error with the code of `INVALID_FIELD` will be returned. For all other institutions, this field is ignored.
	MinLastUpdatedDatetime *time.Time `json:"min_last_updated_datetime,omitempty"`
}

// NewAccountsBalanceGetRequestOptions instantiates a new AccountsBalanceGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsBalanceGetRequestOptions() *AccountsBalanceGetRequestOptions {
	this := AccountsBalanceGetRequestOptions{}
	return &this
}

// NewAccountsBalanceGetRequestOptionsWithDefaults instantiates a new AccountsBalanceGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsBalanceGetRequestOptionsWithDefaults() *AccountsBalanceGetRequestOptions {
	this := AccountsBalanceGetRequestOptions{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *AccountsBalanceGetRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsBalanceGetRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *AccountsBalanceGetRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *AccountsBalanceGetRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

// GetMinLastUpdatedDatetime returns the MinLastUpdatedDatetime field value if set, zero value otherwise.
func (o *AccountsBalanceGetRequestOptions) GetMinLastUpdatedDatetime() time.Time {
	if o == nil || o.MinLastUpdatedDatetime == nil {
		var ret time.Time
		return ret
	}
	return *o.MinLastUpdatedDatetime
}

// GetMinLastUpdatedDatetimeOk returns a tuple with the MinLastUpdatedDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsBalanceGetRequestOptions) GetMinLastUpdatedDatetimeOk() (*time.Time, bool) {
	if o == nil || o.MinLastUpdatedDatetime == nil {
		return nil, false
	}
	return o.MinLastUpdatedDatetime, true
}

// HasMinLastUpdatedDatetime returns a boolean if a field has been set.
func (o *AccountsBalanceGetRequestOptions) HasMinLastUpdatedDatetime() bool {
	if o != nil && o.MinLastUpdatedDatetime != nil {
		return true
	}

	return false
}

// SetMinLastUpdatedDatetime gets a reference to the given time.Time and assigns it to the MinLastUpdatedDatetime field.
func (o *AccountsBalanceGetRequestOptions) SetMinLastUpdatedDatetime(v time.Time) {
	o.MinLastUpdatedDatetime = &v
}

func (o AccountsBalanceGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	if o.MinLastUpdatedDatetime != nil {
		toSerialize["min_last_updated_datetime"] = o.MinLastUpdatedDatetime
	}
	return json.Marshal(toSerialize)
}

type NullableAccountsBalanceGetRequestOptions struct {
	value *AccountsBalanceGetRequestOptions
	isSet bool
}

func (v NullableAccountsBalanceGetRequestOptions) Get() *AccountsBalanceGetRequestOptions {
	return v.value
}

func (v *NullableAccountsBalanceGetRequestOptions) Set(val *AccountsBalanceGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsBalanceGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsBalanceGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsBalanceGetRequestOptions(val *AccountsBalanceGetRequestOptions) *NullableAccountsBalanceGetRequestOptions {
	return &NullableAccountsBalanceGetRequestOptions{value: val, isSet: true}
}

func (v NullableAccountsBalanceGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsBalanceGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


