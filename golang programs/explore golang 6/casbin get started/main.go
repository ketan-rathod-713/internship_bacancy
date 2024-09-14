package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	sqlAdapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Welcome to casbin")

	// connect to the database first.
	db, err := sql.Open("postgres", "postgresql://root:rootpass@localhost:5432/iceline-hosting?sslmode=disable")
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

	sqladapter, err := sqlAdapter.NewAdapter(db, "postgres", "game_server_authorization")
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer("model.conf", sqladapter)
	if err != nil {
		panic(err)
	}

	fmt.Println("Now i can proceed for other changes", e)

	// Add some policies inside the database.

	// TODO: can i change this policy to like :- owner gameservers * // means can do anything so that for each user i don't have to add
	isAdded, err := e.AddPolicies([][]string{
		// {"super_admin", "owner", "*"}, // i can not add this policy as it should have 4 fields // subject, domain, object, action
		{"owner", "gameserver1", "startup_information", "read"},
		{"owner", "gameserver1", "startup_information", "write"},
		{"user", "gameserver1", "startup_information", "read"},
	})

	if err != nil {
		log.Println("Error adding policies", err)
	}

	fmt.Println("isadded: ", isAdded)

	fmt.Println("checking permissions for alice on game server 1")
	ok, err := e.Enforce("alice", "gameserver1", "startup_information", "write")
	if err != nil {
		log.Println("Error enforcing policy", err)
	}

	fmt.Println("allowed for alice: ", ok)

	ok, err = e.Enforce("ketan", "gameserver1", "startup_information", "write")
	if err != nil {
		log.Println("Error enforcing policy", err)
	}

	fmt.Println("allowed for ketan: ", ok)

	// aman is not in the database, so this should return false // ok done
	// matchers should be matched well before.
	ok, err = e.Enforce("aman", "gameserver1", "startup_information", "write")
	if err != nil {
		log.Println("Error enforcing policy", err)
	}

	fmt.Println("allowed for aman: ", ok)

	// let's add super_admin to all of them
	ok, err = e.AddGroupingPolicies([][]string{
		{"super_admin", "owner", "*"},
	})
	if err != nil {
		log.Println("Error enforcing policy", err)
	}
	fmt.Println("is added: ", ok)

	ok, err = e.Enforce("super_admin", "gameserver1", "startup_information", "write")
	if err != nil {
		log.Println("Error enforcing policy", err)
	}
	fmt.Println("allowed for super_admin: ", ok)

}
