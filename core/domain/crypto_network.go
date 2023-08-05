package domain

import "errors"

type CryptoNetwork string
type CryptoNetworkFee struct {
	CryptoNetwork
	Amount float64
}

var Ether CryptoNetworkFee = CryptoNetworkFee{
	CryptoNetwork: EtherNetwork,
	Amount:        0.05,
}

var Solana CryptoNetworkFee = CryptoNetworkFee{
	CryptoNetwork: SolanaNetwork,
	Amount:        0.01,
}

var Apt CryptoNetworkFee = CryptoNetworkFee{
	CryptoNetwork: AptNetwork,
	Amount:        0.05,
}

var NetworkFees []CryptoNetworkFee = []CryptoNetworkFee{Ether, Solana, Apt}

const (
	EtherNetwork  CryptoNetwork = "ether"
	SolanaNetwork CryptoNetwork = "solana"
	AptNetwork    CryptoNetwork = "apt"
)

func (c CryptoNetwork) String() string {
	return string(c)
}

func GasFee(network CryptoNetwork) (*CryptoNetworkFee, error) {
	for _, fee := range NetworkFees {
		if fee.CryptoNetwork == network {
			return &fee, nil
		}
	}

	return nil, errors.New("Network doesn't support yet!")
}
