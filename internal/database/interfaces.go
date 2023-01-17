package database

import (
	"github.com/raul-franca/go-contatos/internal/entity"
)

type ContatoInterface interface {
	Create(contato *entity.Contato) error
	FindAll(page, limit int, sort string) ([]entity.Contato, error)
	FindByID(id string) (*entity.Contato, error)
	FindByName(name string) (*entity.Contato, error)
	Update(contato *entity.Contato) error
	Delete(id string) error
}

type TipoInterface interface {
	Create(tipo *entity.Tipo) error
	ListAll() ([]entity.Tipo, error)
	FindByID(id int) (*entity.Tipo, error)
	FindByName(name string) (*entity.Tipo, error)
	Update(tipo *entity.Tipo) error
	Delete(id string) error
}


