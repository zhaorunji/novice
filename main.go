package main

import (
	"novice/db"
	"novice/router"
)

func main() {
	db.GetDb()
	router.OrderRouter()
}
