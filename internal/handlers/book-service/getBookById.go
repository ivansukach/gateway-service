package book_service

import (
	"context"
	"fmt"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/gateway-service/internal/handlers"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *BookService) GetById(c echo.Context) error {
	log.Info("GetById")
	book := &handlers.BookModel{}
	if err := c.Bind(book); err != nil {
		log.Errorf("echo.Context Error GetByID %s", err)
		return err
	}
	resp, err := a.client.Get(context.Background(), &protocol.GetRequest{Id: book.Id})
	if err != nil {
		log.Errorf("gRPC Error GetByID %s", err)
		return err
	}
	fmt.Println(resp.Book)
	return c.String(http.StatusOK, "")
}
