package book_service

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/gateway-service/internal/handlers"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *BookService) Create(c echo.Context) error {
	log.Info("Create")
	book := &handlers.BookModel{}

	if err := c.Bind(book); err != nil {
		log.Errorf("echo.Context Error Create %s", err)
		return err
	}
	b := &protocol.Book{Id: book.Id, Title: book.Title,
		Author: book.Author, Genre: book.Genre,
		Edition: book.Edition, NumberOfPages: book.NumberOfPages,
		Year: book.Year, Amount: book.Amount,
		IsPopular: book.IsPopular, InStock: book.InStock}

	_, err := a.client.Add(context.Background(),
		&protocol.AddRequest{Book: b})
	if err != nil {
		log.Errorf("gRPC Error Add Book %s", err)
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, "success")
}
