package main

import (
	"amartha-loan-svc/app"
	infra "amartha-loan-svc/infra/postgres"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	if err := infra.InitPostgre(); err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
