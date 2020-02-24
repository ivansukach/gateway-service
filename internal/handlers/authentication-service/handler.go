package authentication_service

import (
	"github.com/ivansukach/gateway-service/internal/middlewares"
	"github.com/labstack/echo"
	protocolAuth "github.com/leshachaplin/grpc-server/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Auth struct {
	client protocolAuth.AuthServiceClient
}

func NewHandlerAuth(url string, version string, e *echo.Echo, opts grpc.DialOption) *Auth {
	clientConnAuthService, err := grpc.Dial(url, opts)
	if err != nil {
		log.Error(err)
	}
	defer clientConnAuthService.Close()
	clientAS := protocolAuth.NewAuthServiceClient(clientConnAuthService)
	as := &Auth{client: clientAS}
	jwt := middlewares.NewJWT(clientAS)
	v := "v" + version
	e.POST(v+"auth/signIn", as.SignIn)
	e.POST(v+"auth/signUp", as.SignUp)
	e.DELETE(v+"auth/deleteUser", as.DeleteUser, jwt.Middleware)
	e.DELETE(v+"auth/deleteClaims", as.DeleteClaims, jwt.Middleware)
	e.POST(v+"auth/addClaims", as.AddClaims, jwt.Middleware)
	return as
}
