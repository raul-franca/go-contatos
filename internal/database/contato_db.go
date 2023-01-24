package database

import (
	"github.com/raul-franca/go-contatos/internal/entity"
	"gorm.io/gorm"
)

type ContatoDB struct {
	DB *gorm.DB
}

func NewContatoDB(db *gorm.DB) *ContatoDB {
	return &ContatoDB{DB: db}
}

// Create cria um novo contato no banco
func (c *ContatoDB) Create(contato *entity.Contato) error {
	return c.DB.Create(&contato).Error
}

// FindAll retornar Todos os contatos do banco
func (c *ContatoDB) FindAll(page, limit int, sort string) ([]entity.Contato, error) {

	var (
		contatos []entity.Contato
		err      error
	)

	//regra de ordenação
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	//regras para a paginação
	if page != 0 && limit != 0 {
		// .Offset((page - 1) * limit)  para corrigir a contagem humana
		err = c.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&contatos).Error
	} else {
		//Ordena apenas pela data de criação.
		err = c.DB.Order("created_at " + sort).Find(&contatos).Error
	}
	return contatos, err
}

func (c *ContatoDB) FindByID(id string) (*entity.Contato, error) {
	var contato entity.Contato
	err := c.DB.Where("id = ?", id).First(&contato).Error
	if err != nil {
		return nil, err
	}
	return &contato, nil
}

func (c *ContatoDB) FindByName(name string) (*[]entity.Contato, error) {
	var contato []entity.Contato
	err := c.DB.Where("nome like ?", "%"+name+"%").Find(&contato).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &contato, nil
}
func (c *ContatoDB) FindByEmail(email string) (*[]entity.Contato, error) {
	var contato []entity.Contato
	err := c.DB.Where("email like ?", "%"+email+"%").Find(&contato).Error
	if err != nil {
		return nil, err
	}
	return &contato, nil
}
func (c *ContatoDB) Update(contato *entity.Contato) error {
	_, err := c.FindByID(contato.ID.String())
	if err != nil {
		return err
	}
	return c.DB.Save(contato).Error

}
func (c *ContatoDB) Delete(id string) error {
	contato, err := c.FindByID(id)
	if err != nil {
		return err
	}
	return c.DB.Delete(contato).Error
}
