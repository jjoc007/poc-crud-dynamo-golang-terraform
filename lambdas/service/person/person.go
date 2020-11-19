package serviceperson

import (
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
	modelperson "github.com/jjoc007/poc-crud-dynamo-golang-terraform/model/person"
	repositoryperson "github.com/jjoc007/poc-crud-dynamo-golang-terraform/repository/person"
	"github.com/pkg/errors"
)

// PersonService describes the structure a Person service.
type PersonService interface {
	Create(*modelperson.Person) error
	Update(*modelperson.Person) error
	Get(string) (*modelperson.Person, error)
	Delete(string) error
}

func New(rep repositoryperson.PersonRepository) PersonService {
	return &personService{
		repository:             rep,
	}
}

type personService struct {
	repository             repositoryperson.PersonRepository
}

func (s *personService) Create(resource *modelperson.Person) error {
	log.Logger.Debug().Msg("Creating person on services")

	err := s.repository.Create(resource)
	if err != nil {
		return errors.Wrapf(err, "Error creating person on services [%s]", resource.UID)
	}
	return nil
}

func (s *personService) Get(id string) (*modelperson.Person, error) {
	log.Logger.Debug().Msg("Getting a person on services")
	animal, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "Error getting person on services [%s]", id)
	}
	return animal, nil
}

func (s *personService) Delete(id string) error {
	log.Logger.Debug().Msg("Deleting a person on services")
	err := s.repository.Delete(id)
	if err != nil {
		return errors.Wrapf(err, "Deleting a person on services [%s]", id)
	}
	return nil
}

func (s *personService) Update(resource *modelperson.Person) error {
	log.Logger.Debug().Msg("Updating person on services")
	err := s.repository.Update(resource)
	if err != nil {
		return errors.Wrapf(err, "Error updating person on services [%s]", resource.UID)
	}
	return nil
}
