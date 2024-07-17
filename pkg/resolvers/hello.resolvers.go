package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/mahmoud-elgammal/raven/graph"
	"github.com/mahmoud-elgammal/raven/graph/model"
)

// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context, name *string) (*model.Hello, error) {
	if len(*name) == 0 {
		return nil, fmt.Errorf("missing name")
	}

	value := fmt.Sprintf("Hello, %s!", *name)

	return &model.Hello{Value: &value}, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }
