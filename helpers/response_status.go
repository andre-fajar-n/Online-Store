package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	InvalidRequestBody = "Invalid request body"
)

type ErrorResponse struct {
	En string `json:"en"`
	Id string `json:"id"`
}

func ReturnError(c *gin.Context, errRes error) {
	errStr := errRes.Error()
	errSplit := strings.SplitN(errStr, ":", 2)
	msg := new(ErrorResponse)
	if len(errSplit) > 1 {
		_ = json.Unmarshal([]byte(errSplit[1]), &msg)
	}
	statusCode, err := strconv.Atoi(errSplit[0])
	if err != nil || msg == nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			En: "Internal server error",
			Id: "Terjadi kesalahan di server",
		})
		return
	}

	c.JSON(statusCode, msg)
}

func ErrorValidation(err *ErrorResponse) error {
	msgByte, _ := json.Marshal(err)
	return fmt.Errorf("%d:%s", http.StatusUnprocessableEntity, string(msgByte))
}

func ErrorBadRequest(err *ErrorResponse) error {
	msgByte, _ := json.Marshal(err)
	return fmt.Errorf("%d:%s", http.StatusUnprocessableEntity, string(msgByte))
}

func DefaultErrorBadRequest(c *gin.Context, errType string) {
	msg := new(ErrorResponse)
	switch errType {
	case InvalidRequestBody:
		msg = &ErrorResponse{
			En: "Invalid request body",
			Id: "Request body tidak valid",
		}
	default:
		msg = &ErrorResponse{
			En: "Bad request",
			Id: "Permintaan buruk",
		}
	}

	c.JSON(http.StatusBadRequest, msg)
}

// SuccessCreate if there is no data, can fill data with nil
func SuccessCreate(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(http.StatusCreated, gin.H{
			"en": "Success",
			"id": "Sukses",
		})
	} else {
		c.JSON(http.StatusCreated, data)
	}
}
