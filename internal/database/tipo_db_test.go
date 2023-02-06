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

func ConexaoTipoDB() (*gorm.DB, *TipoDB) {

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

	_, tipoDB := ConexaoTipoDB()

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

func TestTipoDB_Create2(t *testing.T) {
	_, tipoDB := ConexaoTipoDB()
	novoTipo, _ := entity.NewTipo("Amigos", "#fff3333")

	err := tipoDB.Create(novoTipo)
	assert.Empty(t, err)
}

func TestTipoDB_FindByID(t *testing.T) {
	_, tipoDB := ConexaoTipoDB()
	novoTipo, _ := entity.NewTipo("Amigos", "#fff3333")
	err := tipoDB.Create(novoTipo)
	assert.Empty(t, err)

	listCreatedTipos, err := tipoDB.ListAll()

	for _, tipoFound := range listCreatedTipos {
		tipoFound, err := tipoDB.FindByID(tipoFound.ID)
		assert.Empty(t, err)
		assert.NotNil(t, tipoFound)
	}

	tipoFound, err := tipoDB.FindByID(3)
	assert.Error(t, err)
	assert.Nil(t, tipoFound)

}
func TestTipoDB_FindByName(t *testing.T) {

	_, tipoDB := ConexaoTipoDB()

	tipos := []entity.Tipo{
		{Nome: "Familia", Color: "#fff000"},
		{Nome: "Colegas", Color: "#fff000"},
		{Nome: "Amigos", Color: "#fff000"},
	}
	for _, tipo := range tipos {
		tipoDB.DB.Create(&tipo)
	}

	tiposfound, err := tipoDB.FindByName("a")
	assert.Nil(t, err)
	assert.NotEmpty(t, tiposfound)

	for _, founds := range *tiposfound {
		fmt.Println(founds)
	}

	tiposfound, _ = tipoDB.FindByName("Amigos")
	assert.Nil(t, err)
	assert.NotEmpty(t, tiposfound)

}

func TestTipoDB_Update(t *testing.T) {
	_, tipoDB := ConexaoTipoDB()

	tipo, err := entity.NewTipo("Familia", "#fff")
	assert.Nil(t, err)
	err = tipoDB.Create(tipo)
	if err != nil {
		fmt.Println(err)
	}

	tipo.Nome = "Amigos"
	tipo.Color = "#000"
	err = tipoDB.Update(tipo)
	tipoFound, err := tipoDB.FindByID(1)
	assert.Nil(t, err)
	assert.Equal(t, tipo.Nome, tipoFound.Nome)
	assert.Equal(t, tipo.Color, tipoFound.Color)
}

func TestTipoDB_Delete(t *testing.T) {

	_, tipoDB := ConexaoTipoDB()
	tipo, err := entity.NewTipo("Familia", "#fff")
	assert.NoError(t, err)
	tipoDB.Create(tipo)

	err = tipoDB.Delete(1)
	assert.NoError(t, err)

	_, err = tipoDB.FindByID(tipo.ID)
	assert.NotNil(t, err)

}
