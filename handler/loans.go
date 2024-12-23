package handler

import (
	"amartha-loan-svc/dto"
	respond "amartha-loan-svc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (h AmarthaHttpServer) CreateLoan(c *gin.Context) {
	in := dto.CreateLoan{}

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

	out, appErr := h.app.CreateLoan(in)
	if appErr != nil {
		respond.HandleAppError(c, appErr)
		return
	}

	respond.Created(c, gin.H{
		"loan_id": out.ID,
	})
}

func (h AmarthaHttpServer) ApproveLoan(c *gin.Context) {
	in := dto.ApproveLoan{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		respond.BadRequest(c, err.Error())
		return
	}

	loanIDStr := c.Param("loan_id")

	if loanIDStr == "" {
		respond.BadRequest(c, "loan_id is required")
		return
	}

	loanID, err := strconv.Atoi(loanIDStr)
	if err != nil {
		respond.BadRequest(c, "loan_id must be a numeric")
		return
	}

	in.LoanID = int64(loanID)
	err = validate.Struct(&in)
	if err != nil {
		respond.ValidationError(c, err)
		return
	}

	appErr := h.app.ApproveLoan(in)
	if appErr != nil {
		respond.HandleAppError(c, appErr)
		return
	}

	respond.Ok(c, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully approved loan",
	})
}

func (h AmarthaHttpServer) DisburseLoan(c *gin.Context) {
	in := dto.DisburseLoan{}

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

	appErr := h.app.DisburseLoan(in)
	if appErr != nil {
		respond.HandleAppError(c, appErr)
		return
	}

	respond.Ok(c, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully disbursed loan",
	})
}
