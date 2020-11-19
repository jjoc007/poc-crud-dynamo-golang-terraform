package util

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"

	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"

	"net/http"
)

type Type int

const (
	BadRequest Type = iota
	NotFound
	InternalError
)

var errorsID = map[Type]string{
	BadRequest:    "BAD_REQUEST",
	NotFound:      "NOT_FOUND",
	InternalError: "INTERNAL_ERROR",
}

type GenericError struct {
	// in: body
	Body Body `json:"body"`
}

type Body struct {
	Code         int32  `json:"code"`
	Message      string `json:"message"`
	MessageError error  `json:"messageError"`
}

type App struct {
	Type    Type        `json:"errorType"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func buildGenericError(msg string, err error) []byte {
	genericError := new(GenericError)
	genericError.Body = Body{
		Code:         500,
		Message:      msg,
		MessageError: err,
	}
	data, err := json.Marshal(genericError)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Error Marshal generic errorapp")
		data = []byte(fmt.Sprintf("500 - Internal server errorapp: %s", err))
	}
	return data
}

// ResponseErrorFunction build a generic error response to http
func ResponseErrorFunction(err error, msg string) events.APIGatewayProxyResponse {
	log.Logger.Error().Err(err).Msg(msg)
	data := buildGenericError(msg, err)
	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       string(data),
	}
	return response
}

func ErrorHandler(err error, data interface{}) (*App, error) {
	switch err.Error() {
	case errorsID[NotFound]:
		return New(NotFound, err.Error(), data), nil

	default:
		return nil, err
	}
}

func New(t Type, m string, d interface{}) *App {
	return &App{
		Type:    t,
		Message: m,
		Data:    d,
	}
}