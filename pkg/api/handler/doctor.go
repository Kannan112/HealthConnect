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

// @Summary Doctor Registration
// @Description Register a doctor.
// @Tags Doctor Authentication
// @Accept json
// @Produce json
// @Param categoryid path int true "Category ID"
// @Param registrationData body req.DoctorRegistration true "Doctor Registration Data"
// @Success 202 {object} res.Response "Registration accepted"
// @Failure 400 {object} res.Response "Bad request or registration failure"
// @Router /doctor/{categoryid}/registration [post]
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

// @Summary Doctor Login
// @Description Logs in a doctor.
// @Tags Doctor Authentication
// @Accept json
// @Produce json
// @Param login body req.DoctorLogin true "Doctor Login Request"
// @Success 202 {object} res.Response "Doctor login successful"
// @Failure 400 {object} res.Response "Bad request or login failure"
// @Router /doctor/login [post]
func (c *DoctorHandler) Login(ctx *gin.Context) {
	var doctorLogin req.DoctorLogin
	if err := ctx.Bind(&doctorLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to bind data", err))
		return
	}
	Token, err := c.doctorUseCase.DoctorLogin(ctx, doctorLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to login", err.Error()))
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("DoctorAuth", Token, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusAccepted, res.SuccessResponse(200, "doctor login success", nil))
}

// @Summary Get Doctor Profile
// @Description Get the profile of a doctor.
// @Tags Doctor
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} res.Response "Doctor profile retrieved successfully"
// @Failure 400 {object} res.Response "Bad request or profile retrieval failure"
// @Router /doctor/profile [get]
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

// @Summary List Doctor Categories
// @Description List available categories for doctors.
// @Tags Doctor
// @Produce json
// @Success 200 {object} res.Response "Categories listed successfully"
// @Failure 400 {object} res.Response "Failed to list categories"
// @Router /doctor/categories [get]
func (c *DoctorHandler) ListCategory(ctx *gin.Context) {
	data, err := c.doctorUseCase.ListCategory(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list category", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "Available Categories", data))
}
