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

// WalletTransactionListResponse WalletTransactionListResponse defines the response schema for `/wallet/transaction/list`
type WalletTransactionListResponse struct {
	// An array of transactions of an e-wallet, associated with the given `wallet_id`
	Transactions []WalletTransaction `json:"transactions"`
	// Cursor used for fetching transactions created before the latest transaction provided in this response
	NextCursor *string `json:"next_cursor,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionListResponse WalletTransactionListResponse

// NewWalletTransactionListResponse instantiates a new WalletTransactionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionListResponse(transactions []WalletTransaction, requestId string) *WalletTransactionListResponse {
	this := WalletTransactionListResponse{}
	this.Transactions = transactions
	this.RequestId = requestId
	return &this
}

// NewWalletTransactionListResponseWithDefaults instantiates a new WalletTransactionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionListResponseWithDefaults() *WalletTransactionListResponse {
	this := WalletTransactionListResponse{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *WalletTransactionListResponse) GetTransactions() []WalletTransaction {
	if o == nil {
		var ret []WalletTransaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionListResponse) GetTransactionsOk() (*[]WalletTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *WalletTransactionListResponse) SetTransactions(v []WalletTransaction) {
	o.Transactions = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *WalletTransactionListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *WalletTransactionListResponse) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *WalletTransactionListResponse) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetRequestId returns the RequestId field value
func (o *WalletTransactionListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WalletTransactionListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WalletTransactionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionListResponse := _WalletTransactionListResponse{}

	if err = json.Unmarshal(bytes, &varWalletTransactionListResponse); err == nil {
		*o = WalletTransactionListResponse(varWalletTransactionListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transactions")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionListResponse struct {
	value *WalletTransactionListResponse
	isSet bool
}

func (v NullableWalletTransactionListResponse) Get() *WalletTransactionListResponse {
	return v.value
}

func (v *NullableWalletTransactionListResponse) Set(val *WalletTransactionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionListResponse(val *WalletTransactionListResponse) *NullableWalletTransactionListResponse {
	return &NullableWalletTransactionListResponse{value: val, isSet: true}
}

func (v NullableWalletTransactionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


