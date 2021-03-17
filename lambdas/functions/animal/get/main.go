package main

import (
	"context"
	"encoding/json"
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
	log.Logger.Debug().Msgf("Start lambda get animal by id %s", name)
	animal, err := inj.Instances["animalService"].(serviceanimal.AnimalService).Get(name)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR %v", name)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	outAnimal, err := json.Marshal(animal)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR %v", name)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              string(outAnimal),
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
