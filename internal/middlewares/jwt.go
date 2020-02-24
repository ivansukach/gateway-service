package middlewares

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ivansukach/gateway-service/internal/config"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	"time"
)

type JWT struct {
	client protocol.AuthServiceClient
}

func NewJWT(client protocol.AuthServiceClient) *JWT {
	return &JWT{client: client}
}
func DecryptToken(tokS string, secretkey []byte) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokS, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretkey, nil
	})
	return token, err
}

func (j *JWT) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	cfg := config.Load()
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("Authorization")
		refreshToken := c.Request().Header.Get("RefreshToken")
		decryptedToken, err := DecryptToken(accessToken, []byte(cfg.SecretKeyAuth))
		if err != nil {
			c.String(400, "Unauthorized: inappropriate access token format")
			return nil
		}
		decryptedRToken, err := DecryptToken(refreshToken, []byte(cfg.SecretKeyRefresh))
		if err != nil {
			c.String(400, "Unauthorized: inappropriate refresh token format")
			return nil
		}
		claims := decryptedToken.Claims.(jwt.MapClaims)
		refClaims := decryptedRToken.Claims.(jwt.MapClaims)
		tokenExpirationDate := (claims["exp"]).(float64)
		rTokenExpirationDate := (refClaims["exp"]).(float64)
		fmt.Println("Time now: ", time.Now().Unix())
		fmt.Println("Expiration time for access token: ", int64(tokenExpirationDate))
		fmt.Println("Expiration time for refresh token: ", int64(rTokenExpirationDate))
		if int64(tokenExpirationDate) <= time.Now().Unix() || int64(rTokenExpirationDate) <= time.Now().Unix() {
			fmt.Println("Time is over. You will get new tokens")
			responseRefresh, err := j.client.RefreshToken(context.Background(), &protocol.RefreshTokenRequest{TokenRefresh: refreshToken, Token: accessToken})
			if err != nil {
				c.String(400, "Something wrong during refreshing your tokens")
				return nil
			}
			refreshToken = responseRefresh.GetRefreshToken()
			accessToken = responseRefresh.GetToken()
			c.Request().Header.Set("Authorization", accessToken)
			c.Request().Header.Set("RefreshToken", refreshToken)
		}
		c.Set("claims", &claims)
		return next(c)
	}
}
