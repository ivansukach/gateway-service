package book_service

import (
	"github.com/ivansukach/book-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type BookService struct {
	client protocol.BookServiceClient
	e      *echo.Echo
}

func NewHandlerBS(url string, version string, e *echo.Echo, opts grpc.DialOption) *BookService {
	// create gRPC here
	clientConnBook, err := grpc.Dial(url, opts)
	if err != nil {
		log.Error(err)
	}
	defer clientConnBook.Close()
	client := protocol.NewBookServiceClient(clientConnBook)
	bs := &BookService{client: client}
	v := "v" + version
	e.POST(v+"/book/create", bs.Create)
	e.PUT(v+"/book/update", bs.Update)
	e.DELETE(v+"/book/delete", bs.Delete)
	e.GET(v+"/book/getById", bs.GetById)
	e.GET(v+"/book/listing", bs.Listing)
	return bs
}
