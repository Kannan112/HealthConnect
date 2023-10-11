package handler

import (
	"net/http"

	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"github.com/easy-health/pkg/utils/res"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary Register a new user account
// @Description Register a new user account with the provided details.
// @Tags User
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
	return
}

func (c *UserHandler) Login(ctx *gin.Context) {
	var Login req.UserLogin
	if err := ctx.BindJSON(&Login); err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to bind", err))
		return
	}
	SignedString, err := c.userUseCase.UserLogin(ctx, Login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ErrorResponse(400, "failed to login", err.Error()))
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("UserAuth", SignedString, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, res.SuccessResponse(200, "logined successfuly", nil))
}
