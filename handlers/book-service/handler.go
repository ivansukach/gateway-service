package book_service

import (
	"github.com/ivansukach/book-service/protocol"
)

type BookAccounting struct {
	client protocol.BookServiceClient
}

func NewHandler(client protocol.BookServiceClient) *BookAccounting {
	return &BookAccounting{client: client}
}
