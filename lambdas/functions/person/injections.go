package person

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/config/db"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
	repositoryperson "github.com/jjoc007/poc-crud-dynamo-golang-terraform/repository/person"
	serviceperson "github.com/jjoc007/poc-crud-dynamo-golang-terraform/service/person"
)

// Instances is a global map that contain all object instances of app
var Instances = MakeDependencyInjection()

// MakeDependencyInjection Initialize all dependencies
func MakeDependencyInjection() map[string]interface{} {
	log.Logger.Debug().Msg("Start bootstrap app objects")
	instances := make(map[string]interface{})

	database, err := db.NewDynamoDBStorage()
	if err != nil {
		panic(err)
	}
	instances["dataBase"] = database

	instances["personRepository"] = repositoryperson.NewRepository(database.GetConnection().(*dynamodb.DynamoDB))

	instances["personService"] = serviceperson.New(
		instances["personRepository"].(repositoryperson.PersonRepository))

	log.Logger.Debug().Msg("End bootstrap app objects")
	return instances
}
