package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	inj "github.com/jjoc007/poc-crud-dynamo-golang-terraform/functions/animal"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/service/animal"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/util"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := event.QueryStringParameters["name"]
	log.Logger.Debug().Msg("Start lambda delete animal")
	log.Logger.Debug().Msgf("animal %v", name)
	err := inj.Instances["animalService"].(serviceanimal.AnimalService).Delete(name)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the animal %v", name)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        200,
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
