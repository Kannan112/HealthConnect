package handler

import (
	"net/http"
	"strconv"

	"github.com/easy-health/pkg/api/middleware/handlerurtl"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
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

// @title Go + Gin E-Commerce API
// @version 1.0.0
// @description TechDeck is an E-commerce platform to purchase and sell Electronic itmes
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath
// @query.collection.format multi

// CreateAdmin
// @Summary Create a new admin from admin panel
// @ID AdminSignup
// @Description admin creation
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param admin body req.AdminLogin true "New Admin details"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/createadmin [post]
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
// @Tags User Authentication
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

	Token, err := c.adminUseCase.AdminLogin(ctx, adminLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "Failed to login", err.Error()))
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("AdminAuth", Token, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "logined successfuly", nil))

}

// @Summary Admin Logout
// @Description Logs out an admin user.
// @Tags User Authentication
// @Produce json
// @Success 200 {object} res.Response
// @Router /admin/logout [get]]
func (c *AdminHandler) AdminLogout(ctx *gin.Context) {
	ctx.SetCookie("AdminAuth", "", 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "logoutsuccessfuly", nil))
}

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

// @Summary List Categories
// @Description List categories
// @ID list-categories
// @Tags Admin
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Router /admin/list-category [get]
func (c *AdminHandler) ListCategory(ctx *gin.Context) {
	_, err := handlerurtl.AdminIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "please login", err.Error()))
		return
	}
	//add admin auth token check
	data, err := c.adminUseCase.ListCategory(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list category", err))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "category list", data))
}

// @Summary Delete Category
// @Description Delete a category by ID
// @ID delete-category
// @Tags Admin
// @Produce json
// @Param id query int true "Category ID" Format(int64)
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Router /admin/delete-category [delete]
func (c *AdminHandler) DeleteCategory(ctx *gin.Context) {
	_, err := handlerurtl.AdminIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "please login", err.Error()))
		return
	}
	idParam := ctx.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	if err := c.adminUseCase.DeleteCategory(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed", err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// doctors

// @Summary List Doctors Not Approved
// @Description Get a list of doctors that are not yet approved
// @ID list-doctors-not-approved
// @Tags Admin
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Router /admin/list-doctors-not-approved [get]
func (c *AdminHandler) ListDoctorsNotApproved(ctx *gin.Context) {
	_, err := handlerurtl.AdminIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "please login", err.Error()))
		return
	}
	data, err := c.adminUseCase.ApprovePending(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list", err))
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "list of doctors to approve", data))
}

// @Summary Approve a Doctor
// @Description Approve a doctor by ID
// @ID approve-doctor
// @Tags Admin
// @Produce json
// @Param id path int true "Doctor ID" Format(int64)
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Router /admin/approve-doctor/{id} [post]
func (c *AdminHandler) ApproveDoctor(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	doctorId, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "cant find doctorId", err))
		return
	}
	_, err = handlerurtl.AdminIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "please login", err.Error()))
		return
	}

	err = c.adminUseCase.AdminVerify(ctx, doctorId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to verify", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "approved", nil))
}
