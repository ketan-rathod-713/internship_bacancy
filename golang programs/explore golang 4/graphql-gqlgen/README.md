
# Graphql


## Setting Up Project

```
go get github.com/99designs/gqlgen@latest

printf '// +build tools\npackage tools\nimport _ "github.com/99
designs/gqlgen"' | gofmt > tools.go

go run github.com/99designs/gqlgen init
```

- Remove all uneccessary code and start writting your own schema and then once again generate code for the same. and ig there should be versioning scheme such that no changes should be broken and also it should be backward compatible.
- NOTE : do not generate once again as i have modified joblisting struct with bson field tag for ID. If possible maintain version.

## Running Project

```
go run server.go
```

## Example Grapql Mutations & Queries

```
mutation exampleMutation($data:CreateJobListingInput) {
  createJobListing(input: $data){
    _id,
    title,
    description,
    company,
    url,
    __typename
  }
}

mutation exampleMutationUpdate($id: ID!,$data:UpdateJobListingInput) {
  updateJobListing(id:$id, input: $data){
    _id,
    title,
    description,
    company,
    url
  }
}

mutation exampleDelete($id: ID!) {
  deleteJobListing(id: $id){
    deletedJobId
  }
}

query exampleJobs {
  jobs {
    _id,
    title,
    description,
    company,
    url,
    __typename
  }
}

query exampleJobs($id:ID!) {
  job(id: $id) {
    _id,
    title,
    description
  }
}
```
