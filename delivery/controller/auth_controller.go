package controller

import (
	"go-merchant/config"
	"go-merchant/entity/dto"
	"go-merchant/shared/common"
	"go-merchant/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUC usecase.AuthUsecase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload dto.AuthRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	rsv, err := a.authUC.Login(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(c, rsv, "Ok")
}

func (a *AuthController) logoutHandler(c *gin.Context) {
	var payload dto.AuthRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	rsv, err := a.authUC.Logout(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(c, rsv, "Ok")
}

func (a *AuthController) Route() {
	a.rg.POST(config.AuthLogin, a.loginHandler)
	a.rg.POST(config.AuthLogout, a.logoutHandler)
}

func NewAuthController(authUC usecase.AuthUsecase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUC: authUC, rg: rg}
}
