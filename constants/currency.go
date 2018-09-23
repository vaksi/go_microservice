/*  Currency.go
*
* @Author:             Audy Vaksi <vaksipranata@gmail.com>
* @Date:               September 07, 2018
* @Last Modified by:   @vaksi
* @Last Modified time: 07/09/18 00:29
 */

package constants

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

const (
	// retrieve from http://country.io/currency.json
	currencyCountryJSON = `{"BD":"BDT","BE":"EUR","BF":"XOF","BG":"BGN","BA":"BAM","BB":"BBD","WF":"XPF","BL":"EUR","BM":"BMD","BN":"BND","BO":"BOB","BH":"BHD","BI":"BIF","BJ":"XOF","BT":"BTN","JM":"JMD","BV":"NOK","BW":"BWP","WS":"WST","BQ":"USD","BR":"BRL","BS":"BSD","JE":"GBP","BY":"BYR","BZ":"BZD","RU":"RUB","RW":"RWF","RS":"RSD","TL":"USD","RE":"EUR","TM":"TMT","TJ":"TJS","RO":"RON","TK":"NZD","GW":"XOF","GU":"USD","GT":"GTQ","GS":"GBP","GR":"EUR","GQ":"XAF","GP":"EUR","JP":"JPY","GY":"GYD","GG":"GBP","GF":"EUR","GE":"GEL","GD":"XCD","GB":"GBP","GA":"XAF","SV":"USD","GN":"GNF","GM":"GMD","GL":"DKK","GI":"GIP","GH":"GHS","OM":"OMR","TN":"TND","JO":"JOD","HR":"HRK","HT":"HTG","HU":"HUF","HK":"HKD","HN":"HNL","HM":"AUD","VE":"VEF","PR":"USD","PS":"ILS","PW":"USD","PT":"EUR","SJ":"NOK","PY":"PYG","IQ":"IQD","PA":"PAB","PF":"XPF","PG":"PGK","PE":"PEN","PK":"PKR","PH":"PHP","PN":"NZD","PL":"PLN","PM":"EUR","ZM":"ZMK","EH":"MAD","EE":"EUR","EG":"EGP","ZA":"ZAR","EC":"USD","IT":"EUR","VN":"VND","SB":"SBD","ET":"ETB","SO":"SOS","ZW":"ZWL","SA":"SAR","ES":"EUR","ER":"ERN","ME":"EUR","MD":"MDL","MG":"MGA","MF":"EUR","MA":"MAD","MC":"EUR","UZ":"UZS","MM":"MMK","ML":"XOF","MO":"MOP","MN":"MNT","MH":"USD","MK":"MKD","MU":"MUR","MT":"EUR","MW":"MWK","MV":"MVR","MQ":"EUR","MP":"USD","MS":"XCD","MR":"MRO","IM":"GBP","UG":"UGX","TZ":"TZS","MY":"MYR","MX":"MXN","IL":"ILS","FR":"EUR","IO":"USD","SH":"SHP","FI":"EUR","FJ":"FJD","FK":"FKP","FM":"USD","FO":"DKK","NI":"NIO","NL":"EUR","NO":"NOK","NA":"NAD","VU":"VUV","NC":"XPF","NE":"XOF","NF":"AUD","NG":"NGN","NZ":"NZD","NP":"NPR","NR":"AUD","NU":"NZD","CK":"NZD","XK":"EUR","CI":"XOF","CH":"CHF","CO":"COP","CN":"CNY","CM":"XAF","CL":"CLP","CC":"AUD","CA":"CAD","CG":"XAF","CF":"XAF","CD":"CDF","CZ":"CZK","CY":"EUR","CX":"AUD","CR":"CRC","CW":"ANG","CV":"CVE","CU":"CUP","SZ":"SZL","SY":"SYP","SX":"ANG","KG":"KGS","KE":"KES","SS":"SSP","SR":"SRD","KI":"AUD","KH":"KHR","KN":"XCD","KM":"KMF","ST":"STD","SK":"EUR","KR":"KRW","SI":"EUR","KP":"KPW","KW":"KWD","SN":"XOF","SM":"EUR","SL":"SLL","SC":"SCR","KZ":"KZT","KY":"KYD","SG":"SGD","SE":"SEK","SD":"SDG","DO":"DOP","DM":"XCD","DJ":"DJF","DK":"DKK","VG":"USD","DE":"EUR","YE":"YER","DZ":"DZD","US":"USD","UY":"UYU","YT":"EUR","UM":"USD","LB":"LBP","LC":"XCD","LA":"LAK","TV":"AUD","TW":"TWD","TT":"TTD","TR":"TRY","LK":"LKR","LI":"CHF","LV":"EUR","TO":"TOP","LT":"LTL","LU":"EUR","LR":"LRD","LS":"LSL","TH":"THB","TF":"EUR","TG":"XOF","TD":"XAF","TC":"USD","LY":"LYD","VA":"EUR","VC":"XCD","AE":"AED","AD":"EUR","AG":"XCD","AF":"AFN","AI":"XCD","VI":"USD","IS":"ISK","IR":"IRR","AM":"AMD","AL":"ALL","AO":"AOA","AS":"USD","AR":"ARS","AU":"AUD","AT":"EUR","AW":"AWG","IN":"INR","AX":"EUR","AZ":"AZN","IE":"EUR","ID":"IDR","UA":"UAH","QA":"QAR","MZ":"MZN"}`
	MissingData         = "insufficient data"
)

// Currency define type Currency
type Currency string

// Code of function validation currency
func (c Currency) Code() string {
	currenciesMap := make(map[string]interface{})

	// Decode JSON into our map
	err := json.Unmarshal([]byte(currencyCountryJSON), &currenciesMap)
	if err != nil {
		log.Error(err)
		return ""
	}
	for _, code := range currenciesMap {
		if string(c) == code {
			return string(c)
		}

	}
	return ""
}
