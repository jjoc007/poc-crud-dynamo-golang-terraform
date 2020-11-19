package repositoryperson

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	personmodel "github.com/jjoc007/poc-crud-dynamo-golang-terraform/model/person"

	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/config"
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
)

// PersonRepository describes the lock repository.
type PersonRepository interface {
	Create(*personmodel.Person) error
	Update(*personmodel.Person) error
	GetByID(string) (*personmodel.Person, error)
	Delete(string) error
}

// NewRepository creates and returns a new Websocket repository instance
func NewRepository(database *dynamodb.DynamoDB) PersonRepository {
	return &repository{
		database: database,
		table:    config.POCPersonTable,
	}
}

type repository struct {
	database *dynamodb.DynamoDB
	table    string
}

func (s *repository) Create(resource *personmodel.Person) (err error) {
	log.Logger.Debug().Msgf("Adding a new person [%s] ", resource.UID)

	av, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	})

	log.Logger.Debug().Msgf("ID %v inserted.\n", resource.UID)
	return nil
}

func (s *repository) Update(resource *personmodel.Person) (err error) {
	log.Logger.Debug().Msgf("Adding a new animal [%s] ", resource.UID)

	av, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	})

	log.Logger.Debug().Msgf("ID %v inserted.\n", resource.UID)
	return nil
}

func (s *repository) GetByID(id string) (person *personmodel.Person, err error) {
	log.Logger.Debug().Msgf("Getting Animal by ID")

	result, err := s.database.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(s.table),
		Key: map[string]*dynamodb.AttributeValue{
			"uid": {
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

	person = &personmodel.Person{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &person)
	if err != nil {
		return
	}

	return
}

func (s repository) Delete(id string) (err error) {
	log.Logger.Debug().Msgf("Deleting an person [%s] ", id)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"uid": {
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
