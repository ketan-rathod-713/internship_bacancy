# Important


## Migration CLI Usecase

```
source .env
migrate create -ext sql -dir postgres/migrations create_meetups;

// Finally to migrate
migrate --path "postgres/migrations" --database "$POSTGRESQL_URL" up

// For using dataloaden run below command inside graph directory.
go run github.com/vektah/dataloaden UserLoader string '[]*meetmeup/models.User';
```

## Dataloader

- Define this middleware so that sql calls to database can be minimized.


## Go Pg Package

- Good postgresql orm library.

## Flow of whole program here

- First of all create a starter project for us using golang gqlgen library.
- Update the gqlgen.yml according to requireements add schema to schema.graphql and update the models field in gqlgen.yml file defining which fields to have resolvers and other data.
- Then simply generate code and then implement those methods.
- Add data loader as a middleware and implement it.