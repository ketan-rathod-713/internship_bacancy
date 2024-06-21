
# Graphql

It is a graphql project demonstrating the use of gql-gen package. Here i am not using mutations for now. Rather pushing dummy data to database and quering it in graphql way.


## Setting Up Project

```
go get github.com/99designs/gqlgen@latest

printf '// +build tools\npackage tools\nimport _ "github.com/99
designs/gqlgen"' | gofmt > tools.go

go run github.com/99designs/gqlgen init
```

- Remove all uneccessary code and start writting your own schema and then once again generate code for the same. and ig there should be versioning scheme such that no changes should be broken and also it should be backward compatible.
- NOTE : do not generate once again as i have modified joblisting struct with bson field tag for ID. If possible maintain version.

### Defining Resolvers

add below code to gqlgen.yml to define resolvers for particular fields and to also specify models for it.

```
models:
  User:
    model: facebook/models.User
    fields:
      following:
        resolver: true
      followers:
        resolver: true
```

## Running Project

```
go run server.go
```

## Example Grapql Mutations & Queries

```
query example1 {
  users {
    id,
    name,
    followers {
      id,
      name,
      age
    },
    following {
      id,
      name,
      age
    }
  }
}
```

