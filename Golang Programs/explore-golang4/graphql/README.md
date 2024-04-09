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

Getting issue with the dataloaden.


## Go Pg Package

- Connection


- Basics

## Flow of whole program here

