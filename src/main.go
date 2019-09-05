package main

import (
	"db"
	"router"
)

func main() {
	db.GetDb()
	router.OrderRouter()
}
