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

var (
	DefaultSuccessResponse = Response{
		En: "Success",
		Id: "Sukses",
	}
	DefaultBadRequestResponse = Response{
		En: "Bad request",
		Id: "Permintaan buruk",
	}
	DefaultInvalidRequestBodyOrParam = Response{
		En: "Invalid request body or param",
		Id: "Request body atau param tidak valid",
	}
)

type Response struct {
	En string `json:"en"`
	Id string `json:"id"`
}

func ReturnError(c *gin.Context, errRes error) {
	errStr := errRes.Error()
	errSplit := strings.SplitN(errStr, ":", 2)
	msg := new(Response)
	if len(errSplit) > 1 {
		_ = json.Unmarshal([]byte(errSplit[1]), &msg)
	}
	statusCode, err := strconv.Atoi(errSplit[0])
	if err != nil || msg == nil {
		c.JSON(http.StatusInternalServerError, Response{
			En: "Internal server error",
			Id: "Terjadi kesalahan di server",
		})
		return
	}

	c.JSON(statusCode, msg)
}

func ErrorValidation(err *Response) error {
	msgByte, _ := json.Marshal(err)
	return fmt.Errorf("%d:%s", http.StatusUnprocessableEntity, string(msgByte))
}

func ErrorBadRequest(err *Response) error {
	msgByte, _ := json.Marshal(err)
	return fmt.Errorf("%d:%s", http.StatusBadRequest, string(msgByte))
}

func DefaultErrorBadRequest(c *gin.Context, errType string) {
	var msg = new(Response)
	switch errType {
	case InvalidRequestBody:
		msg = &DefaultInvalidRequestBodyOrParam
	default:
		msg = &DefaultBadRequestResponse
	}

	c.JSON(http.StatusBadRequest, msg)
}

// SuccessCreate if there is no data, can fill data with nil
func SuccessCreate(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(http.StatusCreated, DefaultSuccessResponse)
	} else {
		c.JSON(http.StatusCreated, data)
	}
}
