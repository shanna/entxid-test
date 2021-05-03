package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	xid1 "github.com/rs/xid"
	"github.com/shanna/entxid-test/ent"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input UserInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	return client.User.
		Create().
		SetName(input.Name).
		Save(ctx)
}

func (r *mutationResolver) ClearUsers(ctx context.Context) (int, error) {
	client := ent.FromContext(ctx)
	return client.User.
		Delete().
		Exec(ctx)
}

func (r *queryResolver) Node(ctx context.Context, id xid1.ID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []xid1.ID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
		)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
