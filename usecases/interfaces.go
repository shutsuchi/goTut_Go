package usecases

import "github.com/shutsuchi/goTut_Go/models"

type TodosRepository interface {
	GetAllTodos() ([]models.Todo, error)
}