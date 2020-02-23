package authentication_service

import (
	"context"
	"github.com/ivansukach/gateway-service/handlers"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Auth) SignIn(c echo.Context) error {
	log.Info("SignIn")
	user := new(handlers.UserModel)

	if err := c.Bind(user); err != nil {
		log.Errorf("echo.Context Error SignIn %s", err)
		return err
	}
	responseAuth, err := a.client.SignIn(context.Background(),
		&protocol.SignInRequest{
			Login:    user.Login,
			Password: user.Password,
		})
	if err != nil {
		log.Errorf("GRPC Error SignIn %s", err)
		return echo.ErrUnauthorized
	}
	accessToken := responseAuth.GetToken()
	refreshToken := responseAuth.GetRefreshToken()
	c.Request().Header.Set("Authorization", accessToken)
	c.Request().Header.Set("RefreshToken", refreshToken)
	return c.JSON(http.StatusOK, &handlers.TokenModel{AccessToken: accessToken, RefreshToken: refreshToken})
}
