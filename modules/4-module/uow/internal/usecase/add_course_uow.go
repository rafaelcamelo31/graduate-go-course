package usecase

import (
	"context"

	"github.com/rafaelcamelo31/graduate-go-course/4-module/uow/internal/entity"
	"github.com/rafaelcamelo31/graduate-go-course/4-module/uow/internal/repository"
	uow "github.com/rafaelcamelo31/graduate-go-course/4-module/uow/pkg"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		// Every code in this block is executed with begin or commit
		category := entity.Category{
			Name: input.CategoryName,
		}
		repoCategory := a.getCategoryRepository(ctx)
		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
		repoCourse := a.getCourseRepository(ctx)
		err = repoCourse.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}
