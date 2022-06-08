package users

import (
	"fmt"
)

type User struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Sobrenome string  `json:"sobrenome"`
	Email     string  `json:"email"`
	Idade     string  `json:"idade"`
	Altura    float64 `json:"altura"`
}

var listUser []User
var lastID int

type Repository interface {
	LastID() (int, error)
	GetAll() ([]User, error)
	SaveUser(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error)
	Update(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error)
	UpdateLastNameAndAge(id int, sobrenome string, idade string) (User, error)
	Delete(id int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return listUser, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) SaveUser(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error) {
	u := User{id, nome, sobrenome, email, idade, altura}
	listUser = append(listUser, u)

	lastID = u.ID

	return u, nil
}

func (r *repository) Update(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error) {
	u := User{id, nome, sobrenome, email, idade, altura}
	updated := false

	for i := range listUser {
		if listUser[i].ID == id {
			u.ID = id
			listUser[i] = u
			updated = true
		}
	}

	if !updated {
		return User{}, fmt.Errorf("Usuario %d não encontrado", id)
	}

	return u, nil
}

func (r *repository) UpdateLastNameAndAge(id int, sobrenome string, idade string) (User, error) {
	var u User
	updated := false

	for i := range listUser {
		if listUser[i].ID == id {
			listUser[i].Sobrenome = sobrenome
			listUser[i].Idade = idade
			updated = true
			u = listUser[i]
		}
	}

	if !updated {
		return User{}, fmt.Errorf("Usuario %d não encontrado", id)
	}

	return u, nil
}

func (r *repository) Delete(id int) error {
	deleted := false

	var index int
	for i := range listUser {
		if listUser[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Usuario %d não encontrado", id)
	}

	listUser = append(listUser[:index], listUser[index+1:]...)
	return nil
}
