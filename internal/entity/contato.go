package entity

import (
	"errors"
	"github.com/raul-franca/go-contatos/pkg/entity"
	"time"
)

var (
	ErrNomeIsRequired  = errors.New("Nome é necessario ")
	ErrEmailIsRequired = errors.New("E-mail é necessario ")
	ErrEmailIsInvalid  = errors.New("E-mail invalido ")
)

type Contato struct {
	ID        entity.ID `json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	OBS       string    `json:"obs"`
	Tipo      int       `json:"tipo"`
	Ativo     bool      `json:"ativo"`
	CreatedAt time.Time `json:"created_at"`
}

func NewContato(nome, email, obs string, tipo int) (*Contato, error) {

	// TODO: Implementar validacao do tipo verificar se o tipo exite
	if tipo == 0 {
		tipo = 1
	}

	contato := &Contato{
		ID:        entity.NewID(),
		Nome:      nome,
		Email:     email,
		OBS:       obs,
		Tipo:      tipo,
		Ativo:     true,
		CreatedAt: time.Now(),
	}

	err := contato.Validar()
	if err != nil {
		return nil, err
	}
	return contato, nil
}

func (c *Contato) Validar() error {

	if c.Nome == "" {
		return ErrNomeIsRequired
	}
	if c.Email == "" {
		return ErrEmailIsRequired
	}
	if !entity.IsEmailValid(c.Email) {
		return ErrEmailIsInvalid
	}

	return nil
}
