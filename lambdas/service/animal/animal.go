package serviceanimal

import (
	"github.com/jjoc007/poc-crud-dynamo-golang-terraform/log"
	modelanimal "github.com/jjoc007/poc-crud-dynamo-golang-terraform/model/animal"
	repositoryanimal "github.com/jjoc007/poc-crud-dynamo-golang-terraform/repository/animal"
	"github.com/pkg/errors"
)

// AnimalService describes the structure a Pipeline service.
type AnimalService interface {
	Create(*modelanimal.Animal) error
	Update(*modelanimal.Animal) error
	Get(string) (*modelanimal.Animal, error)
	Delete(string) error
}

// New creates and returns a new lock service instance
func New(rep repositoryanimal.AnimalRepository) AnimalService {
	return &animalService{
		repository:             rep,
	}
}

type animalService struct {
	repository             repositoryanimal.AnimalRepository
}

func (s *animalService) Create(resource *modelanimal.Animal) error {
	log.Logger.Debug().Msg("Creating animal on services")

	err := s.repository.Create(resource)
	if err != nil {
		return errors.Wrapf(err, "Error creating animal on services [%s]", resource.Name)
	}
	return nil
}

func (s *animalService) Get(id string) (*modelanimal.Animal, error) {
	log.Logger.Debug().Msg("Getting a animal on services")
	animal, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "Error getting animal on services [%s]", id)
	}
	return animal, nil
}

func (s *animalService) Delete(id string) error {
	log.Logger.Debug().Msg("Deleting a animal on services")
	err := s.repository.Delete(id)
	if err != nil {
		return errors.Wrapf(err, "Deleting a animal on services [%s]", id)
	}
	return nil
}

func (s *animalService) Update(resource *modelanimal.Animal) error {
	log.Logger.Debug().Msg("Updating pipeline on services")
	err := s.repository.Update(resource)
	if err != nil {
		return errors.Wrapf(err, "Error updating animal on services [%s]", resource.Name)
	}
	return nil
}
