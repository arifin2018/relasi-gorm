package main

import (
	"relasi-gorm/databases"
	"relasi-gorm/routes"
)

func main() {
	databases.DatabaseInit()
	routes.Router()
	
}