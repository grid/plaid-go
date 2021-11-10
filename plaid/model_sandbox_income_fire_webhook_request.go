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

// SandboxIncomeFireWebhookRequest SandboxIncomeFireWebhookRequest defines the request schema for `/sandbox/income/fire_webhook`
type SandboxIncomeFireWebhookRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the verification.
	IncomeVerificationId string `json:"income_verification_id"`
	// The Item ID associated with the verification.
	ItemId string `json:"item_id"`
	// The URL to which the webhook should be sent.
	Webhook string `json:"webhook"`
	// `VERIFICATION_STATUS_PROCESSING_COMPLETE`: The income verification status processing has completed.  `VERIFICATION_STATUS_DOCUMENT_REJECTED`: The documentation uploaded by the end user was recognized as a supported file format, but not recognized as a valid paystub.  `VERIFICATION_STATUS_PROCESSING_FAILED`: A failure occurred when attempting to process the verification documentation.
	VerificationStatus string `json:"verification_status"`
}

// NewSandboxIncomeFireWebhookRequest instantiates a new SandboxIncomeFireWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxIncomeFireWebhookRequest(incomeVerificationId string, itemId string, webhook string, verificationStatus string) *SandboxIncomeFireWebhookRequest {
	this := SandboxIncomeFireWebhookRequest{}
	this.IncomeVerificationId = incomeVerificationId
	this.ItemId = itemId
	this.Webhook = webhook
	this.VerificationStatus = verificationStatus
	return &this
}

// NewSandboxIncomeFireWebhookRequestWithDefaults instantiates a new SandboxIncomeFireWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxIncomeFireWebhookRequestWithDefaults() *SandboxIncomeFireWebhookRequest {
	this := SandboxIncomeFireWebhookRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxIncomeFireWebhookRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxIncomeFireWebhookRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxIncomeFireWebhookRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxIncomeFireWebhookRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxIncomeFireWebhookRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxIncomeFireWebhookRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxIncomeFireWebhookRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxIncomeFireWebhookRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIncomeVerificationId returns the IncomeVerificationId field value
func (o *SandboxIncomeFireWebhookRequest) GetIncomeVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IncomeVerificationId
}

// GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field value
// and a boolean to check if the value has been set.
func (o *SandboxIncomeFireWebhookRequest) GetIncomeVerificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncomeVerificationId, true
}

// SetIncomeVerificationId sets field value
func (o *SandboxIncomeFireWebhookRequest) SetIncomeVerificationId(v string) {
	o.IncomeVerificationId = v
}

// GetItemId returns the ItemId field value
func (o *SandboxIncomeFireWebhookRequest) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *SandboxIncomeFireWebhookRequest) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *SandboxIncomeFireWebhookRequest) SetItemId(v string) {
	o.ItemId = v
}

// GetWebhook returns the Webhook field value
func (o *SandboxIncomeFireWebhookRequest) GetWebhook() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value
// and a boolean to check if the value has been set.
func (o *SandboxIncomeFireWebhookRequest) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Webhook, true
}

// SetWebhook sets field value
func (o *SandboxIncomeFireWebhookRequest) SetWebhook(v string) {
	o.Webhook = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *SandboxIncomeFireWebhookRequest) GetVerificationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *SandboxIncomeFireWebhookRequest) GetVerificationStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *SandboxIncomeFireWebhookRequest) SetVerificationStatus(v string) {
	o.VerificationStatus = v
}

func (o SandboxIncomeFireWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["income_verification_id"] = o.IncomeVerificationId
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["webhook"] = o.Webhook
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxIncomeFireWebhookRequest struct {
	value *SandboxIncomeFireWebhookRequest
	isSet bool
}

func (v NullableSandboxIncomeFireWebhookRequest) Get() *SandboxIncomeFireWebhookRequest {
	return v.value
}

func (v *NullableSandboxIncomeFireWebhookRequest) Set(val *SandboxIncomeFireWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxIncomeFireWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxIncomeFireWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxIncomeFireWebhookRequest(val *SandboxIncomeFireWebhookRequest) *NullableSandboxIncomeFireWebhookRequest {
	return &NullableSandboxIncomeFireWebhookRequest{value: val, isSet: true}
}

func (v NullableSandboxIncomeFireWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxIncomeFireWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


