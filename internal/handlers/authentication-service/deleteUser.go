package authentication_service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/ivansukach/gateway-service/internal/handlers"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Auth) DeleteUser(c echo.Context) error {
	log.Info("DeleteUser")
	user := &handlers.UserModel{}
	if err := c.Bind(user); err != nil {
		log.Errorf("echo.Context Error DeleteUser %s", err)
		return err
	}
	claims := *c.Get("claims").(*jwt.MapClaims)

	if claims["admin"] == "true" {
		_, err := a.client.Delete(context.Background(), &protocol.DeleteRequest{Login: user.Login})
		if err != nil {
			log.Errorf("gRPC Error DeleteClaims %s", err)
			return echo.ErrBadRequest
		}
	} else {
		log.Error("ErrUnauthorized")
		return echo.ErrUnauthorized
	}
	return c.String(http.StatusOK, "")
}
