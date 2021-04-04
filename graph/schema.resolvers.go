package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/nasdf/git-graphql/graph/generated"
	"github.com/nasdf/git-graphql/graph/model"
)

func (r *blobResolver) Data(ctx context.Context, obj *model.Blob) (string, error) {
	blob, err := r.repo.BlobObject(plumbing.NewHash(obj.Hash))
	if err != nil {
		return "", err
	}

	br, err := blob.Reader()
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(br)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (r *commitResolver) Tree(ctx context.Context, obj *model.Commit) (*model.Tree, error) {
	return r.Query().Tree(ctx, obj.TreeHash)
}

func (r *commitResolver) Parents(ctx context.Context, obj *model.Commit) ([]*model.Commit, error) {
	var parents []*model.Commit
	for _, h := range obj.ParentHashes {
		commit, err := r.Query().Commit(ctx, h)
		if err != nil {
			return nil, err
		}

		parents = append(parents, commit)
	}

	return parents, nil
}

func (r *queryResolver) Commit(ctx context.Context, hash string) (*model.Commit, error) {
	commit, err := r.repo.CommitObject(plumbing.NewHash(hash))
	if err != nil {
		return nil, err
	}

	return NewCommit(commit), nil
}

func (r *queryResolver) Blob(ctx context.Context, hash string) (*model.Blob, error) {
	blob, err := r.repo.BlobObject(plumbing.NewHash(hash))
	if err != nil {
		return nil, err
	}

	return &model.Blob{
		Hash: blob.Hash.String(),
		Type: blob.Type().String(),
	}, nil
}

func (r *queryResolver) Tree(ctx context.Context, hash string) (*model.Tree, error) {
	tree, err := r.repo.TreeObject(plumbing.NewHash(hash))
	if err != nil {
		return nil, err
	}

	var entries []*model.TreeEntry
	for _, e := range tree.Entries {
		entries = append(entries, &model.TreeEntry{
			Name: e.Name,
			Mode: e.Mode.String(),
			Hash: e.Hash.String(),
		})
	}

	return &model.Tree{
		Hash:    tree.Hash.String(),
		Type:    tree.Type().String(),
		Entries: entries,
	}, nil
}

func (r *queryResolver) Revision(ctx context.Context, name string) (*model.Commit, error) {
	hash, err := r.repo.ResolveRevision(plumbing.Revision(name))
	if err != nil {
		return nil, err
	}

	return r.Query().Commit(ctx, hash.String())
}

func (r *queryResolver) Commits(ctx context.Context) ([]*model.Commit, error) {
	iter, err := r.repo.CommitObjects()
	if err != nil {
		return nil, err
	}

	var commits []*model.Commit
	err = iter.ForEach(func(commit *object.Commit) error {
		commits = append(commits, NewCommit(commit))
		return nil
	})

	return commits, err
}

func (r *treeEntryResolver) Object(ctx context.Context, obj *model.TreeEntry) (model.Object, error) {
	mode, err := filemode.New(obj.Mode)
	if err != nil {
		return nil, err
	}

	switch {
	case mode.IsFile():
		return r.Query().Blob(ctx, obj.Hash)
	default:
		return r.Query().Tree(ctx, obj.Hash)
	}
}

// Blob returns generated.BlobResolver implementation.
func (r *Resolver) Blob() generated.BlobResolver { return &blobResolver{r} }

// Commit returns generated.CommitResolver implementation.
func (r *Resolver) Commit() generated.CommitResolver { return &commitResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TreeEntry returns generated.TreeEntryResolver implementation.
func (r *Resolver) TreeEntry() generated.TreeEntryResolver { return &treeEntryResolver{r} }

type blobResolver struct{ *Resolver }
type commitResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type treeEntryResolver struct{ *Resolver }
