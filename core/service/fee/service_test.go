package fee

import (
	"github.com/magiconair/properties/assert"
	"github.com/smapig/go-ddd-sample/core/domain/entity"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	loggerMocks "github.com/smapig/go-ddd-sample/core/infrastructure/mock/log"
	repoMocks "github.com/smapig/go-ddd-sample/core/infrastructure/mock/orm"
	"github.com/stretchr/testify/mock"
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
	expectedQueryResult := []entity.FiatPaymentNetwork{
		{FeeValue: "200"},
	}
	s.repo.On("GetBy", []entity.FiatPaymentNetwork{}, mock.Anything, -1, 0).Return(expectedQueryResult, nil)
	actual, err := s.service.FeeCalculation(FeeCalculationRequestDto{})
	expected := FeeCalculationResponseDto{Fee: expectedQueryResult[0].FeeValue}
	require.NoError(s.T(), err)
	assert.Equal(s.T(), actual, expected)
}
