package fee

import (
	"github.com/magiconair/properties/assert"
	"github.com/smapig/go-ddd-sample/core/domain/entity"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	loggerMocks "github.com/smapig/go-ddd-sample/core/infrastructure/mock/log"
	repoMocks "github.com/smapig/go-ddd-sample/core/infrastructure/mock/orm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestFeeService(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite
	service FeeService
	repo    *repoMocks.UnitOfWorkRepository
	logger  *loggerMocks.Logger
	config  config.AppConfig
}

func (s *ServiceTestSuite) SetupSuite() {

	s.logger = new(loggerMocks.Logger)
	s.repo = new(repoMocks.UnitOfWorkRepository)
}

func (s *ServiceTestSuite) SetupTest() {
	s.service = NewFeeService(s.config, s.logger, s.repo)
}

func (s *ServiceTestSuite) Test_FeeCalculation_Success() {
	input := FeeCalculationRequestDto{
		FromNetwork: "usd",
		ToNetwork:   "ether",
	}
	expectedFiatPaymentFees := []entity.FiatPaymentNetwork{
		{FeeValue: "0.5"},
	}
	expected := FeeCalculationResponseDto{
		Fee: "0.575",
	}
	s.repo.On("GetBy", []entity.FiatPaymentNetwork{}, map[string]interface{}{
		"code": input.FromNetwork,
	}, -1, 0).Return(expectedFiatPaymentFees, nil)

	actual, err := s.service.FeeCalculation(input)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), actual, expected)
}
