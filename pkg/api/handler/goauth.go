package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/easy-health/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
)

func SetUpConfig(c *config.Config) *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  "http://localhost:8000/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}

func (c *UserHandler) GoogleLogin(ctx *gin.Context) {
	googleConfig := SetUpConfig(&c.Config)
	url := googleConfig.AuthCodeURL("randomstate")
	fmt.Println(url)
	ctx.Redirect(http.StatusSeeOther, url)
}

func (c *UserHandler) GoogleAuthCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	if state != "randomstate" {
		fmt.Fprintln(ctx.Writer, "state doesn't match")
		return
	}
	code := ctx.Query("code")

	googleConfig := SetUpConfig(&c.Config)
	token, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		// Handle the error
		fmt.Println("Error exchanging code for token:", err)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		// Handle the error
		fmt.Println("Error getting user info from Google:", err)
		return
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(ctx.Writer, "failed to read")
		return
	}
	fmt.Fprintln(ctx.Writer, string(result))

}
