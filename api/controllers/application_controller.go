package controllers

import (
	"github.com/currencyconv/api/models"
	"reflect"
)

var CurData models.JsonCBR

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



