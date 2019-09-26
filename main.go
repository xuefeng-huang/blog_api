package main

import (
	"blog_api/db"
	"blog_api/server"
)

func main() {
	db.Init()
	server.Init()
}
