package app

import (
	"amartha-loan-svc/infra"
	postgres "amartha-loan-svc/infra/postgres"
	"amartha-loan-svc/repository"
	"amartha-loan-svc/usecase"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	amarthaRepo := repository.NewAmarthaRepo(postgres.PSQL.DB.DB)
	app := usecase.NewAmarthaUsecase(amarthaRepo)
	infra.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
