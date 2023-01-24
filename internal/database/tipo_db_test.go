package database

import (
	"fmt"
	"github.com/raul-franca/go-contatos/internal/entity"
	"github.com/stretchr/testify/assert"
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

	db, tipoDB := ConxaoTipoDB()

	novoTipo := entity.Tipo{Nome: "Familia", Color: "#fff000"}

	err := tipoDB.Create(&novoTipo)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)
	tipoFound := entity.Tipo{}
	err = db.Last(&tipoFound).Error
	fmt.Println("Resultado tipoFound: ", tipoFound)

	//criando mais um contato
	outroTipo := entity.Tipo{Nome: "Colegas", Color: "#fff000"}
	err = tipoDB.Create(&outroTipo)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)
	tipoFound2 := entity.Tipo{}
	err = db.Last(&tipoFound2).Error
	fmt.Println("Resultado tipoFound: ", tipoFound2)

	//criando mais um contato
	outroTipo2 := entity.Tipo{Nome: "Amigos", Color: "#fff000"}
	err = tipoDB.Create(&outroTipo2)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)
	tipoFound3 := entity.Tipo{}
	err = db.Last(&tipoFound3).Error
	fmt.Println("Resultado tipoFound: ", tipoFound3)

	tipos, _ := tipoDB.ListAll()

	for _, tipo := range tipos {
		fmt.Println("lisAll tipos", tipo)
	}
}
