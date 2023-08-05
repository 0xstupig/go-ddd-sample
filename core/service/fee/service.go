package fee

import (
	"errors"
	"github.com/smapig/go-ddd-sample/core/domain"
	"github.com/smapig/go-ddd-sample/core/domain/entity"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/core/infrastructure/orm"
	"strconv"
)

type FeeService interface {
	FeeCalculation(data FeeCalculationRequestDto) (FeeCalculationResponseDto, error)
}

type feeServiceImpl struct {
	config config.AppConfig
	logger log.Logger
	repo   orm.UnitOfWorkRepository
}

func (f feeServiceImpl) FeeCalculation(data FeeCalculationRequestDto) (FeeCalculationResponseDto, error) {
	fiatPaymentNetworkFees, err := f.repo.GetBy([]entity.FiatPaymentNetwork{}, map[string]interface{}{
		"code": data.FromNetwork,
	}, -1, 0)
	if err != nil {
		f.logger.Errorf("Fiat Payment Network Fees %+v", err)
		return FeeCalculationResponseDto{}, err
	}

	var gasFee *domain.CryptoNetworkFee
	gasFee, err = domain.GasFee(domain.CryptoNetwork(data.ToNetwork))
	if err != nil {
		f.logger.Errorf("GasFee error %+v", err)
		return FeeCalculationResponseDto{}, err
	}

	var fiatCryptoNetworkFee float64
	fiatCryptoNetworkFee, err = exchange(gasFee, data.FromNetwork)

	if err != nil {
		f.logger.Errorf("Exchange error %+v", err)
		return FeeCalculationResponseDto{}, err
	}

	var fiatPaymentNetworkFee float64
	fiatPaymentNetworkFee, err = fiatPaymentNetworkFees.([]entity.FiatPaymentNetwork)[0].FeeValueGetter()
	if err != nil {
		f.logger.Errorf("Convert Fiat payment network fee error %+v", err)
		return FeeCalculationResponseDto{}, err
	}

	rv := fiatCryptoNetworkFee + fiatPaymentNetworkFee

	return FeeCalculationResponseDto{
		Fee: strconv.FormatFloat(rv, 'f', -1, 64),
	}, nil
}

func exchange(cryptoNetworkFee *domain.CryptoNetworkFee, fiatNetwork string) (float64, error) {
	exchangeRate := map[domain.CryptoNetwork]interface{}{
		domain.EtherNetwork: map[string]interface{}{
			"USD": 1.5,
			"VND": 23000,
		},
		domain.SolanaNetwork: map[string]interface{}{
			"USD": 0.5,
			"VND": 15000,
		},
		domain.AptNetwork: map[string]interface{}{
			"USD": 0.1,
			"VND": 1000,
		},
	}

	corrFiatFees := exchangeRate[cryptoNetworkFee.CryptoNetwork]
	if corrFiatFees == nil {
		return 0, errors.New("Crypto network invalid!")
	}

	exchangeRateValue := corrFiatFees.(map[string]interface{})[fiatNetwork]
	if exchangeRateValue == nil {
		return 0, errors.New("Fiat network invalid!")
	}

	return cryptoNetworkFee.Amount * float64(exchangeRateValue.(float64)), nil
}

func NewFeeService(config config.AppConfig, logger log.Logger, repo orm.UnitOfWorkRepository) FeeService {
	return &feeServiceImpl{
		config, logger, repo,
	}
}
