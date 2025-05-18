package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rafaelcamelo31/graduate-go-course/4-module/uow/internal/db"
	"github.com/rafaelcamelo31/graduate-go-course/4-module/uow/internal/repository"
	uow "github.com/rafaelcamelo31/graduate-go-course/4-module/uow/pkg"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	ctx := context.Background()
	uow := uow.NewUow(ctx, dbt)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	input := InputUseCase{
		CategoryName:     "Category 1", // ID->1
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	useCase := NewAddCourseUseCaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
