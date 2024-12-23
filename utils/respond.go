package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func BadRequest(c *gin.Context, message string, details ...interface{}) {
	obj := gin.H{"status": http.StatusBadRequest, "message": message}
	if len(details) > 0 {
		obj["details"] = details[0]
	}
	c.JSON(http.StatusBadRequest, obj)
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
}

func InternalServerError(c *gin.Context, message string, details ...interface{}) {
	obj := gin.H{"status": http.StatusInternalServerError, "message": message}
	if len(details) > 0 {
		obj["details"] = details[0]
	}
	c.JSON(http.StatusInternalServerError, obj)
}

func ValidationError(c *gin.Context, err error) {
	var errors []string
	for _, fieldErr := range err.(validator.ValidationErrors) {
		errors = append(errors, fieldErr.StructField()+" is "+fieldErr.Tag())
	}
	BadRequest(c, "Validation failed", gin.H{"errors": errors})
}

// HandleAppError sends a response based on an AppError struct
func HandleAppError(c *gin.Context, appError *AppError) {
	// Based on the AppError code, send the appropriate response
	switch appError.StatusCode {
	case http.StatusBadRequest:
		BadRequest(c, appError.Message, appError.Details)
	case http.StatusNotFound:
		NotFound(c, appError.Message)
	default:
		InternalServerError(c, appError.Message, appError.Details)
	}
}
