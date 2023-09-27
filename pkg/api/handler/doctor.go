package handler

import (
	"net/http"

	interfaces "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	doctorUseCase interfaces.DoctorUseCase
}

func NewDoctorHandler(doctorusecase interfaces.DoctorUseCase) *DoctorHandler {
	return &DoctorHandler{
		doctorUseCase: doctorusecase,
	}
}

func (c *DoctorHandler) DoctorRegistration(ctx *gin.Context) {
	var DoctorRegistrationData req.DoctorRegistration
	if err := ctx.BindJSON(&DoctorRegistrationData); err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind the data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	if err := c.doctorUseCase.Registration(ctx, DoctorRegistrationData); err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       nil,
		Errors:     nil,
	})
	return

}
