# Testing In Go


## Important Commands

### Get Coverage Profile Of Test In c.out
go test -coverprofile=c.out

## Display cover profile data
go tool cover -func=c.out

## Display cover profile in web page
go tool cover -html=c.out

## Run all benchmarks in go
go test -bench=.

## For specific benchmark
go test -bench=Add

