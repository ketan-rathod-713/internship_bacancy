hey:
	echo "hey engineers!"

up: 
	sh db/scripts/migration_up.sh
	
down:
	sh db/scripts/migration_down.sh

create:
	migrate create -ext sql -dir db/migration -seq playground_schema

current:
	migrate -path db/migration/ -database "postgresql://bacancy:admin@localhost:5432/bacancy?sslmode=disable" version

	
