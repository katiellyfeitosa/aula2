package users

import (
	"fmt"

	"github.com/katiellyfeitosa/bootcampGO/goWebIII/exercicio/pkg/store"
)

type User struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Sobrenome string  `json:"sobrenome"`
	Email     string  `json:"email"`
	Idade     string  `json:"idade"`
	Altura    float64 `json:"altura"`
}

type Repository interface {
	LastID() (int, error)
	GetAll() ([]User, error)
	SaveUser(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error)
	Update(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error)
	UpdateLastNameAndAge(id int, sobrenome string, idade string) (User, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]User, error) {
	var listUser []User
	r.db.Read(&listUser)

	return listUser, nil
}

func (r *repository) LastID() (int, error) {
	var listUser []User

	if err := r.db.Read(&listUser); err != nil {
		return 0, err
	}
	if len(listUser) == 0 {
		return 0, nil
	}

	return listUser[len(listUser)-1].ID, nil
}

func (r *repository) SaveUser(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error) {
	var listUser []User

	r.db.Read(&listUser)
	u := User{id, nome, sobrenome, email, idade, altura}
	listUser = append(listUser, u)
	if err := r.db.Write(listUser); err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) Update(id int, nome string, sobrenome string, email string, idade string, altura float64) (User, error) {
	var listUser []User
	r.db.Read(&listUser)

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

	if err := r.db.Write(listUser); err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) UpdateLastNameAndAge(id int, sobrenome string, idade string) (User, error) {
	var listUser []User
	r.db.Read(&listUser)

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

	if err := r.db.Write(listUser); err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) Delete(id int) error {
	var listUser []User
	r.db.Read(&listUser)

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

	if err := r.db.Write(listUser); err != nil {
		return err
	}

	return nil
}
