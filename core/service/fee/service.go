package fee

import (
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
	return FeeCalculationResponseDto{}, nil
}

func NewFeeService(config config.AppConfig, logger log.Logger, repo orm.UnitOfWorkRepository) FeeService {
	return &feeServiceImpl{
		config, logger, repo,
	}
}
