package profile_service

import (
	"context"
	"github.com/ivansukach/gateway-service/handlers"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *ProfileAccounting) Create(c echo.Context) error {
	log.Info("Create")
	profile := new(handlers.ProfileModel)

	if err := c.Bind(profile); err != nil {
		log.Errorf("echo.Context Error Create %s", err)
		return err
	}
	p := new(protocol.Profile)
	p.Login = profile.Login
	p.Name = profile.Name
	p.Surname = profile.Surname
	p.Gender = profile.Gender
	p.Age = profile.Age
	p.Password = profile.Password
	p.Employed = profile.Employed
	p.HasAnyPets = profile.HasAnyPets
	_, err := a.client.Create(context.Background(),
		&protocol.CreateRequest{Profile: p})
	if err != nil {
		log.Errorf("GRPC Error Create %s", err)
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, "success")
}
