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
	"time"
)

// AssetReport An object representing an Asset Report
type AssetReport struct {
	// A unique ID identifying an Asset Report. Like all Plaid identifiers, this ID is case sensitive.
	AssetReportId string `json:"asset_report_id"`
	// An identifier you determine and submit for the Asset Report.
	ClientReportId NullableString `json:"client_report_id"`
	// The date and time when the Asset Report was created, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (e.g. \"2018-04-12T03:32:11Z\").
	DateGenerated time.Time `json:"date_generated"`
	// The duration of transaction history you requested
	DaysRequested float32 `json:"days_requested"`
	User AssetReportUser `json:"user"`
	// Data returned by Plaid about each of the Items included in the Asset Report.
	Items []AssetReportItem `json:"items"`
	AdditionalProperties map[string]interface{}
}

type _AssetReport AssetReport

// NewAssetReport instantiates a new AssetReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReport(assetReportId string, clientReportId NullableString, dateGenerated time.Time, daysRequested float32, user AssetReportUser, items []AssetReportItem) *AssetReport {
	this := AssetReport{}
	this.AssetReportId = assetReportId
	this.ClientReportId = clientReportId
	this.DateGenerated = dateGenerated
	this.DaysRequested = daysRequested
	this.User = user
	this.Items = items
	return &this
}

// NewAssetReportWithDefaults instantiates a new AssetReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportWithDefaults() *AssetReport {
	this := AssetReport{}
	return &this
}

// GetAssetReportId returns the AssetReportId field value
func (o *AssetReport) GetAssetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value
// and a boolean to check if the value has been set.
func (o *AssetReport) GetAssetReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportId, true
}

// SetAssetReportId sets field value
func (o *AssetReport) SetAssetReportId(v string) {
	o.AssetReportId = v
}

// GetClientReportId returns the ClientReportId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssetReport) GetClientReportId() string {
	if o == nil || o.ClientReportId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientReportId.Get()
}

// GetClientReportIdOk returns a tuple with the ClientReportId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReport) GetClientReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientReportId.Get(), o.ClientReportId.IsSet()
}

// SetClientReportId sets field value
func (o *AssetReport) SetClientReportId(v string) {
	o.ClientReportId.Set(&v)
}

// GetDateGenerated returns the DateGenerated field value
func (o *AssetReport) GetDateGenerated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateGenerated
}

// GetDateGeneratedOk returns a tuple with the DateGenerated field value
// and a boolean to check if the value has been set.
func (o *AssetReport) GetDateGeneratedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateGenerated, true
}

// SetDateGenerated sets field value
func (o *AssetReport) SetDateGenerated(v time.Time) {
	o.DateGenerated = v
}

// GetDaysRequested returns the DaysRequested field value
func (o *AssetReport) GetDaysRequested() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value
// and a boolean to check if the value has been set.
func (o *AssetReport) GetDaysRequestedOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DaysRequested, true
}

// SetDaysRequested sets field value
func (o *AssetReport) SetDaysRequested(v float32) {
	o.DaysRequested = v
}

// GetUser returns the User field value
func (o *AssetReport) GetUser() AssetReportUser {
	if o == nil {
		var ret AssetReportUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AssetReport) GetUserOk() (*AssetReportUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AssetReport) SetUser(v AssetReportUser) {
	o.User = v
}

// GetItems returns the Items field value
func (o *AssetReport) GetItems() []AssetReportItem {
	if o == nil {
		var ret []AssetReportItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AssetReport) GetItemsOk() (*[]AssetReportItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *AssetReport) SetItems(v []AssetReportItem) {
	o.Items = v
}

func (o AssetReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset_report_id"] = o.AssetReportId
	}
	if true {
		toSerialize["client_report_id"] = o.ClientReportId.Get()
	}
	if true {
		toSerialize["date_generated"] = o.DateGenerated
	}
	if true {
		toSerialize["days_requested"] = o.DaysRequested
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReport) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReport := _AssetReport{}

	if err = json.Unmarshal(bytes, &varAssetReport); err == nil {
		*o = AssetReport(varAssetReport)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_report_id")
		delete(additionalProperties, "client_report_id")
		delete(additionalProperties, "date_generated")
		delete(additionalProperties, "days_requested")
		delete(additionalProperties, "user")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReport struct {
	value *AssetReport
	isSet bool
}

func (v NullableAssetReport) Get() *AssetReport {
	return v.value
}

func (v *NullableAssetReport) Set(val *AssetReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReport(val *AssetReport) *NullableAssetReport {
	return &NullableAssetReport{value: val, isSet: true}
}

func (v NullableAssetReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


