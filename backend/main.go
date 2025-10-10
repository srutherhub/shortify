package main

import (
	"shortify/db"
	"shortify/server"
)

func main() {
	db.Init()
	server.Init()
}
