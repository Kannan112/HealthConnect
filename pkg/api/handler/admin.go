package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/easy-health/pkg/api/middleware/handlerurtl"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase  services.AdminUseCase
	doctorUseCase services.DoctorUseCase
}

func NewAdminHandler(adminUseCase services.AdminUseCase, doctorUseCase services.DoctorUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase:  adminUseCase,
		doctorUseCase: doctorUseCase,
	}
}

// CreateAdmin
// @Summary Create a new admin from admin panel
// @ID AdminSignup
// @Description admin creation
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body req.AdminLogin true "New Admin details"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/create [post]
func (c *AdminHandler) AdminSignup(ctx *gin.Context) {
	var admin req.AdminLogin
	err := ctx.BindJSON(&admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "Failed to bind data", err.Error()))
		return
	}
	err = c.adminUseCase.AdminSignup(ctx, admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "Failed to signup", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "successful", nil))

}

// AdminLogin godoc
// @Summary Admin login
// @Description Logs in an admin user
// @Tags Admin
// @Accept json
// @Produce json
// @Param adminLogin body req.AdminLogin true "Admin login data"
// @Success 200 {object} res.Response
// @failed 400 {object} res.Response
// @Router /admin/login [post]
func (c *AdminHandler) AdminLogin(ctx *gin.Context) {
	var adminLogin req.AdminLogin
	err := ctx.BindJSON(&adminLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "Failed to bind data", err.Error()))
		return
	}

	tokens, err := c.adminUseCase.AdminLogin(ctx, adminLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "Failed to login", err.Error()))
		return
	}

	accessToken, hasAccessToken := tokens["access_token"]
	_, hasRefreshToken := tokens["refresh_token"]

	if !hasAccessToken || !hasRefreshToken {
		// Handle the case where the required tokens are missing.
		ctx.JSON(http.StatusInternalServerError, res.ErrorResponse(500, "Tokens not generated", "Tokens are missing"))
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("AdminAuth", accessToken, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "logined successfuly", tokens))

}

// @Summary Admin Logout
// @Description Logs out an admin user.
// @Tags Admin
// @Produce json
// @Success 200 {object} res.Response
// @Security BearerTokenAuth
// @Router /admin/logout [get]]
func (c *AdminHandler) AdminLogout(ctx *gin.Context) {
	id, err := handlerurtl.AdminIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "you are not login", err.Error()))
		return
	}
	ctx.SetCookie("AdminAuth", "", 3600*24*30, "", "", false, true)
	if err := c.adminUseCase.AdminLogout(ctx, id); err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "please login", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "logoutsuccessfuly", nil))
}

// @Summary create a new category
// @Description Create a new category based on the provided data
// @Tags Categories
// @Accept json
// @Produce json
// @Param data body req.Category true "Category data to create"
// @Success 202 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {string} string "Unauthorized"
// @Failure 422 {object} res.Response
// @Security BearerTokenAuth
// @Router /admin/categories [post]
func (c *AdminHandler) CreateCategory(ctx *gin.Context) {
	_, err := handlerurtl.AdminIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "please login", err.Error()))
		return
	}
	var data req.Category
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to bind", err))
		return
	}
	if err := c.adminUseCase.CreateCateogry(ctx, data); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(422, "failed to create", err.Error()))
		return
	}
	ctx.JSON(http.StatusAccepted, res.SuccessResponse(200, "created", data))
}

// @Summary list categories
// @Description List categories
// @ID list-categories
// @Tags Categories
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Param page query integer false "Page number (default 1)"
// @Param count query integer false "Number of items per page (default 10)"
// @Security BearerTokenAuth
// @Router /admin/categories [get]
func (c *AdminHandler) ListCategory(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")    // Set a default value for page if not provided
	count := ctx.DefaultQuery("count", "10") // Set a default value for count if not provided
	pageNo, err := strconv.Atoi(page)
	if err != nil {
		return
	}

	counts, err := strconv.Atoi(count)
	if err != nil {
		return
	}

	data, err := c.adminUseCase.ListCategory(ctx, pageNo, counts)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list category", err))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "category list", data))
}

// @Summary delete categories
// @Description Delete a category by ID
// @ID delete-category
// @Tags Categories
// @Produce json
// @Param id query int true "Category ID" Format(int64)
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Security BearerTokenAuth
// @Router /admin/categories [delete]
func (c *AdminHandler) DeleteCategory(ctx *gin.Context) {
	_, err := handlerurtl.AdminIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "please login", err.Error()))
		return
	}
	paramsID := ctx.Param("id")

	categoryID, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "cant find categoryId", err))
		return
	}

	if err := c.adminUseCase.DeleteCategory(ctx, categoryID); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed", err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// UpdateCategory updates a category.
// @Summary Update a category by ID
// @Description Update a category with new name and description by providing ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param name query string true "New name of the category"
// @Param description query string true "New description of the category"
// @Success 200 {object} res.Response "Successfully updated category details"
// @Failure 400 {object} res.Response "Please login"
// @Failure 404 {object} res.Response "Failed to get ID" or "Failed to update category"
// @Security BearerTokenAuth
// @Router /admin/categories/{id} [patch]
func (c *AdminHandler) UpdateCategory(ctx *gin.Context) {

	name := ctx.Query("name")
	description := ctx.Query("description")
	category := req.Category{
		Name:        name,
		Description: description,
	}
	id := ctx.Param("id")
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, res.ErrorResponse(404, "failed to get id", err.Error()))
		return
	}
	if err := c.adminUseCase.UpdateCategory(ctx, category, uint(categoryID)); err != nil {
		ctx.JSON(http.StatusBadGateway, res.ErrorResponse(404, "failed to update category", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "updated category details", nil))
}

// @Summary List Doctors Not Approved
// @Description Get a list of doctors that are not yet approved
// @ID list-doctors-not-approved
// @Tags  Admin Dashboard
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Security BearerTokenAuth
// @Router /admin/doctors/not-approved [get]
func (c *AdminHandler) ListDoctorsNotApproved(ctx *gin.Context) {

	data, err := c.adminUseCase.ApprovePending(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "list of doctors to approve", data))
}

// @Summary Approve a Doctor
// @Description Approve a doctor by ID
// @ID approve-doctor
// @Tags Admin Dashboard
// @Produce json
// @Param id path int true "Doctor ID" Format(int64)
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Security BearerTokenAuth
// @Router /admin/doctors/approve/{id} [patch]
func (c *AdminHandler) ApproveDoctor(ctx *gin.Context) {

	paramsId := ctx.Param("id")
	doctorId, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "cant find doctorId", err))
		return
	}

	err = c.adminUseCase.AdminVerify(ctx, doctorId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to verify", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "approved", nil))
}

// GetDoctorProfile godoc
// @Summary Get doctor profile by ID
// @Description Get the profile of a doctor by their ID
// @Tags Admin Dashboard
// @Accept json
// @Produce json
// @Param id path integer true "Doctor ID"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /admin/doctor-profile/{id} [get]
func (c *AdminHandler) GetDoctorProfile(ctx *gin.Context) {

	paramsId := ctx.Param("id")
	doctorId, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "cant find doctorId", err))
		return
	}
	DoctorProfile, err := c.doctorUseCase.DoctorProfile(ctx, doctorId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to get details", err))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "Doctor Profile", DoctorProfile))

}

// VerifiedDoctors godoc
// @Summary Get verified doctors
// @Description Get a list of verified doctors with pagination
// @Tags Admin Dashboard
// @Accept  json
// @Produce  json
// @Param page query integer false "Page number (default 1)"
// @Param count query integer false "Number of items per page (default 10)"
// @Success 202 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /admin/doctors/verified [get]
func (c *AdminHandler) VerifiedDoctors(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")    // Set a default value for page if not provided
	count := ctx.DefaultQuery("count", "10") // Set a default value for count if not provided
	pageNo, pageError := strconv.Atoi(page)
	counts, countError := strconv.Atoi(count)
	if pageError != nil || countError != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to get details", errors.Join(pageError, countError)))
		return
	}

	data, err := c.adminUseCase.ListVerifiedDoctores(ctx, pageNo, counts)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list", err))
		return
	}
	ctx.JSON(http.StatusAccepted, res.SuccessResponse(200, "verified doctors", data))
}
