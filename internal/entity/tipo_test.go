package entity

import (
	"github.com/raul-franca/go-contatos/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTipo(t *testing.T) {
	tipo, err := NewTipo("Familia", "#fff")
	assert.Nil(t, err)
	assert.Equal(t, "Familia", tipo.Nome)
	assert.Equal(t, "#fff", tipo.Color)
	assert.Empty(t, tipo.ID)
	assert.NotEmpty(t, tipo.CreatedAt)
	assert.True(t, tipo.Ativo)

	tipo, err = NewTipo("Familia", "")
	assert.Error(t, err, ErrColorIsInvalid)

	tipo, err = NewTipo("", "#fff")
	assert.Error(t, err, ErrNomeIsRequired)

}

func TestColor(t *testing.T) {
	color := entity.IsHex("#101010")
	assert.True(t, color)
}
