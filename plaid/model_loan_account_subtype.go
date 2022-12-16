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

// LoanAccountSubtype Valid account subtypes for loan accounts. For a list containing descriptions of each subtype, see [Account schemas](https://plaid.com/docs/api/accounts/#StandaloneAccountType-loan).
type LoanAccountSubtype string

var _ = fmt.Printf

// List of LoanAccountSubtype
const (
	LOANACCOUNTSUBTYPE_AUTO LoanAccountSubtype = "auto"
	LOANACCOUNTSUBTYPE_BUSINESS LoanAccountSubtype = "business"
	LOANACCOUNTSUBTYPE_COMMERCIAL LoanAccountSubtype = "commercial"
	LOANACCOUNTSUBTYPE_CONSTRUCTION LoanAccountSubtype = "construction"
	LOANACCOUNTSUBTYPE_CONSUMER LoanAccountSubtype = "consumer"
	LOANACCOUNTSUBTYPE_HOME_EQUITY LoanAccountSubtype = "home equity"
	LOANACCOUNTSUBTYPE_LOAN LoanAccountSubtype = "loan"
	LOANACCOUNTSUBTYPE_MORTGAGE LoanAccountSubtype = "mortgage"
	LOANACCOUNTSUBTYPE_LINE_OF_CREDIT LoanAccountSubtype = "line of credit"
	LOANACCOUNTSUBTYPE_STUDENT LoanAccountSubtype = "student"
	LOANACCOUNTSUBTYPE_OTHER LoanAccountSubtype = "other"
	LOANACCOUNTSUBTYPE_ALL LoanAccountSubtype = "all"
)

var allowedLoanAccountSubtypeEnumValues = []LoanAccountSubtype{
	"auto",
	"business",
	"commercial",
	"construction",
	"consumer",
	"home equity",
	"loan",
	"mortgage",
	"line of credit",
	"student",
	"other",
	"all",
}

func (v *LoanAccountSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := LoanAccountSubtype(value)


	*v = enumTypeValue
	return nil
}

// NewLoanAccountSubtypeFromValue returns a pointer to a valid LoanAccountSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoanAccountSubtypeFromValue(v string) (*LoanAccountSubtype, error) {
	ev := LoanAccountSubtype(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoanAccountSubtype) IsValid() bool {
	for _, existing := range allowedLoanAccountSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoanAccountSubtype value
func (v LoanAccountSubtype) Ptr() *LoanAccountSubtype {
	return &v
}

type NullableLoanAccountSubtype struct {
	value *LoanAccountSubtype
	isSet bool
}

func (v NullableLoanAccountSubtype) Get() *LoanAccountSubtype {
	return v.value
}

func (v *NullableLoanAccountSubtype) Set(val *LoanAccountSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableLoanAccountSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableLoanAccountSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoanAccountSubtype(val *LoanAccountSubtype) *NullableLoanAccountSubtype {
	return &NullableLoanAccountSubtype{value: val, isSet: true}
}

func (v NullableLoanAccountSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoanAccountSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

