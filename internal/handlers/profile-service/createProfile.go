package profile_service

import (
	"context"
	"github.com/ivansukach/gateway-service/internal/handlers"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *ProfileService) Create(c echo.Context) error {
	log.Info("Create")
	profile := &handlers.ProfileModel{}

	if err := c.Bind(profile); err != nil {
		log.Errorf("echo.Context Error Create %s", err)
		return err
	}
	if a.ValidationFields(profile) {
		p := &protocol.Profile{
			Login:      profile.Login,
			Name:       profile.Name,
			Surname:    profile.Surname,
			Gender:     profile.Gender,
			Age:        profile.Age,
			Password:   profile.Password,
			Employed:   profile.Employed,
			HasAnyPets: profile.HasAnyPets}

		_, err := a.client.Create(context.Background(),
			&protocol.CreateRequest{Profile: p})
		if err != nil {
			log.Errorf("gRPC Error Create %s", err)
			return echo.ErrBadRequest
		}
		return c.JSON(http.StatusOK, "success")
	}
	return c.JSON(http.StatusOK, "failure")
}
func (a *ProfileService) ValidationFields(p *handlers.ProfileModel) bool {
	if len(p.Login) < 8 || len(p.Login) > 30 {
		return false
	}
	if len(p.Name) < 2 || len(p.Name) > 15 {
		return false
	}
	if len(p.Surname) < 2 || len(p.Surname) > 30 {
		return false
	}
	if p.Age > 120 || p.Age < 0 {
		return false
	}
	if len(p.Password) < 8 || len(p.Password) > 30 {
		return false
	}
	profile, err := a.client.GetByLogin(context.Background(),
		&protocol.GetByLoginRequest{Login: p.Login})
	if profile != nil || err == nil {
		return false
	}
	return true
}
