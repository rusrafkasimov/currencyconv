package models

type ReqResCurrency struct {
	FrCur string `json:"fr_cur"`
	ToCur string `json:"to_cur"`
	FrVal float64 `json:"fr_val"`
	ToVal float64 `json:"to_val"`
}


type JsonCBR struct {
	Date         string `json:"date"`
	PreviousDate string `json:"previous_date"`
	PreviousURL  string `json:"previous_url"`
	Timestamp    string `json:"timestamp"`
	Valute       Val `json:"valute"`
}

type Val struct {
	AUD Detail `json:"aud"`
	AZN Detail `json:"azn"`
	GBP Detail `json:"gbp"`
	AMD Detail `json:"amd"`
	BYN Detail `json:"byn"`
	BGN Detail `json:"bgn"`
	BRL Detail `json:"brl"`
	HUF Detail `json:"huf"`
	HKD Detail `json:"hkd"`
	DKK Detail `json:"dkk"`
	USD Detail `json:"usd"`
	EUR Detail `json:"eur"`
	INR Detail `json:"inr"`
	KZT Detail `json:"kzt"`
	CAD Detail `json:"cad"`
	KGS Detail `json:"kgs"`
	CNY Detail `json:"cny"`
	MDL Detail `json:"mdl"`
	NOK Detail `json:"nok"`
	PLN Detail `json:"pln"`
	RON Detail `json:"ron"`
	XDR Detail `json:"xdr"`
	SGD Detail `json:"sgd"`
	TJS Detail `json:"tjs"`
	TRY Detail `json:"try"`
	TMT Detail `json:"tmt"`
	UZS Detail `json:"uzs"`
	UAH Detail `json:"uah"`
	CZK Detail `json:"czk"`
	SEK Detail `json:"sek"`
	CHF Detail `json:"chf"`
	ZAR Detail `json:"zar"`
	KRW Detail `json:"krw"`
	JPY Detail `json:"jpy"`
}

type Detail struct {
	ID string `json:"id"`
	NumCode string `json:"num_code"`
	CharCode string `json:"char_code"`
	Nominal int16 `json:"nominal"`
	Name string `json:"name"`
	Value float64 `json:"value"`
	Previous float64 `json:"previous"`
}