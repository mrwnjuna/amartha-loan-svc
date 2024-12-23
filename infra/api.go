package infra

import (
	"amartha-loan-svc/handler"
	"amartha-loan-svc/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app usecase.AmarthaUsecaseInterface) {
	amarthaSvc := handler.NewAmarthaHttpServer(app)
	api := r.Group("/amartha")
	{
		api.POST("/create-loan", amarthaSvc.CreateLoan)
		api.POST("/approve/:loan_id", amarthaSvc.ApproveLoan)
		api.POST("/fund", amarthaSvc.Fund)
		api.POST("/disburse", amarthaSvc.DisburseLoan)
	}
}
