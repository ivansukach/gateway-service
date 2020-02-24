package book_service

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/gateway-service/internal/handlers"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *BookService) Delete(c echo.Context) error {
	log.Info("DeleteBook")
	book := &handlers.BookModel{}
	if err := c.Bind(book); err != nil {
		log.Errorf("echo.Context Error Delete %s", err)
		return err
	}
	_, err := a.client.Delete(context.Background(), &protocol.DeleteRequest{Id: book.Id})
	if err != nil {
		log.Errorf("gRPC Error DeleteClaims %s", err)
		return echo.ErrBadRequest
	}

	return c.String(http.StatusOK, "success")
}
