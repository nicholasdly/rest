package store

import (
	"github.com/nicholasdly/rest/internal/models"
)

type UserStore interface {
	GetAll() ([]models.User, error)
	Get(id int) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id int) error
}
