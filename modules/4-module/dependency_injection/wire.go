//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/rafaelcamelo31/graduate-go-course/4-module/dependency_injection/product"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
