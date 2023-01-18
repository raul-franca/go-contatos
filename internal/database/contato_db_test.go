package database

import (
	"fmt"
	"github.com/raul-franca/go-contatos/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateContato(t *testing.T) {
	//Criar conexão com banco
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	//Criar tabela de contatos
	db.AutoMigrate(&entity.Contato{})
	//cria a Conexão a tabela no banco
	contatoDB := NewContatoDB(db)

	//Criar um contato
	contato, _ := entity.NewContato("raul", "email@em.com", "", 1)

	err = contatoDB.Create(contato)
	assert.Nil(t, err)

	var contFound entity.Contato
	err = db.First(&contFound, "id = ?", contato.ID).Error

	fmt.Print("Resultado contFound: ", contFound)

	assert.Nil(t, err)
	assert.Equal(t, contato.ID, contFound.ID)
	assert.Equal(t, contato.Nome, contFound.Nome)
	assert.Equal(t, contato.Email, contFound.Email)
	assert.NotNil(t, contFound.Tipo)
	assert.NotNil(t, contFound.CreatedAt)
	assert.True(t, contFound.Ativo)
}

func TestFindAll(t *testing.T) {

	//Criar conexão com banco
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	//Criar tabela de contatos
	db.AutoMigrate(&entity.Contato{})
	//cria a Conexão a tabela no banco
	contatoDB := NewContatoDB(db)

	for i := 1; i <= 25; i++ {
		//cria varios contatos e add no db
		contato, err := entity.NewContato(fmt.Sprintf("Nome %d", i), "email@email.com", "", 1)
		assert.NoError(t, err)
		db.Create(contato)
	}
	contatos, err := contatoDB.FindAll(1, 10, "asc")
	fmt.Println("pag 1, limt 10:")
	for _, c := range contatos {
		fmt.Printf("Nome: %s e ID: %s \n", c.Nome, c.ID)
	}
	contatos, err = contatoDB.FindAll(2, 10, "asc")
	fmt.Println("pag 2, limt 10:")
	for _, c := range contatos {
		fmt.Printf("Nome: %s e ID: %s \n", c.Nome, c.ID)
	}
	contatos, err = contatoDB.FindAll(3, 10, "asc")
	fmt.Println("pag 3, limt 10:")
	for _, c := range contatos {
		fmt.Printf("Nome: %s e ID: %s \n", c.Nome, c.ID)
	}

}
