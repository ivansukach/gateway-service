package authentication_service

import (
	"context"
	"fmt"
	"github.com/ivansukach/gateway-service/handlers"
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
	login := c.Request().Header.Get("login")
	_, err := a.client.AddClaims(context.Background(), &protocol.AddClaimsRequest{Claims: claims.Claims})
	fmt.Println("New Claims: ", claims.Claims)
	if err != nil {
		log.Errorf("GRPC Error AddClaims %s", err)
		return err
	}
	responseRefresh, err := a.client.RefreshToken(context.Background(), &protocol.RefreshTokenRequest{Uuid: login, Token: accessToken})
	if err != nil {
		c.String(400, "Something wrong during refreshing your tokens")
		return nil
	}
	refreshToken := responseRefresh.GetRefreshToken()
	accessToken = responseRefresh.GetToken()
	return c.JSON(http.StatusOK, &handlers.TokenModel{AccessToken: accessToken, RefreshToken: refreshToken})
}
