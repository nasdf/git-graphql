package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/nasdf/git-graphql/graph/model"
)

type Resolver struct {
	repo *git.Repository
}

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

func NewReference(r *plumbing.Reference) *model.Reference {
	ref := model.Reference{
		Name: r.Name().String(),
		Type: r.Type().String(),
	}

	switch r.Type() {
	case plumbing.HashReference:
		hash := r.Hash().String()
		ref.Hash = &hash
	case plumbing.SymbolicReference:
		target := r.Target().String()
		ref.Target = &target
	}

	return &ref
}