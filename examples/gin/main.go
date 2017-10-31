package main

import (
	"fmt"
	"net/http"
	"os"

	auth_backend "github.com/Valutac/go-edx-openid-auth"
	"github.com/gin-gonic/gin"
)

func main() {

	auth := auth_backend.Init(
		os.Getenv("EDX_CLIENT_ID"),
		os.Getenv("EDX_CLIENT_SECRET"),
		"http://localhost:18300",
		os.Getenv("EDX_MAIN_URL"),
	)

	state := auth_backend.RandomToken(32)

	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.Redirect(302, auth.GetAuthorizationURL(state))
	})
	r.GET("/complete/edx-oidc/", func(c *gin.Context) {
		userInfo, err := auth.Authenticate(state, c.Request.URL.Query())
		if err != nil {
			fmt.Fprintf(c.Writer, "Authentication failed: %s", err.Error())
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		fmt.Fprintf(c.Writer, "User: %s. Email: %s", userInfo.Username, userInfo.Email)
	})
	r.Run(":18300")
}
