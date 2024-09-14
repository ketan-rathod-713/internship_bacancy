package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	sqlAdapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	_ "github.com/lib/pq"
)

const (
	ADMIN = "admin"
	OWNER = "owner"
)

type Permission struct {
	Object string `json:"object"`
	Action string `json:"action"`
}

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

	a, err := NewAuthorizer(db)

	// a.AddSuperAdmin("ketan")

	// a.AddServerOwner("tridip", "server1")
	// a.AddServerOwner("manav", "server2")

	// Add new role in server 1 with some permissions
	// permissions := []Permission{
	// 	{
	// 		Object: "consol",
	// 		Action: "read",
	// 	},
	// 	{
	// 		Object: "consol",
	// 		Action: "start",
	// 	},
	// 	{
	// 		Object: "consol",
	// 		Action: "update",
	// 	},
	// 	{
	// 		Object: "consol",
	// 		Action: "create",
	// 	},
	// 	{
	// 		Object: "file",
	// 		Action: "create",
	// 	},
	// 	{
	// 		Object: "file",
	// 		Action: "read",
	// 	},
	// 	{
	// 		Object: "file",
	// 		Action: "update",
	// 	},
	// 	{
	// 		Object: "file",
	// 		Action: "delete",
	// 	},
	// 	{
	// 		Object: "file",
	// 		Action: "archive",
	// 	},
	// 	{
	// 		Object: "file",
	// 		Action: "sftp",
	// 	},
	// 	{
	// 		Object: "backup",
	// 		Action: "create",
	// 	},
	// 	{
	// 		Object: "backup",
	// 		Action: "read",
	// 	},
	// 	{
	// 		Object: "backup",
	// 		Action: "delete",
	// 	},
	// 	{
	// 		Object: "backup",
	// 		Action: "restore",
	// 	},
	// 	{
	// 		Object: "backup",
	// 		Action: "download",
	// 	},
	// 	{
	// 		Object: "startup",
	// 		Action: "read",
	// 	},
	// 	{
	// 		Object: "startup",
	// 		Action: "delete",
	// 	},
	// 	{
	// 		Object: "database",
	// 		Action: "create",
	// 	},
	// 	{
	// 		Object: "database",
	// 		Action: "delete",
	// 	},
	// 	{
	// 		Object: "activity",
	// 		Action: "read",
	// 	},
	// }

	// a.AddNewRole("server1", "player", permissions)

	// a.AddNewRole("server1", "player", []Permission{Permission{Object: "*", Action: "*"}})

	ok, err := a.AddUserToRoleInServerDomain("hiten", "server1", "player")
	log.Println(ok, err)

	// working
	p, err := a.Enforcer.GetDomainsForUser("tridip")
	fmt.Println(p, err)

	// p, err = a.Enforcer.GetRolesForUser("tridip")
	// fmt.Println(p, err)

	// working
	permissions := a.Enforcer.GetPermissionsForUserInDomain("tridip", "server1")
	fmt.Println(permissions)

	// check if domain exists or not // it works great
	ok, err = a.Enforcer.HasPolicy(OWNER, "server1", "*", "*")
	fmt.Println(ok, err)

	// check if role exists or not // not working
	roles, err := a.Enforcer.GetAllRoles()
	fmt.Println(roles, err)

	// working correctly but ignoring *
	domains, err := a.Enforcer.GetAllDomains()
	fmt.Println("all domains or servers", domains, err)

	// not working
	actions, err := a.Enforcer.GetAllActions()
	fmt.Println("all actions", actions, err)

	objects, err := a.Enforcer.GetAllObjects()
	fmt.Println("all objects", objects, err)

	// get all roles by domain // working // i can iterate this and check if role exists in that domain or not
	// once we add any user to it the only it works
	rolesbydomain, err := a.Enforcer.GetAllRolesByDomain("server1")
	fmt.Println("all roles in server1", rolesbydomain, err)

	log.Println("display information about server1")

	// all users // it will also return the roles // so we have to remove roles from it ha ha.
	users, err := a.Enforcer.GetAllUsersByDomain("server1")
	fmt.Println("all users in server1", users, err)

	ok, err = a.Enforcer.AddNamedGroupingPolicy("g2", "super_admin")
	fmt.Println(ok, err)

	// enforce
	ok, err = a.Enforcer.Enforce("super_admin", "server1", "read", "consol")
	log.Println(ok, err)
}

type authorizer struct {
	Enforcer *casbin.Enforcer
}

func NewAuthorizer(db *sql.DB) (*authorizer, error) {
	sqladapter, err := sqlAdapter.NewAdapter(db, "postgres", "game_server_authorization")
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewEnforcer("model.conf", sqladapter)
	if err != nil {
		return nil, err
	}

	ok, err := e.AddPolicies([][]string{
		{"admin", "*", "*", "*"},
	})
	if err != nil {
		return nil, err
	}

	log.Println("policies added:", ok)

	return &authorizer{
		Enforcer: e,
	}, nil
}

func (a *authorizer) CheckPermission(subject, resource, action string) (bool, error) {
	return a.Enforcer.Enforce(subject, resource, action)
}

func (a *authorizer) AddSuperAdmin(userId string) (bool, error) {
	return a.Enforcer.AddRoleForUser(userId, ADMIN, "*")
}

func (a *authorizer) AddServerOwner(userId, serverId string) (bool, error) {
	ok, err := a.Enforcer.AddPolicy(OWNER, serverId, "*", "*")
	if err != nil {
		return ok, err
	}

	return a.Enforcer.AddRoleForUser(userId, OWNER, serverId)
}

func (a *authorizer) AddNewRole(serverId string, roleName string, permissions []Permission) (bool, error) {
	// check if any he is owner of any server

	// check if domain exists or not.
	// check if domain exists or not // it works great
	ok, err := a.Enforcer.HasPolicy(OWNER, "server1", "*", "*")
	if err != nil {
		return false, err
	}

	if !ok {
		return ok, errors.New("domain does not exist")
	}

	// for every permission add policy to the role
	for _, p := range permissions {
		ok, err := a.Enforcer.AddPolicy(roleName, serverId, p.Object, p.Action)
		if err != nil {
			return ok, err
		}
	}

	return true, nil
}

func (a *authorizer) AddUserToRoleInServerDomain(userId, serverId, roleName string) (bool, error) {
	// TODO: check if user exists, by seeing users table

	// TODO: do i need to check if domain exists or not.

	// TODO: check if role exists in our server
	rolesbydomain, err := a.Enforcer.GetAllRolesByDomain(serverId)
	fmt.Println("roles by domain", rolesbydomain, err)
	if err != nil {
		return false, err
	}

	for _, r := range rolesbydomain {
		if r == roleName {
			ok, err := a.Enforcer.AddRoleForUser(userId, roleName, serverId)
			if err != nil {
				return ok, err
			}
			return ok, nil
		}
	}

	// fmt.Println("still doing it")
	// a.Enforcer.AddRoleForUser(userId, roleName, serverId)

	return false, errors.New("role does not exist")
}

// Requirements :-
// add server owner
// delete server owner

// add new role with permissions
// edit role permissions // means delete or add new
// delete role ( deleting all permissions and from table data too )

// add user to role
// remove user from role

// if domain deleted, then delete all user roles from that domain.
