package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57-dev

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/_examples/federation/reviews/graph/model"
)

// ManufacturerID is the resolver for the manufacturerID field.
func (r *productResolver) ManufacturerID(ctx context.Context, obj *model.Product, federationRequires map[string]interface{}) (*string, error) {
	manufacturer, ok := federationRequires["manufacturer"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("manufacturer not provided or not an object")
	}

	manufacturerID, ok := manufacturer["id"].(string)
	if !ok {
		return nil, fmt.Errorf("manufacturer.id not provided or not a string")
	}

	return &manufacturerID, nil
}

// Username is the resolver for the username field.
func (r *userResolver) Username(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented: Username - username"))
}

// Reviews is the resolver for the reviews field.
func (r *userResolver) Reviews(ctx context.Context, obj *model.User, federationRequires map[string]interface{}) ([]*model.Review, error) {
	var productReviews []*model.Review
	for _, review := range reviews {
		if review.Author.ID == obj.ID {
			host, ok := federationRequires["host"].(map[string]any)
			if !ok {
				return nil, fmt.Errorf("host not provided or not an object")
			}

			hostID, ok := host["id"].(string)
			if !ok {
				return nil, fmt.Errorf("host.id not provided or not a string")
			}

			email, ok := federationRequires["email"].(string)
			if !ok {
				return nil, fmt.Errorf("email not provided or not a string")
			}

			productReviews = append(productReviews, &model.Review{
				Body:        review.Body,
				Author:      review.Author,
				Product:     review.Product,
				HostIDEmail: hostID + ":" + email,
			})
		}
	}
	return productReviews, nil
}

// Product returns ProductResolver implementation.
func (r *Resolver) Product() ProductResolver { return &productResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type productResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
