// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Currency Indicates the associated currency for an amount of money. Values correspond
// to [ISO 4217](https://wikipedia.org/wiki/ISO_4217).
//
// swagger:model Currency
type Currency string

const (

	// CurrencyUNKNOWNCURRENCY captures enum value "UNKNOWN_CURRENCY"
	CurrencyUNKNOWNCURRENCY Currency = "UNKNOWN_CURRENCY"

	// CurrencyAED captures enum value "AED"
	CurrencyAED Currency = "AED"

	// CurrencyAFN captures enum value "AFN"
	CurrencyAFN Currency = "AFN"

	// CurrencyALL captures enum value "ALL"
	CurrencyALL Currency = "ALL"

	// CurrencyAMD captures enum value "AMD"
	CurrencyAMD Currency = "AMD"

	// CurrencyANG captures enum value "ANG"
	CurrencyANG Currency = "ANG"

	// CurrencyAOA captures enum value "AOA"
	CurrencyAOA Currency = "AOA"

	// CurrencyARS captures enum value "ARS"
	CurrencyARS Currency = "ARS"

	// CurrencyAUD captures enum value "AUD"
	CurrencyAUD Currency = "AUD"

	// CurrencyAWG captures enum value "AWG"
	CurrencyAWG Currency = "AWG"

	// CurrencyAZN captures enum value "AZN"
	CurrencyAZN Currency = "AZN"

	// CurrencyBAM captures enum value "BAM"
	CurrencyBAM Currency = "BAM"

	// CurrencyBBD captures enum value "BBD"
	CurrencyBBD Currency = "BBD"

	// CurrencyBDT captures enum value "BDT"
	CurrencyBDT Currency = "BDT"

	// CurrencyBGN captures enum value "BGN"
	CurrencyBGN Currency = "BGN"

	// CurrencyBHD captures enum value "BHD"
	CurrencyBHD Currency = "BHD"

	// CurrencyBIF captures enum value "BIF"
	CurrencyBIF Currency = "BIF"

	// CurrencyBMD captures enum value "BMD"
	CurrencyBMD Currency = "BMD"

	// CurrencyBND captures enum value "BND"
	CurrencyBND Currency = "BND"

	// CurrencyBOB captures enum value "BOB"
	CurrencyBOB Currency = "BOB"

	// CurrencyBOV captures enum value "BOV"
	CurrencyBOV Currency = "BOV"

	// CurrencyBRL captures enum value "BRL"
	CurrencyBRL Currency = "BRL"

	// CurrencyBSD captures enum value "BSD"
	CurrencyBSD Currency = "BSD"

	// CurrencyBTN captures enum value "BTN"
	CurrencyBTN Currency = "BTN"

	// CurrencyBWP captures enum value "BWP"
	CurrencyBWP Currency = "BWP"

	// CurrencyBYR captures enum value "BYR"
	CurrencyBYR Currency = "BYR"

	// CurrencyBZD captures enum value "BZD"
	CurrencyBZD Currency = "BZD"

	// CurrencyCAD captures enum value "CAD"
	CurrencyCAD Currency = "CAD"

	// CurrencyCDF captures enum value "CDF"
	CurrencyCDF Currency = "CDF"

	// CurrencyCHE captures enum value "CHE"
	CurrencyCHE Currency = "CHE"

	// CurrencyCHF captures enum value "CHF"
	CurrencyCHF Currency = "CHF"

	// CurrencyCHW captures enum value "CHW"
	CurrencyCHW Currency = "CHW"

	// CurrencyCLF captures enum value "CLF"
	CurrencyCLF Currency = "CLF"

	// CurrencyCLP captures enum value "CLP"
	CurrencyCLP Currency = "CLP"

	// CurrencyCNY captures enum value "CNY"
	CurrencyCNY Currency = "CNY"

	// CurrencyCOP captures enum value "COP"
	CurrencyCOP Currency = "COP"

	// CurrencyCOU captures enum value "COU"
	CurrencyCOU Currency = "COU"

	// CurrencyCRC captures enum value "CRC"
	CurrencyCRC Currency = "CRC"

	// CurrencyCUC captures enum value "CUC"
	CurrencyCUC Currency = "CUC"

	// CurrencyCUP captures enum value "CUP"
	CurrencyCUP Currency = "CUP"

	// CurrencyCVE captures enum value "CVE"
	CurrencyCVE Currency = "CVE"

	// CurrencyCZK captures enum value "CZK"
	CurrencyCZK Currency = "CZK"

	// CurrencyDJF captures enum value "DJF"
	CurrencyDJF Currency = "DJF"

	// CurrencyDKK captures enum value "DKK"
	CurrencyDKK Currency = "DKK"

	// CurrencyDOP captures enum value "DOP"
	CurrencyDOP Currency = "DOP"

	// CurrencyDZD captures enum value "DZD"
	CurrencyDZD Currency = "DZD"

	// CurrencyEGP captures enum value "EGP"
	CurrencyEGP Currency = "EGP"

	// CurrencyERN captures enum value "ERN"
	CurrencyERN Currency = "ERN"

	// CurrencyETB captures enum value "ETB"
	CurrencyETB Currency = "ETB"

	// CurrencyEUR captures enum value "EUR"
	CurrencyEUR Currency = "EUR"

	// CurrencyFJD captures enum value "FJD"
	CurrencyFJD Currency = "FJD"

	// CurrencyFKP captures enum value "FKP"
	CurrencyFKP Currency = "FKP"

	// CurrencyGBP captures enum value "GBP"
	CurrencyGBP Currency = "GBP"

	// CurrencyGEL captures enum value "GEL"
	CurrencyGEL Currency = "GEL"

	// CurrencyGHS captures enum value "GHS"
	CurrencyGHS Currency = "GHS"

	// CurrencyGIP captures enum value "GIP"
	CurrencyGIP Currency = "GIP"

	// CurrencyGMD captures enum value "GMD"
	CurrencyGMD Currency = "GMD"

	// CurrencyGNF captures enum value "GNF"
	CurrencyGNF Currency = "GNF"

	// CurrencyGTQ captures enum value "GTQ"
	CurrencyGTQ Currency = "GTQ"

	// CurrencyGYD captures enum value "GYD"
	CurrencyGYD Currency = "GYD"

	// CurrencyHKD captures enum value "HKD"
	CurrencyHKD Currency = "HKD"

	// CurrencyHNL captures enum value "HNL"
	CurrencyHNL Currency = "HNL"

	// CurrencyHRK captures enum value "HRK"
	CurrencyHRK Currency = "HRK"

	// CurrencyHTG captures enum value "HTG"
	CurrencyHTG Currency = "HTG"

	// CurrencyHUF captures enum value "HUF"
	CurrencyHUF Currency = "HUF"

	// CurrencyIDR captures enum value "IDR"
	CurrencyIDR Currency = "IDR"

	// CurrencyILS captures enum value "ILS"
	CurrencyILS Currency = "ILS"

	// CurrencyINR captures enum value "INR"
	CurrencyINR Currency = "INR"

	// CurrencyIQD captures enum value "IQD"
	CurrencyIQD Currency = "IQD"

	// CurrencyIRR captures enum value "IRR"
	CurrencyIRR Currency = "IRR"

	// CurrencyISK captures enum value "ISK"
	CurrencyISK Currency = "ISK"

	// CurrencyJMD captures enum value "JMD"
	CurrencyJMD Currency = "JMD"

	// CurrencyJOD captures enum value "JOD"
	CurrencyJOD Currency = "JOD"

	// CurrencyJPY captures enum value "JPY"
	CurrencyJPY Currency = "JPY"

	// CurrencyKES captures enum value "KES"
	CurrencyKES Currency = "KES"

	// CurrencyKGS captures enum value "KGS"
	CurrencyKGS Currency = "KGS"

	// CurrencyKHR captures enum value "KHR"
	CurrencyKHR Currency = "KHR"

	// CurrencyKMF captures enum value "KMF"
	CurrencyKMF Currency = "KMF"

	// CurrencyKPW captures enum value "KPW"
	CurrencyKPW Currency = "KPW"

	// CurrencyKRW captures enum value "KRW"
	CurrencyKRW Currency = "KRW"

	// CurrencyKWD captures enum value "KWD"
	CurrencyKWD Currency = "KWD"

	// CurrencyKYD captures enum value "KYD"
	CurrencyKYD Currency = "KYD"

	// CurrencyKZT captures enum value "KZT"
	CurrencyKZT Currency = "KZT"

	// CurrencyLAK captures enum value "LAK"
	CurrencyLAK Currency = "LAK"

	// CurrencyLBP captures enum value "LBP"
	CurrencyLBP Currency = "LBP"

	// CurrencyLKR captures enum value "LKR"
	CurrencyLKR Currency = "LKR"

	// CurrencyLRD captures enum value "LRD"
	CurrencyLRD Currency = "LRD"

	// CurrencyLSL captures enum value "LSL"
	CurrencyLSL Currency = "LSL"

	// CurrencyLTL captures enum value "LTL"
	CurrencyLTL Currency = "LTL"

	// CurrencyLVL captures enum value "LVL"
	CurrencyLVL Currency = "LVL"

	// CurrencyLYD captures enum value "LYD"
	CurrencyLYD Currency = "LYD"

	// CurrencyMAD captures enum value "MAD"
	CurrencyMAD Currency = "MAD"

	// CurrencyMDL captures enum value "MDL"
	CurrencyMDL Currency = "MDL"

	// CurrencyMGA captures enum value "MGA"
	CurrencyMGA Currency = "MGA"

	// CurrencyMKD captures enum value "MKD"
	CurrencyMKD Currency = "MKD"

	// CurrencyMMK captures enum value "MMK"
	CurrencyMMK Currency = "MMK"

	// CurrencyMNT captures enum value "MNT"
	CurrencyMNT Currency = "MNT"

	// CurrencyMOP captures enum value "MOP"
	CurrencyMOP Currency = "MOP"

	// CurrencyMRO captures enum value "MRO"
	CurrencyMRO Currency = "MRO"

	// CurrencyMUR captures enum value "MUR"
	CurrencyMUR Currency = "MUR"

	// CurrencyMVR captures enum value "MVR"
	CurrencyMVR Currency = "MVR"

	// CurrencyMWK captures enum value "MWK"
	CurrencyMWK Currency = "MWK"

	// CurrencyMXN captures enum value "MXN"
	CurrencyMXN Currency = "MXN"

	// CurrencyMXV captures enum value "MXV"
	CurrencyMXV Currency = "MXV"

	// CurrencyMYR captures enum value "MYR"
	CurrencyMYR Currency = "MYR"

	// CurrencyMZN captures enum value "MZN"
	CurrencyMZN Currency = "MZN"

	// CurrencyNAD captures enum value "NAD"
	CurrencyNAD Currency = "NAD"

	// CurrencyNGN captures enum value "NGN"
	CurrencyNGN Currency = "NGN"

	// CurrencyNIO captures enum value "NIO"
	CurrencyNIO Currency = "NIO"

	// CurrencyNOK captures enum value "NOK"
	CurrencyNOK Currency = "NOK"

	// CurrencyNPR captures enum value "NPR"
	CurrencyNPR Currency = "NPR"

	// CurrencyNZD captures enum value "NZD"
	CurrencyNZD Currency = "NZD"

	// CurrencyOMR captures enum value "OMR"
	CurrencyOMR Currency = "OMR"

	// CurrencyPAB captures enum value "PAB"
	CurrencyPAB Currency = "PAB"

	// CurrencyPEN captures enum value "PEN"
	CurrencyPEN Currency = "PEN"

	// CurrencyPGK captures enum value "PGK"
	CurrencyPGK Currency = "PGK"

	// CurrencyPHP captures enum value "PHP"
	CurrencyPHP Currency = "PHP"

	// CurrencyPKR captures enum value "PKR"
	CurrencyPKR Currency = "PKR"

	// CurrencyPLN captures enum value "PLN"
	CurrencyPLN Currency = "PLN"

	// CurrencyPYG captures enum value "PYG"
	CurrencyPYG Currency = "PYG"

	// CurrencyQAR captures enum value "QAR"
	CurrencyQAR Currency = "QAR"

	// CurrencyRON captures enum value "RON"
	CurrencyRON Currency = "RON"

	// CurrencyRSD captures enum value "RSD"
	CurrencyRSD Currency = "RSD"

	// CurrencyRUB captures enum value "RUB"
	CurrencyRUB Currency = "RUB"

	// CurrencyRWF captures enum value "RWF"
	CurrencyRWF Currency = "RWF"

	// CurrencySAR captures enum value "SAR"
	CurrencySAR Currency = "SAR"

	// CurrencySBD captures enum value "SBD"
	CurrencySBD Currency = "SBD"

	// CurrencySCR captures enum value "SCR"
	CurrencySCR Currency = "SCR"

	// CurrencySDG captures enum value "SDG"
	CurrencySDG Currency = "SDG"

	// CurrencySEK captures enum value "SEK"
	CurrencySEK Currency = "SEK"

	// CurrencySGD captures enum value "SGD"
	CurrencySGD Currency = "SGD"

	// CurrencySHP captures enum value "SHP"
	CurrencySHP Currency = "SHP"

	// CurrencySLL captures enum value "SLL"
	CurrencySLL Currency = "SLL"

	// CurrencySOS captures enum value "SOS"
	CurrencySOS Currency = "SOS"

	// CurrencySRD captures enum value "SRD"
	CurrencySRD Currency = "SRD"

	// CurrencySSP captures enum value "SSP"
	CurrencySSP Currency = "SSP"

	// CurrencySTD captures enum value "STD"
	CurrencySTD Currency = "STD"

	// CurrencySVC captures enum value "SVC"
	CurrencySVC Currency = "SVC"

	// CurrencySYP captures enum value "SYP"
	CurrencySYP Currency = "SYP"

	// CurrencySZL captures enum value "SZL"
	CurrencySZL Currency = "SZL"

	// CurrencyTHB captures enum value "THB"
	CurrencyTHB Currency = "THB"

	// CurrencyTJS captures enum value "TJS"
	CurrencyTJS Currency = "TJS"

	// CurrencyTMT captures enum value "TMT"
	CurrencyTMT Currency = "TMT"

	// CurrencyTND captures enum value "TND"
	CurrencyTND Currency = "TND"

	// CurrencyTOP captures enum value "TOP"
	CurrencyTOP Currency = "TOP"

	// CurrencyTRY captures enum value "TRY"
	CurrencyTRY Currency = "TRY"

	// CurrencyTTD captures enum value "TTD"
	CurrencyTTD Currency = "TTD"

	// CurrencyTWD captures enum value "TWD"
	CurrencyTWD Currency = "TWD"

	// CurrencyTZS captures enum value "TZS"
	CurrencyTZS Currency = "TZS"

	// CurrencyUAH captures enum value "UAH"
	CurrencyUAH Currency = "UAH"

	// CurrencyUGX captures enum value "UGX"
	CurrencyUGX Currency = "UGX"

	// CurrencyUSD captures enum value "USD"
	CurrencyUSD Currency = "USD"

	// CurrencyUSN captures enum value "USN"
	CurrencyUSN Currency = "USN"

	// CurrencyUSS captures enum value "USS"
	CurrencyUSS Currency = "USS"

	// CurrencyUYI captures enum value "UYI"
	CurrencyUYI Currency = "UYI"

	// CurrencyUYU captures enum value "UYU"
	CurrencyUYU Currency = "UYU"

	// CurrencyUZS captures enum value "UZS"
	CurrencyUZS Currency = "UZS"

	// CurrencyVEF captures enum value "VEF"
	CurrencyVEF Currency = "VEF"

	// CurrencyVND captures enum value "VND"
	CurrencyVND Currency = "VND"

	// CurrencyVUV captures enum value "VUV"
	CurrencyVUV Currency = "VUV"

	// CurrencyWST captures enum value "WST"
	CurrencyWST Currency = "WST"

	// CurrencyXAF captures enum value "XAF"
	CurrencyXAF Currency = "XAF"

	// CurrencyXAG captures enum value "XAG"
	CurrencyXAG Currency = "XAG"

	// CurrencyXAU captures enum value "XAU"
	CurrencyXAU Currency = "XAU"

	// CurrencyXBA captures enum value "XBA"
	CurrencyXBA Currency = "XBA"

	// CurrencyXBB captures enum value "XBB"
	CurrencyXBB Currency = "XBB"

	// CurrencyXBC captures enum value "XBC"
	CurrencyXBC Currency = "XBC"

	// CurrencyXBD captures enum value "XBD"
	CurrencyXBD Currency = "XBD"

	// CurrencyXCD captures enum value "XCD"
	CurrencyXCD Currency = "XCD"

	// CurrencyXDR captures enum value "XDR"
	CurrencyXDR Currency = "XDR"

	// CurrencyXOF captures enum value "XOF"
	CurrencyXOF Currency = "XOF"

	// CurrencyXPD captures enum value "XPD"
	CurrencyXPD Currency = "XPD"

	// CurrencyXPF captures enum value "XPF"
	CurrencyXPF Currency = "XPF"

	// CurrencyXPT captures enum value "XPT"
	CurrencyXPT Currency = "XPT"

	// CurrencyXTS captures enum value "XTS"
	CurrencyXTS Currency = "XTS"

	// CurrencyXXX captures enum value "XXX"
	CurrencyXXX Currency = "XXX"

	// CurrencyYER captures enum value "YER"
	CurrencyYER Currency = "YER"

	// CurrencyZAR captures enum value "ZAR"
	CurrencyZAR Currency = "ZAR"

	// CurrencyZMK captures enum value "ZMK"
	CurrencyZMK Currency = "ZMK"

	// CurrencyZMW captures enum value "ZMW"
	CurrencyZMW Currency = "ZMW"

	// CurrencyBTC captures enum value "BTC"
	CurrencyBTC Currency = "BTC"
)

// for schema
var currencyEnum []interface{}

func init() {
	var res []Currency
	if err := json.Unmarshal([]byte(`["UNKNOWN_CURRENCY","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BOV","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHE","CHF","CHW","CLF","CLP","CNY","COP","COU","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","INR","IQD","IRR","ISK","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MXV","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SRD","SSP","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","USD","USN","USS","UYI","UYU","UZS","VEF","VND","VUV","WST","XAF","XAG","XAU","XBA","XBB","XBC","XBD","XCD","XDR","XOF","XPD","XPF","XPT","XTS","XXX","YER","ZAR","ZMK","ZMW","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		currencyEnum = append(currencyEnum, v)
	}
}

func (m Currency) validateCurrencyEnum(path, location string, value Currency) error {
	if err := validate.Enum(path, location, value, currencyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this currency
func (m Currency) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCurrencyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}