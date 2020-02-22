package profile_service

import (
	"context"
	"fmt"
	"github.com/ivansukach/gateway-service/handlers"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *ProfileAccounting) GetByLogin(c echo.Context) error {
	log.Info("(a *ProfileAccounting) GetByLogin")
	profile := new(handlers.ProfileModel)
	if err := c.Bind(profile); err != nil {
		log.Errorf("echo.Context Error GetByID Profile %s", err)
		return err
	}
	resp, err := a.client.GetByLogin(context.Background(), &protocol.GetByLoginRequest{Login: profile.Login})
	if err != nil {
		log.Errorf("GRPC Error GetByLogin %s", err)
		return err
	}
	fmt.Println(resp.Profile)
	return c.String(http.StatusOK, "")
}
