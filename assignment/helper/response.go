package helper

import (
	"net/http"

	"github.com/claravelita/training-golang-mnc/dtos"
)

func BadRequestResponse() (response dtos.StandardResponse) {
	response.Code = http.StatusBadRequest
	response.Status = "failed"
	response.Data = nil
	return response
}

func SuccessResponse(data interface{}) (response dtos.StandardResponse) {
	response.Code = http.StatusOK
	response.Status = "success"
	response.Data = data
	return response
}
