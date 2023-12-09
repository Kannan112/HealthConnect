package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/easy-health/pkg/domain"
	"github.com/easy-health/pkg/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

// @UserGoogleAuthLoginPage godoc
//
//	@Summary		To load google login page (User)
//	@Description	API for user to load google login page
//	@Id				UserGoogleAuthLoginPage
//	@Tags			User Authentication
//	@Router			/auth/google-auth [get]
//	@Success		200	{object}	res.Response{}	"Successfully google login page loaded"
func (c *UserHandler) UserGoogleAuthLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "google.html", nil)
	// https://shooshtime.com/
}

// UserGoogleAuthInitialize godoc
//
//	@Summary		Initialize google auth (User)
//	@Description	API for user to initialize google auth
//	@Id				UserGoogleAuthInitialize
//	@Tags			User Authentication
//	@Router			/auth/google-auth/initialize [get]
func (c *UserHandler) UserGoogleAuthInitialize(ctx *gin.Context) {

	//setup the google provider
	goauthClientID := c.Config.ClientID
	goauthClientSecret := c.Config.ClientSecret
	goauthRedirectUrl := c.Config.RedirectURL

	log.Fatalln(goauthClientSecret, goauthClientID, goauthRedirectUrl)
	goth.UseProviders(
		google.New(goauthClientID, goauthClientSecret, goauthRedirectUrl, "email", "profile"),
	)
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

// UserGoogleAuthCallBack godoc
//
//	@Summary		Google auth callback (User)
//	@Description	API for google to callback after authentication
//	@Id				UserGoogleAuthCallBack
//	@Tags			User Authentication
//	@Router			/auth/google-auth/callback [post]
//	@Success		200	{object}	res.Response{}	"Successfully logged in with google"
//	@Failure		500	{object}	response.Response{}	"Failed Login with google"
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
