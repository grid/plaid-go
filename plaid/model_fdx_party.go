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

// FDXParty FDX Participant - an entity or person that is a part of a FDX API transaction
type FDXParty struct {
	// Human recognizable common name
	Name string `json:"name"`
	Type FDXPartyType `json:"type"`
	// URI for party, where an end user could learn more about the company or application involved in the data sharing chain
	HomeUri *string `json:"homeUri,omitempty"`
	// URI for a logo asset to be displayed to the end user
	LogoUri *string `json:"logoUri,omitempty"`
	Registry *FDXPartyRegistry `json:"registry,omitempty"`
	// Registered name of party
	RegisteredEntityName *string `json:"registeredEntityName,omitempty"`
	// Registered id of party
	RegisteredEntityId *string `json:"registeredEntityId,omitempty"`
}

// NewFDXParty instantiates a new FDXParty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFDXParty(name string, type_ FDXPartyType) *FDXParty {
	this := FDXParty{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewFDXPartyWithDefaults instantiates a new FDXParty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFDXPartyWithDefaults() *FDXParty {
	this := FDXParty{}
	return &this
}

// GetName returns the Name field value
func (o *FDXParty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FDXParty) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FDXParty) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FDXParty) GetType() FDXPartyType {
	if o == nil {
		var ret FDXPartyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FDXParty) GetTypeOk() (*FDXPartyType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FDXParty) SetType(v FDXPartyType) {
	o.Type = v
}

// GetHomeUri returns the HomeUri field value if set, zero value otherwise.
func (o *FDXParty) GetHomeUri() string {
	if o == nil || o.HomeUri == nil {
		var ret string
		return ret
	}
	return *o.HomeUri
}

// GetHomeUriOk returns a tuple with the HomeUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXParty) GetHomeUriOk() (*string, bool) {
	if o == nil || o.HomeUri == nil {
		return nil, false
	}
	return o.HomeUri, true
}

// HasHomeUri returns a boolean if a field has been set.
func (o *FDXParty) HasHomeUri() bool {
	if o != nil && o.HomeUri != nil {
		return true
	}

	return false
}

// SetHomeUri gets a reference to the given string and assigns it to the HomeUri field.
func (o *FDXParty) SetHomeUri(v string) {
	o.HomeUri = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *FDXParty) GetLogoUri() string {
	if o == nil || o.LogoUri == nil {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXParty) GetLogoUriOk() (*string, bool) {
	if o == nil || o.LogoUri == nil {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *FDXParty) HasLogoUri() bool {
	if o != nil && o.LogoUri != nil {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *FDXParty) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *FDXParty) GetRegistry() FDXPartyRegistry {
	if o == nil || o.Registry == nil {
		var ret FDXPartyRegistry
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXParty) GetRegistryOk() (*FDXPartyRegistry, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *FDXParty) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given FDXPartyRegistry and assigns it to the Registry field.
func (o *FDXParty) SetRegistry(v FDXPartyRegistry) {
	o.Registry = &v
}

// GetRegisteredEntityName returns the RegisteredEntityName field value if set, zero value otherwise.
func (o *FDXParty) GetRegisteredEntityName() string {
	if o == nil || o.RegisteredEntityName == nil {
		var ret string
		return ret
	}
	return *o.RegisteredEntityName
}

// GetRegisteredEntityNameOk returns a tuple with the RegisteredEntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXParty) GetRegisteredEntityNameOk() (*string, bool) {
	if o == nil || o.RegisteredEntityName == nil {
		return nil, false
	}
	return o.RegisteredEntityName, true
}

// HasRegisteredEntityName returns a boolean if a field has been set.
func (o *FDXParty) HasRegisteredEntityName() bool {
	if o != nil && o.RegisteredEntityName != nil {
		return true
	}

	return false
}

// SetRegisteredEntityName gets a reference to the given string and assigns it to the RegisteredEntityName field.
func (o *FDXParty) SetRegisteredEntityName(v string) {
	o.RegisteredEntityName = &v
}

// GetRegisteredEntityId returns the RegisteredEntityId field value if set, zero value otherwise.
func (o *FDXParty) GetRegisteredEntityId() string {
	if o == nil || o.RegisteredEntityId == nil {
		var ret string
		return ret
	}
	return *o.RegisteredEntityId
}

// GetRegisteredEntityIdOk returns a tuple with the RegisteredEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXParty) GetRegisteredEntityIdOk() (*string, bool) {
	if o == nil || o.RegisteredEntityId == nil {
		return nil, false
	}
	return o.RegisteredEntityId, true
}

// HasRegisteredEntityId returns a boolean if a field has been set.
func (o *FDXParty) HasRegisteredEntityId() bool {
	if o != nil && o.RegisteredEntityId != nil {
		return true
	}

	return false
}

// SetRegisteredEntityId gets a reference to the given string and assigns it to the RegisteredEntityId field.
func (o *FDXParty) SetRegisteredEntityId(v string) {
	o.RegisteredEntityId = &v
}

func (o FDXParty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.HomeUri != nil {
		toSerialize["homeUri"] = o.HomeUri
	}
	if o.LogoUri != nil {
		toSerialize["logoUri"] = o.LogoUri
	}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.RegisteredEntityName != nil {
		toSerialize["registeredEntityName"] = o.RegisteredEntityName
	}
	if o.RegisteredEntityId != nil {
		toSerialize["registeredEntityId"] = o.RegisteredEntityId
	}
	return json.Marshal(toSerialize)
}

type NullableFDXParty struct {
	value *FDXParty
	isSet bool
}

func (v NullableFDXParty) Get() *FDXParty {
	return v.value
}

func (v *NullableFDXParty) Set(val *FDXParty) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXParty) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXParty(val *FDXParty) *NullableFDXParty {
	return &NullableFDXParty{value: val, isSet: true}
}

func (v NullableFDXParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


