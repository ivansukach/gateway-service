package profile_service

import (
	"context"
	"fmt"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *ProfileService) Listing(c echo.Context) error {
	log.Info("Listing")
	resp, err := a.client.Listing(context.Background(), &protocol.ListingRequest{})
	if err != nil {
		log.Errorf("gRPC Error Listing Profiles %s", err)
		return err
	}
	for i := range resp.Profiles {
		fmt.Println(resp.Profiles[i])
	}
	return c.String(http.StatusOK, "")
}
