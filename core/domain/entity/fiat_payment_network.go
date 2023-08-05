package entity

import (
	"strconv"
)

type FiatPaymentNetwork struct {
	BaseEntity
	Code        string
	FeeValue    string
	FeeCurrency string
}

func (e *FiatPaymentNetwork) FeeValueGetter() (float64, error) {
	return strconv.ParseFloat(e.FeeValue, 64)
}
