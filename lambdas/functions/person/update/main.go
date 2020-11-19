package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	inj "github.com/jjoc007/poc-crud-dynamo-golang-terraform/functions/person"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
	personmodel "github.com/jjoc007/poc-crud-dynamo-golang-terraform/model/person"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/service/person"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/util"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger.Debug().Msg("Start lambda update component")
	var personPayload *personmodel.Person
	body:= event.Body
	err := json.Unmarshal([]byte(body), &personPayload)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on decoding body %v", personPayload)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	log.Logger.Debug().Msgf("component %v", personPayload)
	err = inj.Instances["componentService"].(serviceperson.PersonService).Update(personPayload)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the Component %v", personPayload)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
