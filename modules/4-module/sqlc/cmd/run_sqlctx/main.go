package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rafaelcamelo31/graduate-go-course/4-module/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTX(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, courses := range courses {
		fmt.Println(courses.ID, courses.Name, courses.Description.String, courses.Price, courses.CategoryID)
	}

	// courseDB := NewCourseDB(dbConn)

	// categoryParams := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// }

	// courseParam := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// 	Price:       100.62,
	// }

	// err = courseDB.CreateCourseAndCategory(ctx, categoryParams, courseParam)
	// if err != nil {
	// 	panic(err)
	// }

}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTX(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
