package controller

import (
	"go-merchant/config"
	"go-merchant/delivery/middleware"
	"go-merchant/entity"
	"go-merchant/shared/common"
	"go-merchant/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentUC      usecase.PaymentUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (p *PaymentController) createHandler(c *gin.Context) {
	var payload entity.Payment
	// bind payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// eksekusi usecase
	payment, err := p.paymentUC.CreatePayment(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// return response jika berhasil
	common.SendCreateResponse(c, payment, "Created")
}

func (p *PaymentController) Route() {
	p.rg.POST(config.PaymentCreate, p.authMiddleware.RequireToken(), p.createHandler)
}

func NewPaymentController(paymentUC usecase.PaymentUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *PaymentController {
	return &PaymentController{paymentUC: paymentUC, rg: rg, authMiddleware: authMiddleware}
}
