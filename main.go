package main

import (
	"blog_api/db"
	"blog_api/server"
)

func main() {
	db.Init("blog")
	server.Init()
}
