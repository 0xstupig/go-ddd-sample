package controller

import (
	"github.com/gin-gonic/gin"
	gin2 "github.com/smapig/go-ddd-sample/core/infrastructure/hosting/gin"
	"github.com/smapig/go-ddd-sample/core/service/fee"
)

type FeeController interface {
	FeeCalculation(ctx *gin.Context)
}

func (c controllerImpl) FeeCalculation(ctx *gin.Context) {
	req := &fee.FeeCalculationRequestDto{}
	if err := gin2.BindData(ctx, &req, gin2.BindType_Query); err != nil {
		gin2.ResponseBadRequest(ctx, err)
		return
	}

	res, err := c.feeService.FeeCalculation(*req)

	if err != nil {
		gin2.ResponseInternalServerError(ctx, err)
		return
	}

	gin2.ResponseSuccess(ctx, res)
}
