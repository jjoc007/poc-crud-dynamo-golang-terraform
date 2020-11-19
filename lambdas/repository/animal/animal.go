package repositoryanimal

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	animalmodel "github.com/jjoc007/poc-crud-dynamo-golang-terraform/model/animal"

	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/config"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
)

// AnimalRepository describes the lock repository.
type AnimalRepository interface {
	Create(*animalmodel.Animal) error
	Update(*animalmodel.Animal) error
	GetByID(string) (*animalmodel.Animal, error)
	Delete(string) error
}

// NewRepository creates and returns a new Websocket repository instance
func NewRepository(database *dynamodb.DynamoDB) AnimalRepository {
	return &repository{
		database: database,
		table:    config.POCAnimalTable,
	}
}

type repository struct {
	database *dynamodb.DynamoDB
	table    string
}

func (s *repository) Create(resource *animalmodel.Animal) (err error) {
	log.Logger.Debug().Msgf("Adding a new animal [%s] ", resource.Name)

	av, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	})

	log.Logger.Debug().Msgf("ID %v inserted.\n", resource.Name)
	return nil
}

func (s *repository) Update(resource *animalmodel.Animal) (err error) {
	log.Logger.Debug().Msgf("Adding a new animal [%s] ", resource.Name)

	av, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	})

	log.Logger.Debug().Msgf("ID %v inserted.\n", resource.Name)
	return nil
}

func (s *repository) GetByID(id string) (animal *animalmodel.Animal, err error) {
	log.Logger.Debug().Msgf("Getting Animal by ID")

	result, err := s.database.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(s.table),
		Key: map[string]*dynamodb.AttributeValue{
			"animal_name": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if result.Item == nil {
		return nil, errors.New("animal not found")
	}

	animal = &animalmodel.Animal{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &animal)
	if err != nil {
		return
	}

	return
}

func (s repository) Delete(id string) (err error) {
	log.Logger.Debug().Msgf("Deleting an animal [%s] ", id)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"animal_name": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(s.table),
	}

	_, err = s.database.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}

	return nil
}
