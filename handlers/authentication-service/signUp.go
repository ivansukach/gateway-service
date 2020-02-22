package authentication_service

import (
	"context"
	"github.com/ivansukach/gateway-service/handlers"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Auth) SignUp(c echo.Context) error {
	log.Info("SignUp")
	user := new(handlers.UserModel)
	if err := c.Bind(user); err != nil {
		log.Errorf("echo.Context Error SignUp %s", err)
		return err
	}
	responseRegistration, err := a.client.SignUp(context.Background(), &protocol.SignUpRequest{
		Login:    user.Login,
		Password: user.Password,
	})
	if err != nil {
		log.Errorf("GRPC Error SignUp %s", err)
		return err
	}
	refreshToken := responseRegistration.GetRefreshToken()
	accessToken := responseRegistration.GetToken()
	c.Request().Header.Set("Authorization", accessToken)
	c.Request().Header.Set("RefreshToken", refreshToken)
	return c.JSON(http.StatusOK, &handlers.TokenModel{AccessToken: accessToken, RefreshToken: refreshToken})
}
