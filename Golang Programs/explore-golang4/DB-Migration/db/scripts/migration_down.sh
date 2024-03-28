echo "Migration Down Starting"

# How much down we want to go

echo "how much down you want to go ?"
read n

migrate -path db/migration/ -database "postgresql://bacancy:admin@localhost:5432/bacancy?sslmode=disable" -verbose down $n

echo "Migration Down Ended"