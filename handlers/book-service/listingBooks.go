package book_service

import (
	"context"
	"fmt"
	"github.com/ivansukach/book-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *BookAccounting) Listing(c echo.Context) error {
	log.Info("Listing")
	resp, err := a.client.Listing(context.Background(), &protocol.EmptyRequest{})
	if err != nil {
		log.Errorf("GRPC Error Listing Books %s", err)
		return err
	}
	for i := range resp.Books {
		fmt.Println(resp.Books[i])
	}
	return c.String(http.StatusOK, "")
}
