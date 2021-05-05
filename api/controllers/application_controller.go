package controllers

import (
	"encoding/json"
	"github.com/currencyconv/api/models"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

var CurData models.JsonCBR

var Currency = [...]string{"RUB", "AUD", "AZN", "GBP", "BYN", "BGN", "BRL", "HUF", "HKD", "DKK", "USD", "EUR", "INR", "KZT", "CAD", "KGS", "CNY", "MDL", "NOK", "PLN", "RON", "XDR", "SGD", "TJS", "TRY", "TMT", "UZS", "UAH", "CZK", "SEK", "CHF", "ZAR", "KRW", "JPY"}

func Convertation (from, to string, val float64) (res float64) {
	frate, torate, fnom, tnom := CheckRates(from, to)
	if fnom != 1 || tnom != 1 {
		frate /= float64(fnom)
		torate /= float64(tnom)
	}
	resp := (frate/torate) * val
	return resp
}

func CheckRates (f, t string) (from, to float64, fromn, ton int64) {

	cbr := reflect.ValueOf(CurData.Valute)
	if f == "RUB" || t == "RUB" {
		if f == "RUB" {
			from = 1.0
			fromn = 1

			to = (cbr.FieldByName(t).FieldByName("Value")).Float()
			ton = (cbr.FieldByName(t).FieldByName("Nominal")).Int()
			return from, to, fromn, ton
		}
		if t == "RUB" {
			from = (cbr.FieldByName(f).FieldByName("Value")).Float()
			fromn = (cbr.FieldByName(f).FieldByName("Nominal")).Int()
			to = 1.0
			ton = 1
			return from, to, fromn, ton
		}

	}

	from = (cbr.FieldByName(f).FieldByName("Value")).Float()
	fromn = (cbr.FieldByName(f).FieldByName("Nominal")).Int()
	to = (cbr.FieldByName(t).FieldByName("Value")).Float()
	ton = (cbr.FieldByName(t).FieldByName("Nominal")).Int()

    return from, to, fromn, ton

}

func CurrencyValid(f, t string) bool {

	var fm, to = false, false
	for _, name := range Currency {
		if name == f {
          fm = true
		}
		if name == t {
			to = true
		}
	}

	if (fm && to) == true {
		return true
	}

	return false
}

func ParseCBR() {

	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	var req = http.Client{
		Timeout: time.Second * 120,
	}

	res, err := req.Get(url)
	if err != nil {
		errorLog.Printf("Parse_CBR_JSON")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		errorLog.Printf("Read_JSON")
	}

	var data models.JsonCBR
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		errorLog.Printf("Unmarshalling")
	}

	CurData = data
	infoLog.Printf("CBR_parse_success")

}
