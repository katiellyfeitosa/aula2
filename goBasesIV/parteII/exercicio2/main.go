package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type cliente struct {
	Arquivo       int
	nomeSobrenome string
	rg            string
	telefone      string
	endereco      string
}

func gerarNumero() int {
	// Use the Seed function to initialize the default Source if different behavior is required for each run.
	rand.Seed(time.Now().UnixNano())
	arquivo := rand.Intn(1000000)

	if arquivo == 0 {
		return 0
	}

	return arquivo
}

func (c cliente) gerarLinha() string {
	return fmt.Sprintf("%d, %s, %s, %s, %s", c.Arquivo, c.nomeSobrenome, c.rg, c.telefone, c.endereco)
}

func gravarArquivo(caminho string, c cliente) error {
	arquivo := gerarNumero()

	if arquivo == 0 {
		panic("Numero do arquivo nao pode ser nulo.")
	}

	c.Arquivo = arquivo

	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Erro ao abrir o arquivo %s.", err)
	}

	_, err = file.WriteString(c.gerarLinha())
	if err != nil {
		return fmt.Errorf("Erro ao abrir o arquivo %s.", err)
	}

	fmt.Printf("Cliente %s gravado com sucesso!\n", c.nomeSobrenome)
	return nil
}

func lerArquivo(nomeArquivo string) string {
	data, err := os.ReadFile(nomeArquivo)

	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado ")
	}

	fmt.Printf("Arquivo %s lido com sucesso!\n", nomeArquivo)

	return string(data)
}

func (c cliente) verificaSeExiste(data string) bool {
	existe := strings.Contains(data, c.nomeSobrenome)
	return existe
}

func main() {

	cliente1 := cliente{
		nomeSobrenome: "Kiko",
		rg:            "870126574",
		telefone:      "11959535619",
		endereco:      "Rua do me ajuda",
	}

	caminhoDoArquivo := "customers.txt"

	var existe bool

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro:", r)
		}

		if existe == true {
			fmt.Printf("Cliente já cadastrado")
		} else {
			gravarArquivo(caminhoDoArquivo, cliente1)
		}
	}()

	data := lerArquivo(caminhoDoArquivo)
	existe = cliente1.verificaSeExiste(data)
}
