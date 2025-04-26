package service

import (
	"context"
	"io"

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

// Client-side streaming
func (c *CategoryService) CreateCategoryStream(stream categorypb.CategoryService_CreateCategoryStreamServer) error {
	categories := &categorypb.ListCategoriesResponse{}

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		categories.Categories = append(categories.Categories, &categorypb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})

	}
}

// Bi-directional streaming
func (c *CategoryService) CreateCategoryBidirectionalStream(stream categorypb.CategoryService_CreateCategoryBidirectionalStreamServer) error {

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.Send(&categorypb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
		if err != nil {
			return err
		}
	}
}
