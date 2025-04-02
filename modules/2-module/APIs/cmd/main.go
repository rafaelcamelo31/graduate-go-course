package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
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

	userDB := database.NewUser(db)
	userHandler := handler.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", conf.TokenAuth))
	r.Use(middleware.WithValue("expiresIn", conf.JWTExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(conf.TokenAuth))
		r.Use(jwtauth.Authenticator(conf.TokenAuth))
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Patch("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/get_token", userHandler.GetJWT)

	log.Println("server starting at port 8000")
	http.ListenAndServe(":8000", r)
}

// Example of custom middleware
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
