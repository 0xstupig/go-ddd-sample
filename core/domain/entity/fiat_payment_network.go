package entity

type FiatPaymentNetwork struct {
	BaseEntity
	Code        string
	FeeValue    string
	FeeCurrency string
}
