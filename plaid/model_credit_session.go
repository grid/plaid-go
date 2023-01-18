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

// CreditSession Metadata and results for a Link session
type CreditSession struct {
	// The unique identifier associated with the Link session. This identifier matches the `link_session_id` returned in the onSuccess/onExit callbacks.
	LinkSessionId *string `json:"link_session_id,omitempty"`
	// The time when the Link session started
	SessionStartTime *time.Time `json:"session_start_time,omitempty"`
	Results *CreditSessionResults `json:"results,omitempty"`
	// The set of errors that occurred during the Link session.
	Errors *[]CreditSessionError `json:"errors,omitempty"`
}

// NewCreditSession instantiates a new CreditSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditSession() *CreditSession {
	this := CreditSession{}
	return &this
}

// NewCreditSessionWithDefaults instantiates a new CreditSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditSessionWithDefaults() *CreditSession {
	this := CreditSession{}
	return &this
}

// GetLinkSessionId returns the LinkSessionId field value if set, zero value otherwise.
func (o *CreditSession) GetLinkSessionId() string {
	if o == nil || o.LinkSessionId == nil {
		var ret string
		return ret
	}
	return *o.LinkSessionId
}

// GetLinkSessionIdOk returns a tuple with the LinkSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSession) GetLinkSessionIdOk() (*string, bool) {
	if o == nil || o.LinkSessionId == nil {
		return nil, false
	}
	return o.LinkSessionId, true
}

// HasLinkSessionId returns a boolean if a field has been set.
func (o *CreditSession) HasLinkSessionId() bool {
	if o != nil && o.LinkSessionId != nil {
		return true
	}

	return false
}

// SetLinkSessionId gets a reference to the given string and assigns it to the LinkSessionId field.
func (o *CreditSession) SetLinkSessionId(v string) {
	o.LinkSessionId = &v
}

// GetSessionStartTime returns the SessionStartTime field value if set, zero value otherwise.
func (o *CreditSession) GetSessionStartTime() time.Time {
	if o == nil || o.SessionStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.SessionStartTime
}

// GetSessionStartTimeOk returns a tuple with the SessionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSession) GetSessionStartTimeOk() (*time.Time, bool) {
	if o == nil || o.SessionStartTime == nil {
		return nil, false
	}
	return o.SessionStartTime, true
}

// HasSessionStartTime returns a boolean if a field has been set.
func (o *CreditSession) HasSessionStartTime() bool {
	if o != nil && o.SessionStartTime != nil {
		return true
	}

	return false
}

// SetSessionStartTime gets a reference to the given time.Time and assigns it to the SessionStartTime field.
func (o *CreditSession) SetSessionStartTime(v time.Time) {
	o.SessionStartTime = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CreditSession) GetResults() CreditSessionResults {
	if o == nil || o.Results == nil {
		var ret CreditSessionResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSession) GetResultsOk() (*CreditSessionResults, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CreditSession) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given CreditSessionResults and assigns it to the Results field.
func (o *CreditSession) SetResults(v CreditSessionResults) {
	o.Results = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreditSession) GetErrors() []CreditSessionError {
	if o == nil || o.Errors == nil {
		var ret []CreditSessionError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSession) GetErrorsOk() (*[]CreditSessionError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreditSession) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []CreditSessionError and assigns it to the Errors field.
func (o *CreditSession) SetErrors(v []CreditSessionError) {
	o.Errors = &v
}

func (o CreditSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkSessionId != nil {
		toSerialize["link_session_id"] = o.LinkSessionId
	}
	if o.SessionStartTime != nil {
		toSerialize["session_start_time"] = o.SessionStartTime
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableCreditSession struct {
	value *CreditSession
	isSet bool
}

func (v NullableCreditSession) Get() *CreditSession {
	return v.value
}

func (v *NullableCreditSession) Set(val *CreditSession) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditSession) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditSession(val *CreditSession) *NullableCreditSession {
	return &NullableCreditSession{value: val, isSet: true}
}

func (v NullableCreditSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


