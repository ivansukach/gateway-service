package book_service

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/gateway-service/handlers"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *BookAccounting) Create(c echo.Context) error {
	log.Info("Create")
	book := new(handlers.BookModel)

	if err := c.Bind(book); err != nil {
		log.Errorf("echo.Context Error Create %s", err)
		return err
	}
	b := new(protocol.Book)
	b.Id = book.Id
	b.Title = book.Title
	b.Author = book.Author
	b.Genre = book.Genre
	b.Edition = book.Edition
	b.NumberOfPages = book.NumberOfPages
	b.Year = book.Year
	b.Amount = book.Amount
	b.IsPopular = book.IsPopular
	b.InStock = book.InStock
	_, err := a.client.Add(context.Background(),
		&protocol.AddRequest{Book: b})
	if err != nil {
		log.Errorf("GRPC Error Add Book %s", err)
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, "success")
}
