package service

import (
	"log"
	"net"

	"github.com/temur-shamshidinov/Content_service/genproto/content_service"
	"github.com/temur-shamshidinov/Content_service/storage"

	"google.golang.org/grpc"
)

func Service() {

	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
		return
	}

	server := grpc.NewServer()

	storage := storage.NewStorage()

	contentService := NewContentServices(storage)

	content_service.RegisterContentServiceServer(server,contentService)

	server.Serve(listen)
}