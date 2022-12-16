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
	"fmt"
)

// WebhookEnvironmentValues The Plaid environment the webhook was sent from
type WebhookEnvironmentValues string

var _ = fmt.Printf

// List of WebhookEnvironmentValues
const (
	WEBHOOKENVIRONMENTVALUES_DEVELOPMENT WebhookEnvironmentValues = "development"
	WEBHOOKENVIRONMENTVALUES_SANDBOX WebhookEnvironmentValues = "sandbox"
	WEBHOOKENVIRONMENTVALUES_PRODUCTION WebhookEnvironmentValues = "production"
)

var allowedWebhookEnvironmentValuesEnumValues = []WebhookEnvironmentValues{
	"development",
	"sandbox",
	"production",
}

func (v *WebhookEnvironmentValues) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WebhookEnvironmentValues(value)


	*v = enumTypeValue
	return nil
}

// NewWebhookEnvironmentValuesFromValue returns a pointer to a valid WebhookEnvironmentValues
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookEnvironmentValuesFromValue(v string) (*WebhookEnvironmentValues, error) {
	ev := WebhookEnvironmentValues(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookEnvironmentValues) IsValid() bool {
	for _, existing := range allowedWebhookEnvironmentValuesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookEnvironmentValues value
func (v WebhookEnvironmentValues) Ptr() *WebhookEnvironmentValues {
	return &v
}

type NullableWebhookEnvironmentValues struct {
	value *WebhookEnvironmentValues
	isSet bool
}

func (v NullableWebhookEnvironmentValues) Get() *WebhookEnvironmentValues {
	return v.value
}

func (v *NullableWebhookEnvironmentValues) Set(val *WebhookEnvironmentValues) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEnvironmentValues) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEnvironmentValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEnvironmentValues(val *WebhookEnvironmentValues) *NullableWebhookEnvironmentValues {
	return &NullableWebhookEnvironmentValues{value: val, isSet: true}
}

func (v NullableWebhookEnvironmentValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEnvironmentValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

