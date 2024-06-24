package main

import (
	"database/sql"
	"log"
	"time"

	sqlAdapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Welcome to casbin")

	// connect to the database first.
	db, err := sql.Open("postgres", "postgresql://bacancy:admin@localhost:5432/bacancy")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 10)

	log.Println("Database Connected")

	sqladapter, err := sqlAdapter.NewAdapter(db, "postgres", "casbin_rule")
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer("model.conf", sqladapter)
	if err != nil {
		panic(err)
	}

	// Load the policy from DB.
	if err = e.LoadPolicy(); err != nil {
		log.Println("LoadPolicy failed, err: ", err)
	}

	// Check the permission.
	has, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		log.Println("Enforce failed, err: ", err)
	}
	if !has {
		log.Println("do not have permission")
	}

	e.AddPolicy("userId", "Resource", "read")

}
