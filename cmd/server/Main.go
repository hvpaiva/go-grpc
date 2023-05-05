package main

import (
	"database/sql"
	"github.com/hvpaiva/go-grpc/internal/database"
	"github.com/hvpaiva/go-grpc/internal/pb"
	"github.com/hvpaiva/go-grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer executeOrPanic(db.Close)

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categoryService)
	reflection.Register(server)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

func executeOrPanic(something func() error) {
	err := something()
	if err != nil {
		panic(err)
	}
}
