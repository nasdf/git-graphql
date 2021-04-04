// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Object interface {
	IsObject()
}

type Blob struct {
	Hash string `json:"hash"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (Blob) IsObject() {}

type Commit struct {
	Hash         string     `json:"hash"`
	Type         string     `json:"type"`
	Author       *Signature `json:"author"`
	Committer    *Signature `json:"committer"`
	Signature    string     `json:"signature"`
	Message      string     `json:"message"`
	Tree         *Tree      `json:"tree"`
	TreeHash     string     `json:"treeHash"`
	Parents      []*Commit  `json:"parents"`
	ParentHashes []string   `json:"parentHashes"`
}

func (Commit) IsObject() {}

type Signature struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	When  string `json:"when"`
}

type Tree struct {
	Hash    string       `json:"hash"`
	Type    string       `json:"type"`
	Entries []*TreeEntry `json:"entries"`
}

func (Tree) IsObject() {}

type TreeEntry struct {
	Hash   string `json:"hash"`
	Name   string `json:"name"`
	Mode   string `json:"mode"`
	Object Object `json:"object"`
}
