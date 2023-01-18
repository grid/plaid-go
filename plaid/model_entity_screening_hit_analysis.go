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

// EntityScreeningHitAnalysis Analysis information describing why a screening hit matched the provided entity information
type EntityScreeningHitAnalysis struct {
	Documents *MatchSummaryCode `json:"documents,omitempty"`
	EmailAddresses *MatchSummaryCode `json:"email_addresses,omitempty"`
	Locations *MatchSummaryCode `json:"locations,omitempty"`
	Names *MatchSummaryCode `json:"names,omitempty"`
	PhoneNumbers *MatchSummaryCode `json:"phone_numbers,omitempty"`
	Urls *MatchSummaryCode `json:"urls,omitempty"`
	// The version of the entity screening's `search_terms` that were compared when the entity screening hit was added. entity screening hits are immutable once they have been reviewed. If changes are detected due to updates to the entity screening's `search_terms`, the associated entity program, or the list's source data prior to review, the entity screening hit will be updated to reflect those changes.
	SearchTermsVersion float32 `json:"search_terms_version"`
}

// NewEntityScreeningHitAnalysis instantiates a new EntityScreeningHitAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitAnalysis(searchTermsVersion float32) *EntityScreeningHitAnalysis {
	this := EntityScreeningHitAnalysis{}
	this.SearchTermsVersion = searchTermsVersion
	return &this
}

// NewEntityScreeningHitAnalysisWithDefaults instantiates a new EntityScreeningHitAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitAnalysisWithDefaults() *EntityScreeningHitAnalysis {
	this := EntityScreeningHitAnalysis{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *EntityScreeningHitAnalysis) GetDocuments() MatchSummaryCode {
	if o == nil || o.Documents == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitAnalysis) GetDocumentsOk() (*MatchSummaryCode, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *EntityScreeningHitAnalysis) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given MatchSummaryCode and assigns it to the Documents field.
func (o *EntityScreeningHitAnalysis) SetDocuments(v MatchSummaryCode) {
	o.Documents = &v
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *EntityScreeningHitAnalysis) GetEmailAddresses() MatchSummaryCode {
	if o == nil || o.EmailAddresses == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitAnalysis) GetEmailAddressesOk() (*MatchSummaryCode, bool) {
	if o == nil || o.EmailAddresses == nil {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *EntityScreeningHitAnalysis) HasEmailAddresses() bool {
	if o != nil && o.EmailAddresses != nil {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given MatchSummaryCode and assigns it to the EmailAddresses field.
func (o *EntityScreeningHitAnalysis) SetEmailAddresses(v MatchSummaryCode) {
	o.EmailAddresses = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *EntityScreeningHitAnalysis) GetLocations() MatchSummaryCode {
	if o == nil || o.Locations == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitAnalysis) GetLocationsOk() (*MatchSummaryCode, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *EntityScreeningHitAnalysis) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given MatchSummaryCode and assigns it to the Locations field.
func (o *EntityScreeningHitAnalysis) SetLocations(v MatchSummaryCode) {
	o.Locations = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *EntityScreeningHitAnalysis) GetNames() MatchSummaryCode {
	if o == nil || o.Names == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitAnalysis) GetNamesOk() (*MatchSummaryCode, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *EntityScreeningHitAnalysis) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given MatchSummaryCode and assigns it to the Names field.
func (o *EntityScreeningHitAnalysis) SetNames(v MatchSummaryCode) {
	o.Names = &v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *EntityScreeningHitAnalysis) GetPhoneNumbers() MatchSummaryCode {
	if o == nil || o.PhoneNumbers == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitAnalysis) GetPhoneNumbersOk() (*MatchSummaryCode, bool) {
	if o == nil || o.PhoneNumbers == nil {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *EntityScreeningHitAnalysis) HasPhoneNumbers() bool {
	if o != nil && o.PhoneNumbers != nil {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given MatchSummaryCode and assigns it to the PhoneNumbers field.
func (o *EntityScreeningHitAnalysis) SetPhoneNumbers(v MatchSummaryCode) {
	o.PhoneNumbers = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *EntityScreeningHitAnalysis) GetUrls() MatchSummaryCode {
	if o == nil || o.Urls == nil {
		var ret MatchSummaryCode
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitAnalysis) GetUrlsOk() (*MatchSummaryCode, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *EntityScreeningHitAnalysis) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given MatchSummaryCode and assigns it to the Urls field.
func (o *EntityScreeningHitAnalysis) SetUrls(v MatchSummaryCode) {
	o.Urls = &v
}

// GetSearchTermsVersion returns the SearchTermsVersion field value
func (o *EntityScreeningHitAnalysis) GetSearchTermsVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SearchTermsVersion
}

// GetSearchTermsVersionOk returns a tuple with the SearchTermsVersion field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitAnalysis) GetSearchTermsVersionOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTermsVersion, true
}

// SetSearchTermsVersion sets field value
func (o *EntityScreeningHitAnalysis) SetSearchTermsVersion(v float32) {
	o.SearchTermsVersion = v
}

func (o EntityScreeningHitAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.EmailAddresses != nil {
		toSerialize["email_addresses"] = o.EmailAddresses
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.PhoneNumbers != nil {
		toSerialize["phone_numbers"] = o.PhoneNumbers
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}
	if true {
		toSerialize["search_terms_version"] = o.SearchTermsVersion
	}
	return json.Marshal(toSerialize)
}

type NullableEntityScreeningHitAnalysis struct {
	value *EntityScreeningHitAnalysis
	isSet bool
}

func (v NullableEntityScreeningHitAnalysis) Get() *EntityScreeningHitAnalysis {
	return v.value
}

func (v *NullableEntityScreeningHitAnalysis) Set(val *EntityScreeningHitAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitAnalysis(val *EntityScreeningHitAnalysis) *NullableEntityScreeningHitAnalysis {
	return &NullableEntityScreeningHitAnalysis{value: val, isSet: true}
}

func (v NullableEntityScreeningHitAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


