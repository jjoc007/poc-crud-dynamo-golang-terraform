package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	inj "github.com/jjoc007/poc-crud-dynamo-golang-terraform/functions/person"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/service/person"

	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/util"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	uid := event.QueryStringParameters["uid"]
	log.Logger.Debug().Msgf("Start lambda get component by id %s", uid)
	person, err := inj.Instances["personService"].(serviceperson.PersonService).Get(uid)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR %v", uid)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	outPerson, err := json.Marshal(person)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR %v", uid)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Body:              string(outPerson),
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
