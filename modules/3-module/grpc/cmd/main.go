package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rafaelcamelo31/graduate-go-course/3-module/grpc/internal/database"
	categorypb "github.com/rafaelcamelo31/graduate-go-course/3-module/grpc/internal/gen/pbs"
	"github.com/rafaelcamelo31/graduate-go-course/3-module/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	categorypb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	log.Println("Starting gRPC server on port :50051...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
