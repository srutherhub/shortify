package main

import (
	"shortify/db"
	"shortify/server"
)

func main() {
	db.Init()
	db.InsertNewURL("https://www.google.com")
	server.Init()
}
