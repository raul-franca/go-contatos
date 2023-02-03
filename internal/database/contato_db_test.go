package database

import (
	"fmt"
	"github.com/raul-franca/go-contatos/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func ConxaoDB() (*gorm.DB, *ContatoDB) {

	//Criar conexão com banco
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	//Criar tabela de contatos
	db.AutoMigrate(&entity.Contato{})
	//cria a Conexão a tabela no banco
	contatoDB := NewContatoDB(db)

	return db, contatoDB
}

func TestCreateContato(t *testing.T) {

	db, contatoDB := ConxaoDB()

	//Criar um contato
	contato, _ := entity.NewContato("raul", "email@em.com", "", 1)

	err := contatoDB.Create(contato)
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

	db, contatoDB := ConxaoDB()

	for i := 1; i <= 25; i++ {
		//cria varios contatos e add no db
		contato, err := entity.NewContato(fmt.Sprintf("Nome %d", i), "email@email.com", "", 1)
		db.Create(contato)
		assert.NoError(t, err)
	}
	contatos, err := contatoDB.FindAll(1, 10, "asc")

	assert.Nil(t, err)
	assert.NotNil(t, contatos)
	assert.NotNil(t, contatos[0])
	assert.Equal(t, "Nome 1", contatos[0].Nome)
	assert.Equal(t, "email@email.com", contatos[0].Email)
	assert.NotNil(t, contatos[0].Tipo)
	assert.NotNil(t, contatos[0].CreatedAt)
	assert.True(t, contatos[0].Ativo)

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

	contatos, err = contatoDB.FindAll(1, 10, "")
	assert.Nil(t, err)
	assert.NotNil(t, contatos)
	assert.NotNil(t, contatos[0])
	assert.Equal(t, "Nome 1", contatos[0].Nome)
	assert.Equal(t, "email@email.com", contatos[0].Email)
	assert.NotNil(t, contatos[0].Tipo)
	assert.NotNil(t, contatos[0].CreatedAt)
	assert.True(t, contatos[0].Ativo)

}

func TestContatoDB_FindAll2(t *testing.T) {
	db, contatoDB := ConxaoDB()

	for i := 1; i <= 8; i++ {
		//cria varios contatos e add no db
		contato, err := entity.NewContato(fmt.Sprintf("Nome %d", i), "email@email.com", "", 1)
		db.Create(contato)
		assert.NoError(t, err)
	}
	contatos, err := contatoDB.FindAll(0, 20, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, contatos)
	assert.NotNil(t, contatos[0])
	assert.Equal(t, "Nome 1", contatos[0].Nome)
	assert.Equal(t, "email@email.com", contatos[0].Email)
	assert.NotNil(t, contatos[0].Tipo)
	assert.NotNil(t, contatos[0].CreatedAt)
	assert.True(t, contatos[0].Ativo)
	for _, c := range contatos {
		fmt.Printf("Nome: %s e ID: %s \n", c.Nome, c.ID)
	}
}

func TestContatoDB_FindByID(t *testing.T) {

	db, contatoDB := ConxaoDB()

	//Criar um contato
	contato, _ := entity.NewContato("raul", "email@em.com", "", 1)

	err := contatoDB.Create(contato)
	assert.Nil(t, err)

	var contID entity.Contato
	err = db.First(&contID).Error

	contatoFound, err := contatoDB.FindByID(contID.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, contato.Nome, contatoFound.Nome)
	assert.Equal(t, contato.Email, contatoFound.Email)
	assert.NotNil(t, contatoFound.Tipo)
	assert.NotNil(t, contatoFound.CreatedAt)
	assert.True(t, contatoFound.Ativo)

}

func TestContatoDB_FindByEmail(t *testing.T) {

	_, contatoDB := ConxaoDB()

	//Criar um contato
	contato, _ := entity.NewContato("raul", "email@em.com", "", 1)

	err := contatoDB.Create(contato)
	assert.Nil(t, err)

	contatoFounds, _ := contatoDB.FindByEmail("email@e")
	assert.Nil(t, err)

	for _, contatoFound := range *contatoFounds {
		assert.Equal(t, contato.Nome, contatoFound.Nome)
		assert.Equal(t, contato.Email, contatoFound.Email)
		assert.NotNil(t, contatoFound.Tipo)
		assert.NotNil(t, contatoFound.CreatedAt)
		assert.True(t, contatoFound.Ativo)
	}
}

func TestContatoDB_FindByName(t *testing.T) {
	_, contatoDB := ConxaoDB()

	//Criar um contato
	contato, _ := entity.NewContato("raul", "email@em.com", "", 1)

	err := contatoDB.Create(contato)
	assert.Nil(t, err)

	contatoFounds, _ := contatoDB.FindByName("ra")
	assert.Nil(t, err)

	for _, contatoFound := range *contatoFounds {
		assert.Equal(t, contato.Nome, contatoFound.Nome)
		assert.Equal(t, contato.Email, contatoFound.Email)
		assert.NotNil(t, contatoFound.Tipo)
		assert.NotNil(t, contatoFound.CreatedAt)
		assert.True(t, contatoFound.Ativo)
	}
}

func TestContatoDB_Update(t *testing.T) {
	_, contatoDB := ConxaoDB()

	contato, _ := entity.NewContato("raul", "email@email.com", "", 1)
	contatoDB.Create(contato)

	contato.Nome = "Aurora"
	contato.Email = "aurora@email.com"
	contatoDB.Update(contato)
	contatoFound, _ := contatoDB.FindByID(contato.ID.String())
	assert.Equal(t, "Aurora", contatoFound.Nome)
	assert.Equal(t, "aurora@email.com", contatoFound.Email)

}
func TestContatoDB_Delete(t *testing.T) {
	_, contatoDB := ConxaoDB()
	contato, err := entity.NewContato("raul", "email@email.com", "", 1)
	assert.NoError(t, err)
	contatoDB.Create(contato)

	err = contatoDB.Delete(contato.ID.String())
	assert.NoError(t, err)

	_, err = contatoDB.FindByID(contato.ID.String())
	assert.Error(t, err)
}
