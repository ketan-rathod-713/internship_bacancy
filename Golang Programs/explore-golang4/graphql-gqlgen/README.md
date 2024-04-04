
# Initial Project Structure

go get github.com/99designs/gqlgen@latest

printf '// +build tools\npackage tools\nimport _ "github.com/99
designs/gqlgen"' | gofmt > tools.go

gqlgen ( graphql generation) with the help of this we haven't need to do all the hard work by ourself.

Now run init command

go run github.com/99designs/gqlgen init

Lots of things happen when we run this command. 

generated.go // also remove all 

server.go file is created
complete graph folder is created.

models_gen.go // we will change it.

resolvers.go
schema.graphqls // we wil change our schema from here
schema.resolvers.go

Now start from writting schema of your own choice.

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

NOTE : do not generate once again as i have modified joblisting struct with bson field tag for ID.