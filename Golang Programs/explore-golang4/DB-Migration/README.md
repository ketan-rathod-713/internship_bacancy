# Database Migrations

To keep track of the different schema versions as we move in our development journey. and also to move back and forth in our schema design it can be very usefull.

## Migrate CLI

### Create new migrate up and down file

```
 migrate create -ext sql -dir db/migration/ -seq user_phone_column
```

Refer scripts for more information about commands.
