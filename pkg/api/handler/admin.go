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
	return
}
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
func (c *AdminHandler) ListDoctorsNotApproved(ctx *gin.Context) {
	_, err := handlerurtl.AdminIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, res.ErrorResponse(400, "please login", err.Error()))
		return
	}
	data, err := c.adminUseCase.ListDoctors(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list", err))
	}
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "list of doctors to approve", data))
}

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
