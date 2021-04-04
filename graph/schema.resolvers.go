package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/nasdf/git-graphql/graph/generated"
	"github.com/nasdf/git-graphql/graph/model"
)

func (r *commitResolver) Tree(ctx context.Context, obj *model.Commit) (*model.Tree, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commitResolver) Parents(ctx context.Context, obj *model.Commit) ([]*model.Commit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Branches(ctx context.Context) ([]*model.Reference, error) {
	var branches []*model.Reference

	iter, err := r.repo.Branches()
	if err != nil {
		return nil, err
	}

	err = iter.ForEach(func(ref *plumbing.Reference) error {
		branches = append(branches, NewReference(ref))
		return nil
	})

	return branches, err
}

func (r *tagResolver) Target(ctx context.Context, obj *model.Tag) (model.Object, error) {
	panic(fmt.Errorf("not implemented"))
}

// Commit returns generated.CommitResolver implementation.
func (r *Resolver) Commit() generated.CommitResolver { return &commitResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

type commitResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tagResolver struct{ *Resolver }
