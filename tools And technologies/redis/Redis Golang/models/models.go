package models

import "github.com/gomodule/redigo/redis"

type User struct {
	Id       string
	Name     string
	Password string
}

type RedisConnection struct {
	Conn *redis.Conn
	Name string
}
