package handler

import (
	"net/http"

	"github.com/easy-health/pkg/api/middleware/handlerurtl"
	"github.com/easy-health/pkg/config"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
	Config      config.Config
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase, Config config.Config) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		Config:      Config,
	}
}

// @Summary Register a new user account
// @Description Register a new user account with the provided details.
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param request body req.UserRegister true "User registration request"
// @Success 202 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/signup [post]
func (c *UserHandler) Register(ctx *gin.Context) {
	var UserReg req.UserRegister
	if err := ctx.BindJSON(&UserReg); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to bind", err))
		return
	}
	if err := c.userUseCase.RegisterUser(ctx, UserReg); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to register", err.Error()))
		return
	}
	ctx.JSON(http.StatusAccepted, res.SuccessResponse(200, "user account registerd", nil))

}

// @Summary User Login
// @Description Logs in a user.
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param login body req.UserLogin true "User Login Request"
// @Success 200 {object} res.Response "Successfully logged in"
// @Failure 400 {object} res.Response "Bad request or login failure"
// @Router /user/login [post]
func (c *UserHandler) Login(ctx *gin.Context) {
	var Login req.UserLogin
	if err := ctx.BindJSON(&Login); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to bind", err))
		return
	}
	Tokens, err := c.userUseCase.UserLogin(ctx, Login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to login", err.Error()))
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("UserAuth", Tokens["access_token"], 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "logined successfuly***", Tokens))
}

// @Summary Logout user from the app
// @Description Logs out a user.
// @Tags User Authentication
// @Produce json
// @Success 200 {object} res.Response "Logged out successfully"
// @Router /user/logout [get]
func (c *UserHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("UserAuth", "", 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "logout successfuly", nil))
}

func (c *UserHandler) ListCategory(ctx *gin.Context) {
	_, err := handlerurtl.UserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to verify", err.Error()))
		return
	}
	categories, err := c.userUseCase.ListCategoryUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to list", err.Error()))
		return
	}
	ctx.JSON(http.StatusBadRequest, res.SuccessResponse(200, "listing categories", categories))

}
