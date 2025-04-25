package service

import (
	"context"

	"github.com/rafaelcamelo31/graduate-go-course/3-module/grpc/internal/database"
	categorypb "github.com/rafaelcamelo31/graduate-go-course/3-module/grpc/internal/gen/pbs"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryService struct {
	categorypb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, req *categorypb.CreateCategoryRequest) (*categorypb.Category, error) {
	cat, err := c.CategoryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	resp := &categorypb.Category{
		Id:          cat.ID,
		Name:        cat.Name,
		Description: cat.Description,
	}

	return resp, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *emptypb.Empty) (*categorypb.ListCategoriesResponse, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	list := []*categorypb.Category{}
	for _, v := range categories {
		list = append(list, &categorypb.Category{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		})
	}

	resp := &categorypb.ListCategoriesResponse{
		Categories: list,
	}

	return resp, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, req *categorypb.GetCategoryRequest) (*categorypb.Category, error) {
	category, err := c.CategoryDB.FindByID(req.Id)
	if err != nil {
		return nil, err
	}

	resp := &categorypb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return resp, nil
}
