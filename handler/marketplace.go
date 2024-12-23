package handler

import (
	"amartha-loan-svc/dto"
	respond "amartha-loan-svc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fund godoc
// @Summary Fund a loan
// @Description Fund a loan after loan approval
// @Tags Marketplace
// @Accept json
// @Produce json
// @Param fundRequest body dto.FundRequest true "Fund Loan Request"
// @Router /amartha/fund [post]
func (h AmarthaHttpServer) Fund(c *gin.Context) {
	in := dto.FundRequest{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		respond.BadRequest(c, err.Error())
		return
	}

	err = validate.Struct(&in)
	if err != nil {
		respond.ValidationError(c, err)
		return
	}

	appErr := h.app.Fund(in)
	if appErr != nil {
		respond.HandleAppError(c, appErr)
		return
	}

	respond.Ok(c, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully funded loan",
	})
}
