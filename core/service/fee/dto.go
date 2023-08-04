package fee

type FeeCalculationRequestDto struct {
	FromNetwork        string   `json:"from_network"`
	FromAmount         int      `json:"from_amount"`
	FromAsset          string   `json:"from_asset"`
	ToAsset            string   `json:"to_asset"`
	ToNetwork          string   `json:"to_network"`
	FeeAsset           string   `json:"fee_asset"`
	CustomerTier       string   `json:"customer_tier"`
	AvailableProviders []string `json:"available_providers"`
}

type FeeCalculationResponseDto struct {
	Fee      string `json:"fee"`
	Provider string `json:"provider"`
	Asset    string `json:"asset"`
}
