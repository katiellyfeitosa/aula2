package main

import "fmt"

type usuario struct {
	Nome      string
	Sobrenome string
	idade     string
	email     string
	senha     string
}

func main() {
	var usuario1 *usuario

	usuario1 = novoUsuario("Katielly", "Rezende", "24", "teste@teste.com", "aaa")

	fmt.Println("Meu usuario", usuario1)

	usuario1.mudarNome("Kati", "Feitosa")
	fmt.Println("Novo nome e sobrenome: ", usuario1.Nome, usuario1.Sobrenome)

	usuario1.mudarIdade("27")
	fmt.Println("Novo idade: ", usuario1.idade)

	usuario1.mudarEmail("kati@teste.com")
	fmt.Println("Novo email: ", usuario1.email)

	usuario1.mudarSenha("01051999")
	fmt.Println("Novo senha: ", usuario1.senha)

	fmt.Println("Meu atualizado", usuario1)
}

func novoUsuario(Nome string, Sobrenome string, idade string, email string, senha string) *usuario {
	return &usuario{
		Nome:      "Katielly",
		Sobrenome: "Rezende",
		idade:     "23",
		email:     "kati@teste.com",
		senha:     "aaaaa",
	}
}

func (u *usuario) mudarNome(novoNome string, novoSobrenome string) {
	if u.Nome != "" {
		u.Nome = novoNome
	}

	if u.Sobrenome != "" {
		u.Sobrenome = novoSobrenome
	}
}

func (u *usuario) mudarIdade(novaIdade string) {
	if u.idade != "" {
		u.idade = novaIdade
	}
}

func (u *usuario) mudarEmail(novoEmail string) {
	if u.email != "" {
		u.email = novoEmail
	}
}

func (u *usuario) mudarSenha(novaSenha string) {
	if u.senha != "" {
		u.senha = novaSenha
	}
}
