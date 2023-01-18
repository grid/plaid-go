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

// PhysicalDocumentImages URLs for downloading original and cropped images for this document submission. The URLs are designed to only allow downloading, not hot linking, so the URL will only serve the document image for 60 seconds before expiring. The expiration time is 60 seconds after the `GET` request for the associated Identity Verification attempt. A new expiring URL is generated with each request, so you can always rerequest the Identity Verification attempt if one of your URLs expires.
type PhysicalDocumentImages struct {
	// Temporary URL that expires after 60 seconds for downloading the uncropped original image of the front of the document.
	OriginalFront string `json:"original_front"`
	// Temporary URL that expires after 60 seconds for downloading the original image of the back of the document. Might be null if the back of the document was not collected.
	OriginalBack NullableString `json:"original_back"`
	// Temporary URL that expires after 60 seconds for downloading a cropped image containing just the front of the document.
	CroppedFront NullableString `json:"cropped_front"`
	// Temporary URL that expires after 60 seconds for downloading a cropped image containing just the back of the document. Might be null if the back of the document was not collected.
	CroppedBack NullableString `json:"cropped_back"`
	// Temporary URL that expires after 60 seconds for downloading a crop of just the user's face from the document image. Might be null if the document does not contain a face photo.
	Face NullableString `json:"face"`
}

// NewPhysicalDocumentImages instantiates a new PhysicalDocumentImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalDocumentImages(originalFront string, originalBack NullableString, croppedFront NullableString, croppedBack NullableString, face NullableString) *PhysicalDocumentImages {
	this := PhysicalDocumentImages{}
	this.OriginalFront = originalFront
	this.OriginalBack = originalBack
	this.CroppedFront = croppedFront
	this.CroppedBack = croppedBack
	this.Face = face
	return &this
}

// NewPhysicalDocumentImagesWithDefaults instantiates a new PhysicalDocumentImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalDocumentImagesWithDefaults() *PhysicalDocumentImages {
	this := PhysicalDocumentImages{}
	return &this
}

// GetOriginalFront returns the OriginalFront field value
func (o *PhysicalDocumentImages) GetOriginalFront() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalFront
}

// GetOriginalFrontOk returns a tuple with the OriginalFront field value
// and a boolean to check if the value has been set.
func (o *PhysicalDocumentImages) GetOriginalFrontOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginalFront, true
}

// SetOriginalFront sets field value
func (o *PhysicalDocumentImages) SetOriginalFront(v string) {
	o.OriginalFront = v
}

// GetOriginalBack returns the OriginalBack field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PhysicalDocumentImages) GetOriginalBack() string {
	if o == nil || o.OriginalBack.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginalBack.Get()
}

// GetOriginalBackOk returns a tuple with the OriginalBack field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalDocumentImages) GetOriginalBackOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalBack.Get(), o.OriginalBack.IsSet()
}

// SetOriginalBack sets field value
func (o *PhysicalDocumentImages) SetOriginalBack(v string) {
	o.OriginalBack.Set(&v)
}

// GetCroppedFront returns the CroppedFront field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PhysicalDocumentImages) GetCroppedFront() string {
	if o == nil || o.CroppedFront.Get() == nil {
		var ret string
		return ret
	}

	return *o.CroppedFront.Get()
}

// GetCroppedFrontOk returns a tuple with the CroppedFront field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalDocumentImages) GetCroppedFrontOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CroppedFront.Get(), o.CroppedFront.IsSet()
}

// SetCroppedFront sets field value
func (o *PhysicalDocumentImages) SetCroppedFront(v string) {
	o.CroppedFront.Set(&v)
}

// GetCroppedBack returns the CroppedBack field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PhysicalDocumentImages) GetCroppedBack() string {
	if o == nil || o.CroppedBack.Get() == nil {
		var ret string
		return ret
	}

	return *o.CroppedBack.Get()
}

// GetCroppedBackOk returns a tuple with the CroppedBack field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalDocumentImages) GetCroppedBackOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CroppedBack.Get(), o.CroppedBack.IsSet()
}

// SetCroppedBack sets field value
func (o *PhysicalDocumentImages) SetCroppedBack(v string) {
	o.CroppedBack.Set(&v)
}

// GetFace returns the Face field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PhysicalDocumentImages) GetFace() string {
	if o == nil || o.Face.Get() == nil {
		var ret string
		return ret
	}

	return *o.Face.Get()
}

// GetFaceOk returns a tuple with the Face field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalDocumentImages) GetFaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Face.Get(), o.Face.IsSet()
}

// SetFace sets field value
func (o *PhysicalDocumentImages) SetFace(v string) {
	o.Face.Set(&v)
}

func (o PhysicalDocumentImages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["original_front"] = o.OriginalFront
	}
	if true {
		toSerialize["original_back"] = o.OriginalBack.Get()
	}
	if true {
		toSerialize["cropped_front"] = o.CroppedFront.Get()
	}
	if true {
		toSerialize["cropped_back"] = o.CroppedBack.Get()
	}
	if true {
		toSerialize["face"] = o.Face.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalDocumentImages struct {
	value *PhysicalDocumentImages
	isSet bool
}

func (v NullablePhysicalDocumentImages) Get() *PhysicalDocumentImages {
	return v.value
}

func (v *NullablePhysicalDocumentImages) Set(val *PhysicalDocumentImages) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalDocumentImages) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalDocumentImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalDocumentImages(val *PhysicalDocumentImages) *NullablePhysicalDocumentImages {
	return &NullablePhysicalDocumentImages{value: val, isSet: true}
}

func (v NullablePhysicalDocumentImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalDocumentImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


