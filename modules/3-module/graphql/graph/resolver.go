package graph

import (
	"github.com/rafaelcamelo31/graduate-go-course/modules/3-module/graphql/internal/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
