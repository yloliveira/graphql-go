package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"graphql-go/graph/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.Create(input.Title)

	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:    category.ID,
		Title: category.Title,
	}, nil
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	product, err := r.ProductDB.Create(input.Title, input.Price, input.CategoryID)

	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:    product.ID,
		Title: product.Title,
		Price: product.Price,
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
		categoriesModel = append(categoriesModel, &model.Category{ID: category.ID, Title: category.Title})
	}

	return categoriesModel, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
