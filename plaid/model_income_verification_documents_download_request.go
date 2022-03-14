/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.84.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeVerificationDocumentsDownloadRequest IncomeVerificationDocumentsDownloadRequest defines the request schema for `/income/verification/documents/download`.
type IncomeVerificationDocumentsDownloadRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the verification.
	IncomeVerificationId NullableString `json:"income_verification_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken NullableString `json:"access_token,omitempty"`
	// The document ID to download. If passed, a single document will be returned in the resulting zip file, rather than all document
	DocumentId NullableString `json:"document_id,omitempty"`
}

// NewIncomeVerificationDocumentsDownloadRequest instantiates a new IncomeVerificationDocumentsDownloadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationDocumentsDownloadRequest() *IncomeVerificationDocumentsDownloadRequest {
	this := IncomeVerificationDocumentsDownloadRequest{}
	return &this
}

// NewIncomeVerificationDocumentsDownloadRequestWithDefaults instantiates a new IncomeVerificationDocumentsDownloadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationDocumentsDownloadRequestWithDefaults() *IncomeVerificationDocumentsDownloadRequest {
	this := IncomeVerificationDocumentsDownloadRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IncomeVerificationDocumentsDownloadRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IncomeVerificationDocumentsDownloadRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IncomeVerificationDocumentsDownloadRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IncomeVerificationDocumentsDownloadRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIncomeVerificationId returns the IncomeVerificationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationDocumentsDownloadRequest) GetIncomeVerificationId() string {
	if o == nil || o.IncomeVerificationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IncomeVerificationId.Get()
}

// GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationDocumentsDownloadRequest) GetIncomeVerificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncomeVerificationId.Get(), o.IncomeVerificationId.IsSet()
}

// HasIncomeVerificationId returns a boolean if a field has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) HasIncomeVerificationId() bool {
	if o != nil && o.IncomeVerificationId.IsSet() {
		return true
	}

	return false
}

// SetIncomeVerificationId gets a reference to the given NullableString and assigns it to the IncomeVerificationId field.
func (o *IncomeVerificationDocumentsDownloadRequest) SetIncomeVerificationId(v string) {
	o.IncomeVerificationId.Set(&v)
}
// SetIncomeVerificationIdNil sets the value for IncomeVerificationId to be an explicit nil
func (o *IncomeVerificationDocumentsDownloadRequest) SetIncomeVerificationIdNil() {
	o.IncomeVerificationId.Set(nil)
}

// UnsetIncomeVerificationId ensures that no value is present for IncomeVerificationId, not even an explicit nil
func (o *IncomeVerificationDocumentsDownloadRequest) UnsetIncomeVerificationId() {
	o.IncomeVerificationId.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationDocumentsDownloadRequest) GetAccessToken() string {
	if o == nil || o.AccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationDocumentsDownloadRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *IncomeVerificationDocumentsDownloadRequest) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *IncomeVerificationDocumentsDownloadRequest) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *IncomeVerificationDocumentsDownloadRequest) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationDocumentsDownloadRequest) GetDocumentId() string {
	if o == nil || o.DocumentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocumentId.Get()
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationDocumentsDownloadRequest) GetDocumentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentId.Get(), o.DocumentId.IsSet()
}

// HasDocumentId returns a boolean if a field has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) HasDocumentId() bool {
	if o != nil && o.DocumentId.IsSet() {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given NullableString and assigns it to the DocumentId field.
func (o *IncomeVerificationDocumentsDownloadRequest) SetDocumentId(v string) {
	o.DocumentId.Set(&v)
}
// SetDocumentIdNil sets the value for DocumentId to be an explicit nil
func (o *IncomeVerificationDocumentsDownloadRequest) SetDocumentIdNil() {
	o.DocumentId.Set(nil)
}

// UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
func (o *IncomeVerificationDocumentsDownloadRequest) UnsetDocumentId() {
	o.DocumentId.Unset()
}

func (o IncomeVerificationDocumentsDownloadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.IncomeVerificationId.IsSet() {
		toSerialize["income_verification_id"] = o.IncomeVerificationId.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	if o.DocumentId.IsSet() {
		toSerialize["document_id"] = o.DocumentId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationDocumentsDownloadRequest struct {
	value *IncomeVerificationDocumentsDownloadRequest
	isSet bool
}

func (v NullableIncomeVerificationDocumentsDownloadRequest) Get() *IncomeVerificationDocumentsDownloadRequest {
	return v.value
}

func (v *NullableIncomeVerificationDocumentsDownloadRequest) Set(val *IncomeVerificationDocumentsDownloadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationDocumentsDownloadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationDocumentsDownloadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationDocumentsDownloadRequest(val *IncomeVerificationDocumentsDownloadRequest) *NullableIncomeVerificationDocumentsDownloadRequest {
	return &NullableIncomeVerificationDocumentsDownloadRequest{value: val, isSet: true}
}

func (v NullableIncomeVerificationDocumentsDownloadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationDocumentsDownloadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


