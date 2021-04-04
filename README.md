# git-graphql

GraphQL server for your git repositories.  Powered by [gqlgen](https://github.com/99designs/gqlgen) and [go-git](https://github.com/go-git/go-git).

## Usage

Go 1.16 or higher required.

```bash
go install github.com/nasdf/git-graphql
# run this from your repo root
git-graphql
```

Query repo head tree entries.

```graphql
query head {
  revision(name: "HEAD") {
    tree {
      entries {
        name
        mode
        object {
          hash
          type
        }
      }
    }
  }
}
```

Query repo commit history.

```graphql
query history {
  commits {
    hash
    message
    committer {
      name
      email
      when
    }
  }
}
```

## License

GPLv3