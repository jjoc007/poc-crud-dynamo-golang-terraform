package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	inj "github.com/jjoc007/poc-crud-dynamo-golang-terraform/functions/animal"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
	animalmodel "github.com/jjoc007/poc-crud-dynamo-golang-terraform/model/animal"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/service/animal"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/util"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger.Debug().Msg("Start lambda create animal")
	var animalPayload *animalmodel.Animal
	body:= event.Body
	err := json.Unmarshal([]byte(body), &animalPayload)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on decoding body %v", animalPayload)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	log.Logger.Debug().Msgf("animal %v , %+v", animalPayload, cxt)

	err = inj.Instances["animalService"].(serviceanimal.AnimalService).Create(animalPayload)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the animal")
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
