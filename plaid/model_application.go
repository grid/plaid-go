/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.161.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Application Metadata about the application
type Application struct {
	// This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect.
	ApplicationId string `json:"application_id"`
	// The name of the application
	Name string `json:"name"`
	// A human-readable name of the application for display purposes
	DisplayName NullableString `json:"display_name"`
	// The date this application was granted production access at Plaid in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC.
	JoinDate string `json:"join_date"`
	// A URL that links to the application logo image.
	LogoUrl NullableString `json:"logo_url"`
	// The URL for the application's website
	ApplicationUrl NullableString `json:"application_url"`
	// A string provided by the connected app stating why they use their respective enabled products.
	ReasonForAccess NullableString `json:"reason_for_access"`
	// A string representing client’s broad use case as assessed by Plaid.
	UseCase NullableString `json:"use_case"`
	// A string representing the name of client’s legal entity.
	CompanyLegalName NullableString `json:"company_legal_name"`
	// A string representing the city of the client’s headquarters.
	City NullableString `json:"city"`
	// A string representing the region of the client’s headquarters.
	Region NullableString `json:"region"`
	// A string representing the postal code of the client’s headquarters.
	PostalCode NullableString `json:"postal_code"`
	// A string representing the country code of the client’s headquarters.
	CountryCode NullableString `json:"country_code"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(applicationId string, name string, displayName NullableString, joinDate string, logoUrl NullableString, applicationUrl NullableString, reasonForAccess NullableString, useCase NullableString, companyLegalName NullableString, city NullableString, region NullableString, postalCode NullableString, countryCode NullableString) *Application {
	this := Application{}
	this.ApplicationId = applicationId
	this.Name = name
	this.DisplayName = displayName
	this.JoinDate = joinDate
	this.LogoUrl = logoUrl
	this.ApplicationUrl = applicationUrl
	this.ReasonForAccess = reasonForAccess
	this.UseCase = useCase
	this.CompanyLegalName = companyLegalName
	this.City = city
	this.Region = region
	this.PostalCode = postalCode
	this.CountryCode = countryCode
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *Application) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *Application) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *Application) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetName returns the Name field value
func (o *Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Application) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// SetDisplayName sets field value
func (o *Application) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// GetJoinDate returns the JoinDate field value
func (o *Application) GetJoinDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinDate
}

// GetJoinDateOk returns a tuple with the JoinDate field value
// and a boolean to check if the value has been set.
func (o *Application) GetJoinDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.JoinDate, true
}

// SetJoinDate sets field value
func (o *Application) SetJoinDate(v string) {
	o.JoinDate = v
}

// GetLogoUrl returns the LogoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// SetLogoUrl sets field value
func (o *Application) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

// GetApplicationUrl returns the ApplicationUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetApplicationUrl() string {
	if o == nil || o.ApplicationUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ApplicationUrl.Get()
}

// GetApplicationUrlOk returns a tuple with the ApplicationUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetApplicationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationUrl.Get(), o.ApplicationUrl.IsSet()
}

// SetApplicationUrl sets field value
func (o *Application) SetApplicationUrl(v string) {
	o.ApplicationUrl.Set(&v)
}

// GetReasonForAccess returns the ReasonForAccess field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetReasonForAccess() string {
	if o == nil || o.ReasonForAccess.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReasonForAccess.Get()
}

// GetReasonForAccessOk returns a tuple with the ReasonForAccess field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetReasonForAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReasonForAccess.Get(), o.ReasonForAccess.IsSet()
}

// SetReasonForAccess sets field value
func (o *Application) SetReasonForAccess(v string) {
	o.ReasonForAccess.Set(&v)
}

// GetUseCase returns the UseCase field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetUseCase() string {
	if o == nil || o.UseCase.Get() == nil {
		var ret string
		return ret
	}

	return *o.UseCase.Get()
}

// GetUseCaseOk returns a tuple with the UseCase field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetUseCaseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UseCase.Get(), o.UseCase.IsSet()
}

// SetUseCase sets field value
func (o *Application) SetUseCase(v string) {
	o.UseCase.Set(&v)
}

// GetCompanyLegalName returns the CompanyLegalName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetCompanyLegalName() string {
	if o == nil || o.CompanyLegalName.Get() == nil {
		var ret string
		return ret
	}

	return *o.CompanyLegalName.Get()
}

// GetCompanyLegalNameOk returns a tuple with the CompanyLegalName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetCompanyLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompanyLegalName.Get(), o.CompanyLegalName.IsSet()
}

// SetCompanyLegalName sets field value
func (o *Application) SetCompanyLegalName(v string) {
	o.CompanyLegalName.Set(&v)
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *Application) SetCity(v string) {
	o.City.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *Application) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *Application) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetCountryCode returns the CountryCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// SetCountryCode sets field value
func (o *Application) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application_id"] = o.ApplicationId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if true {
		toSerialize["join_date"] = o.JoinDate
	}
	if true {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if true {
		toSerialize["application_url"] = o.ApplicationUrl.Get()
	}
	if true {
		toSerialize["reason_for_access"] = o.ReasonForAccess.Get()
	}
	if true {
		toSerialize["use_case"] = o.UseCase.Get()
	}
	if true {
		toSerialize["company_legal_name"] = o.CompanyLegalName.Get()
	}
	if true {
		toSerialize["city"] = o.City.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


