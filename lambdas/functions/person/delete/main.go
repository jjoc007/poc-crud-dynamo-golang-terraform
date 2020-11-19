package main

import (
	"context"
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
	log.Logger.Debug().Msg("Start lambda delete person")
	log.Logger.Debug().Msgf("component %v", uid)
	err := inj.Instances["personService"].(serviceperson.PersonService).Delete(uid)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the person %v", uid)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        200,
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
