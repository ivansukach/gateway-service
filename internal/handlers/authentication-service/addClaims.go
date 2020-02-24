package authentication_service

import (
	"context"
	"fmt"
	"github.com/ivansukach/gateway-service/internal/handlers"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Auth) AddClaims(c echo.Context) error {
	log.Info("AddClaims")
	var claims handlers.ClaimsModel
	if err := c.Bind(&claims); err != nil {
		log.Errorf("echo.Context Error AddClaims %s", err)
		return err
	}
	accessToken := c.Request().Header.Get("Authorization")
	refreshToken := c.Request().Header.Get("Refresh")
	login := c.Request().Header.Get("login")
	_, err := a.client.AddClaims(context.Background(), &protocol.AddClaimsRequest{Login: login, Claims: claims.Claims})
	fmt.Println("New Claims: ", claims.Claims)
	if err != nil {
		log.Errorf("gRPC Error AddClaims %s", err)
		return err
	}
	responseRefresh, err := a.client.RefreshToken(context.Background(), &protocol.RefreshTokenRequest{Token: accessToken, TokenRefresh: refreshToken})
	if err != nil {
		c.String(400, "Something wrong during refreshing your tokens")
		return nil
	}
	refreshToken = responseRefresh.GetRefreshToken()
	accessToken = responseRefresh.GetToken()
	return c.JSON(http.StatusOK, &handlers.TokenModel{AccessToken: accessToken, RefreshToken: refreshToken})
}
