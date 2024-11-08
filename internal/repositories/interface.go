package repositories

import "github.com/titodelerinofilho/categories-msvc/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
