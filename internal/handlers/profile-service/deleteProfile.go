package profile_service

import (
	"context"
	"github.com/ivansukach/gateway-service/internal/handlers"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *ProfileService) Delete(c echo.Context) error {
	log.Info("DeleteProfile")
	profile := &handlers.ProfileModel{}
	if err := c.Bind(profile); err != nil {
		log.Errorf("echo.Context Error Delete User %s", err)
		return err
	}
	_, err := a.client.Delete(context.Background(), &protocol.DeleteRequest{Login: profile.Login})
	if err != nil {
		log.Errorf("пRPC Error Delete User %s", err)
		return echo.ErrBadRequest
	}

	return c.String(http.StatusOK, "success")
}
