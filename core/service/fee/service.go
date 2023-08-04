package fee

import (
	"github.com/smapig/go-ddd-sample/core/domain/entity"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/core/infrastructure/orm"
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
	res, err := f.repo.GetBy([]entity.FiatPaymentNetwork{}, map[string]interface{}{
		"fee_value": "300",
	}, -1, 0)

	if err != nil {
		return FeeCalculationResponseDto{}, err
	}

	return FeeCalculationResponseDto{
		Fee: res.([]entity.FiatPaymentNetwork)[0].FeeValue,
	}, nil
}

func NewFeeService(config config.AppConfig, logger log.Logger, repo orm.UnitOfWorkRepository) FeeService {
	return &feeServiceImpl{
		config, logger, repo,
	}
}
