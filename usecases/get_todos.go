package usecases

import "github.com/shutsuchi/goTut_Go/models"

func GetTodos(repo TodosRepository) ([]models.Todo, error) {
	todos, err := repo.GetAllTodos()
	if err != nil {
		return nil, ErrInternal
	}
  return todos, nil
}