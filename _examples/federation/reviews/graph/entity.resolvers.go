package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57-dev

import (
	"context"

	"github.com/99designs/gqlgen/_examples/federation/reviews/graph/model"
)

// FindManyProductByManufacturerIDAndIDs is the resolver for the findManyProductByManufacturerIDAndIDs field.
func (r *entityResolver) FindManyProductByManufacturerIDAndIDs(ctx context.Context, reps []*model.ProductByManufacturerIDAndIDsInput) ([]*model.Product, error) {
	products := make([]*model.Product, 0, len(reps))
	for idx := range reps {
		rep := reps[len(reps)-idx-1]
		product, err := r.FindProductByManufacturerIDAndID(ctx, rep.ManufacturerID, rep.ID)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{
		ID:   id,
		Host: &model.EmailHost{},
	}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
