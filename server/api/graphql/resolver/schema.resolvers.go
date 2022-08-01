package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mohwa/ci-cd-github-action/internal/handler/graphql/todo"

	"github.com/mohwa/ci-cd-github-action/api/graphql/generated"
	"github.com/mohwa/ci-cd-github-action/api/graphql/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	err := todo.CreateTodo(model.TodoInput{Name: input.Name})

	if err != nil {
		return nil, err
	}

	return &model.Todo{Name: input.Name}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	ret, err := todo.GetTodos()

	if err != nil {
		return nil, err
	}

	result := make([]*model.Todo, len(ret))

	for k, v := range ret {
		result[k] = &model.Todo{
			ID:   v.ID,
			Name: v.Name,
		}
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
