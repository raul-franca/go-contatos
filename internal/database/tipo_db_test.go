package database

import (
	"fmt"
	"github.com/raul-franca/go-contatos/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

//	Create(tipo *entity.Tipo) error
//	ListAll() ([]entity.Tipo, error)
//	FindByID(id int) (*entity.Tipo, error)
//	FindByName(name string) (*[]entity.Tipo, error)
//	Update(tipo *entity.Tipo) error
//	Delete(id int) error

func ConxaoTipoDB() (*gorm.DB, *TipoDB) {

	//Criar conexão com banco
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	//Criar tabela de contatos
	db.AutoMigrate(&entity.Tipo{})
	//cria a Conexão a tabela no banco
	TipoDB := NewTipoDB(db)

	return db, TipoDB
}

func TestTipoDB_Create(t *testing.T) {

	_, tipoDB := ConxaoTipoDB()

	tipos := []entity.Tipo{
		{Nome: "Familia", Color: "#fff000"},
		{Nome: "Colegas", Color: "#fff000"},
		{Nome: "Amigos", Color: "#fff000"},
	}
	for _, tipo := range tipos {
		tipoDB.DB.Create(&tipo)
	}

	listCreatedTipos, err := tipoDB.ListAll()
	if err != nil {
		t.Error("Error ao listar os criados, ", err)
	}
	for _, tipo := range listCreatedTipos {
		fmt.Println(tipo)
	}

}
