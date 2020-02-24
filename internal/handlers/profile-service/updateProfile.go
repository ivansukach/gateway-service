package profile_service

import (
	"context"
	"github.com/ivansukach/gateway-service/internal/handlers"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *ProfileService) Update(c echo.Context) error {
	log.Info("Create")
	profile := &handlers.ProfileModel{}

	if err := c.Bind(profile); err != nil {
		log.Errorf("echo.Context Error Update Profile %s", err)
		return err
	}
	p := &protocol.Profile{
		Login:      profile.Login,
		Name:       profile.Name,
		Surname:    profile.Surname,
		Gender:     profile.Gender,
		Age:        profile.Age,
		Password:   profile.Password,
		Employed:   profile.Employed,
		HasAnyPets: profile.HasAnyPets}

	_, err := a.client.Update(context.Background(),
		&protocol.UpdateRequest{Profile: p})
	if err != nil {
		log.Errorf("gRPC Error Update Profile %s", err)
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, "success")
}
