package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/configs"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/entity"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/infra/database"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/APIs/internal/infra/webserver/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	log.Println(conf.Driver)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handler.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Patch("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	log.Println("server starting at port 8000")
	http.ListenAndServe(":8000", r)
}
