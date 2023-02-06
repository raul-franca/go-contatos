package database

import (
	"github.com/raul-franca/go-contatos/internal/entity"
	"gorm.io/gorm"
)

type TipoDB struct {
	DB *gorm.DB
}

func NewTipoDB(db *gorm.DB) *TipoDB {
	return &TipoDB{DB: db}
}

func (t *TipoDB) Create(tipo *entity.Tipo) error {
	return t.DB.Create(&tipo).Error
}

func (t *TipoDB) ListAll() ([]entity.Tipo, error) {
	var tipos []entity.Tipo
	err := t.DB.Find(&tipos).Order("asc").Error
	if err != nil {
		return nil, err
	}
	return tipos, nil
}

func (t *TipoDB) FindByID(id int) (*entity.Tipo, error) {
	var tipo entity.Tipo
	err := t.DB.Where("id = ?", id).First(&tipo).Error
	if err != nil {
		return nil, err
	}
	return &tipo, nil
}

func (t *TipoDB) FindByName(name string) (*[]entity.Tipo, error) {
	var tipos []entity.Tipo
	err := t.DB.Where("nome like ?", "%"+name+"%").Find(&tipos).Error
	if err != nil {
		return nil, err
	}
	return &tipos, nil
}

func (t *TipoDB) Update(tipo *entity.Tipo) error {
	_, err := t.FindByID(tipo.ID)
	if err != nil {
		return err
	}
	return t.DB.Save(tipo).Error
}
func (t *TipoDB) Delete(id int) error {
	tipo, err := t.FindByID(id)
	if err != nil {
		return err
	}
	return t.DB.Delete(&tipo).Error
}
