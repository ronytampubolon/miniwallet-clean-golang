package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Responses struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type ErrorResponse struct {
	Status string      `json:"status"`
	Data   ErrorDetail `json:"data"`
}
type ErrorDetail struct {
	Error interface{} `json:"error"`
}

func APIResponse(ctx *gin.Context, StatusCode int, Message string, Data interface{}) {

	jsonResponse := Responses{
		Status: Message,
		Data:   Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Error interface{}) {
	errResponse := ErrorResponse{
		Status: "fail",
		Data: ErrorDetail{
			Error: Error,
		},
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}

func UnauthorizedError(ctx *gin.Context) {
	errResponse := ErrorResponse{
		Status: "fail",
		Data: ErrorDetail{
			Error: "unauthorized",
		}}

	ctx.JSON(http.StatusUnauthorized, errResponse)
	defer ctx.AbortWithStatus(http.StatusUnauthorized)
}
