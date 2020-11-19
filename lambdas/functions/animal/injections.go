package animal

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/config/db"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
	repositoryanimal "github.com/jjoc007/poc-crud-dynamo-golang-terraform/repository/animal"
	serviceanimal "github.com/jjoc007/poc-crud-dynamo-golang-terraform/service/animal"
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

	instances["animalRepository"] = repositoryanimal.NewRepository(database.GetConnection().(*dynamodb.DynamoDB))

	instances["animalService"] = serviceanimal.New(
		instances["animalRepository"].(repositoryanimal.AnimalRepository))

	log.Logger.Debug().Msg("End bootstrap app objects")
	return instances
}
