package entity

import (
	"errors"
	"github.com/raul-franca/go-contatos/pkg/entity"
	"time"
)

var (
	ErrColorIsInvalid = errors.New(" Cor invalida")
)

type Tipo struct {
	ID        int       `json:"ID"`
	Nome      string    `json:"nome"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	Ativo     bool      `json:"ativo"`
}

func NewTipo(nome, color string) (*Tipo, error) {
	tipo := Tipo{
		Nome:      nome,
		Color:     color,
		CreatedAt: time.Now(),
		Ativo:     true,
	}
	tipo.Validar()
	return &tipo, nil
}

func (c *Tipo) Validar() error {

	if c.Nome == "" {
		return ErrNomeIsRequired
	}
	if !entity.IsHex(c.Color) {
		return ErrColorIsInvalid
	}

	return nil
}
