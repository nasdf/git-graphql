package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/nasdf/git-graphql/graph/model"
)

// Resolver returns objects from a git repo.
type Resolver struct {
	repo *git.Repository
}

// NewResolver creates a resolver from a git repo in the current directory.
func NewResolver() (*Resolver, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(cwd)
	if err != nil {
		return nil, err
	}

	return &Resolver{repo}, nil
}

// NewCommit returns a commit model from a git commit.
func NewCommit(commit *object.Commit) *model.Commit {
	var parents []string
	for _, p := range commit.ParentHashes {
		parents = append(parents, p.String())
	}

	author := &model.Signature{
		Name:  commit.Author.Name,
		Email: commit.Author.Email,
		When:  commit.Author.When.String(),
	}

	committer := &model.Signature{
		Name:  commit.Committer.Name,
		Email: commit.Committer.Email,
		When:  commit.Author.When.String(),
	}

	return &model.Commit{
		Hash:         commit.Hash.String(),
		Type:         commit.Type().String(),
		Author:       author,
		Committer:    committer,
		Signature:    commit.PGPSignature,
		Message:      commit.Message,
		TreeHash:     commit.TreeHash.String(),
		ParentHashes: parents,
	}
}
