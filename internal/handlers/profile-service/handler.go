package profile_service

import (
	protocolPS "github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ProfileService struct {
	client protocolPS.ProfileServiceClient
}

func NewHandlerPS(url string, version string, e *echo.Echo, opts grpc.DialOption) *ProfileService {
	clientConnProfile, err := grpc.Dial(url, opts)
	if err != nil {
		log.Error(err)
	}

	defer clientConnProfile.Close()

	clientPS := protocolPS.NewProfileServiceClient(clientConnProfile)
	ps := &ProfileService{client: clientPS}
	v := "v" + version
	e.POST(v+"profile/create", ps.Create)
	e.PUT(v+"profile/update", ps.Update)
	e.DELETE(v+"profile/delete", ps.Delete)
	e.GET(v+"profile/getByLogin", ps.GetByLogin)
	e.GET(v+"profile/listing", ps.Listing)
	return ps
}
