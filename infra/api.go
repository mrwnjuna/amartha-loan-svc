package infra

import (
	_ "amartha-loan-svc/docs" // Import generated Swagger docs
	"amartha-loan-svc/handler"
	"amartha-loan-svc/usecase"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app usecase.AmarthaUsecaseInterface) {
	amarthaSvc := handler.NewAmarthaHttpServer(app)
	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/amartha")
	{
		api.POST("/create-loan", amarthaSvc.CreateLoan)
		api.POST("/approve/:loan_id", amarthaSvc.ApproveLoan)
		api.POST("/fund", amarthaSvc.Fund)
		api.POST("/disburse", amarthaSvc.DisburseLoan)
	}
}
