package handler

import (
	"fmt"
	"net/http"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func (c *UserHandler) UserGoogleAuthLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "view/google.html", nil)
}

func (c *UserHandler) UserGoogleAuthInitialize(ctx *gin.Context) {

	//setup the google provider
	goauthClientID := c.Config.ClientID
	goauthClientSecret := c.Config.ClientSecret
	goauthRedirectUrl := c.Config.RedirectURL

	goth.UseProviders(
		google.New(goauthClientID, goauthClientSecret, goauthRedirectUrl, "email", "profile"),
	)
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (c *UserHandler) UserGoogleAuthCallBack(ctx *gin.Context) {
	googleUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		res.ErrorResponse(http.StatusInternalServerError, "Failed to get user details from google", err.Error())
		return
	}
	user := domain.User{
		FirstName: googleUser.FirstName,
		LastName:  googleUser.LastName,
		Email:     googleUser.Email,
	}
	fmt.Println(user)
}
