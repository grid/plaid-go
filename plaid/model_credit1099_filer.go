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

// Credit1099Filer An object representing a filer used by 1099-K tax documents.
type Credit1099Filer struct {
	Address *CreditPayStubAddress `json:"address,omitempty"`
	// Name of filer.
	Name NullableString `json:"name,omitempty"`
	// Tax identification number of filer.
	Tin NullableString `json:"tin,omitempty"`
	// One of the following values will be provided: Payment Settlement Entity (PSE), Electronic Payment Facilitator (EPF), Other Third Party
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Credit1099Filer Credit1099Filer

// NewCredit1099Filer instantiates a new Credit1099Filer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredit1099Filer() *Credit1099Filer {
	this := Credit1099Filer{}
	return &this
}

// NewCredit1099FilerWithDefaults instantiates a new Credit1099Filer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredit1099FilerWithDefaults() *Credit1099Filer {
	this := Credit1099Filer{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Credit1099Filer) GetAddress() CreditPayStubAddress {
	if o == nil || o.Address == nil {
		var ret CreditPayStubAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit1099Filer) GetAddressOk() (*CreditPayStubAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Credit1099Filer) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CreditPayStubAddress and assigns it to the Address field.
func (o *Credit1099Filer) SetAddress(v CreditPayStubAddress) {
	o.Address = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Filer) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Filer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Credit1099Filer) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Credit1099Filer) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Credit1099Filer) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Credit1099Filer) UnsetName() {
	o.Name.Unset()
}

// GetTin returns the Tin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Filer) GetTin() string {
	if o == nil || o.Tin.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tin.Get()
}

// GetTinOk returns a tuple with the Tin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Filer) GetTinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tin.Get(), o.Tin.IsSet()
}

// HasTin returns a boolean if a field has been set.
func (o *Credit1099Filer) HasTin() bool {
	if o != nil && o.Tin.IsSet() {
		return true
	}

	return false
}

// SetTin gets a reference to the given NullableString and assigns it to the Tin field.
func (o *Credit1099Filer) SetTin(v string) {
	o.Tin.Set(&v)
}
// SetTinNil sets the value for Tin to be an explicit nil
func (o *Credit1099Filer) SetTinNil() {
	o.Tin.Set(nil)
}

// UnsetTin ensures that no value is present for Tin, not even an explicit nil
func (o *Credit1099Filer) UnsetTin() {
	o.Tin.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Filer) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Filer) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Credit1099Filer) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Credit1099Filer) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Credit1099Filer) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Credit1099Filer) UnsetType() {
	o.Type.Unset()
}

func (o Credit1099Filer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Tin.IsSet() {
		toSerialize["tin"] = o.Tin.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Credit1099Filer) UnmarshalJSON(bytes []byte) (err error) {
	varCredit1099Filer := _Credit1099Filer{}

	if err = json.Unmarshal(bytes, &varCredit1099Filer); err == nil {
		*o = Credit1099Filer(varCredit1099Filer)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tin")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCredit1099Filer struct {
	value *Credit1099Filer
	isSet bool
}

func (v NullableCredit1099Filer) Get() *Credit1099Filer {
	return v.value
}

func (v *NullableCredit1099Filer) Set(val *Credit1099Filer) {
	v.value = val
	v.isSet = true
}

func (v NullableCredit1099Filer) IsSet() bool {
	return v.isSet
}

func (v *NullableCredit1099Filer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredit1099Filer(val *Credit1099Filer) *NullableCredit1099Filer {
	return &NullableCredit1099Filer{value: val, isSet: true}
}

func (v NullableCredit1099Filer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredit1099Filer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


