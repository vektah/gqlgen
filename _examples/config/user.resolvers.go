package config

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22-dev

import (
	"context"
)

// Name is the resolver for the name field.
func (r *roleResolver) Name(ctx context.Context, obj *UserRole) (string, error) {
	if obj == nil {
		return "", nil
	}
	return obj.RoleName, nil
}

// Role returns RoleResolver implementation.
func (r *Resolver) Role() RoleResolver { return &roleResolver{r} }

type roleResolver struct{ *Resolver }
