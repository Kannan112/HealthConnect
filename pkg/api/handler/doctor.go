package handler

import (
	"net/http"
	"strconv"

	"github.com/easy-health/pkg/api/middleware/handlerurtl"
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
	paramsId := ctx.Param("categoryid")
	CategoryId, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, res.ErrorResponse(400, "failed to get categoryId", err.Error()))
	}
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
	if err := c.doctorUseCase.Registration(ctx, DoctorRegistrationData, uint(CategoryId)); err != nil {
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
		Message:    "waiting for admin to verifiy",
		Data:       nil,
		Errors:     nil,
	})
	return

}

func (c *DoctorHandler) Login(ctx *gin.Context) {
	var doctorLogin req.DoctorLogin
	if err := ctx.Bind(&doctorLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to bind data", err))
		return
	}
	Token, err := c.doctorUseCase.DoctorLogin(ctx, doctorLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to login", err))
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("DoctorAuth", Token, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusAccepted, res.SuccessResponse(200, "doctor login success", nil))
}

func (c *DoctorHandler) Profile(ctx *gin.Context) {
	doctorId, err := handlerurtl.DoctorIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to get doctorId from context", err.Error()))
		return
	}
	data, err := c.doctorUseCase.DoctorProfile(ctx, doctorId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to get profile", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "Doctor Profile", data))

}

func (c *DoctorHandler) ListCategory(ctx *gin.Context) {
	data, err := c.doctorUseCase.ListCategory(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list category", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "Available Categories", data))
}
