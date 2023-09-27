package handler

import (
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(adminUseCase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: adminUseCase,
	}
}

func (c *AdminHandler) AdminSignup(ctx *gin.Context) {
	var adminData req.AdminLogin
	err := ctx.BindJSON(&adminData)

}
