package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"

	"github.com/fabiohsgomes/go-expert-estudos-graphql/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindByCategoryID(obj.ID)
	if err != nil {
		return nil, err
	}
	var coursesModel []*model.Course
	for _, course := range courses {
		coursesModel = append(coursesModel, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}
	return coursesModel, nil
}

// AddCategory is the resolver for the addCategory field.
func (r *mutationResolver) AddCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.Create(input.Name, *input.Description)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

// AddCourse is the resolver for the addCourse field.
func (r *mutationResolver) AddCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	course, err := r.CourseDB.Create(input.Name, *input.Description, input.CategoryID)
	if err != nil {
		return nil, err
	}

	return &model.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesModel []*model.Category
	for _, category := range categories {
		categoriesModel = append(categoriesModel, &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return categoriesModel, nil
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}
	var coursesModel []*model.Course
	for _, course := range courses {
		coursesModel = append(coursesModel, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}
	return coursesModel, nil
}

// Course is the resolver for the course field.
func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: Course - course"))
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
