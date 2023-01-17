package entity

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

// o email
// new retorna um contato valido ...

func TestNewContato(t *testing.T) {
	contato, err := NewContato("Aurora", "aurora@email.com", "", 1)
	assert.Nil(t, err)
	assert.NotEmpty(t, contato)
	assert.NotEmpty(t, contato.ID)
	assert.Equal(t, "Aurora", contato.Nome)
	assert.Equal(t, "aurora@email.com", contato.Email)
	assert.Equal(t, "", contato.OBS)
	assert.NotEmpty(t, contato.Tipo)
	assert.NotEmpty(t, contato.CreatedAt)
	assert.True(t, contato.Ativo)
}
