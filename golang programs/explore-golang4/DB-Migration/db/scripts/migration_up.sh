echo "Migration Up Starting"

echo "how much up you want to go ?"
read n

migrate -path db/migration/ -database "postgresql://bacancy:admin@localhost:5432/bacancy?sslmode=disable" -verbose up $n

echo "Migration Up Ended"