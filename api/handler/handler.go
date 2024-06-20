package handler

import (
	pb "api-geteway/genproto/library_service"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Author   pb.AuthorServiceClient
	Book     pb.BookServiceClient
	Genre    pb.GenreServiceClient
	Borrower pb.BorrowerServiceClient
}

func NewHandler(conn *grpc.ClientConn) *HTTPHandler {
	return &HTTPHandler{
		Author:   pb.NewAuthorServiceClient(conn),
		Book:     pb.NewBookServiceClient(conn),
		Genre:    pb.NewGenreServiceClient(conn),
		Borrower: pb.NewBorrowerServiceClient(conn),
	}
}
